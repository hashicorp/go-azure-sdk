package containerappssourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryInfo struct {
	RegistryPassword *string `json:"registryPassword,omitempty"`
	RegistryURL      *string `json:"registryUrl,omitempty"`
	RegistryUserName *string `json:"registryUserName,omitempty"`
}
