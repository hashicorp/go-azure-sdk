package managedclusterversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedClusterVersionDetails struct {
	ClusterCodeVersion *string `json:"clusterCodeVersion,omitempty"`
	OsType             *OsType `json:"osType,omitempty"`
	SupportExpiryUtc   *string `json:"supportExpiryUtc,omitempty"`
}
