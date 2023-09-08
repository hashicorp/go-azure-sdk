package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminTodo struct {
	Id        *string       `json:"id,omitempty"`
	ODataType *string       `json:"@odata.type,omitempty"`
	Settings  *TodoSettings `json:"settings,omitempty"`
}
