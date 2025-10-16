package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceProperties struct {
	Messaging         *Messaging         `json:"messaging,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Uuid              *string            `json:"uuid,omitempty"`
}
