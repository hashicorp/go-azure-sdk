package applicationwhitelistings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublisherInfo struct {
	BinaryName    *string `json:"binaryName,omitempty"`
	ProductName   *string `json:"productName,omitempty"`
	PublisherName *string `json:"publisherName,omitempty"`
	Version       *string `json:"version,omitempty"`
}
