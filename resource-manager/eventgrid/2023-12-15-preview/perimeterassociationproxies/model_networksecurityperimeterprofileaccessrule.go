package perimeterassociationproxies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterProfileAccessRule struct {
	FullyQualifiedArmId *string                                              `json:"fullyQualifiedArmId,omitempty"`
	Name                *string                                              `json:"name,omitempty"`
	Properties          *NetworkSecurityPerimeterProfileAccessRuleProperties `json:"properties,omitempty"`
	Type                *string                                              `json:"type,omitempty"`
}
