package replicationjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InconsistentVMDetails struct {
	CloudName          *string   `json:"cloudName,omitempty"`
	Details            *[]string `json:"details,omitempty"`
	ErrorIds           *[]string `json:"errorIds,omitempty"`
	VirtualMachineName *string   `json:"vmName,omitempty"`
}
