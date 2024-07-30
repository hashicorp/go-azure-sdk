package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPermissionActionAction string

const (
	AndroidPermissionActionAction_AutoDeny  AndroidPermissionActionAction = "autoDeny"
	AndroidPermissionActionAction_AutoGrant AndroidPermissionActionAction = "autoGrant"
	AndroidPermissionActionAction_Prompt    AndroidPermissionActionAction = "prompt"
)

func PossibleValuesForAndroidPermissionActionAction() []string {
	return []string{
		string(AndroidPermissionActionAction_AutoDeny),
		string(AndroidPermissionActionAction_AutoGrant),
		string(AndroidPermissionActionAction_Prompt),
	}
}

func (s *AndroidPermissionActionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidPermissionActionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidPermissionActionAction(input string) (*AndroidPermissionActionAction, error) {
	vals := map[string]AndroidPermissionActionAction{
		"autodeny":  AndroidPermissionActionAction_AutoDeny,
		"autogrant": AndroidPermissionActionAction_AutoGrant,
		"prompt":    AndroidPermissionActionAction_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPermissionActionAction(input)
	return &out, nil
}
