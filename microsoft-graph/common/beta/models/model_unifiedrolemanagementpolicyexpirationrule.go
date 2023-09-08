package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyExpirationRule struct {
	Id                   *string                                `json:"id,omitempty"`
	IsExpirationRequired *bool                                  `json:"isExpirationRequired,omitempty"`
	MaximumDuration      *string                                `json:"maximumDuration,omitempty"`
	ODataType            *string                                `json:"@odata.type,omitempty"`
	Target               *UnifiedRoleManagementPolicyRuleTarget `json:"target,omitempty"`
}
