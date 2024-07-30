package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HttpRequestEndpoint struct {
	ODataType *string `json:"@odata.type,omitempty"`
	TargetUrl *string `json:"targetUrl,omitempty"`
}
