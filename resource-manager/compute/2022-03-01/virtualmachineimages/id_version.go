package virtualmachineimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VersionId{}

// VersionId is a struct representing the Resource ID for a Version
type VersionId struct {
	SubscriptionId string
	Location       string
	PublisherName  string
	TypeName       string
	Version        string
}

// NewVersionID returns a new VersionId struct
func NewVersionID(subscriptionId string, location string, publisherName string, typeName string, version string) VersionId {
	return VersionId{
		SubscriptionId: subscriptionId,
		Location:       location,
		PublisherName:  publisherName,
		TypeName:       typeName,
		Version:        version,
	}
}

// ParseVersionID parses 'input' into a VersionId
func ParseVersionID(input string) (*VersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.TypeName, ok = parsed.Parsed["typeName"]; !ok {
		return nil, fmt.Errorf("the segment 'typeName' was not found in the resource id %q", input)
	}

	if id.Version, ok = parsed.Parsed["version"]; !ok {
		return nil, fmt.Errorf("the segment 'version' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVersionIDInsensitively parses 'input' case-insensitively into a VersionId
// note: this method should only be used for API response data and not user input
func ParseVersionIDInsensitively(input string) (*VersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	if id.TypeName, ok = parsed.Parsed["typeName"]; !ok {
		return nil, fmt.Errorf("the segment 'typeName' was not found in the resource id %q", input)
	}

	if id.Version, ok = parsed.Parsed["version"]; !ok {
		return nil, fmt.Errorf("the segment 'version' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVersionID checks that 'input' can be parsed as a Version ID
func ValidateVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Version ID
func (id VersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/publishers/%s/artifactTypes/vmExtension/types/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.PublisherName, id.TypeName, id.Version)
}

// Segments returns a slice of Resource ID Segments which comprise this Version ID
func (id VersionId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticVmExtension", "vmExtension", "vmExtension"),
		resourceids.StaticSegment("staticTypes", "types", "types"),
		resourceids.UserSpecifiedSegment("typeName", "typeValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("version", "versionValue"),
	}
}

// String returns a human-readable description of this Version ID
func (id VersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
		fmt.Sprintf("Type Name: %q", id.TypeName),
		fmt.Sprintf("Version: %q", id.Version),
	}
	return fmt.Sprintf("Version (%s)", strings.Join(components, "\n"))
}
