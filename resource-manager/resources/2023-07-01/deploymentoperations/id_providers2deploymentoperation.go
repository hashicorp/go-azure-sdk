package deploymentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2DeploymentOperationId{}

// Providers2DeploymentOperationId is a struct representing the Resource ID for a Providers 2 Deployment Operation
type Providers2DeploymentOperationId struct {
	GroupId        string
	DeploymentName string
	OperationId    string
}

// NewProviders2DeploymentOperationID returns a new Providers2DeploymentOperationId struct
func NewProviders2DeploymentOperationID(groupId string, deploymentName string, operationId string) Providers2DeploymentOperationId {
	return Providers2DeploymentOperationId{
		GroupId:        groupId,
		DeploymentName: deploymentName,
		OperationId:    operationId,
	}
}

// ParseProviders2DeploymentOperationID parses 'input' into a Providers2DeploymentOperationId
func ParseProviders2DeploymentOperationID(input string) (*Providers2DeploymentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2DeploymentOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2DeploymentOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseProviders2DeploymentOperationIDInsensitively parses 'input' case-insensitively into a Providers2DeploymentOperationId
// note: this method should only be used for API response data and not user input
func ParseProviders2DeploymentOperationIDInsensitively(input string) (*Providers2DeploymentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2DeploymentOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2DeploymentOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateProviders2DeploymentOperationID checks that 'input' can be parsed as a Providers 2 Deployment Operation ID
func ValidateProviders2DeploymentOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2DeploymentOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Deployment Operation ID
func (id Providers2DeploymentOperationId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Resources/deployments/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DeploymentName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Deployment Operation ID
func (id Providers2DeploymentOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResources", "Microsoft.Resources", "Microsoft.Resources"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Providers 2 Deployment Operation ID
func (id Providers2DeploymentOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Providers 2 Deployment Operation (%s)", strings.Join(components, "\n"))
}
