package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUserRegistrationDetailsAuthMethods string

const (
	CredentialUserRegistrationDetailsAuthMethods_AlternateMobilePhone CredentialUserRegistrationDetailsAuthMethods = "alternateMobilePhone"
	CredentialUserRegistrationDetailsAuthMethods_AppCode              CredentialUserRegistrationDetailsAuthMethods = "appCode"
	CredentialUserRegistrationDetailsAuthMethods_AppNotification      CredentialUserRegistrationDetailsAuthMethods = "appNotification"
	CredentialUserRegistrationDetailsAuthMethods_AppPassword          CredentialUserRegistrationDetailsAuthMethods = "appPassword"
	CredentialUserRegistrationDetailsAuthMethods_Email                CredentialUserRegistrationDetailsAuthMethods = "email"
	CredentialUserRegistrationDetailsAuthMethods_Fido                 CredentialUserRegistrationDetailsAuthMethods = "fido"
	CredentialUserRegistrationDetailsAuthMethods_MobilePhone          CredentialUserRegistrationDetailsAuthMethods = "mobilePhone"
	CredentialUserRegistrationDetailsAuthMethods_OfficePhone          CredentialUserRegistrationDetailsAuthMethods = "officePhone"
	CredentialUserRegistrationDetailsAuthMethods_SecurityQuestion     CredentialUserRegistrationDetailsAuthMethods = "securityQuestion"
)

func PossibleValuesForCredentialUserRegistrationDetailsAuthMethods() []string {
	return []string{
		string(CredentialUserRegistrationDetailsAuthMethods_AlternateMobilePhone),
		string(CredentialUserRegistrationDetailsAuthMethods_AppCode),
		string(CredentialUserRegistrationDetailsAuthMethods_AppNotification),
		string(CredentialUserRegistrationDetailsAuthMethods_AppPassword),
		string(CredentialUserRegistrationDetailsAuthMethods_Email),
		string(CredentialUserRegistrationDetailsAuthMethods_Fido),
		string(CredentialUserRegistrationDetailsAuthMethods_MobilePhone),
		string(CredentialUserRegistrationDetailsAuthMethods_OfficePhone),
		string(CredentialUserRegistrationDetailsAuthMethods_SecurityQuestion),
	}
}

func (s *CredentialUserRegistrationDetailsAuthMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCredentialUserRegistrationDetailsAuthMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCredentialUserRegistrationDetailsAuthMethods(input string) (*CredentialUserRegistrationDetailsAuthMethods, error) {
	vals := map[string]CredentialUserRegistrationDetailsAuthMethods{
		"alternatemobilephone": CredentialUserRegistrationDetailsAuthMethods_AlternateMobilePhone,
		"appcode":              CredentialUserRegistrationDetailsAuthMethods_AppCode,
		"appnotification":      CredentialUserRegistrationDetailsAuthMethods_AppNotification,
		"apppassword":          CredentialUserRegistrationDetailsAuthMethods_AppPassword,
		"email":                CredentialUserRegistrationDetailsAuthMethods_Email,
		"fido":                 CredentialUserRegistrationDetailsAuthMethods_Fido,
		"mobilephone":          CredentialUserRegistrationDetailsAuthMethods_MobilePhone,
		"officephone":          CredentialUserRegistrationDetailsAuthMethods_OfficePhone,
		"securityquestion":     CredentialUserRegistrationDetailsAuthMethods_SecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialUserRegistrationDetailsAuthMethods(input)
	return &out, nil
}
