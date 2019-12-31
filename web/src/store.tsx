import React, {useState, useEffect} from 'react'
import {createContainer} from 'unstated-next'
import {AuthenticateUserSignin, coinbaseProducts, coinbaseTicker, selectedProductID, FavToggle, UserFavList, CurrentUserID, AuthenticateUserSignup, ResetPasword, ForgotPasword} from "./struct"

export const useStore = () => {

//setup state 
const [enteredUsername, setEnteredUsername]=useState<string>("")
const [enteredEmail, setEnteredEmail] = useState<string>("")
const [enteredPassword, setEnteredPassword] = useState<string>("")
const [isSubmit, setIsSubmit] = useState<boolean>(false)
const [isLogin, setIsLogin] = useState<boolean>(false)
// const [isSignup, setIsSignup] = useState<boolean>(false)
const [signupVerification, setSignupVerification] = useState<boolean>(false)
const [isError, setIsError] = useState<boolean>(false)
const [errorMsg, setErrorMsg] = useState<string>("")
const [username, setUsername] = useState<string>("")
const [ticker, setTicker] = useState<coinbaseTicker | undefined> (undefined) 
const [isPopUp, setPopUp] = useState<boolean>(false)
const [productList, setProductList] = React.useState<coinbaseProducts[] | undefined>(undefined)
const [searchKey, setSearchKey] = useState<string>("")
const [searchResult, setSearchResult] = useState<coinbaseProducts[] | undefined>(undefined)
const [favList, setFavList] = useState<FavToggle [] | undefined>(undefined)
const [currentUser, setCurrentUser] = useState<string>("")
const [userFavList, setUserFavList] = useState<UserFavList [] | undefined>(undefined)
const [openLogoutMsg, setOpenLogoutMsg] = useState<boolean>(false)
const [enteredCurrentPw, setEnteredCurrentPw] = useState<string>("")
const [enteredNewPw, setEnteredNewPw] = useState<string>("")
const [successMsg, setSuccesMsg] = useState<string>("")
const [verifiedEmail, setVerifiedEmail] = useState<boolean>(false) 

useEffect(() => {
    fetchDataFromAPI("/api/auth", "readCookie")
},[])

//reset Login & Signup form input
const ResetFormInput = () => {
    setIsSubmit(false)
    setErrorMsg("")
    setSuccesMsg("")
    setEnteredUsername("")
    setEnteredEmail("")
    setEnteredPassword("")
    setEnteredCurrentPw("")
    setEnteredNewPw("")
    setVerifiedEmail(false)
}

//reset Reset Password Form input 
const ResetResetPwInput = () => {
    setIsSubmit(false)
    setIsError(false)
    setEnteredCurrentPw("")
    setEnteredNewPw("")
    setErrorMsg("")
    setSuccesMsg("")
}

//reset Forgot Password Form input 
const ResetForgotPassInput = () => {
    setVerifiedEmail(false)
    setIsSubmit(false)
    setIsError(false)
    setEnteredEmail("")    
}

const loginValidation = () => {
    if(enteredEmail === ""){
        return false 
    } else if (enteredPassword === ""){
        return false 
    }
    return true 
}

const resetPwValidation = () => {
    if ( enteredCurrentPw === "" ||enteredNewPw === "" ){
        return false 
    } 
    return true 
}

const signupValidation = () => {
    if(enteredUsername === ""){
        return false 
    } else if (enteredEmail === ""){
        return false 
    } else if (enteredPassword === ""){
        return false 
    }
    return true 
}

//set onChange method for email input 
const handleEnteredEmail = (event: React.FormEvent<HTMLInputElement>) => {
    const nonSpaceValue: string = event.currentTarget.value.replace(/\s+/g, "")
    setEnteredEmail(nonSpaceValue)
} 

const handleEnteredPassword = (event:React.FormEvent<HTMLInputElement>) => setEnteredPassword(event.currentTarget.value)

const handleEnteredUsername = (event:React.FormEvent<HTMLInputElement>) => setEnteredUsername(event.currentTarget.value)

async function postData(url: string, body:any, tag: string){
    const response = await fetch (url, {
        method: "POST",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
    })
    const json: any = await response.json()
//    console.log(json.error_msg)

    switch(tag){
        case "login":
            setIsError(!json.is_login)
            setErrorMsg(handleErrorMsg(json.error_msg))
            setIsLogin(json.is_login)
            
            if(json.is_login){
                // const user: string = json.id
                currentUserFavList()
            }   
            break 

        case "ticker":      
            setTicker(json)
            console.log(json)
            break

        case "favouriteToggle":
            setUserFavList(json)
            break

        case "favouriteList":
            setUserFavList(json)
            break

        case "signup": 
        console.log("signup result:", json.is_signup)
            if(!json.is_signup){
                setIsError(!json.is_signup)
                setErrorMsg(handleErrorMsg(json.error_msg))
            }
            else if(json.is_signup === true && json.is_verified === false){
                setSignupVerification(true)
                setSuccesMsg("A verification link has been sent to your email. Please check your email and confirm your email address.")
            }
            break

        case "resetPassword":
            setIsError(!json.success)
            if(json.success){
                setEnteredCurrentPw("")
                setEnteredNewPw("")
                setSuccesMsg("Your password has been reset")
            } else{
                 setIsError(!json.success)
                 setErrorMsg(handleErrorMsg(json.error_msg))
            }
            break

        case "forgotPassword":
            console.log(json.error_msg)
            if (json.error_msg !== ""){
                setIsError(true)
                setErrorMsg(handleErrorMsg(json.error_msg))    
            } else {
                setVerifiedEmail(true)
                setIsError(false)
            }
            break
    }
}

const handleLogin = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setErrorMsg("")
    setIsSubmit(true)
    event.preventDefault()
    if(loginValidation()){
        const authenticateUserSignin: AuthenticateUserSignin = {email: enteredEmail, password: enteredPassword, isSubmit: isSubmit}
        postData("/api/login", authenticateUserSignin, "login")
    } else{
        setIsError(true)
    } 
}

