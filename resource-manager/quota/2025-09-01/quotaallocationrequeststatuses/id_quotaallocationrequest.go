package quotaallocationrequeststatuses

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&QuotaAllocationRequestId{})
}

var _ resourceids.ResourceId = &QuotaAllocationRequestId{}

// QuotaAllocationRequestId is a struct representing the Resource ID for a Quota Allocation Request
type QuotaAllocationRequestId struct {
	ManagementGroupId    string
	SubscriptionId       string
	GroupQuotaName       string
	ResourceProviderName string
	AllocationId         string
}

// NewQuotaAllocationRequestID returns a new QuotaAllocationRequestId struct
func NewQuotaAllocationRequestID(managementGroupId string, subscriptionId string, groupQuotaName string, resourceProviderName string, allocationId string) QuotaAllocationRequestId {
	return QuotaAllocationRequestId{
		ManagementGroupId:    managementGroupId,
		SubscriptionId:       subscriptionId,
		GroupQuotaName:       groupQuotaName,
		ResourceProviderName: resourceProviderName,
		AllocationId:         allocationId,
	}
}

// ParseQuotaAllocationRequestID parses 'input' into a QuotaAllocationRequestId
func ParseQuotaAllocationRequestID(input string) (*QuotaAllocationRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&QuotaAllocationRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := QuotaAllocationRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseQuotaAllocationRequestIDInsensitively parses 'input' case-insensitively into a QuotaAllocationRequestId
// note: this method should only be used for API response data and not user input
func ParseQuotaAllocationRequestIDInsensitively(input string) (*QuotaAllocationRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&QuotaAllocationRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := QuotaAllocationRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *QuotaAllocationRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagementGroupId, ok = input.Parsed["managementGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementGroupId", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.GroupQuotaName, ok = input.Parsed["groupQuotaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupQuotaName", input)
	}

	if id.ResourceProviderName, ok = input.Parsed["resourceProviderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceProviderName", input)
	}

	if id.AllocationId, ok = input.Parsed["allocationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "allocationId", input)
	}

	return nil
}

// ValidateQuotaAllocationRequestID checks that 'input' can be parsed as a Quota Allocation Request ID
func ValidateQuotaAllocationRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQuotaAllocationRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Quota Allocation Request ID
func (id QuotaAllocationRequestId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/subscriptions/%s/providers/Microsoft.Quota/groupQuotas/%s/resourceProviders/%s/quotaAllocationRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.SubscriptionId, id.GroupQuotaName, id.ResourceProviderName, id.AllocationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Quota Allocation Request ID
func (id QuotaAllocationRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupId"),
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftQuota", "Microsoft.Quota", "Microsoft.Quota"),
		resourceids.StaticSegment("staticGroupQuotas", "groupQuotas", "groupQuotas"),
		resourceids.UserSpecifiedSegment("groupQuotaName", "groupQuotaName"),
		resourceids.StaticSegment("staticResourceProviders", "resourceProviders", "resourceProviders"),
		resourceids.UserSpecifiedSegment("resourceProviderName", "resourceProviderName"),
		resourceids.StaticSegment("staticQuotaAllocationRequests", "quotaAllocationRequests", "quotaAllocationRequests"),
		resourceids.UserSpecifiedSegment("allocationId", "allocationId"),
	}
}

// String returns a human-readable description of this Quota Allocation Request ID
func (id QuotaAllocationRequestId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Group Quota Name: %q", id.GroupQuotaName),
		fmt.Sprintf("Resource Provider Name: %q", id.ResourceProviderName),
		fmt.Sprintf("Allocation: %q", id.AllocationId),
	}
	return fmt.Sprintf("Quota Allocation Request (%s)", strings.Join(components, "\n"))
}
