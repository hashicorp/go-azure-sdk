package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItemRiskDetail string

const (
	RiskyUserHistoryItemRiskDetailadminConfirmedServicePrincipalCompromised RiskyUserHistoryItemRiskDetail = "AdminConfirmedServicePrincipalCompromised"
	RiskyUserHistoryItemRiskDetailadminConfirmedSigninCompromised           RiskyUserHistoryItemRiskDetail = "AdminConfirmedSigninCompromised"
	RiskyUserHistoryItemRiskDetailadminConfirmedSigninSafe                  RiskyUserHistoryItemRiskDetail = "AdminConfirmedSigninSafe"
	RiskyUserHistoryItemRiskDetailadminConfirmedUserCompromised             RiskyUserHistoryItemRiskDetail = "AdminConfirmedUserCompromised"
	RiskyUserHistoryItemRiskDetailadminDismissedAllRiskForServicePrincipal  RiskyUserHistoryItemRiskDetail = "AdminDismissedAllRiskForServicePrincipal"
	RiskyUserHistoryItemRiskDetailadminDismissedAllRiskForUser              RiskyUserHistoryItemRiskDetail = "AdminDismissedAllRiskForUser"
	RiskyUserHistoryItemRiskDetailadminGeneratedTemporaryPassword           RiskyUserHistoryItemRiskDetail = "AdminGeneratedTemporaryPassword"
	RiskyUserHistoryItemRiskDetailaiConfirmedSigninSafe                     RiskyUserHistoryItemRiskDetail = "AiConfirmedSigninSafe"
	RiskyUserHistoryItemRiskDetailhidden                                    RiskyUserHistoryItemRiskDetail = "Hidden"
	RiskyUserHistoryItemRiskDetailm365DAdminDismissedDetection              RiskyUserHistoryItemRiskDetail = "M365DAdminDismissedDetection"
	RiskyUserHistoryItemRiskDetailnone                                      RiskyUserHistoryItemRiskDetail = "None"
	RiskyUserHistoryItemRiskDetailuserPassedMFADrivenByRiskBasedPolicy      RiskyUserHistoryItemRiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	RiskyUserHistoryItemRiskDetailuserPerformedSecuredPasswordChange        RiskyUserHistoryItemRiskDetail = "UserPerformedSecuredPasswordChange"
	RiskyUserHistoryItemRiskDetailuserPerformedSecuredPasswordReset         RiskyUserHistoryItemRiskDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskyUserHistoryItemRiskDetail() []string {
	return []string{
		string(RiskyUserHistoryItemRiskDetailadminConfirmedServicePrincipalCompromised),
		string(RiskyUserHistoryItemRiskDetailadminConfirmedSigninCompromised),
		string(RiskyUserHistoryItemRiskDetailadminConfirmedSigninSafe),
		string(RiskyUserHistoryItemRiskDetailadminConfirmedUserCompromised),
		string(RiskyUserHistoryItemRiskDetailadminDismissedAllRiskForServicePrincipal),
		string(RiskyUserHistoryItemRiskDetailadminDismissedAllRiskForUser),
		string(RiskyUserHistoryItemRiskDetailadminGeneratedTemporaryPassword),
		string(RiskyUserHistoryItemRiskDetailaiConfirmedSigninSafe),
		string(RiskyUserHistoryItemRiskDetailhidden),
		string(RiskyUserHistoryItemRiskDetailm365DAdminDismissedDetection),
		string(RiskyUserHistoryItemRiskDetailnone),
		string(RiskyUserHistoryItemRiskDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(RiskyUserHistoryItemRiskDetailuserPerformedSecuredPasswordChange),
		string(RiskyUserHistoryItemRiskDetailuserPerformedSecuredPasswordReset),
	}
}

func parseRiskyUserHistoryItemRiskDetail(input string) (*RiskyUserHistoryItemRiskDetail, error) {
	vals := map[string]RiskyUserHistoryItemRiskDetail{
		"adminconfirmedserviceprincipalcompromised": RiskyUserHistoryItemRiskDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskyUserHistoryItemRiskDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskyUserHistoryItemRiskDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskyUserHistoryItemRiskDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskyUserHistoryItemRiskDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskyUserHistoryItemRiskDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskyUserHistoryItemRiskDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskyUserHistoryItemRiskDetailaiConfirmedSigninSafe,
		"hidden":                                    RiskyUserHistoryItemRiskDetailhidden,
		"m365dadmindismisseddetection":              RiskyUserHistoryItemRiskDetailm365DAdminDismissedDetection,
		"none":                                      RiskyUserHistoryItemRiskDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskyUserHistoryItemRiskDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskyUserHistoryItemRiskDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskyUserHistoryItemRiskDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserHistoryItemRiskDetail(input)
	return &out, nil
}
