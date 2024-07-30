package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitlockerRecoveryKeyVolumeType string

const (
	BitlockerRecoveryKeyVolumeType_FixedDataVolume       BitlockerRecoveryKeyVolumeType = "fixedDataVolume"
	BitlockerRecoveryKeyVolumeType_OperatingSystemVolume BitlockerRecoveryKeyVolumeType = "operatingSystemVolume"
	BitlockerRecoveryKeyVolumeType_RemovableDataVolume   BitlockerRecoveryKeyVolumeType = "removableDataVolume"
)

func PossibleValuesForBitlockerRecoveryKeyVolumeType() []string {
	return []string{
		string(BitlockerRecoveryKeyVolumeType_FixedDataVolume),
		string(BitlockerRecoveryKeyVolumeType_OperatingSystemVolume),
		string(BitlockerRecoveryKeyVolumeType_RemovableDataVolume),
	}
}

func (s *BitlockerRecoveryKeyVolumeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitlockerRecoveryKeyVolumeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitlockerRecoveryKeyVolumeType(input string) (*BitlockerRecoveryKeyVolumeType, error) {
	vals := map[string]BitlockerRecoveryKeyVolumeType{
		"fixeddatavolume":       BitlockerRecoveryKeyVolumeType_FixedDataVolume,
		"operatingsystemvolume": BitlockerRecoveryKeyVolumeType_OperatingSystemVolume,
		"removabledatavolume":   BitlockerRecoveryKeyVolumeType_RemovableDataVolume,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitlockerRecoveryKeyVolumeType(input)
	return &out, nil
}
