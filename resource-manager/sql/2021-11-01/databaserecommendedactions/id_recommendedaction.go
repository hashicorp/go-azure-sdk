package databaserecommendedactions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RecommendedActionId{}

// RecommendedActionId is a struct representing the Resource ID for a Recommended Action
type RecommendedActionId struct {
	SubscriptionId        string
	ResourceGroupName     string
	ServerName            string
	DatabaseName          string
	AdvisorName           string
	RecommendedActionName string
}

// NewRecommendedActionID returns a new RecommendedActionId struct
func NewRecommendedActionID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string) RecommendedActionId {
	return RecommendedActionId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		ServerName:            serverName,
		DatabaseName:          databaseName,
		AdvisorName:           advisorName,
		RecommendedActionName: recommendedActionName,
	}
}

// ParseRecommendedActionID parses 'input' into a RecommendedActionId
func ParseRecommendedActionID(input string) (*RecommendedActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecommendedActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecommendedActionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.AdvisorName, ok = parsed.Parsed["advisorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "advisorName", *parsed)
	}

	if id.RecommendedActionName, ok = parsed.Parsed["recommendedActionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendedActionName", *parsed)
	}

	return &id, nil
}

// ParseRecommendedActionIDInsensitively parses 'input' case-insensitively into a RecommendedActionId
// note: this method should only be used for API response data and not user input
func ParseRecommendedActionIDInsensitively(input string) (*RecommendedActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecommendedActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecommendedActionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.AdvisorName, ok = parsed.Parsed["advisorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "advisorName", *parsed)
	}

	if id.RecommendedActionName, ok = parsed.Parsed["recommendedActionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendedActionName", *parsed)
	}

	return &id, nil
}

// ValidateRecommendedActionID checks that 'input' can be parsed as a Recommended Action ID
func ValidateRecommendedActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecommendedActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recommended Action ID
func (id RecommendedActionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/advisors/%s/recommendedActions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.AdvisorName, id.RecommendedActionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Recommended Action ID
func (id RecommendedActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticAdvisors", "advisors", "advisors"),
		resourceids.UserSpecifiedSegment("advisorName", "advisorValue"),
		resourceids.StaticSegment("staticRecommendedActions", "recommendedActions", "recommendedActions"),
		resourceids.UserSpecifiedSegment("recommendedActionName", "recommendedActionValue"),
	}
}

// String returns a human-readable description of this Recommended Action ID
func (id RecommendedActionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Advisor Name: %q", id.AdvisorName),
		fmt.Sprintf("Recommended Action Name: %q", id.RecommendedActionName),
	}
	return fmt.Sprintf("Recommended Action (%s)", strings.Join(components, "\n"))
}
