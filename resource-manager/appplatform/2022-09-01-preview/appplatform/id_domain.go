package appplatform

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DomainId{}

// DomainId is a struct representing the Resource ID for a Domain
type DomainId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	AppName           string
	DomainName        string
}

// NewDomainID returns a new DomainId struct
func NewDomainID(subscriptionId string, resourceGroupName string, serviceName string, appName string, domainName string) DomainId {
	return DomainId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		AppName:           appName,
		DomainName:        domainName,
	}
}

// ParseDomainID parses 'input' into a DomainId
func ParseDomainID(input string) (*DomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.AppName, ok = parsed.Parsed["appName"]; !ok {
		return nil, fmt.Errorf("the segment 'appName' was not found in the resource id %q", input)
	}

	if id.DomainName, ok = parsed.Parsed["domainName"]; !ok {
		return nil, fmt.Errorf("the segment 'domainName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDomainIDInsensitively parses 'input' case-insensitively into a DomainId
// note: this method should only be used for API response data and not user input
func ParseDomainIDInsensitively(input string) (*DomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.AppName, ok = parsed.Parsed["appName"]; !ok {
		return nil, fmt.Errorf("the segment 'appName' was not found in the resource id %q", input)
	}

	if id.DomainName, ok = parsed.Parsed["domainName"]; !ok {
		return nil, fmt.Errorf("the segment 'domainName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDomainID checks that 'input' can be parsed as a Domain ID
func ValidateDomainID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain ID
func (id DomainId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AppPlatform/spring/%s/apps/%s/domains/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.AppName, id.DomainName)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain ID
func (id DomainId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAppPlatform", "Microsoft.AppPlatform", "Microsoft.AppPlatform"),
		resourceids.StaticSegment("staticSpring", "spring", "spring"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticApps", "apps", "apps"),
		resourceids.UserSpecifiedSegment("appName", "appValue"),
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainName", "domainValue"),
	}
}

// String returns a human-readable description of this Domain ID
func (id DomainId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("App Name: %q", id.AppName),
		fmt.Sprintf("Domain Name: %q", id.DomainName),
	}
	return fmt.Sprintf("Domain (%s)", strings.Join(components, "\n"))
}
