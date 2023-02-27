// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DeploymentId{}

// DeploymentId is a struct representing the Resource ID for a Deployment
type DeploymentId struct {
	DeploymentName string
}

// NewDeploymentID returns a new DeploymentId struct
func NewDeploymentID(deploymentName string) DeploymentId {
	return DeploymentId{
		DeploymentName: deploymentName,
	}
}

// ParseDeploymentID parses 'input' into a DeploymentId
func ParseDeploymentID(input string) (*DeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeploymentId{}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDeploymentIDInsensitively parses 'input' case-insensitively into a DeploymentId
// note: this method should only be used for API response data and not user input
func ParseDeploymentIDInsensitively(input string) (*DeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeploymentId{}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDeploymentID checks that 'input' can be parsed as a Deployment ID
func ValidateDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deployment ID
func (id DeploymentId) ID() string {
	fmtString := "/providers/Microsoft.Resources/deployments/%s"
	return fmt.Sprintf(fmtString, id.DeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Deployment ID
func (id DeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResources", "Microsoft.Resources", "Microsoft.Resources"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
	}
}

// String returns a human-readable description of this Deployment ID
func (id DeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
	}
	return fmt.Sprintf("Deployment (%s)", strings.Join(components, "\n"))
}
