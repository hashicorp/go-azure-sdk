package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode string

const (
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodedeleted      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "Deleted"
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodenoError      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "NoError"
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodenotFound     MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "NotFound"
	MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodeunauthorized MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForMdmWindowsInformationProtectionPolicyPolicySetItemErrorCode() []string {
	return []string{
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodedeleted),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodenoError),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodenotFound),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodeunauthorized),
	}
}

func parseMdmWindowsInformationProtectionPolicyPolicySetItemErrorCode(input string) (*MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode, error) {
	vals := map[string]MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode{
		"deleted":      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodedeleted,
		"noerror":      MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodenoError,
		"notfound":     MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodenotFound,
		"unauthorized": MdmWindowsInformationProtectionPolicyPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmWindowsInformationProtectionPolicyPolicySetItemErrorCode(input)
	return &out, nil
}
