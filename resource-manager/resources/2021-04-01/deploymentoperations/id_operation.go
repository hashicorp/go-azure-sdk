package deploymentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = OperationId{}

// OperationId is a struct representing the Resource ID for a Operation
type OperationId struct {
	DeploymentName string
	OperationId    string
}

// NewOperationID returns a new OperationId struct
func NewOperationID(deploymentName string, operationId string) OperationId {
	return OperationId{
		DeploymentName: deploymentName,
		OperationId:    operationId,
	}
}

// ParseOperationID parses 'input' into a OperationId
func ParseOperationID(input string) (*OperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationId{}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseOperationIDInsensitively parses 'input' case-insensitively into a OperationId
// note: this method should only be used for API response data and not user input
func ParseOperationIDInsensitively(input string) (*OperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationId{}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateOperationID checks that 'input' can be parsed as a Operation ID
func ValidateOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Operation ID
func (id OperationId) ID() string {
	fmtString := "/providers/Microsoft.Resources/deployments/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.DeploymentName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Operation ID
func (id OperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResources", "Microsoft.Resources", "Microsoft.Resources"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Operation ID
func (id OperationId) String() string {
	components := []string{
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Operation (%s)", strings.Join(components, "\n"))
}
