package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceStatusEndpoints struct {
	Inbound *map[string]DeviceStatusEndpoint `json:"inbound,omitempty"`
}
