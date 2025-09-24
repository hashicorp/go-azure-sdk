package groupquotasubscriptionrequeststatuses

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SubscriptionRequestId{})
}

var _ resourceids.ResourceId = &SubscriptionRequestId{}

// SubscriptionRequestId is a struct representing the Resource ID for a Subscription Request
type SubscriptionRequestId struct {
	ManagementGroupId string
	GroupQuotaName    string
	RequestId         string
}

// NewSubscriptionRequestID returns a new SubscriptionRequestId struct
func NewSubscriptionRequestID(managementGroupId string, groupQuotaName string, requestId string) SubscriptionRequestId {
	return SubscriptionRequestId{
		ManagementGroupId: managementGroupId,
		GroupQuotaName:    groupQuotaName,
		RequestId:         requestId,
	}
}

// ParseSubscriptionRequestID parses 'input' into a SubscriptionRequestId
func ParseSubscriptionRequestID(input string) (*SubscriptionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubscriptionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubscriptionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSubscriptionRequestIDInsensitively parses 'input' case-insensitively into a SubscriptionRequestId
// note: this method should only be used for API response data and not user input
func ParseSubscriptionRequestIDInsensitively(input string) (*SubscriptionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubscriptionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubscriptionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SubscriptionRequestId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateSubscriptionRequestID checks that 'input' can be parsed as a Subscription Request ID
func ValidateSubscriptionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSubscriptionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Subscription Request ID
func (id SubscriptionRequestId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Quota/groupQuotas/%s/subscriptionRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.GroupQuotaName, id.RequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Subscription Request ID
func (id SubscriptionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupId"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftQuota", "Microsoft.Quota", "Microsoft.Quota"),
		resourceids.StaticSegment("staticGroupQuotas", "groupQuotas", "groupQuotas"),
		resourceids.UserSpecifiedSegment("groupQuotaName", "groupQuotaName"),
		resourceids.StaticSegment("staticSubscriptionRequests", "subscriptionRequests", "subscriptionRequests"),
		resourceids.UserSpecifiedSegment("requestId", "requestId"),
	}
}

// String returns a human-readable description of this Subscription Request ID
func (id SubscriptionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Group Quota Name: %q", id.GroupQuotaName),
		fmt.Sprintf("Request: %q", id.RequestId),
	}
	return fmt.Sprintf("Subscription Request (%s)", strings.Join(components, "\n"))
}
