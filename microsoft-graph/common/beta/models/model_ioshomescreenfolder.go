package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosHomeScreenFolder struct {
	DisplayName *string                    `json:"displayName,omitempty"`
	ODataType   *string                    `json:"@odata.type,omitempty"`
	Pages       *[]IosHomeScreenFolderPage `json:"pages,omitempty"`
}
