package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyRuleDelta struct {
	Action    *NetworkaccessPolicyRuleDeltaAction `json:"action,omitempty"`
	ODataType *string                             `json:"@odata.type,omitempty"`
	RuleId    *string                             `json:"ruleId,omitempty"`
}
