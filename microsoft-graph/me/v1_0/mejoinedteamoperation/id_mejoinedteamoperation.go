package mejoinedteamoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamOperationId{}

// MeJoinedTeamOperationId is a struct representing the Resource ID for a Me Joined Team Operation
type MeJoinedTeamOperationId struct {
	TeamId                string
	TeamsAsyncOperationId string
}

// NewMeJoinedTeamOperationID returns a new MeJoinedTeamOperationId struct
func NewMeJoinedTeamOperationID(teamId string, teamsAsyncOperationId string) MeJoinedTeamOperationId {
	return MeJoinedTeamOperationId{
		TeamId:                teamId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseMeJoinedTeamOperationID parses 'input' into a MeJoinedTeamOperationId
func ParseMeJoinedTeamOperationID(input string) (*MeJoinedTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamOperationId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamOperationIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamOperationId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamOperationIDInsensitively(input string) (*MeJoinedTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamOperationId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamOperationID checks that 'input' can be parsed as a Me Joined Team Operation ID
func ValidateMeJoinedTeamOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Operation ID
func (id MeJoinedTeamOperationId) ID() string {
	fmtString := "/me/joinedTeams/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Operation ID
func (id MeJoinedTeamOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Operation ID
func (id MeJoinedTeamOperationId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("Me Joined Team Operation (%s)", strings.Join(components, "\n"))
}
