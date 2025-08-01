package guestconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationInfo struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}
