package syncmembers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SyncMemberId{}

// SyncMemberId is a struct representing the Resource ID for a Sync Member
type SyncMemberId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	DatabaseName      string
	SyncGroupName     string
	SyncMemberName    string
}

// NewSyncMemberID returns a new SyncMemberId struct
func NewSyncMemberID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) SyncMemberId {
	return SyncMemberId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		DatabaseName:      databaseName,
		SyncGroupName:     syncGroupName,
		SyncMemberName:    syncMemberName,
	}
}

// ParseSyncMemberID parses 'input' into a SyncMemberId
func ParseSyncMemberID(input string) (*SyncMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(SyncMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SyncMemberId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSyncMemberIDInsensitively parses 'input' case-insensitively into a SyncMemberId
// note: this method should only be used for API response data and not user input
func ParseSyncMemberIDInsensitively(input string) (*SyncMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(SyncMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SyncMemberId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SyncMemberId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DatabaseName, ok = input.Parsed["databaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseName", input)
	}

	if id.SyncGroupName, ok = input.Parsed["syncGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "syncGroupName", input)
	}

	if id.SyncMemberName, ok = input.Parsed["syncMemberName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "syncMemberName", input)
	}

	return nil
}

// ValidateSyncMemberID checks that 'input' can be parsed as a Sync Member ID
func ValidateSyncMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSyncMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sync Member ID
func (id SyncMemberId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/syncGroups/%s/syncMembers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.SyncGroupName, id.SyncMemberName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sync Member ID
func (id SyncMemberId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticSyncGroups", "syncGroups", "syncGroups"),
		resourceids.UserSpecifiedSegment("syncGroupName", "syncGroupValue"),
		resourceids.StaticSegment("staticSyncMembers", "syncMembers", "syncMembers"),
		resourceids.UserSpecifiedSegment("syncMemberName", "syncMemberValue"),
	}
}

// String returns a human-readable description of this Sync Member ID
func (id SyncMemberId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Sync Group Name: %q", id.SyncGroupName),
		fmt.Sprintf("Sync Member Name: %q", id.SyncMemberName),
	}
	return fmt.Sprintf("Sync Member (%s)", strings.Join(components, "\n"))
}
