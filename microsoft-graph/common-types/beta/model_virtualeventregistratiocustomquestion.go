package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistratioCustomQuestion struct {
	AnswerChoices   *[]string                                             `json:"answerChoices,omitempty"`
	AnswerInputType *VirtualEventRegistratioCustomQuestionAnswerInputType `json:"answerInputType,omitempty"`
	DisplayName     *string                                               `json:"displayName,omitempty"`
	Id              *string                                               `json:"id,omitempty"`
	IsRequired      *bool                                                 `json:"isRequired,omitempty"`
	ODataType       *string                                               `json:"@odata.type,omitempty"`
}
