/*
 * 二手交易平台
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type User struct {

	UserID int32 `json:"userID"`

	UserName string `json:"userName"`

	Password string `json:"password"`

	SchoolID int32 `json:"schoolID"`

	Picture string `json:"picture,omitempty"`

	Tel string `json:"tel,omitempty"`

	Mail string `json:"mail"`

	Gender int32 `json:"gender,omitempty"`

	Status int32 `json:"status"`
}
