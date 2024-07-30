package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskServicePrincipalActivityDetail string

const (
	RiskServicePrincipalActivityDetail_AdminConfirmedServicePrincipalCompromised RiskServicePrincipalActivityDetail = "adminConfirmedServicePrincipalCompromised"
	RiskServicePrincipalActivityDetail_AdminConfirmedSigninCompromised           RiskServicePrincipalActivityDetail = "adminConfirmedSigninCompromised"
	RiskServicePrincipalActivityDetail_AdminConfirmedSigninSafe                  RiskServicePrincipalActivityDetail = "adminConfirmedSigninSafe"
	RiskServicePrincipalActivityDetail_AdminConfirmedUserCompromised             RiskServicePrincipalActivityDetail = "adminConfirmedUserCompromised"
	RiskServicePrincipalActivityDetail_AdminDismissedAllRiskForServicePrincipal  RiskServicePrincipalActivityDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskServicePrincipalActivityDetail_AdminDismissedAllRiskForUser              RiskServicePrincipalActivityDetail = "adminDismissedAllRiskForUser"
	RiskServicePrincipalActivityDetail_AdminGeneratedTemporaryPassword           RiskServicePrincipalActivityDetail = "adminGeneratedTemporaryPassword"
	RiskServicePrincipalActivityDetail_AiConfirmedSigninSafe                     RiskServicePrincipalActivityDetail = "aiConfirmedSigninSafe"
	RiskServicePrincipalActivityDetail_Hidden                                    RiskServicePrincipalActivityDetail = "hidden"
	RiskServicePrincipalActivityDetail_M365DAdminDismissedDetection              RiskServicePrincipalActivityDetail = "m365DAdminDismissedDetection"
	RiskServicePrincipalActivityDetail_None                                      RiskServicePrincipalActivityDetail = "none"
	RiskServicePrincipalActivityDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskServicePrincipalActivityDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskServicePrincipalActivityDetail_UserPerformedSecuredPasswordChange        RiskServicePrincipalActivityDetail = "userPerformedSecuredPasswordChange"
	RiskServicePrincipalActivityDetail_UserPerformedSecuredPasswordReset         RiskServicePrincipalActivityDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskServicePrincipalActivityDetail() []string {
	return []string{
		string(RiskServicePrincipalActivityDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskServicePrincipalActivityDetail_AdminConfirmedSigninCompromised),
		string(RiskServicePrincipalActivityDetail_AdminConfirmedSigninSafe),
		string(RiskServicePrincipalActivityDetail_AdminConfirmedUserCompromised),
		string(RiskServicePrincipalActivityDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskServicePrincipalActivityDetail_AdminDismissedAllRiskForUser),
		string(RiskServicePrincipalActivityDetail_AdminGeneratedTemporaryPassword),
		string(RiskServicePrincipalActivityDetail_AiConfirmedSigninSafe),
		string(RiskServicePrincipalActivityDetail_Hidden),
		string(RiskServicePrincipalActivityDetail_M365DAdminDismissedDetection),
		string(RiskServicePrincipalActivityDetail_None),
		string(RiskServicePrincipalActivityDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskServicePrincipalActivityDetail_UserPerformedSecuredPasswordChange),
		string(RiskServicePrincipalActivityDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskServicePrincipalActivityDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskServicePrincipalActivityDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskServicePrincipalActivityDetail(input string) (*RiskServicePrincipalActivityDetail, error) {
	vals := map[string]RiskServicePrincipalActivityDetail{
		"adminconfirmedserviceprincipalcompromised": RiskServicePrincipalActivityDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskServicePrincipalActivityDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskServicePrincipalActivityDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskServicePrincipalActivityDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskServicePrincipalActivityDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskServicePrincipalActivityDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskServicePrincipalActivityDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskServicePrincipalActivityDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskServicePrincipalActivityDetail_Hidden,
		"m365dadmindismisseddetection":              RiskServicePrincipalActivityDetail_M365DAdminDismissedDetection,
		"none":                                      RiskServicePrincipalActivityDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskServicePrincipalActivityDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskServicePrincipalActivityDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskServicePrincipalActivityDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskServicePrincipalActivityDetail(input)
	return &out, nil
}
