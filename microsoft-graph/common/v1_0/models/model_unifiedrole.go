package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRole struct {
	ODataType        *string `json:"@odata.type,omitempty"`
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`
}
