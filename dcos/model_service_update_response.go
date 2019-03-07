/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type ServiceUpdateResponse struct {
	MarathonDeploymentId string               `json:"marathonDeploymentId"`
	Package              V50PackageDefinition `json:"package"`
	// The result of merging the default package options with the user supplied options
	ResolvedOptions map[string]map[string]interface{} `json:"resolvedOptions"`
}
