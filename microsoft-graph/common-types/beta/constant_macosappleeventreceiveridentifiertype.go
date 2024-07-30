package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSAppleEventReceiverIdentifierType string

const (
	MacOSAppleEventReceiverIdentifierType_BundleID MacOSAppleEventReceiverIdentifierType = "bundleID"
	MacOSAppleEventReceiverIdentifierType_Path     MacOSAppleEventReceiverIdentifierType = "path"
)

func PossibleValuesForMacOSAppleEventReceiverIdentifierType() []string {
	return []string{
		string(MacOSAppleEventReceiverIdentifierType_BundleID),
		string(MacOSAppleEventReceiverIdentifierType_Path),
	}
}

func (s *MacOSAppleEventReceiverIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSAppleEventReceiverIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSAppleEventReceiverIdentifierType(input string) (*MacOSAppleEventReceiverIdentifierType, error) {
	vals := map[string]MacOSAppleEventReceiverIdentifierType{
		"bundleid": MacOSAppleEventReceiverIdentifierType_BundleID,
		"path":     MacOSAppleEventReceiverIdentifierType_Path,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSAppleEventReceiverIdentifierType(input)
	return &out, nil
}
