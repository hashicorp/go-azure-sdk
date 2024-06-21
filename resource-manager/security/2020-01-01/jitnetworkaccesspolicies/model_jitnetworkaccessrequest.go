package jitnetworkaccesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessRequest struct {
	Justification   *string                                 `json:"justification,omitempty"`
	Requestor       string                                  `json:"requestor"`
	StartTimeUtc    string                                  `json:"startTimeUtc"`
	VirtualMachines []JitNetworkAccessRequestVirtualMachine `json:"virtualMachines"`
}
