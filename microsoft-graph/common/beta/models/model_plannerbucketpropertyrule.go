package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerBucketPropertyRule struct {
	ODataType *string                            `json:"@odata.type,omitempty"`
	Order     *[]string                          `json:"order,omitempty"`
	RuleKind  *PlannerBucketPropertyRuleRuleKind `json:"ruleKind,omitempty"`
	Title     *[]string                          `json:"title,omitempty"`
}
