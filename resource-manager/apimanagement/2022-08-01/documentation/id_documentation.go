package documentation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DocumentationId{}

// DocumentationId is a struct representing the Resource ID for a Documentation
type DocumentationId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	DocumentationId   string
}

// NewDocumentationID returns a new DocumentationId struct
func NewDocumentationID(subscriptionId string, resourceGroupName string, serviceName string, documentationId string) DocumentationId {
	return DocumentationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		DocumentationId:   documentationId,
	}
}

// ParseDocumentationID parses 'input' into a DocumentationId
func ParseDocumentationID(input string) (*DocumentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DocumentationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DocumentationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.DocumentationId, ok = parsed.Parsed["documentationId"]; !ok {
		return nil, fmt.Errorf("the segment 'documentationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDocumentationIDInsensitively parses 'input' case-insensitively into a DocumentationId
// note: this method should only be used for API response data and not user input
func ParseDocumentationIDInsensitively(input string) (*DocumentationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DocumentationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DocumentationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.DocumentationId, ok = parsed.Parsed["documentationId"]; !ok {
		return nil, fmt.Errorf("the segment 'documentationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDocumentationID checks that 'input' can be parsed as a Documentation ID
func ValidateDocumentationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDocumentationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Documentation ID
func (id DocumentationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/documentations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.DocumentationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Documentation ID
func (id DocumentationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticDocumentations", "documentations", "documentations"),
		resourceids.UserSpecifiedSegment("documentationId", "documentationIdValue"),
	}
}

// String returns a human-readable description of this Documentation ID
func (id DocumentationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Documentation: %q", id.DocumentationId),
	}
	return fmt.Sprintf("Documentation (%s)", strings.Join(components, "\n"))
}
