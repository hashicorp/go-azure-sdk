package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskUserActivityDetail string

const (
	RiskUserActivityDetail_AdminConfirmedServicePrincipalCompromised RiskUserActivityDetail = "adminConfirmedServicePrincipalCompromised"
	RiskUserActivityDetail_AdminConfirmedSigninCompromised           RiskUserActivityDetail = "adminConfirmedSigninCompromised"
	RiskUserActivityDetail_AdminConfirmedSigninSafe                  RiskUserActivityDetail = "adminConfirmedSigninSafe"
	RiskUserActivityDetail_AdminConfirmedUserCompromised             RiskUserActivityDetail = "adminConfirmedUserCompromised"
	RiskUserActivityDetail_AdminDismissedAllRiskForServicePrincipal  RiskUserActivityDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskUserActivityDetail_AdminDismissedAllRiskForUser              RiskUserActivityDetail = "adminDismissedAllRiskForUser"
	RiskUserActivityDetail_AdminGeneratedTemporaryPassword           RiskUserActivityDetail = "adminGeneratedTemporaryPassword"
	RiskUserActivityDetail_AiConfirmedSigninSafe                     RiskUserActivityDetail = "aiConfirmedSigninSafe"
	RiskUserActivityDetail_Hidden                                    RiskUserActivityDetail = "hidden"
	RiskUserActivityDetail_M365DAdminDismissedDetection              RiskUserActivityDetail = "m365DAdminDismissedDetection"
	RiskUserActivityDetail_None                                      RiskUserActivityDetail = "none"
	RiskUserActivityDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskUserActivityDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskUserActivityDetail_UserPerformedSecuredPasswordChange        RiskUserActivityDetail = "userPerformedSecuredPasswordChange"
	RiskUserActivityDetail_UserPerformedSecuredPasswordReset         RiskUserActivityDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskUserActivityDetail() []string {
	return []string{
		string(RiskUserActivityDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskUserActivityDetail_AdminConfirmedSigninCompromised),
		string(RiskUserActivityDetail_AdminConfirmedSigninSafe),
		string(RiskUserActivityDetail_AdminConfirmedUserCompromised),
		string(RiskUserActivityDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskUserActivityDetail_AdminDismissedAllRiskForUser),
		string(RiskUserActivityDetail_AdminGeneratedTemporaryPassword),
		string(RiskUserActivityDetail_AiConfirmedSigninSafe),
		string(RiskUserActivityDetail_Hidden),
		string(RiskUserActivityDetail_M365DAdminDismissedDetection),
		string(RiskUserActivityDetail_None),
		string(RiskUserActivityDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskUserActivityDetail_UserPerformedSecuredPasswordChange),
		string(RiskUserActivityDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskUserActivityDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskUserActivityDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskUserActivityDetail(input string) (*RiskUserActivityDetail, error) {
	vals := map[string]RiskUserActivityDetail{
		"adminconfirmedserviceprincipalcompromised": RiskUserActivityDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskUserActivityDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskUserActivityDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskUserActivityDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskUserActivityDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskUserActivityDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskUserActivityDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskUserActivityDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskUserActivityDetail_Hidden,
		"m365dadmindismisseddetection":              RiskUserActivityDetail_M365DAdminDismissedDetection,
		"none":                                      RiskUserActivityDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskUserActivityDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskUserActivityDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskUserActivityDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskUserActivityDetail(input)
	return &out, nil
}
