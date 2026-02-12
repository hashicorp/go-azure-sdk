package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeploymentManifestId{}

// DeploymentManifestId is a struct representing the Resource ID for a Deployment Manifest
type DeploymentManifestId struct {
	BaseURI              string
	DeploymentManifestId string
}

// NewDeploymentManifestID returns a new DeploymentManifestId struct
func NewDeploymentManifestID(baseURI string, deploymentManifestId string) DeploymentManifestId {
	return DeploymentManifestId{
		BaseURI:              strings.TrimSuffix(baseURI, "/"),
		DeploymentManifestId: deploymentManifestId,
	}
}

// ParseDeploymentManifestID parses 'input' into a DeploymentManifestId
func ParseDeploymentManifestID(input string) (*DeploymentManifestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeploymentManifestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeploymentManifestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeploymentManifestIDInsensitively parses 'input' case-insensitively into a DeploymentManifestId
// note: this method should only be used for API response data and not user input
func ParseDeploymentManifestIDInsensitively(input string) (*DeploymentManifestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeploymentManifestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeploymentManifestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeploymentManifestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeploymentManifestId, ok = input.Parsed["deploymentManifestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deploymentManifestId", input)
	}

	return nil
}

// ValidateDeploymentManifestID checks that 'input' can be parsed as a Deployment Manifest ID
func ValidateDeploymentManifestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeploymentManifestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deployment Manifest ID
func (id DeploymentManifestId) ID() string {
	fmtString := "%s/deploymentManifests/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeploymentManifestId)
}

// Path returns the formatted Deployment Manifest ID without the BaseURI
func (id DeploymentManifestId) Path() string {
	fmtString := "/deploymentManifests/%s"
	return fmt.Sprintf(fmtString, id.DeploymentManifestId)
}

// PathElements returns the values of Deployment Manifest ID Segments without the BaseURI
func (id DeploymentManifestId) PathElements() []any {
	return []any{id.DeploymentManifestId}
}

// Segments returns a slice of Resource ID Segments which comprise this Deployment Manifest ID
func (id DeploymentManifestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDeploymentManifests", "deploymentManifests", "deploymentManifests"),
		resourceids.UserSpecifiedSegment("deploymentManifestId", "deploymentManifestId"),
	}
}

// String returns a human-readable description of this Deployment Manifest ID
func (id DeploymentManifestId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Deployment Manifest: %q", id.DeploymentManifestId),
	}
	return fmt.Sprintf("Deployment Manifest (%s)", strings.Join(components, "\n"))
}
