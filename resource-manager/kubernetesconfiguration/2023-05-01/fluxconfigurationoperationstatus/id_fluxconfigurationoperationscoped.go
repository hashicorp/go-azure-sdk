package fluxconfigurationoperationstatus

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FluxConfigurationOperationScopedId{}

// FluxConfigurationOperationScopedId is a struct representing the Resource ID for a Flux Configuration Operation Scoped
type FluxConfigurationOperationScopedId struct {
	Scope                 string
	FluxConfigurationName string
	OperationId           string
}

// NewFluxConfigurationOperationScopedID returns a new FluxConfigurationOperationScopedId struct
func NewFluxConfigurationOperationScopedID(scope string, fluxConfigurationName string, operationId string) FluxConfigurationOperationScopedId {
	return FluxConfigurationOperationScopedId{
		Scope:                 scope,
		FluxConfigurationName: fluxConfigurationName,
		OperationId:           operationId,
	}
}

// ParseFluxConfigurationOperationScopedID parses 'input' into a FluxConfigurationOperationScopedId
func ParseFluxConfigurationOperationScopedID(input string) (*FluxConfigurationOperationScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(FluxConfigurationOperationScopedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FluxConfigurationOperationScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.FluxConfigurationName, ok = parsed.Parsed["fluxConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fluxConfigurationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseFluxConfigurationOperationScopedIDInsensitively parses 'input' case-insensitively into a FluxConfigurationOperationScopedId
// note: this method should only be used for API response data and not user input
func ParseFluxConfigurationOperationScopedIDInsensitively(input string) (*FluxConfigurationOperationScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(FluxConfigurationOperationScopedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FluxConfigurationOperationScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.FluxConfigurationName, ok = parsed.Parsed["fluxConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fluxConfigurationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateFluxConfigurationOperationScopedID checks that 'input' can be parsed as a Flux Configuration Operation Scoped ID
func ValidateFluxConfigurationOperationScopedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFluxConfigurationOperationScopedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Flux Configuration Operation Scoped ID
func (id FluxConfigurationOperationScopedId) ID() string {
	fmtString := "/%s/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/%s/operations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.FluxConfigurationName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Flux Configuration Operation Scoped ID
func (id FluxConfigurationOperationScopedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKubernetesConfiguration", "Microsoft.KubernetesConfiguration", "Microsoft.KubernetesConfiguration"),
		resourceids.StaticSegment("staticFluxConfigurations", "fluxConfigurations", "fluxConfigurations"),
		resourceids.UserSpecifiedSegment("fluxConfigurationName", "fluxConfigurationValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Flux Configuration Operation Scoped ID
func (id FluxConfigurationOperationScopedId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Flux Configuration Name: %q", id.FluxConfigurationName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Flux Configuration Operation Scoped (%s)", strings.Join(components, "\n"))
}
