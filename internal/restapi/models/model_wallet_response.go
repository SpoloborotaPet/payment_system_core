/*
 * Payment System Rest API
 *
 * Payment System Rest API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type WalletResponse struct {
	Status ResponseStatus `json:"status,omitempty"`

	Id uint `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Balance uint `json:"balance,omitempty"`
}
