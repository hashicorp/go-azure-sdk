package containerappsauthconfigs

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
