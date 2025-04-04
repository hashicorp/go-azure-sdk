package usages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Usage struct {
	CurrentValue float64   `json:"currentValue"`
	Limit        float64   `json:"limit"`
	Name         UsageName `json:"name"`
	Unit         UsageUnit `json:"unit"`
}
