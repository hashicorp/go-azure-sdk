package recommendations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SiteRecommendationId{}

// SiteRecommendationId is a struct representing the Resource ID for a Site Recommendation
type SiteRecommendationId struct {
	SubscriptionId     string
	ResourceGroupName  string
	SiteName           string
	RecommendationName string
}

// NewSiteRecommendationID returns a new SiteRecommendationId struct
func NewSiteRecommendationID(subscriptionId string, resourceGroupName string, siteName string, recommendationName string) SiteRecommendationId {
	return SiteRecommendationId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		SiteName:           siteName,
		RecommendationName: recommendationName,
	}
}

// ParseSiteRecommendationID parses 'input' into a SiteRecommendationId
func ParseSiteRecommendationID(input string) (*SiteRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(SiteRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SiteRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.RecommendationName, ok = parsed.Parsed["recommendationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationName", *parsed)
	}

	return &id, nil
}

// ParseSiteRecommendationIDInsensitively parses 'input' case-insensitively into a SiteRecommendationId
// note: this method should only be used for API response data and not user input
func ParseSiteRecommendationIDInsensitively(input string) (*SiteRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(SiteRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SiteRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.RecommendationName, ok = parsed.Parsed["recommendationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationName", *parsed)
	}

	return &id, nil
}

// ValidateSiteRecommendationID checks that 'input' can be parsed as a Site Recommendation ID
func ValidateSiteRecommendationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSiteRecommendationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Site Recommendation ID
func (id SiteRecommendationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/recommendations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.RecommendationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Site Recommendation ID
func (id SiteRecommendationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticRecommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationName", "recommendationValue"),
	}
}

// String returns a human-readable description of this Site Recommendation ID
func (id SiteRecommendationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Recommendation Name: %q", id.RecommendationName),
	}
	return fmt.Sprintf("Site Recommendation (%s)", strings.Join(components, "\n"))
}
