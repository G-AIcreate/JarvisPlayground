/*
 * Swagger Gjarvis - OpenAPI 3.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	GjarvisAPIService := openapi.NewGjarvisAPIService()
	GjarvisAPIController := openapi.NewGjarvisAPIController(GjarvisAPIService)

	router := openapi.NewRouter(GjarvisAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
