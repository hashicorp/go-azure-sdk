package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintOperation struct {
	CreatedDateTime *string               `json:"createdDateTime,omitempty"`
	Id              *string               `json:"id,omitempty"`
	ODataType       *string               `json:"@odata.type,omitempty"`
	Status          *PrintOperationStatus `json:"status,omitempty"`
}
