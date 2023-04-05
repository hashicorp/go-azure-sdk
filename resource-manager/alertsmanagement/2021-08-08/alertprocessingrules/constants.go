package alertprocessingrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionType string

const (
	ActionTypeAddActionGroups       ActionType = "AddActionGroups"
	ActionTypeRemoveAllActionGroups ActionType = "RemoveAllActionGroups"
)

func PossibleValuesForActionType() []string {
	return []string{
		string(ActionTypeAddActionGroups),
		string(ActionTypeRemoveAllActionGroups),
	}
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

type Field string

const (
	FieldAlertContext        Field = "AlertContext"
	FieldAlertRuleId         Field = "AlertRuleId"
	FieldAlertRuleName       Field = "AlertRuleName"
	FieldDescription         Field = "Description"
	FieldMonitorCondition    Field = "MonitorCondition"
	FieldMonitorService      Field = "MonitorService"
	FieldSeverity            Field = "Severity"
	FieldSignalType          Field = "SignalType"
	FieldTargetResource      Field = "TargetResource"
	FieldTargetResourceGroup Field = "TargetResourceGroup"
	FieldTargetResourceType  Field = "TargetResourceType"
)

func PossibleValuesForField() []string {
	return []string{
		string(FieldAlertContext),
		string(FieldAlertRuleId),
		string(FieldAlertRuleName),
		string(FieldDescription),
		string(FieldMonitorCondition),
		string(FieldMonitorService),
		string(FieldSeverity),
		string(FieldSignalType),
		string(FieldTargetResource),
		string(FieldTargetResourceGroup),
		string(FieldTargetResourceType),
	}
}

type Operator string

const (
	OperatorContains       Operator = "Contains"
	OperatorDoesNotContain Operator = "DoesNotContain"
	OperatorEquals         Operator = "Equals"
	OperatorNotEquals      Operator = "NotEquals"
)

func PossibleValuesForOperator() []string {
	return []string{
		string(OperatorContains),
		string(OperatorDoesNotContain),
		string(OperatorEquals),
		string(OperatorNotEquals),
	}
}

type RecurrenceType string

const (
	RecurrenceTypeDaily   RecurrenceType = "Daily"
	RecurrenceTypeMonthly RecurrenceType = "Monthly"
	RecurrenceTypeWeekly  RecurrenceType = "Weekly"
)

func PossibleValuesForRecurrenceType() []string {
	return []string{
		string(RecurrenceTypeDaily),
		string(RecurrenceTypeMonthly),
		string(RecurrenceTypeWeekly),
	}
}
