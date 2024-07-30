package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUsageSummaryAuthMethod string

const (
	CredentialUsageSummaryAuthMethod_AlternateMobileCall CredentialUsageSummaryAuthMethod = "alternateMobileCall"
	CredentialUsageSummaryAuthMethod_AppCode             CredentialUsageSummaryAuthMethod = "appCode"
	CredentialUsageSummaryAuthMethod_AppNotification     CredentialUsageSummaryAuthMethod = "appNotification"
	CredentialUsageSummaryAuthMethod_AppPassword         CredentialUsageSummaryAuthMethod = "appPassword"
	CredentialUsageSummaryAuthMethod_Email               CredentialUsageSummaryAuthMethod = "email"
	CredentialUsageSummaryAuthMethod_Fido                CredentialUsageSummaryAuthMethod = "fido"
	CredentialUsageSummaryAuthMethod_MobileCall          CredentialUsageSummaryAuthMethod = "mobileCall"
	CredentialUsageSummaryAuthMethod_MobileSMS           CredentialUsageSummaryAuthMethod = "mobileSMS"
	CredentialUsageSummaryAuthMethod_OfficePhone         CredentialUsageSummaryAuthMethod = "officePhone"
	CredentialUsageSummaryAuthMethod_SecurityQuestion    CredentialUsageSummaryAuthMethod = "securityQuestion"
)

func PossibleValuesForCredentialUsageSummaryAuthMethod() []string {
	return []string{
		string(CredentialUsageSummaryAuthMethod_AlternateMobileCall),
		string(CredentialUsageSummaryAuthMethod_AppCode),
		string(CredentialUsageSummaryAuthMethod_AppNotification),
		string(CredentialUsageSummaryAuthMethod_AppPassword),
		string(CredentialUsageSummaryAuthMethod_Email),
		string(CredentialUsageSummaryAuthMethod_Fido),
		string(CredentialUsageSummaryAuthMethod_MobileCall),
		string(CredentialUsageSummaryAuthMethod_MobileSMS),
		string(CredentialUsageSummaryAuthMethod_OfficePhone),
		string(CredentialUsageSummaryAuthMethod_SecurityQuestion),
	}
}

func (s *CredentialUsageSummaryAuthMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCredentialUsageSummaryAuthMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCredentialUsageSummaryAuthMethod(input string) (*CredentialUsageSummaryAuthMethod, error) {
	vals := map[string]CredentialUsageSummaryAuthMethod{
		"alternatemobilecall": CredentialUsageSummaryAuthMethod_AlternateMobileCall,
		"appcode":             CredentialUsageSummaryAuthMethod_AppCode,
		"appnotification":     CredentialUsageSummaryAuthMethod_AppNotification,
		"apppassword":         CredentialUsageSummaryAuthMethod_AppPassword,
		"email":               CredentialUsageSummaryAuthMethod_Email,
		"fido":                CredentialUsageSummaryAuthMethod_Fido,
		"mobilecall":          CredentialUsageSummaryAuthMethod_MobileCall,
		"mobilesms":           CredentialUsageSummaryAuthMethod_MobileSMS,
		"officephone":         CredentialUsageSummaryAuthMethod_OfficePhone,
		"securityquestion":    CredentialUsageSummaryAuthMethod_SecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialUsageSummaryAuthMethod(input)
	return &out, nil
}
