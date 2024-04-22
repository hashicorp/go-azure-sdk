package containerappsauthconfigs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClientCredentialMethod string

const (
	ClientCredentialMethodClientSecretPost ClientCredentialMethod = "ClientSecretPost"
)

func PossibleValuesForClientCredentialMethod() []string {
	return []string{
		string(ClientCredentialMethodClientSecretPost),
	}
}

func (s *ClientCredentialMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClientCredentialMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClientCredentialMethod(input string) (*ClientCredentialMethod, error) {
	vals := map[string]ClientCredentialMethod{
		"clientsecretpost": ClientCredentialMethodClientSecretPost,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClientCredentialMethod(input)
	return &out, nil
}

type CookieExpirationConvention string

const (
	CookieExpirationConventionFixedTime               CookieExpirationConvention = "FixedTime"
	CookieExpirationConventionIdentityProviderDerived CookieExpirationConvention = "IdentityProviderDerived"
)

func PossibleValuesForCookieExpirationConvention() []string {
	return []string{
		string(CookieExpirationConventionFixedTime),
		string(CookieExpirationConventionIdentityProviderDerived),
	}
}

func (s *CookieExpirationConvention) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCookieExpirationConvention(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCookieExpirationConvention(input string) (*CookieExpirationConvention, error) {
	vals := map[string]CookieExpirationConvention{
		"fixedtime":               CookieExpirationConventionFixedTime,
		"identityproviderderived": CookieExpirationConventionIdentityProviderDerived,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CookieExpirationConvention(input)
	return &out, nil
}

type ForwardProxyConvention string

const (
	ForwardProxyConventionCustom   ForwardProxyConvention = "Custom"
	ForwardProxyConventionNoProxy  ForwardProxyConvention = "NoProxy"
	ForwardProxyConventionStandard ForwardProxyConvention = "Standard"
)

func PossibleValuesForForwardProxyConvention() []string {
	return []string{
		string(ForwardProxyConventionCustom),
		string(ForwardProxyConventionNoProxy),
		string(ForwardProxyConventionStandard),
	}
}

func (s *ForwardProxyConvention) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseForwardProxyConvention(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseForwardProxyConvention(input string) (*ForwardProxyConvention, error) {
	vals := map[string]ForwardProxyConvention{
		"custom":   ForwardProxyConventionCustom,
		"noproxy":  ForwardProxyConventionNoProxy,
		"standard": ForwardProxyConventionStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForwardProxyConvention(input)
	return &out, nil
}

type UnauthenticatedClientActionV2 string

const (
	UnauthenticatedClientActionV2AllowAnonymous      UnauthenticatedClientActionV2 = "AllowAnonymous"
	UnauthenticatedClientActionV2RedirectToLoginPage UnauthenticatedClientActionV2 = "RedirectToLoginPage"
	UnauthenticatedClientActionV2ReturnFourZeroOne   UnauthenticatedClientActionV2 = "Return401"
	UnauthenticatedClientActionV2ReturnFourZeroThree UnauthenticatedClientActionV2 = "Return403"
)

func PossibleValuesForUnauthenticatedClientActionV2() []string {
	return []string{
		string(UnauthenticatedClientActionV2AllowAnonymous),
		string(UnauthenticatedClientActionV2RedirectToLoginPage),
		string(UnauthenticatedClientActionV2ReturnFourZeroOne),
		string(UnauthenticatedClientActionV2ReturnFourZeroThree),
	}
}

func (s *UnauthenticatedClientActionV2) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnauthenticatedClientActionV2(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnauthenticatedClientActionV2(input string) (*UnauthenticatedClientActionV2, error) {
	vals := map[string]UnauthenticatedClientActionV2{
		"allowanonymous":      UnauthenticatedClientActionV2AllowAnonymous,
		"redirecttologinpage": UnauthenticatedClientActionV2RedirectToLoginPage,
		"return401":           UnauthenticatedClientActionV2ReturnFourZeroOne,
		"return403":           UnauthenticatedClientActionV2ReturnFourZeroThree,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnauthenticatedClientActionV2(input)
	return &out, nil
}
