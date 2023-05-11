package topquerystatistics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TopQueryStatisticId{}

// TopQueryStatisticId is a struct representing the Resource ID for a Top Query Statistic
type TopQueryStatisticId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	QueryStatisticId  string
}

// NewTopQueryStatisticID returns a new TopQueryStatisticId struct
func NewTopQueryStatisticID(subscriptionId string, resourceGroupName string, serverName string, queryStatisticId string) TopQueryStatisticId {
	return TopQueryStatisticId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		QueryStatisticId:  queryStatisticId,
	}
}

// ParseTopQueryStatisticID parses 'input' into a TopQueryStatisticId
func ParseTopQueryStatisticID(input string) (*TopQueryStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(TopQueryStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TopQueryStatisticId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.QueryStatisticId, ok = parsed.Parsed["queryStatisticId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "queryStatisticId", *parsed)
	}

	return &id, nil
}

// ParseTopQueryStatisticIDInsensitively parses 'input' case-insensitively into a TopQueryStatisticId
// note: this method should only be used for API response data and not user input
func ParseTopQueryStatisticIDInsensitively(input string) (*TopQueryStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(TopQueryStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TopQueryStatisticId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.QueryStatisticId, ok = parsed.Parsed["queryStatisticId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "queryStatisticId", *parsed)
	}

	return &id, nil
}

// ValidateTopQueryStatisticID checks that 'input' can be parsed as a Top Query Statistic ID
func ValidateTopQueryStatisticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTopQueryStatisticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Top Query Statistic ID
func (id TopQueryStatisticId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/servers/%s/topQueryStatistics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.QueryStatisticId)
}

// Segments returns a slice of Resource ID Segments which comprise this Top Query Statistic ID
func (id TopQueryStatisticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticTopQueryStatistics", "topQueryStatistics", "topQueryStatistics"),
		resourceids.UserSpecifiedSegment("queryStatisticId", "queryStatisticIdValue"),
	}
}

// String returns a human-readable description of this Top Query Statistic ID
func (id TopQueryStatisticId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Query Statistic: %q", id.QueryStatisticId),
	}
	return fmt.Sprintf("Top Query Statistic (%s)", strings.Join(components, "\n"))
}
