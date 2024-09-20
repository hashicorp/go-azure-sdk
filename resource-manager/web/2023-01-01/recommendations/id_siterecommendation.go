package recommendations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SiteRecommendationId{})
}

var _ resourceids.ResourceId = &SiteRecommendationId{}

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
	parser := resourceids.NewParserFromResourceIdType(&SiteRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SiteRecommendationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSiteRecommendationIDInsensitively parses 'input' case-insensitively into a SiteRecommendationId
// note: this method should only be used for API response data and not user input
func ParseSiteRecommendationIDInsensitively(input string) (*SiteRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SiteRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SiteRecommendationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SiteRecommendationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SiteName, ok = input.Parsed["siteName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteName", input)
	}

	if id.RecommendationName, ok = input.Parsed["recommendationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recommendationName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("siteName", "siteName"),
		resourceids.StaticSegment("staticRecommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationName", "name"),
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
