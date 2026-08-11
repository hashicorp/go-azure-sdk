package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeeringPassphrases struct {
	ClusterPeeringCommand    string  `json:"clusterPeeringCommand"`
	ClusterPeeringPassphrase string  `json:"clusterPeeringPassphrase"`
	CriticalWarning          *string `json:"criticalWarning,omitempty"`
	VserverPeeringCommand    string  `json:"vserverPeeringCommand"`
}
