package smartgroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertState string

const (
	AlertStateAcknowledged AlertState = "Acknowledged"
	AlertStateClosed       AlertState = "Closed"
	AlertStateNew          AlertState = "New"
)

func PossibleValuesForAlertState() []string {
	return []string{
		string(AlertStateAcknowledged),
		string(AlertStateClosed),
		string(AlertStateNew),
	}
}

func (s *AlertState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertState(input string) (*AlertState, error) {
	vals := map[string]AlertState{
		"acknowledged": AlertStateAcknowledged,
		"closed":       AlertStateClosed,
		"new":          AlertStateNew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertState(input)
	return &out, nil
}

type MonitorCondition string

const (
	MonitorConditionFired    MonitorCondition = "Fired"
	MonitorConditionResolved MonitorCondition = "Resolved"
)

func PossibleValuesForMonitorCondition() []string {
	return []string{
		string(MonitorConditionFired),
		string(MonitorConditionResolved),
	}
}

func (s *MonitorCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitorCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitorCondition(input string) (*MonitorCondition, error) {
	vals := map[string]MonitorCondition{
		"fired":    MonitorConditionFired,
		"resolved": MonitorConditionResolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitorCondition(input)
	return &out, nil
}

type MonitorService string

const (
	MonitorServiceActivityLogAdministrative MonitorService = "ActivityLog Administrative"
	MonitorServiceActivityLogAutoscale      MonitorService = "ActivityLog Autoscale"
	MonitorServiceActivityLogPolicy         MonitorService = "ActivityLog Policy"
	MonitorServiceActivityLogRecommendation MonitorService = "ActivityLog Recommendation"
	MonitorServiceActivityLogSecurity       MonitorService = "ActivityLog Security"
	MonitorServiceApplicationInsights       MonitorService = "Application Insights"
	MonitorServiceLogAnalytics              MonitorService = "Log Analytics"
	MonitorServiceNagios                    MonitorService = "Nagios"
	MonitorServicePlatform                  MonitorService = "Platform"
	MonitorServiceSCOM                      MonitorService = "SCOM"
	MonitorServiceServiceHealth             MonitorService = "ServiceHealth"
	MonitorServiceSmartDetector             MonitorService = "SmartDetector"
	MonitorServiceVMInsights                MonitorService = "VM Insights"
	MonitorServiceZabbix                    MonitorService = "Zabbix"
)

func PossibleValuesForMonitorService() []string {
	return []string{
		string(MonitorServiceActivityLogAdministrative),
		string(MonitorServiceActivityLogAutoscale),
		string(MonitorServiceActivityLogPolicy),
		string(MonitorServiceActivityLogRecommendation),
		string(MonitorServiceActivityLogSecurity),
		string(MonitorServiceApplicationInsights),
		string(MonitorServiceLogAnalytics),
		string(MonitorServiceNagios),
		string(MonitorServicePlatform),
		string(MonitorServiceSCOM),
		string(MonitorServiceServiceHealth),
		string(MonitorServiceSmartDetector),
		string(MonitorServiceVMInsights),
		string(MonitorServiceZabbix),
	}
}

func (s *MonitorService) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitorService(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitorService(input string) (*MonitorService, error) {
	vals := map[string]MonitorService{
		"activitylog administrative": MonitorServiceActivityLogAdministrative,
		"activitylog autoscale":      MonitorServiceActivityLogAutoscale,
		"activitylog policy":         MonitorServiceActivityLogPolicy,
		"activitylog recommendation": MonitorServiceActivityLogRecommendation,
		"activitylog security":       MonitorServiceActivityLogSecurity,
		"application insights":       MonitorServiceApplicationInsights,
		"log analytics":              MonitorServiceLogAnalytics,
		"nagios":                     MonitorServiceNagios,
		"platform":                   MonitorServicePlatform,
		"scom":                       MonitorServiceSCOM,
		"servicehealth":              MonitorServiceServiceHealth,
		"smartdetector":              MonitorServiceSmartDetector,
		"vm insights":                MonitorServiceVMInsights,
		"zabbix":                     MonitorServiceZabbix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitorService(input)
	return &out, nil
}

type Severity string

const (
	SeveritySevFour  Severity = "Sev4"
	SeveritySevOne   Severity = "Sev1"
	SeveritySevThree Severity = "Sev3"
	SeveritySevTwo   Severity = "Sev2"
	SeveritySevZero  Severity = "Sev0"
)

func PossibleValuesForSeverity() []string {
	return []string{
		string(SeveritySevFour),
		string(SeveritySevOne),
		string(SeveritySevThree),
		string(SeveritySevTwo),
		string(SeveritySevZero),
	}
}

func (s *Severity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSeverity(input string) (*Severity, error) {
	vals := map[string]Severity{
		"sev4": SeveritySevFour,
		"sev1": SeveritySevOne,
		"sev3": SeveritySevThree,
		"sev2": SeveritySevTwo,
		"sev0": SeveritySevZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Severity(input)
	return &out, nil
}

type SmartGroupModificationEvent string

const (
	SmartGroupModificationEventAlertAdded        SmartGroupModificationEvent = "AlertAdded"
	SmartGroupModificationEventAlertRemoved      SmartGroupModificationEvent = "AlertRemoved"
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = "SmartGroupCreated"
	SmartGroupModificationEventStateChange       SmartGroupModificationEvent = "StateChange"
)

func PossibleValuesForSmartGroupModificationEvent() []string {
	return []string{
		string(SmartGroupModificationEventAlertAdded),
		string(SmartGroupModificationEventAlertRemoved),
		string(SmartGroupModificationEventSmartGroupCreated),
		string(SmartGroupModificationEventStateChange),
	}
}

func (s *SmartGroupModificationEvent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmartGroupModificationEvent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmartGroupModificationEvent(input string) (*SmartGroupModificationEvent, error) {
	vals := map[string]SmartGroupModificationEvent{
		"alertadded":        SmartGroupModificationEventAlertAdded,
		"alertremoved":      SmartGroupModificationEventAlertRemoved,
		"smartgroupcreated": SmartGroupModificationEventSmartGroupCreated,
		"statechange":       SmartGroupModificationEventStateChange,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmartGroupModificationEvent(input)
	return &out, nil
}

type SmartGroupsSortByFields string

const (
	SmartGroupsSortByFieldsAlertsCount          SmartGroupsSortByFields = "alertsCount"
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = "lastModifiedDateTime"
	SmartGroupsSortByFieldsSeverity             SmartGroupsSortByFields = "severity"
	SmartGroupsSortByFieldsStartDateTime        SmartGroupsSortByFields = "startDateTime"
	SmartGroupsSortByFieldsState                SmartGroupsSortByFields = "state"
)

func PossibleValuesForSmartGroupsSortByFields() []string {
	return []string{
		string(SmartGroupsSortByFieldsAlertsCount),
		string(SmartGroupsSortByFieldsLastModifiedDateTime),
		string(SmartGroupsSortByFieldsSeverity),
		string(SmartGroupsSortByFieldsStartDateTime),
		string(SmartGroupsSortByFieldsState),
	}
}

func (s *SmartGroupsSortByFields) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmartGroupsSortByFields(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmartGroupsSortByFields(input string) (*SmartGroupsSortByFields, error) {
	vals := map[string]SmartGroupsSortByFields{
		"alertscount":          SmartGroupsSortByFieldsAlertsCount,
		"lastmodifieddatetime": SmartGroupsSortByFieldsLastModifiedDateTime,
		"severity":             SmartGroupsSortByFieldsSeverity,
		"startdatetime":        SmartGroupsSortByFieldsStartDateTime,
		"state":                SmartGroupsSortByFieldsState,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmartGroupsSortByFields(input)
	return &out, nil
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

func PossibleValuesForSortOrder() []string {
	return []string{
		string(SortOrderAsc),
		string(SortOrderDesc),
	}
}

func (s *SortOrder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSortOrder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSortOrder(input string) (*SortOrder, error) {
	vals := map[string]SortOrder{
		"asc":  SortOrderAsc,
		"desc": SortOrderDesc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SortOrder(input)
	return &out, nil
}

type State string

const (
	StateAcknowledged State = "Acknowledged"
	StateClosed       State = "Closed"
	StateNew          State = "New"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateAcknowledged),
		string(StateClosed),
		string(StateNew),
	}
}

func (s *State) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseState(input string) (*State, error) {
	vals := map[string]State{
		"acknowledged": StateAcknowledged,
		"closed":       StateClosed,
		"new":          StateNew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := State(input)
	return &out, nil
}

type TimeRange string

const (
	TimeRangeOned       TimeRange = "1d"
	TimeRangeOneh       TimeRange = "1h"
	TimeRangeSevend     TimeRange = "7d"
	TimeRangeThreeZerod TimeRange = "30d"
)

func PossibleValuesForTimeRange() []string {
	return []string{
		string(TimeRangeOned),
		string(TimeRangeOneh),
		string(TimeRangeSevend),
		string(TimeRangeThreeZerod),
	}
}

func (s *TimeRange) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeRange(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeRange(input string) (*TimeRange, error) {
	vals := map[string]TimeRange{
		"1d":  TimeRangeOned,
		"1h":  TimeRangeOneh,
		"7d":  TimeRangeSevend,
		"30d": TimeRangeThreeZerod,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeRange(input)
	return &out, nil
}
