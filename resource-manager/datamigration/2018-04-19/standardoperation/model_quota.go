package standardoperation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Quota struct {
	CurrentValue *float64   `json:"currentValue,omitempty"`
	Id           *string    `json:"id,omitempty"`
	Limit        *float64   `json:"limit,omitempty"`
	Name         *QuotaName `json:"name,omitempty"`
	Unit         *string    `json:"unit,omitempty"`
}
