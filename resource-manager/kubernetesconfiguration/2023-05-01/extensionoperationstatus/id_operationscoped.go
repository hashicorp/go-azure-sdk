package extensionoperationstatus

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = OperationScopedId{}

// OperationScopedId is a struct representing the Resource ID for a Operation Scoped
type OperationScopedId struct {
	Scope         string
	ExtensionName string
	OperationId   string
}

// NewOperationScopedID returns a new OperationScopedId struct
func NewOperationScopedID(scope string, extensionName string, operationId string) OperationScopedId {
	return OperationScopedId{
		Scope:         scope,
		ExtensionName: extensionName,
		OperationId:   operationId,
	}
}

// ParseOperationScopedID parses 'input' into a OperationScopedId
func ParseOperationScopedID(input string) (*OperationScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationScopedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.ExtensionName, ok = parsed.Parsed["extensionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseOperationScopedIDInsensitively parses 'input' case-insensitively into a OperationScopedId
// note: this method should only be used for API response data and not user input
func ParseOperationScopedIDInsensitively(input string) (*OperationScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationScopedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.ExtensionName, ok = parsed.Parsed["extensionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateOperationScopedID checks that 'input' can be parsed as a Operation Scoped ID
func ValidateOperationScopedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOperationScopedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Operation Scoped ID
func (id OperationScopedId) ID() string {
	fmtString := "/%s/providers/Microsoft.KubernetesConfiguration/extensions/%s/operations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.ExtensionName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Operation Scoped ID
func (id OperationScopedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKubernetesConfiguration", "Microsoft.KubernetesConfiguration", "Microsoft.KubernetesConfiguration"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionName", "extensionValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Operation Scoped ID
func (id OperationScopedId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Extension Name: %q", id.ExtensionName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Operation Scoped (%s)", strings.Join(components, "\n"))
}
