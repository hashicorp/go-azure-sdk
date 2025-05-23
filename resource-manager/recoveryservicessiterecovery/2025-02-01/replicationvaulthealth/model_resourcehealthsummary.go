package replicationvaulthealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceHealthSummary struct {
	CategorizedResourceCounts *map[string]int64     `json:"categorizedResourceCounts,omitempty"`
	Issues                    *[]HealthErrorSummary `json:"issues,omitempty"`
	ResourceCount             *int64                `json:"resourceCount,omitempty"`
}
