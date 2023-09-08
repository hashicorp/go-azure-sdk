package userinferenceclassificationoverride

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInferenceClassificationOverrideId{}

// UserInferenceClassificationOverrideId is a struct representing the Resource ID for a User Inference Classification Override
type UserInferenceClassificationOverrideId struct {
	UserId                            string
	InferenceClassificationOverrideId string
}

// NewUserInferenceClassificationOverrideID returns a new UserInferenceClassificationOverrideId struct
func NewUserInferenceClassificationOverrideID(userId string, inferenceClassificationOverrideId string) UserInferenceClassificationOverrideId {
	return UserInferenceClassificationOverrideId{
		UserId:                            userId,
		InferenceClassificationOverrideId: inferenceClassificationOverrideId,
	}
}

// ParseUserInferenceClassificationOverrideID parses 'input' into a UserInferenceClassificationOverrideId
func ParseUserInferenceClassificationOverrideID(input string) (*UserInferenceClassificationOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInferenceClassificationOverrideId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInferenceClassificationOverrideId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.InferenceClassificationOverrideId, ok = parsed.Parsed["inferenceClassificationOverrideId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "inferenceClassificationOverrideId", *parsed)
	}

	return &id, nil
}

// ParseUserInferenceClassificationOverrideIDInsensitively parses 'input' case-insensitively into a UserInferenceClassificationOverrideId
// note: this method should only be used for API response data and not user input
func ParseUserInferenceClassificationOverrideIDInsensitively(input string) (*UserInferenceClassificationOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInferenceClassificationOverrideId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInferenceClassificationOverrideId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.InferenceClassificationOverrideId, ok = parsed.Parsed["inferenceClassificationOverrideId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "inferenceClassificationOverrideId", *parsed)
	}

	return &id, nil
}

// ValidateUserInferenceClassificationOverrideID checks that 'input' can be parsed as a User Inference Classification Override ID
func ValidateUserInferenceClassificationOverrideID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInferenceClassificationOverrideID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Inference Classification Override ID
func (id UserInferenceClassificationOverrideId) ID() string {
	fmtString := "/users/%s/inferenceClassification/overrides/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.InferenceClassificationOverrideId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Inference Classification Override ID
func (id UserInferenceClassificationOverrideId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInferenceClassification", "inferenceClassification", "inferenceClassification"),
		resourceids.StaticSegment("staticOverrides", "overrides", "overrides"),
		resourceids.UserSpecifiedSegment("inferenceClassificationOverrideId", "inferenceClassificationOverrideIdValue"),
	}
}

// String returns a human-readable description of this User Inference Classification Override ID
func (id UserInferenceClassificationOverrideId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Inference Classification Override: %q", id.InferenceClassificationOverrideId),
	}
	return fmt.Sprintf("User Inference Classification Override (%s)", strings.Join(components, "\n"))
}
