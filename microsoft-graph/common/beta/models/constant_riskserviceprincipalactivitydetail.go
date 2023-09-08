package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskServicePrincipalActivityDetail string

const (
	RiskServicePrincipalActivityDetailadminConfirmedServicePrincipalCompromised RiskServicePrincipalActivityDetail = "AdminConfirmedServicePrincipalCompromised"
	RiskServicePrincipalActivityDetailadminConfirmedSigninCompromised           RiskServicePrincipalActivityDetail = "AdminConfirmedSigninCompromised"
	RiskServicePrincipalActivityDetailadminConfirmedSigninSafe                  RiskServicePrincipalActivityDetail = "AdminConfirmedSigninSafe"
	RiskServicePrincipalActivityDetailadminConfirmedUserCompromised             RiskServicePrincipalActivityDetail = "AdminConfirmedUserCompromised"
	RiskServicePrincipalActivityDetailadminDismissedAllRiskForServicePrincipal  RiskServicePrincipalActivityDetail = "AdminDismissedAllRiskForServicePrincipal"
	RiskServicePrincipalActivityDetailadminDismissedAllRiskForUser              RiskServicePrincipalActivityDetail = "AdminDismissedAllRiskForUser"
	RiskServicePrincipalActivityDetailadminGeneratedTemporaryPassword           RiskServicePrincipalActivityDetail = "AdminGeneratedTemporaryPassword"
	RiskServicePrincipalActivityDetailaiConfirmedSigninSafe                     RiskServicePrincipalActivityDetail = "AiConfirmedSigninSafe"
	RiskServicePrincipalActivityDetailhidden                                    RiskServicePrincipalActivityDetail = "Hidden"
	RiskServicePrincipalActivityDetailm365DAdminDismissedDetection              RiskServicePrincipalActivityDetail = "M365DAdminDismissedDetection"
	RiskServicePrincipalActivityDetailnone                                      RiskServicePrincipalActivityDetail = "None"
	RiskServicePrincipalActivityDetailuserPassedMFADrivenByRiskBasedPolicy      RiskServicePrincipalActivityDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	RiskServicePrincipalActivityDetailuserPerformedSecuredPasswordChange        RiskServicePrincipalActivityDetail = "UserPerformedSecuredPasswordChange"
	RiskServicePrincipalActivityDetailuserPerformedSecuredPasswordReset         RiskServicePrincipalActivityDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskServicePrincipalActivityDetail() []string {
	return []string{
		string(RiskServicePrincipalActivityDetailadminConfirmedServicePrincipalCompromised),
		string(RiskServicePrincipalActivityDetailadminConfirmedSigninCompromised),
		string(RiskServicePrincipalActivityDetailadminConfirmedSigninSafe),
		string(RiskServicePrincipalActivityDetailadminConfirmedUserCompromised),
		string(RiskServicePrincipalActivityDetailadminDismissedAllRiskForServicePrincipal),
		string(RiskServicePrincipalActivityDetailadminDismissedAllRiskForUser),
		string(RiskServicePrincipalActivityDetailadminGeneratedTemporaryPassword),
		string(RiskServicePrincipalActivityDetailaiConfirmedSigninSafe),
		string(RiskServicePrincipalActivityDetailhidden),
		string(RiskServicePrincipalActivityDetailm365DAdminDismissedDetection),
		string(RiskServicePrincipalActivityDetailnone),
		string(RiskServicePrincipalActivityDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(RiskServicePrincipalActivityDetailuserPerformedSecuredPasswordChange),
		string(RiskServicePrincipalActivityDetailuserPerformedSecuredPasswordReset),
	}
}

func parseRiskServicePrincipalActivityDetail(input string) (*RiskServicePrincipalActivityDetail, error) {
	vals := map[string]RiskServicePrincipalActivityDetail{
		"adminconfirmedserviceprincipalcompromised": RiskServicePrincipalActivityDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskServicePrincipalActivityDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskServicePrincipalActivityDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskServicePrincipalActivityDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskServicePrincipalActivityDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskServicePrincipalActivityDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           RiskServicePrincipalActivityDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskServicePrincipalActivityDetailaiConfirmedSigninSafe,
		"hidden":                                    RiskServicePrincipalActivityDetailhidden,
		"m365dadmindismisseddetection":              RiskServicePrincipalActivityDetailm365DAdminDismissedDetection,
		"none":                                      RiskServicePrincipalActivityDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskServicePrincipalActivityDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskServicePrincipalActivityDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskServicePrincipalActivityDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskServicePrincipalActivityDetail(input)
	return &out, nil
}
