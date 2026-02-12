package triggers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TriggerId{}

// TriggerId is a struct representing the Resource ID for a Trigger
type TriggerId struct {
	BaseURI     string
	TriggerName string
}

// NewTriggerID returns a new TriggerId struct
func NewTriggerID(baseURI string, triggerName string) TriggerId {
	return TriggerId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		TriggerName: triggerName,
	}
}

// ParseTriggerID parses 'input' into a TriggerId
func ParseTriggerID(input string) (*TriggerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTriggerIDInsensitively parses 'input' case-insensitively into a TriggerId
// note: this method should only be used for API response data and not user input
func ParseTriggerIDInsensitively(input string) (*TriggerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TriggerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.TriggerName, ok = input.Parsed["triggerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "triggerName", input)
	}

	return nil
}

// ValidateTriggerID checks that 'input' can be parsed as a Trigger ID
func ValidateTriggerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trigger ID
func (id TriggerId) ID() string {
	fmtString := "%s/triggers/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.TriggerName)
}

// Path returns the formatted Trigger ID without the BaseURI
func (id TriggerId) Path() string {
	fmtString := "/triggers/%s"
	return fmt.Sprintf(fmtString, id.TriggerName)
}

// PathElements returns the values of Trigger ID Segments without the BaseURI
func (id TriggerId) PathElements() []any {
	return []any{id.TriggerName}
}

// Segments returns a slice of Resource ID Segments which comprise this Trigger ID
func (id TriggerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticTriggers", "triggers", "triggers"),
		resourceids.UserSpecifiedSegment("triggerName", "triggerName"),
	}
}

// String returns a human-readable description of this Trigger ID
func (id TriggerId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Trigger Name: %q", id.TriggerName),
	}
	return fmt.Sprintf("Trigger (%s)", strings.Join(components, "\n"))
}
