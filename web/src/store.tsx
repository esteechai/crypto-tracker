import React, {useState} from 'react'
import {createContainer} from 'unstated-next'
import { create } from 'domain'
import {AuthenticateUserSignin, coinbaseProducts, coinbaseTicker, selectedProductID, FavToggle, UserFavList, CurrentUserID, AuthenticateUserSignup, ResetPasword} from "./struct"
import { string } from 'prop-types'
import { Redirect } from 'react-router'
import ResetPassword from './components/ResetPassword'

export const useStore = () => {

//setup state 
const [enteredUsername, setEnteredUsername]=useState<string>("")
const [enteredEmail, setEnteredEmail] = useState<string>("")
const [enteredPassword, setEnteredPassword] = useState<string>("")
const [isSubmit, setIsSubmit] = useState<boolean>(false)
const [isLogin, setIsLogin] = useState<boolean>(false)
const [isSignUp, setSignUp] = useState<boolean>(false)
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
// const [navBarActiveItem, setNavBarActiveItem] = useState<string>("")
const [enteredCurrentPw, setEnteredCurrenPw] = useState<string>("")
const [enteredNewPw, setEnteredNewPw] = useState<string>("")

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
                const user: string = json.id
                currentUserFavList(user)
            }   
            break 
        case "ticker":      
            setTicker(json)
            console.log(json)
            break
        case "favouriteToggle":
            setUserFavList(json)
            // const newlist =favList?favList.concat(body):[body]
            break
        case "favouriteList":
            setUserFavList(json)
            // console.log("current user fav list: ", json)
            break
        case "signup": 
            setIsError(!json.is_signup)
            setErrorMsg(handleErrorMsg(json.error_msg))
            break
        case "resetPassword":
            setIsError(!json.success)
            if(json.success){
                setEnteredCurrenPw("")
                setEnteredNewPw("")
            }
            setErrorMsg(handleErrorMsg(json.error_msg))

            break
    }
}

const handleLogin = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsSubmit(true)
    event.preventDefault()
    if(loginValidation()){
        const authenticateUserSignin: AuthenticateUserSignin = {email: enteredEmail, password: enteredPassword, isSubmit: isSubmit}
        postData("http://localhost:8080/api/login", authenticateUserSignin, "login")
    } else{
        setIsError(true)
        // setErrorMsg("Invalid Email or Password")
    } 
}

const handleSignup = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsSubmit(true)
    event.preventDefault() 
    if(signupValidation()){
        const authenticateUserSignup: AuthenticateUserSignup = {username: enteredUsername, email: enteredEmail, password: enteredPassword}
        postData("http://localhost:8080/api/signup", authenticateUserSignup, "signup")
    } else {
        setIsError(true)
    }
}

const handleSelectedProduct = (id:string) => {
    const selectedProductID: selectedProductID = {ticker_id: id}
    postData("http://localhost:8080/api/ticker", selectedProductID, "ticker")  
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
        default:
        return "An Unexpected Error Occured"
    }
}

const useFetchProducts = (url: string, options = {}) => {
    const [resp, setResp] = React.useState()
    const [err, setErr] = React.useState()
    React.useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await fetch(url, options)
                const json = await res.json()
                setResp(json)
                console.log(json)
                setProductList(json)
                setSearchResult(json)
            } catch (err) {
                setErr(err)
            }
        }
        fetchData()
    }, [])
    return { resp, err }
}

const fetchDataFromAPI =(url:string, tag:string)=> {
    const fetchData = async () => {
        try {
            const res = await fetch(url)
            const json = await res.json()
            console.log(json)
            
            switch(tag){
                case "product":
                        setProductList(json)
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
    setSearchResult(productList?productList.filter((result: coinbaseProducts)=> result.id.toLowerCase().includes(input.toLowerCase())):undefined)    
}

const handleFavIcon=(productID: string)=>{
    const isfave =userFavList?userFavList.find(fav=>fav.ID===productID):undefined
    console.log(isfave)
    return isfave
}

const handleFavourite = (productID: string, userID: string) => {
        const favToggle: FavToggle = {product_id: productID, user_id: userID}
        postData("http://localhost:8080/api/fav-toggle", favToggle, "favouriteToggle")
        currentUserFavList(favToggle.user_id)
}

const currentUserFavList = (id: string) =>{
    const currentUserID: CurrentUserID = {user_id: id}
    postData("http://localhost:8080/api/fav-list", currentUserID, "favouriteList")
    setCurrentUser(id)
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
    setOpenLogoutMsg(false)
    setIsLogin(false)
}

const handleEnteredCurrentPw = (event:React.FormEvent<HTMLInputElement>) => setEnteredCurrenPw (event.currentTarget.value)
const handleEnteredNewPw = (event:React.FormEvent<HTMLInputElement>) => setEnteredNewPw(event.currentTarget.value)

const handleResetPassword = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsSubmit(true)
    event.preventDefault()
    if(resetPwValidation()){
        const resetPassword: ResetPasword={user_id: currentUser, current_password: enteredCurrentPw, new_password:enteredNewPw}
        postData("http://localhost:8080/api/reset-password", resetPassword, "resetPassword")
    } else {
        setIsError(true)
    }
}   

return {
    setIsSubmit,
    isSubmit,
    isLogin,
    isSignUp,
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
    useFetchProducts,
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
    enteredCurrentPw,
    enteredNewPw,
    handleEnteredCurrentPw,
    handleEnteredNewPw    
}
}

export const StoreContainer = createContainer(useStore)