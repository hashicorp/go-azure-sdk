package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalRiskDetail string

const (
	RiskyServicePrincipalRiskDetailadminConfirmedServicePrincipalCompromised RiskyServicePrincipalRiskDetail = "AdminConfirmedServicePrincipalCompromised"
	RiskyServicePrincipalRiskDetailadminConfirmedSigninCompromised           RiskyServicePrincipalRiskDetail = "AdminConfirmedSigninCompromised"
	RiskyServicePrincipalRiskDetailadminConfirmedSigninSafe                  RiskyServicePrincipalRiskDetail = "AdminConfirmedSigninSafe"
	RiskyServicePrincipalRiskDetailadminConfirmedUserCompromised             RiskyServicePrincipalRiskDetail = "AdminConfirmedUserCompromised"
	RiskyServicePrincipalRiskDetailadminDismissedAllRiskForServicePrincipal  RiskyServicePrincipalRiskDetail = "AdminDismissedAllRiskForServicePrincipal"
	RiskyServicePrincipalRiskDetailadminDismissedAllRiskForUser              RiskyServicePrincipalRiskDetail = "AdminDismissedAllRiskForUser"
	RiskyServicePrincipalRiskDetailadminGeneratedTemporaryPassword           RiskyServicePrincipalRiskDetail = "AdminGeneratedTemporaryPassword"
	RiskyServicePrincipalRiskDetailaiConfirmedSigninSafe                     RiskyServicePrincipalRiskDetail = "AiConfirmedSigninSafe"
	RiskyServicePrincipalRiskDetailhidden                                    RiskyServicePrincipalRiskDetail = "Hidden"
	RiskyServicePrincipalRiskDetailm365DAdminDismissedDetection              RiskyServicePrincipalRiskDetail = "M365DAdminDismissedDetection"
	RiskyServicePrincipalRiskDetailnone                                      RiskyServicePrincipalRiskDetail = "None"
	RiskyServicePrincipalRiskDetailuserPassedMFADrivenByRiskBasedPolicy      RiskyServicePrincipalRiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	RiskyServicePrincipalRiskDetailuserPerformedSecuredPasswordChange        RiskyServicePrincipalRiskDetail = "UserPerformedSecuredPasswordChange"
	RiskyServicePrincipalRiskDetailuserPerformedSecuredPasswordReset         RiskyServicePrincipalRiskDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyServicePrincipalRiskDetail() []string {
	return []string{
		string(RiskyServicePrincipalRiskDetailadminConfirmedServicePrincipalCompromised),
		string(RiskyServicePrincipalRiskDetailadminConfirmedSigninCompromised),
		string(RiskyServicePrincipalRiskDetailadminConfirmedSigninSafe),
		string(RiskyServicePrincipalRiskDetailadminConfirmedUserCompromised),
		string(RiskyServicePrincipalRiskDetailadminDismissedAllRiskForServicePrincipal),
		string(RiskyServicePrincipalRiskDetailadminDismissedAllRiskForUser),
		string(RiskyServicePrincipalRiskDetailadminGeneratedTemporaryPassword),
		string(RiskyServicePrincipalRiskDetailaiConfirmedSigninSafe),
		string(RiskyServicePrincipalRiskDetailhidden),
		string(RiskyServicePrincipalRiskDetailm365DAdminDismissedDetection),
		string(RiskyServicePrincipalRiskDetailnone),
		string(RiskyServicePrincipalRiskDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyServicePrincipalRiskDetailuserPerformedSecuredPasswordChange),
		string(RiskyServicePrincipalRiskDetailuserPerformedSecuredPasswordReset),
	}
}

func parseRiskyServicePrincipalRiskDetail(input string) (*RiskyServicePrincipalRiskDetail, error) {
	vals := map[string]RiskyServicePrincipalRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyServicePrincipalRiskDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyServicePrincipalRiskDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyServicePrincipalRiskDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyServicePrincipalRiskDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyServicePrincipalRiskDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyServicePrincipalRiskDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyServicePrincipalRiskDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyServicePrincipalRiskDetailaiConfirmedSigninSafe,
		"hidden":                                    RiskyServicePrincipalRiskDetailhidden,
		"m365dadmindismisseddetection":              RiskyServicePrincipalRiskDetailm365DAdminDismissedDetection,
		"none":                                      RiskyServicePrincipalRiskDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyServicePrincipalRiskDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyServicePrincipalRiskDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyServicePrincipalRiskDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalRiskDetail(input)
	return &out, nil
}
