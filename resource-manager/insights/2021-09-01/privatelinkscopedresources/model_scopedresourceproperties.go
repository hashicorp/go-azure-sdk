package privatelinkscopedresources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopedResourceProperties struct {
	LinkedResourceId  *string                          `json:"linkedResourceId,omitempty"`
	ProvisioningState *ScopedResourceProvisioningState `json:"provisioningState,omitempty"`
}
