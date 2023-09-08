package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkDetails struct {
	ODataType      *string `json:"@odata.type,omitempty"`
	PolicyId       *string `json:"policyId,omitempty"`
	PolicyName     *string `json:"policyName,omitempty"`
	PolicyTenantId *string `json:"policyTenantId,omitempty"`
	ResourceId     *string `json:"resourceId,omitempty"`
}
