package gates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GateProperties struct {
	DisplayName       *string                `json:"displayName,omitempty"`
	GateType          GateType               `json:"gateType"`
	ProvisioningState *GateProvisioningState `json:"provisioningState,omitempty"`
	State             GateState              `json:"state"`
	Target            GateTarget             `json:"target"`
}
