package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionDesktopApp struct {
	BinaryName        *string `json:"binaryName,omitempty"`
	BinaryVersionHigh *string `json:"binaryVersionHigh,omitempty"`
	BinaryVersionLow  *string `json:"binaryVersionLow,omitempty"`
	Denied            *bool   `json:"denied,omitempty"`
	Description       *string `json:"description,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	ProductName       *string `json:"productName,omitempty"`
	PublisherName     *string `json:"publisherName,omitempty"`
}
