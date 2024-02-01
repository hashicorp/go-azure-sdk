package sqlpoolstransparentdataencryption

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransparentDataEncryptionStatus string

const (
	TransparentDataEncryptionStatusDisabled TransparentDataEncryptionStatus = "Disabled"
	TransparentDataEncryptionStatusEnabled  TransparentDataEncryptionStatus = "Enabled"
)

func PossibleValuesForTransparentDataEncryptionStatus() []string {
	return []string{
		string(TransparentDataEncryptionStatusDisabled),
		string(TransparentDataEncryptionStatusEnabled),
	}
}

func (s *TransparentDataEncryptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransparentDataEncryptionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransparentDataEncryptionStatus(input string) (*TransparentDataEncryptionStatus, error) {
	vals := map[string]TransparentDataEncryptionStatus{
		"disabled": TransparentDataEncryptionStatusDisabled,
		"enabled":  TransparentDataEncryptionStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransparentDataEncryptionStatus(input)
	return &out, nil
}
