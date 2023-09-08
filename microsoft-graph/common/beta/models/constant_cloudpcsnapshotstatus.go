package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSnapshotStatus string

const (
	CloudPCSnapshotStatusready CloudPCSnapshotStatus = "Ready"
)

func PossibleValuesForCloudPCSnapshotStatus() []string {
	return []string{
		string(CloudPCSnapshotStatusready),
	}
}

func parseCloudPCSnapshotStatus(input string) (*CloudPCSnapshotStatus, error) {
	vals := map[string]CloudPCSnapshotStatus{
		"ready": CloudPCSnapshotStatusready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSnapshotStatus(input)
	return &out, nil
}
