package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingQuestionAssignment struct {
	IsRequired *bool   `json:"isRequired,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	QuestionId *string `json:"questionId,omitempty"`
}
