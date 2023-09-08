package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCredentialUsageDetailsAuthMethod string

const (
	UserCredentialUsageDetailsAuthMethodalternateMobileCall UserCredentialUsageDetailsAuthMethod = "AlternateMobileCall"
	UserCredentialUsageDetailsAuthMethodappCode             UserCredentialUsageDetailsAuthMethod = "AppCode"
	UserCredentialUsageDetailsAuthMethodappNotification     UserCredentialUsageDetailsAuthMethod = "AppNotification"
	UserCredentialUsageDetailsAuthMethodappPassword         UserCredentialUsageDetailsAuthMethod = "AppPassword"
	UserCredentialUsageDetailsAuthMethodemail               UserCredentialUsageDetailsAuthMethod = "Email"
	UserCredentialUsageDetailsAuthMethodfido                UserCredentialUsageDetailsAuthMethod = "Fido"
	UserCredentialUsageDetailsAuthMethodmobileCall          UserCredentialUsageDetailsAuthMethod = "MobileCall"
	UserCredentialUsageDetailsAuthMethodmobileSMS           UserCredentialUsageDetailsAuthMethod = "MobileSMS"
	UserCredentialUsageDetailsAuthMethodofficePhone         UserCredentialUsageDetailsAuthMethod = "OfficePhone"
	UserCredentialUsageDetailsAuthMethodsecurityQuestion    UserCredentialUsageDetailsAuthMethod = "SecurityQuestion"
)

func PossibleValuesForUserCredentialUsageDetailsAuthMethod() []string {
	return []string{
		string(UserCredentialUsageDetailsAuthMethodalternateMobileCall),
		string(UserCredentialUsageDetailsAuthMethodappCode),
		string(UserCredentialUsageDetailsAuthMethodappNotification),
		string(UserCredentialUsageDetailsAuthMethodappPassword),
		string(UserCredentialUsageDetailsAuthMethodemail),
		string(UserCredentialUsageDetailsAuthMethodfido),
		string(UserCredentialUsageDetailsAuthMethodmobileCall),
		string(UserCredentialUsageDetailsAuthMethodmobileSMS),
		string(UserCredentialUsageDetailsAuthMethodofficePhone),
		string(UserCredentialUsageDetailsAuthMethodsecurityQuestion),
	}
}

func parseUserCredentialUsageDetailsAuthMethod(input string) (*UserCredentialUsageDetailsAuthMethod, error) {
	vals := map[string]UserCredentialUsageDetailsAuthMethod{
		"alternatemobilecall": UserCredentialUsageDetailsAuthMethodalternateMobileCall,
		"appcode":             UserCredentialUsageDetailsAuthMethodappCode,
		"appnotification":     UserCredentialUsageDetailsAuthMethodappNotification,
		"apppassword":         UserCredentialUsageDetailsAuthMethodappPassword,
		"email":               UserCredentialUsageDetailsAuthMethodemail,
		"fido":                UserCredentialUsageDetailsAuthMethodfido,
		"mobilecall":          UserCredentialUsageDetailsAuthMethodmobileCall,
		"mobilesms":           UserCredentialUsageDetailsAuthMethodmobileSMS,
		"officephone":         UserCredentialUsageDetailsAuthMethodofficePhone,
		"securityquestion":    UserCredentialUsageDetailsAuthMethodsecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserCredentialUsageDetailsAuthMethod(input)
	return &out, nil
}
