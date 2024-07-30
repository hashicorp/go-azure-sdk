package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreLocalizedDescription struct {
	Description *string `json:"description,omitempty"`
	LanguageTag *string `json:"languageTag,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
