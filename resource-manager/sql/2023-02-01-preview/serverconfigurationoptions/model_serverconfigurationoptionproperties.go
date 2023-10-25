package serverconfigurationoptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerConfigurationOptionProperties struct {
	ProvisioningState              *ProvisioningState `json:"provisioningState,omitempty"`
	ServerConfigurationOptionValue int64              `json:"serverConfigurationOptionValue"`
}
