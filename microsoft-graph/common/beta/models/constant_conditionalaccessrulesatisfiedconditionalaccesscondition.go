package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRuleSatisfiedConditionalAccessCondition string

const (
	ConditionalAccessRuleSatisfiedConditionalAccessConditionapplication                     ConditionalAccessRuleSatisfiedConditionalAccessCondition = "Application"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionauthenticationFlows             ConditionalAccessRuleSatisfiedConditionalAccessCondition = "AuthenticationFlows"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionclient                          ConditionalAccessRuleSatisfiedConditionalAccessCondition = "Client"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionclientType                      ConditionalAccessRuleSatisfiedConditionalAccessCondition = "ClientType"
	ConditionalAccessRuleSatisfiedConditionalAccessConditiondevicePlatform                  ConditionalAccessRuleSatisfiedConditionalAccessCondition = "DevicePlatform"
	ConditionalAccessRuleSatisfiedConditionalAccessConditiondeviceState                     ConditionalAccessRuleSatisfiedConditionalAccessCondition = "DeviceState"
	ConditionalAccessRuleSatisfiedConditionalAccessConditioninsiderRisk                     ConditionalAccessRuleSatisfiedConditionalAccessCondition = "InsiderRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionipAddressSeenByAzureAD          ConditionalAccessRuleSatisfiedConditionalAccessCondition = "IpAddressSeenByAzureAD"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionipAddressSeenByResourceProvider ConditionalAccessRuleSatisfiedConditionalAccessCondition = "IpAddressSeenByResourceProvider"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionlocation                        ConditionalAccessRuleSatisfiedConditionalAccessCondition = "Location"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionnone                            ConditionalAccessRuleSatisfiedConditionalAccessCondition = "None"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionservicePrincipalRisk            ConditionalAccessRuleSatisfiedConditionalAccessCondition = "ServicePrincipalRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionservicePrincipals               ConditionalAccessRuleSatisfiedConditionalAccessCondition = "ServicePrincipals"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionsignInRisk                      ConditionalAccessRuleSatisfiedConditionalAccessCondition = "SignInRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessConditiontime                            ConditionalAccessRuleSatisfiedConditionalAccessCondition = "Time"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionuserRisk                        ConditionalAccessRuleSatisfiedConditionalAccessCondition = "UserRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessConditionusers                           ConditionalAccessRuleSatisfiedConditionalAccessCondition = "Users"
)

func PossibleValuesForConditionalAccessRuleSatisfiedConditionalAccessCondition() []string {
	return []string{
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionapplication),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionauthenticationFlows),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionclient),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionclientType),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditiondevicePlatform),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditiondeviceState),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditioninsiderRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionipAddressSeenByAzureAD),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionipAddressSeenByResourceProvider),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionlocation),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionnone),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionservicePrincipalRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionservicePrincipals),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionsignInRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditiontime),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionuserRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessConditionusers),
	}
}

func parseConditionalAccessRuleSatisfiedConditionalAccessCondition(input string) (*ConditionalAccessRuleSatisfiedConditionalAccessCondition, error) {
	vals := map[string]ConditionalAccessRuleSatisfiedConditionalAccessCondition{
		"application":                     ConditionalAccessRuleSatisfiedConditionalAccessConditionapplication,
		"authenticationflows":             ConditionalAccessRuleSatisfiedConditionalAccessConditionauthenticationFlows,
		"client":                          ConditionalAccessRuleSatisfiedConditionalAccessConditionclient,
		"clienttype":                      ConditionalAccessRuleSatisfiedConditionalAccessConditionclientType,
		"deviceplatform":                  ConditionalAccessRuleSatisfiedConditionalAccessConditiondevicePlatform,
		"devicestate":                     ConditionalAccessRuleSatisfiedConditionalAccessConditiondeviceState,
		"insiderrisk":                     ConditionalAccessRuleSatisfiedConditionalAccessConditioninsiderRisk,
		"ipaddressseenbyazuread":          ConditionalAccessRuleSatisfiedConditionalAccessConditionipAddressSeenByAzureAD,
		"ipaddressseenbyresourceprovider": ConditionalAccessRuleSatisfiedConditionalAccessConditionipAddressSeenByResourceProvider,
		"location":                        ConditionalAccessRuleSatisfiedConditionalAccessConditionlocation,
		"none":                            ConditionalAccessRuleSatisfiedConditionalAccessConditionnone,
		"serviceprincipalrisk":            ConditionalAccessRuleSatisfiedConditionalAccessConditionservicePrincipalRisk,
		"serviceprincipals":               ConditionalAccessRuleSatisfiedConditionalAccessConditionservicePrincipals,
		"signinrisk":                      ConditionalAccessRuleSatisfiedConditionalAccessConditionsignInRisk,
		"time":                            ConditionalAccessRuleSatisfiedConditionalAccessConditiontime,
		"userrisk":                        ConditionalAccessRuleSatisfiedConditionalAccessConditionuserRisk,
		"users":                           ConditionalAccessRuleSatisfiedConditionalAccessConditionusers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessRuleSatisfiedConditionalAccessCondition(input)
	return &out, nil
}
