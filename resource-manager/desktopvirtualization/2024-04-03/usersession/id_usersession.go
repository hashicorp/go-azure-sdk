package usersession

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&UserSessionId{})
}

var _ resourceids.ResourceId = &UserSessionId{}

// UserSessionId is a struct representing the Resource ID for a User Session
type UserSessionId struct {
	SubscriptionId    string
	ResourceGroupName string
	HostPoolName      string
	SessionHostName   string
	UserSessionId     string
}

// NewUserSessionID returns a new UserSessionId struct
func NewUserSessionID(subscriptionId string, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionId string) UserSessionId {
	return UserSessionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		HostPoolName:      hostPoolName,
		SessionHostName:   sessionHostName,
		UserSessionId:     userSessionId,
	}
}

// ParseUserSessionID parses 'input' into a UserSessionId
func ParseUserSessionID(input string) (*UserSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserSessionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserSessionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserSessionIDInsensitively parses 'input' case-insensitively into a UserSessionId
// note: this method should only be used for API response data and not user input
func ParseUserSessionIDInsensitively(input string) (*UserSessionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserSessionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserSessionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserSessionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.HostPoolName, ok = input.Parsed["hostPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hostPoolName", input)
	}

	if id.SessionHostName, ok = input.Parsed["sessionHostName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sessionHostName", input)
	}

	if id.UserSessionId, ok = input.Parsed["userSessionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userSessionId", input)
	}

	return nil
}

// ValidateUserSessionID checks that 'input' can be parsed as a User Session ID
func ValidateUserSessionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserSessionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Session ID
func (id UserSessionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/hostPools/%s/sessionHosts/%s/userSessions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HostPoolName, id.SessionHostName, id.UserSessionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Session ID
func (id UserSessionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDesktopVirtualization", "Microsoft.DesktopVirtualization", "Microsoft.DesktopVirtualization"),
		resourceids.StaticSegment("staticHostPools", "hostPools", "hostPools"),
		resourceids.UserSpecifiedSegment("hostPoolName", "hostPoolName"),
		resourceids.StaticSegment("staticSessionHosts", "sessionHosts", "sessionHosts"),
		resourceids.UserSpecifiedSegment("sessionHostName", "sessionHostName"),
		resourceids.StaticSegment("staticUserSessions", "userSessions", "userSessions"),
		resourceids.UserSpecifiedSegment("userSessionId", "userSessionId"),
	}
}

// String returns a human-readable description of this User Session ID
func (id UserSessionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Host Pool Name: %q", id.HostPoolName),
		fmt.Sprintf("Session Host Name: %q", id.SessionHostName),
		fmt.Sprintf("User Session: %q", id.UserSessionId),
	}
	return fmt.Sprintf("User Session (%s)", strings.Join(components, "\n"))
}
