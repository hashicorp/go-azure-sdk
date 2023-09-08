package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskUserActivityDetail string

const (
	RiskUserActivityDetailadminConfirmedServicePrincipalCompromised RiskUserActivityDetail = "AdminConfirmedServicePrincipalCompromised"
	RiskUserActivityDetailadminConfirmedSigninCompromised           RiskUserActivityDetail = "AdminConfirmedSigninCompromised"
	RiskUserActivityDetailadminConfirmedSigninSafe                  RiskUserActivityDetail = "AdminConfirmedSigninSafe"
	RiskUserActivityDetailadminConfirmedUserCompromised             RiskUserActivityDetail = "AdminConfirmedUserCompromised"
	RiskUserActivityDetailadminDismissedAllRiskForServicePrincipal  RiskUserActivityDetail = "AdminDismissedAllRiskForServicePrincipal"
	RiskUserActivityDetailadminDismissedAllRiskForUser              RiskUserActivityDetail = "AdminDismissedAllRiskForUser"
	RiskUserActivityDetailadminGeneratedTemporaryPassword           RiskUserActivityDetail = "AdminGeneratedTemporaryPassword"
	RiskUserActivityDetailaiConfirmedSigninSafe                     RiskUserActivityDetail = "AiConfirmedSigninSafe"
	RiskUserActivityDetailhidden                                    RiskUserActivityDetail = "Hidden"
	RiskUserActivityDetailm365DAdminDismissedDetection              RiskUserActivityDetail = "M365DAdminDismissedDetection"
	RiskUserActivityDetailnone                                      RiskUserActivityDetail = "None"
	RiskUserActivityDetailuserPassedMFADrivenByRiskBasedPolicy      RiskUserActivityDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	RiskUserActivityDetailuserPerformedSecuredPasswordChange        RiskUserActivityDetail = "UserPerformedSecuredPasswordChange"
	RiskUserActivityDetailuserPerformedSecuredPasswordReset         RiskUserActivityDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskUserActivityDetail() []string {
	return []string{
		string(RiskUserActivityDetailadminConfirmedServicePrincipalCompromised),
		string(RiskUserActivityDetailadminConfirmedSigninCompromised),
		string(RiskUserActivityDetailadminConfirmedSigninSafe),
		string(RiskUserActivityDetailadminConfirmedUserCompromised),
		string(RiskUserActivityDetailadminDismissedAllRiskForServicePrincipal),
		string(RiskUserActivityDetailadminDismissedAllRiskForUser),
		string(RiskUserActivityDetailadminGeneratedTemporaryPassword),
		string(RiskUserActivityDetailaiConfirmedSigninSafe),
		string(RiskUserActivityDetailhidden),
		string(RiskUserActivityDetailm365DAdminDismissedDetection),
		string(RiskUserActivityDetailnone),
		string(RiskUserActivityDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(RiskUserActivityDetailuserPerformedSecuredPasswordChange),
		string(RiskUserActivityDetailuserPerformedSecuredPasswordReset),
	}
}

func parseRiskUserActivityDetail(input string) (*RiskUserActivityDetail, error) {
	vals := map[string]RiskUserActivityDetail{
		"adminconfirmedserviceprincipalcompromised": RiskUserActivityDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskUserActivityDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskUserActivityDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskUserActivityDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskUserActivityDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskUserActivityDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskUserActivityDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskUserActivityDetailaiConfirmedSigninSafe,
		"hidden":                                    RiskUserActivityDetailhidden,
		"m365dadmindismisseddetection":              RiskUserActivityDetailm365DAdminDismissedDetection,
		"none":                                      RiskUserActivityDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskUserActivityDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskUserActivityDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskUserActivityDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskUserActivityDetail(input)
	return &out, nil
}
