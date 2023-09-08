package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditResource struct {
	DisplayName        *string                 `json:"displayName,omitempty"`
	ModifiedProperties *[]CloudPCAuditProperty `json:"modifiedProperties,omitempty"`
	ODataType          *string                 `json:"@odata.type,omitempty"`
	ResourceId         *string                 `json:"resourceId,omitempty"`
	Type               *string                 `json:"type,omitempty"`
}
