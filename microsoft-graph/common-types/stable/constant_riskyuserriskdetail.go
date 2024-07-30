package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserRiskDetail string

const (
	RiskyUserRiskDetail_AdminConfirmedServicePrincipalCompromised RiskyUserRiskDetail = "adminConfirmedServicePrincipalCompromised"
	RiskyUserRiskDetail_AdminConfirmedSigninCompromised           RiskyUserRiskDetail = "adminConfirmedSigninCompromised"
	RiskyUserRiskDetail_AdminConfirmedSigninSafe                  RiskyUserRiskDetail = "adminConfirmedSigninSafe"
	RiskyUserRiskDetail_AdminConfirmedUserCompromised             RiskyUserRiskDetail = "adminConfirmedUserCompromised"
	RiskyUserRiskDetail_AdminDismissedAllRiskForServicePrincipal  RiskyUserRiskDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskyUserRiskDetail_AdminDismissedAllRiskForUser              RiskyUserRiskDetail = "adminDismissedAllRiskForUser"
	RiskyUserRiskDetail_AdminGeneratedTemporaryPassword           RiskyUserRiskDetail = "adminGeneratedTemporaryPassword"
	RiskyUserRiskDetail_AiConfirmedSigninSafe                     RiskyUserRiskDetail = "aiConfirmedSigninSafe"
	RiskyUserRiskDetail_Hidden                                    RiskyUserRiskDetail = "hidden"
	RiskyUserRiskDetail_M365DAdminDismissedDetection              RiskyUserRiskDetail = "m365DAdminDismissedDetection"
	RiskyUserRiskDetail_None                                      RiskyUserRiskDetail = "none"
	RiskyUserRiskDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskyUserRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskyUserRiskDetail_UserPerformedSecuredPasswordChange        RiskyUserRiskDetail = "userPerformedSecuredPasswordChange"
	RiskyUserRiskDetail_UserPerformedSecuredPasswordReset         RiskyUserRiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyUserRiskDetail() []string {
	return []string{
		string(RiskyUserRiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskyUserRiskDetail_AdminConfirmedSigninCompromised),
		string(RiskyUserRiskDetail_AdminConfirmedSigninSafe),
		string(RiskyUserRiskDetail_AdminConfirmedUserCompromised),
		string(RiskyUserRiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskyUserRiskDetail_AdminDismissedAllRiskForUser),
		string(RiskyUserRiskDetail_AdminGeneratedTemporaryPassword),
		string(RiskyUserRiskDetail_AiConfirmedSigninSafe),
		string(RiskyUserRiskDetail_Hidden),
		string(RiskyUserRiskDetail_M365DAdminDismissedDetection),
		string(RiskyUserRiskDetail_None),
		string(RiskyUserRiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyUserRiskDetail_UserPerformedSecuredPasswordChange),
		string(RiskyUserRiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskyUserRiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyUserRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyUserRiskDetail(input string) (*RiskyUserRiskDetail, error) {
	vals := map[string]RiskyUserRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyUserRiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyUserRiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyUserRiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyUserRiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyUserRiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyUserRiskDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyUserRiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyUserRiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskyUserRiskDetail_Hidden,
		"m365dadmindismisseddetection":              RiskyUserRiskDetail_M365DAdminDismissedDetection,
		"none":                                      RiskyUserRiskDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyUserRiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyUserRiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyUserRiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserRiskDetail(input)
	return &out, nil
}
