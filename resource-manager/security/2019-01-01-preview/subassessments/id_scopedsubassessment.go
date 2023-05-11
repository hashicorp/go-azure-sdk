package subassessments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedSubAssessmentId{}

// ScopedSubAssessmentId is a struct representing the Resource ID for a Scoped Sub Assessment
type ScopedSubAssessmentId struct {
	Scope             string
	AssessmentName    string
	SubAssessmentName string
}

// NewScopedSubAssessmentID returns a new ScopedSubAssessmentId struct
func NewScopedSubAssessmentID(scope string, assessmentName string, subAssessmentName string) ScopedSubAssessmentId {
	return ScopedSubAssessmentId{
		Scope:             scope,
		AssessmentName:    assessmentName,
		SubAssessmentName: subAssessmentName,
	}
}

// ParseScopedSubAssessmentID parses 'input' into a ScopedSubAssessmentId
func ParseScopedSubAssessmentID(input string) (*ScopedSubAssessmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedSubAssessmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedSubAssessmentId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.AssessmentName, ok = parsed.Parsed["assessmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assessmentName", *parsed)
	}

	if id.SubAssessmentName, ok = parsed.Parsed["subAssessmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subAssessmentName", *parsed)
	}

	return &id, nil
}

// ParseScopedSubAssessmentIDInsensitively parses 'input' case-insensitively into a ScopedSubAssessmentId
// note: this method should only be used for API response data and not user input
func ParseScopedSubAssessmentIDInsensitively(input string) (*ScopedSubAssessmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedSubAssessmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedSubAssessmentId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.AssessmentName, ok = parsed.Parsed["assessmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assessmentName", *parsed)
	}

	if id.SubAssessmentName, ok = parsed.Parsed["subAssessmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subAssessmentName", *parsed)
	}

	return &id, nil
}

// ValidateScopedSubAssessmentID checks that 'input' can be parsed as a Scoped Sub Assessment ID
func ValidateScopedSubAssessmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedSubAssessmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Sub Assessment ID
func (id ScopedSubAssessmentId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/assessments/%s/subAssessments/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.AssessmentName, id.SubAssessmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Sub Assessment ID
func (id ScopedSubAssessmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticAssessments", "assessments", "assessments"),
		resourceids.UserSpecifiedSegment("assessmentName", "assessmentValue"),
		resourceids.StaticSegment("staticSubAssessments", "subAssessments", "subAssessments"),
		resourceids.UserSpecifiedSegment("subAssessmentName", "subAssessmentValue"),
	}
}

// String returns a human-readable description of this Scoped Sub Assessment ID
func (id ScopedSubAssessmentId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Assessment Name: %q", id.AssessmentName),
		fmt.Sprintf("Sub Assessment Name: %q", id.SubAssessmentName),
	}
	return fmt.Sprintf("Scoped Sub Assessment (%s)", strings.Join(components, "\n"))
}
