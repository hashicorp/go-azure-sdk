package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyNotebookModelUserRole string

const (
	CopyNotebookModelUserRoleContributor CopyNotebookModelUserRole = "Contributor"
	CopyNotebookModelUserRoleNone        CopyNotebookModelUserRole = "None"
	CopyNotebookModelUserRoleOwner       CopyNotebookModelUserRole = "Owner"
	CopyNotebookModelUserRoleReader      CopyNotebookModelUserRole = "Reader"
)

func PossibleValuesForCopyNotebookModelUserRole() []string {
	return []string{
		string(CopyNotebookModelUserRoleContributor),
		string(CopyNotebookModelUserRoleNone),
		string(CopyNotebookModelUserRoleOwner),
		string(CopyNotebookModelUserRoleReader),
	}
}

func parseCopyNotebookModelUserRole(input string) (*CopyNotebookModelUserRole, error) {
	vals := map[string]CopyNotebookModelUserRole{
		"contributor": CopyNotebookModelUserRoleContributor,
		"none":        CopyNotebookModelUserRoleNone,
		"owner":       CopyNotebookModelUserRoleOwner,
		"reader":      CopyNotebookModelUserRoleReader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CopyNotebookModelUserRole(input)
	return &out, nil
}
