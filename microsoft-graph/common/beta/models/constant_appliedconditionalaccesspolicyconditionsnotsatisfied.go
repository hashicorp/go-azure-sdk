package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicyConditionsNotSatisfied string

const (
	AppliedConditionalAccessPolicyConditionsNotSatisfiedapplication                     AppliedConditionalAccessPolicyConditionsNotSatisfied = "Application"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedauthenticationFlows             AppliedConditionalAccessPolicyConditionsNotSatisfied = "AuthenticationFlows"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedclient                          AppliedConditionalAccessPolicyConditionsNotSatisfied = "Client"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedclientType                      AppliedConditionalAccessPolicyConditionsNotSatisfied = "ClientType"
	AppliedConditionalAccessPolicyConditionsNotSatisfieddevicePlatform                  AppliedConditionalAccessPolicyConditionsNotSatisfied = "DevicePlatform"
	AppliedConditionalAccessPolicyConditionsNotSatisfieddeviceState                     AppliedConditionalAccessPolicyConditionsNotSatisfied = "DeviceState"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedinsiderRisk                     AppliedConditionalAccessPolicyConditionsNotSatisfied = "InsiderRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedipAddressSeenByAzureAD          AppliedConditionalAccessPolicyConditionsNotSatisfied = "IpAddressSeenByAzureAD"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedipAddressSeenByResourceProvider AppliedConditionalAccessPolicyConditionsNotSatisfied = "IpAddressSeenByResourceProvider"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedlocation                        AppliedConditionalAccessPolicyConditionsNotSatisfied = "Location"
	AppliedConditionalAccessPolicyConditionsNotSatisfiednone                            AppliedConditionalAccessPolicyConditionsNotSatisfied = "None"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedservicePrincipalRisk            AppliedConditionalAccessPolicyConditionsNotSatisfied = "ServicePrincipalRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedservicePrincipals               AppliedConditionalAccessPolicyConditionsNotSatisfied = "ServicePrincipals"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedsignInRisk                      AppliedConditionalAccessPolicyConditionsNotSatisfied = "SignInRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedtime                            AppliedConditionalAccessPolicyConditionsNotSatisfied = "Time"
	AppliedConditionalAccessPolicyConditionsNotSatisfieduserRisk                        AppliedConditionalAccessPolicyConditionsNotSatisfied = "UserRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfiedusers                           AppliedConditionalAccessPolicyConditionsNotSatisfied = "Users"
)

func PossibleValuesForAppliedConditionalAccessPolicyConditionsNotSatisfied() []string {
	return []string{
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedapplication),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedauthenticationFlows),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedclient),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedclientType),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfieddevicePlatform),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfieddeviceState),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedinsiderRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedipAddressSeenByAzureAD),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedipAddressSeenByResourceProvider),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedlocation),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiednone),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedservicePrincipalRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedservicePrincipals),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedsignInRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedtime),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfieduserRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfiedusers),
	}
}

func parseAppliedConditionalAccessPolicyConditionsNotSatisfied(input string) (*AppliedConditionalAccessPolicyConditionsNotSatisfied, error) {
	vals := map[string]AppliedConditionalAccessPolicyConditionsNotSatisfied{
		"application":                     AppliedConditionalAccessPolicyConditionsNotSatisfiedapplication,
		"authenticationflows":             AppliedConditionalAccessPolicyConditionsNotSatisfiedauthenticationFlows,
		"client":                          AppliedConditionalAccessPolicyConditionsNotSatisfiedclient,
		"clienttype":                      AppliedConditionalAccessPolicyConditionsNotSatisfiedclientType,
		"deviceplatform":                  AppliedConditionalAccessPolicyConditionsNotSatisfieddevicePlatform,
		"devicestate":                     AppliedConditionalAccessPolicyConditionsNotSatisfieddeviceState,
		"insiderrisk":                     AppliedConditionalAccessPolicyConditionsNotSatisfiedinsiderRisk,
		"ipaddressseenbyazuread":          AppliedConditionalAccessPolicyConditionsNotSatisfiedipAddressSeenByAzureAD,
		"ipaddressseenbyresourceprovider": AppliedConditionalAccessPolicyConditionsNotSatisfiedipAddressSeenByResourceProvider,
		"location":                        AppliedConditionalAccessPolicyConditionsNotSatisfiedlocation,
		"none":                            AppliedConditionalAccessPolicyConditionsNotSatisfiednone,
		"serviceprincipalrisk":            AppliedConditionalAccessPolicyConditionsNotSatisfiedservicePrincipalRisk,
		"serviceprincipals":               AppliedConditionalAccessPolicyConditionsNotSatisfiedservicePrincipals,
		"signinrisk":                      AppliedConditionalAccessPolicyConditionsNotSatisfiedsignInRisk,
		"time":                            AppliedConditionalAccessPolicyConditionsNotSatisfiedtime,
		"userrisk":                        AppliedConditionalAccessPolicyConditionsNotSatisfieduserRisk,
		"users":                           AppliedConditionalAccessPolicyConditionsNotSatisfiedusers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedConditionalAccessPolicyConditionsNotSatisfied(input)
	return &out, nil
}
