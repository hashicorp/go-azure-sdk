package recommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Category string

const (
	CategoryCostOptimization Category = "CostOptimization"
	CategoryDemo             Category = "Demo"
	CategoryNewFeature       Category = "NewFeature"
	CategoryOnboarding       Category = "Onboarding"
	CategorySocEfficiency    Category = "SocEfficiency"
)

func PossibleValuesForCategory() []string {
	return []string{
		string(CategoryCostOptimization),
		string(CategoryDemo),
		string(CategoryNewFeature),
		string(CategoryOnboarding),
		string(CategorySocEfficiency),
	}
}

func (s *Category) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategory(input string) (*Category, error) {
	vals := map[string]Category{
		"costoptimization": CategoryCostOptimization,
		"demo":             CategoryDemo,
		"newfeature":       CategoryNewFeature,
		"onboarding":       CategoryOnboarding,
		"socefficiency":    CategorySocEfficiency,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Category(input)
	return &out, nil
}

type Context string

const (
	ContextAnalytics Context = "Analytics"
	ContextIncidents Context = "Incidents"
	ContextNone      Context = "None"
	ContextOverview  Context = "Overview"
)

func PossibleValuesForContext() []string {
	return []string{
		string(ContextAnalytics),
		string(ContextIncidents),
		string(ContextNone),
		string(ContextOverview),
	}
}

func (s *Context) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContext(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContext(input string) (*Context, error) {
	vals := map[string]Context{
		"analytics": ContextAnalytics,
		"incidents": ContextIncidents,
		"none":      ContextNone,
		"overview":  ContextOverview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Context(input)
	return &out, nil
}

type Priority string

const (
	PriorityActive     Priority = "Active"
	PriorityDone       Priority = "Done"
	PriorityHigh       Priority = "High"
	PriorityInProgress Priority = "InProgress"
	PriorityLow        Priority = "Low"
	PriorityMedium     Priority = "Medium"
)

func PossibleValuesForPriority() []string {
	return []string{
		string(PriorityActive),
		string(PriorityDone),
		string(PriorityHigh),
		string(PriorityInProgress),
		string(PriorityLow),
		string(PriorityMedium),
	}
}

func (s *Priority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePriority(input string) (*Priority, error) {
	vals := map[string]Priority{
		"active":     PriorityActive,
		"done":       PriorityDone,
		"high":       PriorityHigh,
		"inprogress": PriorityInProgress,
		"low":        PriorityLow,
		"medium":     PriorityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Priority(input)
	return &out, nil
}

type State string

const (
	StateActive            State = "Active"
	StateCompletedByAction State = "CompletedByAction"
	StateCompletedByUser   State = "CompletedByUser"
	StateDisabled          State = "Disabled"
	StateHidden            State = "Hidden"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateActive),
		string(StateCompletedByAction),
		string(StateCompletedByUser),
		string(StateDisabled),
		string(StateHidden),
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
		"active":            StateActive,
		"completedbyaction": StateCompletedByAction,
		"completedbyuser":   StateCompletedByUser,
		"disabled":          StateDisabled,
		"hidden":            StateHidden,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := State(input)
	return &out, nil
}
