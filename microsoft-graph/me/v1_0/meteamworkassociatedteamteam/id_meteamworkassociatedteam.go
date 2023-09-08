package meteamworkassociatedteamteam

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeTeamworkAssociatedTeamId{}

// MeTeamworkAssociatedTeamId is a struct representing the Resource ID for a Me Teamwork Associated Team
type MeTeamworkAssociatedTeamId struct {
	AssociatedTeamInfoId string
}

// NewMeTeamworkAssociatedTeamID returns a new MeTeamworkAssociatedTeamId struct
func NewMeTeamworkAssociatedTeamID(associatedTeamInfoId string) MeTeamworkAssociatedTeamId {
	return MeTeamworkAssociatedTeamId{
		AssociatedTeamInfoId: associatedTeamInfoId,
	}
}

// ParseMeTeamworkAssociatedTeamID parses 'input' into a MeTeamworkAssociatedTeamId
func ParseMeTeamworkAssociatedTeamID(input string) (*MeTeamworkAssociatedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTeamworkAssociatedTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTeamworkAssociatedTeamId{}

	if id.AssociatedTeamInfoId, ok = parsed.Parsed["associatedTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "associatedTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseMeTeamworkAssociatedTeamIDInsensitively parses 'input' case-insensitively into a MeTeamworkAssociatedTeamId
// note: this method should only be used for API response data and not user input
func ParseMeTeamworkAssociatedTeamIDInsensitively(input string) (*MeTeamworkAssociatedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeTeamworkAssociatedTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeTeamworkAssociatedTeamId{}

	if id.AssociatedTeamInfoId, ok = parsed.Parsed["associatedTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "associatedTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateMeTeamworkAssociatedTeamID checks that 'input' can be parsed as a Me Teamwork Associated Team ID
func ValidateMeTeamworkAssociatedTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeTeamworkAssociatedTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Teamwork Associated Team ID
func (id MeTeamworkAssociatedTeamId) ID() string {
	fmtString := "/me/teamwork/associatedTeams/%s"
	return fmt.Sprintf(fmtString, id.AssociatedTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Teamwork Associated Team ID
func (id MeTeamworkAssociatedTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticTeamwork", "teamwork", "teamwork"),
		resourceids.StaticSegment("staticAssociatedTeams", "associatedTeams", "associatedTeams"),
		resourceids.UserSpecifiedSegment("associatedTeamInfoId", "associatedTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this Me Teamwork Associated Team ID
func (id MeTeamworkAssociatedTeamId) String() string {
	components := []string{
		fmt.Sprintf("Associated Team Info: %q", id.AssociatedTeamInfoId),
	}
	return fmt.Sprintf("Me Teamwork Associated Team (%s)", strings.Join(components, "\n"))
}
