package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddWatermark struct {
	FontColor   *string                  `json:"fontColor,omitempty"`
	FontSize    *int64                   `json:"fontSize,omitempty"`
	Name        *string                  `json:"name,omitempty"`
	ODataType   *string                  `json:"@odata.type,omitempty"`
	Orientation *AddWatermarkOrientation `json:"orientation,omitempty"`
	Text        *string                  `json:"text,omitempty"`
}