const handleSignup = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsSubmit(true)
    event.preventDefault() 
    if(signupValidation()){
        const authenticateUserSignup: AuthenticateUserSignup = {username: enteredUsername, email: enteredEmail, password: enteredPassword}
        postData("/api/signup", authenticateUserSignup, "signup")
    } else {
        setIsError(true)
    }
    ResetFormInput()
}

const handleSelectedProduct = (id:string) => {
    const selectedProductID: selectedProductID = {ticker_id: id}
    postData("/api/ticker", selectedProductID, "ticker")  
}

const handleErrorMsg = (errorFlag: string) => {
    switch(errorFlag){
        case "Invalid_Email": 
            return "Invalid email"
        case "Incorrect_Password":
            return "Incorrect password"
        case "Add_Fav_Product_Error":
            return "Error occured when adding product to favourites"
        case "Remove_Fav_Product_Error":
            return "Error occured when removing product from favourites"
        case "Empty_Fav_Product_List":
            return "No favourites"
        case "Violate_UN_Username":
            return "Username has already been used"
        case "Violate_UN_Email":
            return "Email address has already been used"
        case "Invalid_Email_Or_Password":
            return "Invalid email and/or password"
        case "Incorrect_New_Password_Format":
            return "Invalid format for new password"
        case "Password_Matching_Issue":
            return "Your current password does not match. Please try again"
        case "Weak_Password":
            return "Weak password"
        case "Reset_Password_Error":
            return "Unexpected error occured on reset password"
        case "Same_Reset_Password_Input":
            return "Your new password cannot be the same as current password"
        case "Request_Reset_Pass_Token_Error":
            return "There's no email address found in our database. Please try again."
        case "DB_Query_Error":
            return "Requested data does not exist"
        case "Signup_Error":
            return "Error occured on signup"
        case "User_Verification_Error":
            return "Error occured on verifying your account"
        case "JSON_Parse_Error":
            return "Error occured when sending request to server"
        case "Add_Product_Error":
            return "Error occured on listing all products"
        // case "Update_Tickers_Error":
            // return "Error occured on updating tickers"
        default: 
        return "An Unexpected Error Occured"
    }
}

// const useFetchProducts = (url: string, options = {}) => {
//     const [resp, setResp] = React.useState()
//     const [err, setErr] = React.useState()
//     React.useEffect(() => {
//         const fetchData = async () => {
//             try {
//                 const res = await fetch(url, options)
//                 const json = await res.json()
//                 setResp(json)
//                 // console.log(json)
//                 setProductList(json)
//                 console.log("set product list in store:", json)
//                 setSearchResult(json)
//             } catch (err) {
//                 setErr(err)
//             }
//         }
//         fetchData()
//     }, [])
//     return { resp, err }
// }

