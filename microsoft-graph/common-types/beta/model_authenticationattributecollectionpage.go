package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionPage struct {
	CustomStringsFileId *string                                                   `json:"customStringsFileId,omitempty"`
	ODataType           *string                                                   `json:"@odata.type,omitempty"`
	Views               *[]AuthenticationAttributeCollectionPageViewConfiguration `json:"views,omitempty"`
}
