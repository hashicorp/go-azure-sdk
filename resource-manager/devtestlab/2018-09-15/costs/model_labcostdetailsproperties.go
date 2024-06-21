package costs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabCostDetailsProperties struct {
	Cost     *float64  `json:"cost,omitempty"`
	CostType *CostType `json:"costType,omitempty"`
	Date     *string   `json:"date,omitempty"`
}
