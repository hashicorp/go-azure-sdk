package groupsiteinformationprotectionsensitivitylabelsublabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionSensitivityLabelSublabelId{}

// GroupSiteInformationProtectionSensitivityLabelSublabelId is a struct representing the Resource ID for a Group Site Information Protection Sensitivity Label Sublabel
type GroupSiteInformationProtectionSensitivityLabelSublabelId struct {
	GroupId             string
	SiteId              string
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewGroupSiteInformationProtectionSensitivityLabelSublabelID returns a new GroupSiteInformationProtectionSensitivityLabelSublabelId struct
func NewGroupSiteInformationProtectionSensitivityLabelSublabelID(groupId string, siteId string, sensitivityLabelId string, sensitivityLabelId1 string) GroupSiteInformationProtectionSensitivityLabelSublabelId {
	return GroupSiteInformationProtectionSensitivityLabelSublabelId{
		GroupId:             groupId,
		SiteId:              siteId,
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseGroupSiteInformationProtectionSensitivityLabelSublabelID parses 'input' into a GroupSiteInformationProtectionSensitivityLabelSublabelId
func ParseGroupSiteInformationProtectionSensitivityLabelSublabelID(input string) (*GroupSiteInformationProtectionSensitivityLabelSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionSensitivityLabelSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionSensitivityLabelSublabelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	if id.SensitivityLabelId1, ok = parsed.Parsed["sensitivityLabelId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteInformationProtectionSensitivityLabelSublabelIDInsensitively parses 'input' case-insensitively into a GroupSiteInformationProtectionSensitivityLabelSublabelId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteInformationProtectionSensitivityLabelSublabelIDInsensitively(input string) (*GroupSiteInformationProtectionSensitivityLabelSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionSensitivityLabelSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionSensitivityLabelSublabelId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	if id.SensitivityLabelId1, ok = parsed.Parsed["sensitivityLabelId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteInformationProtectionSensitivityLabelSublabelID checks that 'input' can be parsed as a Group Site Information Protection Sensitivity Label Sublabel ID
func ValidateGroupSiteInformationProtectionSensitivityLabelSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteInformationProtectionSensitivityLabelSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Information Protection Sensitivity Label Sublabel ID
func (id GroupSiteInformationProtectionSensitivityLabelSublabelId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Information Protection Sensitivity Label Sublabel ID
func (id GroupSiteInformationProtectionSensitivityLabelSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticSensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelIdValue"),
		resourceids.StaticSegment("staticSublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1Value"),
	}
}

// String returns a human-readable description of this Group Site Information Protection Sensitivity Label Sublabel ID
func (id GroupSiteInformationProtectionSensitivityLabelSublabelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("Group Site Information Protection Sensitivity Label Sublabel (%s)", strings.Join(components, "\n"))
}
