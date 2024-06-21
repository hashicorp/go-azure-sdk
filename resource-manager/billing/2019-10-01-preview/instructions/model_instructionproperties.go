package instructions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstructionProperties struct {
	Amount       float64 `json:"amount"`
	CreationDate *string `json:"creationDate,omitempty"`
	EndDate      string  `json:"endDate"`
	StartDate    string  `json:"startDate"`
}
