package virtualmachineimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = OfferSkuId{}

// OfferSkuId is a struct representing the Resource ID for a Offer Sku
type OfferSkuId struct {
	SubscriptionId string
	Location       string
	EdgeZone       string
	PublisherName  string
	Offer          string
	Skus           string
}

// NewOfferSkuID returns a new OfferSkuId struct
func NewOfferSkuID(subscriptionId string, location string, edgeZone string, publisherName string, offer string, skus string) OfferSkuId {
	return OfferSkuId{
		SubscriptionId: subscriptionId,
		Location:       location,
		EdgeZone:       edgeZone,
		PublisherName:  publisherName,
		Offer:          offer,
		Skus:           skus,
	}
}

// ParseOfferSkuID parses 'input' into a OfferSkuId
func ParseOfferSkuID(input string) (*OfferSkuId, error) {
	parser := resourceids.NewParserFromResourceIdType(OfferSkuId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OfferSkuId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.EdgeZone, ok = parsed.Parsed["edgeZone"]; !ok {
		return nil, fmt.Errorf("the segment 'edgeZone' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.Offer, ok = parsed.Parsed["offer"]; !ok {
		return nil, fmt.Errorf("the segment 'offer' was not found in the resource id %q", input)
	}

	if id.Skus, ok = parsed.Parsed["skus"]; !ok {
		return nil, fmt.Errorf("the segment 'skus' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseOfferSkuIDInsensitively parses 'input' case-insensitively into a OfferSkuId
// note: this method should only be used for API response data and not user input
func ParseOfferSkuIDInsensitively(input string) (*OfferSkuId, error) {
	parser := resourceids.NewParserFromResourceIdType(OfferSkuId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OfferSkuId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.EdgeZone, ok = parsed.Parsed["edgeZone"]; !ok {
		return nil, fmt.Errorf("the segment 'edgeZone' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.Offer, ok = parsed.Parsed["offer"]; !ok {
		return nil, fmt.Errorf("the segment 'offer' was not found in the resource id %q", input)
	}

	if id.Skus, ok = parsed.Parsed["skus"]; !ok {
		return nil, fmt.Errorf("the segment 'skus' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateOfferSkuID checks that 'input' can be parsed as a Offer Sku ID
func ValidateOfferSkuID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOfferSkuID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Offer Sku ID
func (id OfferSkuId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/edgeZones/%s/publishers/%s/artifactTypes/vmImage/offers/%s/skus/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.EdgeZone, id.PublisherName, id.Offer, id.Skus)
}

// Segments returns a slice of Resource ID Segments which comprise this Offer Sku ID
func (id OfferSkuId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticEdgeZones", "edgeZones", "edgeZones"),
		resourceids.UserSpecifiedSegment("edgeZone", "edgeZoneValue"),
		resourceids.StaticSegment("staticPublishers", "publishers", "publishers"),
		resourceids.UserSpecifiedSegment("publisherName", "publisherValue"),
		resourceids.StaticSegment("staticArtifactTypes", "artifactTypes", "artifactTypes"),
		resourceids.StaticSegment("staticVmImage", "vmImage", "vmImage"),
		resourceids.StaticSegment("staticOffers", "offers", "offers"),
		resourceids.UserSpecifiedSegment("offer", "offerValue"),
		resourceids.StaticSegment("staticSkus", "skus", "skus"),
		resourceids.UserSpecifiedSegment("skus", "skusValue"),
	}
}

// String returns a human-readable description of this Offer Sku ID
func (id OfferSkuId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Edge Zone: %q", id.EdgeZone),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
		fmt.Sprintf("Offer: %q", id.Offer),
		fmt.Sprintf("Skus: %q", id.Skus),
	}
	return fmt.Sprintf("Offer Sku (%s)", strings.Join(components, "\n"))
}
