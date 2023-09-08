package userjoinedteamprimarychannelmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelMessageHostedContentId{}

// UserJoinedTeamPrimaryChannelMessageHostedContentId is a struct representing the Resource ID for a User Joined Team Primary Channel Message Hosted Content
type UserJoinedTeamPrimaryChannelMessageHostedContentId struct {
	UserId                     string
	TeamId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewUserJoinedTeamPrimaryChannelMessageHostedContentID returns a new UserJoinedTeamPrimaryChannelMessageHostedContentId struct
func NewUserJoinedTeamPrimaryChannelMessageHostedContentID(userId string, teamId string, chatMessageId string, chatMessageHostedContentId string) UserJoinedTeamPrimaryChannelMessageHostedContentId {
	return UserJoinedTeamPrimaryChannelMessageHostedContentId{
		UserId:                     userId,
		TeamId:                     teamId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserJoinedTeamPrimaryChannelMessageHostedContentID parses 'input' into a UserJoinedTeamPrimaryChannelMessageHostedContentId
func ParseUserJoinedTeamPrimaryChannelMessageHostedContentID(input string) (*UserJoinedTeamPrimaryChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMessageHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPrimaryChannelMessageHostedContentIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPrimaryChannelMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPrimaryChannelMessageHostedContentIDInsensitively(input string) (*UserJoinedTeamPrimaryChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelMessageHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPrimaryChannelMessageHostedContentID checks that 'input' can be parsed as a User Joined Team Primary Channel Message Hosted Content ID
func ValidateUserJoinedTeamPrimaryChannelMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPrimaryChannelMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Primary Channel Message Hosted Content ID
func (id UserJoinedTeamPrimaryChannelMessageHostedContentId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Primary Channel Message Hosted Content ID
func (id UserJoinedTeamPrimaryChannelMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Primary Channel Message Hosted Content ID
func (id UserJoinedTeamPrimaryChannelMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Joined Team Primary Channel Message Hosted Content (%s)", strings.Join(components, "\n"))
}
