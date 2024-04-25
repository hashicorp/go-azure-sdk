package linkers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedDryrunId{})
}

var _ resourceids.ResourceId = &ScopedDryrunId{}

// ScopedDryrunId is a struct representing the Resource ID for a Scoped Dryrun
type ScopedDryrunId struct {
	ResourceUri string
	DryrunName  string
}

// NewScopedDryrunID returns a new ScopedDryrunId struct
func NewScopedDryrunID(resourceUri string, dryrunName string) ScopedDryrunId {
	return ScopedDryrunId{
		ResourceUri: resourceUri,
		DryrunName:  dryrunName,
	}
}

// ParseScopedDryrunID parses 'input' into a ScopedDryrunId
func ParseScopedDryrunID(input string) (*ScopedDryrunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedDryrunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDryrunId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedDryrunIDInsensitively parses 'input' case-insensitively into a ScopedDryrunId
// note: this method should only be used for API response data and not user input
func ParseScopedDryrunIDInsensitively(input string) (*ScopedDryrunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedDryrunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDryrunId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedDryrunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceUri, ok = input.Parsed["resourceUri"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceUri", input)
	}

	if id.DryrunName, ok = input.Parsed["dryrunName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dryrunName", input)
	}

	return nil
}

// ValidateScopedDryrunID checks that 'input' can be parsed as a Scoped Dryrun ID
func ValidateScopedDryrunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedDryrunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Dryrun ID
func (id ScopedDryrunId) ID() string {
	fmtString := "/%s/providers/Microsoft.ServiceLinker/dryruns/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceUri, "/"), id.DryrunName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Dryrun ID
func (id ScopedDryrunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceUri", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceLinker", "Microsoft.ServiceLinker", "Microsoft.ServiceLinker"),
		resourceids.StaticSegment("staticDryruns", "dryruns", "dryruns"),
		resourceids.UserSpecifiedSegment("dryrunName", "dryrunValue"),
	}
}

// String returns a human-readable description of this Scoped Dryrun ID
func (id ScopedDryrunId) String() string {
	components := []string{
		fmt.Sprintf("Resource Uri: %q", id.ResourceUri),
		fmt.Sprintf("Dryrun Name: %q", id.DryrunName),
	}
	return fmt.Sprintf("Scoped Dryrun (%s)", strings.Join(components, "\n"))
}
