package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterPatchProperties struct {
	AdministratorConfiguration    *AdministratorConfigurationPatch    `json:"administratorConfiguration,omitempty"`
	ControlPlaneNodeConfiguration *ControlPlaneNodePatchConfiguration `json:"controlPlaneNodeConfiguration,omitempty"`
	KubernetesVersion             *string                             `json:"kubernetesVersion,omitempty"`
}
