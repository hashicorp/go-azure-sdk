package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeInfo struct {
	IPConfiguration *[]KubernetesIPConfiguration `json:"ipConfiguration,omitempty"`
	Name            *string                      `json:"name,omitempty"`
	Type            *KubernetesNodeType          `json:"type,omitempty"`
}
