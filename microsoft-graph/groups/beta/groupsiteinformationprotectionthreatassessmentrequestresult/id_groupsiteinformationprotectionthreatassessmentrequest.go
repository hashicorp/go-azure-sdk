package groupsiteinformationprotectionthreatassessmentrequestresult

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionThreatAssessmentRequestId{}

// GroupSiteInformationProtectionThreatAssessmentRequestId is a struct representing the Resource ID for a Group Site Information Protection Threat Assessment Request
type GroupSiteInformationProtectionThreatAssessmentRequestId struct {
	GroupId                   string
	SiteId                    string
	ThreatAssessmentRequestId string
}

// NewGroupSiteInformationProtectionThreatAssessmentRequestID returns a new GroupSiteInformationProtectionThreatAssessmentRequestId struct
func NewGroupSiteInformationProtectionThreatAssessmentRequestID(groupId string, siteId string, threatAssessmentRequestId string) GroupSiteInformationProtectionThreatAssessmentRequestId {
	return GroupSiteInformationProtectionThreatAssessmentRequestId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
	}
}

// ParseGroupSiteInformationProtectionThreatAssessmentRequestID parses 'input' into a GroupSiteInformationProtectionThreatAssessmentRequestId
func ParseGroupSiteInformationProtectionThreatAssessmentRequestID(input string) (*GroupSiteInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionThreatAssessmentRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	return &id, nil
}

// ParseGroupSiteInformationProtectionThreatAssessmentRequestIDInsensitively parses 'input' case-insensitively into a GroupSiteInformationProtectionThreatAssessmentRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupSiteInformationProtectionThreatAssessmentRequestIDInsensitively(input string) (*GroupSiteInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupSiteInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupSiteInformationProtectionThreatAssessmentRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	return &id, nil
}

// ValidateGroupSiteInformationProtectionThreatAssessmentRequestID checks that 'input' can be parsed as a Group Site Information Protection Threat Assessment Request ID
func ValidateGroupSiteInformationProtectionThreatAssessmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupSiteInformationProtectionThreatAssessmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Site Information Protection Threat Assessment Request ID
func (id GroupSiteInformationProtectionThreatAssessmentRequestId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/threatAssessmentRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ThreatAssessmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Site Information Protection Threat Assessment Request ID
func (id GroupSiteInformationProtectionThreatAssessmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticThreatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestIdValue"),
	}
}

// String returns a human-readable description of this Group Site Information Protection Threat Assessment Request ID
func (id GroupSiteInformationProtectionThreatAssessmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
	}
	return fmt.Sprintf("Group Site Information Protection Threat Assessment Request (%s)", strings.Join(components, "\n"))
}
