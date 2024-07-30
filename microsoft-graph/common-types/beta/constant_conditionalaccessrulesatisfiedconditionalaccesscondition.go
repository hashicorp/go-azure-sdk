package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRuleSatisfiedConditionalAccessCondition string

const (
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_Application                     ConditionalAccessRuleSatisfiedConditionalAccessCondition = "application"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_AuthenticationFlows             ConditionalAccessRuleSatisfiedConditionalAccessCondition = "authenticationFlows"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_Client                          ConditionalAccessRuleSatisfiedConditionalAccessCondition = "client"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_ClientType                      ConditionalAccessRuleSatisfiedConditionalAccessCondition = "clientType"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_DevicePlatform                  ConditionalAccessRuleSatisfiedConditionalAccessCondition = "devicePlatform"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_DeviceState                     ConditionalAccessRuleSatisfiedConditionalAccessCondition = "deviceState"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_InsiderRisk                     ConditionalAccessRuleSatisfiedConditionalAccessCondition = "insiderRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_IpAddressSeenByAzureAD          ConditionalAccessRuleSatisfiedConditionalAccessCondition = "ipAddressSeenByAzureAD"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_IpAddressSeenByResourceProvider ConditionalAccessRuleSatisfiedConditionalAccessCondition = "ipAddressSeenByResourceProvider"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_Location                        ConditionalAccessRuleSatisfiedConditionalAccessCondition = "location"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_None                            ConditionalAccessRuleSatisfiedConditionalAccessCondition = "none"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_ServicePrincipalRisk            ConditionalAccessRuleSatisfiedConditionalAccessCondition = "servicePrincipalRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_ServicePrincipals               ConditionalAccessRuleSatisfiedConditionalAccessCondition = "servicePrincipals"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_SignInRisk                      ConditionalAccessRuleSatisfiedConditionalAccessCondition = "signInRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_Time                            ConditionalAccessRuleSatisfiedConditionalAccessCondition = "time"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_UserRisk                        ConditionalAccessRuleSatisfiedConditionalAccessCondition = "userRisk"
	ConditionalAccessRuleSatisfiedConditionalAccessCondition_Users                           ConditionalAccessRuleSatisfiedConditionalAccessCondition = "users"
)

func PossibleValuesForConditionalAccessRuleSatisfiedConditionalAccessCondition() []string {
	return []string{
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_Application),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_AuthenticationFlows),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_Client),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_ClientType),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_DevicePlatform),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_DeviceState),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_InsiderRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_IpAddressSeenByAzureAD),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_IpAddressSeenByResourceProvider),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_Location),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_None),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_ServicePrincipalRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_ServicePrincipals),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_SignInRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_Time),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_UserRisk),
		string(ConditionalAccessRuleSatisfiedConditionalAccessCondition_Users),
	}
}

func (s *ConditionalAccessRuleSatisfiedConditionalAccessCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessRuleSatisfiedConditionalAccessCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessRuleSatisfiedConditionalAccessCondition(input string) (*ConditionalAccessRuleSatisfiedConditionalAccessCondition, error) {
	vals := map[string]ConditionalAccessRuleSatisfiedConditionalAccessCondition{
		"application":                     ConditionalAccessRuleSatisfiedConditionalAccessCondition_Application,
		"authenticationflows":             ConditionalAccessRuleSatisfiedConditionalAccessCondition_AuthenticationFlows,
		"client":                          ConditionalAccessRuleSatisfiedConditionalAccessCondition_Client,
		"clienttype":                      ConditionalAccessRuleSatisfiedConditionalAccessCondition_ClientType,
		"deviceplatform":                  ConditionalAccessRuleSatisfiedConditionalAccessCondition_DevicePlatform,
		"devicestate":                     ConditionalAccessRuleSatisfiedConditionalAccessCondition_DeviceState,
		"insiderrisk":                     ConditionalAccessRuleSatisfiedConditionalAccessCondition_InsiderRisk,
		"ipaddressseenbyazuread":          ConditionalAccessRuleSatisfiedConditionalAccessCondition_IpAddressSeenByAzureAD,
		"ipaddressseenbyresourceprovider": ConditionalAccessRuleSatisfiedConditionalAccessCondition_IpAddressSeenByResourceProvider,
		"location":                        ConditionalAccessRuleSatisfiedConditionalAccessCondition_Location,
		"none":                            ConditionalAccessRuleSatisfiedConditionalAccessCondition_None,
		"serviceprincipalrisk":            ConditionalAccessRuleSatisfiedConditionalAccessCondition_ServicePrincipalRisk,
		"serviceprincipals":               ConditionalAccessRuleSatisfiedConditionalAccessCondition_ServicePrincipals,
		"signinrisk":                      ConditionalAccessRuleSatisfiedConditionalAccessCondition_SignInRisk,
		"time":                            ConditionalAccessRuleSatisfiedConditionalAccessCondition_Time,
		"userrisk":                        ConditionalAccessRuleSatisfiedConditionalAccessCondition_UserRisk,
		"users":                           ConditionalAccessRuleSatisfiedConditionalAccessCondition_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessRuleSatisfiedConditionalAccessCondition(input)
	return &out, nil
}
