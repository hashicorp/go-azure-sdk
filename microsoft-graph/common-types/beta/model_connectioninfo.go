package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionInfo struct {
	ODataType *string `json:"@odata.type,omitempty"`
	Url       *string `json:"url,omitempty"`
}
