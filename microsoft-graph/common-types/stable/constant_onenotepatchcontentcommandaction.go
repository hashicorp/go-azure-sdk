package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchContentCommandAction string

const (
	OnenotePatchContentCommandAction_Append  OnenotePatchContentCommandAction = "Append"
	OnenotePatchContentCommandAction_Delete  OnenotePatchContentCommandAction = "Delete"
	OnenotePatchContentCommandAction_Insert  OnenotePatchContentCommandAction = "Insert"
	OnenotePatchContentCommandAction_Prepend OnenotePatchContentCommandAction = "Prepend"
	OnenotePatchContentCommandAction_Replace OnenotePatchContentCommandAction = "Replace"
)

func PossibleValuesForOnenotePatchContentCommandAction() []string {
	return []string{
		string(OnenotePatchContentCommandAction_Append),
		string(OnenotePatchContentCommandAction_Delete),
		string(OnenotePatchContentCommandAction_Insert),
		string(OnenotePatchContentCommandAction_Prepend),
		string(OnenotePatchContentCommandAction_Replace),
	}
}

func (s *OnenotePatchContentCommandAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenotePatchContentCommandAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenotePatchContentCommandAction(input string) (*OnenotePatchContentCommandAction, error) {
	vals := map[string]OnenotePatchContentCommandAction{
		"append":  OnenotePatchContentCommandAction_Append,
		"delete":  OnenotePatchContentCommandAction_Delete,
		"insert":  OnenotePatchContentCommandAction_Insert,
		"prepend": OnenotePatchContentCommandAction_Prepend,
		"replace": OnenotePatchContentCommandAction_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenotePatchContentCommandAction(input)
	return &out, nil
}
