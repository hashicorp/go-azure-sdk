package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpScope struct {
	ODataType    *string                         `json:"@odata.type,omitempty"`
	ResourceType *string                         `json:"resourceType,omitempty"`
	Service      *AuthorizationSystemTypeService `json:"service,omitempty"`
}
