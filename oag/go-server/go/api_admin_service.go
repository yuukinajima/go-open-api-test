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

// AdminApiService is a service that implents the logic for the AdminApiServicer
// This service should implement the business logic for every endpoint for the AdminApi API. 
// Include any external packages or services that will be required by this service.
type AdminApiService struct {
}

// NewAdminApiService creates a default api service
func NewAdminApiService() AdminApiServicer {
	return &AdminApiService{}
}

// DeleteAuthersId - 
func (s *AdminApiService) DeleteAuthersId(id string) (interface{}, error) {
	// TODO - update DeleteAuthersId with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteAuthersId' not implemented")
}

// DeleteBooksId - 
func (s *AdminApiService) DeleteBooksId(id string) (interface{}, error) {
	// TODO - update DeleteBooksId with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteBooksId' not implemented")
}

// DeleteCategoriesId - 
func (s *AdminApiService) DeleteCategoriesId(id string) (interface{}, error) {
	// TODO - update DeleteCategoriesId with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteCategoriesId' not implemented")
}

// GetBooksId - Your GET endpoint
func (s *AdminApiService) GetBooksId(id string) (interface{}, error) {
	// TODO - update GetBooksId with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetBooksId' not implemented")
}

// PatchAuthersId - 
func (s *AdminApiService) PatchAuthersId(id string, auther Auther) (interface{}, error) {
	// TODO - update PatchAuthersId with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PatchAuthersId' not implemented")
}

// PatchBooksId - 
func (s *AdminApiService) PatchBooksId(id string, book Book) (interface{}, error) {
	// TODO - update PatchBooksId with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PatchBooksId' not implemented")
}

// PatchCategoriesId - 
func (s *AdminApiService) PatchCategoriesId(id string, category Category) (interface{}, error) {
	// TODO - update PatchCategoriesId with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PatchCategoriesId' not implemented")
}

// PostAuthers - 
func (s *AdminApiService) PostAuthers(auther Auther) (interface{}, error) {
	// TODO - update PostAuthers with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PostAuthers' not implemented")
}

// PostBooks - 
func (s *AdminApiService) PostBooks(book Book) (interface{}, error) {
	// TODO - update PostBooks with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PostBooks' not implemented")
}

// PostCategories - 
func (s *AdminApiService) PostCategories(category Category) (interface{}, error) {
	// TODO - update PostCategories with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PostCategories' not implemented")
}
