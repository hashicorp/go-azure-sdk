package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosHomeScreenApp struct {
	BundleID    *string `json:"bundleID,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	IsWebClip   *bool   `json:"isWebClip,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
