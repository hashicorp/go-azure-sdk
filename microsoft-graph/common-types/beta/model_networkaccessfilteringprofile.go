package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringProfile struct {
	ConditionalAccessPolicies *[]NetworkaccessConditionalAccessPolicy `json:"conditionalAccessPolicies,omitempty"`
	CreatedDateTime           *string                                 `json:"createdDateTime,omitempty"`
	Description               *string                                 `json:"description,omitempty"`
	Id                        *string                                 `json:"id,omitempty"`
	LastModifiedDateTime      *string                                 `json:"lastModifiedDateTime,omitempty"`
	Name                      *string                                 `json:"name,omitempty"`
	ODataType                 *string                                 `json:"@odata.type,omitempty"`
	Policies                  *[]NetworkaccessPolicyLink              `json:"policies,omitempty"`
	Priority                  *int64                                  `json:"priority,omitempty"`
	State                     *NetworkaccessFilteringProfileState     `json:"state,omitempty"`
	Version                   *string                                 `json:"version,omitempty"`
}
