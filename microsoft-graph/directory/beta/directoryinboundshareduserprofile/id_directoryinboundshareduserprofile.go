package directoryinboundshareduserprofile

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryInboundSharedUserProfileId{}

// DirectoryInboundSharedUserProfileId is a struct representing the Resource ID for a Directory Inbound Shared User Profile
type DirectoryInboundSharedUserProfileId struct {
	InboundSharedUserProfileUserId string
}

// NewDirectoryInboundSharedUserProfileID returns a new DirectoryInboundSharedUserProfileId struct
func NewDirectoryInboundSharedUserProfileID(inboundSharedUserProfileUserId string) DirectoryInboundSharedUserProfileId {
	return DirectoryInboundSharedUserProfileId{
		InboundSharedUserProfileUserId: inboundSharedUserProfileUserId,
	}
}

// ParseDirectoryInboundSharedUserProfileID parses 'input' into a DirectoryInboundSharedUserProfileId
func ParseDirectoryInboundSharedUserProfileID(input string) (*DirectoryInboundSharedUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryInboundSharedUserProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryInboundSharedUserProfileId{}

	if id.InboundSharedUserProfileUserId, ok = parsed.Parsed["inboundSharedUserProfileUserId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "inboundSharedUserProfileUserId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryInboundSharedUserProfileIDInsensitively parses 'input' case-insensitively into a DirectoryInboundSharedUserProfileId
// note: this method should only be used for API response data and not user input
func ParseDirectoryInboundSharedUserProfileIDInsensitively(input string) (*DirectoryInboundSharedUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryInboundSharedUserProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryInboundSharedUserProfileId{}

	if id.InboundSharedUserProfileUserId, ok = parsed.Parsed["inboundSharedUserProfileUserId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "inboundSharedUserProfileUserId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryInboundSharedUserProfileID checks that 'input' can be parsed as a Directory Inbound Shared User Profile ID
func ValidateDirectoryInboundSharedUserProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryInboundSharedUserProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Inbound Shared User Profile ID
func (id DirectoryInboundSharedUserProfileId) ID() string {
	fmtString := "/directory/inboundSharedUserProfiles/%s"
	return fmt.Sprintf(fmtString, id.InboundSharedUserProfileUserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Inbound Shared User Profile ID
func (id DirectoryInboundSharedUserProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticInboundSharedUserProfiles", "inboundSharedUserProfiles", "inboundSharedUserProfiles"),
		resourceids.UserSpecifiedSegment("inboundSharedUserProfileUserId", "inboundSharedUserProfileUserIdValue"),
	}
}

// String returns a human-readable description of this Directory Inbound Shared User Profile ID
func (id DirectoryInboundSharedUserProfileId) String() string {
	components := []string{
		fmt.Sprintf("Inbound Shared User Profile User: %q", id.InboundSharedUserProfileUserId),
	}
	return fmt.Sprintf("Directory Inbound Shared User Profile (%s)", strings.Join(components, "\n"))
}
