package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionRiskDetail string

const (
	ServicePrincipalRiskDetectionRiskDetailadminConfirmedServicePrincipalCompromised ServicePrincipalRiskDetectionRiskDetail = "AdminConfirmedServicePrincipalCompromised"
	ServicePrincipalRiskDetectionRiskDetailadminConfirmedSigninCompromised           ServicePrincipalRiskDetectionRiskDetail = "AdminConfirmedSigninCompromised"
	ServicePrincipalRiskDetectionRiskDetailadminConfirmedSigninSafe                  ServicePrincipalRiskDetectionRiskDetail = "AdminConfirmedSigninSafe"
	ServicePrincipalRiskDetectionRiskDetailadminConfirmedUserCompromised             ServicePrincipalRiskDetectionRiskDetail = "AdminConfirmedUserCompromised"
	ServicePrincipalRiskDetectionRiskDetailadminDismissedAllRiskForServicePrincipal  ServicePrincipalRiskDetectionRiskDetail = "AdminDismissedAllRiskForServicePrincipal"
	ServicePrincipalRiskDetectionRiskDetailadminDismissedAllRiskForUser              ServicePrincipalRiskDetectionRiskDetail = "AdminDismissedAllRiskForUser"
	ServicePrincipalRiskDetectionRiskDetailadminGeneratedTemporaryPassword           ServicePrincipalRiskDetectionRiskDetail = "AdminGeneratedTemporaryPassword"
	ServicePrincipalRiskDetectionRiskDetailaiConfirmedSigninSafe                     ServicePrincipalRiskDetectionRiskDetail = "AiConfirmedSigninSafe"
	ServicePrincipalRiskDetectionRiskDetailhidden                                    ServicePrincipalRiskDetectionRiskDetail = "Hidden"
	ServicePrincipalRiskDetectionRiskDetailm365DAdminDismissedDetection              ServicePrincipalRiskDetectionRiskDetail = "M365DAdminDismissedDetection"
	ServicePrincipalRiskDetectionRiskDetailnone                                      ServicePrincipalRiskDetectionRiskDetail = "None"
	ServicePrincipalRiskDetectionRiskDetailuserPassedMFADrivenByRiskBasedPolicy      ServicePrincipalRiskDetectionRiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	ServicePrincipalRiskDetectionRiskDetailuserPerformedSecuredPasswordChange        ServicePrincipalRiskDetectionRiskDetail = "UserPerformedSecuredPasswordChange"
	ServicePrincipalRiskDetectionRiskDetailuserPerformedSecuredPasswordReset         ServicePrincipalRiskDetectionRiskDetail = "UserPerformedSecuredPasswordReset"
)

func PossibleValuesForServicePrincipalRiskDetectionRiskDetail() []string {
	return []string{
		string(ServicePrincipalRiskDetectionRiskDetailadminConfirmedServicePrincipalCompromised),
		string(ServicePrincipalRiskDetectionRiskDetailadminConfirmedSigninCompromised),
		string(ServicePrincipalRiskDetectionRiskDetailadminConfirmedSigninSafe),
		string(ServicePrincipalRiskDetectionRiskDetailadminConfirmedUserCompromised),
		string(ServicePrincipalRiskDetectionRiskDetailadminDismissedAllRiskForServicePrincipal),
		string(ServicePrincipalRiskDetectionRiskDetailadminDismissedAllRiskForUser),
		string(ServicePrincipalRiskDetectionRiskDetailadminGeneratedTemporaryPassword),
		string(ServicePrincipalRiskDetectionRiskDetailaiConfirmedSigninSafe),
		string(ServicePrincipalRiskDetectionRiskDetailhidden),
		string(ServicePrincipalRiskDetectionRiskDetailm365DAdminDismissedDetection),
		string(ServicePrincipalRiskDetectionRiskDetailnone),
		string(ServicePrincipalRiskDetectionRiskDetailuserPassedMFADrivenByRiskBasedPolicy),
		string(ServicePrincipalRiskDetectionRiskDetailuserPerformedSecuredPasswordChange),
		string(ServicePrincipalRiskDetectionRiskDetailuserPerformedSecuredPasswordReset),
	}
}

func parseServicePrincipalRiskDetectionRiskDetail(input string) (*ServicePrincipalRiskDetectionRiskDetail, error) {
	vals := map[string]ServicePrincipalRiskDetectionRiskDetail{
		"adminconfirmedserviceprincipalcompromised": ServicePrincipalRiskDetectionRiskDetailadminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           ServicePrincipalRiskDetectionRiskDetailadminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  ServicePrincipalRiskDetectionRiskDetailadminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             ServicePrincipalRiskDetectionRiskDetailadminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  ServicePrincipalRiskDetectionRiskDetailadminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              ServicePrincipalRiskDetectionRiskDetailadminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           ServicePrincipalRiskDetectionRiskDetailadminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     ServicePrincipalRiskDetectionRiskDetailaiConfirmedSigninSafe,
		"hidden":                                    ServicePrincipalRiskDetectionRiskDetailhidden,
		"m365dadmindismisseddetection":              ServicePrincipalRiskDetectionRiskDetailm365DAdminDismissedDetection,
		"none":                                      ServicePrincipalRiskDetectionRiskDetailnone,
		"userpassedmfadrivenbyriskbasedpolicy":      ServicePrincipalRiskDetectionRiskDetailuserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        ServicePrincipalRiskDetectionRiskDetailuserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         ServicePrincipalRiskDetectionRiskDetailuserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionRiskDetail(input)
	return &out, nil
}
