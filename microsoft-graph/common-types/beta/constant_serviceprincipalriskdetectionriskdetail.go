package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionRiskDetail string

const (
	ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedServicePrincipalCompromised ServicePrincipalRiskDetectionRiskDetail = "adminConfirmedServicePrincipalCompromised"
	ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedSigninCompromised           ServicePrincipalRiskDetectionRiskDetail = "adminConfirmedSigninCompromised"
	ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedSigninSafe                  ServicePrincipalRiskDetectionRiskDetail = "adminConfirmedSigninSafe"
	ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedUserCompromised             ServicePrincipalRiskDetectionRiskDetail = "adminConfirmedUserCompromised"
	ServicePrincipalRiskDetectionRiskDetail_AdminDismissedAllRiskForServicePrincipal  ServicePrincipalRiskDetectionRiskDetail = "adminDismissedAllRiskForServicePrincipal"
	ServicePrincipalRiskDetectionRiskDetail_AdminDismissedAllRiskForUser              ServicePrincipalRiskDetectionRiskDetail = "adminDismissedAllRiskForUser"
	ServicePrincipalRiskDetectionRiskDetail_AdminGeneratedTemporaryPassword           ServicePrincipalRiskDetectionRiskDetail = "adminGeneratedTemporaryPassword"
	ServicePrincipalRiskDetectionRiskDetail_AiConfirmedSigninSafe                     ServicePrincipalRiskDetectionRiskDetail = "aiConfirmedSigninSafe"
	ServicePrincipalRiskDetectionRiskDetail_Hidden                                    ServicePrincipalRiskDetectionRiskDetail = "hidden"
	ServicePrincipalRiskDetectionRiskDetail_M365DAdminDismissedDetection              ServicePrincipalRiskDetectionRiskDetail = "m365DAdminDismissedDetection"
	ServicePrincipalRiskDetectionRiskDetail_None                                      ServicePrincipalRiskDetectionRiskDetail = "none"
	ServicePrincipalRiskDetectionRiskDetail_UserPassedMFADrivenByRiskBasedPolicy      ServicePrincipalRiskDetectionRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	ServicePrincipalRiskDetectionRiskDetail_UserPerformedSecuredPasswordChange        ServicePrincipalRiskDetectionRiskDetail = "userPerformedSecuredPasswordChange"
	ServicePrincipalRiskDetectionRiskDetail_UserPerformedSecuredPasswordReset         ServicePrincipalRiskDetectionRiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForServicePrincipalRiskDetectionRiskDetail() []string {
	return []string{
		string(ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedSigninCompromised),
		string(ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedSigninSafe),
		string(ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedUserCompromised),
		string(ServicePrincipalRiskDetectionRiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(ServicePrincipalRiskDetectionRiskDetail_AdminDismissedAllRiskForUser),
		string(ServicePrincipalRiskDetectionRiskDetail_AdminGeneratedTemporaryPassword),
		string(ServicePrincipalRiskDetectionRiskDetail_AiConfirmedSigninSafe),
		string(ServicePrincipalRiskDetectionRiskDetail_Hidden),
		string(ServicePrincipalRiskDetectionRiskDetail_M365DAdminDismissedDetection),
		string(ServicePrincipalRiskDetectionRiskDetail_None),
		string(ServicePrincipalRiskDetectionRiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(ServicePrincipalRiskDetectionRiskDetail_UserPerformedSecuredPasswordChange),
		string(ServicePrincipalRiskDetectionRiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *ServicePrincipalRiskDetectionRiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePrincipalRiskDetectionRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePrincipalRiskDetectionRiskDetail(input string) (*ServicePrincipalRiskDetectionRiskDetail, error) {
	vals := map[string]ServicePrincipalRiskDetectionRiskDetail{
		"adminconfirmedserviceprincipalcompromised": ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             ServicePrincipalRiskDetectionRiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  ServicePrincipalRiskDetectionRiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              ServicePrincipalRiskDetectionRiskDetail_AdminDismissedAllRiskForUser,
		"admingeneratedtemporarypassword":           ServicePrincipalRiskDetectionRiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     ServicePrincipalRiskDetectionRiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    ServicePrincipalRiskDetectionRiskDetail_Hidden,
		"m365dadmindismisseddetection":              ServicePrincipalRiskDetectionRiskDetail_M365DAdminDismissedDetection,
		"none":                                      ServicePrincipalRiskDetectionRiskDetail_None,
		"userpassedmfadrivenbyriskbasedpolicy":      ServicePrincipalRiskDetectionRiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        ServicePrincipalRiskDetectionRiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         ServicePrincipalRiskDetectionRiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionRiskDetail(input)
	return &out, nil
}
