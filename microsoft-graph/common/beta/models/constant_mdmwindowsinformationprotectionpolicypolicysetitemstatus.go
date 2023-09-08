package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmWindowsInformationProtectionPolicyPolicySetItemStatus string

const (
	MdmWindowsInformationProtectionPolicyPolicySetItemStatuserror          MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "Error"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatusnotAssigned    MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "NotAssigned"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatuspartialSuccess MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "PartialSuccess"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatussuccess        MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "Success"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatusunknown        MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "Unknown"
	MdmWindowsInformationProtectionPolicyPolicySetItemStatusvalidating     MdmWindowsInformationProtectionPolicyPolicySetItemStatus = "Validating"
)

func PossibleValuesForMdmWindowsInformationProtectionPolicyPolicySetItemStatus() []string {
	return []string{
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatuserror),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatusnotAssigned),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatuspartialSuccess),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatussuccess),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatusunknown),
		string(MdmWindowsInformationProtectionPolicyPolicySetItemStatusvalidating),
	}
}

func parseMdmWindowsInformationProtectionPolicyPolicySetItemStatus(input string) (*MdmWindowsInformationProtectionPolicyPolicySetItemStatus, error) {
	vals := map[string]MdmWindowsInformationProtectionPolicyPolicySetItemStatus{
		"error":          MdmWindowsInformationProtectionPolicyPolicySetItemStatuserror,
		"notassigned":    MdmWindowsInformationProtectionPolicyPolicySetItemStatusnotAssigned,
		"partialsuccess": MdmWindowsInformationProtectionPolicyPolicySetItemStatuspartialSuccess,
		"success":        MdmWindowsInformationProtectionPolicyPolicySetItemStatussuccess,
		"unknown":        MdmWindowsInformationProtectionPolicyPolicySetItemStatusunknown,
		"validating":     MdmWindowsInformationProtectionPolicyPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MdmWindowsInformationProtectionPolicyPolicySetItemStatus(input)
	return &out, nil
}
