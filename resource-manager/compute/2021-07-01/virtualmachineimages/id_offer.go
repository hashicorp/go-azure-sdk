package virtualmachineimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = OfferId{}

// OfferId is a struct representing the Resource ID for a Offer
type OfferId struct {
	SubscriptionId string
	Location       string
	PublisherName  string
	Offer          string
}

// NewOfferID returns a new OfferId struct
func NewOfferID(subscriptionId string, location string, publisherName string, offer string) OfferId {
	return OfferId{
		SubscriptionId: subscriptionId,
		Location:       location,
		PublisherName:  publisherName,
		Offer:          offer,
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
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.Offer, ok = parsed.Parsed["offer"]; !ok {
		return nil, fmt.Errorf("the segment 'offer' was not found in the resource id %q", input)
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
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.Offer, ok = parsed.Parsed["offer"]; !ok {
		return nil, fmt.Errorf("the segment 'offer' was not found in the resource id %q", input)
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
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.PublisherName, id.Offer)
}

// Segments returns a slice of Resource ID Segments which comprise this Offer ID
func (id OfferId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticPublishers", "publishers", "publishers"),
		resourceids.UserSpecifiedSegment("publisherName", "publisherValue"),
		resourceids.StaticSegment("staticArtifactTypes", "artifactTypes", "artifactTypes"),
		resourceids.StaticSegment("staticVmImage", "vmImage", "vmImage"),
		resourceids.StaticSegment("staticOffers", "offers", "offers"),
		resourceids.UserSpecifiedSegment("offer", "offerValue"),
	}
}

// String returns a human-readable description of this Offer ID
func (id OfferId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
		fmt.Sprintf("Offer: %q", id.Offer),
	}
	return fmt.Sprintf("Offer (%s)", strings.Join(components, "\n"))
}
