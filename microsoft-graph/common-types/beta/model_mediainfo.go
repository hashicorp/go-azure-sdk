package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaInfo struct {
	ODataType  *string `json:"@odata.type,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
	Uri        *string `json:"uri,omitempty"`
}
