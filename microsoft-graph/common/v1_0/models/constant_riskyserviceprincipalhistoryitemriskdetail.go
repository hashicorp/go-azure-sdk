package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalHistoryItemRiskDetail string

const (
	RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedServicePrincipalCompromised RiskyServicePrincipalHistoryItemRiskDetail = "AdminConfirmedServicePrincipalCompromised"
	RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedSigninCompromised           RiskyServicePrincipalHistoryItemRiskDetail = "AdminConfirmedSigninCompromised"
	RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedSigninSafe                  RiskyServicePrincipalHistoryItemRiskDetail = "AdminConfirmedSigninSafe"
	RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedUserCompromised             RiskyServicePrincipalHistoryItemRiskDetail = "AdminConfirmedUserCompromised"
	RiskyServicePrincipalHistoryItemRiskDetailadminDismissedAllRiskForServicePrincipal  RiskyServicePrincipalHistoryItemRiskDetail = "AdminDismissedAllRiskForServicePrincipal"
	RiskyServicePrincipalHistoryItemRiskDetailadminDismissedAllRiskForUser              RiskyServicePrincipalHistoryItemRiskDetail = "AdminDismissedAllRiskForUser"
	RiskyServicePrincipalHistoryItemRiskDetailadminGeneratedTemporaryPassword           RiskyServicePrincipalHistoryItemRiskDetail = "AdminGeneratedTemporaryPassword"
	RiskyServicePrincipalHistoryItemRiskDetailaiConfirmedSigninSafe                     RiskyServicePrincipalHistoryItemRiskDetail = "AiConfirmedSigninSafe"
	RiskyServicePrincipalHistoryItemRiskDetailhidden                                    RiskyServicePrincipalHistoryItemRiskDetail = "Hidden"
	RiskyServicePrincipalHistoryItemRiskDetailm365DAdminDismissedDetection              RiskyServicePrincipalHistoryItemRiskDetail = "M365DAdminDismissedDetection"
	RiskyServicePrincipalHistoryItemRiskDetailnone                                      RiskyServicePrincipalHistoryItemRiskDetail = "None"
	RiskyServicePrincipalHistoryItemRiskDetailuserPassedMFADrivenByRiskBasedPolicy      RiskyServicePrincipalHistoryItemRiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	RiskyServicePrincipalHistoryItemRiskDetailuserPerformedSecuredPasswordChange        RiskyServicePrincipalHistoryItemRiskDetail = "UserPerformedSecuredPasswordChange"
	RiskyServicePrincipalHistoryItemRiskDetailuserPerformedSecuredPasswordReset         RiskyServicePrincipalHistoryItemRiskDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyServicePrincipalHistoryItemRiskDetail() []string {
	return []string{
		string(RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedServicePrincipalCompromised),
		string(RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedSigninCompromised),
		string(RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedSigninSafe),
		string(RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedUserCompromised),
		string(RiskyServicePrincipalHistoryItemRiskDetailadminDismissedAllRiskForServicePrincipal),
		string(RiskyServicePrincipalHistoryItemRiskDetailadminDismissedAllRiskForUser),
		string(RiskyServicePrincipalHistoryItemRiskDetailadminGeneratedTemporaryPassword),
		string(RiskyServicePrincipalHistoryItemRiskDetailaiConfirmedSigninSafe),
		string(RiskyServicePrincipalHistoryItemRiskDetailhidden),
		string(RiskyServicePrincipalHistoryItemRiskDetailm365DAdminDismissedDetection),
		string(RiskyServicePrincipalHistoryItemRiskDetailnone),
		string(RiskyServicePrincipalHistoryItemRiskDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyServicePrincipalHistoryItemRiskDetailuserPerformedSecuredPasswordChange),
		string(RiskyServicePrincipalHistoryItemRiskDetailuserPerformedSecuredPasswordReset),
	}
}

func parseRiskyServicePrincipalHistoryItemRiskDetail(input string) (*RiskyServicePrincipalHistoryItemRiskDetail, error) {
	vals := map[string]RiskyServicePrincipalHistoryItemRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyServicePrincipalHistoryItemRiskDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyServicePrincipalHistoryItemRiskDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyServicePrincipalHistoryItemRiskDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyServicePrincipalHistoryItemRiskDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyServicePrincipalHistoryItemRiskDetailaiConfirmedSigninSafe,
		"hidden":                                    RiskyServicePrincipalHistoryItemRiskDetailhidden,
		"m365dadmindismisseddetection":              RiskyServicePrincipalHistoryItemRiskDetailm365DAdminDismissedDetection,
		"none":                                      RiskyServicePrincipalHistoryItemRiskDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyServicePrincipalHistoryItemRiskDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyServicePrincipalHistoryItemRiskDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyServicePrincipalHistoryItemRiskDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalHistoryItemRiskDetail(input)
	return &out, nil
}