const fetchDataFromAPI =(url:string, tag:string)=> {
    const fetchData = async () => {
        try {
            // fetchData
            const res = await fetch(url)
            const json = await res.json()
            console.log(json)
            
            switch(tag){
                case "product":
                        setProductList(json)
                        setSearchResult(json)
                        console.log("fetch data from api: ", productList)
                        break

                case "readCookie":
                    setIsLogin(json.checked_cookie)
                    console.log("checked cookie: ", json.checked_cookie)
                    break 

                case "favouriteList":
                    setUserFavList(json)
                    console.log("fav list:", json)
                    break

                case "logout":
                    if(json.success){
                        setOpenLogoutMsg(false)
                        ResetFormInput()
                        ResetResetPwInput()
                        setIsLogin(false)
                    }
                    break

                default:
                    break
            }
        } catch (err) {
            console.log(err)
        }
    }
    fetchData()
}

const handleSearch = (event: React.FormEvent<HTMLInputElement>) => {
   const input: string = event.currentTarget.value
    setSearchKey(input)    
    setSearchResult(productList?productList.filter((result: coinbaseProducts)=> result.ID.toLowerCase().includes(input.toLowerCase())):undefined)    
}

const handleFavIcon=(productID: string)=>{
    const isfave =userFavList?userFavList.find(fav=>fav.ID===productID):undefined
    console.log(isfave)
    return isfave
}

const handleFavourite = (productID: string, userID: string) => {
        const favToggle: FavToggle = {product_id: productID}
        postData("/api/fav-toggle", favToggle, "favouriteToggle")
        currentUserFavList()
}

    const currentUserFavList = () => {
    fetchDataFromAPI("/api/fav-list", "favouriteList")
}

const handleLogoutMsg = () => {
    setOpenLogoutMsg(true)
}

const CancelLogout = () => {
    console.log("cancel logout")
    setOpenLogoutMsg(false)
}

const ConfirmLogout = () => {
    console.log("confirm logout")
    fetchDataFromAPI("/api/logout", "logout")
    
}

const handleEnteredCurrentPw = (event:React.FormEvent<HTMLInputElement>) => setEnteredCurrentPw (event.currentTarget.value)
const handleEnteredNewPw = (event:React.FormEvent<HTMLInputElement>) => setEnteredNewPw(event.currentTarget.value)

const handleResetPassword = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsSubmit(true)
    event.preventDefault()
    if(resetPwValidation()){
        const resetPassword: ResetPasword={current_password: enteredCurrentPw, new_password:enteredNewPw}
        postData("/api/reset-password", resetPassword, "resetPassword") 
    } else {
        setIsError(true)
    }
}   

const HandleForgotPassword = () => {
    setIsSubmit(true)
    console.log("forgot pass email:", enteredEmail)
    if (enteredEmail !== ""){
        const forgotPassword : ForgotPasword={email: enteredEmail}
        postData("/api/forgot-password", forgotPassword, "forgotPassword")
    } else {
        setIsError(true)
    }
}

return {
    setIsSubmit,
    isSubmit,
    isLogin,
    signupVerification,
    enteredUsername,
    enteredEmail,
    enteredPassword,
    handleEnteredUsername,
    handleEnteredEmail,
    handleEnteredPassword,
    handleLogin,
    handleSignup,
    isError,
    errorMsg,
    productList,
    // useFetchProducts,
    handleSelectedProduct,
    ticker,
    setTicker,
    isPopUp,
    setPopUp,
    handleSearch,
    searchKey,
    searchResult,
    handleFavourite,
    setCurrentUser,
    currentUser,
    handleFavIcon,
    userFavList,
    CancelLogout,
    openLogoutMsg,
    setOpenLogoutMsg,
    ConfirmLogout,
    handleLogoutMsg,
    handleResetPassword,
    fetchDataFromAPI,
    enteredCurrentPw,
    enteredNewPw,
    handleEnteredCurrentPw,
    handleEnteredNewPw,
    ResetFormInput,
    successMsg,
    ResetResetPwInput,
    ResetForgotPassInput,
    HandleForgotPassword,
    verifiedEmail,
}
}

export const StoreContainer = createContainer(useStore)