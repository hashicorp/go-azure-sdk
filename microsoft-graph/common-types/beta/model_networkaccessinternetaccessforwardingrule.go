package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessInternetAccessForwardingRule struct {
	Action       *NetworkaccessInternetAccessForwardingRuleAction   `json:"action,omitempty"`
	Destinations *[]NetworkaccessRuleDestination                    `json:"destinations,omitempty"`
	Id           *string                                            `json:"id,omitempty"`
	Name         *string                                            `json:"name,omitempty"`
	ODataType    *string                                            `json:"@odata.type,omitempty"`
	RuleType     *NetworkaccessInternetAccessForwardingRuleRuleType `json:"ruleType,omitempty"`
}
