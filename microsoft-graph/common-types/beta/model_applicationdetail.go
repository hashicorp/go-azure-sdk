package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationDetail struct {
	FileDescription     *string `json:"fileDescription,omitempty"`
	FileHash            *string `json:"fileHash,omitempty"`
	FileName            *string `json:"fileName,omitempty"`
	FilePath            *string `json:"filePath,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	ProductInternalName *string `json:"productInternalName,omitempty"`
	ProductName         *string `json:"productName,omitempty"`
	ProductVersion      *string `json:"productVersion,omitempty"`
	PublisherName       *string `json:"publisherName,omitempty"`
}
