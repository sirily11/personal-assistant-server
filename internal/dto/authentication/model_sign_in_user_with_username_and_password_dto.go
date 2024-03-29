/*
 * sme API
 *
 * this is sme Api doc
 *
 * API version: 2.0.0+2024-02-15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dto

// SignInUserWithUsernameAndPasswordDto - User sign in with username and password
type SignInUserWithUsernameAndPasswordDto struct {

	// The user source
	Type string `json:"type,omitempty"`

	// User name
	Username string `json:"username,omitempty"`

	// User password
	Password string `json:"password,omitempty"`
}
