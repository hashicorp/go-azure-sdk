package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUsageSummaryAuthMethod string

const (
	CredentialUsageSummaryAuthMethodalternateMobileCall CredentialUsageSummaryAuthMethod = "AlternateMobileCall"
	CredentialUsageSummaryAuthMethodappCode             CredentialUsageSummaryAuthMethod = "AppCode"
	CredentialUsageSummaryAuthMethodappNotification     CredentialUsageSummaryAuthMethod = "AppNotification"
	CredentialUsageSummaryAuthMethodappPassword         CredentialUsageSummaryAuthMethod = "AppPassword"
	CredentialUsageSummaryAuthMethodemail               CredentialUsageSummaryAuthMethod = "Email"
	CredentialUsageSummaryAuthMethodfido                CredentialUsageSummaryAuthMethod = "Fido"
	CredentialUsageSummaryAuthMethodmobileCall          CredentialUsageSummaryAuthMethod = "MobileCall"
	CredentialUsageSummaryAuthMethodmobileSMS           CredentialUsageSummaryAuthMethod = "MobileSMS"
	CredentialUsageSummaryAuthMethodofficePhone         CredentialUsageSummaryAuthMethod = "OfficePhone"
	CredentialUsageSummaryAuthMethodsecurityQuestion    CredentialUsageSummaryAuthMethod = "SecurityQuestion"
)

func PossibleValuesForCredentialUsageSummaryAuthMethod() []string {
	return []string{
		string(CredentialUsageSummaryAuthMethodalternateMobileCall),
		string(CredentialUsageSummaryAuthMethodappCode),
		string(CredentialUsageSummaryAuthMethodappNotification),
		string(CredentialUsageSummaryAuthMethodappPassword),
		string(CredentialUsageSummaryAuthMethodemail),
		string(CredentialUsageSummaryAuthMethodfido),
		string(CredentialUsageSummaryAuthMethodmobileCall),
		string(CredentialUsageSummaryAuthMethodmobileSMS),
		string(CredentialUsageSummaryAuthMethodofficePhone),
		string(CredentialUsageSummaryAuthMethodsecurityQuestion),
	}
}

func parseCredentialUsageSummaryAuthMethod(input string) (*CredentialUsageSummaryAuthMethod, error) {
	vals := map[string]CredentialUsageSummaryAuthMethod{
		"alternatemobilecall": CredentialUsageSummaryAuthMethodalternateMobileCall,
		"appcode":             CredentialUsageSummaryAuthMethodappCode,
		"appnotification":     CredentialUsageSummaryAuthMethodappNotification,
		"apppassword":         CredentialUsageSummaryAuthMethodappPassword,
		"email":               CredentialUsageSummaryAuthMethodemail,
		"fido":                CredentialUsageSummaryAuthMethodfido,
		"mobilecall":          CredentialUsageSummaryAuthMethodmobileCall,
		"mobilesms":           CredentialUsageSummaryAuthMethodmobileSMS,
		"officephone":         CredentialUsageSummaryAuthMethodofficePhone,
		"securityquestion":    CredentialUsageSummaryAuthMethodsecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialUsageSummaryAuthMethod(input)
	return &out, nil
}
