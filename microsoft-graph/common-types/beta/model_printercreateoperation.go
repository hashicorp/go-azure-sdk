package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCreateOperation struct {
	Certificate     *string               `json:"certificate,omitempty"`
	CreatedDateTime *string               `json:"createdDateTime,omitempty"`
	Id              *string               `json:"id,omitempty"`
	ODataType       *string               `json:"@odata.type,omitempty"`
	Printer         *Printer              `json:"printer,omitempty"`
	Status          *PrintOperationStatus `json:"status,omitempty"`
}
