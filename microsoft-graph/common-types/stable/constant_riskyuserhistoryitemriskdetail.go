package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItemRiskDetail string

const (
	RiskyUserHistoryItemRiskDetail_AdminConfirmedServicePrincipalCompromised RiskyUserHistoryItemRiskDetail = "adminConfirmedServicePrincipalCompromised"
	RiskyUserHistoryItemRiskDetail_AdminConfirmedSigninCompromised           RiskyUserHistoryItemRiskDetail = "adminConfirmedSigninCompromised"
	RiskyUserHistoryItemRiskDetail_AdminConfirmedSigninSafe                  RiskyUserHistoryItemRiskDetail = "adminConfirmedSigninSafe"
	RiskyUserHistoryItemRiskDetail_AdminConfirmedUserCompromised             RiskyUserHistoryItemRiskDetail = "adminConfirmedUserCompromised"
	RiskyUserHistoryItemRiskDetail_AdminDismissedAllRiskForServicePrincipal  RiskyUserHistoryItemRiskDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskyUserHistoryItemRiskDetail_AdminDismissedAllRiskForUser              RiskyUserHistoryItemRiskDetail = "adminDismissedAllRiskForUser"
	RiskyUserHistoryItemRiskDetail_AdminGeneratedTemporaryPassword           RiskyUserHistoryItemRiskDetail = "adminGeneratedTemporaryPassword"
	RiskyUserHistoryItemRiskDetail_AiConfirmedSigninSafe                     RiskyUserHistoryItemRiskDetail = "aiConfirmedSigninSafe"
	RiskyUserHistoryItemRiskDetail_Hidden                                    RiskyUserHistoryItemRiskDetail = "hidden"
	RiskyUserHistoryItemRiskDetail_M365DAdminDismissedDetection              RiskyUserHistoryItemRiskDetail = "m365DAdminDismissedDetection"
	RiskyUserHistoryItemRiskDetail_None                                      RiskyUserHistoryItemRiskDetail = "none"
	RiskyUserHistoryItemRiskDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskyUserHistoryItemRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskyUserHistoryItemRiskDetail_UserPerformedSecuredPasswordChange        RiskyUserHistoryItemRiskDetail = "userPerformedSecuredPasswordChange"
	RiskyUserHistoryItemRiskDetail_UserPerformedSecuredPasswordReset         RiskyUserHistoryItemRiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyUserHistoryItemRiskDetail() []string {
	return []string{
		string(RiskyUserHistoryItemRiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskyUserHistoryItemRiskDetail_AdminConfirmedSigninCompromised),
		string(RiskyUserHistoryItemRiskDetail_AdminConfirmedSigninSafe),
		string(RiskyUserHistoryItemRiskDetail_AdminConfirmedUserCompromised),
		string(RiskyUserHistoryItemRiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskyUserHistoryItemRiskDetail_AdminDismissedAllRiskForUser),
		string(RiskyUserHistoryItemRiskDetail_AdminGeneratedTemporaryPassword),
		string(RiskyUserHistoryItemRiskDetail_AiConfirmedSigninSafe),
		string(RiskyUserHistoryItemRiskDetail_Hidden),
		string(RiskyUserHistoryItemRiskDetail_M365DAdminDismissedDetection),
		string(RiskyUserHistoryItemRiskDetail_None),
		string(RiskyUserHistoryItemRiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyUserHistoryItemRiskDetail_UserPerformedSecuredPasswordChange),
		string(RiskyUserHistoryItemRiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskyUserHistoryItemRiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyUserHistoryItemRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyUserHistoryItemRiskDetail(input string) (*RiskyUserHistoryItemRiskDetail, error) {
	vals := map[string]RiskyUserHistoryItemRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyUserHistoryItemRiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyUserHistoryItemRiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyUserHistoryItemRiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyUserHistoryItemRiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyUserHistoryItemRiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyUserHistoryItemRiskDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyUserHistoryItemRiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyUserHistoryItemRiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskyUserHistoryItemRiskDetail_Hidden,
		"m365dadmindismisseddetection":              RiskyUserHistoryItemRiskDetail_M365DAdminDismissedDetection,
		"none":                                      RiskyUserHistoryItemRiskDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyUserHistoryItemRiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyUserHistoryItemRiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyUserHistoryItemRiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserHistoryItemRiskDetail(input)
	return &out, nil
}
