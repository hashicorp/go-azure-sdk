package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageProperties struct {
	CurrentValue *int64  `json:"currentValue,omitempty"`
	Limit        *int64  `json:"limit,omitempty"`
	Unit         *string `json:"unit,omitempty"`
}
