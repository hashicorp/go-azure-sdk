package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectorySizeQuota struct {
	ODataType *string `json:"@odata.type,omitempty"`
	Total     *int64  `json:"total,omitempty"`
	Used      *int64  `json:"used,omitempty"`
}
