package watchlists

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceType string

const (
	SourceTypeLocalFile     SourceType = "Local file"
	SourceTypeRemoteStorage SourceType = "Remote storage"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeLocalFile),
		string(SourceTypeRemoteStorage),
	}
}

func parseSourceType(input string) (*SourceType, error) {
	vals := map[string]SourceType{
		"local file":     SourceTypeLocalFile,
		"remote storage": SourceTypeRemoteStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceType(input)
	return &out, nil
}
