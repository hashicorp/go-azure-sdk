package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceEnvironment struct {
	AccessPackageResources *[]AccessPackageResource `json:"accessPackageResources,omitempty"`
	ConnectionInfo         *ConnectionInfo          `json:"connectionInfo,omitempty"`
	CreatedBy              *string                  `json:"createdBy,omitempty"`
	CreatedDateTime        *string                  `json:"createdDateTime,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	DisplayName            *string                  `json:"displayName,omitempty"`
	Id                     *string                  `json:"id,omitempty"`
	IsDefaultEnvironment   *bool                    `json:"isDefaultEnvironment,omitempty"`
	ModifiedBy             *string                  `json:"modifiedBy,omitempty"`
	ModifiedDateTime       *string                  `json:"modifiedDateTime,omitempty"`
	ODataType              *string                  `json:"@odata.type,omitempty"`
	OriginId               *string                  `json:"originId,omitempty"`
	OriginSystem           *string                  `json:"originSystem,omitempty"`
}
