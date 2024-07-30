package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAuthorizationSystem struct {
	Actions                 *[]AzureAuthorizationSystemTypeAction `json:"actions,omitempty"`
	AssociatedIdentities    *AzureAssociatedIdentities            `json:"associatedIdentities,omitempty"`
	AuthorizationSystemId   *string                               `json:"authorizationSystemId,omitempty"`
	AuthorizationSystemName *string                               `json:"authorizationSystemName,omitempty"`
	AuthorizationSystemType *string                               `json:"authorizationSystemType,omitempty"`
	DataCollectionInfo      *DataCollectionInfo                   `json:"dataCollectionInfo,omitempty"`
	Id                      *string                               `json:"id,omitempty"`
	ODataType               *string                               `json:"@odata.type,omitempty"`
	Resources               *[]AzureAuthorizationSystemResource   `json:"resources,omitempty"`
	RoleDefinitions         *[]AzureRoleDefinition                `json:"roleDefinitions,omitempty"`
	Services                *[]AuthorizationSystemTypeService     `json:"services,omitempty"`
}
