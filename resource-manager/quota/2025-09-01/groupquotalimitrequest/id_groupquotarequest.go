package groupquotalimitrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&GroupQuotaRequestId{})
}

var _ resourceids.ResourceId = &GroupQuotaRequestId{}

// GroupQuotaRequestId is a struct representing the Resource ID for a Group Quota Request
type GroupQuotaRequestId struct {
	ManagementGroupId string
	GroupQuotaName    string
	RequestId         string
}

// NewGroupQuotaRequestID returns a new GroupQuotaRequestId struct
func NewGroupQuotaRequestID(managementGroupId string, groupQuotaName string, requestId string) GroupQuotaRequestId {
	return GroupQuotaRequestId{
		ManagementGroupId: managementGroupId,
		GroupQuotaName:    groupQuotaName,
		RequestId:         requestId,
	}
}

// ParseGroupQuotaRequestID parses 'input' into a GroupQuotaRequestId
func ParseGroupQuotaRequestID(input string) (*GroupQuotaRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupQuotaRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupQuotaRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupQuotaRequestIDInsensitively parses 'input' case-insensitively into a GroupQuotaRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupQuotaRequestIDInsensitively(input string) (*GroupQuotaRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupQuotaRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupQuotaRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupQuotaRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagementGroupId, ok = input.Parsed["managementGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementGroupId", input)
	}

	if id.GroupQuotaName, ok = input.Parsed["groupQuotaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupQuotaName", input)
	}

	if id.RequestId, ok = input.Parsed["requestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "requestId", input)
	}

	return nil
}

// ValidateGroupQuotaRequestID checks that 'input' can be parsed as a Group Quota Request ID
func ValidateGroupQuotaRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupQuotaRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Quota Request ID
func (id GroupQuotaRequestId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Quota/groupQuotas/%s/groupQuotaRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.GroupQuotaName, id.RequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Quota Request ID
func (id GroupQuotaRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupId"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftQuota", "Microsoft.Quota", "Microsoft.Quota"),
		resourceids.StaticSegment("staticGroupQuotas", "groupQuotas", "groupQuotas"),
		resourceids.UserSpecifiedSegment("groupQuotaName", "groupQuotaName"),
		resourceids.StaticSegment("staticGroupQuotaRequests", "groupQuotaRequests", "groupQuotaRequests"),
		resourceids.UserSpecifiedSegment("requestId", "requestId"),
	}
}

// String returns a human-readable description of this Group Quota Request ID
func (id GroupQuotaRequestId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Group Quota Name: %q", id.GroupQuotaName),
		fmt.Sprintf("Request: %q", id.RequestId),
	}
	return fmt.Sprintf("Group Quota Request (%s)", strings.Join(components, "\n"))
}
