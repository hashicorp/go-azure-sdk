package videos

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VideoId{}

// VideoId is a struct representing the Resource ID for a Video
type VideoId struct {
	SubscriptionId    string
	ResourceGroupName string
	VideoAnalyzerName string
	VideoName         string
}

// NewVideoID returns a new VideoId struct
func NewVideoID(subscriptionId string, resourceGroupName string, videoAnalyzerName string, videoName string) VideoId {
	return VideoId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VideoAnalyzerName: videoAnalyzerName,
		VideoName:         videoName,
	}
}

// ParseVideoID parses 'input' into a VideoId
func ParseVideoID(input string) (*VideoId, error) {
	parser := resourceids.NewParserFromResourceIdType(VideoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VideoId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VideoAnalyzerName, ok = parsed.Parsed["videoAnalyzerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "videoAnalyzerName", *parsed)
	}

	if id.VideoName, ok = parsed.Parsed["videoName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "videoName", *parsed)
	}

	return &id, nil
}

// ParseVideoIDInsensitively parses 'input' case-insensitively into a VideoId
// note: this method should only be used for API response data and not user input
func ParseVideoIDInsensitively(input string) (*VideoId, error) {
	parser := resourceids.NewParserFromResourceIdType(VideoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VideoId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VideoAnalyzerName, ok = parsed.Parsed["videoAnalyzerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "videoAnalyzerName", *parsed)
	}

	if id.VideoName, ok = parsed.Parsed["videoName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "videoName", *parsed)
	}

	return &id, nil
}

// ValidateVideoID checks that 'input' can be parsed as a Video ID
func ValidateVideoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVideoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Video ID
func (id VideoId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Media/videoAnalyzers/%s/videos/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VideoAnalyzerName, id.VideoName)
}

// Segments returns a slice of Resource ID Segments which comprise this Video ID
func (id VideoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMedia", "Microsoft.Media", "Microsoft.Media"),
		resourceids.StaticSegment("staticVideoAnalyzers", "videoAnalyzers", "videoAnalyzers"),
		resourceids.UserSpecifiedSegment("videoAnalyzerName", "videoAnalyzerValue"),
		resourceids.StaticSegment("staticVideos", "videos", "videos"),
		resourceids.UserSpecifiedSegment("videoName", "videoValue"),
	}
}

// String returns a human-readable description of this Video ID
func (id VideoId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Video Analyzer Name: %q", id.VideoAnalyzerName),
		fmt.Sprintf("Video Name: %q", id.VideoName),
	}
	return fmt.Sprintf("Video (%s)", strings.Join(components, "\n"))
}
