package recommendedactions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RecommendedActionId{}

// RecommendedActionId is a struct representing the Resource ID for a Recommended Action
type RecommendedActionId struct {
	SubscriptionId        string
	ResourceGroupName     string
	ServerName            string
	AdvisorName           string
	RecommendedActionName string
}

// NewRecommendedActionID returns a new RecommendedActionId struct
func NewRecommendedActionID(subscriptionId string, resourceGroupName string, serverName string, advisorName string, recommendedActionName string) RecommendedActionId {
	return RecommendedActionId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		ServerName:            serverName,
		AdvisorName:           advisorName,
		RecommendedActionName: recommendedActionName,
	}
}

// ParseRecommendedActionID parses 'input' into a RecommendedActionId
func ParseRecommendedActionID(input string) (*RecommendedActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RecommendedActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RecommendedActionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRecommendedActionIDInsensitively parses 'input' case-insensitively into a RecommendedActionId
// note: this method should only be used for API response data and not user input
func ParseRecommendedActionIDInsensitively(input string) (*RecommendedActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RecommendedActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RecommendedActionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RecommendedActionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AdvisorName, ok = input.Parsed["advisorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "advisorName", input)
	}

	if id.RecommendedActionName, ok = input.Parsed["recommendedActionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recommendedActionName", input)
	}

	return nil
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/servers/%s/advisors/%s/recommendedActions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.AdvisorName, id.RecommendedActionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Recommended Action ID
func (id RecommendedActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
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
		fmt.Sprintf("Advisor Name: %q", id.AdvisorName),
		fmt.Sprintf("Recommended Action Name: %q", id.RecommendedActionName),
	}
	return fmt.Sprintf("Recommended Action (%s)", strings.Join(components, "\n"))
}
