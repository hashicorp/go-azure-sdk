package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResizeValidationResultValidationResult string

const (
	CloudPCResizeValidationResultValidationResult_CloudPCNotFound          CloudPCResizeValidationResultValidationResult = "cloudPcNotFound"
	CloudPCResizeValidationResultValidationResult_InternalServerError      CloudPCResizeValidationResultValidationResult = "internalServerError"
	CloudPCResizeValidationResultValidationResult_OperationConflict        CloudPCResizeValidationResultValidationResult = "operationConflict"
	CloudPCResizeValidationResultValidationResult_OperationNotSupported    CloudPCResizeValidationResultValidationResult = "operationNotSupported"
	CloudPCResizeValidationResultValidationResult_Success                  CloudPCResizeValidationResultValidationResult = "success"
	CloudPCResizeValidationResultValidationResult_TargetLicenseHasAssigned CloudPCResizeValidationResultValidationResult = "targetLicenseHasAssigned"
)

func PossibleValuesForCloudPCResizeValidationResultValidationResult() []string {
	return []string{
		string(CloudPCResizeValidationResultValidationResult_CloudPCNotFound),
		string(CloudPCResizeValidationResultValidationResult_InternalServerError),
		string(CloudPCResizeValidationResultValidationResult_OperationConflict),
		string(CloudPCResizeValidationResultValidationResult_OperationNotSupported),
		string(CloudPCResizeValidationResultValidationResult_Success),
		string(CloudPCResizeValidationResultValidationResult_TargetLicenseHasAssigned),
	}
}

func (s *CloudPCResizeValidationResultValidationResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCResizeValidationResultValidationResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCResizeValidationResultValidationResult(input string) (*CloudPCResizeValidationResultValidationResult, error) {
	vals := map[string]CloudPCResizeValidationResultValidationResult{
		"cloudpcnotfound":          CloudPCResizeValidationResultValidationResult_CloudPCNotFound,
		"internalservererror":      CloudPCResizeValidationResultValidationResult_InternalServerError,
		"operationconflict":        CloudPCResizeValidationResultValidationResult_OperationConflict,
		"operationnotsupported":    CloudPCResizeValidationResultValidationResult_OperationNotSupported,
		"success":                  CloudPCResizeValidationResultValidationResult_Success,
		"targetlicensehasassigned": CloudPCResizeValidationResultValidationResult_TargetLicenseHasAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCResizeValidationResultValidationResult(input)
	return &out, nil
}
