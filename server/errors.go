package main

import "errors"

var BadRequestError = errors.New("Bad_Request_Error")

var InvalidEmailOrPassword = errors.New("Invalid_Email_Or_Password")

var InvalidEmail = errors.New("Invalid_Email")

var IncorrectPassword = errors.New("Incorrect_Password")

var EmptyRows = errors.New("Empty_Rows")

var ViolateUNEmail = errors.New("Violate_UN_Email")

var ViolateUNUsername = errors.New("Violate_UN_Username")

var AddFavProductError = errors.New("Add_Fav_Product_Error")

var RemoveFavProductError = errors.New("Remove_Fav_Product_Error")

var EmptyFavProductList = errors.New("Empty_Fav_Product_List")

var VerifyEmailError = errors.New("Verify_Email_Error")

var IncorrectNewPasswordFormat = errors.New("Incorrect_New_Password_Format")

var PasswordMatchingIssue = errors.New("Password_Matching_Issue")

var WeakPassword = errors.New("Weak_Password")

var ResetPasswordError = errors.New("Reset_Password_Error")

var SameResetPwInput = errors.New("Same_Reset_Password_Input")

var RequestResetPassTokenError = errors.New("Request_Reset_Pass_Token_Error")

var DbQueryError = errors.New("DB_Query_Error")
