package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingQuestionAnswer struct {
	Answer          *string                               `json:"answer,omitempty"`
	AnswerInputType *BookingQuestionAnswerAnswerInputType `json:"answerInputType,omitempty"`
	AnswerOptions   *[]string                             `json:"answerOptions,omitempty"`
	IsRequired      *bool                                 `json:"isRequired,omitempty"`
	ODataType       *string                               `json:"@odata.type,omitempty"`
	Question        *string                               `json:"question,omitempty"`
	QuestionId      *string                               `json:"questionId,omitempty"`
	SelectedOptions *[]string                             `json:"selectedOptions,omitempty"`
}
