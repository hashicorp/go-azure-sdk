package usages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Usage struct {
	CurrentValue   *int64  `json:"currentValue,omitempty"`
	Id             *string `json:"id,omitempty"`
	Limit          *int64  `json:"limit,omitempty"`
	Name           *Name   `json:"name,omitempty"`
	RequestedLimit *int64  `json:"requestedLimit,omitempty"`
	Type           *string `json:"type,omitempty"`
	Unit           *string `json:"unit,omitempty"`
}
