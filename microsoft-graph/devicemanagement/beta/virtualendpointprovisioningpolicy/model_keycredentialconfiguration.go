package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyCredentialConfiguration struct {
	CertificateBasedApplicationConfigurationIds *[]string                                  `json:"certificateBasedApplicationConfigurationIds,omitempty"`
	MaxLifetime                                 *string                                    `json:"maxLifetime,omitempty"`
	ODataType                                   *string                                    `json:"@odata.type,omitempty"`
	RestrictForAppsCreatedAfterDateTime         *string                                    `json:"restrictForAppsCreatedAfterDateTime,omitempty"`
	RestrictionType                             *KeyCredentialConfigurationRestrictionType `json:"restrictionType,omitempty"`
}
