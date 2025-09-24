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
	recaser.RegisterResourceId(&GroupQuotaResourceProviderId{})
}

var _ resourceids.ResourceId = &GroupQuotaResourceProviderId{}

// GroupQuotaResourceProviderId is a struct representing the Resource ID for a Group Quota Resource Provider
type GroupQuotaResourceProviderId struct {
	ManagementGroupId    string
	SubscriptionId       string
	GroupQuotaName       string
	ResourceProviderName string
}

// NewGroupQuotaResourceProviderID returns a new GroupQuotaResourceProviderId struct
func NewGroupQuotaResourceProviderID(managementGroupId string, subscriptionId string, groupQuotaName string, resourceProviderName string) GroupQuotaResourceProviderId {
	return GroupQuotaResourceProviderId{
		ManagementGroupId:    managementGroupId,
		SubscriptionId:       subscriptionId,
		GroupQuotaName:       groupQuotaName,
		ResourceProviderName: resourceProviderName,
	}
}

// ParseGroupQuotaResourceProviderID parses 'input' into a GroupQuotaResourceProviderId
func ParseGroupQuotaResourceProviderID(input string) (*GroupQuotaResourceProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupQuotaResourceProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupQuotaResourceProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupQuotaResourceProviderIDInsensitively parses 'input' case-insensitively into a GroupQuotaResourceProviderId
// note: this method should only be used for API response data and not user input
func ParseGroupQuotaResourceProviderIDInsensitively(input string) (*GroupQuotaResourceProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupQuotaResourceProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupQuotaResourceProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupQuotaResourceProviderId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupQuotaResourceProviderID checks that 'input' can be parsed as a Group Quota Resource Provider ID
func ValidateGroupQuotaResourceProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupQuotaResourceProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Quota Resource Provider ID
func (id GroupQuotaResourceProviderId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/subscriptions/%s/providers/Microsoft.Quota/groupQuotas/%s/resourceProviders/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.SubscriptionId, id.GroupQuotaName, id.ResourceProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Quota Resource Provider ID
func (id GroupQuotaResourceProviderId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Group Quota Resource Provider ID
func (id GroupQuotaResourceProviderId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Group Quota Name: %q", id.GroupQuotaName),
		fmt.Sprintf("Resource Provider Name: %q", id.ResourceProviderName),
	}
	return fmt.Sprintf("Group Quota Resource Provider (%s)", strings.Join(components, "\n"))
}
