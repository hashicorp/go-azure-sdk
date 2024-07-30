package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSnapshotSnapshotType string

const (
	CloudPCSnapshotSnapshotType_Automatic CloudPCSnapshotSnapshotType = "automatic"
	CloudPCSnapshotSnapshotType_Manual    CloudPCSnapshotSnapshotType = "manual"
)

func PossibleValuesForCloudPCSnapshotSnapshotType() []string {
	return []string{
		string(CloudPCSnapshotSnapshotType_Automatic),
		string(CloudPCSnapshotSnapshotType_Manual),
	}
}

func (s *CloudPCSnapshotSnapshotType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSnapshotSnapshotType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSnapshotSnapshotType(input string) (*CloudPCSnapshotSnapshotType, error) {
	vals := map[string]CloudPCSnapshotSnapshotType{
		"automatic": CloudPCSnapshotSnapshotType_Automatic,
		"manual":    CloudPCSnapshotSnapshotType_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSnapshotSnapshotType(input)
	return &out, nil
}
