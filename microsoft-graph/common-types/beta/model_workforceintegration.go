package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegration struct {
	ApiVersion                          *int64                                                   `json:"apiVersion,omitempty"`
	CreatedBy                           *IdentitySet                                             `json:"createdBy,omitempty"`
	CreatedDateTime                     *string                                                  `json:"createdDateTime,omitempty"`
	DisplayName                         *string                                                  `json:"displayName,omitempty"`
	EligibilityFilteringEnabledEntities *WorkforceIntegrationEligibilityFilteringEnabledEntities `json:"eligibilityFilteringEnabledEntities,omitempty"`
	Encryption                          *WorkforceIntegrationEncryption                          `json:"encryption,omitempty"`
	Id                                  *string                                                  `json:"id,omitempty"`
	IsActive                            *bool                                                    `json:"isActive,omitempty"`
	LastModifiedBy                      *IdentitySet                                             `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime                *string                                                  `json:"lastModifiedDateTime,omitempty"`
	ODataType                           *string                                                  `json:"@odata.type,omitempty"`
	SupportedEntities                   *WorkforceIntegrationSupportedEntities                   `json:"supportedEntities,omitempty"`
	Supports                            *WorkforceIntegrationSupports                            `json:"supports,omitempty"`
	Url                                 *string                                                  `json:"url,omitempty"`
}
