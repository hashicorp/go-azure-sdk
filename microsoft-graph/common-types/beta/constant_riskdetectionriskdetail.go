package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskDetail string

const (
	RiskDetectionRiskDetail_AdminConfirmedServicePrincipalCompromised RiskDetectionRiskDetail = "adminConfirmedServicePrincipalCompromised"
	RiskDetectionRiskDetail_AdminConfirmedSigninCompromised           RiskDetectionRiskDetail = "adminConfirmedSigninCompromised"
	RiskDetectionRiskDetail_AdminConfirmedSigninSafe                  RiskDetectionRiskDetail = "adminConfirmedSigninSafe"
	RiskDetectionRiskDetail_AdminConfirmedUserCompromised             RiskDetectionRiskDetail = "adminConfirmedUserCompromised"
	RiskDetectionRiskDetail_AdminDismissedAllRiskForServicePrincipal  RiskDetectionRiskDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskDetectionRiskDetail_AdminDismissedAllRiskForUser              RiskDetectionRiskDetail = "adminDismissedAllRiskForUser"
	RiskDetectionRiskDetail_AdminGeneratedTemporaryPassword           RiskDetectionRiskDetail = "adminGeneratedTemporaryPassword"
	RiskDetectionRiskDetail_AiConfirmedSigninSafe                     RiskDetectionRiskDetail = "aiConfirmedSigninSafe"
	RiskDetectionRiskDetail_Hidden                                    RiskDetectionRiskDetail = "hidden"
	RiskDetectionRiskDetail_M365DAdminDismissedDetection              RiskDetectionRiskDetail = "m365DAdminDismissedDetection"
	RiskDetectionRiskDetail_None                                      RiskDetectionRiskDetail = "none"
	RiskDetectionRiskDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskDetectionRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskDetectionRiskDetail_UserPerformedSecuredPasswordChange        RiskDetectionRiskDetail = "userPerformedSecuredPasswordChange"
	RiskDetectionRiskDetail_UserPerformedSecuredPasswordReset         RiskDetectionRiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskDetectionRiskDetail() []string {
	return []string{
		string(RiskDetectionRiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskDetectionRiskDetail_AdminConfirmedSigninCompromised),
		string(RiskDetectionRiskDetail_AdminConfirmedSigninSafe),
		string(RiskDetectionRiskDetail_AdminConfirmedUserCompromised),
		string(RiskDetectionRiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskDetectionRiskDetail_AdminDismissedAllRiskForUser),
		string(RiskDetectionRiskDetail_AdminGeneratedTemporaryPassword),
		string(RiskDetectionRiskDetail_AiConfirmedSigninSafe),
		string(RiskDetectionRiskDetail_Hidden),
		string(RiskDetectionRiskDetail_M365DAdminDismissedDetection),
		string(RiskDetectionRiskDetail_None),
		string(RiskDetectionRiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskDetectionRiskDetail_UserPerformedSecuredPasswordChange),
		string(RiskDetectionRiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskDetectionRiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionRiskDetail(input string) (*RiskDetectionRiskDetail, error) {
	vals := map[string]RiskDetectionRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskDetectionRiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskDetectionRiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskDetectionRiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskDetectionRiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskDetectionRiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskDetectionRiskDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskDetectionRiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskDetectionRiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskDetectionRiskDetail_Hidden,
		"m365dadmindismisseddetection":              RiskDetectionRiskDetail_M365DAdminDismissedDetection,
		"none":                                      RiskDetectionRiskDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskDetectionRiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskDetectionRiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskDetectionRiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskDetail(input)
	return &out, nil
}
