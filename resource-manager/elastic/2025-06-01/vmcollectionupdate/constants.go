package vmcollectionupdate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationName string

const (
	OperationNameAdd    OperationName = "Add"
	OperationNameDelete OperationName = "Delete"
)

func PossibleValuesForOperationName() []string {
	return []string{
		string(OperationNameAdd),
		string(OperationNameDelete),
	}
}

func (s *OperationName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationName(input string) (*OperationName, error) {
	vals := map[string]OperationName{
		"add":    OperationNameAdd,
		"delete": OperationNameDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationName(input)
	return &out, nil
}
