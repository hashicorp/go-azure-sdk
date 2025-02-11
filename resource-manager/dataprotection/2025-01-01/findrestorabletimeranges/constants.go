package findrestorabletimeranges

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreSourceDataStoreType string

const (
	RestoreSourceDataStoreTypeArchiveStore     RestoreSourceDataStoreType = "ArchiveStore"
	RestoreSourceDataStoreTypeOperationalStore RestoreSourceDataStoreType = "OperationalStore"
	RestoreSourceDataStoreTypeVaultStore       RestoreSourceDataStoreType = "VaultStore"
)

func PossibleValuesForRestoreSourceDataStoreType() []string {
	return []string{
		string(RestoreSourceDataStoreTypeArchiveStore),
		string(RestoreSourceDataStoreTypeOperationalStore),
		string(RestoreSourceDataStoreTypeVaultStore),
	}
}

func (s *RestoreSourceDataStoreType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestoreSourceDataStoreType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestoreSourceDataStoreType(input string) (*RestoreSourceDataStoreType, error) {
	vals := map[string]RestoreSourceDataStoreType{
		"archivestore":     RestoreSourceDataStoreTypeArchiveStore,
		"operationalstore": RestoreSourceDataStoreTypeOperationalStore,
		"vaultstore":       RestoreSourceDataStoreTypeVaultStore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestoreSourceDataStoreType(input)
	return &out, nil
}
