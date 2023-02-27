package apiissueattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AttachmentId{}

// AttachmentId is a struct representing the Resource ID for a Attachment
type AttachmentId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ApiId             string
	IssueId           string
	AttachmentId      string
}

// NewAttachmentID returns a new AttachmentId struct
func NewAttachmentID(subscriptionId string, resourceGroupName string, serviceName string, apiId string, issueId string, attachmentId string) AttachmentId {
	return AttachmentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ApiId:             apiId,
		IssueId:           issueId,
		AttachmentId:      attachmentId,
	}
}

// ParseAttachmentID parses 'input' into a AttachmentId
func ParseAttachmentID(input string) (*AttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(AttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AttachmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, fmt.Errorf("the segment 'apiId' was not found in the resource id %q", input)
	}

	if id.IssueId, ok = parsed.Parsed["issueId"]; !ok {
		return nil, fmt.Errorf("the segment 'issueId' was not found in the resource id %q", input)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, fmt.Errorf("the segment 'attachmentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAttachmentIDInsensitively parses 'input' case-insensitively into a AttachmentId
// note: this method should only be used for API response data and not user input
func ParseAttachmentIDInsensitively(input string) (*AttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(AttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AttachmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, fmt.Errorf("the segment 'apiId' was not found in the resource id %q", input)
	}

	if id.IssueId, ok = parsed.Parsed["issueId"]; !ok {
		return nil, fmt.Errorf("the segment 'issueId' was not found in the resource id %q", input)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, fmt.Errorf("the segment 'attachmentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAttachmentID checks that 'input' can be parsed as a Attachment ID
func ValidateAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Attachment ID
func (id AttachmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/apis/%s/issues/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ApiId, id.IssueId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Attachment ID
func (id AttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticApis", "apis", "apis"),
		resourceids.UserSpecifiedSegment("apiId", "apiIdValue"),
		resourceids.StaticSegment("staticIssues", "issues", "issues"),
		resourceids.UserSpecifiedSegment("issueId", "issueIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Attachment ID
func (id AttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Api: %q", id.ApiId),
		fmt.Sprintf("Issue: %q", id.IssueId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Attachment (%s)", strings.Join(components, "\n"))
}
