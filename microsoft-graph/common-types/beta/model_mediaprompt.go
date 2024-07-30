package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaPrompt struct {
	Loop      *int64     `json:"loop,omitempty"`
	MediaInfo *MediaInfo `json:"mediaInfo,omitempty"`
	ODataType *string    `json:"@odata.type,omitempty"`
}
