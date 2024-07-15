package builds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildProperties struct {
	BuildStatus                  *BuildStatus                      `json:"buildStatus,omitempty"`
	Configuration                *BuildConfiguration               `json:"configuration,omitempty"`
	DestinationContainerRegistry *ContainerRegistryWithCustomImage `json:"destinationContainerRegistry,omitempty"`
	LogStreamEndpoint            *string                           `json:"logStreamEndpoint,omitempty"`
	ProvisioningState            *BuildProvisioningState           `json:"provisioningState,omitempty"`
	TokenEndpoint                *string                           `json:"tokenEndpoint,omitempty"`
	UploadEndpoint               *string                           `json:"uploadEndpoint,omitempty"`
}
