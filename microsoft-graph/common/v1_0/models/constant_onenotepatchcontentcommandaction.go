package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchContentCommandAction string

const (
	OnenotePatchContentCommandActionAppend  OnenotePatchContentCommandAction = "Append"
	OnenotePatchContentCommandActionDelete  OnenotePatchContentCommandAction = "Delete"
	OnenotePatchContentCommandActionInsert  OnenotePatchContentCommandAction = "Insert"
	OnenotePatchContentCommandActionPrepend OnenotePatchContentCommandAction = "Prepend"
	OnenotePatchContentCommandActionReplace OnenotePatchContentCommandAction = "Replace"
)

func PossibleValuesForOnenotePatchContentCommandAction() []string {
	return []string{
		string(OnenotePatchContentCommandActionAppend),
		string(OnenotePatchContentCommandActionDelete),
		string(OnenotePatchContentCommandActionInsert),
		string(OnenotePatchContentCommandActionPrepend),
		string(OnenotePatchContentCommandActionReplace),
	}
}

func parseOnenotePatchContentCommandAction(input string) (*OnenotePatchContentCommandAction, error) {
	vals := map[string]OnenotePatchContentCommandAction{
		"append":  OnenotePatchContentCommandActionAppend,
		"delete":  OnenotePatchContentCommandActionDelete,
		"insert":  OnenotePatchContentCommandActionInsert,
		"prepend": OnenotePatchContentCommandActionPrepend,
		"replace": OnenotePatchContentCommandActionReplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenotePatchContentCommandAction(input)
	return &out, nil
}
