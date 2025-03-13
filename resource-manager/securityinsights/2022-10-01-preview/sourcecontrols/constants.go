package sourcecontrols

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentType string

const (
	ContentTypeAnalyticsRule ContentType = "AnalyticsRule"
	ContentTypeWorkbook      ContentType = "Workbook"
)

func PossibleValuesForContentType() []string {
	return []string{
		string(ContentTypeAnalyticsRule),
		string(ContentTypeWorkbook),
	}
}

func (s *ContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentType(input string) (*ContentType, error) {
	vals := map[string]ContentType{
		"analyticsrule": ContentTypeAnalyticsRule,
		"workbook":      ContentTypeWorkbook,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentType(input)
	return &out, nil
}

type DeploymentFetchStatus string

const (
	DeploymentFetchStatusNotFound     DeploymentFetchStatus = "NotFound"
	DeploymentFetchStatusSuccess      DeploymentFetchStatus = "Success"
	DeploymentFetchStatusUnauthorized DeploymentFetchStatus = "Unauthorized"
)

func PossibleValuesForDeploymentFetchStatus() []string {
	return []string{
		string(DeploymentFetchStatusNotFound),
		string(DeploymentFetchStatusSuccess),
		string(DeploymentFetchStatusUnauthorized),
	}
}

func (s *DeploymentFetchStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentFetchStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentFetchStatus(input string) (*DeploymentFetchStatus, error) {
	vals := map[string]DeploymentFetchStatus{
		"notfound":     DeploymentFetchStatusNotFound,
		"success":      DeploymentFetchStatusSuccess,
		"unauthorized": DeploymentFetchStatusUnauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentFetchStatus(input)
	return &out, nil
}

type DeploymentResult string

const (
	DeploymentResultCanceled DeploymentResult = "Canceled"
	DeploymentResultFailed   DeploymentResult = "Failed"
	DeploymentResultSuccess  DeploymentResult = "Success"
)

func PossibleValuesForDeploymentResult() []string {
	return []string{
		string(DeploymentResultCanceled),
		string(DeploymentResultFailed),
		string(DeploymentResultSuccess),
	}
}

func (s *DeploymentResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentResult(input string) (*DeploymentResult, error) {
	vals := map[string]DeploymentResult{
		"canceled": DeploymentResultCanceled,
		"failed":   DeploymentResultFailed,
		"success":  DeploymentResultSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentResult(input)
	return &out, nil
}

type DeploymentState string

const (
	DeploymentStateCanceling  DeploymentState = "Canceling"
	DeploymentStateCompleted  DeploymentState = "Completed"
	DeploymentStateInProgress DeploymentState = "In_Progress"
	DeploymentStateQueued     DeploymentState = "Queued"
)

func PossibleValuesForDeploymentState() []string {
	return []string{
		string(DeploymentStateCanceling),
		string(DeploymentStateCompleted),
		string(DeploymentStateInProgress),
		string(DeploymentStateQueued),
	}
}

func (s *DeploymentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentState(input string) (*DeploymentState, error) {
	vals := map[string]DeploymentState{
		"canceling":   DeploymentStateCanceling,
		"completed":   DeploymentStateCompleted,
		"in_progress": DeploymentStateInProgress,
		"queued":      DeploymentStateQueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentState(input)
	return &out, nil
}

type RepoType string

const (
	RepoTypeDevOps RepoType = "DevOps"
	RepoTypeGithub RepoType = "Github"
)

func PossibleValuesForRepoType() []string {
	return []string{
		string(RepoTypeDevOps),
		string(RepoTypeGithub),
	}
}

func (s *RepoType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRepoType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRepoType(input string) (*RepoType, error) {
	vals := map[string]RepoType{
		"devops": RepoTypeDevOps,
		"github": RepoTypeGithub,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RepoType(input)
	return &out, nil
}

type Version string

const (
	VersionVOne Version = "V1"
	VersionVTwo Version = "V2"
)

func PossibleValuesForVersion() []string {
	return []string{
		string(VersionVOne),
		string(VersionVTwo),
	}
}

func (s *Version) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVersion(input string) (*Version, error) {
	vals := map[string]Version{
		"v1": VersionVOne,
		"v2": VersionVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Version(input)
	return &out, nil
}
