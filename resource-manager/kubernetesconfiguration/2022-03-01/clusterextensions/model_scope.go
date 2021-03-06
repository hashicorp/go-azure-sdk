package clusterextensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Scope struct {
	Cluster   *ScopeCluster   `json:"cluster,omitempty"`
	Namespace *ScopeNamespace `json:"namespace,omitempty"`
}
