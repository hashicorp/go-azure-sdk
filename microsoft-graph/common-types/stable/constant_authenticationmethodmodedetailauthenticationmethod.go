package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodModeDetailAuthenticationMethod string

const (
	AuthenticationMethodModeDetailAuthenticationMethod_Email                   AuthenticationMethodModeDetailAuthenticationMethod = "email"
	AuthenticationMethodModeDetailAuthenticationMethod_Federation              AuthenticationMethodModeDetailAuthenticationMethod = "federation"
	AuthenticationMethodModeDetailAuthenticationMethod_Fido2                   AuthenticationMethodModeDetailAuthenticationMethod = "fido2"
	AuthenticationMethodModeDetailAuthenticationMethod_HardwareOath            AuthenticationMethodModeDetailAuthenticationMethod = "hardwareOath"
	AuthenticationMethodModeDetailAuthenticationMethod_MicrosoftAuthenticator  AuthenticationMethodModeDetailAuthenticationMethod = "microsoftAuthenticator"
	AuthenticationMethodModeDetailAuthenticationMethod_Password                AuthenticationMethodModeDetailAuthenticationMethod = "password"
	AuthenticationMethodModeDetailAuthenticationMethod_Sms                     AuthenticationMethodModeDetailAuthenticationMethod = "sms"
	AuthenticationMethodModeDetailAuthenticationMethod_SoftwareOath            AuthenticationMethodModeDetailAuthenticationMethod = "softwareOath"
	AuthenticationMethodModeDetailAuthenticationMethod_TemporaryAccessPass     AuthenticationMethodModeDetailAuthenticationMethod = "temporaryAccessPass"
	AuthenticationMethodModeDetailAuthenticationMethod_Voice                   AuthenticationMethodModeDetailAuthenticationMethod = "voice"
	AuthenticationMethodModeDetailAuthenticationMethod_WindowsHelloForBusiness AuthenticationMethodModeDetailAuthenticationMethod = "windowsHelloForBusiness"
	AuthenticationMethodModeDetailAuthenticationMethod_X509Certificate         AuthenticationMethodModeDetailAuthenticationMethod = "x509Certificate"
)

func PossibleValuesForAuthenticationMethodModeDetailAuthenticationMethod() []string {
	return []string{
		string(AuthenticationMethodModeDetailAuthenticationMethod_Email),
		string(AuthenticationMethodModeDetailAuthenticationMethod_Federation),
		string(AuthenticationMethodModeDetailAuthenticationMethod_Fido2),
		string(AuthenticationMethodModeDetailAuthenticationMethod_HardwareOath),
		string(AuthenticationMethodModeDetailAuthenticationMethod_MicrosoftAuthenticator),
		string(AuthenticationMethodModeDetailAuthenticationMethod_Password),
		string(AuthenticationMethodModeDetailAuthenticationMethod_Sms),
		string(AuthenticationMethodModeDetailAuthenticationMethod_SoftwareOath),
		string(AuthenticationMethodModeDetailAuthenticationMethod_TemporaryAccessPass),
		string(AuthenticationMethodModeDetailAuthenticationMethod_Voice),
		string(AuthenticationMethodModeDetailAuthenticationMethod_WindowsHelloForBusiness),
		string(AuthenticationMethodModeDetailAuthenticationMethod_X509Certificate),
	}
}

func (s *AuthenticationMethodModeDetailAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodModeDetailAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodModeDetailAuthenticationMethod(input string) (*AuthenticationMethodModeDetailAuthenticationMethod, error) {
	vals := map[string]AuthenticationMethodModeDetailAuthenticationMethod{
		"email":                   AuthenticationMethodModeDetailAuthenticationMethod_Email,
		"federation":              AuthenticationMethodModeDetailAuthenticationMethod_Federation,
		"fido2":                   AuthenticationMethodModeDetailAuthenticationMethod_Fido2,
		"hardwareoath":            AuthenticationMethodModeDetailAuthenticationMethod_HardwareOath,
		"microsoftauthenticator":  AuthenticationMethodModeDetailAuthenticationMethod_MicrosoftAuthenticator,
		"password":                AuthenticationMethodModeDetailAuthenticationMethod_Password,
		"sms":                     AuthenticationMethodModeDetailAuthenticationMethod_Sms,
		"softwareoath":            AuthenticationMethodModeDetailAuthenticationMethod_SoftwareOath,
		"temporaryaccesspass":     AuthenticationMethodModeDetailAuthenticationMethod_TemporaryAccessPass,
		"voice":                   AuthenticationMethodModeDetailAuthenticationMethod_Voice,
		"windowshelloforbusiness": AuthenticationMethodModeDetailAuthenticationMethod_WindowsHelloForBusiness,
		"x509certificate":         AuthenticationMethodModeDetailAuthenticationMethod_X509Certificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodModeDetailAuthenticationMethod(input)
	return &out, nil
}
