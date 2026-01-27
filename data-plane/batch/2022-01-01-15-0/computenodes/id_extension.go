package computenodes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExtensionId{}

// ExtensionId is a struct representing the Resource ID for a Extension
type ExtensionId struct {
	BaseURI       string
	PoolId        string
	NodeId        string
	ExtensionName string
}

// NewExtensionID returns a new ExtensionId struct
func NewExtensionID(baseURI string, poolId string, nodeId string, extensionName string) ExtensionId {
	return ExtensionId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		PoolId:        poolId,
		NodeId:        nodeId,
		ExtensionName: extensionName,
	}
}

// ParseExtensionID parses 'input' into a ExtensionId
func ParseExtensionID(input string) (*ExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExtensionIDInsensitively parses 'input' case-insensitively into a ExtensionId
// note: this method should only be used for API response data and not user input
func ParseExtensionIDInsensitively(input string) (*ExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.PoolId, ok = input.Parsed["poolId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "poolId", input)
	}

	if id.NodeId, ok = input.Parsed["nodeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "nodeId", input)
	}

	if id.ExtensionName, ok = input.Parsed["extensionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionName", input)
	}

	return nil
}

// ValidateExtensionID checks that 'input' can be parsed as a Extension ID
func ValidateExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Extension ID
func (id ExtensionId) ID() string {
	fmtString := "%s/pools/%s/nodes/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.BaseURI, id.PoolId, id.NodeId, id.ExtensionName)
}

// Path returns the formatted Extension ID without the BaseURI
func (id ExtensionId) Path() string {
	fmtString := "/pools/%s/nodes/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.PoolId, id.NodeId, id.ExtensionName)
}

// PathElements returns the values of Extension ID Segments without the BaseURI
func (id ExtensionId) PathElements() []any {
	return []any{id.PoolId, id.NodeId, id.ExtensionName}
}

// Segments returns a slice of Resource ID Segments which comprise this Extension ID
func (id ExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint_url"),
		resourceids.StaticSegment("staticPools", "pools", "pools"),
		resourceids.UserSpecifiedSegment("poolId", "poolId"),
		resourceids.StaticSegment("staticNodes", "nodes", "nodes"),
		resourceids.UserSpecifiedSegment("nodeId", "nodeId"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionName", "extensionName"),
	}
}

// String returns a human-readable description of this Extension ID
func (id ExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Pool: %q", id.PoolId),
		fmt.Sprintf("Node: %q", id.NodeId),
		fmt.Sprintf("Extension Name: %q", id.ExtensionName),
	}
	return fmt.Sprintf("Extension (%s)", strings.Join(components, "\n"))
}
