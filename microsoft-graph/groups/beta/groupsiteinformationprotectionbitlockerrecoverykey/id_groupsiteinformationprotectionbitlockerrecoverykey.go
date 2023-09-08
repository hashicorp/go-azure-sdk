package groupsiteinformationprotectionbitlockerrecoverykey

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionBitlockerRecoveryKeyId{}

// GroupSiteInformationProtectionBitlockerRecoveryKeyId is a struct representing the Resource ID for a Group Site Information Protection Bitlocker Recovery Key
type GroupSiteInformationProtectionBitlockerRecoveryKeyId struct {
	GroupId                string
	SiteId                 string
	BitlockerRecoveryKeyId string
}

// NewGroupSiteInformationProtectionBitlockerRecoveryKeyID returns a new GroupSiteInformationProtectionBitlockerRecoveryKeyId struct
func NewGroupSiteInformationProtectionBitlockerRecoveryKeyID(groupId string, siteId string, bitlockerRecoveryKeyId string) GroupSiteInformationProtectionBitlockerRecoveryKeyId {
	return GroupSiteInformationProtectionBitlockerRecoveryKeyId{
		GroupId:                groupId,
		SiteId:                 siteId,
		BitlockerRecoveryKeyId: bitlockerRecoveryKeyId,
	}
}

// ParseGroupSiteInformationProtectionBitlockerRecoveryKeyID parses 'input' into a GroupSiteInformationProtectionBitlockerRecoveryKeyId
func ParseGroupSiteInformationProtectionBitlockerRecoveryKeyID(input string) (*GroupSiteInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionBitlockerRecoveryKeyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.BitlockerRecoveryKeyId, ok = parsed.Parsed["bitlockerRecoveryKeyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "bitlockerRecoveryKeyId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteInformationProtectionBitlockerRecoveryKeyIDInsensitively parses 'input' case-insensitively into a GroupSiteInformationProtectionBitlockerRecoveryKeyId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteInformationProtectionBitlockerRecoveryKeyIDInsensitively(input string) (*GroupSiteInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionBitlockerRecoveryKeyId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.BitlockerRecoveryKeyId, ok = parsed.Parsed["bitlockerRecoveryKeyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "bitlockerRecoveryKeyId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteInformationProtectionBitlockerRecoveryKeyID checks that 'input' can be parsed as a Group Site Information Protection Bitlocker Recovery Key ID
func ValidateGroupSiteInformationProtectionBitlockerRecoveryKeyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteInformationProtectionBitlockerRecoveryKeyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Information Protection Bitlocker Recovery Key ID
func (id GroupSiteInformationProtectionBitlockerRecoveryKeyId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/bitlocker/recoveryKeys/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.BitlockerRecoveryKeyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Information Protection Bitlocker Recovery Key ID
func (id GroupSiteInformationProtectionBitlockerRecoveryKeyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticBitlocker", "bitlocker", "bitlocker"),
		resourceids.StaticSegment("staticRecoveryKeys", "recoveryKeys", "recoveryKeys"),
		resourceids.UserSpecifiedSegment("bitlockerRecoveryKeyId", "bitlockerRecoveryKeyIdValue"),
	}
}

// String returns a human-readable description of this Group Site Information Protection Bitlocker Recovery Key ID
func (id GroupSiteInformationProtectionBitlockerRecoveryKeyId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Bitlocker Recovery Key: %q", id.BitlockerRecoveryKeyId),
	}
	return fmt.Sprintf("Group Site Information Protection Bitlocker Recovery Key (%s)", strings.Join(components, "\n"))
}
