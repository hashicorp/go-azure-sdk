package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCredentialUsageDetailsAuthMethod string

const (
	UserCredentialUsageDetailsAuthMethod_AlternateMobileCall UserCredentialUsageDetailsAuthMethod = "alternateMobileCall"
	UserCredentialUsageDetailsAuthMethod_AppCode             UserCredentialUsageDetailsAuthMethod = "appCode"
	UserCredentialUsageDetailsAuthMethod_AppNotification     UserCredentialUsageDetailsAuthMethod = "appNotification"
	UserCredentialUsageDetailsAuthMethod_AppPassword         UserCredentialUsageDetailsAuthMethod = "appPassword"
	UserCredentialUsageDetailsAuthMethod_Email               UserCredentialUsageDetailsAuthMethod = "email"
	UserCredentialUsageDetailsAuthMethod_Fido                UserCredentialUsageDetailsAuthMethod = "fido"
	UserCredentialUsageDetailsAuthMethod_MobileCall          UserCredentialUsageDetailsAuthMethod = "mobileCall"
	UserCredentialUsageDetailsAuthMethod_MobileSMS           UserCredentialUsageDetailsAuthMethod = "mobileSMS"
	UserCredentialUsageDetailsAuthMethod_OfficePhone         UserCredentialUsageDetailsAuthMethod = "officePhone"
	UserCredentialUsageDetailsAuthMethod_SecurityQuestion    UserCredentialUsageDetailsAuthMethod = "securityQuestion"
)

func PossibleValuesForUserCredentialUsageDetailsAuthMethod() []string {
	return []string{
		string(UserCredentialUsageDetailsAuthMethod_AlternateMobileCall),
		string(UserCredentialUsageDetailsAuthMethod_AppCode),
		string(UserCredentialUsageDetailsAuthMethod_AppNotification),
		string(UserCredentialUsageDetailsAuthMethod_AppPassword),
		string(UserCredentialUsageDetailsAuthMethod_Email),
		string(UserCredentialUsageDetailsAuthMethod_Fido),
		string(UserCredentialUsageDetailsAuthMethod_MobileCall),
		string(UserCredentialUsageDetailsAuthMethod_MobileSMS),
		string(UserCredentialUsageDetailsAuthMethod_OfficePhone),
		string(UserCredentialUsageDetailsAuthMethod_SecurityQuestion),
	}
}

func (s *UserCredentialUsageDetailsAuthMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserCredentialUsageDetailsAuthMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserCredentialUsageDetailsAuthMethod(input string) (*UserCredentialUsageDetailsAuthMethod, error) {
	vals := map[string]UserCredentialUsageDetailsAuthMethod{
		"alternatemobilecall": UserCredentialUsageDetailsAuthMethod_AlternateMobileCall,
		"appcode":             UserCredentialUsageDetailsAuthMethod_AppCode,
		"appnotification":     UserCredentialUsageDetailsAuthMethod_AppNotification,
		"apppassword":         UserCredentialUsageDetailsAuthMethod_AppPassword,
		"email":               UserCredentialUsageDetailsAuthMethod_Email,
		"fido":                UserCredentialUsageDetailsAuthMethod_Fido,
		"mobilecall":          UserCredentialUsageDetailsAuthMethod_MobileCall,
		"mobilesms":           UserCredentialUsageDetailsAuthMethod_MobileSMS,
		"officephone":         UserCredentialUsageDetailsAuthMethod_OfficePhone,
		"securityquestion":    UserCredentialUsageDetailsAuthMethod_SecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserCredentialUsageDetailsAuthMethod(input)
	return &out, nil
}
