/*
file: goapi/cmd/application/responses/structures.go
description: holds structures that are used to create HTTP responses
Author: david@gofiliate.com
Modified by: david:gofiliate.com
Last Modified: 2022-03-10
© 2021 Gotech Solutions Malta Limited (C72157), Trading as Gofiliate
*/
package responses

type ErrorResponse struct {
	Error_code    int    `json:"error_code"`
	Error_message string `json:"error_message"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PayloadResponse struct {
	Response_code int                    `json:"response_code"`
	Payload       map[string]interface{} `json:"payload"`
}
