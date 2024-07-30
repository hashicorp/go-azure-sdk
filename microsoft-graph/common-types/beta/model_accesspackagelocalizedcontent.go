package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageLocalizedContent struct {
	DefaultText    *string                       `json:"defaultText,omitempty"`
	LocalizedTexts *[]AccessPackageLocalizedText `json:"localizedTexts,omitempty"`
	ODataType      *string                       `json:"@odata.type,omitempty"`
}
