package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegration struct {
	ApiVersion           *int64                                 `json:"apiVersion,omitempty"`
	CreatedDateTime      *string                                `json:"createdDateTime,omitempty"`
	DisplayName          *string                                `json:"displayName,omitempty"`
	Encryption           *WorkforceIntegrationEncryption        `json:"encryption,omitempty"`
	Id                   *string                                `json:"id,omitempty"`
	IsActive             *bool                                  `json:"isActive,omitempty"`
	LastModifiedBy       *IdentitySet                           `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                                `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                `json:"@odata.type,omitempty"`
	SupportedEntities    *WorkforceIntegrationSupportedEntities `json:"supportedEntities,omitempty"`
	Url                  *string                                `json:"url,omitempty"`
}
