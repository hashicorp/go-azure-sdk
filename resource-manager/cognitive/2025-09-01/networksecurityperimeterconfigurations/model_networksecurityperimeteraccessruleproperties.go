package networksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterAccessRuleProperties struct {
	AddressPrefixes           *[]string                                                           `json:"addressPrefixes,omitempty"`
	Direction                 *NspAccessRuleDirection                                             `json:"direction,omitempty"`
	FullyQualifiedDomainNames *[]string                                                           `json:"fullyQualifiedDomainNames,omitempty"`
	NetworkSecurityPerimeters *[]NetworkSecurityPerimeter                                         `json:"networkSecurityPerimeters,omitempty"`
	Subscriptions             *[]NetworkSecurityPerimeterAccessRulePropertiesSubscriptionsInlined `json:"subscriptions,omitempty"`
}
