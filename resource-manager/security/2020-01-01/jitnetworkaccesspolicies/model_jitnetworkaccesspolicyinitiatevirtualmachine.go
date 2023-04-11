package jitnetworkaccesspolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessPolicyInitiateVirtualMachine struct {
	Id    string                               `json:"id"`
	Ports []JitNetworkAccessPolicyInitiatePort `json:"ports"`
}
