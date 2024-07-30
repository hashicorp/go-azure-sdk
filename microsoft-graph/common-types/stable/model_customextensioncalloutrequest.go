package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutRequest struct {
	Data      *CustomExtensionData `json:"data,omitempty"`
	ODataType *string              `json:"@odata.type,omitempty"`
	Source    *string              `json:"source,omitempty"`
	Type      *string              `json:"type,omitempty"`
}
