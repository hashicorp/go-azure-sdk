package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Media struct {
	IsTranscriptionShown *bool        `json:"isTranscriptionShown,omitempty"`
	MediaSource          *MediaSource `json:"mediaSource,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
}
