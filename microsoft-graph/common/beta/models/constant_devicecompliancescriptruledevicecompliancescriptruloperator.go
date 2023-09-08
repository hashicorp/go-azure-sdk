package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator string

const (
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorallOf            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "AllOf"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorand              DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "And"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorbeginsWith       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "BeginsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorbetween          DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "Between"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorcontains         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "Contains"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatordayTimeBetween   DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "DayTimeBetween"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorendsWith         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "EndsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorexcludesAll      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "ExcludesAll"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorgreaterEquals    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "GreaterEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorgreaterThan      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "GreaterThan"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorisEquals         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "IsEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorlessEquals       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "LessEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorlessThan         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "LessThan"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornone             DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "None"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornoneOf           DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "NoneOf"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotBeginsWith    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "NotBeginsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotBetween       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "NotBetween"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotContains      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "NotContains"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotEndsWith      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "NotEndsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotEquals        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "NotEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatoroneOf            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "OneOf"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatoror               DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "Or"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatororderedSetEquals DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "OrderedSetEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorsetEquals        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "SetEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorsubsetOf         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "SubsetOf"
)

func PossibleValuesForDeviceComplianceScriptRuleDeviceComplianceScriptRulOperator() []string {
	return []string{
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorallOf),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorand),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorbeginsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorbetween),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorcontains),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatordayTimeBetween),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorendsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorexcludesAll),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorgreaterEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorgreaterThan),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorisEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorlessEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorlessThan),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornone),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornoneOf),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotBeginsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotBetween),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotContains),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotEndsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatoroneOf),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatoror),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatororderedSetEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorsetEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorsubsetOf),
	}
}

func parseDeviceComplianceScriptRuleDeviceComplianceScriptRulOperator(input string) (*DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator, error) {
	vals := map[string]DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator{
		"allof":            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorallOf,
		"and":              DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorand,
		"beginswith":       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorbeginsWith,
		"between":          DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorbetween,
		"contains":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorcontains,
		"daytimebetween":   DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatordayTimeBetween,
		"endswith":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorendsWith,
		"excludesall":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorexcludesAll,
		"greaterequals":    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorgreaterEquals,
		"greaterthan":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorgreaterThan,
		"isequals":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorisEquals,
		"lessequals":       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorlessEquals,
		"lessthan":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorlessThan,
		"none":             DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornone,
		"noneof":           DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornoneOf,
		"notbeginswith":    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotBeginsWith,
		"notbetween":       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotBetween,
		"notcontains":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotContains,
		"notendswith":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotEndsWith,
		"notequals":        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatornotEquals,
		"oneof":            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatoroneOf,
		"or":               DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatoror,
		"orderedsetequals": DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatororderedSetEquals,
		"setequals":        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorsetEquals,
		"subsetof":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperatorsubsetOf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator(input)
	return &out, nil
}
