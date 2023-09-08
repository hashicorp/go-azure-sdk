package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersistentBrowserSessionControlMode string

const (
	PersistentBrowserSessionControlModealways PersistentBrowserSessionControlMode = "Always"
	PersistentBrowserSessionControlModenever  PersistentBrowserSessionControlMode = "Never"
)

func PossibleValuesForPersistentBrowserSessionControlMode() []string {
	return []string{
		string(PersistentBrowserSessionControlModealways),
		string(PersistentBrowserSessionControlModenever),
	}
}

func parsePersistentBrowserSessionControlMode(input string) (*PersistentBrowserSessionControlMode, error) {
	vals := map[string]PersistentBrowserSessionControlMode{
		"always": PersistentBrowserSessionControlModealways,
		"never":  PersistentBrowserSessionControlModenever,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersistentBrowserSessionControlMode(input)
	return &out, nil
}
