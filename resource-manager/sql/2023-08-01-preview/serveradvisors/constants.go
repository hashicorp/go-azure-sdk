package serveradvisors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorStatus string

const (
	AdvisorStatusGA                   AdvisorStatus = "GA"
	AdvisorStatusLimitedPublicPreview AdvisorStatus = "LimitedPublicPreview"
	AdvisorStatusPrivatePreview       AdvisorStatus = "PrivatePreview"
	AdvisorStatusPublicPreview        AdvisorStatus = "PublicPreview"
)

func PossibleValuesForAdvisorStatus() []string {
	return []string{
		string(AdvisorStatusGA),
		string(AdvisorStatusLimitedPublicPreview),
		string(AdvisorStatusPrivatePreview),
		string(AdvisorStatusPublicPreview),
	}
}

func (s *AdvisorStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvisorStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvisorStatus(input string) (*AdvisorStatus, error) {
	vals := map[string]AdvisorStatus{
		"ga":                   AdvisorStatusGA,
		"limitedpublicpreview": AdvisorStatusLimitedPublicPreview,
		"privatepreview":       AdvisorStatusPrivatePreview,
		"publicpreview":        AdvisorStatusPublicPreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvisorStatus(input)
	return &out, nil
}

type AutoExecuteStatus string

const (
	AutoExecuteStatusDefault  AutoExecuteStatus = "Default"
	AutoExecuteStatusDisabled AutoExecuteStatus = "Disabled"
	AutoExecuteStatusEnabled  AutoExecuteStatus = "Enabled"
)

func PossibleValuesForAutoExecuteStatus() []string {
	return []string{
		string(AutoExecuteStatusDefault),
		string(AutoExecuteStatusDisabled),
		string(AutoExecuteStatusEnabled),
	}
}

func (s *AutoExecuteStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoExecuteStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoExecuteStatus(input string) (*AutoExecuteStatus, error) {
	vals := map[string]AutoExecuteStatus{
		"default":  AutoExecuteStatusDefault,
		"disabled": AutoExecuteStatusDisabled,
		"enabled":  AutoExecuteStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoExecuteStatus(input)
	return &out, nil
}

type AutoExecuteStatusInheritedFrom string

const (
	AutoExecuteStatusInheritedFromDatabase     AutoExecuteStatusInheritedFrom = "Database"
	AutoExecuteStatusInheritedFromDefault      AutoExecuteStatusInheritedFrom = "Default"
	AutoExecuteStatusInheritedFromElasticPool  AutoExecuteStatusInheritedFrom = "ElasticPool"
	AutoExecuteStatusInheritedFromServer       AutoExecuteStatusInheritedFrom = "Server"
	AutoExecuteStatusInheritedFromSubscription AutoExecuteStatusInheritedFrom = "Subscription"
)

func PossibleValuesForAutoExecuteStatusInheritedFrom() []string {
	return []string{
		string(AutoExecuteStatusInheritedFromDatabase),
		string(AutoExecuteStatusInheritedFromDefault),
		string(AutoExecuteStatusInheritedFromElasticPool),
		string(AutoExecuteStatusInheritedFromServer),
		string(AutoExecuteStatusInheritedFromSubscription),
	}
}

func (s *AutoExecuteStatusInheritedFrom) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoExecuteStatusInheritedFrom(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoExecuteStatusInheritedFrom(input string) (*AutoExecuteStatusInheritedFrom, error) {
	vals := map[string]AutoExecuteStatusInheritedFrom{
		"database":     AutoExecuteStatusInheritedFromDatabase,
		"default":      AutoExecuteStatusInheritedFromDefault,
		"elasticpool":  AutoExecuteStatusInheritedFromElasticPool,
		"server":       AutoExecuteStatusInheritedFromServer,
		"subscription": AutoExecuteStatusInheritedFromSubscription,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoExecuteStatusInheritedFrom(input)
	return &out, nil
}

type ImplementationMethod string

const (
	ImplementationMethodAzurePowerShell ImplementationMethod = "AzurePowerShell"
	ImplementationMethodTSql            ImplementationMethod = "TSql"
)

func PossibleValuesForImplementationMethod() []string {
	return []string{
		string(ImplementationMethodAzurePowerShell),
		string(ImplementationMethodTSql),
	}
}

func (s *ImplementationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImplementationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImplementationMethod(input string) (*ImplementationMethod, error) {
	vals := map[string]ImplementationMethod{
		"azurepowershell": ImplementationMethodAzurePowerShell,
		"tsql":            ImplementationMethodTSql,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImplementationMethod(input)
	return &out, nil
}

type IsRetryable string

const (
	IsRetryableNo  IsRetryable = "No"
	IsRetryableYes IsRetryable = "Yes"
)

func PossibleValuesForIsRetryable() []string {
	return []string{
		string(IsRetryableNo),
		string(IsRetryableYes),
	}
}

func (s *IsRetryable) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIsRetryable(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIsRetryable(input string) (*IsRetryable, error) {
	vals := map[string]IsRetryable{
		"no":  IsRetryableNo,
		"yes": IsRetryableYes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IsRetryable(input)
	return &out, nil
}

type RecommendedActionCurrentState string

const (
	RecommendedActionCurrentStateActive          RecommendedActionCurrentState = "Active"
	RecommendedActionCurrentStateError           RecommendedActionCurrentState = "Error"
	RecommendedActionCurrentStateExecuting       RecommendedActionCurrentState = "Executing"
	RecommendedActionCurrentStateExpired         RecommendedActionCurrentState = "Expired"
	RecommendedActionCurrentStateIgnored         RecommendedActionCurrentState = "Ignored"
	RecommendedActionCurrentStateMonitoring      RecommendedActionCurrentState = "Monitoring"
	RecommendedActionCurrentStatePending         RecommendedActionCurrentState = "Pending"
	RecommendedActionCurrentStatePendingRevert   RecommendedActionCurrentState = "PendingRevert"
	RecommendedActionCurrentStateResolved        RecommendedActionCurrentState = "Resolved"
	RecommendedActionCurrentStateRevertCancelled RecommendedActionCurrentState = "RevertCancelled"
	RecommendedActionCurrentStateReverted        RecommendedActionCurrentState = "Reverted"
	RecommendedActionCurrentStateReverting       RecommendedActionCurrentState = "Reverting"
	RecommendedActionCurrentStateSuccess         RecommendedActionCurrentState = "Success"
	RecommendedActionCurrentStateVerifying       RecommendedActionCurrentState = "Verifying"
)

func PossibleValuesForRecommendedActionCurrentState() []string {
	return []string{
		string(RecommendedActionCurrentStateActive),
		string(RecommendedActionCurrentStateError),
		string(RecommendedActionCurrentStateExecuting),
		string(RecommendedActionCurrentStateExpired),
		string(RecommendedActionCurrentStateIgnored),
		string(RecommendedActionCurrentStateMonitoring),
		string(RecommendedActionCurrentStatePending),
		string(RecommendedActionCurrentStatePendingRevert),
		string(RecommendedActionCurrentStateResolved),
		string(RecommendedActionCurrentStateRevertCancelled),
		string(RecommendedActionCurrentStateReverted),
		string(RecommendedActionCurrentStateReverting),
		string(RecommendedActionCurrentStateSuccess),
		string(RecommendedActionCurrentStateVerifying),
	}
}

func (s *RecommendedActionCurrentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendedActionCurrentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendedActionCurrentState(input string) (*RecommendedActionCurrentState, error) {
	vals := map[string]RecommendedActionCurrentState{
		"active":          RecommendedActionCurrentStateActive,
		"error":           RecommendedActionCurrentStateError,
		"executing":       RecommendedActionCurrentStateExecuting,
		"expired":         RecommendedActionCurrentStateExpired,
		"ignored":         RecommendedActionCurrentStateIgnored,
		"monitoring":      RecommendedActionCurrentStateMonitoring,
		"pending":         RecommendedActionCurrentStatePending,
		"pendingrevert":   RecommendedActionCurrentStatePendingRevert,
		"resolved":        RecommendedActionCurrentStateResolved,
		"revertcancelled": RecommendedActionCurrentStateRevertCancelled,
		"reverted":        RecommendedActionCurrentStateReverted,
		"reverting":       RecommendedActionCurrentStateReverting,
		"success":         RecommendedActionCurrentStateSuccess,
		"verifying":       RecommendedActionCurrentStateVerifying,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendedActionCurrentState(input)
	return &out, nil
}

type RecommendedActionInitiatedBy string

const (
	RecommendedActionInitiatedBySystem RecommendedActionInitiatedBy = "System"
	RecommendedActionInitiatedByUser   RecommendedActionInitiatedBy = "User"
)

func PossibleValuesForRecommendedActionInitiatedBy() []string {
	return []string{
		string(RecommendedActionInitiatedBySystem),
		string(RecommendedActionInitiatedByUser),
	}
}

func (s *RecommendedActionInitiatedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendedActionInitiatedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendedActionInitiatedBy(input string) (*RecommendedActionInitiatedBy, error) {
	vals := map[string]RecommendedActionInitiatedBy{
		"system": RecommendedActionInitiatedBySystem,
		"user":   RecommendedActionInitiatedByUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendedActionInitiatedBy(input)
	return &out, nil
}
