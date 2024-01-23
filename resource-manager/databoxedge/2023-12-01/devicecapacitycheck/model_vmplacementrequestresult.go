package devicecapacitycheck

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMPlacementRequestResult struct {
	IsFeasible  *bool     `json:"isFeasible,omitempty"`
	Message     *string   `json:"message,omitempty"`
	MessageCode *string   `json:"messageCode,omitempty"`
	VMSize      *[]string `json:"vmSize,omitempty"`
}
