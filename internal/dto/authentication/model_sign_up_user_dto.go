/*
 * sme API
 *
 * this is sme Api doc
 *
 * API version: 2.0.0+2024-02-15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dto

type Source string

const (
	SourceWebApp   Source = "WEBAPP"
	SourceWhatsApp Source = "WHATSAPP"
)

type SignUpUserDto[T SignUpUserDtoCredentialData] struct {
	// The user source
	Source Source `json:"source,omitempty"`

	// The user permission
	Permissions []string `json:"permissions,omitempty"`

	CredentialData T `json:"credentialData,omitempty"`
}
