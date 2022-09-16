package hcireports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportResource struct {
	Error  *ErrorDetail `json:"error,omitempty"`
	Id     *string      `json:"id,omitempty"`
	Name   *string      `json:"name,omitempty"`
	Status *string      `json:"status,omitempty"`
	Type   *string      `json:"type,omitempty"`
}
