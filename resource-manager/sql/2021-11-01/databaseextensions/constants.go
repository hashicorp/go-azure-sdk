package databaseextensions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationMode string

const (
	OperationModeExport         OperationMode = "Export"
	OperationModeImport         OperationMode = "Import"
	OperationModePolybaseImport OperationMode = "PolybaseImport"
)

func PossibleValuesForOperationMode() []string {
	return []string{
		string(OperationModeExport),
		string(OperationModeImport),
		string(OperationModePolybaseImport),
	}
}

func (s *OperationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationMode(input string) (*OperationMode, error) {
	vals := map[string]OperationMode{
		"export":         OperationModeExport,
		"import":         OperationModeImport,
		"polybaseimport": OperationModePolybaseImport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationMode(input)
	return &out, nil
}

type StorageKeyType string

const (
	StorageKeyTypeSharedAccessKey  StorageKeyType = "SharedAccessKey"
	StorageKeyTypeStorageAccessKey StorageKeyType = "StorageAccessKey"
)

func PossibleValuesForStorageKeyType() []string {
	return []string{
		string(StorageKeyTypeSharedAccessKey),
		string(StorageKeyTypeStorageAccessKey),
	}
}

func (s *StorageKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageKeyType(input string) (*StorageKeyType, error) {
	vals := map[string]StorageKeyType{
		"sharedaccesskey":  StorageKeyTypeSharedAccessKey,
		"storageaccesskey": StorageKeyTypeStorageAccessKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageKeyType(input)
	return &out, nil
}
