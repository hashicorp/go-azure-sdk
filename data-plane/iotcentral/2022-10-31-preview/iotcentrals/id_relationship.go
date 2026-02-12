package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RelationshipId{}

// RelationshipId is a struct representing the Resource ID for a Relationship
type RelationshipId struct {
	BaseURI        string
	DeviceId       string
	RelationshipId string
}

// NewRelationshipID returns a new RelationshipId struct
func NewRelationshipID(baseURI string, deviceId string, relationshipId string) RelationshipId {
	return RelationshipId{
		BaseURI:        strings.TrimSuffix(baseURI, "/"),
		DeviceId:       deviceId,
		RelationshipId: relationshipId,
	}
}

// ParseRelationshipID parses 'input' into a RelationshipId
func ParseRelationshipID(input string) (*RelationshipId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RelationshipId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RelationshipId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRelationshipIDInsensitively parses 'input' case-insensitively into a RelationshipId
// note: this method should only be used for API response data and not user input
func ParseRelationshipIDInsensitively(input string) (*RelationshipId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RelationshipId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RelationshipId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RelationshipId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.RelationshipId, ok = input.Parsed["relationshipId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationshipId", input)
	}

	return nil
}

// ValidateRelationshipID checks that 'input' can be parsed as a Relationship ID
func ValidateRelationshipID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRelationshipID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Relationship ID
func (id RelationshipId) ID() string {
	fmtString := "%s/devices/%s/relationships/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.RelationshipId)
}

// Path returns the formatted Relationship ID without the BaseURI
func (id RelationshipId) Path() string {
	fmtString := "/devices/%s/relationships/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.RelationshipId)
}

// PathElements returns the values of Relationship ID Segments without the BaseURI
func (id RelationshipId) PathElements() []any {
	return []any{id.DeviceId, id.RelationshipId}
}

// Segments returns a slice of Resource ID Segments which comprise this Relationship ID
func (id RelationshipId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticRelationships", "relationships", "relationships"),
		resourceids.UserSpecifiedSegment("relationshipId", "relationshipId"),
	}
}

// String returns a human-readable description of this Relationship ID
func (id RelationshipId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Relationship: %q", id.RelationshipId),
	}
	return fmt.Sprintf("Relationship (%s)", strings.Join(components, "\n"))
}
