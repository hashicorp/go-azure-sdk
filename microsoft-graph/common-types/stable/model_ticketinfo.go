package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TicketInfo struct {
	ODataType    *string `json:"@odata.type,omitempty"`
	TicketNumber *string `json:"ticketNumber,omitempty"`
	TicketSystem *string `json:"ticketSystem,omitempty"`
}
