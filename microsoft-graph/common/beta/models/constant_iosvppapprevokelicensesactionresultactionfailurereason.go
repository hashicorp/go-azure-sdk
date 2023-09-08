package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppRevokeLicensesActionResultActionFailureReason string

const (
	IosVppAppRevokeLicensesActionResultActionFailureReasonappleFailure                            IosVppAppRevokeLicensesActionResultActionFailureReason = "AppleFailure"
	IosVppAppRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate IosVppAppRevokeLicensesActionResultActionFailureReason = "ExpiredApplePushNotificationCertificate"
	IosVppAppRevokeLicensesActionResultActionFailureReasonexpiredVppToken                         IosVppAppRevokeLicensesActionResultActionFailureReason = "ExpiredVppToken"
	IosVppAppRevokeLicensesActionResultActionFailureReasoninternalError                           IosVppAppRevokeLicensesActionResultActionFailureReason = "InternalError"
	IosVppAppRevokeLicensesActionResultActionFailureReasonnone                                    IosVppAppRevokeLicensesActionResultActionFailureReason = "None"
)

func PossibleValuesForIosVppAppRevokeLicensesActionResultActionFailureReason() []string {
	return []string{
		string(IosVppAppRevokeLicensesActionResultActionFailureReasonappleFailure),
		string(IosVppAppRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate),
		string(IosVppAppRevokeLicensesActionResultActionFailureReasonexpiredVppToken),
		string(IosVppAppRevokeLicensesActionResultActionFailureReasoninternalError),
		string(IosVppAppRevokeLicensesActionResultActionFailureReasonnone),
	}
}

func parseIosVppAppRevokeLicensesActionResultActionFailureReason(input string) (*IosVppAppRevokeLicensesActionResultActionFailureReason, error) {
	vals := map[string]IosVppAppRevokeLicensesActionResultActionFailureReason{
		"applefailure": IosVppAppRevokeLicensesActionResultActionFailureReasonappleFailure,
		"expiredapplepushnotificationcertificate": IosVppAppRevokeLicensesActionResultActionFailureReasonexpiredApplePushNotificationCertificate,
		"expiredvpptoken":                         IosVppAppRevokeLicensesActionResultActionFailureReasonexpiredVppToken,
		"internalerror":                           IosVppAppRevokeLicensesActionResultActionFailureReasoninternalError,
		"none":                                    IosVppAppRevokeLicensesActionResultActionFailureReasonnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppRevokeLicensesActionResultActionFailureReason(input)
	return &out, nil
}
