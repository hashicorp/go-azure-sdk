package backend

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackendProtocol string

const (
	BackendProtocolHTTP BackendProtocol = "http"
	BackendProtocolSoap BackendProtocol = "soap"
)

func PossibleValuesForBackendProtocol() []string {
	return []string{
		string(BackendProtocolHTTP),
		string(BackendProtocolSoap),
	}
}

func parseBackendProtocol(input string) (*BackendProtocol, error) {
	vals := map[string]BackendProtocol{
		"http": BackendProtocolHTTP,
		"soap": BackendProtocolSoap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackendProtocol(input)
	return &out, nil
}
