package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesIPConfiguration struct {
	IPAddress *string `json:"ipAddress,omitempty"`
	Port      *string `json:"port,omitempty"`
}
