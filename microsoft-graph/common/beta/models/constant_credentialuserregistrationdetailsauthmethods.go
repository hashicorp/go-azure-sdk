package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUserRegistrationDetailsAuthMethods string

const (
	CredentialUserRegistrationDetailsAuthMethodsalternateMobilePhone CredentialUserRegistrationDetailsAuthMethods = "AlternateMobilePhone"
	CredentialUserRegistrationDetailsAuthMethodsappCode              CredentialUserRegistrationDetailsAuthMethods = "AppCode"
	CredentialUserRegistrationDetailsAuthMethodsappNotification      CredentialUserRegistrationDetailsAuthMethods = "AppNotification"
	CredentialUserRegistrationDetailsAuthMethodsappPassword          CredentialUserRegistrationDetailsAuthMethods = "AppPassword"
	CredentialUserRegistrationDetailsAuthMethodsemail                CredentialUserRegistrationDetailsAuthMethods = "Email"
	CredentialUserRegistrationDetailsAuthMethodsfido                 CredentialUserRegistrationDetailsAuthMethods = "Fido"
	CredentialUserRegistrationDetailsAuthMethodsmobilePhone          CredentialUserRegistrationDetailsAuthMethods = "MobilePhone"
	CredentialUserRegistrationDetailsAuthMethodsofficePhone          CredentialUserRegistrationDetailsAuthMethods = "OfficePhone"
	CredentialUserRegistrationDetailsAuthMethodssecurityQuestion     CredentialUserRegistrationDetailsAuthMethods = "SecurityQuestion"
)

func PossibleValuesForCredentialUserRegistrationDetailsAuthMethods() []string {
	return []string{
		string(CredentialUserRegistrationDetailsAuthMethodsalternateMobilePhone),
		string(CredentialUserRegistrationDetailsAuthMethodsappCode),
		string(CredentialUserRegistrationDetailsAuthMethodsappNotification),
		string(CredentialUserRegistrationDetailsAuthMethodsappPassword),
		string(CredentialUserRegistrationDetailsAuthMethodsemail),
		string(CredentialUserRegistrationDetailsAuthMethodsfido),
		string(CredentialUserRegistrationDetailsAuthMethodsmobilePhone),
		string(CredentialUserRegistrationDetailsAuthMethodsofficePhone),
		string(CredentialUserRegistrationDetailsAuthMethodssecurityQuestion),
	}
}

func parseCredentialUserRegistrationDetailsAuthMethods(input string) (*CredentialUserRegistrationDetailsAuthMethods, error) {
	vals := map[string]CredentialUserRegistrationDetailsAuthMethods{
		"alternatemobilephone": CredentialUserRegistrationDetailsAuthMethodsalternateMobilePhone,
		"appcode":              CredentialUserRegistrationDetailsAuthMethodsappCode,
		"appnotification":      CredentialUserRegistrationDetailsAuthMethodsappNotification,
		"apppassword":          CredentialUserRegistrationDetailsAuthMethodsappPassword,
		"email":                CredentialUserRegistrationDetailsAuthMethodsemail,
		"fido":                 CredentialUserRegistrationDetailsAuthMethodsfido,
		"mobilephone":          CredentialUserRegistrationDetailsAuthMethodsmobilePhone,
		"officephone":          CredentialUserRegistrationDetailsAuthMethodsofficePhone,
		"securityquestion":     CredentialUserRegistrationDetailsAuthMethodssecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialUserRegistrationDetailsAuthMethods(input)
	return &out, nil
}
