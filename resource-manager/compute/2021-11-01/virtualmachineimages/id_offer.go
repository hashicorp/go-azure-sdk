package virtualmachineimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = OfferId{}

// OfferId is a struct representing the Resource ID for a Offer
type OfferId struct {
	SubscriptionId string
	LocationName   string
	PublisherName  string
	OfferName      string
}

// NewOfferID returns a new OfferId struct
func NewOfferID(subscriptionId string, locationName string, publisherName string, offerName string) OfferId {
	return OfferId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		PublisherName:  publisherName,
		OfferName:      offerName,
	}
}

// ParseOfferID parses 'input' into a OfferId
func ParseOfferID(input string) (*OfferId, error) {
	parser := resourceids.NewParserFromResourceIdType(OfferId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OfferId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "publisherName", *parsed)
	}

	if id.OfferName, ok = parsed.Parsed["offerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerName", *parsed)
	}

	return &id, nil
}

// ParseOfferIDInsensitively parses 'input' case-insensitively into a OfferId
// note: this method should only be used for API response data and not user input
func ParseOfferIDInsensitively(input string) (*OfferId, error) {
	parser := resourceids.NewParserFromResourceIdType(OfferId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OfferId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "publisherName", *parsed)
	}

	if id.OfferName, ok = parsed.Parsed["offerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerName", *parsed)
	}

	return &id, nil
}

// ValidateOfferID checks that 'input' can be parsed as a Offer ID
func ValidateOfferID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOfferID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Offer ID
func (id OfferId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/publishers/%s/artifactTypes/vmImage/offers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.PublisherName, id.OfferName)
}

// Segments returns a slice of Resource ID Segments which comprise this Offer ID
func (id OfferId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticPublishers", "publishers", "publishers"),
		resourceids.UserSpecifiedSegment("publisherName", "publisherValue"),
		resourceids.StaticSegment("staticArtifactTypes", "artifactTypes", "artifactTypes"),
		resourceids.StaticSegment("staticVmImage", "vmImage", "vmImage"),
		resourceids.StaticSegment("staticOffers", "offers", "offers"),
		resourceids.UserSpecifiedSegment("offerName", "offerValue"),
	}
}

// String returns a human-readable description of this Offer ID
func (id OfferId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
		fmt.Sprintf("Offer Name: %q", id.OfferName),
	}
	return fmt.Sprintf("Offer (%s)", strings.Join(components, "\n"))
}
