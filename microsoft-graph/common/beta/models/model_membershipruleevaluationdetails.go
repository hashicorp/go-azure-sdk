package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipRuleEvaluationDetails struct {
	MembershipRuleEvaluationDetails *ExpressionEvaluationDetails `json:"membershipRuleEvaluationDetails,omitempty"`
	ODataType                       *string                      `json:"@odata.type,omitempty"`
}
