package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalRiskDetail string

const (
	RiskyServicePrincipalRiskDetail_AdminConfirmedServicePrincipalCompromised RiskyServicePrincipalRiskDetail = "adminConfirmedServicePrincipalCompromised"
	RiskyServicePrincipalRiskDetail_AdminConfirmedSigninCompromised           RiskyServicePrincipalRiskDetail = "adminConfirmedSigninCompromised"
	RiskyServicePrincipalRiskDetail_AdminConfirmedSigninSafe                  RiskyServicePrincipalRiskDetail = "adminConfirmedSigninSafe"
	RiskyServicePrincipalRiskDetail_AdminConfirmedUserCompromised             RiskyServicePrincipalRiskDetail = "adminConfirmedUserCompromised"
	RiskyServicePrincipalRiskDetail_AdminDismissedAllRiskForServicePrincipal  RiskyServicePrincipalRiskDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskyServicePrincipalRiskDetail_AdminDismissedAllRiskForUser              RiskyServicePrincipalRiskDetail = "adminDismissedAllRiskForUser"
	RiskyServicePrincipalRiskDetail_AdminGeneratedTemporaryPassword           RiskyServicePrincipalRiskDetail = "adminGeneratedTemporaryPassword"
	RiskyServicePrincipalRiskDetail_AiConfirmedSigninSafe                     RiskyServicePrincipalRiskDetail = "aiConfirmedSigninSafe"
	RiskyServicePrincipalRiskDetail_Hidden                                    RiskyServicePrincipalRiskDetail = "hidden"
	RiskyServicePrincipalRiskDetail_M365DAdminDismissedDetection              RiskyServicePrincipalRiskDetail = "m365DAdminDismissedDetection"
	RiskyServicePrincipalRiskDetail_None                                      RiskyServicePrincipalRiskDetail = "none"
	RiskyServicePrincipalRiskDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskyServicePrincipalRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskyServicePrincipalRiskDetail_UserPerformedSecuredPasswordChange        RiskyServicePrincipalRiskDetail = "userPerformedSecuredPasswordChange"
	RiskyServicePrincipalRiskDetail_UserPerformedSecuredPasswordReset         RiskyServicePrincipalRiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyServicePrincipalRiskDetail() []string {
	return []string{
		string(RiskyServicePrincipalRiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskyServicePrincipalRiskDetail_AdminConfirmedSigninCompromised),
		string(RiskyServicePrincipalRiskDetail_AdminConfirmedSigninSafe),
		string(RiskyServicePrincipalRiskDetail_AdminConfirmedUserCompromised),
		string(RiskyServicePrincipalRiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskyServicePrincipalRiskDetail_AdminDismissedAllRiskForUser),
		string(RiskyServicePrincipalRiskDetail_AdminGeneratedTemporaryPassword),
		string(RiskyServicePrincipalRiskDetail_AiConfirmedSigninSafe),
		string(RiskyServicePrincipalRiskDetail_Hidden),
		string(RiskyServicePrincipalRiskDetail_M365DAdminDismissedDetection),
		string(RiskyServicePrincipalRiskDetail_None),
		string(RiskyServicePrincipalRiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyServicePrincipalRiskDetail_UserPerformedSecuredPasswordChange),
		string(RiskyServicePrincipalRiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskyServicePrincipalRiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyServicePrincipalRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyServicePrincipalRiskDetail(input string) (*RiskyServicePrincipalRiskDetail, error) {
	vals := map[string]RiskyServicePrincipalRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyServicePrincipalRiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyServicePrincipalRiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyServicePrincipalRiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyServicePrincipalRiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyServicePrincipalRiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyServicePrincipalRiskDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyServicePrincipalRiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyServicePrincipalRiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskyServicePrincipalRiskDetail_Hidden,
		"m365dadmindismisseddetection":              RiskyServicePrincipalRiskDetail_M365DAdminDismissedDetection,
		"none":                                      RiskyServicePrincipalRiskDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyServicePrincipalRiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyServicePrincipalRiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyServicePrincipalRiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalRiskDetail(input)
	return &out, nil
}
