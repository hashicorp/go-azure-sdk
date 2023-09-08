package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenRevokeLicensesActionResultActionFailureReason string

const (
	VppTokenRevokeLicensesActionResultActionFailureReasonappleFailure                            VppTokenRevokeLicensesActionResultActionFailureReason = "AppleFailure"
	VppTokenRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate VppTokenRevokeLicensesActionResultActionFailureReason = "ExpiredApplePushNotificationCertificate"
	VppTokenRevokeLicensesActionResultActionFailureReasonexpiredVppToken                         VppTokenRevokeLicensesActionResultActionFailureReason = "ExpiredVppToken"
	VppTokenRevokeLicensesActionResultActionFailureReasoninternalError                           VppTokenRevokeLicensesActionResultActionFailureReason = "InternalError"
	VppTokenRevokeLicensesActionResultActionFailureReasonnone                                    VppTokenRevokeLicensesActionResultActionFailureReason = "None"
)

func PossibleValuesForVppTokenRevokeLicensesActionResultActionFailureReason() []string {
	return []string{
		string(VppTokenRevokeLicensesActionResultActionFailureReasonappleFailure),
		string(VppTokenRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate),
		string(VppTokenRevokeLicensesActionResultActionFailureReasonexpiredVppToken),
		string(VppTokenRevokeLicensesActionResultActionFailureReasoninternalError),
		string(VppTokenRevokeLicensesActionResultActionFailureReasonnone),
	}
}

func parseVppTokenRevokeLicensesActionResultActionFailureReason(input string) (*VppTokenRevokeLicensesActionResultActionFailureReason, error) {
	vals := map[string]VppTokenRevokeLicensesActionResultActionFailureReason{
		"applefailure": VppTokenRevokeLicensesActionResultActionFailureReasonappleFailure,
		"expiredapplepushnotificationcertificate": VppTokenRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate,
		"expiredvpptoken":                         VppTokenRevokeLicensesActionResultActionFailureReasonexpiredVppToken,
		"internalerror":                           VppTokenRevokeLicensesActionResultActionFailureReasoninternalError,
		"none":                                    VppTokenRevokeLicensesActionResultActionFailureReasonnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenRevokeLicensesActionResultActionFailureReason(input)
	return &out, nil
}
