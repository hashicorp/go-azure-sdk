package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestion struct {
	AnswerChoices   *[]string                                        `json:"answerChoices,omitempty"`
	AnswerInputType *VirtualEventRegistrationQuestionAnswerInputType `json:"answerInputType,omitempty"`
	DisplayName     *string                                          `json:"displayName,omitempty"`
	Id              *string                                          `json:"id,omitempty"`
	IsRequired      *bool                                            `json:"isRequired,omitempty"`
	ODataType       *string                                          `json:"@odata.type,omitempty"`
}
