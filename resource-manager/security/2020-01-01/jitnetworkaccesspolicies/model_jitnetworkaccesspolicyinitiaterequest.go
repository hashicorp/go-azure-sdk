package jitnetworkaccesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessPolicyInitiateRequest struct {
	Justification   *string                                        `json:"justification,omitempty"`
	VirtualMachines []JitNetworkAccessPolicyInitiateVirtualMachine `json:"virtualMachines"`
}
