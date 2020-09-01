/*
 * test-api dev
 *
 * test api 
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// AllApiService is a service that implents the logic for the AllApiServicer
// This service should implement the business logic for every endpoint for the AllApi API. 
// Include any external packages or services that will be required by this service.
type AllApiService struct {
}

// NewAllApiService creates a default api service
func NewAllApiService() AllApiServicer {
	return &AllApiService{}
}

// GetAuthers - Your GET endpoint
func (s *AllApiService) GetAuthers() (interface{}, error) {
	// TODO - update GetAuthers with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetAuthers' not implemented")
}

// GetAuthersId - Your GET endpoint
func (s *AllApiService) GetAuthersId(id string) (interface{}, error) {
	// TODO - update GetAuthersId with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetAuthersId' not implemented")
}

// GetBooks - Your GET endpoint
func (s *AllApiService) GetBooks() (interface{}, error) {
	// TODO - update GetBooks with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetBooks' not implemented")
}

// GetBooksId - Your GET endpoint
func (s *AllApiService) GetBooksId(id string) (interface{}, error) {
	// TODO - update GetBooksId with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetBooksId' not implemented")
}

// GetCategories - Your GET endpoint
func (s *AllApiService) GetCategories() (interface{}, error) {
	// TODO - update GetCategories with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetCategories' not implemented")
}

// GetCategoriesId - Your GET endpoint
func (s *AllApiService) GetCategoriesId(id string) (interface{}, error) {
	// TODO - update GetCategoriesId with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetCategoriesId' not implemented")
}

// PostAuthers - 
func (s *AllApiService) PostAuthers(auther Auther) (interface{}, error) {
	// TODO - update PostAuthers with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PostAuthers' not implemented")
}

// PostBooks - 
func (s *AllApiService) PostBooks(book Book) (interface{}, error) {
	// TODO - update PostBooks with the required logic for this service method.
	// Add api_all_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PostBooks' not implemented")
}
