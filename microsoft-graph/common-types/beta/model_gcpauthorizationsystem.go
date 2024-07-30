package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpAuthorizationSystem struct {
	Actions                 *[]GcpAuthorizationSystemTypeAction `json:"actions,omitempty"`
	AssociatedIdentities    *GcpAssociatedIdentities            `json:"associatedIdentities,omitempty"`
	AuthorizationSystemId   *string                             `json:"authorizationSystemId,omitempty"`
	AuthorizationSystemName *string                             `json:"authorizationSystemName,omitempty"`
	AuthorizationSystemType *string                             `json:"authorizationSystemType,omitempty"`
	DataCollectionInfo      *DataCollectionInfo                 `json:"dataCollectionInfo,omitempty"`
	Id                      *string                             `json:"id,omitempty"`
	ODataType               *string                             `json:"@odata.type,omitempty"`
	Resources               *[]GcpAuthorizationSystemResource   `json:"resources,omitempty"`
	Roles                   *[]GcpRole                          `json:"roles,omitempty"`
	Services                *[]AuthorizationSystemTypeService   `json:"services,omitempty"`
}
