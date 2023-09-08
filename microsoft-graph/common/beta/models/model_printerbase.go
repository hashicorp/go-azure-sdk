package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterBase struct {
	Capabilities    *PrinterCapabilities `json:"capabilities,omitempty"`
	Defaults        *PrinterDefaults     `json:"defaults,omitempty"`
	DisplayName     *string              `json:"displayName,omitempty"`
	Id              *string              `json:"id,omitempty"`
	IsAcceptingJobs *bool                `json:"isAcceptingJobs,omitempty"`
	Jobs            *[]PrintJob          `json:"jobs,omitempty"`
	Location        *PrinterLocation     `json:"location,omitempty"`
	Manufacturer    *string              `json:"manufacturer,omitempty"`
	Model           *string              `json:"model,omitempty"`
	Name            *string              `json:"name,omitempty"`
	ODataType       *string              `json:"@odata.type,omitempty"`
	Status          *PrinterStatus       `json:"status,omitempty"`
}
