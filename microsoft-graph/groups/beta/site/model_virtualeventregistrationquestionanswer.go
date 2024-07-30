package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestionAnswer struct {
	BooleanValue      *bool     `json:"booleanValue,omitempty"`
	DisplayName       *string   `json:"displayName,omitempty"`
	MultiChoiceValues *[]string `json:"multiChoiceValues,omitempty"`
	ODataType         *string   `json:"@odata.type,omitempty"`
	QuestionId        *string   `json:"questionId,omitempty"`
	Value             *string   `json:"value,omitempty"`
}
