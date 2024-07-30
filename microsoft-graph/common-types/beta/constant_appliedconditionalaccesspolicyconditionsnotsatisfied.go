package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicyConditionsNotSatisfied string

const (
	AppliedConditionalAccessPolicyConditionsNotSatisfied_Application                     AppliedConditionalAccessPolicyConditionsNotSatisfied = "application"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_AuthenticationFlows             AppliedConditionalAccessPolicyConditionsNotSatisfied = "authenticationFlows"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_Client                          AppliedConditionalAccessPolicyConditionsNotSatisfied = "client"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_ClientType                      AppliedConditionalAccessPolicyConditionsNotSatisfied = "clientType"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_DevicePlatform                  AppliedConditionalAccessPolicyConditionsNotSatisfied = "devicePlatform"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_DeviceState                     AppliedConditionalAccessPolicyConditionsNotSatisfied = "deviceState"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_InsiderRisk                     AppliedConditionalAccessPolicyConditionsNotSatisfied = "insiderRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_IpAddressSeenByAzureAD          AppliedConditionalAccessPolicyConditionsNotSatisfied = "ipAddressSeenByAzureAD"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_IpAddressSeenByResourceProvider AppliedConditionalAccessPolicyConditionsNotSatisfied = "ipAddressSeenByResourceProvider"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_Location                        AppliedConditionalAccessPolicyConditionsNotSatisfied = "location"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_None                            AppliedConditionalAccessPolicyConditionsNotSatisfied = "none"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_ServicePrincipalRisk            AppliedConditionalAccessPolicyConditionsNotSatisfied = "servicePrincipalRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_ServicePrincipals               AppliedConditionalAccessPolicyConditionsNotSatisfied = "servicePrincipals"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_SignInRisk                      AppliedConditionalAccessPolicyConditionsNotSatisfied = "signInRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_Time                            AppliedConditionalAccessPolicyConditionsNotSatisfied = "time"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_UserRisk                        AppliedConditionalAccessPolicyConditionsNotSatisfied = "userRisk"
	AppliedConditionalAccessPolicyConditionsNotSatisfied_Users                           AppliedConditionalAccessPolicyConditionsNotSatisfied = "users"
)

func PossibleValuesForAppliedConditionalAccessPolicyConditionsNotSatisfied() []string {
	return []string{
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_Application),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_AuthenticationFlows),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_Client),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_ClientType),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_DevicePlatform),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_DeviceState),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_InsiderRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_IpAddressSeenByAzureAD),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_IpAddressSeenByResourceProvider),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_Location),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_None),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_ServicePrincipalRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_ServicePrincipals),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_SignInRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_Time),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_UserRisk),
		string(AppliedConditionalAccessPolicyConditionsNotSatisfied_Users),
	}
}

func (s *AppliedConditionalAccessPolicyConditionsNotSatisfied) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppliedConditionalAccessPolicyConditionsNotSatisfied(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppliedConditionalAccessPolicyConditionsNotSatisfied(input string) (*AppliedConditionalAccessPolicyConditionsNotSatisfied, error) {
	vals := map[string]AppliedConditionalAccessPolicyConditionsNotSatisfied{
		"application":                     AppliedConditionalAccessPolicyConditionsNotSatisfied_Application,
		"authenticationflows":             AppliedConditionalAccessPolicyConditionsNotSatisfied_AuthenticationFlows,
		"client":                          AppliedConditionalAccessPolicyConditionsNotSatisfied_Client,
		"clienttype":                      AppliedConditionalAccessPolicyConditionsNotSatisfied_ClientType,
		"deviceplatform":                  AppliedConditionalAccessPolicyConditionsNotSatisfied_DevicePlatform,
		"devicestate":                     AppliedConditionalAccessPolicyConditionsNotSatisfied_DeviceState,
		"insiderrisk":                     AppliedConditionalAccessPolicyConditionsNotSatisfied_InsiderRisk,
		"ipaddressseenbyazuread":          AppliedConditionalAccessPolicyConditionsNotSatisfied_IpAddressSeenByAzureAD,
		"ipaddressseenbyresourceprovider": AppliedConditionalAccessPolicyConditionsNotSatisfied_IpAddressSeenByResourceProvider,
		"location":                        AppliedConditionalAccessPolicyConditionsNotSatisfied_Location,
		"none":                            AppliedConditionalAccessPolicyConditionsNotSatisfied_None,
		"serviceprincipalrisk":            AppliedConditionalAccessPolicyConditionsNotSatisfied_ServicePrincipalRisk,
		"serviceprincipals":               AppliedConditionalAccessPolicyConditionsNotSatisfied_ServicePrincipals,
		"signinrisk":                      AppliedConditionalAccessPolicyConditionsNotSatisfied_SignInRisk,
		"time":                            AppliedConditionalAccessPolicyConditionsNotSatisfied_Time,
		"userrisk":                        AppliedConditionalAccessPolicyConditionsNotSatisfied_UserRisk,
		"users":                           AppliedConditionalAccessPolicyConditionsNotSatisfied_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedConditionalAccessPolicyConditionsNotSatisfied(input)
	return &out, nil
}
