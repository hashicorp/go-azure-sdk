// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiissueattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = IssueId{}

// IssueId is a struct representing the Resource ID for a Issue
type IssueId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	IssueId           string
}

// NewIssueID returns a new IssueId struct
func NewIssueID(subscriptionId string, resourceGroupName string, serviceName string, issueId string) IssueId {
	return IssueId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		IssueId:           issueId,
	}
}

// ParseIssueID parses 'input' into a IssueId
func ParseIssueID(input string) (*IssueId, error) {
	parser := resourceids.NewParserFromResourceIdType(IssueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IssueId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.IssueId, ok = parsed.Parsed["issueId"]; !ok {
		return nil, fmt.Errorf("the segment 'issueId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseIssueIDInsensitively parses 'input' case-insensitively into a IssueId
// note: this method should only be used for API response data and not user input
func ParseIssueIDInsensitively(input string) (*IssueId, error) {
	parser := resourceids.NewParserFromResourceIdType(IssueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IssueId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.IssueId, ok = parsed.Parsed["issueId"]; !ok {
		return nil, fmt.Errorf("the segment 'issueId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateIssueID checks that 'input' can be parsed as a Issue ID
func ValidateIssueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIssueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Issue ID
func (id IssueId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/issues/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.IssueId)
}

// Segments returns a slice of Resource ID Segments which comprise this Issue ID
func (id IssueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticIssues", "issues", "issues"),
		resourceids.UserSpecifiedSegment("issueId", "issueIdValue"),
	}
}

// String returns a human-readable description of this Issue ID
func (id IssueId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Issue: %q", id.IssueId),
	}
	return fmt.Sprintf("Issue (%s)", strings.Join(components, "\n"))
}
