package provideroperationsmetadata

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProviderOperationId{})
}

var _ resourceids.ResourceId = &ProviderOperationId{}

// ProviderOperationId is a struct representing the Resource ID for a Provider Operation
type ProviderOperationId struct {
	ProviderOperationName string
}

// NewProviderOperationID returns a new ProviderOperationId struct
func NewProviderOperationID(providerOperationName string) ProviderOperationId {
	return ProviderOperationId{
		ProviderOperationName: providerOperationName,
	}
}

// ParseProviderOperationID parses 'input' into a ProviderOperationId
func ParseProviderOperationID(input string) (*ProviderOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderOperationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviderOperationIDInsensitively parses 'input' case-insensitively into a ProviderOperationId
// note: this method should only be used for API response data and not user input
func ParseProviderOperationIDInsensitively(input string) (*ProviderOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderOperationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProviderOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ProviderOperationName, ok = input.Parsed["providerOperationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "providerOperationName", input)
	}

	return nil
}

// ValidateProviderOperationID checks that 'input' can be parsed as a Provider Operation ID
func ValidateProviderOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Operation ID
func (id ProviderOperationId) ID() string {
	fmtString := "/providers/Microsoft.Authorization/providerOperations/%s"
	return fmt.Sprintf(fmtString, id.ProviderOperationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Operation ID
func (id ProviderOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticProviderOperations", "providerOperations", "providerOperations"),
		resourceids.UserSpecifiedSegment("providerOperationName", "providerOperationValue"),
	}
}

// String returns a human-readable description of this Provider Operation ID
func (id ProviderOperationId) String() string {
	components := []string{
		fmt.Sprintf("Provider Operation Name: %q", id.ProviderOperationName),
	}
	return fmt.Sprintf("Provider Operation (%s)", strings.Join(components, "\n"))
}
