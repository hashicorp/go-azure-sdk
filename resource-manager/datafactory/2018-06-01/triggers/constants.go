package triggers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobEventTypes string

const (
	BlobEventTypesMicrosoftPointStoragePointBlobCreated BlobEventTypes = "Microsoft.Storage.BlobCreated"
	BlobEventTypesMicrosoftPointStoragePointBlobDeleted BlobEventTypes = "Microsoft.Storage.BlobDeleted"
)

func PossibleValuesForBlobEventTypes() []string {
	return []string{
		string(BlobEventTypesMicrosoftPointStoragePointBlobCreated),
		string(BlobEventTypesMicrosoftPointStoragePointBlobDeleted),
	}
}

func (s *BlobEventTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBlobEventTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBlobEventTypes(input string) (*BlobEventTypes, error) {
	vals := map[string]BlobEventTypes{
		"microsoft.storage.blobcreated": BlobEventTypesMicrosoftPointStoragePointBlobCreated,
		"microsoft.storage.blobdeleted": BlobEventTypesMicrosoftPointStoragePointBlobDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BlobEventTypes(input)
	return &out, nil
}

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
	}
}

func (s *DayOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDayOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDayOfWeek(input string) (*DayOfWeek, error) {
	vals := map[string]DayOfWeek{
		"friday":    DayOfWeekFriday,
		"monday":    DayOfWeekMonday,
		"saturday":  DayOfWeekSaturday,
		"sunday":    DayOfWeekSunday,
		"thursday":  DayOfWeekThursday,
		"tuesday":   DayOfWeekTuesday,
		"wednesday": DayOfWeekWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DayOfWeek(input)
	return &out, nil
}

type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

func PossibleValuesForDaysOfWeek() []string {
	return []string{
		string(DaysOfWeekFriday),
		string(DaysOfWeekMonday),
		string(DaysOfWeekSaturday),
		string(DaysOfWeekSunday),
		string(DaysOfWeekThursday),
		string(DaysOfWeekTuesday),
		string(DaysOfWeekWednesday),
	}
}

func (s *DaysOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDaysOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDaysOfWeek(input string) (*DaysOfWeek, error) {
	vals := map[string]DaysOfWeek{
		"friday":    DaysOfWeekFriday,
		"monday":    DaysOfWeekMonday,
		"saturday":  DaysOfWeekSaturday,
		"sunday":    DaysOfWeekSunday,
		"thursday":  DaysOfWeekThursday,
		"tuesday":   DaysOfWeekTuesday,
		"wednesday": DaysOfWeekWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DaysOfWeek(input)
	return &out, nil
}

type EventSubscriptionStatus string

const (
	EventSubscriptionStatusDeprovisioning EventSubscriptionStatus = "Deprovisioning"
	EventSubscriptionStatusDisabled       EventSubscriptionStatus = "Disabled"
	EventSubscriptionStatusEnabled        EventSubscriptionStatus = "Enabled"
	EventSubscriptionStatusProvisioning   EventSubscriptionStatus = "Provisioning"
	EventSubscriptionStatusUnknown        EventSubscriptionStatus = "Unknown"
)

func PossibleValuesForEventSubscriptionStatus() []string {
	return []string{
		string(EventSubscriptionStatusDeprovisioning),
		string(EventSubscriptionStatusDisabled),
		string(EventSubscriptionStatusEnabled),
		string(EventSubscriptionStatusProvisioning),
		string(EventSubscriptionStatusUnknown),
	}
}

func (s *EventSubscriptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventSubscriptionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventSubscriptionStatus(input string) (*EventSubscriptionStatus, error) {
	vals := map[string]EventSubscriptionStatus{
		"deprovisioning": EventSubscriptionStatusDeprovisioning,
		"disabled":       EventSubscriptionStatusDisabled,
		"enabled":        EventSubscriptionStatusEnabled,
		"provisioning":   EventSubscriptionStatusProvisioning,
		"unknown":        EventSubscriptionStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSubscriptionStatus(input)
	return &out, nil
}

type PipelineReferenceType string

const (
	PipelineReferenceTypePipelineReference PipelineReferenceType = "PipelineReference"
)

func PossibleValuesForPipelineReferenceType() []string {
	return []string{
		string(PipelineReferenceTypePipelineReference),
	}
}

func (s *PipelineReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePipelineReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePipelineReferenceType(input string) (*PipelineReferenceType, error) {
	vals := map[string]PipelineReferenceType{
		"pipelinereference": PipelineReferenceTypePipelineReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PipelineReferenceType(input)
	return &out, nil
}

type RecurrenceFrequency string

const (
	RecurrenceFrequencyDay          RecurrenceFrequency = "Day"
	RecurrenceFrequencyHour         RecurrenceFrequency = "Hour"
	RecurrenceFrequencyMinute       RecurrenceFrequency = "Minute"
	RecurrenceFrequencyMonth        RecurrenceFrequency = "Month"
	RecurrenceFrequencyNotSpecified RecurrenceFrequency = "NotSpecified"
	RecurrenceFrequencyWeek         RecurrenceFrequency = "Week"
	RecurrenceFrequencyYear         RecurrenceFrequency = "Year"
)

func PossibleValuesForRecurrenceFrequency() []string {
	return []string{
		string(RecurrenceFrequencyDay),
		string(RecurrenceFrequencyHour),
		string(RecurrenceFrequencyMinute),
		string(RecurrenceFrequencyMonth),
		string(RecurrenceFrequencyNotSpecified),
		string(RecurrenceFrequencyWeek),
		string(RecurrenceFrequencyYear),
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
		"day":          RecurrenceFrequencyDay,
		"hour":         RecurrenceFrequencyHour,
		"minute":       RecurrenceFrequencyMinute,
		"month":        RecurrenceFrequencyMonth,
		"notspecified": RecurrenceFrequencyNotSpecified,
		"week":         RecurrenceFrequencyWeek,
		"year":         RecurrenceFrequencyYear,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceFrequency(input)
	return &out, nil
}

type TriggerReferenceType string

const (
	TriggerReferenceTypeTriggerReference TriggerReferenceType = "TriggerReference"
)

func PossibleValuesForTriggerReferenceType() []string {
	return []string{
		string(TriggerReferenceTypeTriggerReference),
	}
}

func (s *TriggerReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerReferenceType(input string) (*TriggerReferenceType, error) {
	vals := map[string]TriggerReferenceType{
		"triggerreference": TriggerReferenceTypeTriggerReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerReferenceType(input)
	return &out, nil
}

type TriggerRuntimeState string

const (
	TriggerRuntimeStateDisabled TriggerRuntimeState = "Disabled"
	TriggerRuntimeStateStarted  TriggerRuntimeState = "Started"
	TriggerRuntimeStateStopped  TriggerRuntimeState = "Stopped"
)

func PossibleValuesForTriggerRuntimeState() []string {
	return []string{
		string(TriggerRuntimeStateDisabled),
		string(TriggerRuntimeStateStarted),
		string(TriggerRuntimeStateStopped),
	}
}

func (s *TriggerRuntimeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerRuntimeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerRuntimeState(input string) (*TriggerRuntimeState, error) {
	vals := map[string]TriggerRuntimeState{
		"disabled": TriggerRuntimeStateDisabled,
		"started":  TriggerRuntimeStateStarted,
		"stopped":  TriggerRuntimeStateStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerRuntimeState(input)
	return &out, nil
}

type TumblingWindowFrequency string

const (
	TumblingWindowFrequencyHour   TumblingWindowFrequency = "Hour"
	TumblingWindowFrequencyMinute TumblingWindowFrequency = "Minute"
	TumblingWindowFrequencyMonth  TumblingWindowFrequency = "Month"
)

func PossibleValuesForTumblingWindowFrequency() []string {
	return []string{
		string(TumblingWindowFrequencyHour),
		string(TumblingWindowFrequencyMinute),
		string(TumblingWindowFrequencyMonth),
	}
}

func (s *TumblingWindowFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTumblingWindowFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTumblingWindowFrequency(input string) (*TumblingWindowFrequency, error) {
	vals := map[string]TumblingWindowFrequency{
		"hour":   TumblingWindowFrequencyHour,
		"minute": TumblingWindowFrequencyMinute,
		"month":  TumblingWindowFrequencyMonth,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TumblingWindowFrequency(input)
	return &out, nil
}

type Type string

const (
	TypeLinkedServiceReference Type = "LinkedServiceReference"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeLinkedServiceReference),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"linkedservicereference": TypeLinkedServiceReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
