package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionApp struct {
	Denied        *bool   `json:"denied,omitempty"`
	Description   *string `json:"description,omitempty"`
	DisplayName   *string `json:"displayName,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	ProductName   *string `json:"productName,omitempty"`
	PublisherName *string `json:"publisherName,omitempty"`
}
