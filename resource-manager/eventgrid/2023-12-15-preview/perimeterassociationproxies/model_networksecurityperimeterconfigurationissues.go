package perimeterassociationproxies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationIssues struct {
	Name       *string                                                `json:"name,omitempty"`
	Properties *NetworkSecurityPerimeterConfigurationIssuesProperties `json:"properties,omitempty"`
}
