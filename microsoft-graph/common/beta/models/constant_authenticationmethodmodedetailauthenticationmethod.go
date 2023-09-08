package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodModeDetailAuthenticationMethod string

const (
	AuthenticationMethodModeDetailAuthenticationMethodemail                   AuthenticationMethodModeDetailAuthenticationMethod = "Email"
	AuthenticationMethodModeDetailAuthenticationMethodfederation              AuthenticationMethodModeDetailAuthenticationMethod = "Federation"
	AuthenticationMethodModeDetailAuthenticationMethodfido2                   AuthenticationMethodModeDetailAuthenticationMethod = "Fido2"
	AuthenticationMethodModeDetailAuthenticationMethodhardwareOath            AuthenticationMethodModeDetailAuthenticationMethod = "HardwareOath"
	AuthenticationMethodModeDetailAuthenticationMethodmicrosoftAuthenticator  AuthenticationMethodModeDetailAuthenticationMethod = "MicrosoftAuthenticator"
	AuthenticationMethodModeDetailAuthenticationMethodpassword                AuthenticationMethodModeDetailAuthenticationMethod = "Password"
	AuthenticationMethodModeDetailAuthenticationMethodsms                     AuthenticationMethodModeDetailAuthenticationMethod = "Sms"
	AuthenticationMethodModeDetailAuthenticationMethodsoftwareOath            AuthenticationMethodModeDetailAuthenticationMethod = "SoftwareOath"
	AuthenticationMethodModeDetailAuthenticationMethodtemporaryAccessPass     AuthenticationMethodModeDetailAuthenticationMethod = "TemporaryAccessPass"
	AuthenticationMethodModeDetailAuthenticationMethodvoice                   AuthenticationMethodModeDetailAuthenticationMethod = "Voice"
	AuthenticationMethodModeDetailAuthenticationMethodwindowsHelloForBusiness AuthenticationMethodModeDetailAuthenticationMethod = "WindowsHelloForBusiness"
	AuthenticationMethodModeDetailAuthenticationMethodx509Certificate         AuthenticationMethodModeDetailAuthenticationMethod = "X509Certificate"
)

func PossibleValuesForAuthenticationMethodModeDetailAuthenticationMethod() []string {
	return []string{
		string(AuthenticationMethodModeDetailAuthenticationMethodemail),
		string(AuthenticationMethodModeDetailAuthenticationMethodfederation),
		string(AuthenticationMethodModeDetailAuthenticationMethodfido2),
		string(AuthenticationMethodModeDetailAuthenticationMethodhardwareOath),
		string(AuthenticationMethodModeDetailAuthenticationMethodmicrosoftAuthenticator),
		string(AuthenticationMethodModeDetailAuthenticationMethodpassword),
		string(AuthenticationMethodModeDetailAuthenticationMethodsms),
		string(AuthenticationMethodModeDetailAuthenticationMethodsoftwareOath),
		string(AuthenticationMethodModeDetailAuthenticationMethodtemporaryAccessPass),
		string(AuthenticationMethodModeDetailAuthenticationMethodvoice),
		string(AuthenticationMethodModeDetailAuthenticationMethodwindowsHelloForBusiness),
		string(AuthenticationMethodModeDetailAuthenticationMethodx509Certificate),
	}
}

func parseAuthenticationMethodModeDetailAuthenticationMethod(input string) (*AuthenticationMethodModeDetailAuthenticationMethod, error) {
	vals := map[string]AuthenticationMethodModeDetailAuthenticationMethod{
		"email":                   AuthenticationMethodModeDetailAuthenticationMethodemail,
		"federation":              AuthenticationMethodModeDetailAuthenticationMethodfederation,
		"fido2":                   AuthenticationMethodModeDetailAuthenticationMethodfido2,
		"hardwareoath":            AuthenticationMethodModeDetailAuthenticationMethodhardwareOath,
		"microsoftauthenticator":  AuthenticationMethodModeDetailAuthenticationMethodmicrosoftAuthenticator,
		"password":                AuthenticationMethodModeDetailAuthenticationMethodpassword,
		"sms":                     AuthenticationMethodModeDetailAuthenticationMethodsms,
		"softwareoath":            AuthenticationMethodModeDetailAuthenticationMethodsoftwareOath,
		"temporaryaccesspass":     AuthenticationMethodModeDetailAuthenticationMethodtemporaryAccessPass,
		"voice":                   AuthenticationMethodModeDetailAuthenticationMethodvoice,
		"windowshelloforbusiness": AuthenticationMethodModeDetailAuthenticationMethodwindowsHelloForBusiness,
		"x509certificate":         AuthenticationMethodModeDetailAuthenticationMethodx509Certificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodModeDetailAuthenticationMethod(input)
	return &out, nil
}
