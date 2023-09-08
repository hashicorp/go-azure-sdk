package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookUserRole string

const (
	NotebookUserRoleContributor NotebookUserRole = "Contributor"
	NotebookUserRoleNone        NotebookUserRole = "None"
	NotebookUserRoleOwner       NotebookUserRole = "Owner"
	NotebookUserRoleReader      NotebookUserRole = "Reader"
)

func PossibleValuesForNotebookUserRole() []string {
	return []string{
		string(NotebookUserRoleContributor),
		string(NotebookUserRoleNone),
		string(NotebookUserRoleOwner),
		string(NotebookUserRoleReader),
	}
}

func parseNotebookUserRole(input string) (*NotebookUserRole, error) {
	vals := map[string]NotebookUserRole{
		"contributor": NotebookUserRoleContributor,
		"none":        NotebookUserRoleNone,
		"owner":       NotebookUserRoleOwner,
		"reader":      NotebookUserRoleReader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotebookUserRole(input)
	return &out, nil
}
