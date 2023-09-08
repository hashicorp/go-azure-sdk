package groupsiteinformationprotectionthreatassessmentrequestresult

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionThreatAssessmentRequestResultId{}

// GroupSiteInformationProtectionThreatAssessmentRequestResultId is a struct representing the Resource ID for a Group Site Information Protection Threat Assessment Request Result
type GroupSiteInformationProtectionThreatAssessmentRequestResultId struct {
	GroupId                   string
	SiteId                    string
	ThreatAssessmentRequestId string
	ThreatAssessmentResultId  string
}

// NewGroupSiteInformationProtectionThreatAssessmentRequestResultID returns a new GroupSiteInformationProtectionThreatAssessmentRequestResultId struct
func NewGroupSiteInformationProtectionThreatAssessmentRequestResultID(groupId string, siteId string, threatAssessmentRequestId string, threatAssessmentResultId string) GroupSiteInformationProtectionThreatAssessmentRequestResultId {
	return GroupSiteInformationProtectionThreatAssessmentRequestResultId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
		ThreatAssessmentResultId:  threatAssessmentResultId,
	}
}

// ParseGroupSiteInformationProtectionThreatAssessmentRequestResultID parses 'input' into a GroupSiteInformationProtectionThreatAssessmentRequestResultId
func ParseGroupSiteInformationProtectionThreatAssessmentRequestResultID(input string) (*GroupSiteInformationProtectionThreatAssessmentRequestResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionThreatAssessmentRequestResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionThreatAssessmentRequestResultId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	if id.ThreatAssessmentResultId, ok = parsed.Parsed["threatAssessmentResultId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteInformationProtectionThreatAssessmentRequestResultIDInsensitively parses 'input' case-insensitively into a GroupSiteInformationProtectionThreatAssessmentRequestResultId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteInformationProtectionThreatAssessmentRequestResultIDInsensitively(input string) (*GroupSiteInformationProtectionThreatAssessmentRequestResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionThreatAssessmentRequestResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionThreatAssessmentRequestResultId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	if id.ThreatAssessmentResultId, ok = parsed.Parsed["threatAssessmentResultId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteInformationProtectionThreatAssessmentRequestResultID checks that 'input' can be parsed as a Group Site Information Protection Threat Assessment Request Result ID
func ValidateGroupSiteInformationProtectionThreatAssessmentRequestResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteInformationProtectionThreatAssessmentRequestResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Information Protection Threat Assessment Request Result ID
func (id GroupSiteInformationProtectionThreatAssessmentRequestResultId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/threatAssessmentRequests/%s/results/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ThreatAssessmentRequestId, id.ThreatAssessmentResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Information Protection Threat Assessment Request Result ID
func (id GroupSiteInformationProtectionThreatAssessmentRequestResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticThreatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestIdValue"),
		resourceids.StaticSegment("staticResults", "results", "results"),
		resourceids.UserSpecifiedSegment("threatAssessmentResultId", "threatAssessmentResultIdValue"),
	}
}

// String returns a human-readable description of this Group Site Information Protection Threat Assessment Request Result ID
func (id GroupSiteInformationProtectionThreatAssessmentRequestResultId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
		fmt.Sprintf("Threat Assessment Result: %q", id.ThreatAssessmentResultId),
	}
	return fmt.Sprintf("Group Site Information Protection Threat Assessment Request Result (%s)", strings.Join(components, "\n"))
}
