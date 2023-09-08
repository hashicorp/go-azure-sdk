package groupsiteinformationprotectionpolicylabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionPolicyLabelId{}

// GroupSiteInformationProtectionPolicyLabelId is a struct representing the Resource ID for a Group Site Information Protection Policy Label
type GroupSiteInformationProtectionPolicyLabelId struct {
	GroupId                      string
	SiteId                       string
	InformationProtectionLabelId string
}

// NewGroupSiteInformationProtectionPolicyLabelID returns a new GroupSiteInformationProtectionPolicyLabelId struct
func NewGroupSiteInformationProtectionPolicyLabelID(groupId string, siteId string, informationProtectionLabelId string) GroupSiteInformationProtectionPolicyLabelId {
	return GroupSiteInformationProtectionPolicyLabelId{
		GroupId:                      groupId,
		SiteId:                       siteId,
		InformationProtectionLabelId: informationProtectionLabelId,
	}
}

// ParseGroupSiteInformationProtectionPolicyLabelID parses 'input' into a GroupSiteInformationProtectionPolicyLabelId
func ParseGroupSiteInformationProtectionPolicyLabelID(input string) (*GroupSiteInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionPolicyLabelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.InformationProtectionLabelId, ok = parsed.Parsed["informationProtectionLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionLabelId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteInformationProtectionPolicyLabelIDInsensitively parses 'input' case-insensitively into a GroupSiteInformationProtectionPolicyLabelId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteInformationProtectionPolicyLabelIDInsensitively(input string) (*GroupSiteInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionPolicyLabelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.InformationProtectionLabelId, ok = parsed.Parsed["informationProtectionLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionLabelId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteInformationProtectionPolicyLabelID checks that 'input' can be parsed as a Group Site Information Protection Policy Label ID
func ValidateGroupSiteInformationProtectionPolicyLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteInformationProtectionPolicyLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Information Protection Policy Label ID
func (id GroupSiteInformationProtectionPolicyLabelId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/policy/labels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.InformationProtectionLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Information Protection Policy Label ID
func (id GroupSiteInformationProtectionPolicyLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticPolicy", "policy", "policy"),
		resourceids.StaticSegment("staticLabels", "labels", "labels"),
		resourceids.UserSpecifiedSegment("informationProtectionLabelId", "informationProtectionLabelIdValue"),
	}
}

// String returns a human-readable description of this Group Site Information Protection Policy Label ID
func (id GroupSiteInformationProtectionPolicyLabelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Information Protection Label: %q", id.InformationProtectionLabelId),
	}
	return fmt.Sprintf("Group Site Information Protection Policy Label (%s)", strings.Join(components, "\n"))
}
