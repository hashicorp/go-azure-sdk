package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitlockerRecoveryKeyVolumeType string

const (
	BitlockerRecoveryKeyVolumeTypefixedDataVolume       BitlockerRecoveryKeyVolumeType = "FixedDataVolume"
	BitlockerRecoveryKeyVolumeTypeoperatingSystemVolume BitlockerRecoveryKeyVolumeType = "OperatingSystemVolume"
	BitlockerRecoveryKeyVolumeTyperemovableDataVolume   BitlockerRecoveryKeyVolumeType = "RemovableDataVolume"
)

func PossibleValuesForBitlockerRecoveryKeyVolumeType() []string {
	return []string{
		string(BitlockerRecoveryKeyVolumeTypefixedDataVolume),
		string(BitlockerRecoveryKeyVolumeTypeoperatingSystemVolume),
		string(BitlockerRecoveryKeyVolumeTyperemovableDataVolume),
	}
}

func parseBitlockerRecoveryKeyVolumeType(input string) (*BitlockerRecoveryKeyVolumeType, error) {
	vals := map[string]BitlockerRecoveryKeyVolumeType{
		"fixeddatavolume":       BitlockerRecoveryKeyVolumeTypefixedDataVolume,
		"operatingsystemvolume": BitlockerRecoveryKeyVolumeTypeoperatingSystemVolume,
		"removabledatavolume":   BitlockerRecoveryKeyVolumeTyperemovableDataVolume,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitlockerRecoveryKeyVolumeType(input)
	return &out, nil
}
