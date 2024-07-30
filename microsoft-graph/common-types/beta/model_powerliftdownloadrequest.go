package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerliftDownloadRequest struct {
	Files       *[]string `json:"files,omitempty"`
	ODataType   *string   `json:"@odata.type,omitempty"`
	PowerliftId *string   `json:"powerliftId,omitempty"`
}
