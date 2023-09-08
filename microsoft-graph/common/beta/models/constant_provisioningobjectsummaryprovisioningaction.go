package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningObjectSummaryProvisioningAction string

const (
	ProvisioningObjectSummaryProvisioningActioncreate       ProvisioningObjectSummaryProvisioningAction = "Create"
	ProvisioningObjectSummaryProvisioningActiondelete       ProvisioningObjectSummaryProvisioningAction = "Delete"
	ProvisioningObjectSummaryProvisioningActiondisable      ProvisioningObjectSummaryProvisioningAction = "Disable"
	ProvisioningObjectSummaryProvisioningActionother        ProvisioningObjectSummaryProvisioningAction = "Other"
	ProvisioningObjectSummaryProvisioningActionstagedDelete ProvisioningObjectSummaryProvisioningAction = "StagedDelete"
	ProvisioningObjectSummaryProvisioningActionupdate       ProvisioningObjectSummaryProvisioningAction = "Update"
)

func PossibleValuesForProvisioningObjectSummaryProvisioningAction() []string {
	return []string{
		string(ProvisioningObjectSummaryProvisioningActioncreate),
		string(ProvisioningObjectSummaryProvisioningActiondelete),
		string(ProvisioningObjectSummaryProvisioningActiondisable),
		string(ProvisioningObjectSummaryProvisioningActionother),
		string(ProvisioningObjectSummaryProvisioningActionstagedDelete),
		string(ProvisioningObjectSummaryProvisioningActionupdate),
	}
}

func parseProvisioningObjectSummaryProvisioningAction(input string) (*ProvisioningObjectSummaryProvisioningAction, error) {
	vals := map[string]ProvisioningObjectSummaryProvisioningAction{
		"create":       ProvisioningObjectSummaryProvisioningActioncreate,
		"delete":       ProvisioningObjectSummaryProvisioningActiondelete,
		"disable":      ProvisioningObjectSummaryProvisioningActiondisable,
		"other":        ProvisioningObjectSummaryProvisioningActionother,
		"stageddelete": ProvisioningObjectSummaryProvisioningActionstagedDelete,
		"update":       ProvisioningObjectSummaryProvisioningActionupdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningObjectSummaryProvisioningAction(input)
	return &out, nil
}
