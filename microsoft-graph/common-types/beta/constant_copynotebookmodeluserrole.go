package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyNotebookModelUserRole string

const (
	CopyNotebookModelUserRole_Contributor CopyNotebookModelUserRole = "Contributor"
	CopyNotebookModelUserRole_None        CopyNotebookModelUserRole = "None"
	CopyNotebookModelUserRole_Owner       CopyNotebookModelUserRole = "Owner"
	CopyNotebookModelUserRole_Reader      CopyNotebookModelUserRole = "Reader"
)

func PossibleValuesForCopyNotebookModelUserRole() []string {
	return []string{
		string(CopyNotebookModelUserRole_Contributor),
		string(CopyNotebookModelUserRole_None),
		string(CopyNotebookModelUserRole_Owner),
		string(CopyNotebookModelUserRole_Reader),
	}
}

func (s *CopyNotebookModelUserRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCopyNotebookModelUserRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCopyNotebookModelUserRole(input string) (*CopyNotebookModelUserRole, error) {
	vals := map[string]CopyNotebookModelUserRole{
		"contributor": CopyNotebookModelUserRole_Contributor,
		"none":        CopyNotebookModelUserRole_None,
		"owner":       CopyNotebookModelUserRole_Owner,
		"reader":      CopyNotebookModelUserRole_Reader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CopyNotebookModelUserRole(input)
	return &out, nil
}
