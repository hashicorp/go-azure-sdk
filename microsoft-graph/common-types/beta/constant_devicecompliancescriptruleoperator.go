package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleOperator string

const (
	DeviceComplianceScriptRuleOperator_AllOf            DeviceComplianceScriptRuleOperator = "allOf"
	DeviceComplianceScriptRuleOperator_And              DeviceComplianceScriptRuleOperator = "and"
	DeviceComplianceScriptRuleOperator_BeginsWith       DeviceComplianceScriptRuleOperator = "beginsWith"
	DeviceComplianceScriptRuleOperator_Between          DeviceComplianceScriptRuleOperator = "between"
	DeviceComplianceScriptRuleOperator_Contains         DeviceComplianceScriptRuleOperator = "contains"
	DeviceComplianceScriptRuleOperator_DayTimeBetween   DeviceComplianceScriptRuleOperator = "dayTimeBetween"
	DeviceComplianceScriptRuleOperator_EndsWith         DeviceComplianceScriptRuleOperator = "endsWith"
	DeviceComplianceScriptRuleOperator_ExcludesAll      DeviceComplianceScriptRuleOperator = "excludesAll"
	DeviceComplianceScriptRuleOperator_GreaterEquals    DeviceComplianceScriptRuleOperator = "greaterEquals"
	DeviceComplianceScriptRuleOperator_GreaterThan      DeviceComplianceScriptRuleOperator = "greaterThan"
	DeviceComplianceScriptRuleOperator_IsEquals         DeviceComplianceScriptRuleOperator = "isEquals"
	DeviceComplianceScriptRuleOperator_LessEquals       DeviceComplianceScriptRuleOperator = "lessEquals"
	DeviceComplianceScriptRuleOperator_LessThan         DeviceComplianceScriptRuleOperator = "lessThan"
	DeviceComplianceScriptRuleOperator_None             DeviceComplianceScriptRuleOperator = "none"
	DeviceComplianceScriptRuleOperator_NoneOf           DeviceComplianceScriptRuleOperator = "noneOf"
	DeviceComplianceScriptRuleOperator_NotBeginsWith    DeviceComplianceScriptRuleOperator = "notBeginsWith"
	DeviceComplianceScriptRuleOperator_NotBetween       DeviceComplianceScriptRuleOperator = "notBetween"
	DeviceComplianceScriptRuleOperator_NotContains      DeviceComplianceScriptRuleOperator = "notContains"
	DeviceComplianceScriptRuleOperator_NotEndsWith      DeviceComplianceScriptRuleOperator = "notEndsWith"
	DeviceComplianceScriptRuleOperator_NotEquals        DeviceComplianceScriptRuleOperator = "notEquals"
	DeviceComplianceScriptRuleOperator_OneOf            DeviceComplianceScriptRuleOperator = "oneOf"
	DeviceComplianceScriptRuleOperator_Or               DeviceComplianceScriptRuleOperator = "or"
	DeviceComplianceScriptRuleOperator_OrderedSetEquals DeviceComplianceScriptRuleOperator = "orderedSetEquals"
	DeviceComplianceScriptRuleOperator_SetEquals        DeviceComplianceScriptRuleOperator = "setEquals"
	DeviceComplianceScriptRuleOperator_SubsetOf         DeviceComplianceScriptRuleOperator = "subsetOf"
)

func PossibleValuesForDeviceComplianceScriptRuleOperator() []string {
	return []string{
		string(DeviceComplianceScriptRuleOperator_AllOf),
		string(DeviceComplianceScriptRuleOperator_And),
		string(DeviceComplianceScriptRuleOperator_BeginsWith),
		string(DeviceComplianceScriptRuleOperator_Between),
		string(DeviceComplianceScriptRuleOperator_Contains),
		string(DeviceComplianceScriptRuleOperator_DayTimeBetween),
		string(DeviceComplianceScriptRuleOperator_EndsWith),
		string(DeviceComplianceScriptRuleOperator_ExcludesAll),
		string(DeviceComplianceScriptRuleOperator_GreaterEquals),
		string(DeviceComplianceScriptRuleOperator_GreaterThan),
		string(DeviceComplianceScriptRuleOperator_IsEquals),
		string(DeviceComplianceScriptRuleOperator_LessEquals),
		string(DeviceComplianceScriptRuleOperator_LessThan),
		string(DeviceComplianceScriptRuleOperator_None),
		string(DeviceComplianceScriptRuleOperator_NoneOf),
		string(DeviceComplianceScriptRuleOperator_NotBeginsWith),
		string(DeviceComplianceScriptRuleOperator_NotBetween),
		string(DeviceComplianceScriptRuleOperator_NotContains),
		string(DeviceComplianceScriptRuleOperator_NotEndsWith),
		string(DeviceComplianceScriptRuleOperator_NotEquals),
		string(DeviceComplianceScriptRuleOperator_OneOf),
		string(DeviceComplianceScriptRuleOperator_Or),
		string(DeviceComplianceScriptRuleOperator_OrderedSetEquals),
		string(DeviceComplianceScriptRuleOperator_SetEquals),
		string(DeviceComplianceScriptRuleOperator_SubsetOf),
	}
}

func (s *DeviceComplianceScriptRuleOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRuleOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRuleOperator(input string) (*DeviceComplianceScriptRuleOperator, error) {
	vals := map[string]DeviceComplianceScriptRuleOperator{
		"allof":            DeviceComplianceScriptRuleOperator_AllOf,
		"and":              DeviceComplianceScriptRuleOperator_And,
		"beginswith":       DeviceComplianceScriptRuleOperator_BeginsWith,
		"between":          DeviceComplianceScriptRuleOperator_Between,
		"contains":         DeviceComplianceScriptRuleOperator_Contains,
		"daytimebetween":   DeviceComplianceScriptRuleOperator_DayTimeBetween,
		"endswith":         DeviceComplianceScriptRuleOperator_EndsWith,
		"excludesall":      DeviceComplianceScriptRuleOperator_ExcludesAll,
		"greaterequals":    DeviceComplianceScriptRuleOperator_GreaterEquals,
		"greaterthan":      DeviceComplianceScriptRuleOperator_GreaterThan,
		"isequals":         DeviceComplianceScriptRuleOperator_IsEquals,
		"lessequals":       DeviceComplianceScriptRuleOperator_LessEquals,
		"lessthan":         DeviceComplianceScriptRuleOperator_LessThan,
		"none":             DeviceComplianceScriptRuleOperator_None,
		"noneof":           DeviceComplianceScriptRuleOperator_NoneOf,
		"notbeginswith":    DeviceComplianceScriptRuleOperator_NotBeginsWith,
		"notbetween":       DeviceComplianceScriptRuleOperator_NotBetween,
		"notcontains":      DeviceComplianceScriptRuleOperator_NotContains,
		"notendswith":      DeviceComplianceScriptRuleOperator_NotEndsWith,
		"notequals":        DeviceComplianceScriptRuleOperator_NotEquals,
		"oneof":            DeviceComplianceScriptRuleOperator_OneOf,
		"or":               DeviceComplianceScriptRuleOperator_Or,
		"orderedsetequals": DeviceComplianceScriptRuleOperator_OrderedSetEquals,
		"setequals":        DeviceComplianceScriptRuleOperator_SetEquals,
		"subsetof":         DeviceComplianceScriptRuleOperator_SubsetOf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleOperator(input)
	return &out, nil
}
