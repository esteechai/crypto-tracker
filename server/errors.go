package main

import "errors"

var BadRequestError = errors.New("Bad_Request_Error")

var InvalidEmail = errors.New("Invalid_Email")

var IncorrectPasswordFormat = errors.New("Incorrect_Password_Format")

var EmptyRows = errors.New("Empty_Rows")

var ViolateUNEmail = errors.New("Violate_UN_Email")

var ViolateUNUsername = errors.New("Violate_UN_Username")
