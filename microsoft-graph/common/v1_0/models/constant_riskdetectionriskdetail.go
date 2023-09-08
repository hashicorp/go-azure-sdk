package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskDetail string

const (
	RiskDetectionRiskDetailadminConfirmedServicePrincipalCompromised RiskDetectionRiskDetail = "AdminConfirmedServicePrincipalCompromised"
	RiskDetectionRiskDetailadminConfirmedSigninCompromised           RiskDetectionRiskDetail = "AdminConfirmedSigninCompromised"
	RiskDetectionRiskDetailadminConfirmedSigninSafe                  RiskDetectionRiskDetail = "AdminConfirmedSigninSafe"
	RiskDetectionRiskDetailadminConfirmedUserCompromised             RiskDetectionRiskDetail = "AdminConfirmedUserCompromised"
	RiskDetectionRiskDetailadminDismissedAllRiskForServicePrincipal  RiskDetectionRiskDetail = "AdminDismissedAllRiskForServicePrincipal"
	RiskDetectionRiskDetailadminDismissedAllRiskForUser              RiskDetectionRiskDetail = "AdminDismissedAllRiskForUser"
	RiskDetectionRiskDetailadminGeneratedTemporaryPassword           RiskDetectionRiskDetail = "AdminGeneratedTemporaryPassword"
	RiskDetectionRiskDetailaiConfirmedSigninSafe                     RiskDetectionRiskDetail = "AiConfirmedSigninSafe"
	RiskDetectionRiskDetailhidden                                    RiskDetectionRiskDetail = "Hidden"
	RiskDetectionRiskDetailm365DAdminDismissedDetection              RiskDetectionRiskDetail = "M365DAdminDismissedDetection"
	RiskDetectionRiskDetailnone                                      RiskDetectionRiskDetail = "None"
	RiskDetectionRiskDetailuserPassedMFADrivenByRiskBasedPolicy      RiskDetectionRiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	RiskDetectionRiskDetailuserPerformedSecuredPasswordChange        RiskDetectionRiskDetail = "UserPerformedSecuredPasswordChange"
	RiskDetectionRiskDetailuserPerformedSecuredPasswordReset         RiskDetectionRiskDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskDetectionRiskDetail() []string {
	return []string{
		string(RiskDetectionRiskDetailadminConfirmedServicePrincipalCompromised),
		string(RiskDetectionRiskDetailadminConfirmedSigninCompromised),
		string(RiskDetectionRiskDetailadminConfirmedSigninSafe),
		string(RiskDetectionRiskDetailadminConfirmedUserCompromised),
		string(RiskDetectionRiskDetailadminDismissedAllRiskForServicePrincipal),
		string(RiskDetectionRiskDetailadminDismissedAllRiskForUser),
		string(RiskDetectionRiskDetailadminGeneratedTemporaryPassword),
		string(RiskDetectionRiskDetailaiConfirmedSigninSafe),
		string(RiskDetectionRiskDetailhidden),
		string(RiskDetectionRiskDetailm365DAdminDismissedDetection),
		string(RiskDetectionRiskDetailnone),
		string(RiskDetectionRiskDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(RiskDetectionRiskDetailuserPerformedSecuredPasswordChange),
		string(RiskDetectionRiskDetailuserPerformedSecuredPasswordReset),
	}
}

func parseRiskDetectionRiskDetail(input string) (*RiskDetectionRiskDetail, error) {
	vals := map[string]RiskDetectionRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskDetectionRiskDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskDetectionRiskDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskDetectionRiskDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskDetectionRiskDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskDetectionRiskDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskDetectionRiskDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskDetectionRiskDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskDetectionRiskDetailaiConfirmedSigninSafe,
		"hidden":                                    RiskDetectionRiskDetailhidden,
		"m365dadmindismisseddetection":              RiskDetectionRiskDetailm365DAdminDismissedDetection,
		"none":                                      RiskDetectionRiskDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskDetectionRiskDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskDetectionRiskDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskDetectionRiskDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskDetail(input)
	return &out, nil
}
