package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageStatus string

const (
	LoginPageStatusarchive LoginPageStatus = "Archive"
	LoginPageStatusdelete  LoginPageStatus = "Delete"
	LoginPageStatusdraft   LoginPageStatus = "Draft"
	LoginPageStatusready   LoginPageStatus = "Ready"
	LoginPageStatusunknown LoginPageStatus = "Unknown"
)

func PossibleValuesForLoginPageStatus() []string {
	return []string{
		string(LoginPageStatusarchive),
		string(LoginPageStatusdelete),
		string(LoginPageStatusdraft),
		string(LoginPageStatusready),
		string(LoginPageStatusunknown),
	}
}

func parseLoginPageStatus(input string) (*LoginPageStatus, error) {
	vals := map[string]LoginPageStatus{
		"archive": LoginPageStatusarchive,
		"delete":  LoginPageStatusdelete,
		"draft":   LoginPageStatusdraft,
		"ready":   LoginPageStatusready,
		"unknown": LoginPageStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginPageStatus(input)
	return &out, nil
}
