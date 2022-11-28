package videos

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VideoProperties struct {
	Description *string         `json:"description,omitempty"`
	Flags       *VideoFlags     `json:"flags"`
	MediaInfo   *VideoMediaInfo `json:"mediaInfo"`
	Streaming   *VideoStreaming `json:"streaming"`
	Title       *string         `json:"title,omitempty"`
	Type        *VideoType      `json:"type,omitempty"`
}
