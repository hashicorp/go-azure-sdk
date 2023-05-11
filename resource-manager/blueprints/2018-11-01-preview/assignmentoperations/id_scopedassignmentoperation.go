package assignmentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedAssignmentOperationId{}

// ScopedAssignmentOperationId is a struct representing the Resource ID for a Scoped Assignment Operation
type ScopedAssignmentOperationId struct {
	ResourceScope           string
	BlueprintAssignmentName string
	AssignmentOperationName string
}

// NewScopedAssignmentOperationID returns a new ScopedAssignmentOperationId struct
func NewScopedAssignmentOperationID(resourceScope string, blueprintAssignmentName string, assignmentOperationName string) ScopedAssignmentOperationId {
	return ScopedAssignmentOperationId{
		ResourceScope:           resourceScope,
		BlueprintAssignmentName: blueprintAssignmentName,
		AssignmentOperationName: assignmentOperationName,
	}
}

// ParseScopedAssignmentOperationID parses 'input' into a ScopedAssignmentOperationId
func ParseScopedAssignmentOperationID(input string) (*ScopedAssignmentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedAssignmentOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedAssignmentOperationId{}

	if id.ResourceScope, ok = parsed.Parsed["resourceScope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceScope", *parsed)
	}

	if id.BlueprintAssignmentName, ok = parsed.Parsed["blueprintAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "blueprintAssignmentName", *parsed)
	}

	if id.AssignmentOperationName, ok = parsed.Parsed["assignmentOperationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assignmentOperationName", *parsed)
	}

	return &id, nil
}

// ParseScopedAssignmentOperationIDInsensitively parses 'input' case-insensitively into a ScopedAssignmentOperationId
// note: this method should only be used for API response data and not user input
func ParseScopedAssignmentOperationIDInsensitively(input string) (*ScopedAssignmentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedAssignmentOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedAssignmentOperationId{}

	if id.ResourceScope, ok = parsed.Parsed["resourceScope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceScope", *parsed)
	}

	if id.BlueprintAssignmentName, ok = parsed.Parsed["blueprintAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "blueprintAssignmentName", *parsed)
	}

	if id.AssignmentOperationName, ok = parsed.Parsed["assignmentOperationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assignmentOperationName", *parsed)
	}

	return &id, nil
}

// ValidateScopedAssignmentOperationID checks that 'input' can be parsed as a Scoped Assignment Operation ID
func ValidateScopedAssignmentOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedAssignmentOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Assignment Operation ID
func (id ScopedAssignmentOperationId) ID() string {
	fmtString := "/%s/providers/Microsoft.Blueprint/blueprintAssignments/%s/assignmentOperations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceScope, "/"), id.BlueprintAssignmentName, id.AssignmentOperationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Assignment Operation ID
func (id ScopedAssignmentOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceScope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBlueprint", "Microsoft.Blueprint", "Microsoft.Blueprint"),
		resourceids.StaticSegment("staticBlueprintAssignments", "blueprintAssignments", "blueprintAssignments"),
		resourceids.UserSpecifiedSegment("blueprintAssignmentName", "blueprintAssignmentValue"),
		resourceids.StaticSegment("staticAssignmentOperations", "assignmentOperations", "assignmentOperations"),
		resourceids.UserSpecifiedSegment("assignmentOperationName", "assignmentOperationValue"),
	}
}

// String returns a human-readable description of this Scoped Assignment Operation ID
func (id ScopedAssignmentOperationId) String() string {
	components := []string{
		fmt.Sprintf("Resource Scope: %q", id.ResourceScope),
		fmt.Sprintf("Blueprint Assignment Name: %q", id.BlueprintAssignmentName),
		fmt.Sprintf("Assignment Operation Name: %q", id.AssignmentOperationName),
	}
	return fmt.Sprintf("Scoped Assignment Operation (%s)", strings.Join(components, "\n"))
}
