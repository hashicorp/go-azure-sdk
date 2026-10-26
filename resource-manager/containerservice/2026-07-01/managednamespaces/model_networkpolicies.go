package managednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkPolicies struct {
	Egress  *PolicyRule `json:"egress,omitempty"`
	Ingress *PolicyRule `json:"ingress,omitempty"`
}
