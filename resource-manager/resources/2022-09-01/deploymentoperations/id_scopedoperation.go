package deploymentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedOperationId{})
}

var _ resourceids.ResourceId = &ScopedOperationId{}

// ScopedOperationId is a struct representing the Resource ID for a Scoped Operation
type ScopedOperationId struct {
	Scope          string
	DeploymentName string
	OperationId    string
}

// NewScopedOperationID returns a new ScopedOperationId struct
func NewScopedOperationID(scope string, deploymentName string, operationId string) ScopedOperationId {
	return ScopedOperationId{
		Scope:          scope,
		DeploymentName: deploymentName,
		OperationId:    operationId,
	}
}

// ParseScopedOperationID parses 'input' into a ScopedOperationId
func ParseScopedOperationID(input string) (*ScopedOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedOperationIDInsensitively parses 'input' case-insensitively into a ScopedOperationId
// note: this method should only be used for API response data and not user input
func ParseScopedOperationIDInsensitively(input string) (*ScopedOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.DeploymentName, ok = input.Parsed["deploymentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", input)
	}

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateScopedOperationID checks that 'input' can be parsed as a Scoped Operation ID
func ValidateScopedOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Operation ID
func (id ScopedOperationId) ID() string {
	fmtString := "/%s/providers/Microsoft.Resources/deployments/%s/operations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.DeploymentName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Operation ID
func (id ScopedOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResources", "Microsoft.Resources", "Microsoft.Resources"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Scoped Operation ID
func (id ScopedOperationId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Scoped Operation (%s)", strings.Join(components, "\n"))
}
