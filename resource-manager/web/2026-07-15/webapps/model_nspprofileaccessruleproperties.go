package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NspProfileAccessRuleProperties struct {
	AddressPrefixes           *[]string                              `json:"addressPrefixes,omitempty"`
	AppliesTo                 *[]NspProfileAccessRuleGranularFeature `json:"appliesTo,omitempty"`
	Direction                 *string                                `json:"direction,omitempty"`
	EmailAddresses            *[]string                              `json:"emailAddresses,omitempty"`
	FullyQualifiedDomainNames *[]string                              `json:"fullyQualifiedDomainNames,omitempty"`
	NetworkSecurityPerimeters *[]NetworkSecurityPerimeter            `json:"networkSecurityPerimeters,omitempty"`
	PhoneNumbers              *[]string                              `json:"phoneNumbers,omitempty"`
	ServiceTags               *[]string                              `json:"serviceTags,omitempty"`
	Subscriptions             *[]NspSubscription                     `json:"subscriptions,omitempty"`
}
