package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskDetail string

const (
	SignInRiskDetail_AdminConfirmedServicePrincipalCompromised SignInRiskDetail = "adminConfirmedServicePrincipalCompromised"
	SignInRiskDetail_AdminConfirmedSigninCompromised           SignInRiskDetail = "adminConfirmedSigninCompromised"
	SignInRiskDetail_AdminConfirmedSigninSafe                  SignInRiskDetail = "adminConfirmedSigninSafe"
	SignInRiskDetail_AdminConfirmedUserCompromised             SignInRiskDetail = "adminConfirmedUserCompromised"
	SignInRiskDetail_AdminDismissedAllRiskForServicePrincipal  SignInRiskDetail = "adminDismissedAllRiskForServicePrincipal"
	SignInRiskDetail_AdminDismissedAllRiskForUser              SignInRiskDetail = "adminDismissedAllRiskForUser"
	SignInRiskDetail_AdminGeneratedTemporaryPassword           SignInRiskDetail = "adminGeneratedTemporaryPassword"
	SignInRiskDetail_AiConfirmedSigninSafe                     SignInRiskDetail = "aiConfirmedSigninSafe"
	SignInRiskDetail_Hidden                                    SignInRiskDetail = "hidden"
	SignInRiskDetail_M365DAdminDismissedDetection              SignInRiskDetail = "m365DAdminDismissedDetection"
	SignInRiskDetail_None                                      SignInRiskDetail = "none"
	SignInRiskDetail_UserPassedMFADrivenByRiskBasedPolicy      SignInRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	SignInRiskDetail_UserPerformedSecuredPasswordChange        SignInRiskDetail = "userPerformedSecuredPasswordChange"
	SignInRiskDetail_UserPerformedSecuredPasswordReset         SignInRiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForSignInRiskDetail() []string {
	return []string{
		string(SignInRiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(SignInRiskDetail_AdminConfirmedSigninCompromised),
		string(SignInRiskDetail_AdminConfirmedSigninSafe),
		string(SignInRiskDetail_AdminConfirmedUserCompromised),
		string(SignInRiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(SignInRiskDetail_AdminDismissedAllRiskForUser),
		string(SignInRiskDetail_AdminGeneratedTemporaryPassword),
		string(SignInRiskDetail_AiConfirmedSigninSafe),
		string(SignInRiskDetail_Hidden),
		string(SignInRiskDetail_M365DAdminDismissedDetection),
		string(SignInRiskDetail_None),
		string(SignInRiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(SignInRiskDetail_UserPerformedSecuredPasswordChange),
		string(SignInRiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *SignInRiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInRiskDetail(input string) (*SignInRiskDetail, error) {
	vals := map[string]SignInRiskDetail{
		"adminconfirmedserviceprincipalcompromised": SignInRiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           SignInRiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  SignInRiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             SignInRiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  SignInRiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              SignInRiskDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           SignInRiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     SignInRiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    SignInRiskDetail_Hidden,
		"m365dadmindismisseddetection":              SignInRiskDetail_M365DAdminDismissedDetection,
		"none":                                      SignInRiskDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      SignInRiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        SignInRiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         SignInRiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskDetail(input)
	return &out, nil
}
