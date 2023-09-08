package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSnapshotSnapshotType string

const (
	CloudPCSnapshotSnapshotTypeautomatic CloudPCSnapshotSnapshotType = "Automatic"
	CloudPCSnapshotSnapshotTypemanual    CloudPCSnapshotSnapshotType = "Manual"
)

func PossibleValuesForCloudPCSnapshotSnapshotType() []string {
	return []string{
		string(CloudPCSnapshotSnapshotTypeautomatic),
		string(CloudPCSnapshotSnapshotTypemanual),
	}
}

func parseCloudPCSnapshotSnapshotType(input string) (*CloudPCSnapshotSnapshotType, error) {
	vals := map[string]CloudPCSnapshotSnapshotType{
		"automatic": CloudPCSnapshotSnapshotTypeautomatic,
		"manual":    CloudPCSnapshotSnapshotTypemanual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSnapshotSnapshotType(input)
	return &out, nil
}
