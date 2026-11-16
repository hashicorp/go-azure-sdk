package discoveryrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveryError struct {
	Context *[]string `json:"context,omitempty"`
	Message string    `json:"message"`
}
