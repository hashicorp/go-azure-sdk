package backups

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Origin string

const (
	OriginFull Origin = "Full"
)

func PossibleValuesForOrigin() []string {
	return []string{
		string(OriginFull),
	}
}

func parseOrigin(input string) (*Origin, error) {
	vals := map[string]Origin{
		"full": OriginFull,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Origin(input)
	return &out, nil
}
