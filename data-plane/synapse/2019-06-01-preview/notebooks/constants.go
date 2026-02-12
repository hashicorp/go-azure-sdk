package notebooks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BigDataPoolReferenceType string

const (
	BigDataPoolReferenceTypeBigDataPoolReference BigDataPoolReferenceType = "BigDataPoolReference"
)

func PossibleValuesForBigDataPoolReferenceType() []string {
	return []string{
		string(BigDataPoolReferenceTypeBigDataPoolReference),
	}
}

func (s *BigDataPoolReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBigDataPoolReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBigDataPoolReferenceType(input string) (*BigDataPoolReferenceType, error) {
	vals := map[string]BigDataPoolReferenceType{
		"bigdatapoolreference": BigDataPoolReferenceTypeBigDataPoolReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BigDataPoolReferenceType(input)
	return &out, nil
}

type CellOutputType string

const (
	CellOutputTypeDisplayData   CellOutputType = "display_data"
	CellOutputTypeError         CellOutputType = "error"
	CellOutputTypeExecuteResult CellOutputType = "execute_result"
	CellOutputTypeStream        CellOutputType = "stream"
)

func PossibleValuesForCellOutputType() []string {
	return []string{
		string(CellOutputTypeDisplayData),
		string(CellOutputTypeError),
		string(CellOutputTypeExecuteResult),
		string(CellOutputTypeStream),
	}
}

func (s *CellOutputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCellOutputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCellOutputType(input string) (*CellOutputType, error) {
	vals := map[string]CellOutputType{
		"display_data":   CellOutputTypeDisplayData,
		"error":          CellOutputTypeError,
		"execute_result": CellOutputTypeExecuteResult,
		"stream":         CellOutputTypeStream,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CellOutputType(input)
	return &out, nil
}
