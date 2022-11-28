package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigServerProperties struct {
	ConfigServer      *ConfigServerSettings `json:"configServer"`
	Error             *Error                `json:"error"`
	ProvisioningState *ConfigServerState    `json:"provisioningState,omitempty"`
}
