package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskDetail string

const (
	SignInRiskDetailadminConfirmedServicePrincipalCompromised SignInRiskDetail = "AdminConfirmedServicePrincipalCompromised"
	SignInRiskDetailadminConfirmedSigninCompromised           SignInRiskDetail = "AdminConfirmedSigninCompromised"
	SignInRiskDetailadminConfirmedSigninSafe                  SignInRiskDetail = "AdminConfirmedSigninSafe"
	SignInRiskDetailadminConfirmedUserCompromised             SignInRiskDetail = "AdminConfirmedUserCompromised"
	SignInRiskDetailadminDismissedAllRiskForServicePrincipal  SignInRiskDetail = "AdminDismissedAllRiskForServicePrincipal"
	SignInRiskDetailadminDismissedAllRiskForUser              SignInRiskDetail = "AdminDismissedAllRiskForUser"
	SignInRiskDetailadminGeneratedTemporaryPassword           SignInRiskDetail = "AdminGeneratedTemporaryPassword"
	SignInRiskDetailaiConfirmedSigninSafe                     SignInRiskDetail = "AiConfirmedSigninSafe"
	SignInRiskDetailhidden                                    SignInRiskDetail = "Hidden"
	SignInRiskDetailm365DAdminDismissedDetection              SignInRiskDetail = "M365DAdminDismissedDetection"
	SignInRiskDetailnone                                      SignInRiskDetail = "None"
	SignInRiskDetailuserPassedMFADrivenByRiskBasedPolicy      SignInRiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	SignInRiskDetailuserPerformedSecuredPasswordChange        SignInRiskDetail = "UserPerformedSecuredPasswordChange"
	SignInRiskDetailuserPerformedSecuredPasswordReset         SignInRiskDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForSignInRiskDetail() []string {
	return []string{
		string(SignInRiskDetailadminConfirmedServicePrincipalCompromised),
		string(SignInRiskDetailadminConfirmedSigninCompromised),
		string(SignInRiskDetailadminConfirmedSigninSafe),
		string(SignInRiskDetailadminConfirmedUserCompromised),
		string(SignInRiskDetailadminDismissedAllRiskForServicePrincipal),
		string(SignInRiskDetailadminDismissedAllRiskForUser),
		string(SignInRiskDetailadminGeneratedTemporaryPassword),
		string(SignInRiskDetailaiConfirmedSigninSafe),
		string(SignInRiskDetailhidden),
		string(SignInRiskDetailm365DAdminDismissedDetection),
		string(SignInRiskDetailnone),
		string(SignInRiskDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(SignInRiskDetailuserPerformedSecuredPasswordChange),
		string(SignInRiskDetailuserPerformedSecuredPasswordReset),
	}
}

func parseSignInRiskDetail(input string) (*SignInRiskDetail, error) {
	vals := map[string]SignInRiskDetail{
		"adminconfirmedserviceprincipalcompromised": SignInRiskDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           SignInRiskDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  SignInRiskDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             SignInRiskDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  SignInRiskDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              SignInRiskDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           SignInRiskDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     SignInRiskDetailaiConfirmedSigninSafe,
		"hidden":                                    SignInRiskDetailhidden,
		"m365dadmindismisseddetection":              SignInRiskDetailm365DAdminDismissedDetection,
		"none":                                      SignInRiskDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      SignInRiskDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        SignInRiskDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         SignInRiskDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskDetail(input)
	return &out, nil
}
