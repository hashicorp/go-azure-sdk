package invitationinvitedusersponsor

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = InvitationInvitedUserSponsorId{}

// InvitationInvitedUserSponsorId is a struct representing the Resource ID for a Invitation Invited User Sponsor
type InvitationInvitedUserSponsorId struct {
	InvitationId      string
	DirectoryObjectId string
}

// NewInvitationInvitedUserSponsorID returns a new InvitationInvitedUserSponsorId struct
func NewInvitationInvitedUserSponsorID(invitationId string, directoryObjectId string) InvitationInvitedUserSponsorId {
	return InvitationInvitedUserSponsorId{
		InvitationId:      invitationId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseInvitationInvitedUserSponsorID parses 'input' into a InvitationInvitedUserSponsorId
func ParseInvitationInvitedUserSponsorID(input string) (*InvitationInvitedUserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(InvitationInvitedUserSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := InvitationInvitedUserSponsorId{}

	if id.InvitationId, ok = parsed.Parsed["invitationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "invitationId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseInvitationInvitedUserSponsorIDInsensitively parses 'input' case-insensitively into a InvitationInvitedUserSponsorId
// note: this method should only be used for API response data and not user input
func ParseInvitationInvitedUserSponsorIDInsensitively(input string) (*InvitationInvitedUserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(InvitationInvitedUserSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := InvitationInvitedUserSponsorId{}

	if id.InvitationId, ok = parsed.Parsed["invitationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "invitationId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateInvitationInvitedUserSponsorID checks that 'input' can be parsed as a Invitation Invited User Sponsor ID
func ValidateInvitationInvitedUserSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvitationInvitedUserSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invitation Invited User Sponsor ID
func (id InvitationInvitedUserSponsorId) ID() string {
	fmtString := "/invitations/%s/invitedUserSponsors/%s"
	return fmt.Sprintf(fmtString, id.InvitationId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Invitation Invited User Sponsor ID
func (id InvitationInvitedUserSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticInvitations", "invitations", "invitations"),
		resourceids.UserSpecifiedSegment("invitationId", "invitationIdValue"),
		resourceids.StaticSegment("staticInvitedUserSponsors", "invitedUserSponsors", "invitedUserSponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Invitation Invited User Sponsor ID
func (id InvitationInvitedUserSponsorId) String() string {
	components := []string{
		fmt.Sprintf("Invitation: %q", id.InvitationId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Invitation Invited User Sponsor (%s)", strings.Join(components, "\n"))
}
