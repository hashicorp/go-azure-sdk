package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonNamePronounciation struct {
	DisplayName *string `json:"displayName,omitempty"`
	First       *string `json:"first,omitempty"`
	Last        *string `json:"last,omitempty"`
	Maiden      *string `json:"maiden,omitempty"`
	Middle      *string `json:"middle,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
