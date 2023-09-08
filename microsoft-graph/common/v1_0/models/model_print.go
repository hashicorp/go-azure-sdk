package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Print struct {
	Connectors      *[]PrintConnector      `json:"connectors,omitempty"`
	ODataType       *string                `json:"@odata.type,omitempty"`
	Operations      *[]PrintOperation      `json:"operations,omitempty"`
	Printers        *[]Printer             `json:"printers,omitempty"`
	Services        *[]PrintService        `json:"services,omitempty"`
	Settings        *PrintSettings         `json:"settings,omitempty"`
	Shares          *[]PrinterShare        `json:"shares,omitempty"`
	TaskDefinitions *[]PrintTaskDefinition `json:"taskDefinitions,omitempty"`
}
