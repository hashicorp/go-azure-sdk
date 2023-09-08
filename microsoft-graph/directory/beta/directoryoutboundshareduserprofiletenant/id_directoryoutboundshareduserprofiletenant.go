package directoryoutboundshareduserprofiletenant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryOutboundSharedUserProfileTenantId{}

// DirectoryOutboundSharedUserProfileTenantId is a struct representing the Resource ID for a Directory Outbound Shared User Profile Tenant
type DirectoryOutboundSharedUserProfileTenantId struct {
	OutboundSharedUserProfileUserId string
	TenantReferenceTenantId         string
}

// NewDirectoryOutboundSharedUserProfileTenantID returns a new DirectoryOutboundSharedUserProfileTenantId struct
func NewDirectoryOutboundSharedUserProfileTenantID(outboundSharedUserProfileUserId string, tenantReferenceTenantId string) DirectoryOutboundSharedUserProfileTenantId {
	return DirectoryOutboundSharedUserProfileTenantId{
		OutboundSharedUserProfileUserId: outboundSharedUserProfileUserId,
		TenantReferenceTenantId:         tenantReferenceTenantId,
	}
}

// ParseDirectoryOutboundSharedUserProfileTenantID parses 'input' into a DirectoryOutboundSharedUserProfileTenantId
func ParseDirectoryOutboundSharedUserProfileTenantID(input string) (*DirectoryOutboundSharedUserProfileTenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryOutboundSharedUserProfileTenantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryOutboundSharedUserProfileTenantId{}

	if id.OutboundSharedUserProfileUserId, ok = parsed.Parsed["outboundSharedUserProfileUserId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outboundSharedUserProfileUserId", *parsed)
	}

	if id.TenantReferenceTenantId, ok = parsed.Parsed["tenantReferenceTenantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tenantReferenceTenantId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryOutboundSharedUserProfileTenantIDInsensitively parses 'input' case-insensitively into a DirectoryOutboundSharedUserProfileTenantId
// note: this method should only be used for API response data and not user input
func ParseDirectoryOutboundSharedUserProfileTenantIDInsensitively(input string) (*DirectoryOutboundSharedUserProfileTenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryOutboundSharedUserProfileTenantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryOutboundSharedUserProfileTenantId{}

	if id.OutboundSharedUserProfileUserId, ok = parsed.Parsed["outboundSharedUserProfileUserId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outboundSharedUserProfileUserId", *parsed)
	}

	if id.TenantReferenceTenantId, ok = parsed.Parsed["tenantReferenceTenantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tenantReferenceTenantId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryOutboundSharedUserProfileTenantID checks that 'input' can be parsed as a Directory Outbound Shared User Profile Tenant ID
func ValidateDirectoryOutboundSharedUserProfileTenantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryOutboundSharedUserProfileTenantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Outbound Shared User Profile Tenant ID
func (id DirectoryOutboundSharedUserProfileTenantId) ID() string {
	fmtString := "/directory/outboundSharedUserProfiles/%s/tenants/%s"
	return fmt.Sprintf(fmtString, id.OutboundSharedUserProfileUserId, id.TenantReferenceTenantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Outbound Shared User Profile Tenant ID
func (id DirectoryOutboundSharedUserProfileTenantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticOutboundSharedUserProfiles", "outboundSharedUserProfiles", "outboundSharedUserProfiles"),
		resourceids.UserSpecifiedSegment("outboundSharedUserProfileUserId", "outboundSharedUserProfileUserIdValue"),
		resourceids.StaticSegment("staticTenants", "tenants", "tenants"),
		resourceids.UserSpecifiedSegment("tenantReferenceTenantId", "tenantReferenceTenantIdValue"),
	}
}

// String returns a human-readable description of this Directory Outbound Shared User Profile Tenant ID
func (id DirectoryOutboundSharedUserProfileTenantId) String() string {
	components := []string{
		fmt.Sprintf("Outbound Shared User Profile User: %q", id.OutboundSharedUserProfileUserId),
		fmt.Sprintf("Tenant Reference Tenant: %q", id.TenantReferenceTenantId),
	}
	return fmt.Sprintf("Directory Outbound Shared User Profile Tenant (%s)", strings.Join(components, "\n"))
}
