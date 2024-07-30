package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicy struct {
	AllowedCloudEndpoints *[]string                                      `json:"allowedCloudEndpoints,omitempty"`
	Default               *CrossTenantAccessPolicyConfigurationDefault   `json:"default,omitempty"`
	Definition            *[]string                                      `json:"definition,omitempty"`
	DeletedDateTime       *string                                        `json:"deletedDateTime,omitempty"`
	Description           *string                                        `json:"description,omitempty"`
	DisplayName           *string                                        `json:"displayName,omitempty"`
	Id                    *string                                        `json:"id,omitempty"`
	ODataType             *string                                        `json:"@odata.type,omitempty"`
	Partners              *[]CrossTenantAccessPolicyConfigurationPartner `json:"partners,omitempty"`
	Templates             *PolicyTemplate                                `json:"templates,omitempty"`
}
