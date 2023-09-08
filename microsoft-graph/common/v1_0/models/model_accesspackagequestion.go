package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageQuestion struct {
	Id               *string                       `json:"id,omitempty"`
	IsAnswerEditable *bool                         `json:"isAnswerEditable,omitempty"`
	IsRequired       *bool                         `json:"isRequired,omitempty"`
	Localizations    *[]AccessPackageLocalizedText `json:"localizations,omitempty"`
	ODataType        *string                       `json:"@odata.type,omitempty"`
	Sequence         *int64                        `json:"sequence,omitempty"`
	Text             *string                       `json:"text,omitempty"`
}
