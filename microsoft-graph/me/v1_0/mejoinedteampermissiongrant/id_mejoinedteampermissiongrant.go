package mejoinedteampermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPermissionGrantId{}

// MeJoinedTeamPermissionGrantId is a struct representing the Resource ID for a Me Joined Team Permission Grant
type MeJoinedTeamPermissionGrantId struct {
	TeamId                            string
	ResourceSpecificPermissionGrantId string
}

// NewMeJoinedTeamPermissionGrantID returns a new MeJoinedTeamPermissionGrantId struct
func NewMeJoinedTeamPermissionGrantID(teamId string, resourceSpecificPermissionGrantId string) MeJoinedTeamPermissionGrantId {
	return MeJoinedTeamPermissionGrantId{
		TeamId:                            teamId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseMeJoinedTeamPermissionGrantID parses 'input' into a MeJoinedTeamPermissionGrantId
func ParseMeJoinedTeamPermissionGrantID(input string) (*MeJoinedTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPermissionGrantId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPermissionGrantIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPermissionGrantIDInsensitively(input string) (*MeJoinedTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPermissionGrantId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPermissionGrantID checks that 'input' can be parsed as a Me Joined Team Permission Grant ID
func ValidateMeJoinedTeamPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Permission Grant ID
func (id MeJoinedTeamPermissionGrantId) ID() string {
	fmtString := "/me/joinedTeams/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Permission Grant ID
func (id MeJoinedTeamPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPermissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Permission Grant ID
func (id MeJoinedTeamPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Me Joined Team Permission Grant (%s)", strings.Join(components, "\n"))
}
