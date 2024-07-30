package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Printer struct {
	AcceptingJobs      *bool                `json:"acceptingJobs,omitempty"`
	Capabilities       *PrinterCapabilities `json:"capabilities,omitempty"`
	Connectors         *[]PrintConnector    `json:"connectors,omitempty"`
	Defaults           *PrinterDefaults     `json:"defaults,omitempty"`
	DisplayName        *string              `json:"displayName,omitempty"`
	HasPhysicalDevice  *bool                `json:"hasPhysicalDevice,omitempty"`
	Id                 *string              `json:"id,omitempty"`
	IsAcceptingJobs    *bool                `json:"isAcceptingJobs,omitempty"`
	IsShared           *bool                `json:"isShared,omitempty"`
	Jobs               *[]PrintJob          `json:"jobs,omitempty"`
	LastSeenDateTime   *string              `json:"lastSeenDateTime,omitempty"`
	Location           *PrinterLocation     `json:"location,omitempty"`
	Manufacturer       *string              `json:"manufacturer,omitempty"`
	Model              *string              `json:"model,omitempty"`
	Name               *string              `json:"name,omitempty"`
	ODataType          *string              `json:"@odata.type,omitempty"`
	RegisteredDateTime *string              `json:"registeredDateTime,omitempty"`
	Share              *PrinterShare        `json:"share,omitempty"`
	Shares             *[]PrinterShare      `json:"shares,omitempty"`
	Status             *PrinterStatus       `json:"status,omitempty"`
	TaskTriggers       *[]PrintTaskTrigger  `json:"taskTriggers,omitempty"`
}
