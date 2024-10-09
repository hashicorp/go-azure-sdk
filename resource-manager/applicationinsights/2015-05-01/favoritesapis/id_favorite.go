package favoritesapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&FavoriteId{})
}

var _ resourceids.ResourceId = &FavoriteId{}

// FavoriteId is a struct representing the Resource ID for a Favorite
type FavoriteId struct {
	SubscriptionId    string
	ResourceGroupName string
	ComponentName     string
	FavoriteId        string
}

// NewFavoriteID returns a new FavoriteId struct
func NewFavoriteID(subscriptionId string, resourceGroupName string, componentName string, favoriteId string) FavoriteId {
	return FavoriteId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ComponentName:     componentName,
		FavoriteId:        favoriteId,
	}
}

// ParseFavoriteID parses 'input' into a FavoriteId
func ParseFavoriteID(input string) (*FavoriteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FavoriteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FavoriteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFavoriteIDInsensitively parses 'input' case-insensitively into a FavoriteId
// note: this method should only be used for API response data and not user input
func ParseFavoriteIDInsensitively(input string) (*FavoriteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FavoriteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FavoriteId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FavoriteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ComponentName, ok = input.Parsed["componentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "componentName", input)
	}

	if id.FavoriteId, ok = input.Parsed["favoriteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "favoriteId", input)
	}

	return nil
}

// ValidateFavoriteID checks that 'input' can be parsed as a Favorite ID
func ValidateFavoriteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFavoriteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Favorite ID
func (id FavoriteId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/components/%s/favorites/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ComponentName, id.FavoriteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Favorite ID
func (id FavoriteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
		resourceids.StaticSegment("staticFavorites", "favorites", "favorites"),
		resourceids.UserSpecifiedSegment("favoriteId", "favoriteId"),
	}
}

// String returns a human-readable description of this Favorite ID
func (id FavoriteId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Favorite: %q", id.FavoriteId),
	}
	return fmt.Sprintf("Favorite (%s)", strings.Join(components, "\n"))
}
