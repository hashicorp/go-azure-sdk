package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookUserRole string

const (
	NotebookUserRole_Contributor NotebookUserRole = "Contributor"
	NotebookUserRole_None        NotebookUserRole = "None"
	NotebookUserRole_Owner       NotebookUserRole = "Owner"
	NotebookUserRole_Reader      NotebookUserRole = "Reader"
)

func PossibleValuesForNotebookUserRole() []string {
	return []string{
		string(NotebookUserRole_Contributor),
		string(NotebookUserRole_None),
		string(NotebookUserRole_Owner),
		string(NotebookUserRole_Reader),
	}
}

func (s *NotebookUserRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotebookUserRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotebookUserRole(input string) (*NotebookUserRole, error) {
	vals := map[string]NotebookUserRole{
		"contributor": NotebookUserRole_Contributor,
		"none":        NotebookUserRole_None,
		"owner":       NotebookUserRole_Owner,
		"reader":      NotebookUserRole_Reader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotebookUserRole(input)
	return &out, nil
}
