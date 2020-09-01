/*
 * test-api dev
 *
 * test api 
 *
 * API version: 1.0
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

	AdminApiService := openapi.NewAdminApiService()
	AdminApiController := openapi.NewAdminApiController(AdminApiService)

	AllApiService := openapi.NewAllApiService()
	AllApiController := openapi.NewAllApiController(AllApiService)

	router := openapi.NewRouter(AdminApiController, AllApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
