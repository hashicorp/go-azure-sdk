package quotabycounterkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaCounterContract struct {
	CounterKey      string                               `json:"counterKey"`
	PeriodEndTime   string                               `json:"periodEndTime"`
	PeriodKey       string                               `json:"periodKey"`
	PeriodStartTime string                               `json:"periodStartTime"`
	Value           *QuotaCounterValueContractProperties `json:"value,omitempty"`
}
