package topquerystatistics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TopQueryStatisticId{})
}

var _ resourceids.ResourceId = &TopQueryStatisticId{}

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
	parser := resourceids.NewParserFromResourceIdType(&TopQueryStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TopQueryStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTopQueryStatisticIDInsensitively parses 'input' case-insensitively into a TopQueryStatisticId
// note: this method should only be used for API response data and not user input
func ParseTopQueryStatisticIDInsensitively(input string) (*TopQueryStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TopQueryStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TopQueryStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TopQueryStatisticId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServerName, ok = input.Parsed["serverName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverName", input)
	}

	if id.QueryStatisticId, ok = input.Parsed["queryStatisticId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "queryStatisticId", input)
	}

	return nil
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforMariaDB/servers/%s/topQueryStatistics/%s"
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
		resourceids.ResourceProviderSegment("staticMicrosoftDBforMariaDB", "Microsoft.DBforMariaDB", "Microsoft.DBforMariaDB"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverName"),
		resourceids.StaticSegment("staticTopQueryStatistics", "topQueryStatistics", "topQueryStatistics"),
		resourceids.UserSpecifiedSegment("queryStatisticId", "queryStatisticId"),
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
