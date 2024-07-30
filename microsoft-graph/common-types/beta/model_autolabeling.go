package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoLabeling struct {
	Message          *string   `json:"message,omitempty"`
	ODataType        *string   `json:"@odata.type,omitempty"`
	SensitiveTypeIds *[]string `json:"sensitiveTypeIds,omitempty"`
}
