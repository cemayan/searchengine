// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package types

// Defines values for SEErrorKind.
const (
	Db        SEErrorKind = "db"
	Messaging SEErrorKind = "messaging"
)

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Code *int                    `json:"code,omitempty"`
	Data *map[string]interface{} `json:"data,omitempty"`
	Msg  *string                 `json:"msg,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Msg string `json:"msg"`
}

// Items defines model for Items.
type Items struct {
	Title *string `json:"title,omitempty"`
	Url   *string `json:"url,omitempty"`
}

// RecordRequest defines model for RecordRequest.
type RecordRequest struct {
	Data string `json:"data"`
}

// ResultsResponse defines model for ResultsResponse.
type ResultsResponse struct {
	Items *Items `json:"items,omitempty"`
}

// SEError defines model for SEError.
type SEError struct {
	DbName string      `json:"dbName,omitempty"`
	Error  string      `json:"error,omitempty"`
	Key    string      `json:"key,omitempty"`
	Kind   SEErrorKind `json:"kind,omitempty"`
	Value  string      `json:"value,omitempty"`
}

// SEErrorKind defines model for SEError.Kind.
type SEErrorKind string

// SearchResponse defines model for SearchResponse.
type SearchResponse = []string

// SelectionRequest defines model for SelectionRequest.
type SelectionRequest struct {
	Query       string `json:"query"`
	SelectedKey string `json:"selectedKey"`
}
