/*
 * sme API
 *
 * this is sme Api doc
 *
 * API version: 2.0.0+2024-02-15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dto

type SignUpUserDtoCredentialData interface {
	any | SignUpUserDtoUserNamePasswordCredentialData
}

type SignUpUserDtoUserNamePasswordCredentialData struct {
	Type string `json:"type,omitempty"`
	// User name
	Username string `json:"username,omitempty"`
	// User password
	Password string `json:"password,omitempty"`
}
