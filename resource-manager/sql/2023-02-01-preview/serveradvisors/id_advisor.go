package serveradvisors

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdvisorId{}

// AdvisorId is a struct representing the Resource ID for a Advisor
type AdvisorId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	AdvisorName       string
}

// NewAdvisorID returns a new AdvisorId struct
func NewAdvisorID(subscriptionId string, resourceGroupName string, serverName string, advisorName string) AdvisorId {
	return AdvisorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		AdvisorName:       advisorName,
	}
}

// ParseAdvisorID parses 'input' into a AdvisorId
func ParseAdvisorID(input string) (*AdvisorId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdvisorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdvisorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdvisorIDInsensitively parses 'input' case-insensitively into a AdvisorId
// note: this method should only be used for API response data and not user input
func ParseAdvisorIDInsensitively(input string) (*AdvisorId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdvisorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdvisorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdvisorId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateAdvisorID checks that 'input' can be parsed as a Advisor ID
func ValidateAdvisorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdvisorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Advisor ID
func (id AdvisorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/advisors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.AdvisorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Advisor ID
func (id AdvisorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticAdvisors", "advisors", "advisors"),
		resourceids.UserSpecifiedSegment("advisorName", "advisorValue"),
	}
}

// String returns a human-readable description of this Advisor ID
func (id AdvisorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Advisor Name: %q", id.AdvisorName),
	}
	return fmt.Sprintf("Advisor (%s)", strings.Join(components, "\n"))
}
