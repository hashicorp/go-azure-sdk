package files

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NodeFileId{}

// NodeFileId is a struct representing the Resource ID for a Node File
type NodeFileId struct {
	BaseURI  string
	PoolId   string
	NodeId   string
	FileName string
}

// NewNodeFileID returns a new NodeFileId struct
func NewNodeFileID(baseURI string, poolId string, nodeId string, fileName string) NodeFileId {
	return NodeFileId{
		BaseURI:  strings.TrimSuffix(baseURI, "/"),
		PoolId:   poolId,
		NodeId:   nodeId,
		FileName: fileName,
	}
}

// ParseNodeFileID parses 'input' into a NodeFileId
func ParseNodeFileID(input string) (*NodeFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NodeFileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NodeFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseNodeFileIDInsensitively parses 'input' case-insensitively into a NodeFileId
// note: this method should only be used for API response data and not user input
func ParseNodeFileIDInsensitively(input string) (*NodeFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NodeFileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NodeFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *NodeFileId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.FileName, ok = input.Parsed["fileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "fileName", input)
	}

	return nil
}

// ValidateNodeFileID checks that 'input' can be parsed as a Node File ID
func ValidateNodeFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNodeFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Node File ID
func (id NodeFileId) ID() string {
	fmtString := "%s/pools/%s/nodes/%s/files/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.PoolId, id.NodeId, id.FileName)
}

// Path returns the formatted Node File ID without the BaseURI
func (id NodeFileId) Path() string {
	fmtString := "/pools/%s/nodes/%s/files/%s"
	return fmt.Sprintf(fmtString, id.PoolId, id.NodeId, id.FileName)
}

// PathElements returns the values of Node File ID Segments without the BaseURI
func (id NodeFileId) PathElements() []any {
	return []any{id.PoolId, id.NodeId, id.FileName}
}

// Segments returns a slice of Resource ID Segments which comprise this Node File ID
func (id NodeFileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticPools", "pools", "pools"),
		resourceids.UserSpecifiedSegment("poolId", "poolId"),
		resourceids.StaticSegment("staticNodes", "nodes", "nodes"),
		resourceids.UserSpecifiedSegment("nodeId", "nodeId"),
		resourceids.StaticSegment("staticFiles", "files", "files"),
		resourceids.UserSpecifiedSegment("fileName", "fileName"),
	}
}

// String returns a human-readable description of this Node File ID
func (id NodeFileId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Pool: %q", id.PoolId),
		fmt.Sprintf("Node: %q", id.NodeId),
		fmt.Sprintf("File Name: %q", id.FileName),
	}
	return fmt.Sprintf("Node File (%s)", strings.Join(components, "\n"))
}
