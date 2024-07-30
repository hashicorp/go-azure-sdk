package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceVisualization struct {
	ContainerDisplayName *string `json:"containerDisplayName,omitempty"`
	ContainerType        *string `json:"containerType,omitempty"`
	ContainerWebUrl      *string `json:"containerWebUrl,omitempty"`
	MediaType            *string `json:"mediaType,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	PreviewImageUrl      *string `json:"previewImageUrl,omitempty"`
	PreviewText          *string `json:"previewText,omitempty"`
	Title                *string `json:"title,omitempty"`
	Type                 *string `json:"type,omitempty"`
}
