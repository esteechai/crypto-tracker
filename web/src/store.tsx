import React, {useState} from 'react'
import {createContainer} from 'unstated-next'
import { create } from 'domain'
import {authenticateData, coinbaseProducts, coinbaseTicker} from "./struct"
import { string } from 'prop-types'

export const useStore = () => {
const [isLogin, setIsLogin] = useState<boolean>(false)
const [isSignUp, setSignUp] = useState<boolean>(false)
const [isError, setIsError] = useState<boolean>(false)
const [errorMsg, setErrorMsg] = useState<string>("")
const [username, setUsername] = useState<string>("")

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
    console.log(json.error_msg)

    switch(tag){
        case "login":
        setIsError(!json.success)
        setErrorMsg(handleErrorMsg(json.error_msg))
        setIsLogin(json.success)
        if(json.success){
            console.log("sign in success")
        }
    }
}

const handleLogin = (event:React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsSubmit(true)
    event.preventDefault()
    if(loginValidation()){
        const authenticateData: authenticateData = {email: enteredEmail, password: enteredPassword, isSubmit: isSubmit}
        postData("http://localhost:8080/api/login", authenticateData, "login")
    } else{
        setIsError(true)
        setErrorMsg("Invalid Email or Password")
    }
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
const [productTicker, setProductTicker] = React.useState<coinbaseTicker[] | undefined>(undefined)

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
                case "ticker":
                        setProductTicker(json, id:string)
                default:
                    break
            }
        } catch (err) {
            console.log(err)
        }
    }
    fetchData()
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
    productTicker
}

}

export const StoreContainer = createContainer(useStore)