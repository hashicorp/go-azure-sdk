package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisplayNameLocalization struct {
	DisplayName *string `json:"displayName,omitempty"`
	LanguageTag *string `json:"languageTag,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
