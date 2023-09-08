package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEventActivityResult string

const (
	CloudPCAuditEventActivityResultclientError CloudPCAuditEventActivityResult = "ClientError"
	CloudPCAuditEventActivityResultfailure     CloudPCAuditEventActivityResult = "Failure"
	CloudPCAuditEventActivityResultother       CloudPCAuditEventActivityResult = "Other"
	CloudPCAuditEventActivityResultsuccess     CloudPCAuditEventActivityResult = "Success"
	CloudPCAuditEventActivityResulttimeout     CloudPCAuditEventActivityResult = "Timeout"
)

func PossibleValuesForCloudPCAuditEventActivityResult() []string {
	return []string{
		string(CloudPCAuditEventActivityResultclientError),
		string(CloudPCAuditEventActivityResultfailure),
		string(CloudPCAuditEventActivityResultother),
		string(CloudPCAuditEventActivityResultsuccess),
		string(CloudPCAuditEventActivityResulttimeout),
	}
}

func parseCloudPCAuditEventActivityResult(input string) (*CloudPCAuditEventActivityResult, error) {
	vals := map[string]CloudPCAuditEventActivityResult{
		"clienterror": CloudPCAuditEventActivityResultclientError,
		"failure":     CloudPCAuditEventActivityResultfailure,
		"other":       CloudPCAuditEventActivityResultother,
		"success":     CloudPCAuditEventActivityResultsuccess,
		"timeout":     CloudPCAuditEventActivityResulttimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditEventActivityResult(input)
	return &out, nil
}
