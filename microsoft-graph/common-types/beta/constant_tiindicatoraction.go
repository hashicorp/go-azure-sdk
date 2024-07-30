package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorAction string

const (
	TiIndicatorAction_Alert   TiIndicatorAction = "alert"
	TiIndicatorAction_Allow   TiIndicatorAction = "allow"
	TiIndicatorAction_Block   TiIndicatorAction = "block"
	TiIndicatorAction_Unknown TiIndicatorAction = "unknown"
)

func PossibleValuesForTiIndicatorAction() []string {
	return []string{
		string(TiIndicatorAction_Alert),
		string(TiIndicatorAction_Allow),
		string(TiIndicatorAction_Block),
		string(TiIndicatorAction_Unknown),
	}
}

func (s *TiIndicatorAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTiIndicatorAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTiIndicatorAction(input string) (*TiIndicatorAction, error) {
	vals := map[string]TiIndicatorAction{
		"alert":   TiIndicatorAction_Alert,
		"allow":   TiIndicatorAction_Allow,
		"block":   TiIndicatorAction_Block,
		"unknown": TiIndicatorAction_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorAction(input)
	return &out, nil
}
