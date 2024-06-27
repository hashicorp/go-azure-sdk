package savingsplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanValidResponseProperty struct {
	Reason     *string `json:"reason,omitempty"`
	ReasonCode *string `json:"reasonCode,omitempty"`
	Valid      *bool   `json:"valid,omitempty"`
}
