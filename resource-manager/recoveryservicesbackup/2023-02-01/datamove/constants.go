package datamove

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMoveLevel string

const (
	DataMoveLevelContainer DataMoveLevel = "Container"
	DataMoveLevelInvalid   DataMoveLevel = "Invalid"
	DataMoveLevelVault     DataMoveLevel = "Vault"
)

func PossibleValuesForDataMoveLevel() []string {
	return []string{
		string(DataMoveLevelContainer),
		string(DataMoveLevelInvalid),
		string(DataMoveLevelVault),
	}
}

func parseDataMoveLevel(input string) (*DataMoveLevel, error) {
	vals := map[string]DataMoveLevel{
		"container": DataMoveLevelContainer,
		"invalid":   DataMoveLevelInvalid,
		"vault":     DataMoveLevelVault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataMoveLevel(input)
	return &out, nil
}
