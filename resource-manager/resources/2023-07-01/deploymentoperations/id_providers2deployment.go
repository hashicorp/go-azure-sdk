package deploymentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2DeploymentId{}

// Providers2DeploymentId is a struct representing the Resource ID for a Providers 2 Deployment
type Providers2DeploymentId struct {
	GroupId        string
	DeploymentName string
}

// NewProviders2DeploymentID returns a new Providers2DeploymentId struct
func NewProviders2DeploymentID(groupId string, deploymentName string) Providers2DeploymentId {
	return Providers2DeploymentId{
		GroupId:        groupId,
		DeploymentName: deploymentName,
	}
}

// ParseProviders2DeploymentID parses 'input' into a Providers2DeploymentId
func ParseProviders2DeploymentID(input string) (*Providers2DeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2DeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2DeploymentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", *parsed)
	}

	return &id, nil
}

// ParseProviders2DeploymentIDInsensitively parses 'input' case-insensitively into a Providers2DeploymentId
// note: this method should only be used for API response data and not user input
func ParseProviders2DeploymentIDInsensitively(input string) (*Providers2DeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2DeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2DeploymentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", *parsed)
	}

	return &id, nil
}

// ValidateProviders2DeploymentID checks that 'input' can be parsed as a Providers 2 Deployment ID
func ValidateProviders2DeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2DeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Deployment ID
func (id Providers2DeploymentId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Resources/deployments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.DeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Deployment ID
func (id Providers2DeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResources", "Microsoft.Resources", "Microsoft.Resources"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
	}
}

// String returns a human-readable description of this Providers 2 Deployment ID
func (id Providers2DeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
	}
	return fmt.Sprintf("Providers 2 Deployment (%s)", strings.Join(components, "\n"))
}
