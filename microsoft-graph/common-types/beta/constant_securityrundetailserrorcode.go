package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunDetailsErrorCode string

const (
	SecurityRunDetailsErrorCode_AlertCreationFailed      SecurityRunDetailsErrorCode = "alertCreationFailed"
	SecurityRunDetailsErrorCode_AlertReportNotFound      SecurityRunDetailsErrorCode = "alertReportNotFound"
	SecurityRunDetailsErrorCode_PartialRowsFailed        SecurityRunDetailsErrorCode = "partialRowsFailed"
	SecurityRunDetailsErrorCode_QueryExceededResultSize  SecurityRunDetailsErrorCode = "queryExceededResultSize"
	SecurityRunDetailsErrorCode_QueryExecutionFailed     SecurityRunDetailsErrorCode = "queryExecutionFailed"
	SecurityRunDetailsErrorCode_QueryExecutionThrottling SecurityRunDetailsErrorCode = "queryExecutionThrottling"
	SecurityRunDetailsErrorCode_QueryLimitsExceeded      SecurityRunDetailsErrorCode = "queryLimitsExceeded"
	SecurityRunDetailsErrorCode_QueryTimeout             SecurityRunDetailsErrorCode = "queryTimeout"
)

func PossibleValuesForSecurityRunDetailsErrorCode() []string {
	return []string{
		string(SecurityRunDetailsErrorCode_AlertCreationFailed),
		string(SecurityRunDetailsErrorCode_AlertReportNotFound),
		string(SecurityRunDetailsErrorCode_PartialRowsFailed),
		string(SecurityRunDetailsErrorCode_QueryExceededResultSize),
		string(SecurityRunDetailsErrorCode_QueryExecutionFailed),
		string(SecurityRunDetailsErrorCode_QueryExecutionThrottling),
		string(SecurityRunDetailsErrorCode_QueryLimitsExceeded),
		string(SecurityRunDetailsErrorCode_QueryTimeout),
	}
}

func (s *SecurityRunDetailsErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRunDetailsErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRunDetailsErrorCode(input string) (*SecurityRunDetailsErrorCode, error) {
	vals := map[string]SecurityRunDetailsErrorCode{
		"alertcreationfailed":      SecurityRunDetailsErrorCode_AlertCreationFailed,
		"alertreportnotfound":      SecurityRunDetailsErrorCode_AlertReportNotFound,
		"partialrowsfailed":        SecurityRunDetailsErrorCode_PartialRowsFailed,
		"queryexceededresultsize":  SecurityRunDetailsErrorCode_QueryExceededResultSize,
		"queryexecutionfailed":     SecurityRunDetailsErrorCode_QueryExecutionFailed,
		"queryexecutionthrottling": SecurityRunDetailsErrorCode_QueryExecutionThrottling,
		"querylimitsexceeded":      SecurityRunDetailsErrorCode_QueryLimitsExceeded,
		"querytimeout":             SecurityRunDetailsErrorCode_QueryTimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRunDetailsErrorCode(input)
	return &out, nil
}
