package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadTypes struct {
	ODataType     *string           `json:"@odata.type,omitempty"`
	RawContent    *string           `json:"rawContent,omitempty"`
	VisualContent *VisualProperties `json:"visualContent,omitempty"`
}
