package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringPolicy struct {
	Action               *NetworkaccessFilteringPolicyAction `json:"action,omitempty"`
	CreatedDateTime      *string                             `json:"createdDateTime,omitempty"`
	Description          *string                             `json:"description,omitempty"`
	Id                   *string                             `json:"id,omitempty"`
	LastModifiedDateTime *string                             `json:"lastModifiedDateTime,omitempty"`
	Name                 *string                             `json:"name,omitempty"`
	ODataType            *string                             `json:"@odata.type,omitempty"`
	PolicyRules          *[]NetworkaccessPolicyRule          `json:"policyRules,omitempty"`
	Version              *string                             `json:"version,omitempty"`
}
