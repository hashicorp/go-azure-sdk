package groupsiteinformationprotectionsensitivitylabelsublabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionSensitivityLabelId{}

// GroupSiteInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a Group Site Information Protection Sensitivity Label
type GroupSiteInformationProtectionSensitivityLabelId struct {
	GroupId            string
	SiteId             string
	SensitivityLabelId string
}

// NewGroupSiteInformationProtectionSensitivityLabelID returns a new GroupSiteInformationProtectionSensitivityLabelId struct
func NewGroupSiteInformationProtectionSensitivityLabelID(groupId string, siteId string, sensitivityLabelId string) GroupSiteInformationProtectionSensitivityLabelId {
	return GroupSiteInformationProtectionSensitivityLabelId{
		GroupId:            groupId,
		SiteId:             siteId,
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseGroupSiteInformationProtectionSensitivityLabelID parses 'input' into a GroupSiteInformationProtectionSensitivityLabelId
func ParseGroupSiteInformationProtectionSensitivityLabelID(input string) (*GroupSiteInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionSensitivityLabelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a GroupSiteInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteInformationProtectionSensitivityLabelIDInsensitively(input string) (*GroupSiteInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionSensitivityLabelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a Group Site Information Protection Sensitivity Label ID
func ValidateGroupSiteInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Information Protection Sensitivity Label ID
func (id GroupSiteInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Information Protection Sensitivity Label ID
func (id GroupSiteInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticSensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelIdValue"),
	}
}

// String returns a human-readable description of this Group Site Information Protection Sensitivity Label ID
func (id GroupSiteInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("Group Site Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
