package majorversionupgradeprecheck

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MajorVersionUpgradePrecheckStatus string

const (
	MajorVersionUpgradePrecheckStatusCanceled   MajorVersionUpgradePrecheckStatus = "Canceled"
	MajorVersionUpgradePrecheckStatusFailed     MajorVersionUpgradePrecheckStatus = "Failed"
	MajorVersionUpgradePrecheckStatusSucceeded  MajorVersionUpgradePrecheckStatus = "Succeeded"
	MajorVersionUpgradePrecheckStatusValidating MajorVersionUpgradePrecheckStatus = "Validating"
)

func PossibleValuesForMajorVersionUpgradePrecheckStatus() []string {
	return []string{
		string(MajorVersionUpgradePrecheckStatusCanceled),
		string(MajorVersionUpgradePrecheckStatusFailed),
		string(MajorVersionUpgradePrecheckStatusSucceeded),
		string(MajorVersionUpgradePrecheckStatusValidating),
	}
}

func (s *MajorVersionUpgradePrecheckStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMajorVersionUpgradePrecheckStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMajorVersionUpgradePrecheckStatus(input string) (*MajorVersionUpgradePrecheckStatus, error) {
	vals := map[string]MajorVersionUpgradePrecheckStatus{
		"canceled":   MajorVersionUpgradePrecheckStatusCanceled,
		"failed":     MajorVersionUpgradePrecheckStatusFailed,
		"succeeded":  MajorVersionUpgradePrecheckStatusSucceeded,
		"validating": MajorVersionUpgradePrecheckStatusValidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MajorVersionUpgradePrecheckStatus(input)
	return &out, nil
}
