package containerappspatches

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectionStatus string

const (
	DetectionStatusFailed              DetectionStatus = "Failed"
	DetectionStatusRegistryLoginFailed DetectionStatus = "RegistryLoginFailed"
	DetectionStatusSucceeded           DetectionStatus = "Succeeded"
)

func PossibleValuesForDetectionStatus() []string {
	return []string{
		string(DetectionStatusFailed),
		string(DetectionStatusRegistryLoginFailed),
		string(DetectionStatusSucceeded),
	}
}

func (s *DetectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDetectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDetectionStatus(input string) (*DetectionStatus, error) {
	vals := map[string]DetectionStatus{
		"failed":              DetectionStatusFailed,
		"registryloginfailed": DetectionStatusRegistryLoginFailed,
		"succeeded":           DetectionStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectionStatus(input)
	return &out, nil
}

type PatchApplyStatus string

const (
	PatchApplyStatusCanceled               PatchApplyStatus = "Canceled"
	PatchApplyStatusCreatingRevision       PatchApplyStatus = "CreatingRevision"
	PatchApplyStatusImagePushPullFailed    PatchApplyStatus = "ImagePushPullFailed"
	PatchApplyStatusManuallySkipped        PatchApplyStatus = "ManuallySkipped"
	PatchApplyStatusNotStarted             PatchApplyStatus = "NotStarted"
	PatchApplyStatusRebaseFailed           PatchApplyStatus = "RebaseFailed"
	PatchApplyStatusRebaseInProgress       PatchApplyStatus = "RebaseInProgress"
	PatchApplyStatusRevisionCreationFailed PatchApplyStatus = "RevisionCreationFailed"
	PatchApplyStatusSucceeded              PatchApplyStatus = "Succeeded"
)

func PossibleValuesForPatchApplyStatus() []string {
	return []string{
		string(PatchApplyStatusCanceled),
		string(PatchApplyStatusCreatingRevision),
		string(PatchApplyStatusImagePushPullFailed),
		string(PatchApplyStatusManuallySkipped),
		string(PatchApplyStatusNotStarted),
		string(PatchApplyStatusRebaseFailed),
		string(PatchApplyStatusRebaseInProgress),
		string(PatchApplyStatusRevisionCreationFailed),
		string(PatchApplyStatusSucceeded),
	}
}

func (s *PatchApplyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePatchApplyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePatchApplyStatus(input string) (*PatchApplyStatus, error) {
	vals := map[string]PatchApplyStatus{
		"canceled":               PatchApplyStatusCanceled,
		"creatingrevision":       PatchApplyStatusCreatingRevision,
		"imagepushpullfailed":    PatchApplyStatusImagePushPullFailed,
		"manuallyskipped":        PatchApplyStatusManuallySkipped,
		"notstarted":             PatchApplyStatusNotStarted,
		"rebasefailed":           PatchApplyStatusRebaseFailed,
		"rebaseinprogress":       PatchApplyStatusRebaseInProgress,
		"revisioncreationfailed": PatchApplyStatusRevisionCreationFailed,
		"succeeded":              PatchApplyStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PatchApplyStatus(input)
	return &out, nil
}

type PatchType string

const (
	PatchTypeFrameworkAndOSSecurity PatchType = "FrameworkAndOSSecurity"
	PatchTypeFrameworkSecurity      PatchType = "FrameworkSecurity"
	PatchTypeOSSecurity             PatchType = "OSSecurity"
	PatchTypeOther                  PatchType = "Other"
)

func PossibleValuesForPatchType() []string {
	return []string{
		string(PatchTypeFrameworkAndOSSecurity),
		string(PatchTypeFrameworkSecurity),
		string(PatchTypeOSSecurity),
		string(PatchTypeOther),
	}
}

func (s *PatchType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePatchType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePatchType(input string) (*PatchType, error) {
	vals := map[string]PatchType{
		"frameworkandossecurity": PatchTypeFrameworkAndOSSecurity,
		"frameworksecurity":      PatchTypeFrameworkSecurity,
		"ossecurity":             PatchTypeOSSecurity,
		"other":                  PatchTypeOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PatchType(input)
	return &out, nil
}
