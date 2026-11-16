package authentications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationProperties struct {
	ProvisioningState *ProvisioningState  `json:"provisioningState,omitempty"`
	State             AuthenticationState `json:"state"`
}
