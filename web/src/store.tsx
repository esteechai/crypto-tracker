import React, {useState} from 'react'
import {createContainer} from 'unstated-next'
import { create } from 'domain'
import {AuthenticateData, coinbaseProducts, coinbaseTicker, selectedProductID} from "./struct"
import { string } from 'prop-types'

export const useStore = () => {
const [isLogin, setIsLogin] = useState<boolean>(false)
const [isSignUp, setSignUp] = useState<boolean>(false)
const [isError, setIsError] = useState<boolean>(false)
const [errorMsg, setErrorMsg] = useState<string>("")
const [username, setUsername] = useState<string>("")
const [ticker, setTicker] = useState<coinbaseTicker | undefined> (undefined) 
const [isPopUp, setPopUp] = useState<boolean>(false)

const loginValidation = () => {
    if(enteredEmail === ""){
        return false 
    } else if (enteredPassword === ""){
        return false 
    }
    return true 
}

//setup state 
const [enteredEmail, setEnteredEmail] = useState<string>("")
const [enteredPassword, setEnteredPassword] = useState<string>("")
const [isSubmit, setIsSubmit] = useState<boolean>(false)

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
        setIsError(!json.success)
        setErrorMsg(handleErrorMsg(json.error_msg))
        setIsLogin(json.success)
        if(json.success){
            console.log("sign in success")
        }
        case "ticker":
          
            setTicker(json)
            console.log(json)
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
    setPopUp(true)
    const selectedProductID: selectedProductID = {ticker_id: id}
    console.log("store.tsx:", id)
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


const [productList, setProductList] = React.useState<coinbaseProducts[] | undefined>(undefined)


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

const [searchKey, setSearchKey] = useState<string>("")
const [searchResult, setSearchResult] = useState<coinbaseProducts[] | undefined>(undefined)

const handleSearch = (event: React.FormEvent<HTMLInputElement>) => {
   const input: string = event.currentTarget.value
    setSearchKey(input)    
    setSearchResult(productList?productList.filter((result: coinbaseProducts)=> result.id.toLowerCase().includes(input.toLowerCase())):undefined)
    // setSearchResult(productList)

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
    searchResult
}

}

export const StoreContainer = createContainer(useStore)