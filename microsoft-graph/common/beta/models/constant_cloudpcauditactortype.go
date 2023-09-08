package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditActorType string

const (
	CloudPCAuditActorTypeapplication CloudPCAuditActorType = "Application"
	CloudPCAuditActorTypeitPro       CloudPCAuditActorType = "ItPro"
	CloudPCAuditActorTypepartner     CloudPCAuditActorType = "Partner"
	CloudPCAuditActorTypeunknown     CloudPCAuditActorType = "Unknown"
)

func PossibleValuesForCloudPCAuditActorType() []string {
	return []string{
		string(CloudPCAuditActorTypeapplication),
		string(CloudPCAuditActorTypeitPro),
		string(CloudPCAuditActorTypepartner),
		string(CloudPCAuditActorTypeunknown),
	}
}

func parseCloudPCAuditActorType(input string) (*CloudPCAuditActorType, error) {
	vals := map[string]CloudPCAuditActorType{
		"application": CloudPCAuditActorTypeapplication,
		"itpro":       CloudPCAuditActorTypeitPro,
		"partner":     CloudPCAuditActorTypepartner,
		"unknown":     CloudPCAuditActorTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditActorType(input)
	return &out, nil
}
