package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResizeValidationResultValidationResult string

const (
	CloudPCResizeValidationResultValidationResultcloudPcNotFound          CloudPCResizeValidationResultValidationResult = "CloudPCNotFound"
	CloudPCResizeValidationResultValidationResultinternalServerError      CloudPCResizeValidationResultValidationResult = "InternalServerError"
	CloudPCResizeValidationResultValidationResultoperationConflict        CloudPCResizeValidationResultValidationResult = "OperationConflict"
	CloudPCResizeValidationResultValidationResultoperationNotSupported    CloudPCResizeValidationResultValidationResult = "OperationNotSupported"
	CloudPCResizeValidationResultValidationResultsuccess                  CloudPCResizeValidationResultValidationResult = "Success"
	CloudPCResizeValidationResultValidationResulttargetLicenseHasAssigned CloudPCResizeValidationResultValidationResult = "TargetLicenseHasAssigned"
)

func PossibleValuesForCloudPCResizeValidationResultValidationResult() []string {
	return []string{
		string(CloudPCResizeValidationResultValidationResultcloudPcNotFound),
		string(CloudPCResizeValidationResultValidationResultinternalServerError),
		string(CloudPCResizeValidationResultValidationResultoperationConflict),
		string(CloudPCResizeValidationResultValidationResultoperationNotSupported),
		string(CloudPCResizeValidationResultValidationResultsuccess),
		string(CloudPCResizeValidationResultValidationResulttargetLicenseHasAssigned),
	}
}

func parseCloudPCResizeValidationResultValidationResult(input string) (*CloudPCResizeValidationResultValidationResult, error) {
	vals := map[string]CloudPCResizeValidationResultValidationResult{
		"cloudpcnotfound":          CloudPCResizeValidationResultValidationResultcloudPcNotFound,
		"internalservererror":      CloudPCResizeValidationResultValidationResultinternalServerError,
		"operationconflict":        CloudPCResizeValidationResultValidationResultoperationConflict,
		"operationnotsupported":    CloudPCResizeValidationResultValidationResultoperationNotSupported,
		"success":                  CloudPCResizeValidationResultValidationResultsuccess,
		"targetlicensehasassigned": CloudPCResizeValidationResultValidationResulttargetLicenseHasAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCResizeValidationResultValidationResult(input)
	return &out, nil
}
