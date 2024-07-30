package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRule struct {
	Action       *NetworkaccessM365ForwardingRuleAction   `json:"action,omitempty"`
	Category     *NetworkaccessM365ForwardingRuleCategory `json:"category,omitempty"`
	Destinations *[]NetworkaccessRuleDestination          `json:"destinations,omitempty"`
	Id           *string                                  `json:"id,omitempty"`
	Name         *string                                  `json:"name,omitempty"`
	ODataType    *string                                  `json:"@odata.type,omitempty"`
	Ports        *[]string                                `json:"ports,omitempty"`
	Protocol     *NetworkaccessM365ForwardingRuleProtocol `json:"protocol,omitempty"`
	RuleType     *NetworkaccessM365ForwardingRuleRuleType `json:"ruleType,omitempty"`
}
