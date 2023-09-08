package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingRule struct {
	Action       *NetworkaccessForwardingRuleAction   `json:"action,omitempty"`
	Destinations *[]NetworkaccessRuleDestination      `json:"destinations,omitempty"`
	Id           *string                              `json:"id,omitempty"`
	Name         *string                              `json:"name,omitempty"`
	ODataType    *string                              `json:"@odata.type,omitempty"`
	RuleType     *NetworkaccessForwardingRuleRuleType `json:"ruleType,omitempty"`
}
