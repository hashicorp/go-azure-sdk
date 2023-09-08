package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSAppleEventReceiverIdentifierType string

const (
	MacOSAppleEventReceiverIdentifierTypebundleID MacOSAppleEventReceiverIdentifierType = "BundleID"
	MacOSAppleEventReceiverIdentifierTypepath     MacOSAppleEventReceiverIdentifierType = "Path"
)

func PossibleValuesForMacOSAppleEventReceiverIdentifierType() []string {
	return []string{
		string(MacOSAppleEventReceiverIdentifierTypebundleID),
		string(MacOSAppleEventReceiverIdentifierTypepath),
	}
}

func parseMacOSAppleEventReceiverIdentifierType(input string) (*MacOSAppleEventReceiverIdentifierType, error) {
	vals := map[string]MacOSAppleEventReceiverIdentifierType{
		"bundleid": MacOSAppleEventReceiverIdentifierTypebundleID,
		"path":     MacOSAppleEventReceiverIdentifierTypepath,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSAppleEventReceiverIdentifierType(input)
	return &out, nil
}
