package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRuleSatisfied struct {
	ConditionalAccessCondition *ConditionalAccessRuleSatisfiedConditionalAccessCondition `json:"conditionalAccessCondition,omitempty"`
	ODataType                  *string                                                   `json:"@odata.type,omitempty"`
	RuleSatisfied              *ConditionalAccessRuleSatisfiedRuleSatisfied              `json:"ruleSatisfied,omitempty"`
}
