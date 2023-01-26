package clusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ClusterVersionId{}

// ClusterVersionId is a struct representing the Resource ID for a Cluster Version
type ClusterVersionId struct {
	SubscriptionId     string
	LocationName       string
	ClusterVersionName string
}

// NewClusterVersionID returns a new ClusterVersionId struct
func NewClusterVersionID(subscriptionId string, locationName string, clusterVersionName string) ClusterVersionId {
	return ClusterVersionId{
		SubscriptionId:     subscriptionId,
		LocationName:       locationName,
		ClusterVersionName: clusterVersionName,
	}
}

// ParseClusterVersionID parses 'input' into a ClusterVersionId
func ParseClusterVersionID(input string) (*ClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ClusterVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ClusterVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.ClusterVersionName, ok = parsed.Parsed["clusterVersionName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterVersionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseClusterVersionIDInsensitively parses 'input' case-insensitively into a ClusterVersionId
// note: this method should only be used for API response data and not user input
func ParseClusterVersionIDInsensitively(input string) (*ClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ClusterVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ClusterVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.ClusterVersionName, ok = parsed.Parsed["clusterVersionName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterVersionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateClusterVersionID checks that 'input' can be parsed as a Cluster Version ID
func ValidateClusterVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseClusterVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cluster Version ID
func (id ClusterVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/clusterVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.ClusterVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cluster Version ID
func (id ClusterVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticClusterVersions", "clusterVersions", "clusterVersions"),
		resourceids.UserSpecifiedSegment("clusterVersionName", "clusterVersionValue"),
	}
}

// String returns a human-readable description of this Cluster Version ID
func (id ClusterVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Cluster Version Name: %q", id.ClusterVersionName),
	}
	return fmt.Sprintf("Cluster Version (%s)", strings.Join(components, "\n"))
}
