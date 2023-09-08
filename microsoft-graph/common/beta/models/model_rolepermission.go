package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RolePermission struct {
	Actions         *[]string         `json:"actions,omitempty"`
	ODataType       *string           `json:"@odata.type,omitempty"`
	ResourceActions *[]ResourceAction `json:"resourceActions,omitempty"`
}
