package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppListItem struct {
	AppId       *string `json:"appId,omitempty"`
	AppStoreUrl *string `json:"appStoreUrl,omitempty"`
	Name        *string `json:"name,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Publisher   *string `json:"publisher,omitempty"`
}
