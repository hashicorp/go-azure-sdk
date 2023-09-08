package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorAction string

const (
	TiIndicatorActionalert   TiIndicatorAction = "Alert"
	TiIndicatorActionallow   TiIndicatorAction = "Allow"
	TiIndicatorActionblock   TiIndicatorAction = "Block"
	TiIndicatorActionunknown TiIndicatorAction = "Unknown"
)

func PossibleValuesForTiIndicatorAction() []string {
	return []string{
		string(TiIndicatorActionalert),
		string(TiIndicatorActionallow),
		string(TiIndicatorActionblock),
		string(TiIndicatorActionunknown),
	}
}

func parseTiIndicatorAction(input string) (*TiIndicatorAction, error) {
	vals := map[string]TiIndicatorAction{
		"alert":   TiIndicatorActionalert,
		"allow":   TiIndicatorActionallow,
		"block":   TiIndicatorActionblock,
		"unknown": TiIndicatorActionunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorAction(input)
	return &out, nil
}
