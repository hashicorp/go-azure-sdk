package apimconfig

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApiCollectionId{}

// ApiCollectionId is a struct representing the Resource ID for a Api Collection
type ApiCollectionId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ApiId             string
}

// NewApiCollectionID returns a new ApiCollectionId struct
func NewApiCollectionID(subscriptionId string, resourceGroupName string, serviceName string, apiId string) ApiCollectionId {
	return ApiCollectionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ApiId:             apiId,
	}
}

// ParseApiCollectionID parses 'input' into a ApiCollectionId
func ParseApiCollectionID(input string) (*ApiCollectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApiCollectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApiCollectionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApiCollectionIDInsensitively parses 'input' case-insensitively into a ApiCollectionId
// note: this method should only be used for API response data and not user input
func ParseApiCollectionIDInsensitively(input string) (*ApiCollectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApiCollectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApiCollectionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApiCollectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServiceName, ok = input.Parsed["serviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceName", input)
	}

	if id.ApiId, ok = input.Parsed["apiId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "apiId", input)
	}

	return nil
}

// ValidateApiCollectionID checks that 'input' can be parsed as a Api Collection ID
func ValidateApiCollectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApiCollectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Api Collection ID
func (id ApiCollectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/providers/Microsoft.Security/apiCollections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ApiId)
}

// Segments returns a slice of Resource ID Segments which comprise this Api Collection ID
func (id ApiCollectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticApiCollections", "apiCollections", "apiCollections"),
		resourceids.UserSpecifiedSegment("apiId", "apiIdValue"),
	}
}

// String returns a human-readable description of this Api Collection ID
func (id ApiCollectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Api: %q", id.ApiId),
	}
	return fmt.Sprintf("Api Collection (%s)", strings.Join(components, "\n"))
}
