package secrets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretProperties struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
	UniqueIdentifier  *string `json:"uniqueIdentifier,omitempty"`
	Value             *string `json:"value,omitempty"`
}
