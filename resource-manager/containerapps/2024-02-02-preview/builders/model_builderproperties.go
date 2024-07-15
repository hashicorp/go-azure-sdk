package builders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuilderProperties struct {
	ContainerRegistries *[]ContainerRegistry      `json:"containerRegistries,omitempty"`
	EnvironmentId       string                    `json:"environmentId"`
	ProvisioningState   *BuilderProvisioningState `json:"provisioningState,omitempty"`
}
