package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserRiskDetail string

const (
	RiskyUserRiskDetailadminConfirmedServicePrincipalCompromised RiskyUserRiskDetail = "AdminConfirmedServicePrincipalCompromised"
	RiskyUserRiskDetailadminConfirmedSigninCompromised           RiskyUserRiskDetail = "AdminConfirmedSigninCompromised"
	RiskyUserRiskDetailadminConfirmedSigninSafe                  RiskyUserRiskDetail = "AdminConfirmedSigninSafe"
	RiskyUserRiskDetailadminConfirmedUserCompromised             RiskyUserRiskDetail = "AdminConfirmedUserCompromised"
	RiskyUserRiskDetailadminDismissedAllRiskForServicePrincipal  RiskyUserRiskDetail = "AdminDismissedAllRiskForServicePrincipal"
	RiskyUserRiskDetailadminDismissedAllRiskForUser              RiskyUserRiskDetail = "AdminDismissedAllRiskForUser"
	RiskyUserRiskDetailadminGeneratedTemporaryPassword           RiskyUserRiskDetail = "AdminGeneratedTemporaryPassword"
	RiskyUserRiskDetailaiConfirmedSigninSafe                     RiskyUserRiskDetail = "AiConfirmedSigninSafe"
	RiskyUserRiskDetailhidden                                    RiskyUserRiskDetail = "Hidden"
	RiskyUserRiskDetailm365DAdminDismissedDetection              RiskyUserRiskDetail = "M365DAdminDismissedDetection"
	RiskyUserRiskDetailnone                                      RiskyUserRiskDetail = "None"
	RiskyUserRiskDetailuserPassedMFADrivenByRiskBasedPolicy      RiskyUserRiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	RiskyUserRiskDetailuserPerformedSecuredPasswordChange        RiskyUserRiskDetail = "UserPerformedSecuredPasswordChange"
	RiskyUserRiskDetailuserPerformedSecuredPasswordReset         RiskyUserRiskDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyUserRiskDetail() []string {
	return []string{
		string(RiskyUserRiskDetailadminConfirmedServicePrincipalCompromised),
		string(RiskyUserRiskDetailadminConfirmedSigninCompromised),
		string(RiskyUserRiskDetailadminConfirmedSigninSafe),
		string(RiskyUserRiskDetailadminConfirmedUserCompromised),
		string(RiskyUserRiskDetailadminDismissedAllRiskForServicePrincipal),
		string(RiskyUserRiskDetailadminDismissedAllRiskForUser),
		string(RiskyUserRiskDetailadminGeneratedTemporaryPassword),
		string(RiskyUserRiskDetailaiConfirmedSigninSafe),
		string(RiskyUserRiskDetailhidden),
		string(RiskyUserRiskDetailm365DAdminDismissedDetection),
		string(RiskyUserRiskDetailnone),
		string(RiskyUserRiskDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyUserRiskDetailuserPerformedSecuredPasswordChange),
		string(RiskyUserRiskDetailuserPerformedSecuredPasswordReset),
	}
}

func parseRiskyUserRiskDetail(input string) (*RiskyUserRiskDetail, error) {
	vals := map[string]RiskyUserRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyUserRiskDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyUserRiskDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyUserRiskDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyUserRiskDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyUserRiskDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyUserRiskDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyUserRiskDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyUserRiskDetailaiConfirmedSigninSafe,
		"hidden":                                    RiskyUserRiskDetailhidden,
		"m365dadmindismisseddetection":              RiskyUserRiskDetailm365DAdminDismissedDetection,
		"none":                                      RiskyUserRiskDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyUserRiskDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyUserRiskDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyUserRiskDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserRiskDetail(input)
	return &out, nil
}
