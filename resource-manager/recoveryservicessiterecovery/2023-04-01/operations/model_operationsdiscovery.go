package operations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationsDiscovery struct {
	Display    *Display     `json:"display,omitempty"`
	Name       *string      `json:"name,omitempty"`
	Origin     *string      `json:"origin,omitempty"`
	Properties *interface{} `json:"properties,omitempty"`
}
