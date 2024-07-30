package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaStream struct {
	Direction   *MediaStreamDirection `json:"direction,omitempty"`
	Label       *string               `json:"label,omitempty"`
	MediaType   *MediaStreamMediaType `json:"mediaType,omitempty"`
	ODataType   *string               `json:"@odata.type,omitempty"`
	ServerMuted *bool                 `json:"serverMuted,omitempty"`
	SourceId    *string               `json:"sourceId,omitempty"`
}
