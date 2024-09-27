package standardassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedStandardAssignmentId{})
}

var _ resourceids.ResourceId = &ScopedStandardAssignmentId{}

// ScopedStandardAssignmentId is a struct representing the Resource ID for a Scoped Standard Assignment
type ScopedStandardAssignmentId struct {
	ResourceId             string
	StandardAssignmentName string
}

// NewScopedStandardAssignmentID returns a new ScopedStandardAssignmentId struct
func NewScopedStandardAssignmentID(resourceId string, standardAssignmentName string) ScopedStandardAssignmentId {
	return ScopedStandardAssignmentId{
		ResourceId:             resourceId,
		StandardAssignmentName: standardAssignmentName,
	}
}

// ParseScopedStandardAssignmentID parses 'input' into a ScopedStandardAssignmentId
func ParseScopedStandardAssignmentID(input string) (*ScopedStandardAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedStandardAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedStandardAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedStandardAssignmentIDInsensitively parses 'input' case-insensitively into a ScopedStandardAssignmentId
// note: this method should only be used for API response data and not user input
func ParseScopedStandardAssignmentIDInsensitively(input string) (*ScopedStandardAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedStandardAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedStandardAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedStandardAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceId, ok = input.Parsed["resourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceId", input)
	}

	if id.StandardAssignmentName, ok = input.Parsed["standardAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "standardAssignmentName", input)
	}

	return nil
}

// ValidateScopedStandardAssignmentID checks that 'input' can be parsed as a Scoped Standard Assignment ID
func ValidateScopedStandardAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedStandardAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Standard Assignment ID
func (id ScopedStandardAssignmentId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/standardAssignments/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceId, "/"), id.StandardAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Standard Assignment ID
func (id ScopedStandardAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceId", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticStandardAssignments", "standardAssignments", "standardAssignments"),
		resourceids.UserSpecifiedSegment("standardAssignmentName", "standardAssignmentName"),
	}
}

// String returns a human-readable description of this Scoped Standard Assignment ID
func (id ScopedStandardAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Resource: %q", id.ResourceId),
		fmt.Sprintf("Standard Assignment Name: %q", id.StandardAssignmentName),
	}
	return fmt.Sprintf("Scoped Standard Assignment (%s)", strings.Join(components, "\n"))
}
