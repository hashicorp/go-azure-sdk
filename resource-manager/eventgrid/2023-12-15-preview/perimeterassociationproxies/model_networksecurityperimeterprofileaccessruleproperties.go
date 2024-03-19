package perimeterassociationproxies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterProfileAccessRuleProperties struct {
	AddressPrefixes           *[]string                                           `json:"addressPrefixes,omitempty"`
	Direction                 *NetworkSecurityPerimeterProfileAccessRuleDirection `json:"direction,omitempty"`
	EmailAddresses            *[]string                                           `json:"emailAddresses,omitempty"`
	FullyQualifiedDomainNames *[]string                                           `json:"fullyQualifiedDomainNames,omitempty"`
	NetworkSecurityPerimeters *[]NetworkSecurityPerimeterInfo                     `json:"networkSecurityPerimeters,omitempty"`
	PhoneNumbers              *[]string                                           `json:"phoneNumbers,omitempty"`
	Subscriptions             *[]string                                           `json:"subscriptions,omitempty"`
}
