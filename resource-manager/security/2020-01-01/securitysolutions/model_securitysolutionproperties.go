package securitysolutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySolutionProperties struct {
	ProtectionStatus  string            `json:"protectionStatus"`
	ProvisioningState ProvisioningState `json:"provisioningState"`
	SecurityFamily    SecurityFamily    `json:"securityFamily"`
	Template          string            `json:"template"`
}
