package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEventActivityOperationType string

const (
	CloudPCAuditEventActivityOperationTypecreate CloudPCAuditEventActivityOperationType = "Create"
	CloudPCAuditEventActivityOperationTypedelete CloudPCAuditEventActivityOperationType = "Delete"
	CloudPCAuditEventActivityOperationTypeother  CloudPCAuditEventActivityOperationType = "Other"
	CloudPCAuditEventActivityOperationTypepatch  CloudPCAuditEventActivityOperationType = "Patch"
)

func PossibleValuesForCloudPCAuditEventActivityOperationType() []string {
	return []string{
		string(CloudPCAuditEventActivityOperationTypecreate),
		string(CloudPCAuditEventActivityOperationTypedelete),
		string(CloudPCAuditEventActivityOperationTypeother),
		string(CloudPCAuditEventActivityOperationTypepatch),
	}
}

func parseCloudPCAuditEventActivityOperationType(input string) (*CloudPCAuditEventActivityOperationType, error) {
	vals := map[string]CloudPCAuditEventActivityOperationType{
		"create": CloudPCAuditEventActivityOperationTypecreate,
		"delete": CloudPCAuditEventActivityOperationTypedelete,
		"other":  CloudPCAuditEventActivityOperationTypeother,
		"patch":  CloudPCAuditEventActivityOperationTypepatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditEventActivityOperationType(input)
	return &out, nil
}
