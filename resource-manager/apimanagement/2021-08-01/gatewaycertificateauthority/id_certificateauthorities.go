package gatewaycertificateauthority

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CertificateAuthoritiesId{}

// CertificateAuthoritiesId is a struct representing the Resource ID for a Certificate Authorities
type CertificateAuthoritiesId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	GatewayId         string
	CertificateId     string
}

// NewCertificateAuthoritiesID returns a new CertificateAuthoritiesId struct
func NewCertificateAuthoritiesID(subscriptionId string, resourceGroupName string, serviceName string, gatewayId string, certificateId string) CertificateAuthoritiesId {
	return CertificateAuthoritiesId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		GatewayId:         gatewayId,
		CertificateId:     certificateId,
	}
}

// ParseCertificateAuthoritiesID parses 'input' into a CertificateAuthoritiesId
func ParseCertificateAuthoritiesID(input string) (*CertificateAuthoritiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(CertificateAuthoritiesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CertificateAuthoritiesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.GatewayId, ok = parsed.Parsed["gatewayId"]; !ok {
		return nil, fmt.Errorf("the segment 'gatewayId' was not found in the resource id %q", input)
	}

	if id.CertificateId, ok = parsed.Parsed["certificateId"]; !ok {
		return nil, fmt.Errorf("the segment 'certificateId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCertificateAuthoritiesIDInsensitively parses 'input' case-insensitively into a CertificateAuthoritiesId
// note: this method should only be used for API response data and not user input
func ParseCertificateAuthoritiesIDInsensitively(input string) (*CertificateAuthoritiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(CertificateAuthoritiesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CertificateAuthoritiesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.GatewayId, ok = parsed.Parsed["gatewayId"]; !ok {
		return nil, fmt.Errorf("the segment 'gatewayId' was not found in the resource id %q", input)
	}

	if id.CertificateId, ok = parsed.Parsed["certificateId"]; !ok {
		return nil, fmt.Errorf("the segment 'certificateId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCertificateAuthoritiesID checks that 'input' can be parsed as a Certificate Authorities ID
func ValidateCertificateAuthoritiesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCertificateAuthoritiesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Certificate Authorities ID
func (id CertificateAuthoritiesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/gateways/%s/certificateAuthorities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.GatewayId, id.CertificateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Certificate Authorities ID
func (id CertificateAuthoritiesId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticGateways", "gateways", "gateways"),
		resourceids.UserSpecifiedSegment("gatewayId", "gatewayIdValue"),
		resourceids.StaticSegment("staticCertificateAuthorities", "certificateAuthorities", "certificateAuthorities"),
		resourceids.UserSpecifiedSegment("certificateId", "certificateIdValue"),
	}
}

// String returns a human-readable description of this Certificate Authorities ID
func (id CertificateAuthoritiesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Gateway: %q", id.GatewayId),
		fmt.Sprintf("Certificate: %q", id.CertificateId),
	}
	return fmt.Sprintf("Certificate Authorities (%s)", strings.Join(components, "\n"))
}
