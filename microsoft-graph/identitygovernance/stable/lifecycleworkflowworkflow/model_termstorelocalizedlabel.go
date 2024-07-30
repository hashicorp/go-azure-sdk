package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreLocalizedLabel struct {
	IsDefault   *bool   `json:"isDefault,omitempty"`
	LanguageTag *string `json:"languageTag,omitempty"`
	Name        *string `json:"name,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
