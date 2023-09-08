package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppRevokeLicensesActionResultActionFailureReason string

const (
	MacOsVppAppRevokeLicensesActionResultActionFailureReasonappleFailure                            MacOsVppAppRevokeLicensesActionResultActionFailureReason = "AppleFailure"
	MacOsVppAppRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate MacOsVppAppRevokeLicensesActionResultActionFailureReason = "ExpiredApplePushNotificationCertificate"
	MacOsVppAppRevokeLicensesActionResultActionFailureReasonexpiredVppToken                         MacOsVppAppRevokeLicensesActionResultActionFailureReason = "ExpiredVppToken"
	MacOsVppAppRevokeLicensesActionResultActionFailureReasoninternalError                           MacOsVppAppRevokeLicensesActionResultActionFailureReason = "InternalError"
	MacOsVppAppRevokeLicensesActionResultActionFailureReasonnone                                    MacOsVppAppRevokeLicensesActionResultActionFailureReason = "None"
)

func PossibleValuesForMacOsVppAppRevokeLicensesActionResultActionFailureReason() []string {
	return []string{
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReasonappleFailure),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReasonexpiredVppToken),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReasoninternalError),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReasonnone),
	}
}

func parseMacOsVppAppRevokeLicensesActionResultActionFailureReason(input string) (*MacOsVppAppRevokeLicensesActionResultActionFailureReason, error) {
	vals := map[string]MacOsVppAppRevokeLicensesActionResultActionFailureReason{
		"applefailure": MacOsVppAppRevokeLicensesActionResultActionFailureReasonappleFailure,
		"expiredapplepushnotificationcertificate": MacOsVppAppRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate,
		"expiredvpptoken":                         MacOsVppAppRevokeLicensesActionResultActionFailureReasonexpiredVppToken,
		"internalerror":                           MacOsVppAppRevokeLicensesActionResultActionFailureReasoninternalError,
		"none":                                    MacOsVppAppRevokeLicensesActionResultActionFailureReasonnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppRevokeLicensesActionResultActionFailureReason(input)
	return &out, nil
}
