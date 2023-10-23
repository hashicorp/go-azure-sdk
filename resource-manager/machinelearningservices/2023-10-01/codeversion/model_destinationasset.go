package codeversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DestinationAsset struct {
	DestinationName    *string `json:"destinationName,omitempty"`
	DestinationVersion *string `json:"destinationVersion,omitempty"`
	RegistryName       *string `json:"registryName,omitempty"`
}
