import React, {useState} from 'react'
import {createContainer} from 'unstated-next'
import { create } from 'domain'
import {AuthenticateData, coinbaseProducts, coinbaseTicker, selectedProductID, FavToggle, favTicker} from "./struct"
import { string } from 'prop-types'

export const useStore = () => {

//setup state 
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
const [isFavourite, setIsFavourite] = useState<boolean>(false)
const [favList, setFavList] = useState<FavToggle [] | undefined>(undefined)
const [currentUser, setCurrentUser] = useState<string>("")
const [list, setList] = useState<favTicker [] | undefined>(undefined)

const loginValidation = () => {
    if(enteredEmail === ""){
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
   // console.log(json.error_msg)

    switch(tag){
        case "login":
        setIsError(!json.is_login)
        setErrorMsg(handleErrorMsg(json.error_msg))
        setIsLogin(json.is_login)
        if(json.is_login){
            setCurrentUser(json.id)
            console.log("current user: ", currentUser)
        }   
        break 
        case "ticker":      
            setTicker(json)
            console.log(json)
            break
        case "favouriteToggle":
            console.log("fav toggle: ", json)
            setList(json)
            break
    }
}

const handleLogin = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsSubmit(true)
    event.preventDefault()
    if(loginValidation()){
        const authenticateData: AuthenticateData = {email: enteredEmail, password: enteredPassword, isSubmit: isSubmit}
        postData("http://localhost:8080/api/login", authenticateData, "login")
    } else{
        setIsError(true)
        setErrorMsg("Invalid Email or Password")
    }
}

const handleSelectedProduct = (id:string) => {
    const selectedProductID: selectedProductID = {ticker_id: id}
    postData("http://localhost:8080/api/ticker", selectedProductID, "ticker")  
}

const handleErrorMsg = (errorFlag: string) => {
    switch(errorFlag){
        case "Invalid_Email": 
            return "Incorrect email"
        case "Invalid_Password":
            return "Incorrect password"
        case "Incorrect_Password_Format":
            return "Incorrect password format"
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

const handleFavourite = (productID: string, userID: string) => {
    setIsFavourite(!isFavourite)
    const favToggle: FavToggle = {product_id: productID, user_id: userID,is_fav: !isFavourite}
    console.log("selected product:",favToggle)
    postData("http://localhost:8080/api/favToggle", favToggle, "favouriteToggle")  
}


return {
    isSubmit,
    isLogin,
    isSignUp,
    enteredEmail,
    enteredPassword,
    handleEnteredEmail,
    handleEnteredPassword,
    handleLogin,
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
    isFavourite,
    currentUser,
    setList,
    list
}
}

export const StoreContainer = createContainer(useStore)