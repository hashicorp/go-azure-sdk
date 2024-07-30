package inviteduser

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InvitationId{}

// InvitationId is a struct representing the Resource ID for a Invitation
type InvitationId struct {
	InvitationId string
}

// NewInvitationID returns a new InvitationId struct
func NewInvitationID(invitationId string) InvitationId {
	return InvitationId{
		InvitationId: invitationId,
	}
}

// ParseInvitationID parses 'input' into a InvitationId
func ParseInvitationID(input string) (*InvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvitationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvitationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvitationIDInsensitively parses 'input' case-insensitively into a InvitationId
// note: this method should only be used for API response data and not user input
func ParseInvitationIDInsensitively(input string) (*InvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvitationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvitationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvitationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.InvitationId, ok = input.Parsed["invitationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invitationId", input)
	}

	return nil
}

// ValidateInvitationID checks that 'input' can be parsed as a Invitation ID
func ValidateInvitationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvitationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invitation ID
func (id InvitationId) ID() string {
	fmtString := "/invitations/%s"
	return fmt.Sprintf(fmtString, id.InvitationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Invitation ID
func (id InvitationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("invitations", "invitations", "invitations"),
		resourceids.UserSpecifiedSegment("invitationId", "invitationIdValue"),
	}
}

// String returns a human-readable description of this Invitation ID
func (id InvitationId) String() string {
	components := []string{
		fmt.Sprintf("Invitation: %q", id.InvitationId),
	}
	return fmt.Sprintf("Invitation (%s)", strings.Join(components, "\n"))
}
