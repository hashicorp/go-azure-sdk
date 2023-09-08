package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption string

const (
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionadSite                 DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "AdSite"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionauthenticatedDomainSid DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "AuthenticatedDomainSid"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptiondhcpUserOption         DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "DhcpUserOption"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptiondnsSuffix              DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "DnsSuffix"
	DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionnotConfigured          DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption = "NotConfigured"
)

func PossibleValuesForDeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption() []string {
	return []string{
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionadSite),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionauthenticatedDomainSid),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptiondhcpUserOption),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptiondnsSuffix),
		string(DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionnotConfigured),
	}
}

func parseDeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption(input string) (*DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption, error) {
	vals := map[string]DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption{
		"adsite":                 DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionadSite,
		"authenticateddomainsid": DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionauthenticatedDomainSid,
		"dhcpuseroption":         DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptiondhcpUserOption,
		"dnssuffix":              DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptiondnsSuffix,
		"notconfigured":          DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOptionnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryOptimizationGroupIdSourceOptionsGroupIdSourceOption(input)
	return &out, nil
}
