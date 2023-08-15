package disks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageType string

const (
	StorageTypePremium     StorageType = "Premium"
	StorageTypeStandard    StorageType = "Standard"
	StorageTypeStandardSSD StorageType = "StandardSSD"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypePremium),
		string(StorageTypeStandard),
		string(StorageTypeStandardSSD),
	}
}

func (s *StorageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageType(input string) (*StorageType, error) {
	vals := map[string]StorageType{
		"premium":     StorageTypePremium,
		"standard":    StorageTypeStandard,
		"standardssd": StorageTypeStandardSSD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageType(input)
	return &out, nil
}
