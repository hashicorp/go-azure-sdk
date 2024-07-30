package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRule struct {
	Actions     *MessageRuleActions    `json:"actions,omitempty"`
	Conditions  *MessageRulePredicates `json:"conditions,omitempty"`
	DisplayName *string                `json:"displayName,omitempty"`
	Exceptions  *MessageRulePredicates `json:"exceptions,omitempty"`
	HasError    *bool                  `json:"hasError,omitempty"`
	Id          *string                `json:"id,omitempty"`
	IsEnabled   *bool                  `json:"isEnabled,omitempty"`
	IsReadOnly  *bool                  `json:"isReadOnly,omitempty"`
	ODataType   *string                `json:"@odata.type,omitempty"`
	Sequence    *int64                 `json:"sequence,omitempty"`
}
