package syncagents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SyncAgentId{}

// SyncAgentId is a struct representing the Resource ID for a Sync Agent
type SyncAgentId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	SyncAgentName     string
}

// NewSyncAgentID returns a new SyncAgentId struct
func NewSyncAgentID(subscriptionId string, resourceGroupName string, serverName string, syncAgentName string) SyncAgentId {
	return SyncAgentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		SyncAgentName:     syncAgentName,
	}
}

// ParseSyncAgentID parses 'input' into a SyncAgentId
func ParseSyncAgentID(input string) (*SyncAgentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SyncAgentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SyncAgentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSyncAgentIDInsensitively parses 'input' case-insensitively into a SyncAgentId
// note: this method should only be used for API response data and not user input
func ParseSyncAgentIDInsensitively(input string) (*SyncAgentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SyncAgentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SyncAgentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SyncAgentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SyncAgentName, ok = input.Parsed["syncAgentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "syncAgentName", input)
	}

	return nil
}

// ValidateSyncAgentID checks that 'input' can be parsed as a Sync Agent ID
func ValidateSyncAgentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSyncAgentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sync Agent ID
func (id SyncAgentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/syncAgents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.SyncAgentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sync Agent ID
func (id SyncAgentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticSyncAgents", "syncAgents", "syncAgents"),
		resourceids.UserSpecifiedSegment("syncAgentName", "syncAgentValue"),
	}
}

// String returns a human-readable description of this Sync Agent ID
func (id SyncAgentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Sync Agent Name: %q", id.SyncAgentName),
	}
	return fmt.Sprintf("Sync Agent (%s)", strings.Join(components, "\n"))
}
