package jobtargetgroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TargetGroupId{}

// TargetGroupId is a struct representing the Resource ID for a Target Group
type TargetGroupId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	JobAgentName      string
	TargetGroupName   string
}

// NewTargetGroupID returns a new TargetGroupId struct
func NewTargetGroupID(subscriptionId string, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string) TargetGroupId {
	return TargetGroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		JobAgentName:      jobAgentName,
		TargetGroupName:   targetGroupName,
	}
}

// ParseTargetGroupID parses 'input' into a TargetGroupId
func ParseTargetGroupID(input string) (*TargetGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(TargetGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TargetGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.JobAgentName, ok = parsed.Parsed["jobAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobAgentName", *parsed)
	}

	if id.TargetGroupName, ok = parsed.Parsed["targetGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "targetGroupName", *parsed)
	}

	return &id, nil
}

// ParseTargetGroupIDInsensitively parses 'input' case-insensitively into a TargetGroupId
// note: this method should only be used for API response data and not user input
func ParseTargetGroupIDInsensitively(input string) (*TargetGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(TargetGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TargetGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.JobAgentName, ok = parsed.Parsed["jobAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobAgentName", *parsed)
	}

	if id.TargetGroupName, ok = parsed.Parsed["targetGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "targetGroupName", *parsed)
	}

	return &id, nil
}

// ValidateTargetGroupID checks that 'input' can be parsed as a Target Group ID
func ValidateTargetGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTargetGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Target Group ID
func (id TargetGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/jobAgents/%s/targetGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.JobAgentName, id.TargetGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Target Group ID
func (id TargetGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticJobAgents", "jobAgents", "jobAgents"),
		resourceids.UserSpecifiedSegment("jobAgentName", "jobAgentValue"),
		resourceids.StaticSegment("staticTargetGroups", "targetGroups", "targetGroups"),
		resourceids.UserSpecifiedSegment("targetGroupName", "targetGroupValue"),
	}
}

// String returns a human-readable description of this Target Group ID
func (id TargetGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Job Agent Name: %q", id.JobAgentName),
		fmt.Sprintf("Target Group Name: %q", id.TargetGroupName),
	}
	return fmt.Sprintf("Target Group (%s)", strings.Join(components, "\n"))
}
