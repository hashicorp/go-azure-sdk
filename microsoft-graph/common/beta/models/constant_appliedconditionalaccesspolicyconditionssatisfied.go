package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicyConditionsSatisfied string

const (
	AppliedConditionalAccessPolicyConditionsSatisfiedapplication                     AppliedConditionalAccessPolicyConditionsSatisfied = "Application"
	AppliedConditionalAccessPolicyConditionsSatisfiedauthenticationFlows             AppliedConditionalAccessPolicyConditionsSatisfied = "AuthenticationFlows"
	AppliedConditionalAccessPolicyConditionsSatisfiedclient                          AppliedConditionalAccessPolicyConditionsSatisfied = "Client"
	AppliedConditionalAccessPolicyConditionsSatisfiedclientType                      AppliedConditionalAccessPolicyConditionsSatisfied = "ClientType"
	AppliedConditionalAccessPolicyConditionsSatisfieddevicePlatform                  AppliedConditionalAccessPolicyConditionsSatisfied = "DevicePlatform"
	AppliedConditionalAccessPolicyConditionsSatisfieddeviceState                     AppliedConditionalAccessPolicyConditionsSatisfied = "DeviceState"
	AppliedConditionalAccessPolicyConditionsSatisfiedinsiderRisk                     AppliedConditionalAccessPolicyConditionsSatisfied = "InsiderRisk"
	AppliedConditionalAccessPolicyConditionsSatisfiedipAddressSeenByAzureAD          AppliedConditionalAccessPolicyConditionsSatisfied = "IpAddressSeenByAzureAD"
	AppliedConditionalAccessPolicyConditionsSatisfiedipAddressSeenByResourceProvider AppliedConditionalAccessPolicyConditionsSatisfied = "IpAddressSeenByResourceProvider"
	AppliedConditionalAccessPolicyConditionsSatisfiedlocation                        AppliedConditionalAccessPolicyConditionsSatisfied = "Location"
	AppliedConditionalAccessPolicyConditionsSatisfiednone                            AppliedConditionalAccessPolicyConditionsSatisfied = "None"
	AppliedConditionalAccessPolicyConditionsSatisfiedservicePrincipalRisk            AppliedConditionalAccessPolicyConditionsSatisfied = "ServicePrincipalRisk"
	AppliedConditionalAccessPolicyConditionsSatisfiedservicePrincipals               AppliedConditionalAccessPolicyConditionsSatisfied = "ServicePrincipals"
	AppliedConditionalAccessPolicyConditionsSatisfiedsignInRisk                      AppliedConditionalAccessPolicyConditionsSatisfied = "SignInRisk"
	AppliedConditionalAccessPolicyConditionsSatisfiedtime                            AppliedConditionalAccessPolicyConditionsSatisfied = "Time"
	AppliedConditionalAccessPolicyConditionsSatisfieduserRisk                        AppliedConditionalAccessPolicyConditionsSatisfied = "UserRisk"
	AppliedConditionalAccessPolicyConditionsSatisfiedusers                           AppliedConditionalAccessPolicyConditionsSatisfied = "Users"
)

func PossibleValuesForAppliedConditionalAccessPolicyConditionsSatisfied() []string {
	return []string{
		string(AppliedConditionalAccessPolicyConditionsSatisfiedapplication),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedauthenticationFlows),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedclient),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedclientType),
		string(AppliedConditionalAccessPolicyConditionsSatisfieddevicePlatform),
		string(AppliedConditionalAccessPolicyConditionsSatisfieddeviceState),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedinsiderRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedipAddressSeenByAzureAD),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedipAddressSeenByResourceProvider),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedlocation),
		string(AppliedConditionalAccessPolicyConditionsSatisfiednone),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedservicePrincipalRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedservicePrincipals),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedsignInRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedtime),
		string(AppliedConditionalAccessPolicyConditionsSatisfieduserRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfiedusers),
	}
}

func parseAppliedConditionalAccessPolicyConditionsSatisfied(input string) (*AppliedConditionalAccessPolicyConditionsSatisfied, error) {
	vals := map[string]AppliedConditionalAccessPolicyConditionsSatisfied{
		"application":                     AppliedConditionalAccessPolicyConditionsSatisfiedapplication,
		"authenticationflows":             AppliedConditionalAccessPolicyConditionsSatisfiedauthenticationFlows,
		"client":                          AppliedConditionalAccessPolicyConditionsSatisfiedclient,
		"clienttype":                      AppliedConditionalAccessPolicyConditionsSatisfiedclientType,
		"deviceplatform":                  AppliedConditionalAccessPolicyConditionsSatisfieddevicePlatform,
		"devicestate":                     AppliedConditionalAccessPolicyConditionsSatisfieddeviceState,
		"insiderrisk":                     AppliedConditionalAccessPolicyConditionsSatisfiedinsiderRisk,
		"ipaddressseenbyazuread":          AppliedConditionalAccessPolicyConditionsSatisfiedipAddressSeenByAzureAD,
		"ipaddressseenbyresourceprovider": AppliedConditionalAccessPolicyConditionsSatisfiedipAddressSeenByResourceProvider,
		"location":                        AppliedConditionalAccessPolicyConditionsSatisfiedlocation,
		"none":                            AppliedConditionalAccessPolicyConditionsSatisfiednone,
		"serviceprincipalrisk":            AppliedConditionalAccessPolicyConditionsSatisfiedservicePrincipalRisk,
		"serviceprincipals":               AppliedConditionalAccessPolicyConditionsSatisfiedservicePrincipals,
		"signinrisk":                      AppliedConditionalAccessPolicyConditionsSatisfiedsignInRisk,
		"time":                            AppliedConditionalAccessPolicyConditionsSatisfiedtime,
		"userrisk":                        AppliedConditionalAccessPolicyConditionsSatisfieduserRisk,
		"users":                           AppliedConditionalAccessPolicyConditionsSatisfiedusers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedConditionalAccessPolicyConditionsSatisfied(input)
	return &out, nil
}
