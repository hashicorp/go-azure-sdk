package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalHistoryItemRiskDetail string

const (
	RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedServicePrincipalCompromised RiskyServicePrincipalHistoryItemRiskDetail = "adminConfirmedServicePrincipalCompromised"
	RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedSigninCompromised           RiskyServicePrincipalHistoryItemRiskDetail = "adminConfirmedSigninCompromised"
	RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedSigninSafe                  RiskyServicePrincipalHistoryItemRiskDetail = "adminConfirmedSigninSafe"
	RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedUserCompromised             RiskyServicePrincipalHistoryItemRiskDetail = "adminConfirmedUserCompromised"
	RiskyServicePrincipalHistoryItemRiskDetail_AdminDismissedAllRiskForServicePrincipal  RiskyServicePrincipalHistoryItemRiskDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskyServicePrincipalHistoryItemRiskDetail_AdminDismissedAllRiskForUser              RiskyServicePrincipalHistoryItemRiskDetail = "adminDismissedAllRiskForUser"
	RiskyServicePrincipalHistoryItemRiskDetail_AdminGeneratedTemporaryPassword           RiskyServicePrincipalHistoryItemRiskDetail = "adminGeneratedTemporaryPassword"
	RiskyServicePrincipalHistoryItemRiskDetail_AiConfirmedSigninSafe                     RiskyServicePrincipalHistoryItemRiskDetail = "aiConfirmedSigninSafe"
	RiskyServicePrincipalHistoryItemRiskDetail_Hidden                                    RiskyServicePrincipalHistoryItemRiskDetail = "hidden"
	RiskyServicePrincipalHistoryItemRiskDetail_M365DAdminDismissedDetection              RiskyServicePrincipalHistoryItemRiskDetail = "m365DAdminDismissedDetection"
	RiskyServicePrincipalHistoryItemRiskDetail_None                                      RiskyServicePrincipalHistoryItemRiskDetail = "none"
	RiskyServicePrincipalHistoryItemRiskDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskyServicePrincipalHistoryItemRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskyServicePrincipalHistoryItemRiskDetail_UserPerformedSecuredPasswordChange        RiskyServicePrincipalHistoryItemRiskDetail = "userPerformedSecuredPasswordChange"
	RiskyServicePrincipalHistoryItemRiskDetail_UserPerformedSecuredPasswordReset         RiskyServicePrincipalHistoryItemRiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyServicePrincipalHistoryItemRiskDetail() []string {
	return []string{
		string(RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedSigninCompromised),
		string(RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedSigninSafe),
		string(RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedUserCompromised),
		string(RiskyServicePrincipalHistoryItemRiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskyServicePrincipalHistoryItemRiskDetail_AdminDismissedAllRiskForUser),
		string(RiskyServicePrincipalHistoryItemRiskDetail_AdminGeneratedTemporaryPassword),
		string(RiskyServicePrincipalHistoryItemRiskDetail_AiConfirmedSigninSafe),
		string(RiskyServicePrincipalHistoryItemRiskDetail_Hidden),
		string(RiskyServicePrincipalHistoryItemRiskDetail_M365DAdminDismissedDetection),
		string(RiskyServicePrincipalHistoryItemRiskDetail_None),
		string(RiskyServicePrincipalHistoryItemRiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyServicePrincipalHistoryItemRiskDetail_UserPerformedSecuredPasswordChange),
		string(RiskyServicePrincipalHistoryItemRiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskyServicePrincipalHistoryItemRiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyServicePrincipalHistoryItemRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyServicePrincipalHistoryItemRiskDetail(input string) (*RiskyServicePrincipalHistoryItemRiskDetail, error) {
	vals := map[string]RiskyServicePrincipalHistoryItemRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyServicePrincipalHistoryItemRiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyServicePrincipalHistoryItemRiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyServicePrincipalHistoryItemRiskDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyServicePrincipalHistoryItemRiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyServicePrincipalHistoryItemRiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskyServicePrincipalHistoryItemRiskDetail_Hidden,
		"m365dadmindismisseddetection":              RiskyServicePrincipalHistoryItemRiskDetail_M365DAdminDismissedDetection,
		"none":                                      RiskyServicePrincipalHistoryItemRiskDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyServicePrincipalHistoryItemRiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyServicePrincipalHistoryItemRiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyServicePrincipalHistoryItemRiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalHistoryItemRiskDetail(input)
	return &out, nil
}
