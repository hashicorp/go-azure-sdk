package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesRoleResources struct {
	Compute KubernetesRoleCompute  `json:"compute"`
	Network *KubernetesRoleNetwork `json:"network,omitempty"`
	Storage *KubernetesRoleStorage `json:"storage,omitempty"`
}
