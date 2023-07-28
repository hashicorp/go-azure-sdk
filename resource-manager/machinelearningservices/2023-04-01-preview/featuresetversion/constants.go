package featuresetversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetProvisioningState string

const (
	AssetProvisioningStateCanceled  AssetProvisioningState = "Canceled"
	AssetProvisioningStateCreating  AssetProvisioningState = "Creating"
	AssetProvisioningStateDeleting  AssetProvisioningState = "Deleting"
	AssetProvisioningStateFailed    AssetProvisioningState = "Failed"
	AssetProvisioningStateSucceeded AssetProvisioningState = "Succeeded"
	AssetProvisioningStateUpdating  AssetProvisioningState = "Updating"
)

func PossibleValuesForAssetProvisioningState() []string {
	return []string{
		string(AssetProvisioningStateCanceled),
		string(AssetProvisioningStateCreating),
		string(AssetProvisioningStateDeleting),
		string(AssetProvisioningStateFailed),
		string(AssetProvisioningStateSucceeded),
		string(AssetProvisioningStateUpdating),
	}
}

func (s *AssetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssetProvisioningState(input string) (*AssetProvisioningState, error) {
	vals := map[string]AssetProvisioningState{
		"canceled":  AssetProvisioningStateCanceled,
		"creating":  AssetProvisioningStateCreating,
		"deleting":  AssetProvisioningStateDeleting,
		"failed":    AssetProvisioningStateFailed,
		"succeeded": AssetProvisioningStateSucceeded,
		"updating":  AssetProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssetProvisioningState(input)
	return &out, nil
}

type AutoDeleteCondition string

const (
	AutoDeleteConditionCreatedGreaterThan      AutoDeleteCondition = "CreatedGreaterThan"
	AutoDeleteConditionLastAccessedGreaterThan AutoDeleteCondition = "LastAccessedGreaterThan"
)

func PossibleValuesForAutoDeleteCondition() []string {
	return []string{
		string(AutoDeleteConditionCreatedGreaterThan),
		string(AutoDeleteConditionLastAccessedGreaterThan),
	}
}

func (s *AutoDeleteCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoDeleteCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoDeleteCondition(input string) (*AutoDeleteCondition, error) {
	vals := map[string]AutoDeleteCondition{
		"createdgreaterthan":      AutoDeleteConditionCreatedGreaterThan,
		"lastaccessedgreaterthan": AutoDeleteConditionLastAccessedGreaterThan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoDeleteCondition(input)
	return &out, nil
}

type EmailNotificationEnableType string

const (
	EmailNotificationEnableTypeJobCancelled EmailNotificationEnableType = "JobCancelled"
	EmailNotificationEnableTypeJobCompleted EmailNotificationEnableType = "JobCompleted"
	EmailNotificationEnableTypeJobFailed    EmailNotificationEnableType = "JobFailed"
)

func PossibleValuesForEmailNotificationEnableType() []string {
	return []string{
		string(EmailNotificationEnableTypeJobCancelled),
		string(EmailNotificationEnableTypeJobCompleted),
		string(EmailNotificationEnableTypeJobFailed),
	}
}

func (s *EmailNotificationEnableType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailNotificationEnableType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailNotificationEnableType(input string) (*EmailNotificationEnableType, error) {
	vals := map[string]EmailNotificationEnableType{
		"jobcancelled": EmailNotificationEnableTypeJobCancelled,
		"jobcompleted": EmailNotificationEnableTypeJobCompleted,
		"jobfailed":    EmailNotificationEnableTypeJobFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailNotificationEnableType(input)
	return &out, nil
}

type FeaturestoreJobType string

const (
	FeaturestoreJobTypeBackfillMaterialization  FeaturestoreJobType = "BackfillMaterialization"
	FeaturestoreJobTypeRecurrentMaterialization FeaturestoreJobType = "RecurrentMaterialization"
)

func PossibleValuesForFeaturestoreJobType() []string {
	return []string{
		string(FeaturestoreJobTypeBackfillMaterialization),
		string(FeaturestoreJobTypeRecurrentMaterialization),
	}
}

func (s *FeaturestoreJobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeaturestoreJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeaturestoreJobType(input string) (*FeaturestoreJobType, error) {
	vals := map[string]FeaturestoreJobType{
		"backfillmaterialization":  FeaturestoreJobTypeBackfillMaterialization,
		"recurrentmaterialization": FeaturestoreJobTypeRecurrentMaterialization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeaturestoreJobType(input)
	return &out, nil
}

type JobStatus string

const (
	JobStatusCancelRequested JobStatus = "CancelRequested"
	JobStatusCanceled        JobStatus = "Canceled"
	JobStatusCompleted       JobStatus = "Completed"
	JobStatusFailed          JobStatus = "Failed"
	JobStatusFinalizing      JobStatus = "Finalizing"
	JobStatusNotResponding   JobStatus = "NotResponding"
	JobStatusNotStarted      JobStatus = "NotStarted"
	JobStatusPaused          JobStatus = "Paused"
	JobStatusPreparing       JobStatus = "Preparing"
	JobStatusProvisioning    JobStatus = "Provisioning"
	JobStatusQueued          JobStatus = "Queued"
	JobStatusRunning         JobStatus = "Running"
	JobStatusScheduled       JobStatus = "Scheduled"
	JobStatusStarting        JobStatus = "Starting"
	JobStatusUnknown         JobStatus = "Unknown"
)

func PossibleValuesForJobStatus() []string {
	return []string{
		string(JobStatusCancelRequested),
		string(JobStatusCanceled),
		string(JobStatusCompleted),
		string(JobStatusFailed),
		string(JobStatusFinalizing),
		string(JobStatusNotResponding),
		string(JobStatusNotStarted),
		string(JobStatusPaused),
		string(JobStatusPreparing),
		string(JobStatusProvisioning),
		string(JobStatusQueued),
		string(JobStatusRunning),
		string(JobStatusScheduled),
		string(JobStatusStarting),
		string(JobStatusUnknown),
	}
}

func (s *JobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobStatus(input string) (*JobStatus, error) {
	vals := map[string]JobStatus{
		"cancelrequested": JobStatusCancelRequested,
		"canceled":        JobStatusCanceled,
		"completed":       JobStatusCompleted,
		"failed":          JobStatusFailed,
		"finalizing":      JobStatusFinalizing,
		"notresponding":   JobStatusNotResponding,
		"notstarted":      JobStatusNotStarted,
		"paused":          JobStatusPaused,
		"preparing":       JobStatusPreparing,
		"provisioning":    JobStatusProvisioning,
		"queued":          JobStatusQueued,
		"running":         JobStatusRunning,
		"scheduled":       JobStatusScheduled,
		"starting":        JobStatusStarting,
		"unknown":         JobStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobStatus(input)
	return &out, nil
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

func (s *ListViewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseListViewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
}

type MaterializationStoreType string

const (
	MaterializationStoreTypeNone             MaterializationStoreType = "None"
	MaterializationStoreTypeOffline          MaterializationStoreType = "Offline"
	MaterializationStoreTypeOnline           MaterializationStoreType = "Online"
	MaterializationStoreTypeOnlineAndOffline MaterializationStoreType = "OnlineAndOffline"
)

func PossibleValuesForMaterializationStoreType() []string {
	return []string{
		string(MaterializationStoreTypeNone),
		string(MaterializationStoreTypeOffline),
		string(MaterializationStoreTypeOnline),
		string(MaterializationStoreTypeOnlineAndOffline),
	}
}

func (s *MaterializationStoreType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMaterializationStoreType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMaterializationStoreType(input string) (*MaterializationStoreType, error) {
	vals := map[string]MaterializationStoreType{
		"none":             MaterializationStoreTypeNone,
		"offline":          MaterializationStoreTypeOffline,
		"online":           MaterializationStoreTypeOnline,
		"onlineandoffline": MaterializationStoreTypeOnlineAndOffline,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MaterializationStoreType(input)
	return &out, nil
}

type RecurrenceFrequency string

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = "Day"
	RecurrenceFrequencyHour   RecurrenceFrequency = "Hour"
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	RecurrenceFrequencyMonth  RecurrenceFrequency = "Month"
	RecurrenceFrequencyWeek   RecurrenceFrequency = "Week"
)

func PossibleValuesForRecurrenceFrequency() []string {
	return []string{
		string(RecurrenceFrequencyDay),
		string(RecurrenceFrequencyHour),
		string(RecurrenceFrequencyMinute),
		string(RecurrenceFrequencyMonth),
		string(RecurrenceFrequencyWeek),
	}
}

func (s *RecurrenceFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrenceFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrenceFrequency(input string) (*RecurrenceFrequency, error) {
	vals := map[string]RecurrenceFrequency{
		"day":    RecurrenceFrequencyDay,
		"hour":   RecurrenceFrequencyHour,
		"minute": RecurrenceFrequencyMinute,
		"month":  RecurrenceFrequencyMonth,
		"week":   RecurrenceFrequencyWeek,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceFrequency(input)
	return &out, nil
}

type TriggerType string

const (
	TriggerTypeCron       TriggerType = "Cron"
	TriggerTypeRecurrence TriggerType = "Recurrence"
)

func PossibleValuesForTriggerType() []string {
	return []string{
		string(TriggerTypeCron),
		string(TriggerTypeRecurrence),
	}
}

func (s *TriggerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerType(input string) (*TriggerType, error) {
	vals := map[string]TriggerType{
		"cron":       TriggerTypeCron,
		"recurrence": TriggerTypeRecurrence,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerType(input)
	return &out, nil
}

type WebhookType string

const (
	WebhookTypeAzureDevOps WebhookType = "AzureDevOps"
)

func PossibleValuesForWebhookType() []string {
	return []string{
		string(WebhookTypeAzureDevOps),
	}
}

func (s *WebhookType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebhookType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebhookType(input string) (*WebhookType, error) {
	vals := map[string]WebhookType{
		"azuredevops": WebhookTypeAzureDevOps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebhookType(input)
	return &out, nil
}

type WeekDay string

const (
	WeekDayFriday    WeekDay = "Friday"
	WeekDayMonday    WeekDay = "Monday"
	WeekDaySaturday  WeekDay = "Saturday"
	WeekDaySunday    WeekDay = "Sunday"
	WeekDayThursday  WeekDay = "Thursday"
	WeekDayTuesday   WeekDay = "Tuesday"
	WeekDayWednesday WeekDay = "Wednesday"
)

func PossibleValuesForWeekDay() []string {
	return []string{
		string(WeekDayFriday),
		string(WeekDayMonday),
		string(WeekDaySaturday),
		string(WeekDaySunday),
		string(WeekDayThursday),
		string(WeekDayTuesday),
		string(WeekDayWednesday),
	}
}

func (s *WeekDay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWeekDay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWeekDay(input string) (*WeekDay, error) {
	vals := map[string]WeekDay{
		"friday":    WeekDayFriday,
		"monday":    WeekDayMonday,
		"saturday":  WeekDaySaturday,
		"sunday":    WeekDaySunday,
		"thursday":  WeekDayThursday,
		"tuesday":   WeekDayTuesday,
		"wednesday": WeekDayWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WeekDay(input)
	return &out, nil
}
