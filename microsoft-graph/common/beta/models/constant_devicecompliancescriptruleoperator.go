package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleOperator string

const (
	DeviceComplianceScriptRuleOperatorallOf            DeviceComplianceScriptRuleOperator = "AllOf"
	DeviceComplianceScriptRuleOperatorand              DeviceComplianceScriptRuleOperator = "And"
	DeviceComplianceScriptRuleOperatorbeginsWith       DeviceComplianceScriptRuleOperator = "BeginsWith"
	DeviceComplianceScriptRuleOperatorbetween          DeviceComplianceScriptRuleOperator = "Between"
	DeviceComplianceScriptRuleOperatorcontains         DeviceComplianceScriptRuleOperator = "Contains"
	DeviceComplianceScriptRuleOperatordayTimeBetween   DeviceComplianceScriptRuleOperator = "DayTimeBetween"
	DeviceComplianceScriptRuleOperatorendsWith         DeviceComplianceScriptRuleOperator = "EndsWith"
	DeviceComplianceScriptRuleOperatorexcludesAll      DeviceComplianceScriptRuleOperator = "ExcludesAll"
	DeviceComplianceScriptRuleOperatorgreaterEquals    DeviceComplianceScriptRuleOperator = "GreaterEquals"
	DeviceComplianceScriptRuleOperatorgreaterThan      DeviceComplianceScriptRuleOperator = "GreaterThan"
	DeviceComplianceScriptRuleOperatorisEquals         DeviceComplianceScriptRuleOperator = "IsEquals"
	DeviceComplianceScriptRuleOperatorlessEquals       DeviceComplianceScriptRuleOperator = "LessEquals"
	DeviceComplianceScriptRuleOperatorlessThan         DeviceComplianceScriptRuleOperator = "LessThan"
	DeviceComplianceScriptRuleOperatornone             DeviceComplianceScriptRuleOperator = "None"
	DeviceComplianceScriptRuleOperatornoneOf           DeviceComplianceScriptRuleOperator = "NoneOf"
	DeviceComplianceScriptRuleOperatornotBeginsWith    DeviceComplianceScriptRuleOperator = "NotBeginsWith"
	DeviceComplianceScriptRuleOperatornotBetween       DeviceComplianceScriptRuleOperator = "NotBetween"
	DeviceComplianceScriptRuleOperatornotContains      DeviceComplianceScriptRuleOperator = "NotContains"
	DeviceComplianceScriptRuleOperatornotEndsWith      DeviceComplianceScriptRuleOperator = "NotEndsWith"
	DeviceComplianceScriptRuleOperatornotEquals        DeviceComplianceScriptRuleOperator = "NotEquals"
	DeviceComplianceScriptRuleOperatoroneOf            DeviceComplianceScriptRuleOperator = "OneOf"
	DeviceComplianceScriptRuleOperatoror               DeviceComplianceScriptRuleOperator = "Or"
	DeviceComplianceScriptRuleOperatororderedSetEquals DeviceComplianceScriptRuleOperator = "OrderedSetEquals"
	DeviceComplianceScriptRuleOperatorsetEquals        DeviceComplianceScriptRuleOperator = "SetEquals"
	DeviceComplianceScriptRuleOperatorsubsetOf         DeviceComplianceScriptRuleOperator = "SubsetOf"
)

func PossibleValuesForDeviceComplianceScriptRuleOperator() []string {
	return []string{
		string(DeviceComplianceScriptRuleOperatorallOf),
		string(DeviceComplianceScriptRuleOperatorand),
		string(DeviceComplianceScriptRuleOperatorbeginsWith),
		string(DeviceComplianceScriptRuleOperatorbetween),
		string(DeviceComplianceScriptRuleOperatorcontains),
		string(DeviceComplianceScriptRuleOperatordayTimeBetween),
		string(DeviceComplianceScriptRuleOperatorendsWith),
		string(DeviceComplianceScriptRuleOperatorexcludesAll),
		string(DeviceComplianceScriptRuleOperatorgreaterEquals),
		string(DeviceComplianceScriptRuleOperatorgreaterThan),
		string(DeviceComplianceScriptRuleOperatorisEquals),
		string(DeviceComplianceScriptRuleOperatorlessEquals),
		string(DeviceComplianceScriptRuleOperatorlessThan),
		string(DeviceComplianceScriptRuleOperatornone),
		string(DeviceComplianceScriptRuleOperatornoneOf),
		string(DeviceComplianceScriptRuleOperatornotBeginsWith),
		string(DeviceComplianceScriptRuleOperatornotBetween),
		string(DeviceComplianceScriptRuleOperatornotContains),
		string(DeviceComplianceScriptRuleOperatornotEndsWith),
		string(DeviceComplianceScriptRuleOperatornotEquals),
		string(DeviceComplianceScriptRuleOperatoroneOf),
		string(DeviceComplianceScriptRuleOperatoror),
		string(DeviceComplianceScriptRuleOperatororderedSetEquals),
		string(DeviceComplianceScriptRuleOperatorsetEquals),
		string(DeviceComplianceScriptRuleOperatorsubsetOf),
	}
}

func parseDeviceComplianceScriptRuleOperator(input string) (*DeviceComplianceScriptRuleOperator, error) {
	vals := map[string]DeviceComplianceScriptRuleOperator{
		"allof":            DeviceComplianceScriptRuleOperatorallOf,
		"and":              DeviceComplianceScriptRuleOperatorand,
		"beginswith":       DeviceComplianceScriptRuleOperatorbeginsWith,
		"between":          DeviceComplianceScriptRuleOperatorbetween,
		"contains":         DeviceComplianceScriptRuleOperatorcontains,
		"daytimebetween":   DeviceComplianceScriptRuleOperatordayTimeBetween,
		"endswith":         DeviceComplianceScriptRuleOperatorendsWith,
		"excludesall":      DeviceComplianceScriptRuleOperatorexcludesAll,
		"greaterequals":    DeviceComplianceScriptRuleOperatorgreaterEquals,
		"greaterthan":      DeviceComplianceScriptRuleOperatorgreaterThan,
		"isequals":         DeviceComplianceScriptRuleOperatorisEquals,
		"lessequals":       DeviceComplianceScriptRuleOperatorlessEquals,
		"lessthan":         DeviceComplianceScriptRuleOperatorlessThan,
		"none":             DeviceComplianceScriptRuleOperatornone,
		"noneof":           DeviceComplianceScriptRuleOperatornoneOf,
		"notbeginswith":    DeviceComplianceScriptRuleOperatornotBeginsWith,
		"notbetween":       DeviceComplianceScriptRuleOperatornotBetween,
		"notcontains":      DeviceComplianceScriptRuleOperatornotContains,
		"notendswith":      DeviceComplianceScriptRuleOperatornotEndsWith,
		"notequals":        DeviceComplianceScriptRuleOperatornotEquals,
		"oneof":            DeviceComplianceScriptRuleOperatoroneOf,
		"or":               DeviceComplianceScriptRuleOperatoror,
		"orderedsetequals": DeviceComplianceScriptRuleOperatororderedSetEquals,
		"setequals":        DeviceComplianceScriptRuleOperatorsetEquals,
		"subsetof":         DeviceComplianceScriptRuleOperatorsubsetOf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleOperator(input)
	return &out, nil
}
