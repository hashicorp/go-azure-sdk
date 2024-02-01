package sqlpools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMode string

const (
	CreateModeDefault            CreateMode = "Default"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeRecovery           CreateMode = "Recovery"
	CreateModeRestore            CreateMode = "Restore"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeDefault),
		string(CreateModePointInTimeRestore),
		string(CreateModeRecovery),
		string(CreateModeRestore),
	}
}

func (s *CreateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMode(input string) (*CreateMode, error) {
	vals := map[string]CreateMode{
		"default":            CreateModeDefault,
		"pointintimerestore": CreateModePointInTimeRestore,
		"recovery":           CreateModeRecovery,
		"restore":            CreateModeRestore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMode(input)
	return &out, nil
}

type StorageAccountType string

const (
	StorageAccountTypeGRS StorageAccountType = "GRS"
	StorageAccountTypeLRS StorageAccountType = "LRS"
)

func PossibleValuesForStorageAccountType() []string {
	return []string{
		string(StorageAccountTypeGRS),
		string(StorageAccountTypeLRS),
	}
}

func (s *StorageAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageAccountType(input string) (*StorageAccountType, error) {
	vals := map[string]StorageAccountType{
		"grs": StorageAccountTypeGRS,
		"lrs": StorageAccountTypeLRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAccountType(input)
	return &out, nil
}
