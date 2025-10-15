package akriconnectortemplate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorsContainerRegistrySettings struct {
	ImagePullSecrets *[]AkriConnectorsImagePullSecret `json:"imagePullSecrets,omitempty"`
	Registry         string                           `json:"registry"`
}
