package cacherules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheRuleProperties struct {
	CreationDate            *string            `json:"creationDate,omitempty"`
	CredentialSetResourceId *string            `json:"credentialSetResourceId,omitempty"`
	ProvisioningState       *ProvisioningState `json:"provisioningState,omitempty"`
	SourceRepository        *string            `json:"sourceRepository,omitempty"`
	TargetRepository        *string            `json:"targetRepository,omitempty"`
}
