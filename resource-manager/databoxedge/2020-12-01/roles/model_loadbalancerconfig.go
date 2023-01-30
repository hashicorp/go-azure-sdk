package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoadBalancerConfig struct {
	Type    *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}
