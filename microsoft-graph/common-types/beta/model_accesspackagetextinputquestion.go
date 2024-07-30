package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageTextInputQuestion struct {
	Id                   *string                        `json:"id,omitempty"`
	IsAnswerEditable     *bool                          `json:"isAnswerEditable,omitempty"`
	IsRequired           *bool                          `json:"isRequired,omitempty"`
	IsSingleLineQuestion *bool                          `json:"isSingleLineQuestion,omitempty"`
	ODataType            *string                        `json:"@odata.type,omitempty"`
	RegexPattern         *string                        `json:"regexPattern,omitempty"`
	Sequence             *int64                         `json:"sequence,omitempty"`
	Text                 *AccessPackageLocalizedContent `json:"text,omitempty"`
}
