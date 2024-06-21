package registries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftDeletePolicy struct {
	LastUpdatedTime *string       `json:"lastUpdatedTime,omitempty"`
	RetentionDays   *int64        `json:"retentionDays,omitempty"`
	Status          *PolicyStatus `json:"status,omitempty"`
}
