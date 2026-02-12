package attestation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataType string

const (
	DataTypeBinary DataType = "Binary"
	DataTypeJSON   DataType = "JSON"
)

func PossibleValuesForDataType() []string {
	return []string{
		string(DataTypeBinary),
		string(DataTypeJSON),
	}
}

func (s *DataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataType(input string) (*DataType, error) {
	vals := map[string]DataType{
		"binary": DataTypeBinary,
		"json":   DataTypeJSON,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataType(input)
	return &out, nil
}
