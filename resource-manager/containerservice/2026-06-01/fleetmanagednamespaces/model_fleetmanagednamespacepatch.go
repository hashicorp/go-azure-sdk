package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FleetManagedNamespacePatch struct {
	Tags *map[string]string `json:"tags,omitempty"`
}
