package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLaunchItem struct {
	Hide      *bool   `json:"hide,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Path      *string `json:"path,omitempty"`
}
