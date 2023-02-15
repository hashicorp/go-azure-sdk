package contenttypecontentitem

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ContentItemId{}

// ContentItemId is a struct representing the Resource ID for a Content Item
type ContentItemId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ContentTypeId     string
	ContentItemId     string
}

// NewContentItemID returns a new ContentItemId struct
func NewContentItemID(subscriptionId string, resourceGroupName string, serviceName string, contentTypeId string, contentItemId string) ContentItemId {
	return ContentItemId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ContentTypeId:     contentTypeId,
		ContentItemId:     contentItemId,
	}
}

// ParseContentItemID parses 'input' into a ContentItemId
func ParseContentItemID(input string) (*ContentItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(ContentItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ContentItemId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, fmt.Errorf("the segment 'contentTypeId' was not found in the resource id %q", input)
	}

	if id.ContentItemId, ok = parsed.Parsed["contentItemId"]; !ok {
		return nil, fmt.Errorf("the segment 'contentItemId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseContentItemIDInsensitively parses 'input' case-insensitively into a ContentItemId
// note: this method should only be used for API response data and not user input
func ParseContentItemIDInsensitively(input string) (*ContentItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(ContentItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ContentItemId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ContentTypeId, ok = parsed.Parsed["contentTypeId"]; !ok {
		return nil, fmt.Errorf("the segment 'contentTypeId' was not found in the resource id %q", input)
	}

	if id.ContentItemId, ok = parsed.Parsed["contentItemId"]; !ok {
		return nil, fmt.Errorf("the segment 'contentItemId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateContentItemID checks that 'input' can be parsed as a Content Item ID
func ValidateContentItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContentItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Content Item ID
func (id ContentItemId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/contentTypes/%s/contentItems/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ContentTypeId, id.ContentItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Content Item ID
func (id ContentItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticContentTypes", "contentTypes", "contentTypes"),
		resourceids.UserSpecifiedSegment("contentTypeId", "contentTypeIdValue"),
		resourceids.StaticSegment("staticContentItems", "contentItems", "contentItems"),
		resourceids.UserSpecifiedSegment("contentItemId", "contentItemIdValue"),
	}
}

// String returns a human-readable description of this Content Item ID
func (id ContentItemId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Content Type: %q", id.ContentTypeId),
		fmt.Sprintf("Content Item: %q", id.ContentItemId),
	}
	return fmt.Sprintf("Content Item (%s)", strings.Join(components, "\n"))
}
