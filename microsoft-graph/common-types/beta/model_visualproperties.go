package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VisualProperties struct {
	Body      *string `json:"body,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Title     *string `json:"title,omitempty"`
}
