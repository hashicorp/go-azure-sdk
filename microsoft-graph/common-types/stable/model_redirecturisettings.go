package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedirectUriSettings struct {
	Index     *int64  `json:"index,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Uri       *string `json:"uri,omitempty"`
}
