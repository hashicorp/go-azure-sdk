package domainsharedemaildomaininvitation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainSharedEmailDomainInvitationId{}

// DomainSharedEmailDomainInvitationId is a struct representing the Resource ID for a Domain Shared Email Domain Invitation
type DomainSharedEmailDomainInvitationId struct {
	DomainId                      string
	SharedEmailDomainInvitationId string
}

// NewDomainSharedEmailDomainInvitationID returns a new DomainSharedEmailDomainInvitationId struct
func NewDomainSharedEmailDomainInvitationID(domainId string, sharedEmailDomainInvitationId string) DomainSharedEmailDomainInvitationId {
	return DomainSharedEmailDomainInvitationId{
		DomainId:                      domainId,
		SharedEmailDomainInvitationId: sharedEmailDomainInvitationId,
	}
}

// ParseDomainSharedEmailDomainInvitationID parses 'input' into a DomainSharedEmailDomainInvitationId
func ParseDomainSharedEmailDomainInvitationID(input string) (*DomainSharedEmailDomainInvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainSharedEmailDomainInvitationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainSharedEmailDomainInvitationId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.SharedEmailDomainInvitationId, ok = parsed.Parsed["sharedEmailDomainInvitationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedEmailDomainInvitationId", *parsed)
	}

	return &id, nil
}

// ParseDomainSharedEmailDomainInvitationIDInsensitively parses 'input' case-insensitively into a DomainSharedEmailDomainInvitationId
// note: this method should only be used for API response data and not user input
func ParseDomainSharedEmailDomainInvitationIDInsensitively(input string) (*DomainSharedEmailDomainInvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainSharedEmailDomainInvitationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainSharedEmailDomainInvitationId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.SharedEmailDomainInvitationId, ok = parsed.Parsed["sharedEmailDomainInvitationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedEmailDomainInvitationId", *parsed)
	}

	return &id, nil
}

// ValidateDomainSharedEmailDomainInvitationID checks that 'input' can be parsed as a Domain Shared Email Domain Invitation ID
func ValidateDomainSharedEmailDomainInvitationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainSharedEmailDomainInvitationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Shared Email Domain Invitation ID
func (id DomainSharedEmailDomainInvitationId) ID() string {
	fmtString := "/domains/%s/sharedEmailDomainInvitations/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.SharedEmailDomainInvitationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Shared Email Domain Invitation ID
func (id DomainSharedEmailDomainInvitationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainIdValue"),
		resourceids.StaticSegment("staticSharedEmailDomainInvitations", "sharedEmailDomainInvitations", "sharedEmailDomainInvitations"),
		resourceids.UserSpecifiedSegment("sharedEmailDomainInvitationId", "sharedEmailDomainInvitationIdValue"),
	}
}

// String returns a human-readable description of this Domain Shared Email Domain Invitation ID
func (id DomainSharedEmailDomainInvitationId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Shared Email Domain Invitation: %q", id.SharedEmailDomainInvitationId),
	}
	return fmt.Sprintf("Domain Shared Email Domain Invitation (%s)", strings.Join(components, "\n"))
}
