package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunDetailsErrorCode string

const (
	SecurityRunDetailsErrorCodealertCreationFailed      SecurityRunDetailsErrorCode = "AlertCreationFailed"
	SecurityRunDetailsErrorCodealertReportNotFound      SecurityRunDetailsErrorCode = "AlertReportNotFound"
	SecurityRunDetailsErrorCodepartialRowsFailed        SecurityRunDetailsErrorCode = "PartialRowsFailed"
	SecurityRunDetailsErrorCodequeryExceededResultSize  SecurityRunDetailsErrorCode = "QueryExceededResultSize"
	SecurityRunDetailsErrorCodequeryExecutionFailed     SecurityRunDetailsErrorCode = "QueryExecutionFailed"
	SecurityRunDetailsErrorCodequeryExecutionThrottling SecurityRunDetailsErrorCode = "QueryExecutionThrottling"
	SecurityRunDetailsErrorCodequeryLimitsExceeded      SecurityRunDetailsErrorCode = "QueryLimitsExceeded"
	SecurityRunDetailsErrorCodequeryTimeout             SecurityRunDetailsErrorCode = "QueryTimeout"
)

func PossibleValuesForSecurityRunDetailsErrorCode() []string {
	return []string{
		string(SecurityRunDetailsErrorCodealertCreationFailed),
		string(SecurityRunDetailsErrorCodealertReportNotFound),
		string(SecurityRunDetailsErrorCodepartialRowsFailed),
		string(SecurityRunDetailsErrorCodequeryExceededResultSize),
		string(SecurityRunDetailsErrorCodequeryExecutionFailed),
		string(SecurityRunDetailsErrorCodequeryExecutionThrottling),
		string(SecurityRunDetailsErrorCodequeryLimitsExceeded),
		string(SecurityRunDetailsErrorCodequeryTimeout),
	}
}

func parseSecurityRunDetailsErrorCode(input string) (*SecurityRunDetailsErrorCode, error) {
	vals := map[string]SecurityRunDetailsErrorCode{
		"alertcreationfailed":      SecurityRunDetailsErrorCodealertCreationFailed,
		"alertreportnotfound":      SecurityRunDetailsErrorCodealertReportNotFound,
		"partialrowsfailed":        SecurityRunDetailsErrorCodepartialRowsFailed,
		"queryexceededresultsize":  SecurityRunDetailsErrorCodequeryExceededResultSize,
		"queryexecutionfailed":     SecurityRunDetailsErrorCodequeryExecutionFailed,
		"queryexecutionthrottling": SecurityRunDetailsErrorCodequeryExecutionThrottling,
		"querylimitsexceeded":      SecurityRunDetailsErrorCodequeryLimitsExceeded,
		"querytimeout":             SecurityRunDetailsErrorCodequeryTimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRunDetailsErrorCode(input)
	return &out, nil
}
