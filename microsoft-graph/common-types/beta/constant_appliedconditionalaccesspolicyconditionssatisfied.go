package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicyConditionsSatisfied string

const (
	AppliedConditionalAccessPolicyConditionsSatisfied_Application                     AppliedConditionalAccessPolicyConditionsSatisfied = "application"
	AppliedConditionalAccessPolicyConditionsSatisfied_AuthenticationFlows             AppliedConditionalAccessPolicyConditionsSatisfied = "authenticationFlows"
	AppliedConditionalAccessPolicyConditionsSatisfied_Client                          AppliedConditionalAccessPolicyConditionsSatisfied = "client"
	AppliedConditionalAccessPolicyConditionsSatisfied_ClientType                      AppliedConditionalAccessPolicyConditionsSatisfied = "clientType"
	AppliedConditionalAccessPolicyConditionsSatisfied_DevicePlatform                  AppliedConditionalAccessPolicyConditionsSatisfied = "devicePlatform"
	AppliedConditionalAccessPolicyConditionsSatisfied_DeviceState                     AppliedConditionalAccessPolicyConditionsSatisfied = "deviceState"
	AppliedConditionalAccessPolicyConditionsSatisfied_InsiderRisk                     AppliedConditionalAccessPolicyConditionsSatisfied = "insiderRisk"
	AppliedConditionalAccessPolicyConditionsSatisfied_IpAddressSeenByAzureAD          AppliedConditionalAccessPolicyConditionsSatisfied = "ipAddressSeenByAzureAD"
	AppliedConditionalAccessPolicyConditionsSatisfied_IpAddressSeenByResourceProvider AppliedConditionalAccessPolicyConditionsSatisfied = "ipAddressSeenByResourceProvider"
	AppliedConditionalAccessPolicyConditionsSatisfied_Location                        AppliedConditionalAccessPolicyConditionsSatisfied = "location"
	AppliedConditionalAccessPolicyConditionsSatisfied_None                            AppliedConditionalAccessPolicyConditionsSatisfied = "none"
	AppliedConditionalAccessPolicyConditionsSatisfied_ServicePrincipalRisk            AppliedConditionalAccessPolicyConditionsSatisfied = "servicePrincipalRisk"
	AppliedConditionalAccessPolicyConditionsSatisfied_ServicePrincipals               AppliedConditionalAccessPolicyConditionsSatisfied = "servicePrincipals"
	AppliedConditionalAccessPolicyConditionsSatisfied_SignInRisk                      AppliedConditionalAccessPolicyConditionsSatisfied = "signInRisk"
	AppliedConditionalAccessPolicyConditionsSatisfied_Time                            AppliedConditionalAccessPolicyConditionsSatisfied = "time"
	AppliedConditionalAccessPolicyConditionsSatisfied_UserRisk                        AppliedConditionalAccessPolicyConditionsSatisfied = "userRisk"
	AppliedConditionalAccessPolicyConditionsSatisfied_Users                           AppliedConditionalAccessPolicyConditionsSatisfied = "users"
)

func PossibleValuesForAppliedConditionalAccessPolicyConditionsSatisfied() []string {
	return []string{
		string(AppliedConditionalAccessPolicyConditionsSatisfied_Application),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_AuthenticationFlows),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_Client),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_ClientType),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_DevicePlatform),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_DeviceState),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_InsiderRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_IpAddressSeenByAzureAD),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_IpAddressSeenByResourceProvider),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_Location),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_None),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_ServicePrincipalRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_ServicePrincipals),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_SignInRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_Time),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_UserRisk),
		string(AppliedConditionalAccessPolicyConditionsSatisfied_Users),
	}
}

func (s *AppliedConditionalAccessPolicyConditionsSatisfied) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppliedConditionalAccessPolicyConditionsSatisfied(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppliedConditionalAccessPolicyConditionsSatisfied(input string) (*AppliedConditionalAccessPolicyConditionsSatisfied, error) {
	vals := map[string]AppliedConditionalAccessPolicyConditionsSatisfied{
		"application":                     AppliedConditionalAccessPolicyConditionsSatisfied_Application,
		"authenticationflows":             AppliedConditionalAccessPolicyConditionsSatisfied_AuthenticationFlows,
		"client":                          AppliedConditionalAccessPolicyConditionsSatisfied_Client,
		"clienttype":                      AppliedConditionalAccessPolicyConditionsSatisfied_ClientType,
		"deviceplatform":                  AppliedConditionalAccessPolicyConditionsSatisfied_DevicePlatform,
		"devicestate":                     AppliedConditionalAccessPolicyConditionsSatisfied_DeviceState,
		"insiderrisk":                     AppliedConditionalAccessPolicyConditionsSatisfied_InsiderRisk,
		"ipaddressseenbyazuread":          AppliedConditionalAccessPolicyConditionsSatisfied_IpAddressSeenByAzureAD,
		"ipaddressseenbyresourceprovider": AppliedConditionalAccessPolicyConditionsSatisfied_IpAddressSeenByResourceProvider,
		"location":                        AppliedConditionalAccessPolicyConditionsSatisfied_Location,
		"none":                            AppliedConditionalAccessPolicyConditionsSatisfied_None,
		"serviceprincipalrisk":            AppliedConditionalAccessPolicyConditionsSatisfied_ServicePrincipalRisk,
		"serviceprincipals":               AppliedConditionalAccessPolicyConditionsSatisfied_ServicePrincipals,
		"signinrisk":                      AppliedConditionalAccessPolicyConditionsSatisfied_SignInRisk,
		"time":                            AppliedConditionalAccessPolicyConditionsSatisfied_Time,
		"userrisk":                        AppliedConditionalAccessPolicyConditionsSatisfied_UserRisk,
		"users":                           AppliedConditionalAccessPolicyConditionsSatisfied_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedConditionalAccessPolicyConditionsSatisfied(input)
	return &out, nil
}
