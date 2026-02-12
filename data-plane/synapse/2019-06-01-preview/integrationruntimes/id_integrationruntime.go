package integrationruntimes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IntegrationRuntimeId{}

// IntegrationRuntimeId is a struct representing the Resource ID for a Integration Runtime
type IntegrationRuntimeId struct {
	BaseURI                string
	IntegrationRuntimeName string
}

// NewIntegrationRuntimeID returns a new IntegrationRuntimeId struct
func NewIntegrationRuntimeID(baseURI string, integrationRuntimeName string) IntegrationRuntimeId {
	return IntegrationRuntimeId{
		BaseURI:                strings.TrimSuffix(baseURI, "/"),
		IntegrationRuntimeName: integrationRuntimeName,
	}
}

// ParseIntegrationRuntimeID parses 'input' into a IntegrationRuntimeId
func ParseIntegrationRuntimeID(input string) (*IntegrationRuntimeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IntegrationRuntimeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IntegrationRuntimeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIntegrationRuntimeIDInsensitively parses 'input' case-insensitively into a IntegrationRuntimeId
// note: this method should only be used for API response data and not user input
func ParseIntegrationRuntimeIDInsensitively(input string) (*IntegrationRuntimeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IntegrationRuntimeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IntegrationRuntimeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IntegrationRuntimeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.IntegrationRuntimeName, ok = input.Parsed["integrationRuntimeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "integrationRuntimeName", input)
	}

	return nil
}

// ValidateIntegrationRuntimeID checks that 'input' can be parsed as a Integration Runtime ID
func ValidateIntegrationRuntimeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIntegrationRuntimeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Integration Runtime ID
func (id IntegrationRuntimeId) ID() string {
	fmtString := "%s/integrationRuntimes/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.IntegrationRuntimeName)
}

// Path returns the formatted Integration Runtime ID without the BaseURI
func (id IntegrationRuntimeId) Path() string {
	fmtString := "/integrationRuntimes/%s"
	return fmt.Sprintf(fmtString, id.IntegrationRuntimeName)
}

// PathElements returns the values of Integration Runtime ID Segments without the BaseURI
func (id IntegrationRuntimeId) PathElements() []any {
	return []any{id.IntegrationRuntimeName}
}

// Segments returns a slice of Resource ID Segments which comprise this Integration Runtime ID
func (id IntegrationRuntimeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticIntegrationRuntimes", "integrationRuntimes", "integrationRuntimes"),
		resourceids.UserSpecifiedSegment("integrationRuntimeName", "integrationRuntimeName"),
	}
}

// String returns a human-readable description of this Integration Runtime ID
func (id IntegrationRuntimeId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Integration Runtime Name: %q", id.IntegrationRuntimeName),
	}
	return fmt.Sprintf("Integration Runtime (%s)", strings.Join(components, "\n"))
}
