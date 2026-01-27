package files

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NodeId{}

// NodeId is a struct representing the Resource ID for a Node
type NodeId struct {
	BaseURI string
	PoolId  string
	NodeId  string
}

// NewNodeID returns a new NodeId struct
func NewNodeID(baseURI string, poolId string, nodeId string) NodeId {
	return NodeId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		PoolId:  poolId,
		NodeId:  nodeId,
	}
}

// ParseNodeID parses 'input' into a NodeId
func ParseNodeID(input string) (*NodeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NodeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NodeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseNodeIDInsensitively parses 'input' case-insensitively into a NodeId
// note: this method should only be used for API response data and not user input
func ParseNodeIDInsensitively(input string) (*NodeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NodeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NodeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *NodeId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateNodeID checks that 'input' can be parsed as a Node ID
func ValidateNodeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNodeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Node ID
func (id NodeId) ID() string {
	fmtString := "%s/pools/%s/nodes/%s"
	return fmt.Sprintf(fmtString, id.BaseURI, id.PoolId, id.NodeId)
}

// Path returns the formatted Node ID without the BaseURI
func (id NodeId) Path() string {
	fmtString := "/pools/%s/nodes/%s"
	return fmt.Sprintf(fmtString, id.PoolId, id.NodeId)
}

// PathElements returns the values of Node ID Segments without the BaseURI
func (id NodeId) PathElements() []any {
	return []any{id.PoolId, id.NodeId}
}

// Segments returns a slice of Resource ID Segments which comprise this Node ID
func (id NodeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint_url"),
		resourceids.StaticSegment("staticPools", "pools", "pools"),
		resourceids.UserSpecifiedSegment("poolId", "poolId"),
		resourceids.StaticSegment("staticNodes", "nodes", "nodes"),
		resourceids.UserSpecifiedSegment("nodeId", "nodeId"),
	}
}

// String returns a human-readable description of this Node ID
func (id NodeId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Pool: %q", id.PoolId),
		fmt.Sprintf("Node: %q", id.NodeId),
	}
	return fmt.Sprintf("Node (%s)", strings.Join(components, "\n"))
}
