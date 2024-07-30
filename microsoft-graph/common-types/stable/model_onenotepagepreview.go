package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePagePreview struct {
	Links       *OnenotePagePreviewLinks `json:"links,omitempty"`
	ODataType   *string                  `json:"@odata.type,omitempty"`
	PreviewText *string                  `json:"previewText,omitempty"`
}
