package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationDetailsDefaultMfaMethod string

const (
	UserRegistrationDetailsDefaultMfaMethod_AlternateMobilePhone       UserRegistrationDetailsDefaultMfaMethod = "alternateMobilePhone"
	UserRegistrationDetailsDefaultMfaMethod_MicrosoftAuthenticatorPush UserRegistrationDetailsDefaultMfaMethod = "microsoftAuthenticatorPush"
	UserRegistrationDetailsDefaultMfaMethod_MobilePhone                UserRegistrationDetailsDefaultMfaMethod = "mobilePhone"
	UserRegistrationDetailsDefaultMfaMethod_None                       UserRegistrationDetailsDefaultMfaMethod = "none"
	UserRegistrationDetailsDefaultMfaMethod_OfficePhone                UserRegistrationDetailsDefaultMfaMethod = "officePhone"
	UserRegistrationDetailsDefaultMfaMethod_SoftwareOneTimePasscode    UserRegistrationDetailsDefaultMfaMethod = "softwareOneTimePasscode"
)

func PossibleValuesForUserRegistrationDetailsDefaultMfaMethod() []string {
	return []string{
		string(UserRegistrationDetailsDefaultMfaMethod_AlternateMobilePhone),
		string(UserRegistrationDetailsDefaultMfaMethod_MicrosoftAuthenticatorPush),
		string(UserRegistrationDetailsDefaultMfaMethod_MobilePhone),
		string(UserRegistrationDetailsDefaultMfaMethod_None),
		string(UserRegistrationDetailsDefaultMfaMethod_OfficePhone),
		string(UserRegistrationDetailsDefaultMfaMethod_SoftwareOneTimePasscode),
	}
}

func (s *UserRegistrationDetailsDefaultMfaMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationDetailsDefaultMfaMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationDetailsDefaultMfaMethod(input string) (*UserRegistrationDetailsDefaultMfaMethod, error) {
	vals := map[string]UserRegistrationDetailsDefaultMfaMethod{
		"alternatemobilephone":       UserRegistrationDetailsDefaultMfaMethod_AlternateMobilePhone,
		"microsoftauthenticatorpush": UserRegistrationDetailsDefaultMfaMethod_MicrosoftAuthenticatorPush,
		"mobilephone":                UserRegistrationDetailsDefaultMfaMethod_MobilePhone,
		"none":                       UserRegistrationDetailsDefaultMfaMethod_None,
		"officephone":                UserRegistrationDetailsDefaultMfaMethod_OfficePhone,
		"softwareonetimepasscode":    UserRegistrationDetailsDefaultMfaMethod_SoftwareOneTimePasscode,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationDetailsDefaultMfaMethod(input)
	return &out, nil
}
