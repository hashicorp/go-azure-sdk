package authenticationsettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AuthenticationsettingId{})
}

var _ resourceids.ResourceId = &AuthenticationsettingId{}

// AuthenticationsettingId is a struct representing the Resource ID for a Authenticationsetting
type AuthenticationsettingId struct {
	SubscriptionId            string
	ResourceGroupName         string
	HealthmodelName           string
	AuthenticationsettingName string
}

// NewAuthenticationsettingID returns a new AuthenticationsettingId struct
func NewAuthenticationsettingID(subscriptionId string, resourceGroupName string, healthmodelName string, authenticationsettingName string) AuthenticationsettingId {
	return AuthenticationsettingId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		HealthmodelName:           healthmodelName,
		AuthenticationsettingName: authenticationsettingName,
	}
}

// ParseAuthenticationsettingID parses 'input' into a AuthenticationsettingId
func ParseAuthenticationsettingID(input string) (*AuthenticationsettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthenticationsettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthenticationsettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuthenticationsettingIDInsensitively parses 'input' case-insensitively into a AuthenticationsettingId
// note: this method should only be used for API response data and not user input
func ParseAuthenticationsettingIDInsensitively(input string) (*AuthenticationsettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthenticationsettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthenticationsettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuthenticationsettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.HealthmodelName, ok = input.Parsed["healthmodelName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "healthmodelName", input)
	}

	if id.AuthenticationsettingName, ok = input.Parsed["authenticationsettingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationsettingName", input)
	}

	return nil
}

// ValidateAuthenticationsettingID checks that 'input' can be parsed as a Authenticationsetting ID
func ValidateAuthenticationsettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuthenticationsettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Authenticationsetting ID
func (id AuthenticationsettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CloudHealth/healthmodels/%s/authenticationsettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HealthmodelName, id.AuthenticationsettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Authenticationsetting ID
func (id AuthenticationsettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCloudHealth", "Microsoft.CloudHealth", "Microsoft.CloudHealth"),
		resourceids.StaticSegment("staticHealthmodels", "healthmodels", "healthmodels"),
		resourceids.UserSpecifiedSegment("healthmodelName", "healthmodelName"),
		resourceids.StaticSegment("staticAuthenticationsettings", "authenticationsettings", "authenticationsettings"),
		resourceids.UserSpecifiedSegment("authenticationsettingName", "authenticationsettingName"),
	}
}

// String returns a human-readable description of this Authenticationsetting ID
func (id AuthenticationsettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Healthmodel Name: %q", id.HealthmodelName),
		fmt.Sprintf("Authenticationsetting Name: %q", id.AuthenticationsettingName),
	}
	return fmt.Sprintf("Authenticationsetting (%s)", strings.Join(components, "\n"))
}
