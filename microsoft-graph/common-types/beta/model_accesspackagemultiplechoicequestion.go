package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageMultipleChoiceQuestion struct {
	AllowsMultipleSelection *bool                          `json:"allowsMultipleSelection,omitempty"`
	Choices                 *[]AccessPackageAnswerChoice   `json:"choices,omitempty"`
	Id                      *string                        `json:"id,omitempty"`
	IsAnswerEditable        *bool                          `json:"isAnswerEditable,omitempty"`
	IsRequired              *bool                          `json:"isRequired,omitempty"`
	ODataType               *string                        `json:"@odata.type,omitempty"`
	Sequence                *int64                         `json:"sequence,omitempty"`
	Text                    *AccessPackageLocalizedContent `json:"text,omitempty"`
}
