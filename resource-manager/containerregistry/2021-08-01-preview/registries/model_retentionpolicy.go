package registries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionPolicy struct {
	Days            *int64        `json:"days,omitempty"`
	LastUpdatedTime *string       `json:"lastUpdatedTime,omitempty"`
	Status          *PolicyStatus `json:"status,omitempty"`
}
