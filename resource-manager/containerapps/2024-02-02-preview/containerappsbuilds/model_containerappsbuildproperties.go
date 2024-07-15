package containerappsbuilds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsBuildProperties struct {
	BuildStatus                  *BuildStatus                      `json:"buildStatus,omitempty"`
	Configuration                *ContainerAppsBuildConfiguration  `json:"configuration,omitempty"`
	DestinationContainerRegistry *ContainerRegistryWithCustomImage `json:"destinationContainerRegistry,omitempty"`
	LogStreamEndpoint            *string                           `json:"logStreamEndpoint,omitempty"`
	ProvisioningState            *BuildProvisioningState           `json:"provisioningState,omitempty"`
}
