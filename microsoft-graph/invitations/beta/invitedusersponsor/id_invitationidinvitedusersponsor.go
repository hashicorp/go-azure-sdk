package invitedusersponsor

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InvitationIdInvitedUserSponsorId{}

// InvitationIdInvitedUserSponsorId is a struct representing the Resource ID for a Invitation Id Invited User Sponsor
type InvitationIdInvitedUserSponsorId struct {
	InvitationId      string
	DirectoryObjectId string
}

// NewInvitationIdInvitedUserSponsorID returns a new InvitationIdInvitedUserSponsorId struct
func NewInvitationIdInvitedUserSponsorID(invitationId string, directoryObjectId string) InvitationIdInvitedUserSponsorId {
	return InvitationIdInvitedUserSponsorId{
		InvitationId:      invitationId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseInvitationIdInvitedUserSponsorID parses 'input' into a InvitationIdInvitedUserSponsorId
func ParseInvitationIdInvitedUserSponsorID(input string) (*InvitationIdInvitedUserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvitationIdInvitedUserSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvitationIdInvitedUserSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvitationIdInvitedUserSponsorIDInsensitively parses 'input' case-insensitively into a InvitationIdInvitedUserSponsorId
// note: this method should only be used for API response data and not user input
func ParseInvitationIdInvitedUserSponsorIDInsensitively(input string) (*InvitationIdInvitedUserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvitationIdInvitedUserSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvitationIdInvitedUserSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvitationIdInvitedUserSponsorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.InvitationId, ok = input.Parsed["invitationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invitationId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateInvitationIdInvitedUserSponsorID checks that 'input' can be parsed as a Invitation Id Invited User Sponsor ID
func ValidateInvitationIdInvitedUserSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvitationIdInvitedUserSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invitation Id Invited User Sponsor ID
func (id InvitationIdInvitedUserSponsorId) ID() string {
	fmtString := "/invitations/%s/invitedUserSponsors/%s"
	return fmt.Sprintf(fmtString, id.InvitationId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Invitation Id Invited User Sponsor ID
func (id InvitationIdInvitedUserSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("invitations", "invitations", "invitations"),
		resourceids.UserSpecifiedSegment("invitationId", "invitationIdValue"),
		resourceids.StaticSegment("invitedUserSponsors", "invitedUserSponsors", "invitedUserSponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Invitation Id Invited User Sponsor ID
func (id InvitationIdInvitedUserSponsorId) String() string {
	components := []string{
		fmt.Sprintf("Invitation: %q", id.InvitationId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Invitation Id Invited User Sponsor (%s)", strings.Join(components, "\n"))
}
