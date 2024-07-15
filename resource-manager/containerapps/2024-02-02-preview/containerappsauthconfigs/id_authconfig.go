package containerappsauthconfigs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AuthConfigId{})
}

var _ resourceids.ResourceId = &AuthConfigId{}

// AuthConfigId is a struct representing the Resource ID for a Auth Config
type AuthConfigId struct {
	SubscriptionId    string
	ResourceGroupName string
	ContainerAppName  string
	AuthConfigName    string
}

// NewAuthConfigID returns a new AuthConfigId struct
func NewAuthConfigID(subscriptionId string, resourceGroupName string, containerAppName string, authConfigName string) AuthConfigId {
	return AuthConfigId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ContainerAppName:  containerAppName,
		AuthConfigName:    authConfigName,
	}
}

// ParseAuthConfigID parses 'input' into a AuthConfigId
func ParseAuthConfigID(input string) (*AuthConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthConfigId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthConfigId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuthConfigIDInsensitively parses 'input' case-insensitively into a AuthConfigId
// note: this method should only be used for API response data and not user input
func ParseAuthConfigIDInsensitively(input string) (*AuthConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthConfigId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthConfigId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuthConfigId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ContainerAppName, ok = input.Parsed["containerAppName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "containerAppName", input)
	}

	if id.AuthConfigName, ok = input.Parsed["authConfigName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authConfigName", input)
	}

	return nil
}

// ValidateAuthConfigID checks that 'input' can be parsed as a Auth Config ID
func ValidateAuthConfigID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuthConfigID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Auth Config ID
func (id AuthConfigId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/containerApps/%s/authConfigs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ContainerAppName, id.AuthConfigName)
}

// Segments returns a slice of Resource ID Segments which comprise this Auth Config ID
func (id AuthConfigId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticContainerApps", "containerApps", "containerApps"),
		resourceids.UserSpecifiedSegment("containerAppName", "containerAppValue"),
		resourceids.StaticSegment("staticAuthConfigs", "authConfigs", "authConfigs"),
		resourceids.UserSpecifiedSegment("authConfigName", "authConfigValue"),
	}
}

// String returns a human-readable description of this Auth Config ID
func (id AuthConfigId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Container App Name: %q", id.ContainerAppName),
		fmt.Sprintf("Auth Config Name: %q", id.AuthConfigName),
	}
	return fmt.Sprintf("Auth Config (%s)", strings.Join(components, "\n"))
}
