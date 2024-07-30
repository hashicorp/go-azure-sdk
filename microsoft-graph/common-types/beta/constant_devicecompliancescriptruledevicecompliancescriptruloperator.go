package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator string

const (
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_AllOf            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "allOf"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_And              DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "and"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_BeginsWith       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "beginsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Between          DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "between"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Contains         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "contains"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_DayTimeBetween   DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "dayTimeBetween"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_EndsWith         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "endsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_ExcludesAll      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "excludesAll"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_GreaterEquals    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "greaterEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_GreaterThan      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "greaterThan"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_IsEquals         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "isEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_LessEquals       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "lessEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_LessThan         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "lessThan"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_None             DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "none"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NoneOf           DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "noneOf"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotBeginsWith    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "notBeginsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotBetween       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "notBetween"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotContains      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "notContains"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotEndsWith      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "notEndsWith"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotEquals        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "notEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_OneOf            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "oneOf"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Or               DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "or"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_OrderedSetEquals DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "orderedSetEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_SetEquals        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "setEquals"
	DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_SubsetOf         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator = "subsetOf"
)

func PossibleValuesForDeviceComplianceScriptRuleDeviceComplianceScriptRulOperator() []string {
	return []string{
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_AllOf),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_And),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_BeginsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Between),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Contains),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_DayTimeBetween),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_EndsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_ExcludesAll),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_GreaterEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_GreaterThan),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_IsEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_LessEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_LessThan),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_None),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NoneOf),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotBeginsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotBetween),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotContains),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotEndsWith),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_OneOf),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Or),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_OrderedSetEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_SetEquals),
		string(DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_SubsetOf),
	}
}

func (s *DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRuleDeviceComplianceScriptRulOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRuleDeviceComplianceScriptRulOperator(input string) (*DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator, error) {
	vals := map[string]DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator{
		"allof":            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_AllOf,
		"and":              DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_And,
		"beginswith":       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_BeginsWith,
		"between":          DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Between,
		"contains":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Contains,
		"daytimebetween":   DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_DayTimeBetween,
		"endswith":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_EndsWith,
		"excludesall":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_ExcludesAll,
		"greaterequals":    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_GreaterEquals,
		"greaterthan":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_GreaterThan,
		"isequals":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_IsEquals,
		"lessequals":       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_LessEquals,
		"lessthan":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_LessThan,
		"none":             DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_None,
		"noneof":           DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NoneOf,
		"notbeginswith":    DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotBeginsWith,
		"notbetween":       DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotBetween,
		"notcontains":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotContains,
		"notendswith":      DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotEndsWith,
		"notequals":        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_NotEquals,
		"oneof":            DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_OneOf,
		"or":               DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_Or,
		"orderedsetequals": DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_OrderedSetEquals,
		"setequals":        DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_SetEquals,
		"subsetof":         DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator_SubsetOf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator(input)
	return &out, nil
}
