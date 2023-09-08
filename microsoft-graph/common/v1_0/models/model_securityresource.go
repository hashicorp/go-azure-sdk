package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityResource struct {
	ODataType    *string                       `json:"@odata.type,omitempty"`
	Resource     *string                       `json:"resource,omitempty"`
	ResourceType *SecurityResourceResourceType `json:"resourceType,omitempty"`
}
