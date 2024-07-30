package lifecycleworkflowdeleteditemworkflow

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityHistoryItemStatus string

const (
	ActivityHistoryItemStatusActive  ActivityHistoryItemStatus = "active"
	ActivityHistoryItemStatusDeleted ActivityHistoryItemStatus = "deleted"
	ActivityHistoryItemStatusIgnored ActivityHistoryItemStatus = "ignored"
	ActivityHistoryItemStatusUpdated ActivityHistoryItemStatus = "updated"
)

func PossibleValuesForActivityHistoryItemStatus() []string {
	return []string{
		string(ActivityHistoryItemStatusActive),
		string(ActivityHistoryItemStatusDeleted),
		string(ActivityHistoryItemStatusIgnored),
		string(ActivityHistoryItemStatusUpdated),
	}
}

func (s *ActivityHistoryItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActivityHistoryItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActivityHistoryItemStatus(input string) (*ActivityHistoryItemStatus, error) {
	vals := map[string]ActivityHistoryItemStatus{
		"active":  ActivityHistoryItemStatusActive,
		"deleted": ActivityHistoryItemStatusDeleted,
		"ignored": ActivityHistoryItemStatusIgnored,
		"updated": ActivityHistoryItemStatusUpdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityHistoryItemStatus(input)
	return &out, nil
}

type AgreementAcceptanceState string

const (
	AgreementAcceptanceStateAccepted AgreementAcceptanceState = "accepted"
	AgreementAcceptanceStateDeclined AgreementAcceptanceState = "declined"
)

func PossibleValuesForAgreementAcceptanceState() []string {
	return []string{
		string(AgreementAcceptanceStateAccepted),
		string(AgreementAcceptanceStateDeclined),
	}
}

func (s *AgreementAcceptanceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAgreementAcceptanceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAgreementAcceptanceState(input string) (*AgreementAcceptanceState, error) {
	vals := map[string]AgreementAcceptanceState{
		"accepted": AgreementAcceptanceStateAccepted,
		"declined": AgreementAcceptanceStateDeclined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgreementAcceptanceState(input)
	return &out, nil
}

type AttendeeType string

const (
	AttendeeTypeOptional AttendeeType = "optional"
	AttendeeTypeRequired AttendeeType = "required"
	AttendeeTypeResource AttendeeType = "resource"
)

func PossibleValuesForAttendeeType() []string {
	return []string{
		string(AttendeeTypeOptional),
		string(AttendeeTypeRequired),
		string(AttendeeTypeResource),
	}
}

func (s *AttendeeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttendeeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttendeeType(input string) (*AttendeeType, error) {
	vals := map[string]AttendeeType{
		"optional": AttendeeTypeOptional,
		"required": AttendeeTypeRequired,
		"resource": AttendeeTypeResource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeType(input)
	return &out, nil
}

type AutomaticRepliesSettingExternalAudience string

const (
	AutomaticRepliesSettingExternalAudienceAll          AutomaticRepliesSettingExternalAudience = "all"
	AutomaticRepliesSettingExternalAudienceContactsOnly AutomaticRepliesSettingExternalAudience = "contactsOnly"
	AutomaticRepliesSettingExternalAudienceNone         AutomaticRepliesSettingExternalAudience = "none"
)

func PossibleValuesForAutomaticRepliesSettingExternalAudience() []string {
	return []string{
		string(AutomaticRepliesSettingExternalAudienceAll),
		string(AutomaticRepliesSettingExternalAudienceContactsOnly),
		string(AutomaticRepliesSettingExternalAudienceNone),
	}
}

func (s *AutomaticRepliesSettingExternalAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticRepliesSettingExternalAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticRepliesSettingExternalAudience(input string) (*AutomaticRepliesSettingExternalAudience, error) {
	vals := map[string]AutomaticRepliesSettingExternalAudience{
		"all":          AutomaticRepliesSettingExternalAudienceAll,
		"contactsonly": AutomaticRepliesSettingExternalAudienceContactsOnly,
		"none":         AutomaticRepliesSettingExternalAudienceNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticRepliesSettingExternalAudience(input)
	return &out, nil
}

type AutomaticRepliesSettingStatus string

const (
	AutomaticRepliesSettingStatusAlwaysEnabled AutomaticRepliesSettingStatus = "alwaysEnabled"
	AutomaticRepliesSettingStatusDisabled      AutomaticRepliesSettingStatus = "disabled"
	AutomaticRepliesSettingStatusScheduled     AutomaticRepliesSettingStatus = "scheduled"
)

func PossibleValuesForAutomaticRepliesSettingStatus() []string {
	return []string{
		string(AutomaticRepliesSettingStatusAlwaysEnabled),
		string(AutomaticRepliesSettingStatusDisabled),
		string(AutomaticRepliesSettingStatusScheduled),
	}
}

func (s *AutomaticRepliesSettingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticRepliesSettingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticRepliesSettingStatus(input string) (*AutomaticRepliesSettingStatus, error) {
	vals := map[string]AutomaticRepliesSettingStatus{
		"alwaysenabled": AutomaticRepliesSettingStatusAlwaysEnabled,
		"disabled":      AutomaticRepliesSettingStatusDisabled,
		"scheduled":     AutomaticRepliesSettingStatusScheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticRepliesSettingStatus(input)
	return &out, nil
}

type BroadcastMeetingSettingsAllowedAudience string

const (
	BroadcastMeetingSettingsAllowedAudienceEveryone       BroadcastMeetingSettingsAllowedAudience = "everyone"
	BroadcastMeetingSettingsAllowedAudienceOrganization   BroadcastMeetingSettingsAllowedAudience = "organization"
	BroadcastMeetingSettingsAllowedAudienceRoleIsAttendee BroadcastMeetingSettingsAllowedAudience = "roleIsAttendee"
)

func PossibleValuesForBroadcastMeetingSettingsAllowedAudience() []string {
	return []string{
		string(BroadcastMeetingSettingsAllowedAudienceEveryone),
		string(BroadcastMeetingSettingsAllowedAudienceOrganization),
		string(BroadcastMeetingSettingsAllowedAudienceRoleIsAttendee),
	}
}

func (s *BroadcastMeetingSettingsAllowedAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBroadcastMeetingSettingsAllowedAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBroadcastMeetingSettingsAllowedAudience(input string) (*BroadcastMeetingSettingsAllowedAudience, error) {
	vals := map[string]BroadcastMeetingSettingsAllowedAudience{
		"everyone":       BroadcastMeetingSettingsAllowedAudienceEveryone,
		"organization":   BroadcastMeetingSettingsAllowedAudienceOrganization,
		"roleisattendee": BroadcastMeetingSettingsAllowedAudienceRoleIsAttendee,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BroadcastMeetingSettingsAllowedAudience(input)
	return &out, nil
}

type CalendarAllowedOnlineMeetingProviders string

const (
	CalendarAllowedOnlineMeetingProvidersSkypeForBusiness CalendarAllowedOnlineMeetingProviders = "skypeForBusiness"
	CalendarAllowedOnlineMeetingProvidersSkypeForConsumer CalendarAllowedOnlineMeetingProviders = "skypeForConsumer"
	CalendarAllowedOnlineMeetingProvidersTeamsForBusiness CalendarAllowedOnlineMeetingProviders = "teamsForBusiness"
	CalendarAllowedOnlineMeetingProvidersUnknown          CalendarAllowedOnlineMeetingProviders = "unknown"
)

func PossibleValuesForCalendarAllowedOnlineMeetingProviders() []string {
	return []string{
		string(CalendarAllowedOnlineMeetingProvidersSkypeForBusiness),
		string(CalendarAllowedOnlineMeetingProvidersSkypeForConsumer),
		string(CalendarAllowedOnlineMeetingProvidersTeamsForBusiness),
		string(CalendarAllowedOnlineMeetingProvidersUnknown),
	}
}

func (s *CalendarAllowedOnlineMeetingProviders) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarAllowedOnlineMeetingProviders(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarAllowedOnlineMeetingProviders(input string) (*CalendarAllowedOnlineMeetingProviders, error) {
	vals := map[string]CalendarAllowedOnlineMeetingProviders{
		"skypeforbusiness": CalendarAllowedOnlineMeetingProvidersSkypeForBusiness,
		"skypeforconsumer": CalendarAllowedOnlineMeetingProvidersSkypeForConsumer,
		"teamsforbusiness": CalendarAllowedOnlineMeetingProvidersTeamsForBusiness,
		"unknown":          CalendarAllowedOnlineMeetingProvidersUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarAllowedOnlineMeetingProviders(input)
	return &out, nil
}

type CalendarColor string

const (
	CalendarColorAuto        CalendarColor = "auto"
	CalendarColorLightBlue   CalendarColor = "lightBlue"
	CalendarColorLightBrown  CalendarColor = "lightBrown"
	CalendarColorLightGray   CalendarColor = "lightGray"
	CalendarColorLightGreen  CalendarColor = "lightGreen"
	CalendarColorLightOrange CalendarColor = "lightOrange"
	CalendarColorLightPink   CalendarColor = "lightPink"
	CalendarColorLightRed    CalendarColor = "lightRed"
	CalendarColorLightTeal   CalendarColor = "lightTeal"
	CalendarColorLightYellow CalendarColor = "lightYellow"
	CalendarColorMaxColor    CalendarColor = "maxColor"
)

func PossibleValuesForCalendarColor() []string {
	return []string{
		string(CalendarColorAuto),
		string(CalendarColorLightBlue),
		string(CalendarColorLightBrown),
		string(CalendarColorLightGray),
		string(CalendarColorLightGreen),
		string(CalendarColorLightOrange),
		string(CalendarColorLightPink),
		string(CalendarColorLightRed),
		string(CalendarColorLightTeal),
		string(CalendarColorLightYellow),
		string(CalendarColorMaxColor),
	}
}

func (s *CalendarColor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarColor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarColor(input string) (*CalendarColor, error) {
	vals := map[string]CalendarColor{
		"auto":        CalendarColorAuto,
		"lightblue":   CalendarColorLightBlue,
		"lightbrown":  CalendarColorLightBrown,
		"lightgray":   CalendarColorLightGray,
		"lightgreen":  CalendarColorLightGreen,
		"lightorange": CalendarColorLightOrange,
		"lightpink":   CalendarColorLightPink,
		"lightred":    CalendarColorLightRed,
		"lightteal":   CalendarColorLightTeal,
		"lightyellow": CalendarColorLightYellow,
		"maxcolor":    CalendarColorMaxColor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarColor(input)
	return &out, nil
}

type CalendarDefaultOnlineMeetingProvider string

const (
	CalendarDefaultOnlineMeetingProviderSkypeForBusiness CalendarDefaultOnlineMeetingProvider = "skypeForBusiness"
	CalendarDefaultOnlineMeetingProviderSkypeForConsumer CalendarDefaultOnlineMeetingProvider = "skypeForConsumer"
	CalendarDefaultOnlineMeetingProviderTeamsForBusiness CalendarDefaultOnlineMeetingProvider = "teamsForBusiness"
	CalendarDefaultOnlineMeetingProviderUnknown          CalendarDefaultOnlineMeetingProvider = "unknown"
)

func PossibleValuesForCalendarDefaultOnlineMeetingProvider() []string {
	return []string{
		string(CalendarDefaultOnlineMeetingProviderSkypeForBusiness),
		string(CalendarDefaultOnlineMeetingProviderSkypeForConsumer),
		string(CalendarDefaultOnlineMeetingProviderTeamsForBusiness),
		string(CalendarDefaultOnlineMeetingProviderUnknown),
	}
}

func (s *CalendarDefaultOnlineMeetingProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarDefaultOnlineMeetingProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarDefaultOnlineMeetingProvider(input string) (*CalendarDefaultOnlineMeetingProvider, error) {
	vals := map[string]CalendarDefaultOnlineMeetingProvider{
		"skypeforbusiness": CalendarDefaultOnlineMeetingProviderSkypeForBusiness,
		"skypeforconsumer": CalendarDefaultOnlineMeetingProviderSkypeForConsumer,
		"teamsforbusiness": CalendarDefaultOnlineMeetingProviderTeamsForBusiness,
		"unknown":          CalendarDefaultOnlineMeetingProviderUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarDefaultOnlineMeetingProvider(input)
	return &out, nil
}

type CalendarPermissionAllowedRoles string

const (
	CalendarPermissionAllowedRolesCustom                            CalendarPermissionAllowedRoles = "custom"
	CalendarPermissionAllowedRolesDelegateWithPrivateEventAccess    CalendarPermissionAllowedRoles = "delegateWithPrivateEventAccess"
	CalendarPermissionAllowedRolesDelegateWithoutPrivateEventAccess CalendarPermissionAllowedRoles = "delegateWithoutPrivateEventAccess"
	CalendarPermissionAllowedRolesFreeBusyRead                      CalendarPermissionAllowedRoles = "freeBusyRead"
	CalendarPermissionAllowedRolesLimitedRead                       CalendarPermissionAllowedRoles = "limitedRead"
	CalendarPermissionAllowedRolesNone                              CalendarPermissionAllowedRoles = "none"
	CalendarPermissionAllowedRolesRead                              CalendarPermissionAllowedRoles = "read"
	CalendarPermissionAllowedRolesWrite                             CalendarPermissionAllowedRoles = "write"
)

func PossibleValuesForCalendarPermissionAllowedRoles() []string {
	return []string{
		string(CalendarPermissionAllowedRolesCustom),
		string(CalendarPermissionAllowedRolesDelegateWithPrivateEventAccess),
		string(CalendarPermissionAllowedRolesDelegateWithoutPrivateEventAccess),
		string(CalendarPermissionAllowedRolesFreeBusyRead),
		string(CalendarPermissionAllowedRolesLimitedRead),
		string(CalendarPermissionAllowedRolesNone),
		string(CalendarPermissionAllowedRolesRead),
		string(CalendarPermissionAllowedRolesWrite),
	}
}

func (s *CalendarPermissionAllowedRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarPermissionAllowedRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarPermissionAllowedRoles(input string) (*CalendarPermissionAllowedRoles, error) {
	vals := map[string]CalendarPermissionAllowedRoles{
		"custom":                            CalendarPermissionAllowedRolesCustom,
		"delegatewithprivateeventaccess":    CalendarPermissionAllowedRolesDelegateWithPrivateEventAccess,
		"delegatewithoutprivateeventaccess": CalendarPermissionAllowedRolesDelegateWithoutPrivateEventAccess,
		"freebusyread":                      CalendarPermissionAllowedRolesFreeBusyRead,
		"limitedread":                       CalendarPermissionAllowedRolesLimitedRead,
		"none":                              CalendarPermissionAllowedRolesNone,
		"read":                              CalendarPermissionAllowedRolesRead,
		"write":                             CalendarPermissionAllowedRolesWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarPermissionAllowedRoles(input)
	return &out, nil
}

type CalendarPermissionRole string

const (
	CalendarPermissionRoleCustom                            CalendarPermissionRole = "custom"
	CalendarPermissionRoleDelegateWithPrivateEventAccess    CalendarPermissionRole = "delegateWithPrivateEventAccess"
	CalendarPermissionRoleDelegateWithoutPrivateEventAccess CalendarPermissionRole = "delegateWithoutPrivateEventAccess"
	CalendarPermissionRoleFreeBusyRead                      CalendarPermissionRole = "freeBusyRead"
	CalendarPermissionRoleLimitedRead                       CalendarPermissionRole = "limitedRead"
	CalendarPermissionRoleNone                              CalendarPermissionRole = "none"
	CalendarPermissionRoleRead                              CalendarPermissionRole = "read"
	CalendarPermissionRoleWrite                             CalendarPermissionRole = "write"
)

func PossibleValuesForCalendarPermissionRole() []string {
	return []string{
		string(CalendarPermissionRoleCustom),
		string(CalendarPermissionRoleDelegateWithPrivateEventAccess),
		string(CalendarPermissionRoleDelegateWithoutPrivateEventAccess),
		string(CalendarPermissionRoleFreeBusyRead),
		string(CalendarPermissionRoleLimitedRead),
		string(CalendarPermissionRoleNone),
		string(CalendarPermissionRoleRead),
		string(CalendarPermissionRoleWrite),
	}
}

func (s *CalendarPermissionRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarPermissionRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarPermissionRole(input string) (*CalendarPermissionRole, error) {
	vals := map[string]CalendarPermissionRole{
		"custom":                            CalendarPermissionRoleCustom,
		"delegatewithprivateeventaccess":    CalendarPermissionRoleDelegateWithPrivateEventAccess,
		"delegatewithoutprivateeventaccess": CalendarPermissionRoleDelegateWithoutPrivateEventAccess,
		"freebusyread":                      CalendarPermissionRoleFreeBusyRead,
		"limitedread":                       CalendarPermissionRoleLimitedRead,
		"none":                              CalendarPermissionRoleNone,
		"read":                              CalendarPermissionRoleRead,
		"write":                             CalendarPermissionRoleWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarPermissionRole(input)
	return &out, nil
}

type ChannelMembershipType string

const (
	ChannelMembershipTypePrivate  ChannelMembershipType = "private"
	ChannelMembershipTypeShared   ChannelMembershipType = "shared"
	ChannelMembershipTypeStandard ChannelMembershipType = "standard"
)

func PossibleValuesForChannelMembershipType() []string {
	return []string{
		string(ChannelMembershipTypePrivate),
		string(ChannelMembershipTypeShared),
		string(ChannelMembershipTypeStandard),
	}
}

func (s *ChannelMembershipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChannelMembershipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChannelMembershipType(input string) (*ChannelMembershipType, error) {
	vals := map[string]ChannelMembershipType{
		"private":  ChannelMembershipTypePrivate,
		"shared":   ChannelMembershipTypeShared,
		"standard": ChannelMembershipTypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelMembershipType(input)
	return &out, nil
}

type ChatChatType string

const (
	ChatChatTypeGroup    ChatChatType = "group"
	ChatChatTypeMeeting  ChatChatType = "meeting"
	ChatChatTypeOneOnOne ChatChatType = "oneOnOne"
)

func PossibleValuesForChatChatType() []string {
	return []string{
		string(ChatChatTypeGroup),
		string(ChatChatTypeMeeting),
		string(ChatChatTypeOneOnOne),
	}
}

func (s *ChatChatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatChatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatChatType(input string) (*ChatChatType, error) {
	vals := map[string]ChatChatType{
		"group":    ChatChatTypeGroup,
		"meeting":  ChatChatTypeMeeting,
		"oneonone": ChatChatTypeOneOnOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatChatType(input)
	return &out, nil
}

type ChatMessageHistoryItemActions string

const (
	ChatMessageHistoryItemActionsActionUndefined ChatMessageHistoryItemActions = "actionUndefined"
	ChatMessageHistoryItemActionsReactionAdded   ChatMessageHistoryItemActions = "reactionAdded"
	ChatMessageHistoryItemActionsReactionRemoved ChatMessageHistoryItemActions = "reactionRemoved"
)

func PossibleValuesForChatMessageHistoryItemActions() []string {
	return []string{
		string(ChatMessageHistoryItemActionsActionUndefined),
		string(ChatMessageHistoryItemActionsReactionAdded),
		string(ChatMessageHistoryItemActionsReactionRemoved),
	}
}

func (s *ChatMessageHistoryItemActions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageHistoryItemActions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageHistoryItemActions(input string) (*ChatMessageHistoryItemActions, error) {
	vals := map[string]ChatMessageHistoryItemActions{
		"actionundefined": ChatMessageHistoryItemActionsActionUndefined,
		"reactionadded":   ChatMessageHistoryItemActionsReactionAdded,
		"reactionremoved": ChatMessageHistoryItemActionsReactionRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageHistoryItemActions(input)
	return &out, nil
}

type ChatMessageImportance string

const (
	ChatMessageImportanceHigh   ChatMessageImportance = "high"
	ChatMessageImportanceNormal ChatMessageImportance = "normal"
	ChatMessageImportanceUrgent ChatMessageImportance = "urgent"
)

func PossibleValuesForChatMessageImportance() []string {
	return []string{
		string(ChatMessageImportanceHigh),
		string(ChatMessageImportanceNormal),
		string(ChatMessageImportanceUrgent),
	}
}

func (s *ChatMessageImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageImportance(input string) (*ChatMessageImportance, error) {
	vals := map[string]ChatMessageImportance{
		"high":   ChatMessageImportanceHigh,
		"normal": ChatMessageImportanceNormal,
		"urgent": ChatMessageImportanceUrgent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageImportance(input)
	return &out, nil
}

type ChatMessageInfoMessageType string

const (
	ChatMessageInfoMessageTypeChatEvent          ChatMessageInfoMessageType = "chatEvent"
	ChatMessageInfoMessageTypeMessage            ChatMessageInfoMessageType = "message"
	ChatMessageInfoMessageTypeSystemEventMessage ChatMessageInfoMessageType = "systemEventMessage"
	ChatMessageInfoMessageTypeTyping             ChatMessageInfoMessageType = "typing"
)

func PossibleValuesForChatMessageInfoMessageType() []string {
	return []string{
		string(ChatMessageInfoMessageTypeChatEvent),
		string(ChatMessageInfoMessageTypeMessage),
		string(ChatMessageInfoMessageTypeSystemEventMessage),
		string(ChatMessageInfoMessageTypeTyping),
	}
}

func (s *ChatMessageInfoMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageInfoMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageInfoMessageType(input string) (*ChatMessageInfoMessageType, error) {
	vals := map[string]ChatMessageInfoMessageType{
		"chatevent":          ChatMessageInfoMessageTypeChatEvent,
		"message":            ChatMessageInfoMessageTypeMessage,
		"systemeventmessage": ChatMessageInfoMessageTypeSystemEventMessage,
		"typing":             ChatMessageInfoMessageTypeTyping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageInfoMessageType(input)
	return &out, nil
}

type ChatMessageMessageType string

const (
	ChatMessageMessageTypeChatEvent          ChatMessageMessageType = "chatEvent"
	ChatMessageMessageTypeMessage            ChatMessageMessageType = "message"
	ChatMessageMessageTypeSystemEventMessage ChatMessageMessageType = "systemEventMessage"
	ChatMessageMessageTypeTyping             ChatMessageMessageType = "typing"
)

func PossibleValuesForChatMessageMessageType() []string {
	return []string{
		string(ChatMessageMessageTypeChatEvent),
		string(ChatMessageMessageTypeMessage),
		string(ChatMessageMessageTypeSystemEventMessage),
		string(ChatMessageMessageTypeTyping),
	}
}

func (s *ChatMessageMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessageMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessageMessageType(input string) (*ChatMessageMessageType, error) {
	vals := map[string]ChatMessageMessageType{
		"chatevent":          ChatMessageMessageTypeChatEvent,
		"message":            ChatMessageMessageTypeMessage,
		"systemeventmessage": ChatMessageMessageTypeSystemEventMessage,
		"typing":             ChatMessageMessageTypeTyping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessageMessageType(input)
	return &out, nil
}

type ChatMessagePolicyViolationDlpAction string

const (
	ChatMessagePolicyViolationDlpActionBlockAccess         ChatMessagePolicyViolationDlpAction = "blockAccess"
	ChatMessagePolicyViolationDlpActionBlockAccessExternal ChatMessagePolicyViolationDlpAction = "blockAccessExternal"
	ChatMessagePolicyViolationDlpActionNone                ChatMessagePolicyViolationDlpAction = "none"
	ChatMessagePolicyViolationDlpActionNotifySender        ChatMessagePolicyViolationDlpAction = "notifySender"
)

func PossibleValuesForChatMessagePolicyViolationDlpAction() []string {
	return []string{
		string(ChatMessagePolicyViolationDlpActionBlockAccess),
		string(ChatMessagePolicyViolationDlpActionBlockAccessExternal),
		string(ChatMessagePolicyViolationDlpActionNone),
		string(ChatMessagePolicyViolationDlpActionNotifySender),
	}
}

func (s *ChatMessagePolicyViolationDlpAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationDlpAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationDlpAction(input string) (*ChatMessagePolicyViolationDlpAction, error) {
	vals := map[string]ChatMessagePolicyViolationDlpAction{
		"blockaccess":         ChatMessagePolicyViolationDlpActionBlockAccess,
		"blockaccessexternal": ChatMessagePolicyViolationDlpActionBlockAccessExternal,
		"none":                ChatMessagePolicyViolationDlpActionNone,
		"notifysender":        ChatMessagePolicyViolationDlpActionNotifySender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationDlpAction(input)
	return &out, nil
}

type ChatMessagePolicyViolationUserAction string

const (
	ChatMessagePolicyViolationUserActionNone                ChatMessagePolicyViolationUserAction = "none"
	ChatMessagePolicyViolationUserActionOverride            ChatMessagePolicyViolationUserAction = "override"
	ChatMessagePolicyViolationUserActionReportFalsePositive ChatMessagePolicyViolationUserAction = "reportFalsePositive"
)

func PossibleValuesForChatMessagePolicyViolationUserAction() []string {
	return []string{
		string(ChatMessagePolicyViolationUserActionNone),
		string(ChatMessagePolicyViolationUserActionOverride),
		string(ChatMessagePolicyViolationUserActionReportFalsePositive),
	}
}

func (s *ChatMessagePolicyViolationUserAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationUserAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationUserAction(input string) (*ChatMessagePolicyViolationUserAction, error) {
	vals := map[string]ChatMessagePolicyViolationUserAction{
		"none":                ChatMessagePolicyViolationUserActionNone,
		"override":            ChatMessagePolicyViolationUserActionOverride,
		"reportfalsepositive": ChatMessagePolicyViolationUserActionReportFalsePositive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationUserAction(input)
	return &out, nil
}

type ChatMessagePolicyViolationVerdictDetails string

const (
	ChatMessagePolicyViolationVerdictDetailsAllowFalsePositiveOverride        ChatMessagePolicyViolationVerdictDetails = "allowFalsePositiveOverride"
	ChatMessagePolicyViolationVerdictDetailsAllowOverrideWithJustification    ChatMessagePolicyViolationVerdictDetails = "allowOverrideWithJustification"
	ChatMessagePolicyViolationVerdictDetailsAllowOverrideWithoutJustification ChatMessagePolicyViolationVerdictDetails = "allowOverrideWithoutJustification"
	ChatMessagePolicyViolationVerdictDetailsNone                              ChatMessagePolicyViolationVerdictDetails = "none"
)

func PossibleValuesForChatMessagePolicyViolationVerdictDetails() []string {
	return []string{
		string(ChatMessagePolicyViolationVerdictDetailsAllowFalsePositiveOverride),
		string(ChatMessagePolicyViolationVerdictDetailsAllowOverrideWithJustification),
		string(ChatMessagePolicyViolationVerdictDetailsAllowOverrideWithoutJustification),
		string(ChatMessagePolicyViolationVerdictDetailsNone),
	}
}

func (s *ChatMessagePolicyViolationVerdictDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChatMessagePolicyViolationVerdictDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChatMessagePolicyViolationVerdictDetails(input string) (*ChatMessagePolicyViolationVerdictDetails, error) {
	vals := map[string]ChatMessagePolicyViolationVerdictDetails{
		"allowfalsepositiveoverride":        ChatMessagePolicyViolationVerdictDetailsAllowFalsePositiveOverride,
		"allowoverridewithjustification":    ChatMessagePolicyViolationVerdictDetailsAllowOverrideWithJustification,
		"allowoverridewithoutjustification": ChatMessagePolicyViolationVerdictDetailsAllowOverrideWithoutJustification,
		"none":                              ChatMessagePolicyViolationVerdictDetailsNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChatMessagePolicyViolationVerdictDetails(input)
	return &out, nil
}

type ColumnDefinitionType string

const (
	ColumnDefinitionTypeApprovalStatus ColumnDefinitionType = "approvalStatus"
	ColumnDefinitionTypeBoolean        ColumnDefinitionType = "boolean"
	ColumnDefinitionTypeCalculated     ColumnDefinitionType = "calculated"
	ColumnDefinitionTypeChoice         ColumnDefinitionType = "choice"
	ColumnDefinitionTypeCurrency       ColumnDefinitionType = "currency"
	ColumnDefinitionTypeDateTime       ColumnDefinitionType = "dateTime"
	ColumnDefinitionTypeGeolocation    ColumnDefinitionType = "geolocation"
	ColumnDefinitionTypeLocation       ColumnDefinitionType = "location"
	ColumnDefinitionTypeLookup         ColumnDefinitionType = "lookup"
	ColumnDefinitionTypeMultichoice    ColumnDefinitionType = "multichoice"
	ColumnDefinitionTypeMultiterm      ColumnDefinitionType = "multiterm"
	ColumnDefinitionTypeNote           ColumnDefinitionType = "note"
	ColumnDefinitionTypeNumber         ColumnDefinitionType = "number"
	ColumnDefinitionTypeTerm           ColumnDefinitionType = "term"
	ColumnDefinitionTypeText           ColumnDefinitionType = "text"
	ColumnDefinitionTypeThumbnail      ColumnDefinitionType = "thumbnail"
	ColumnDefinitionTypeUrl            ColumnDefinitionType = "url"
	ColumnDefinitionTypeUser           ColumnDefinitionType = "user"
)

func PossibleValuesForColumnDefinitionType() []string {
	return []string{
		string(ColumnDefinitionTypeApprovalStatus),
		string(ColumnDefinitionTypeBoolean),
		string(ColumnDefinitionTypeCalculated),
		string(ColumnDefinitionTypeChoice),
		string(ColumnDefinitionTypeCurrency),
		string(ColumnDefinitionTypeDateTime),
		string(ColumnDefinitionTypeGeolocation),
		string(ColumnDefinitionTypeLocation),
		string(ColumnDefinitionTypeLookup),
		string(ColumnDefinitionTypeMultichoice),
		string(ColumnDefinitionTypeMultiterm),
		string(ColumnDefinitionTypeNote),
		string(ColumnDefinitionTypeNumber),
		string(ColumnDefinitionTypeTerm),
		string(ColumnDefinitionTypeText),
		string(ColumnDefinitionTypeThumbnail),
		string(ColumnDefinitionTypeUrl),
		string(ColumnDefinitionTypeUser),
	}
}

func (s *ColumnDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseColumnDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseColumnDefinitionType(input string) (*ColumnDefinitionType, error) {
	vals := map[string]ColumnDefinitionType{
		"approvalstatus": ColumnDefinitionTypeApprovalStatus,
		"boolean":        ColumnDefinitionTypeBoolean,
		"calculated":     ColumnDefinitionTypeCalculated,
		"choice":         ColumnDefinitionTypeChoice,
		"currency":       ColumnDefinitionTypeCurrency,
		"datetime":       ColumnDefinitionTypeDateTime,
		"geolocation":    ColumnDefinitionTypeGeolocation,
		"location":       ColumnDefinitionTypeLocation,
		"lookup":         ColumnDefinitionTypeLookup,
		"multichoice":    ColumnDefinitionTypeMultichoice,
		"multiterm":      ColumnDefinitionTypeMultiterm,
		"note":           ColumnDefinitionTypeNote,
		"number":         ColumnDefinitionTypeNumber,
		"term":           ColumnDefinitionTypeTerm,
		"text":           ColumnDefinitionTypeText,
		"thumbnail":      ColumnDefinitionTypeThumbnail,
		"url":            ColumnDefinitionTypeUrl,
		"user":           ColumnDefinitionTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ColumnDefinitionType(input)
	return &out, nil
}

type DeviceActionResultActionState string

const (
	DeviceActionResultActionStateActive       DeviceActionResultActionState = "active"
	DeviceActionResultActionStateCanceled     DeviceActionResultActionState = "canceled"
	DeviceActionResultActionStateDone         DeviceActionResultActionState = "done"
	DeviceActionResultActionStateFailed       DeviceActionResultActionState = "failed"
	DeviceActionResultActionStateNone         DeviceActionResultActionState = "none"
	DeviceActionResultActionStateNotSupported DeviceActionResultActionState = "notSupported"
	DeviceActionResultActionStatePending      DeviceActionResultActionState = "pending"
)

func PossibleValuesForDeviceActionResultActionState() []string {
	return []string{
		string(DeviceActionResultActionStateActive),
		string(DeviceActionResultActionStateCanceled),
		string(DeviceActionResultActionStateDone),
		string(DeviceActionResultActionStateFailed),
		string(DeviceActionResultActionStateNone),
		string(DeviceActionResultActionStateNotSupported),
		string(DeviceActionResultActionStatePending),
	}
}

func (s *DeviceActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceActionResultActionState(input string) (*DeviceActionResultActionState, error) {
	vals := map[string]DeviceActionResultActionState{
		"active":       DeviceActionResultActionStateActive,
		"canceled":     DeviceActionResultActionStateCanceled,
		"done":         DeviceActionResultActionStateDone,
		"failed":       DeviceActionResultActionStateFailed,
		"none":         DeviceActionResultActionStateNone,
		"notsupported": DeviceActionResultActionStateNotSupported,
		"pending":      DeviceActionResultActionStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceActionResultActionState(input)
	return &out, nil
}

type DeviceCompliancePolicySettingStateState string

const (
	DeviceCompliancePolicySettingStateStateCompliant     DeviceCompliancePolicySettingStateState = "compliant"
	DeviceCompliancePolicySettingStateStateConflict      DeviceCompliancePolicySettingStateState = "conflict"
	DeviceCompliancePolicySettingStateStateError         DeviceCompliancePolicySettingStateState = "error"
	DeviceCompliancePolicySettingStateStateNonCompliant  DeviceCompliancePolicySettingStateState = "nonCompliant"
	DeviceCompliancePolicySettingStateStateNotApplicable DeviceCompliancePolicySettingStateState = "notApplicable"
	DeviceCompliancePolicySettingStateStateNotAssigned   DeviceCompliancePolicySettingStateState = "notAssigned"
	DeviceCompliancePolicySettingStateStateRemediated    DeviceCompliancePolicySettingStateState = "remediated"
	DeviceCompliancePolicySettingStateStateUnknown       DeviceCompliancePolicySettingStateState = "unknown"
)

func PossibleValuesForDeviceCompliancePolicySettingStateState() []string {
	return []string{
		string(DeviceCompliancePolicySettingStateStateCompliant),
		string(DeviceCompliancePolicySettingStateStateConflict),
		string(DeviceCompliancePolicySettingStateStateError),
		string(DeviceCompliancePolicySettingStateStateNonCompliant),
		string(DeviceCompliancePolicySettingStateStateNotApplicable),
		string(DeviceCompliancePolicySettingStateStateNotAssigned),
		string(DeviceCompliancePolicySettingStateStateRemediated),
		string(DeviceCompliancePolicySettingStateStateUnknown),
	}
}

func (s *DeviceCompliancePolicySettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicySettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicySettingStateState(input string) (*DeviceCompliancePolicySettingStateState, error) {
	vals := map[string]DeviceCompliancePolicySettingStateState{
		"compliant":     DeviceCompliancePolicySettingStateStateCompliant,
		"conflict":      DeviceCompliancePolicySettingStateStateConflict,
		"error":         DeviceCompliancePolicySettingStateStateError,
		"noncompliant":  DeviceCompliancePolicySettingStateStateNonCompliant,
		"notapplicable": DeviceCompliancePolicySettingStateStateNotApplicable,
		"notassigned":   DeviceCompliancePolicySettingStateStateNotAssigned,
		"remediated":    DeviceCompliancePolicySettingStateStateRemediated,
		"unknown":       DeviceCompliancePolicySettingStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicySettingStateState(input)
	return &out, nil
}

type DeviceCompliancePolicyStatePlatformType string

const (
	DeviceCompliancePolicyStatePlatformTypeAll               DeviceCompliancePolicyStatePlatformType = "all"
	DeviceCompliancePolicyStatePlatformTypeAndroid           DeviceCompliancePolicyStatePlatformType = "android"
	DeviceCompliancePolicyStatePlatformTypeAndroidForWork    DeviceCompliancePolicyStatePlatformType = "androidForWork"
	DeviceCompliancePolicyStatePlatformTypeIOS               DeviceCompliancePolicyStatePlatformType = "iOS"
	DeviceCompliancePolicyStatePlatformTypeMacOS             DeviceCompliancePolicyStatePlatformType = "macOS"
	DeviceCompliancePolicyStatePlatformTypeWindows10AndLater DeviceCompliancePolicyStatePlatformType = "windows10AndLater"
	DeviceCompliancePolicyStatePlatformTypeWindows81AndLater DeviceCompliancePolicyStatePlatformType = "windows81AndLater"
	DeviceCompliancePolicyStatePlatformTypeWindowsPhone81    DeviceCompliancePolicyStatePlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceCompliancePolicyStatePlatformType() []string {
	return []string{
		string(DeviceCompliancePolicyStatePlatformTypeAll),
		string(DeviceCompliancePolicyStatePlatformTypeAndroid),
		string(DeviceCompliancePolicyStatePlatformTypeAndroidForWork),
		string(DeviceCompliancePolicyStatePlatformTypeIOS),
		string(DeviceCompliancePolicyStatePlatformTypeMacOS),
		string(DeviceCompliancePolicyStatePlatformTypeWindows10AndLater),
		string(DeviceCompliancePolicyStatePlatformTypeWindows81AndLater),
		string(DeviceCompliancePolicyStatePlatformTypeWindowsPhone81),
	}
}

func (s *DeviceCompliancePolicyStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicyStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicyStatePlatformType(input string) (*DeviceCompliancePolicyStatePlatformType, error) {
	vals := map[string]DeviceCompliancePolicyStatePlatformType{
		"all":               DeviceCompliancePolicyStatePlatformTypeAll,
		"android":           DeviceCompliancePolicyStatePlatformTypeAndroid,
		"androidforwork":    DeviceCompliancePolicyStatePlatformTypeAndroidForWork,
		"ios":               DeviceCompliancePolicyStatePlatformTypeIOS,
		"macos":             DeviceCompliancePolicyStatePlatformTypeMacOS,
		"windows10andlater": DeviceCompliancePolicyStatePlatformTypeWindows10AndLater,
		"windows81andlater": DeviceCompliancePolicyStatePlatformTypeWindows81AndLater,
		"windowsphone81":    DeviceCompliancePolicyStatePlatformTypeWindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyStatePlatformType(input)
	return &out, nil
}

type DeviceCompliancePolicyStateState string

const (
	DeviceCompliancePolicyStateStateCompliant     DeviceCompliancePolicyStateState = "compliant"
	DeviceCompliancePolicyStateStateConflict      DeviceCompliancePolicyStateState = "conflict"
	DeviceCompliancePolicyStateStateError         DeviceCompliancePolicyStateState = "error"
	DeviceCompliancePolicyStateStateNonCompliant  DeviceCompliancePolicyStateState = "nonCompliant"
	DeviceCompliancePolicyStateStateNotApplicable DeviceCompliancePolicyStateState = "notApplicable"
	DeviceCompliancePolicyStateStateNotAssigned   DeviceCompliancePolicyStateState = "notAssigned"
	DeviceCompliancePolicyStateStateRemediated    DeviceCompliancePolicyStateState = "remediated"
	DeviceCompliancePolicyStateStateUnknown       DeviceCompliancePolicyStateState = "unknown"
)

func PossibleValuesForDeviceCompliancePolicyStateState() []string {
	return []string{
		string(DeviceCompliancePolicyStateStateCompliant),
		string(DeviceCompliancePolicyStateStateConflict),
		string(DeviceCompliancePolicyStateStateError),
		string(DeviceCompliancePolicyStateStateNonCompliant),
		string(DeviceCompliancePolicyStateStateNotApplicable),
		string(DeviceCompliancePolicyStateStateNotAssigned),
		string(DeviceCompliancePolicyStateStateRemediated),
		string(DeviceCompliancePolicyStateStateUnknown),
	}
}

func (s *DeviceCompliancePolicyStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicyStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicyStateState(input string) (*DeviceCompliancePolicyStateState, error) {
	vals := map[string]DeviceCompliancePolicyStateState{
		"compliant":     DeviceCompliancePolicyStateStateCompliant,
		"conflict":      DeviceCompliancePolicyStateStateConflict,
		"error":         DeviceCompliancePolicyStateStateError,
		"noncompliant":  DeviceCompliancePolicyStateStateNonCompliant,
		"notapplicable": DeviceCompliancePolicyStateStateNotApplicable,
		"notassigned":   DeviceCompliancePolicyStateStateNotAssigned,
		"remediated":    DeviceCompliancePolicyStateStateRemediated,
		"unknown":       DeviceCompliancePolicyStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyStateState(input)
	return &out, nil
}

type DeviceConfigurationSettingStateState string

const (
	DeviceConfigurationSettingStateStateCompliant     DeviceConfigurationSettingStateState = "compliant"
	DeviceConfigurationSettingStateStateConflict      DeviceConfigurationSettingStateState = "conflict"
	DeviceConfigurationSettingStateStateError         DeviceConfigurationSettingStateState = "error"
	DeviceConfigurationSettingStateStateNonCompliant  DeviceConfigurationSettingStateState = "nonCompliant"
	DeviceConfigurationSettingStateStateNotApplicable DeviceConfigurationSettingStateState = "notApplicable"
	DeviceConfigurationSettingStateStateNotAssigned   DeviceConfigurationSettingStateState = "notAssigned"
	DeviceConfigurationSettingStateStateRemediated    DeviceConfigurationSettingStateState = "remediated"
	DeviceConfigurationSettingStateStateUnknown       DeviceConfigurationSettingStateState = "unknown"
)

func PossibleValuesForDeviceConfigurationSettingStateState() []string {
	return []string{
		string(DeviceConfigurationSettingStateStateCompliant),
		string(DeviceConfigurationSettingStateStateConflict),
		string(DeviceConfigurationSettingStateStateError),
		string(DeviceConfigurationSettingStateStateNonCompliant),
		string(DeviceConfigurationSettingStateStateNotApplicable),
		string(DeviceConfigurationSettingStateStateNotAssigned),
		string(DeviceConfigurationSettingStateStateRemediated),
		string(DeviceConfigurationSettingStateStateUnknown),
	}
}

func (s *DeviceConfigurationSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationSettingStateState(input string) (*DeviceConfigurationSettingStateState, error) {
	vals := map[string]DeviceConfigurationSettingStateState{
		"compliant":     DeviceConfigurationSettingStateStateCompliant,
		"conflict":      DeviceConfigurationSettingStateStateConflict,
		"error":         DeviceConfigurationSettingStateStateError,
		"noncompliant":  DeviceConfigurationSettingStateStateNonCompliant,
		"notapplicable": DeviceConfigurationSettingStateStateNotApplicable,
		"notassigned":   DeviceConfigurationSettingStateStateNotAssigned,
		"remediated":    DeviceConfigurationSettingStateStateRemediated,
		"unknown":       DeviceConfigurationSettingStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationSettingStateState(input)
	return &out, nil
}

type DeviceConfigurationStatePlatformType string

const (
	DeviceConfigurationStatePlatformTypeAll               DeviceConfigurationStatePlatformType = "all"
	DeviceConfigurationStatePlatformTypeAndroid           DeviceConfigurationStatePlatformType = "android"
	DeviceConfigurationStatePlatformTypeAndroidForWork    DeviceConfigurationStatePlatformType = "androidForWork"
	DeviceConfigurationStatePlatformTypeIOS               DeviceConfigurationStatePlatformType = "iOS"
	DeviceConfigurationStatePlatformTypeMacOS             DeviceConfigurationStatePlatformType = "macOS"
	DeviceConfigurationStatePlatformTypeWindows10AndLater DeviceConfigurationStatePlatformType = "windows10AndLater"
	DeviceConfigurationStatePlatformTypeWindows81AndLater DeviceConfigurationStatePlatformType = "windows81AndLater"
	DeviceConfigurationStatePlatformTypeWindowsPhone81    DeviceConfigurationStatePlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceConfigurationStatePlatformType() []string {
	return []string{
		string(DeviceConfigurationStatePlatformTypeAll),
		string(DeviceConfigurationStatePlatformTypeAndroid),
		string(DeviceConfigurationStatePlatformTypeAndroidForWork),
		string(DeviceConfigurationStatePlatformTypeIOS),
		string(DeviceConfigurationStatePlatformTypeMacOS),
		string(DeviceConfigurationStatePlatformTypeWindows10AndLater),
		string(DeviceConfigurationStatePlatformTypeWindows81AndLater),
		string(DeviceConfigurationStatePlatformTypeWindowsPhone81),
	}
}

func (s *DeviceConfigurationStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationStatePlatformType(input string) (*DeviceConfigurationStatePlatformType, error) {
	vals := map[string]DeviceConfigurationStatePlatformType{
		"all":               DeviceConfigurationStatePlatformTypeAll,
		"android":           DeviceConfigurationStatePlatformTypeAndroid,
		"androidforwork":    DeviceConfigurationStatePlatformTypeAndroidForWork,
		"ios":               DeviceConfigurationStatePlatformTypeIOS,
		"macos":             DeviceConfigurationStatePlatformTypeMacOS,
		"windows10andlater": DeviceConfigurationStatePlatformTypeWindows10AndLater,
		"windows81andlater": DeviceConfigurationStatePlatformTypeWindows81AndLater,
		"windowsphone81":    DeviceConfigurationStatePlatformTypeWindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationStatePlatformType(input)
	return &out, nil
}

type DeviceConfigurationStateState string

const (
	DeviceConfigurationStateStateCompliant     DeviceConfigurationStateState = "compliant"
	DeviceConfigurationStateStateConflict      DeviceConfigurationStateState = "conflict"
	DeviceConfigurationStateStateError         DeviceConfigurationStateState = "error"
	DeviceConfigurationStateStateNonCompliant  DeviceConfigurationStateState = "nonCompliant"
	DeviceConfigurationStateStateNotApplicable DeviceConfigurationStateState = "notApplicable"
	DeviceConfigurationStateStateNotAssigned   DeviceConfigurationStateState = "notAssigned"
	DeviceConfigurationStateStateRemediated    DeviceConfigurationStateState = "remediated"
	DeviceConfigurationStateStateUnknown       DeviceConfigurationStateState = "unknown"
)

func PossibleValuesForDeviceConfigurationStateState() []string {
	return []string{
		string(DeviceConfigurationStateStateCompliant),
		string(DeviceConfigurationStateStateConflict),
		string(DeviceConfigurationStateStateError),
		string(DeviceConfigurationStateStateNonCompliant),
		string(DeviceConfigurationStateStateNotApplicable),
		string(DeviceConfigurationStateStateNotAssigned),
		string(DeviceConfigurationStateStateRemediated),
		string(DeviceConfigurationStateStateUnknown),
	}
}

func (s *DeviceConfigurationStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationStateState(input string) (*DeviceConfigurationStateState, error) {
	vals := map[string]DeviceConfigurationStateState{
		"compliant":     DeviceConfigurationStateStateCompliant,
		"conflict":      DeviceConfigurationStateStateConflict,
		"error":         DeviceConfigurationStateStateError,
		"noncompliant":  DeviceConfigurationStateStateNonCompliant,
		"notapplicable": DeviceConfigurationStateStateNotApplicable,
		"notassigned":   DeviceConfigurationStateStateNotAssigned,
		"remediated":    DeviceConfigurationStateStateRemediated,
		"unknown":       DeviceConfigurationStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationStateState(input)
	return &out, nil
}

type DeviceLogCollectionResponseStatus string

const (
	DeviceLogCollectionResponseStatusCompleted DeviceLogCollectionResponseStatus = "completed"
	DeviceLogCollectionResponseStatusFailed    DeviceLogCollectionResponseStatus = "failed"
	DeviceLogCollectionResponseStatusPending   DeviceLogCollectionResponseStatus = "pending"
)

func PossibleValuesForDeviceLogCollectionResponseStatus() []string {
	return []string{
		string(DeviceLogCollectionResponseStatusCompleted),
		string(DeviceLogCollectionResponseStatusFailed),
		string(DeviceLogCollectionResponseStatusPending),
	}
}

func (s *DeviceLogCollectionResponseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceLogCollectionResponseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceLogCollectionResponseStatus(input string) (*DeviceLogCollectionResponseStatus, error) {
	vals := map[string]DeviceLogCollectionResponseStatus{
		"completed": DeviceLogCollectionResponseStatusCompleted,
		"failed":    DeviceLogCollectionResponseStatusFailed,
		"pending":   DeviceLogCollectionResponseStatusPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLogCollectionResponseStatus(input)
	return &out, nil
}

type EventImportance string

const (
	EventImportanceHigh   EventImportance = "high"
	EventImportanceLow    EventImportance = "low"
	EventImportanceNormal EventImportance = "normal"
)

func PossibleValuesForEventImportance() []string {
	return []string{
		string(EventImportanceHigh),
		string(EventImportanceLow),
		string(EventImportanceNormal),
	}
}

func (s *EventImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventImportance(input string) (*EventImportance, error) {
	vals := map[string]EventImportance{
		"high":   EventImportanceHigh,
		"low":    EventImportanceLow,
		"normal": EventImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventImportance(input)
	return &out, nil
}

type EventOnlineMeetingProvider string

const (
	EventOnlineMeetingProviderSkypeForBusiness EventOnlineMeetingProvider = "skypeForBusiness"
	EventOnlineMeetingProviderSkypeForConsumer EventOnlineMeetingProvider = "skypeForConsumer"
	EventOnlineMeetingProviderTeamsForBusiness EventOnlineMeetingProvider = "teamsForBusiness"
	EventOnlineMeetingProviderUnknown          EventOnlineMeetingProvider = "unknown"
)

func PossibleValuesForEventOnlineMeetingProvider() []string {
	return []string{
		string(EventOnlineMeetingProviderSkypeForBusiness),
		string(EventOnlineMeetingProviderSkypeForConsumer),
		string(EventOnlineMeetingProviderTeamsForBusiness),
		string(EventOnlineMeetingProviderUnknown),
	}
}

func (s *EventOnlineMeetingProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventOnlineMeetingProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventOnlineMeetingProvider(input string) (*EventOnlineMeetingProvider, error) {
	vals := map[string]EventOnlineMeetingProvider{
		"skypeforbusiness": EventOnlineMeetingProviderSkypeForBusiness,
		"skypeforconsumer": EventOnlineMeetingProviderSkypeForConsumer,
		"teamsforbusiness": EventOnlineMeetingProviderTeamsForBusiness,
		"unknown":          EventOnlineMeetingProviderUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventOnlineMeetingProvider(input)
	return &out, nil
}

type EventSensitivity string

const (
	EventSensitivityConfidential EventSensitivity = "confidential"
	EventSensitivityNormal       EventSensitivity = "normal"
	EventSensitivityPersonal     EventSensitivity = "personal"
	EventSensitivityPrivate      EventSensitivity = "private"
)

func PossibleValuesForEventSensitivity() []string {
	return []string{
		string(EventSensitivityConfidential),
		string(EventSensitivityNormal),
		string(EventSensitivityPersonal),
		string(EventSensitivityPrivate),
	}
}

func (s *EventSensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventSensitivity(input string) (*EventSensitivity, error) {
	vals := map[string]EventSensitivity{
		"confidential": EventSensitivityConfidential,
		"normal":       EventSensitivityNormal,
		"personal":     EventSensitivityPersonal,
		"private":      EventSensitivityPrivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSensitivity(input)
	return &out, nil
}

type EventShowAs string

const (
	EventShowAsBusy             EventShowAs = "busy"
	EventShowAsFree             EventShowAs = "free"
	EventShowAsOof              EventShowAs = "oof"
	EventShowAsTentative        EventShowAs = "tentative"
	EventShowAsUnknown          EventShowAs = "unknown"
	EventShowAsWorkingElsewhere EventShowAs = "workingElsewhere"
)

func PossibleValuesForEventShowAs() []string {
	return []string{
		string(EventShowAsBusy),
		string(EventShowAsFree),
		string(EventShowAsOof),
		string(EventShowAsTentative),
		string(EventShowAsUnknown),
		string(EventShowAsWorkingElsewhere),
	}
}

func (s *EventShowAs) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventShowAs(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventShowAs(input string) (*EventShowAs, error) {
	vals := map[string]EventShowAs{
		"busy":             EventShowAsBusy,
		"free":             EventShowAsFree,
		"oof":              EventShowAsOof,
		"tentative":        EventShowAsTentative,
		"unknown":          EventShowAsUnknown,
		"workingelsewhere": EventShowAsWorkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventShowAs(input)
	return &out, nil
}

type EventType string

const (
	EventTypeException      EventType = "exception"
	EventTypeOccurrence     EventType = "occurrence"
	EventTypeSeriesMaster   EventType = "seriesMaster"
	EventTypeSingleInstance EventType = "singleInstance"
)

func PossibleValuesForEventType() []string {
	return []string{
		string(EventTypeException),
		string(EventTypeOccurrence),
		string(EventTypeSeriesMaster),
		string(EventTypeSingleInstance),
	}
}

func (s *EventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventType(input string) (*EventType, error) {
	vals := map[string]EventType{
		"exception":      EventTypeException,
		"occurrence":     EventTypeOccurrence,
		"seriesmaster":   EventTypeSeriesMaster,
		"singleinstance": EventTypeSingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventType(input)
	return &out, nil
}

type Fido2AuthenticationMethodAttestationLevel string

const (
	Fido2AuthenticationMethodAttestationLevelAttested    Fido2AuthenticationMethodAttestationLevel = "attested"
	Fido2AuthenticationMethodAttestationLevelNotAttested Fido2AuthenticationMethodAttestationLevel = "notAttested"
)

func PossibleValuesForFido2AuthenticationMethodAttestationLevel() []string {
	return []string{
		string(Fido2AuthenticationMethodAttestationLevelAttested),
		string(Fido2AuthenticationMethodAttestationLevelNotAttested),
	}
}

func (s *Fido2AuthenticationMethodAttestationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFido2AuthenticationMethodAttestationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFido2AuthenticationMethodAttestationLevel(input string) (*Fido2AuthenticationMethodAttestationLevel, error) {
	vals := map[string]Fido2AuthenticationMethodAttestationLevel{
		"attested":    Fido2AuthenticationMethodAttestationLevelAttested,
		"notattested": Fido2AuthenticationMethodAttestationLevelNotAttested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2AuthenticationMethodAttestationLevel(input)
	return &out, nil
}

type FollowupFlagFlagStatus string

const (
	FollowupFlagFlagStatusComplete   FollowupFlagFlagStatus = "complete"
	FollowupFlagFlagStatusFlagged    FollowupFlagFlagStatus = "flagged"
	FollowupFlagFlagStatusNotFlagged FollowupFlagFlagStatus = "notFlagged"
)

func PossibleValuesForFollowupFlagFlagStatus() []string {
	return []string{
		string(FollowupFlagFlagStatusComplete),
		string(FollowupFlagFlagStatusFlagged),
		string(FollowupFlagFlagStatusNotFlagged),
	}
}

func (s *FollowupFlagFlagStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFollowupFlagFlagStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFollowupFlagFlagStatus(input string) (*FollowupFlagFlagStatus, error) {
	vals := map[string]FollowupFlagFlagStatus{
		"complete":   FollowupFlagFlagStatusComplete,
		"flagged":    FollowupFlagFlagStatusFlagged,
		"notflagged": FollowupFlagFlagStatusNotFlagged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FollowupFlagFlagStatus(input)
	return &out, nil
}

type IdentityGovernanceParameterValueType string

const (
	IdentityGovernanceParameterValueTypeBool   IdentityGovernanceParameterValueType = "bool"
	IdentityGovernanceParameterValueTypeEnum   IdentityGovernanceParameterValueType = "enum"
	IdentityGovernanceParameterValueTypeInt    IdentityGovernanceParameterValueType = "int"
	IdentityGovernanceParameterValueTypeString IdentityGovernanceParameterValueType = "string"
)

func PossibleValuesForIdentityGovernanceParameterValueType() []string {
	return []string{
		string(IdentityGovernanceParameterValueTypeBool),
		string(IdentityGovernanceParameterValueTypeEnum),
		string(IdentityGovernanceParameterValueTypeInt),
		string(IdentityGovernanceParameterValueTypeString),
	}
}

func (s *IdentityGovernanceParameterValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceParameterValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceParameterValueType(input string) (*IdentityGovernanceParameterValueType, error) {
	vals := map[string]IdentityGovernanceParameterValueType{
		"bool":   IdentityGovernanceParameterValueTypeBool,
		"enum":   IdentityGovernanceParameterValueTypeEnum,
		"int":    IdentityGovernanceParameterValueTypeInt,
		"string": IdentityGovernanceParameterValueTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceParameterValueType(input)
	return &out, nil
}

type IdentityGovernanceRunProcessingStatus string

const (
	IdentityGovernanceRunProcessingStatusCanceled            IdentityGovernanceRunProcessingStatus = "canceled"
	IdentityGovernanceRunProcessingStatusCompleted           IdentityGovernanceRunProcessingStatus = "completed"
	IdentityGovernanceRunProcessingStatusCompletedWithErrors IdentityGovernanceRunProcessingStatus = "completedWithErrors"
	IdentityGovernanceRunProcessingStatusFailed              IdentityGovernanceRunProcessingStatus = "failed"
	IdentityGovernanceRunProcessingStatusInProgress          IdentityGovernanceRunProcessingStatus = "inProgress"
	IdentityGovernanceRunProcessingStatusQueued              IdentityGovernanceRunProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceRunProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceRunProcessingStatusCanceled),
		string(IdentityGovernanceRunProcessingStatusCompleted),
		string(IdentityGovernanceRunProcessingStatusCompletedWithErrors),
		string(IdentityGovernanceRunProcessingStatusFailed),
		string(IdentityGovernanceRunProcessingStatusInProgress),
		string(IdentityGovernanceRunProcessingStatusQueued),
	}
}

func (s *IdentityGovernanceRunProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceRunProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceRunProcessingStatus(input string) (*IdentityGovernanceRunProcessingStatus, error) {
	vals := map[string]IdentityGovernanceRunProcessingStatus{
		"canceled":            IdentityGovernanceRunProcessingStatusCanceled,
		"completed":           IdentityGovernanceRunProcessingStatusCompleted,
		"completedwitherrors": IdentityGovernanceRunProcessingStatusCompletedWithErrors,
		"failed":              IdentityGovernanceRunProcessingStatusFailed,
		"inprogress":          IdentityGovernanceRunProcessingStatusInProgress,
		"queued":              IdentityGovernanceRunProcessingStatusQueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceRunProcessingStatus(input)
	return &out, nil
}

type IdentityGovernanceRunWorkflowExecutionType string

const (
	IdentityGovernanceRunWorkflowExecutionTypeOnDemand  IdentityGovernanceRunWorkflowExecutionType = "onDemand"
	IdentityGovernanceRunWorkflowExecutionTypeScheduled IdentityGovernanceRunWorkflowExecutionType = "scheduled"
)

func PossibleValuesForIdentityGovernanceRunWorkflowExecutionType() []string {
	return []string{
		string(IdentityGovernanceRunWorkflowExecutionTypeOnDemand),
		string(IdentityGovernanceRunWorkflowExecutionTypeScheduled),
	}
}

func (s *IdentityGovernanceRunWorkflowExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceRunWorkflowExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceRunWorkflowExecutionType(input string) (*IdentityGovernanceRunWorkflowExecutionType, error) {
	vals := map[string]IdentityGovernanceRunWorkflowExecutionType{
		"ondemand":  IdentityGovernanceRunWorkflowExecutionTypeOnDemand,
		"scheduled": IdentityGovernanceRunWorkflowExecutionTypeScheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceRunWorkflowExecutionType(input)
	return &out, nil
}

type IdentityGovernanceTaskCategory string

const (
	IdentityGovernanceTaskCategoryJoiner IdentityGovernanceTaskCategory = "joiner"
	IdentityGovernanceTaskCategoryLeaver IdentityGovernanceTaskCategory = "leaver"
	IdentityGovernanceTaskCategoryMover  IdentityGovernanceTaskCategory = "mover"
)

func PossibleValuesForIdentityGovernanceTaskCategory() []string {
	return []string{
		string(IdentityGovernanceTaskCategoryJoiner),
		string(IdentityGovernanceTaskCategoryLeaver),
		string(IdentityGovernanceTaskCategoryMover),
	}
}

func (s *IdentityGovernanceTaskCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskCategory(input string) (*IdentityGovernanceTaskCategory, error) {
	vals := map[string]IdentityGovernanceTaskCategory{
		"joiner": IdentityGovernanceTaskCategoryJoiner,
		"leaver": IdentityGovernanceTaskCategoryLeaver,
		"mover":  IdentityGovernanceTaskCategoryMover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskCategory(input)
	return &out, nil
}

type IdentityGovernanceTaskDefinitionCategory string

const (
	IdentityGovernanceTaskDefinitionCategoryJoiner IdentityGovernanceTaskDefinitionCategory = "joiner"
	IdentityGovernanceTaskDefinitionCategoryLeaver IdentityGovernanceTaskDefinitionCategory = "leaver"
	IdentityGovernanceTaskDefinitionCategoryMover  IdentityGovernanceTaskDefinitionCategory = "mover"
)

func PossibleValuesForIdentityGovernanceTaskDefinitionCategory() []string {
	return []string{
		string(IdentityGovernanceTaskDefinitionCategoryJoiner),
		string(IdentityGovernanceTaskDefinitionCategoryLeaver),
		string(IdentityGovernanceTaskDefinitionCategoryMover),
	}
}

func (s *IdentityGovernanceTaskDefinitionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskDefinitionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskDefinitionCategory(input string) (*IdentityGovernanceTaskDefinitionCategory, error) {
	vals := map[string]IdentityGovernanceTaskDefinitionCategory{
		"joiner": IdentityGovernanceTaskDefinitionCategoryJoiner,
		"leaver": IdentityGovernanceTaskDefinitionCategoryLeaver,
		"mover":  IdentityGovernanceTaskDefinitionCategoryMover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskDefinitionCategory(input)
	return &out, nil
}

type IdentityGovernanceTaskProcessingResultProcessingStatus string

const (
	IdentityGovernanceTaskProcessingResultProcessingStatusCanceled            IdentityGovernanceTaskProcessingResultProcessingStatus = "canceled"
	IdentityGovernanceTaskProcessingResultProcessingStatusCompleted           IdentityGovernanceTaskProcessingResultProcessingStatus = "completed"
	IdentityGovernanceTaskProcessingResultProcessingStatusCompletedWithErrors IdentityGovernanceTaskProcessingResultProcessingStatus = "completedWithErrors"
	IdentityGovernanceTaskProcessingResultProcessingStatusFailed              IdentityGovernanceTaskProcessingResultProcessingStatus = "failed"
	IdentityGovernanceTaskProcessingResultProcessingStatusInProgress          IdentityGovernanceTaskProcessingResultProcessingStatus = "inProgress"
	IdentityGovernanceTaskProcessingResultProcessingStatusQueued              IdentityGovernanceTaskProcessingResultProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceTaskProcessingResultProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceTaskProcessingResultProcessingStatusCanceled),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusCompleted),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusCompletedWithErrors),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusFailed),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusInProgress),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusQueued),
	}
}

func (s *IdentityGovernanceTaskProcessingResultProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskProcessingResultProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskProcessingResultProcessingStatus(input string) (*IdentityGovernanceTaskProcessingResultProcessingStatus, error) {
	vals := map[string]IdentityGovernanceTaskProcessingResultProcessingStatus{
		"canceled":            IdentityGovernanceTaskProcessingResultProcessingStatusCanceled,
		"completed":           IdentityGovernanceTaskProcessingResultProcessingStatusCompleted,
		"completedwitherrors": IdentityGovernanceTaskProcessingResultProcessingStatusCompletedWithErrors,
		"failed":              IdentityGovernanceTaskProcessingResultProcessingStatusFailed,
		"inprogress":          IdentityGovernanceTaskProcessingResultProcessingStatusInProgress,
		"queued":              IdentityGovernanceTaskProcessingResultProcessingStatusQueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskProcessingResultProcessingStatus(input)
	return &out, nil
}

type IdentityGovernanceTaskReportProcessingStatus string

const (
	IdentityGovernanceTaskReportProcessingStatusCanceled            IdentityGovernanceTaskReportProcessingStatus = "canceled"
	IdentityGovernanceTaskReportProcessingStatusCompleted           IdentityGovernanceTaskReportProcessingStatus = "completed"
	IdentityGovernanceTaskReportProcessingStatusCompletedWithErrors IdentityGovernanceTaskReportProcessingStatus = "completedWithErrors"
	IdentityGovernanceTaskReportProcessingStatusFailed              IdentityGovernanceTaskReportProcessingStatus = "failed"
	IdentityGovernanceTaskReportProcessingStatusInProgress          IdentityGovernanceTaskReportProcessingStatus = "inProgress"
	IdentityGovernanceTaskReportProcessingStatusQueued              IdentityGovernanceTaskReportProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceTaskReportProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceTaskReportProcessingStatusCanceled),
		string(IdentityGovernanceTaskReportProcessingStatusCompleted),
		string(IdentityGovernanceTaskReportProcessingStatusCompletedWithErrors),
		string(IdentityGovernanceTaskReportProcessingStatusFailed),
		string(IdentityGovernanceTaskReportProcessingStatusInProgress),
		string(IdentityGovernanceTaskReportProcessingStatusQueued),
	}
}

func (s *IdentityGovernanceTaskReportProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskReportProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskReportProcessingStatus(input string) (*IdentityGovernanceTaskReportProcessingStatus, error) {
	vals := map[string]IdentityGovernanceTaskReportProcessingStatus{
		"canceled":            IdentityGovernanceTaskReportProcessingStatusCanceled,
		"completed":           IdentityGovernanceTaskReportProcessingStatusCompleted,
		"completedwitherrors": IdentityGovernanceTaskReportProcessingStatusCompletedWithErrors,
		"failed":              IdentityGovernanceTaskReportProcessingStatusFailed,
		"inprogress":          IdentityGovernanceTaskReportProcessingStatusInProgress,
		"queued":              IdentityGovernanceTaskReportProcessingStatusQueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskReportProcessingStatus(input)
	return &out, nil
}

type IdentityGovernanceUserProcessingResultProcessingStatus string

const (
	IdentityGovernanceUserProcessingResultProcessingStatusCanceled            IdentityGovernanceUserProcessingResultProcessingStatus = "canceled"
	IdentityGovernanceUserProcessingResultProcessingStatusCompleted           IdentityGovernanceUserProcessingResultProcessingStatus = "completed"
	IdentityGovernanceUserProcessingResultProcessingStatusCompletedWithErrors IdentityGovernanceUserProcessingResultProcessingStatus = "completedWithErrors"
	IdentityGovernanceUserProcessingResultProcessingStatusFailed              IdentityGovernanceUserProcessingResultProcessingStatus = "failed"
	IdentityGovernanceUserProcessingResultProcessingStatusInProgress          IdentityGovernanceUserProcessingResultProcessingStatus = "inProgress"
	IdentityGovernanceUserProcessingResultProcessingStatusQueued              IdentityGovernanceUserProcessingResultProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceUserProcessingResultProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceUserProcessingResultProcessingStatusCanceled),
		string(IdentityGovernanceUserProcessingResultProcessingStatusCompleted),
		string(IdentityGovernanceUserProcessingResultProcessingStatusCompletedWithErrors),
		string(IdentityGovernanceUserProcessingResultProcessingStatusFailed),
		string(IdentityGovernanceUserProcessingResultProcessingStatusInProgress),
		string(IdentityGovernanceUserProcessingResultProcessingStatusQueued),
	}
}

func (s *IdentityGovernanceUserProcessingResultProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceUserProcessingResultProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceUserProcessingResultProcessingStatus(input string) (*IdentityGovernanceUserProcessingResultProcessingStatus, error) {
	vals := map[string]IdentityGovernanceUserProcessingResultProcessingStatus{
		"canceled":            IdentityGovernanceUserProcessingResultProcessingStatusCanceled,
		"completed":           IdentityGovernanceUserProcessingResultProcessingStatusCompleted,
		"completedwitherrors": IdentityGovernanceUserProcessingResultProcessingStatusCompletedWithErrors,
		"failed":              IdentityGovernanceUserProcessingResultProcessingStatusFailed,
		"inprogress":          IdentityGovernanceUserProcessingResultProcessingStatusInProgress,
		"queued":              IdentityGovernanceUserProcessingResultProcessingStatusQueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceUserProcessingResultProcessingStatus(input)
	return &out, nil
}

type IdentityGovernanceUserProcessingResultWorkflowExecutionType string

const (
	IdentityGovernanceUserProcessingResultWorkflowExecutionTypeOnDemand  IdentityGovernanceUserProcessingResultWorkflowExecutionType = "onDemand"
	IdentityGovernanceUserProcessingResultWorkflowExecutionTypeScheduled IdentityGovernanceUserProcessingResultWorkflowExecutionType = "scheduled"
)

func PossibleValuesForIdentityGovernanceUserProcessingResultWorkflowExecutionType() []string {
	return []string{
		string(IdentityGovernanceUserProcessingResultWorkflowExecutionTypeOnDemand),
		string(IdentityGovernanceUserProcessingResultWorkflowExecutionTypeScheduled),
	}
}

func (s *IdentityGovernanceUserProcessingResultWorkflowExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceUserProcessingResultWorkflowExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceUserProcessingResultWorkflowExecutionType(input string) (*IdentityGovernanceUserProcessingResultWorkflowExecutionType, error) {
	vals := map[string]IdentityGovernanceUserProcessingResultWorkflowExecutionType{
		"ondemand":  IdentityGovernanceUserProcessingResultWorkflowExecutionTypeOnDemand,
		"scheduled": IdentityGovernanceUserProcessingResultWorkflowExecutionTypeScheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceUserProcessingResultWorkflowExecutionType(input)
	return &out, nil
}

type IdentityGovernanceWorkflowCategory string

const (
	IdentityGovernanceWorkflowCategoryJoiner IdentityGovernanceWorkflowCategory = "joiner"
	IdentityGovernanceWorkflowCategoryLeaver IdentityGovernanceWorkflowCategory = "leaver"
	IdentityGovernanceWorkflowCategoryMover  IdentityGovernanceWorkflowCategory = "mover"
)

func PossibleValuesForIdentityGovernanceWorkflowCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowCategoryJoiner),
		string(IdentityGovernanceWorkflowCategoryLeaver),
		string(IdentityGovernanceWorkflowCategoryMover),
	}
}

func (s *IdentityGovernanceWorkflowCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowCategory(input string) (*IdentityGovernanceWorkflowCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowCategory{
		"joiner": IdentityGovernanceWorkflowCategoryJoiner,
		"leaver": IdentityGovernanceWorkflowCategoryLeaver,
		"mover":  IdentityGovernanceWorkflowCategoryMover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowCategory(input)
	return &out, nil
}

type IdentityGovernanceWorkflowVersionCategory string

const (
	IdentityGovernanceWorkflowVersionCategoryJoiner IdentityGovernanceWorkflowVersionCategory = "joiner"
	IdentityGovernanceWorkflowVersionCategoryLeaver IdentityGovernanceWorkflowVersionCategory = "leaver"
	IdentityGovernanceWorkflowVersionCategoryMover  IdentityGovernanceWorkflowVersionCategory = "mover"
)

func PossibleValuesForIdentityGovernanceWorkflowVersionCategory() []string {
	return []string{
		string(IdentityGovernanceWorkflowVersionCategoryJoiner),
		string(IdentityGovernanceWorkflowVersionCategoryLeaver),
		string(IdentityGovernanceWorkflowVersionCategoryMover),
	}
}

func (s *IdentityGovernanceWorkflowVersionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowVersionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowVersionCategory(input string) (*IdentityGovernanceWorkflowVersionCategory, error) {
	vals := map[string]IdentityGovernanceWorkflowVersionCategory{
		"joiner": IdentityGovernanceWorkflowVersionCategoryJoiner,
		"leaver": IdentityGovernanceWorkflowVersionCategoryLeaver,
		"mover":  IdentityGovernanceWorkflowVersionCategoryMover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowVersionCategory(input)
	return &out, nil
}

type InferenceClassificationOverrideClassifyAs string

const (
	InferenceClassificationOverrideClassifyAsFocused InferenceClassificationOverrideClassifyAs = "focused"
	InferenceClassificationOverrideClassifyAsOther   InferenceClassificationOverrideClassifyAs = "other"
)

func PossibleValuesForInferenceClassificationOverrideClassifyAs() []string {
	return []string{
		string(InferenceClassificationOverrideClassifyAsFocused),
		string(InferenceClassificationOverrideClassifyAsOther),
	}
}

func (s *InferenceClassificationOverrideClassifyAs) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInferenceClassificationOverrideClassifyAs(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInferenceClassificationOverrideClassifyAs(input string) (*InferenceClassificationOverrideClassifyAs, error) {
	vals := map[string]InferenceClassificationOverrideClassifyAs{
		"focused": InferenceClassificationOverrideClassifyAsFocused,
		"other":   InferenceClassificationOverrideClassifyAsOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InferenceClassificationOverrideClassifyAs(input)
	return &out, nil
}

type ItemBodyContentType string

const (
	ItemBodyContentTypeHtml ItemBodyContentType = "html"
	ItemBodyContentTypeText ItemBodyContentType = "text"
)

func PossibleValuesForItemBodyContentType() []string {
	return []string{
		string(ItemBodyContentTypeHtml),
		string(ItemBodyContentTypeText),
	}
}

func (s *ItemBodyContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemBodyContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemBodyContentType(input string) (*ItemBodyContentType, error) {
	vals := map[string]ItemBodyContentType{
		"html": ItemBodyContentTypeHtml,
		"text": ItemBodyContentTypeText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemBodyContentType(input)
	return &out, nil
}

type LearningCourseActivityStatus string

const (
	LearningCourseActivityStatusCompleted  LearningCourseActivityStatus = "completed"
	LearningCourseActivityStatusInProgress LearningCourseActivityStatus = "inProgress"
	LearningCourseActivityStatusNotStarted LearningCourseActivityStatus = "notStarted"
)

func PossibleValuesForLearningCourseActivityStatus() []string {
	return []string{
		string(LearningCourseActivityStatusCompleted),
		string(LearningCourseActivityStatusInProgress),
		string(LearningCourseActivityStatusNotStarted),
	}
}

func (s *LearningCourseActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLearningCourseActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLearningCourseActivityStatus(input string) (*LearningCourseActivityStatus, error) {
	vals := map[string]LearningCourseActivityStatus{
		"completed":  LearningCourseActivityStatusCompleted,
		"inprogress": LearningCourseActivityStatusInProgress,
		"notstarted": LearningCourseActivityStatusNotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningCourseActivityStatus(input)
	return &out, nil
}

type LobbyBypassSettingsScope string

const (
	LobbyBypassSettingsScopeEveryone                    LobbyBypassSettingsScope = "everyone"
	LobbyBypassSettingsScopeInvited                     LobbyBypassSettingsScope = "invited"
	LobbyBypassSettingsScopeOrganization                LobbyBypassSettingsScope = "organization"
	LobbyBypassSettingsScopeOrganizationAndFederated    LobbyBypassSettingsScope = "organizationAndFederated"
	LobbyBypassSettingsScopeOrganizationExcludingGuests LobbyBypassSettingsScope = "organizationExcludingGuests"
	LobbyBypassSettingsScopeOrganizer                   LobbyBypassSettingsScope = "organizer"
)

func PossibleValuesForLobbyBypassSettingsScope() []string {
	return []string{
		string(LobbyBypassSettingsScopeEveryone),
		string(LobbyBypassSettingsScopeInvited),
		string(LobbyBypassSettingsScopeOrganization),
		string(LobbyBypassSettingsScopeOrganizationAndFederated),
		string(LobbyBypassSettingsScopeOrganizationExcludingGuests),
		string(LobbyBypassSettingsScopeOrganizer),
	}
}

func (s *LobbyBypassSettingsScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLobbyBypassSettingsScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLobbyBypassSettingsScope(input string) (*LobbyBypassSettingsScope, error) {
	vals := map[string]LobbyBypassSettingsScope{
		"everyone":                    LobbyBypassSettingsScopeEveryone,
		"invited":                     LobbyBypassSettingsScopeInvited,
		"organization":                LobbyBypassSettingsScopeOrganization,
		"organizationandfederated":    LobbyBypassSettingsScopeOrganizationAndFederated,
		"organizationexcludingguests": LobbyBypassSettingsScopeOrganizationExcludingGuests,
		"organizer":                   LobbyBypassSettingsScopeOrganizer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LobbyBypassSettingsScope(input)
	return &out, nil
}

type LocationLocationType string

const (
	LocationLocationTypeBusinessAddress LocationLocationType = "businessAddress"
	LocationLocationTypeConferenceRoom  LocationLocationType = "conferenceRoom"
	LocationLocationTypeDefault         LocationLocationType = "default"
	LocationLocationTypeGeoCoordinates  LocationLocationType = "geoCoordinates"
	LocationLocationTypeHomeAddress     LocationLocationType = "homeAddress"
	LocationLocationTypeHotel           LocationLocationType = "hotel"
	LocationLocationTypeLocalBusiness   LocationLocationType = "localBusiness"
	LocationLocationTypePostalAddress   LocationLocationType = "postalAddress"
	LocationLocationTypeRestaurant      LocationLocationType = "restaurant"
	LocationLocationTypeStreetAddress   LocationLocationType = "streetAddress"
)

func PossibleValuesForLocationLocationType() []string {
	return []string{
		string(LocationLocationTypeBusinessAddress),
		string(LocationLocationTypeConferenceRoom),
		string(LocationLocationTypeDefault),
		string(LocationLocationTypeGeoCoordinates),
		string(LocationLocationTypeHomeAddress),
		string(LocationLocationTypeHotel),
		string(LocationLocationTypeLocalBusiness),
		string(LocationLocationTypePostalAddress),
		string(LocationLocationTypeRestaurant),
		string(LocationLocationTypeStreetAddress),
	}
}

func (s *LocationLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationLocationType(input string) (*LocationLocationType, error) {
	vals := map[string]LocationLocationType{
		"businessaddress": LocationLocationTypeBusinessAddress,
		"conferenceroom":  LocationLocationTypeConferenceRoom,
		"default":         LocationLocationTypeDefault,
		"geocoordinates":  LocationLocationTypeGeoCoordinates,
		"homeaddress":     LocationLocationTypeHomeAddress,
		"hotel":           LocationLocationTypeHotel,
		"localbusiness":   LocationLocationTypeLocalBusiness,
		"postaladdress":   LocationLocationTypePostalAddress,
		"restaurant":      LocationLocationTypeRestaurant,
		"streetaddress":   LocationLocationTypeStreetAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationLocationType(input)
	return &out, nil
}

type LocationUniqueIdType string

const (
	LocationUniqueIdTypeBing          LocationUniqueIdType = "bing"
	LocationUniqueIdTypeDirectory     LocationUniqueIdType = "directory"
	LocationUniqueIdTypeLocationStore LocationUniqueIdType = "locationStore"
	LocationUniqueIdTypePrivate       LocationUniqueIdType = "private"
	LocationUniqueIdTypeUnknown       LocationUniqueIdType = "unknown"
)

func PossibleValuesForLocationUniqueIdType() []string {
	return []string{
		string(LocationUniqueIdTypeBing),
		string(LocationUniqueIdTypeDirectory),
		string(LocationUniqueIdTypeLocationStore),
		string(LocationUniqueIdTypePrivate),
		string(LocationUniqueIdTypeUnknown),
	}
}

func (s *LocationUniqueIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationUniqueIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationUniqueIdType(input string) (*LocationUniqueIdType, error) {
	vals := map[string]LocationUniqueIdType{
		"bing":          LocationUniqueIdTypeBing,
		"directory":     LocationUniqueIdTypeDirectory,
		"locationstore": LocationUniqueIdTypeLocationStore,
		"private":       LocationUniqueIdTypePrivate,
		"unknown":       LocationUniqueIdTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationUniqueIdType(input)
	return &out, nil
}

type LongRunningOperationStatus string

const (
	LongRunningOperationStatusFailed     LongRunningOperationStatus = "failed"
	LongRunningOperationStatusNotStarted LongRunningOperationStatus = "notStarted"
	LongRunningOperationStatusRunning    LongRunningOperationStatus = "running"
	LongRunningOperationStatusSucceeded  LongRunningOperationStatus = "succeeded"
)

func PossibleValuesForLongRunningOperationStatus() []string {
	return []string{
		string(LongRunningOperationStatusFailed),
		string(LongRunningOperationStatusNotStarted),
		string(LongRunningOperationStatusRunning),
		string(LongRunningOperationStatusSucceeded),
	}
}

func (s *LongRunningOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLongRunningOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLongRunningOperationStatus(input string) (*LongRunningOperationStatus, error) {
	vals := map[string]LongRunningOperationStatus{
		"failed":     LongRunningOperationStatusFailed,
		"notstarted": LongRunningOperationStatusNotStarted,
		"running":    LongRunningOperationStatusRunning,
		"succeeded":  LongRunningOperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LongRunningOperationStatus(input)
	return &out, nil
}

type MailboxSettingsDelegateMeetingMessageDeliveryOptions string

const (
	MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateAndInformationToPrincipal MailboxSettingsDelegateMeetingMessageDeliveryOptions = "sendToDelegateAndInformationToPrincipal"
	MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateAndPrincipal              MailboxSettingsDelegateMeetingMessageDeliveryOptions = "sendToDelegateAndPrincipal"
	MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateOnly                      MailboxSettingsDelegateMeetingMessageDeliveryOptions = "sendToDelegateOnly"
)

func PossibleValuesForMailboxSettingsDelegateMeetingMessageDeliveryOptions() []string {
	return []string{
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateAndInformationToPrincipal),
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateAndPrincipal),
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateOnly),
	}
}

func (s *MailboxSettingsDelegateMeetingMessageDeliveryOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxSettingsDelegateMeetingMessageDeliveryOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxSettingsDelegateMeetingMessageDeliveryOptions(input string) (*MailboxSettingsDelegateMeetingMessageDeliveryOptions, error) {
	vals := map[string]MailboxSettingsDelegateMeetingMessageDeliveryOptions{
		"sendtodelegateandinformationtoprincipal": MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateAndInformationToPrincipal,
		"sendtodelegateandprincipal":              MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateAndPrincipal,
		"sendtodelegateonly":                      MailboxSettingsDelegateMeetingMessageDeliveryOptionsSendToDelegateOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsDelegateMeetingMessageDeliveryOptions(input)
	return &out, nil
}

type MailboxSettingsUserPurpose string

const (
	MailboxSettingsUserPurposeEquipment MailboxSettingsUserPurpose = "equipment"
	MailboxSettingsUserPurposeLinked    MailboxSettingsUserPurpose = "linked"
	MailboxSettingsUserPurposeOthers    MailboxSettingsUserPurpose = "others"
	MailboxSettingsUserPurposeRoom      MailboxSettingsUserPurpose = "room"
	MailboxSettingsUserPurposeShared    MailboxSettingsUserPurpose = "shared"
	MailboxSettingsUserPurposeUser      MailboxSettingsUserPurpose = "user"
)

func PossibleValuesForMailboxSettingsUserPurpose() []string {
	return []string{
		string(MailboxSettingsUserPurposeEquipment),
		string(MailboxSettingsUserPurposeLinked),
		string(MailboxSettingsUserPurposeOthers),
		string(MailboxSettingsUserPurposeRoom),
		string(MailboxSettingsUserPurposeShared),
		string(MailboxSettingsUserPurposeUser),
	}
}

func (s *MailboxSettingsUserPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxSettingsUserPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxSettingsUserPurpose(input string) (*MailboxSettingsUserPurpose, error) {
	vals := map[string]MailboxSettingsUserPurpose{
		"equipment": MailboxSettingsUserPurposeEquipment,
		"linked":    MailboxSettingsUserPurposeLinked,
		"others":    MailboxSettingsUserPurposeOthers,
		"room":      MailboxSettingsUserPurposeRoom,
		"shared":    MailboxSettingsUserPurposeShared,
		"user":      MailboxSettingsUserPurposeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsUserPurpose(input)
	return &out, nil
}

type ManagedAppRegistrationFlaggedReasons string

const (
	ManagedAppRegistrationFlaggedReasonsNone         ManagedAppRegistrationFlaggedReasons = "none"
	ManagedAppRegistrationFlaggedReasonsRootedDevice ManagedAppRegistrationFlaggedReasons = "rootedDevice"
)

func PossibleValuesForManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(ManagedAppRegistrationFlaggedReasonsNone),
		string(ManagedAppRegistrationFlaggedReasonsRootedDevice),
	}
}

func (s *ManagedAppRegistrationFlaggedReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppRegistrationFlaggedReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppRegistrationFlaggedReasons(input string) (*ManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]ManagedAppRegistrationFlaggedReasons{
		"none":         ManagedAppRegistrationFlaggedReasonsNone,
		"rooteddevice": ManagedAppRegistrationFlaggedReasonsRootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}

type ManagedDeviceComplianceState string

const (
	ManagedDeviceComplianceStateCompliant     ManagedDeviceComplianceState = "compliant"
	ManagedDeviceComplianceStateConfigManager ManagedDeviceComplianceState = "configManager"
	ManagedDeviceComplianceStateConflict      ManagedDeviceComplianceState = "conflict"
	ManagedDeviceComplianceStateError         ManagedDeviceComplianceState = "error"
	ManagedDeviceComplianceStateInGracePeriod ManagedDeviceComplianceState = "inGracePeriod"
	ManagedDeviceComplianceStateNoncompliant  ManagedDeviceComplianceState = "noncompliant"
	ManagedDeviceComplianceStateUnknown       ManagedDeviceComplianceState = "unknown"
)

func PossibleValuesForManagedDeviceComplianceState() []string {
	return []string{
		string(ManagedDeviceComplianceStateCompliant),
		string(ManagedDeviceComplianceStateConfigManager),
		string(ManagedDeviceComplianceStateConflict),
		string(ManagedDeviceComplianceStateError),
		string(ManagedDeviceComplianceStateInGracePeriod),
		string(ManagedDeviceComplianceStateNoncompliant),
		string(ManagedDeviceComplianceStateUnknown),
	}
}

func (s *ManagedDeviceComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceComplianceState(input string) (*ManagedDeviceComplianceState, error) {
	vals := map[string]ManagedDeviceComplianceState{
		"compliant":     ManagedDeviceComplianceStateCompliant,
		"configmanager": ManagedDeviceComplianceStateConfigManager,
		"conflict":      ManagedDeviceComplianceStateConflict,
		"error":         ManagedDeviceComplianceStateError,
		"ingraceperiod": ManagedDeviceComplianceStateInGracePeriod,
		"noncompliant":  ManagedDeviceComplianceStateNoncompliant,
		"unknown":       ManagedDeviceComplianceStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceComplianceState(input)
	return &out, nil
}

type ManagedDeviceDeviceEnrollmentType string

const (
	ManagedDeviceDeviceEnrollmentTypeAppleBulkWithUser                     ManagedDeviceDeviceEnrollmentType = "appleBulkWithUser"
	ManagedDeviceDeviceEnrollmentTypeAppleBulkWithoutUser                  ManagedDeviceDeviceEnrollmentType = "appleBulkWithoutUser"
	ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollment                   ManagedDeviceDeviceEnrollmentType = "appleUserEnrollment"
	ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollmentWithServiceAccount ManagedDeviceDeviceEnrollmentType = "appleUserEnrollmentWithServiceAccount"
	ManagedDeviceDeviceEnrollmentTypeDeviceEnrollmentManager               ManagedDeviceDeviceEnrollmentType = "deviceEnrollmentManager"
	ManagedDeviceDeviceEnrollmentTypeUnknown                               ManagedDeviceDeviceEnrollmentType = "unknown"
	ManagedDeviceDeviceEnrollmentTypeUserEnrollment                        ManagedDeviceDeviceEnrollmentType = "userEnrollment"
	ManagedDeviceDeviceEnrollmentTypeWindowsAutoEnrollment                 ManagedDeviceDeviceEnrollmentType = "windowsAutoEnrollment"
	ManagedDeviceDeviceEnrollmentTypeWindowsAzureADJoin                    ManagedDeviceDeviceEnrollmentType = "windowsAzureADJoin"
	ManagedDeviceDeviceEnrollmentTypeWindowsAzureADJoinUsingDeviceAuth     ManagedDeviceDeviceEnrollmentType = "windowsAzureADJoinUsingDeviceAuth"
	ManagedDeviceDeviceEnrollmentTypeWindowsBulkAzureDomainJoin            ManagedDeviceDeviceEnrollmentType = "windowsBulkAzureDomainJoin"
	ManagedDeviceDeviceEnrollmentTypeWindowsBulkUserless                   ManagedDeviceDeviceEnrollmentType = "windowsBulkUserless"
	ManagedDeviceDeviceEnrollmentTypeWindowsCoManagement                   ManagedDeviceDeviceEnrollmentType = "windowsCoManagement"
)

func PossibleValuesForManagedDeviceDeviceEnrollmentType() []string {
	return []string{
		string(ManagedDeviceDeviceEnrollmentTypeAppleBulkWithUser),
		string(ManagedDeviceDeviceEnrollmentTypeAppleBulkWithoutUser),
		string(ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollment),
		string(ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollmentWithServiceAccount),
		string(ManagedDeviceDeviceEnrollmentTypeDeviceEnrollmentManager),
		string(ManagedDeviceDeviceEnrollmentTypeUnknown),
		string(ManagedDeviceDeviceEnrollmentTypeUserEnrollment),
		string(ManagedDeviceDeviceEnrollmentTypeWindowsAutoEnrollment),
		string(ManagedDeviceDeviceEnrollmentTypeWindowsAzureADJoin),
		string(ManagedDeviceDeviceEnrollmentTypeWindowsAzureADJoinUsingDeviceAuth),
		string(ManagedDeviceDeviceEnrollmentTypeWindowsBulkAzureDomainJoin),
		string(ManagedDeviceDeviceEnrollmentTypeWindowsBulkUserless),
		string(ManagedDeviceDeviceEnrollmentTypeWindowsCoManagement),
	}
}

func (s *ManagedDeviceDeviceEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceDeviceEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceDeviceEnrollmentType(input string) (*ManagedDeviceDeviceEnrollmentType, error) {
	vals := map[string]ManagedDeviceDeviceEnrollmentType{
		"applebulkwithuser":                     ManagedDeviceDeviceEnrollmentTypeAppleBulkWithUser,
		"applebulkwithoutuser":                  ManagedDeviceDeviceEnrollmentTypeAppleBulkWithoutUser,
		"appleuserenrollment":                   ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollmentWithServiceAccount,
		"deviceenrollmentmanager":               ManagedDeviceDeviceEnrollmentTypeDeviceEnrollmentManager,
		"unknown":                               ManagedDeviceDeviceEnrollmentTypeUnknown,
		"userenrollment":                        ManagedDeviceDeviceEnrollmentTypeUserEnrollment,
		"windowsautoenrollment":                 ManagedDeviceDeviceEnrollmentTypeWindowsAutoEnrollment,
		"windowsazureadjoin":                    ManagedDeviceDeviceEnrollmentTypeWindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     ManagedDeviceDeviceEnrollmentTypeWindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            ManagedDeviceDeviceEnrollmentTypeWindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   ManagedDeviceDeviceEnrollmentTypeWindowsBulkUserless,
		"windowscomanagement":                   ManagedDeviceDeviceEnrollmentTypeWindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceEnrollmentType(input)
	return &out, nil
}

type ManagedDeviceDeviceRegistrationState string

const (
	ManagedDeviceDeviceRegistrationStateApprovalPending                ManagedDeviceDeviceRegistrationState = "approvalPending"
	ManagedDeviceDeviceRegistrationStateCertificateReset               ManagedDeviceDeviceRegistrationState = "certificateReset"
	ManagedDeviceDeviceRegistrationStateKeyConflict                    ManagedDeviceDeviceRegistrationState = "keyConflict"
	ManagedDeviceDeviceRegistrationStateNotRegistered                  ManagedDeviceDeviceRegistrationState = "notRegistered"
	ManagedDeviceDeviceRegistrationStateNotRegisteredPendingEnrollment ManagedDeviceDeviceRegistrationState = "notRegisteredPendingEnrollment"
	ManagedDeviceDeviceRegistrationStateRegistered                     ManagedDeviceDeviceRegistrationState = "registered"
	ManagedDeviceDeviceRegistrationStateRevoked                        ManagedDeviceDeviceRegistrationState = "revoked"
	ManagedDeviceDeviceRegistrationStateUnknown                        ManagedDeviceDeviceRegistrationState = "unknown"
)

func PossibleValuesForManagedDeviceDeviceRegistrationState() []string {
	return []string{
		string(ManagedDeviceDeviceRegistrationStateApprovalPending),
		string(ManagedDeviceDeviceRegistrationStateCertificateReset),
		string(ManagedDeviceDeviceRegistrationStateKeyConflict),
		string(ManagedDeviceDeviceRegistrationStateNotRegistered),
		string(ManagedDeviceDeviceRegistrationStateNotRegisteredPendingEnrollment),
		string(ManagedDeviceDeviceRegistrationStateRegistered),
		string(ManagedDeviceDeviceRegistrationStateRevoked),
		string(ManagedDeviceDeviceRegistrationStateUnknown),
	}
}

func (s *ManagedDeviceDeviceRegistrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceDeviceRegistrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceDeviceRegistrationState(input string) (*ManagedDeviceDeviceRegistrationState, error) {
	vals := map[string]ManagedDeviceDeviceRegistrationState{
		"approvalpending":                ManagedDeviceDeviceRegistrationStateApprovalPending,
		"certificatereset":               ManagedDeviceDeviceRegistrationStateCertificateReset,
		"keyconflict":                    ManagedDeviceDeviceRegistrationStateKeyConflict,
		"notregistered":                  ManagedDeviceDeviceRegistrationStateNotRegistered,
		"notregisteredpendingenrollment": ManagedDeviceDeviceRegistrationStateNotRegisteredPendingEnrollment,
		"registered":                     ManagedDeviceDeviceRegistrationStateRegistered,
		"revoked":                        ManagedDeviceDeviceRegistrationStateRevoked,
		"unknown":                        ManagedDeviceDeviceRegistrationStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceRegistrationState(input)
	return &out, nil
}

type ManagedDeviceExchangeAccessState string

const (
	ManagedDeviceExchangeAccessStateAllowed     ManagedDeviceExchangeAccessState = "allowed"
	ManagedDeviceExchangeAccessStateBlocked     ManagedDeviceExchangeAccessState = "blocked"
	ManagedDeviceExchangeAccessStateNone        ManagedDeviceExchangeAccessState = "none"
	ManagedDeviceExchangeAccessStateQuarantined ManagedDeviceExchangeAccessState = "quarantined"
	ManagedDeviceExchangeAccessStateUnknown     ManagedDeviceExchangeAccessState = "unknown"
)

func PossibleValuesForManagedDeviceExchangeAccessState() []string {
	return []string{
		string(ManagedDeviceExchangeAccessStateAllowed),
		string(ManagedDeviceExchangeAccessStateBlocked),
		string(ManagedDeviceExchangeAccessStateNone),
		string(ManagedDeviceExchangeAccessStateQuarantined),
		string(ManagedDeviceExchangeAccessStateUnknown),
	}
}

func (s *ManagedDeviceExchangeAccessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceExchangeAccessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceExchangeAccessState(input string) (*ManagedDeviceExchangeAccessState, error) {
	vals := map[string]ManagedDeviceExchangeAccessState{
		"allowed":     ManagedDeviceExchangeAccessStateAllowed,
		"blocked":     ManagedDeviceExchangeAccessStateBlocked,
		"none":        ManagedDeviceExchangeAccessStateNone,
		"quarantined": ManagedDeviceExchangeAccessStateQuarantined,
		"unknown":     ManagedDeviceExchangeAccessStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceExchangeAccessState(input)
	return &out, nil
}

type ManagedDeviceExchangeAccessStateReason string

const (
	ManagedDeviceExchangeAccessStateReasonAzureADBlockDueToAccessPolicy ManagedDeviceExchangeAccessStateReason = "azureADBlockDueToAccessPolicy"
	ManagedDeviceExchangeAccessStateReasonCompliant                     ManagedDeviceExchangeAccessStateReason = "compliant"
	ManagedDeviceExchangeAccessStateReasonCompromisedPassword           ManagedDeviceExchangeAccessStateReason = "compromisedPassword"
	ManagedDeviceExchangeAccessStateReasonDeviceNotKnownWithManagedApp  ManagedDeviceExchangeAccessStateReason = "deviceNotKnownWithManagedApp"
	ManagedDeviceExchangeAccessStateReasonExchangeDeviceRule            ManagedDeviceExchangeAccessStateReason = "exchangeDeviceRule"
	ManagedDeviceExchangeAccessStateReasonExchangeGlobalRule            ManagedDeviceExchangeAccessStateReason = "exchangeGlobalRule"
	ManagedDeviceExchangeAccessStateReasonExchangeIndividualRule        ManagedDeviceExchangeAccessStateReason = "exchangeIndividualRule"
	ManagedDeviceExchangeAccessStateReasonExchangeMailboxPolicy         ManagedDeviceExchangeAccessStateReason = "exchangeMailboxPolicy"
	ManagedDeviceExchangeAccessStateReasonExchangeUpgrade               ManagedDeviceExchangeAccessStateReason = "exchangeUpgrade"
	ManagedDeviceExchangeAccessStateReasonMfaRequired                   ManagedDeviceExchangeAccessStateReason = "mfaRequired"
	ManagedDeviceExchangeAccessStateReasonNone                          ManagedDeviceExchangeAccessStateReason = "none"
	ManagedDeviceExchangeAccessStateReasonNotCompliant                  ManagedDeviceExchangeAccessStateReason = "notCompliant"
	ManagedDeviceExchangeAccessStateReasonNotEnrolled                   ManagedDeviceExchangeAccessStateReason = "notEnrolled"
	ManagedDeviceExchangeAccessStateReasonOther                         ManagedDeviceExchangeAccessStateReason = "other"
	ManagedDeviceExchangeAccessStateReasonUnknown                       ManagedDeviceExchangeAccessStateReason = "unknown"
	ManagedDeviceExchangeAccessStateReasonUnknownLocation               ManagedDeviceExchangeAccessStateReason = "unknownLocation"
)

func PossibleValuesForManagedDeviceExchangeAccessStateReason() []string {
	return []string{
		string(ManagedDeviceExchangeAccessStateReasonAzureADBlockDueToAccessPolicy),
		string(ManagedDeviceExchangeAccessStateReasonCompliant),
		string(ManagedDeviceExchangeAccessStateReasonCompromisedPassword),
		string(ManagedDeviceExchangeAccessStateReasonDeviceNotKnownWithManagedApp),
		string(ManagedDeviceExchangeAccessStateReasonExchangeDeviceRule),
		string(ManagedDeviceExchangeAccessStateReasonExchangeGlobalRule),
		string(ManagedDeviceExchangeAccessStateReasonExchangeIndividualRule),
		string(ManagedDeviceExchangeAccessStateReasonExchangeMailboxPolicy),
		string(ManagedDeviceExchangeAccessStateReasonExchangeUpgrade),
		string(ManagedDeviceExchangeAccessStateReasonMfaRequired),
		string(ManagedDeviceExchangeAccessStateReasonNone),
		string(ManagedDeviceExchangeAccessStateReasonNotCompliant),
		string(ManagedDeviceExchangeAccessStateReasonNotEnrolled),
		string(ManagedDeviceExchangeAccessStateReasonOther),
		string(ManagedDeviceExchangeAccessStateReasonUnknown),
		string(ManagedDeviceExchangeAccessStateReasonUnknownLocation),
	}
}

func (s *ManagedDeviceExchangeAccessStateReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceExchangeAccessStateReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceExchangeAccessStateReason(input string) (*ManagedDeviceExchangeAccessStateReason, error) {
	vals := map[string]ManagedDeviceExchangeAccessStateReason{
		"azureadblockduetoaccesspolicy": ManagedDeviceExchangeAccessStateReasonAzureADBlockDueToAccessPolicy,
		"compliant":                     ManagedDeviceExchangeAccessStateReasonCompliant,
		"compromisedpassword":           ManagedDeviceExchangeAccessStateReasonCompromisedPassword,
		"devicenotknownwithmanagedapp":  ManagedDeviceExchangeAccessStateReasonDeviceNotKnownWithManagedApp,
		"exchangedevicerule":            ManagedDeviceExchangeAccessStateReasonExchangeDeviceRule,
		"exchangeglobalrule":            ManagedDeviceExchangeAccessStateReasonExchangeGlobalRule,
		"exchangeindividualrule":        ManagedDeviceExchangeAccessStateReasonExchangeIndividualRule,
		"exchangemailboxpolicy":         ManagedDeviceExchangeAccessStateReasonExchangeMailboxPolicy,
		"exchangeupgrade":               ManagedDeviceExchangeAccessStateReasonExchangeUpgrade,
		"mfarequired":                   ManagedDeviceExchangeAccessStateReasonMfaRequired,
		"none":                          ManagedDeviceExchangeAccessStateReasonNone,
		"notcompliant":                  ManagedDeviceExchangeAccessStateReasonNotCompliant,
		"notenrolled":                   ManagedDeviceExchangeAccessStateReasonNotEnrolled,
		"other":                         ManagedDeviceExchangeAccessStateReasonOther,
		"unknown":                       ManagedDeviceExchangeAccessStateReasonUnknown,
		"unknownlocation":               ManagedDeviceExchangeAccessStateReasonUnknownLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceExchangeAccessStateReason(input)
	return &out, nil
}

type ManagedDeviceManagedDeviceOwnerType string

const (
	ManagedDeviceManagedDeviceOwnerTypeCompany  ManagedDeviceManagedDeviceOwnerType = "company"
	ManagedDeviceManagedDeviceOwnerTypePersonal ManagedDeviceManagedDeviceOwnerType = "personal"
	ManagedDeviceManagedDeviceOwnerTypeUnknown  ManagedDeviceManagedDeviceOwnerType = "unknown"
)

func PossibleValuesForManagedDeviceManagedDeviceOwnerType() []string {
	return []string{
		string(ManagedDeviceManagedDeviceOwnerTypeCompany),
		string(ManagedDeviceManagedDeviceOwnerTypePersonal),
		string(ManagedDeviceManagedDeviceOwnerTypeUnknown),
	}
}

func (s *ManagedDeviceManagedDeviceOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceManagedDeviceOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceManagedDeviceOwnerType(input string) (*ManagedDeviceManagedDeviceOwnerType, error) {
	vals := map[string]ManagedDeviceManagedDeviceOwnerType{
		"company":  ManagedDeviceManagedDeviceOwnerTypeCompany,
		"personal": ManagedDeviceManagedDeviceOwnerTypePersonal,
		"unknown":  ManagedDeviceManagedDeviceOwnerTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagedDeviceOwnerType(input)
	return &out, nil
}

type ManagedDeviceManagementAgent string

const (
	ManagedDeviceManagementAgentConfigurationManagerClient        ManagedDeviceManagementAgent = "configurationManagerClient"
	ManagedDeviceManagementAgentConfigurationManagerClientMdm     ManagedDeviceManagementAgent = "configurationManagerClientMdm"
	ManagedDeviceManagementAgentConfigurationManagerClientMdmEas  ManagedDeviceManagementAgent = "configurationManagerClientMdmEas"
	ManagedDeviceManagementAgentEas                               ManagedDeviceManagementAgent = "eas"
	ManagedDeviceManagementAgentEasIntuneClient                   ManagedDeviceManagementAgent = "easIntuneClient"
	ManagedDeviceManagementAgentEasMdm                            ManagedDeviceManagementAgent = "easMdm"
	ManagedDeviceManagementAgentGoogleCloudDevicePolicyController ManagedDeviceManagementAgent = "googleCloudDevicePolicyController"
	ManagedDeviceManagementAgentIntuneClient                      ManagedDeviceManagementAgent = "intuneClient"
	ManagedDeviceManagementAgentJamf                              ManagedDeviceManagementAgent = "jamf"
	ManagedDeviceManagementAgentMdm                               ManagedDeviceManagementAgent = "mdm"
	ManagedDeviceManagementAgentMicrosoft365ManagedMdm            ManagedDeviceManagementAgent = "microsoft365ManagedMdm"
	ManagedDeviceManagementAgentMsSense                           ManagedDeviceManagementAgent = "msSense"
	ManagedDeviceManagementAgentUnknown                           ManagedDeviceManagementAgent = "unknown"
)

func PossibleValuesForManagedDeviceManagementAgent() []string {
	return []string{
		string(ManagedDeviceManagementAgentConfigurationManagerClient),
		string(ManagedDeviceManagementAgentConfigurationManagerClientMdm),
		string(ManagedDeviceManagementAgentConfigurationManagerClientMdmEas),
		string(ManagedDeviceManagementAgentEas),
		string(ManagedDeviceManagementAgentEasIntuneClient),
		string(ManagedDeviceManagementAgentEasMdm),
		string(ManagedDeviceManagementAgentGoogleCloudDevicePolicyController),
		string(ManagedDeviceManagementAgentIntuneClient),
		string(ManagedDeviceManagementAgentJamf),
		string(ManagedDeviceManagementAgentMdm),
		string(ManagedDeviceManagementAgentMicrosoft365ManagedMdm),
		string(ManagedDeviceManagementAgentMsSense),
		string(ManagedDeviceManagementAgentUnknown),
	}
}

func (s *ManagedDeviceManagementAgent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceManagementAgent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceManagementAgent(input string) (*ManagedDeviceManagementAgent, error) {
	vals := map[string]ManagedDeviceManagementAgent{
		"configurationmanagerclient":        ManagedDeviceManagementAgentConfigurationManagerClient,
		"configurationmanagerclientmdm":     ManagedDeviceManagementAgentConfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  ManagedDeviceManagementAgentConfigurationManagerClientMdmEas,
		"eas":                               ManagedDeviceManagementAgentEas,
		"easintuneclient":                   ManagedDeviceManagementAgentEasIntuneClient,
		"easmdm":                            ManagedDeviceManagementAgentEasMdm,
		"googleclouddevicepolicycontroller": ManagedDeviceManagementAgentGoogleCloudDevicePolicyController,
		"intuneclient":                      ManagedDeviceManagementAgentIntuneClient,
		"jamf":                              ManagedDeviceManagementAgentJamf,
		"mdm":                               ManagedDeviceManagementAgentMdm,
		"microsoft365managedmdm":            ManagedDeviceManagementAgentMicrosoft365ManagedMdm,
		"mssense":                           ManagedDeviceManagementAgentMsSense,
		"unknown":                           ManagedDeviceManagementAgentUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementAgent(input)
	return &out, nil
}

type ManagedDevicePartnerReportedThreatState string

const (
	ManagedDevicePartnerReportedThreatStateActivated      ManagedDevicePartnerReportedThreatState = "activated"
	ManagedDevicePartnerReportedThreatStateCompromised    ManagedDevicePartnerReportedThreatState = "compromised"
	ManagedDevicePartnerReportedThreatStateDeactivated    ManagedDevicePartnerReportedThreatState = "deactivated"
	ManagedDevicePartnerReportedThreatStateHighSeverity   ManagedDevicePartnerReportedThreatState = "highSeverity"
	ManagedDevicePartnerReportedThreatStateLowSeverity    ManagedDevicePartnerReportedThreatState = "lowSeverity"
	ManagedDevicePartnerReportedThreatStateMediumSeverity ManagedDevicePartnerReportedThreatState = "mediumSeverity"
	ManagedDevicePartnerReportedThreatStateMisconfigured  ManagedDevicePartnerReportedThreatState = "misconfigured"
	ManagedDevicePartnerReportedThreatStateSecured        ManagedDevicePartnerReportedThreatState = "secured"
	ManagedDevicePartnerReportedThreatStateUnknown        ManagedDevicePartnerReportedThreatState = "unknown"
	ManagedDevicePartnerReportedThreatStateUnresponsive   ManagedDevicePartnerReportedThreatState = "unresponsive"
)

func PossibleValuesForManagedDevicePartnerReportedThreatState() []string {
	return []string{
		string(ManagedDevicePartnerReportedThreatStateActivated),
		string(ManagedDevicePartnerReportedThreatStateCompromised),
		string(ManagedDevicePartnerReportedThreatStateDeactivated),
		string(ManagedDevicePartnerReportedThreatStateHighSeverity),
		string(ManagedDevicePartnerReportedThreatStateLowSeverity),
		string(ManagedDevicePartnerReportedThreatStateMediumSeverity),
		string(ManagedDevicePartnerReportedThreatStateMisconfigured),
		string(ManagedDevicePartnerReportedThreatStateSecured),
		string(ManagedDevicePartnerReportedThreatStateUnknown),
		string(ManagedDevicePartnerReportedThreatStateUnresponsive),
	}
}

func (s *ManagedDevicePartnerReportedThreatState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDevicePartnerReportedThreatState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDevicePartnerReportedThreatState(input string) (*ManagedDevicePartnerReportedThreatState, error) {
	vals := map[string]ManagedDevicePartnerReportedThreatState{
		"activated":      ManagedDevicePartnerReportedThreatStateActivated,
		"compromised":    ManagedDevicePartnerReportedThreatStateCompromised,
		"deactivated":    ManagedDevicePartnerReportedThreatStateDeactivated,
		"highseverity":   ManagedDevicePartnerReportedThreatStateHighSeverity,
		"lowseverity":    ManagedDevicePartnerReportedThreatStateLowSeverity,
		"mediumseverity": ManagedDevicePartnerReportedThreatStateMediumSeverity,
		"misconfigured":  ManagedDevicePartnerReportedThreatStateMisconfigured,
		"secured":        ManagedDevicePartnerReportedThreatStateSecured,
		"unknown":        ManagedDevicePartnerReportedThreatStateUnknown,
		"unresponsive":   ManagedDevicePartnerReportedThreatStateUnresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDevicePartnerReportedThreatState(input)
	return &out, nil
}

type MeetingParticipantInfoRole string

const (
	MeetingParticipantInfoRoleAttendee    MeetingParticipantInfoRole = "attendee"
	MeetingParticipantInfoRoleCoorganizer MeetingParticipantInfoRole = "coorganizer"
	MeetingParticipantInfoRolePresenter   MeetingParticipantInfoRole = "presenter"
	MeetingParticipantInfoRoleProducer    MeetingParticipantInfoRole = "producer"
)

func PossibleValuesForMeetingParticipantInfoRole() []string {
	return []string{
		string(MeetingParticipantInfoRoleAttendee),
		string(MeetingParticipantInfoRoleCoorganizer),
		string(MeetingParticipantInfoRolePresenter),
		string(MeetingParticipantInfoRoleProducer),
	}
}

func (s *MeetingParticipantInfoRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingParticipantInfoRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingParticipantInfoRole(input string) (*MeetingParticipantInfoRole, error) {
	vals := map[string]MeetingParticipantInfoRole{
		"attendee":    MeetingParticipantInfoRoleAttendee,
		"coorganizer": MeetingParticipantInfoRoleCoorganizer,
		"presenter":   MeetingParticipantInfoRolePresenter,
		"producer":    MeetingParticipantInfoRoleProducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingParticipantInfoRole(input)
	return &out, nil
}

type MessageImportance string

const (
	MessageImportanceHigh   MessageImportance = "high"
	MessageImportanceLow    MessageImportance = "low"
	MessageImportanceNormal MessageImportance = "normal"
)

func PossibleValuesForMessageImportance() []string {
	return []string{
		string(MessageImportanceHigh),
		string(MessageImportanceLow),
		string(MessageImportanceNormal),
	}
}

func (s *MessageImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageImportance(input string) (*MessageImportance, error) {
	vals := map[string]MessageImportance{
		"high":   MessageImportanceHigh,
		"low":    MessageImportanceLow,
		"normal": MessageImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageImportance(input)
	return &out, nil
}

type MessageInferenceClassification string

const (
	MessageInferenceClassificationFocused MessageInferenceClassification = "focused"
	MessageInferenceClassificationOther   MessageInferenceClassification = "other"
)

func PossibleValuesForMessageInferenceClassification() []string {
	return []string{
		string(MessageInferenceClassificationFocused),
		string(MessageInferenceClassificationOther),
	}
}

func (s *MessageInferenceClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageInferenceClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageInferenceClassification(input string) (*MessageInferenceClassification, error) {
	vals := map[string]MessageInferenceClassification{
		"focused": MessageInferenceClassificationFocused,
		"other":   MessageInferenceClassificationOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageInferenceClassification(input)
	return &out, nil
}

type MessageRuleActionsMarkImportance string

const (
	MessageRuleActionsMarkImportanceHigh   MessageRuleActionsMarkImportance = "high"
	MessageRuleActionsMarkImportanceLow    MessageRuleActionsMarkImportance = "low"
	MessageRuleActionsMarkImportanceNormal MessageRuleActionsMarkImportance = "normal"
)

func PossibleValuesForMessageRuleActionsMarkImportance() []string {
	return []string{
		string(MessageRuleActionsMarkImportanceHigh),
		string(MessageRuleActionsMarkImportanceLow),
		string(MessageRuleActionsMarkImportanceNormal),
	}
}

func (s *MessageRuleActionsMarkImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRuleActionsMarkImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRuleActionsMarkImportance(input string) (*MessageRuleActionsMarkImportance, error) {
	vals := map[string]MessageRuleActionsMarkImportance{
		"high":   MessageRuleActionsMarkImportanceHigh,
		"low":    MessageRuleActionsMarkImportanceLow,
		"normal": MessageRuleActionsMarkImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRuleActionsMarkImportance(input)
	return &out, nil
}

type MessageRulePredicatesImportance string

const (
	MessageRulePredicatesImportanceHigh   MessageRulePredicatesImportance = "high"
	MessageRulePredicatesImportanceLow    MessageRulePredicatesImportance = "low"
	MessageRulePredicatesImportanceNormal MessageRulePredicatesImportance = "normal"
)

func PossibleValuesForMessageRulePredicatesImportance() []string {
	return []string{
		string(MessageRulePredicatesImportanceHigh),
		string(MessageRulePredicatesImportanceLow),
		string(MessageRulePredicatesImportanceNormal),
	}
}

func (s *MessageRulePredicatesImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRulePredicatesImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRulePredicatesImportance(input string) (*MessageRulePredicatesImportance, error) {
	vals := map[string]MessageRulePredicatesImportance{
		"high":   MessageRulePredicatesImportanceHigh,
		"low":    MessageRulePredicatesImportanceLow,
		"normal": MessageRulePredicatesImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesImportance(input)
	return &out, nil
}

type MessageRulePredicatesMessageActionFlag string

const (
	MessageRulePredicatesMessageActionFlagAny                 MessageRulePredicatesMessageActionFlag = "any"
	MessageRulePredicatesMessageActionFlagCall                MessageRulePredicatesMessageActionFlag = "call"
	MessageRulePredicatesMessageActionFlagDoNotForward        MessageRulePredicatesMessageActionFlag = "doNotForward"
	MessageRulePredicatesMessageActionFlagFollowUp            MessageRulePredicatesMessageActionFlag = "followUp"
	MessageRulePredicatesMessageActionFlagForward             MessageRulePredicatesMessageActionFlag = "forward"
	MessageRulePredicatesMessageActionFlagFyi                 MessageRulePredicatesMessageActionFlag = "fyi"
	MessageRulePredicatesMessageActionFlagNoResponseNecessary MessageRulePredicatesMessageActionFlag = "noResponseNecessary"
	MessageRulePredicatesMessageActionFlagRead                MessageRulePredicatesMessageActionFlag = "read"
	MessageRulePredicatesMessageActionFlagReply               MessageRulePredicatesMessageActionFlag = "reply"
	MessageRulePredicatesMessageActionFlagReplyToAll          MessageRulePredicatesMessageActionFlag = "replyToAll"
	MessageRulePredicatesMessageActionFlagReview              MessageRulePredicatesMessageActionFlag = "review"
)

func PossibleValuesForMessageRulePredicatesMessageActionFlag() []string {
	return []string{
		string(MessageRulePredicatesMessageActionFlagAny),
		string(MessageRulePredicatesMessageActionFlagCall),
		string(MessageRulePredicatesMessageActionFlagDoNotForward),
		string(MessageRulePredicatesMessageActionFlagFollowUp),
		string(MessageRulePredicatesMessageActionFlagForward),
		string(MessageRulePredicatesMessageActionFlagFyi),
		string(MessageRulePredicatesMessageActionFlagNoResponseNecessary),
		string(MessageRulePredicatesMessageActionFlagRead),
		string(MessageRulePredicatesMessageActionFlagReply),
		string(MessageRulePredicatesMessageActionFlagReplyToAll),
		string(MessageRulePredicatesMessageActionFlagReview),
	}
}

func (s *MessageRulePredicatesMessageActionFlag) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRulePredicatesMessageActionFlag(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRulePredicatesMessageActionFlag(input string) (*MessageRulePredicatesMessageActionFlag, error) {
	vals := map[string]MessageRulePredicatesMessageActionFlag{
		"any":                 MessageRulePredicatesMessageActionFlagAny,
		"call":                MessageRulePredicatesMessageActionFlagCall,
		"donotforward":        MessageRulePredicatesMessageActionFlagDoNotForward,
		"followup":            MessageRulePredicatesMessageActionFlagFollowUp,
		"forward":             MessageRulePredicatesMessageActionFlagForward,
		"fyi":                 MessageRulePredicatesMessageActionFlagFyi,
		"noresponsenecessary": MessageRulePredicatesMessageActionFlagNoResponseNecessary,
		"read":                MessageRulePredicatesMessageActionFlagRead,
		"reply":               MessageRulePredicatesMessageActionFlagReply,
		"replytoall":          MessageRulePredicatesMessageActionFlagReplyToAll,
		"review":              MessageRulePredicatesMessageActionFlagReview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesMessageActionFlag(input)
	return &out, nil
}

type MessageRulePredicatesSensitivity string

const (
	MessageRulePredicatesSensitivityConfidential MessageRulePredicatesSensitivity = "confidential"
	MessageRulePredicatesSensitivityNormal       MessageRulePredicatesSensitivity = "normal"
	MessageRulePredicatesSensitivityPersonal     MessageRulePredicatesSensitivity = "personal"
	MessageRulePredicatesSensitivityPrivate      MessageRulePredicatesSensitivity = "private"
)

func PossibleValuesForMessageRulePredicatesSensitivity() []string {
	return []string{
		string(MessageRulePredicatesSensitivityConfidential),
		string(MessageRulePredicatesSensitivityNormal),
		string(MessageRulePredicatesSensitivityPersonal),
		string(MessageRulePredicatesSensitivityPrivate),
	}
}

func (s *MessageRulePredicatesSensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMessageRulePredicatesSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMessageRulePredicatesSensitivity(input string) (*MessageRulePredicatesSensitivity, error) {
	vals := map[string]MessageRulePredicatesSensitivity{
		"confidential": MessageRulePredicatesSensitivityConfidential,
		"normal":       MessageRulePredicatesSensitivityNormal,
		"personal":     MessageRulePredicatesSensitivityPersonal,
		"private":      MessageRulePredicatesSensitivityPrivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MessageRulePredicatesSensitivity(input)
	return &out, nil
}

type NotebookUserRole string

const (
	NotebookUserRoleContributor NotebookUserRole = "Contributor"
	NotebookUserRoleNone        NotebookUserRole = "None"
	NotebookUserRoleOwner       NotebookUserRole = "Owner"
	NotebookUserRoleReader      NotebookUserRole = "Reader"
)

func PossibleValuesForNotebookUserRole() []string {
	return []string{
		string(NotebookUserRoleContributor),
		string(NotebookUserRoleNone),
		string(NotebookUserRoleOwner),
		string(NotebookUserRoleReader),
	}
}

func (s *NotebookUserRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotebookUserRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotebookUserRole(input string) (*NotebookUserRole, error) {
	vals := map[string]NotebookUserRole{
		"contributor": NotebookUserRoleContributor,
		"none":        NotebookUserRoleNone,
		"owner":       NotebookUserRoleOwner,
		"reader":      NotebookUserRoleReader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotebookUserRole(input)
	return &out, nil
}

type OfferShiftRequestAssignedTo string

const (
	OfferShiftRequestAssignedToManager   OfferShiftRequestAssignedTo = "manager"
	OfferShiftRequestAssignedToRecipient OfferShiftRequestAssignedTo = "recipient"
	OfferShiftRequestAssignedToSender    OfferShiftRequestAssignedTo = "sender"
	OfferShiftRequestAssignedToSystem    OfferShiftRequestAssignedTo = "system"
)

func PossibleValuesForOfferShiftRequestAssignedTo() []string {
	return []string{
		string(OfferShiftRequestAssignedToManager),
		string(OfferShiftRequestAssignedToRecipient),
		string(OfferShiftRequestAssignedToSender),
		string(OfferShiftRequestAssignedToSystem),
	}
}

func (s *OfferShiftRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfferShiftRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfferShiftRequestAssignedTo(input string) (*OfferShiftRequestAssignedTo, error) {
	vals := map[string]OfferShiftRequestAssignedTo{
		"manager":   OfferShiftRequestAssignedToManager,
		"recipient": OfferShiftRequestAssignedToRecipient,
		"sender":    OfferShiftRequestAssignedToSender,
		"system":    OfferShiftRequestAssignedToSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfferShiftRequestAssignedTo(input)
	return &out, nil
}

type OfferShiftRequestState string

const (
	OfferShiftRequestStateApproved OfferShiftRequestState = "approved"
	OfferShiftRequestStateDeclined OfferShiftRequestState = "declined"
	OfferShiftRequestStatePending  OfferShiftRequestState = "pending"
)

func PossibleValuesForOfferShiftRequestState() []string {
	return []string{
		string(OfferShiftRequestStateApproved),
		string(OfferShiftRequestStateDeclined),
		string(OfferShiftRequestStatePending),
	}
}

func (s *OfferShiftRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfferShiftRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfferShiftRequestState(input string) (*OfferShiftRequestState, error) {
	vals := map[string]OfferShiftRequestState{
		"approved": OfferShiftRequestStateApproved,
		"declined": OfferShiftRequestStateDeclined,
		"pending":  OfferShiftRequestStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfferShiftRequestState(input)
	return &out, nil
}

type OnenoteOperationStatus string

const (
	OnenoteOperationStatusCompleted  OnenoteOperationStatus = "Completed"
	OnenoteOperationStatusFailed     OnenoteOperationStatus = "Failed"
	OnenoteOperationStatusNotStarted OnenoteOperationStatus = "NotStarted"
	OnenoteOperationStatusRunning    OnenoteOperationStatus = "Running"
)

func PossibleValuesForOnenoteOperationStatus() []string {
	return []string{
		string(OnenoteOperationStatusCompleted),
		string(OnenoteOperationStatusFailed),
		string(OnenoteOperationStatusNotStarted),
		string(OnenoteOperationStatusRunning),
	}
}

func (s *OnenoteOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenoteOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenoteOperationStatus(input string) (*OnenoteOperationStatus, error) {
	vals := map[string]OnenoteOperationStatus{
		"completed":  OnenoteOperationStatusCompleted,
		"failed":     OnenoteOperationStatusFailed,
		"notstarted": OnenoteOperationStatusNotStarted,
		"running":    OnenoteOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenoteOperationStatus(input)
	return &out, nil
}

type OnlineMeetingAllowMeetingChat string

const (
	OnlineMeetingAllowMeetingChatDisabled OnlineMeetingAllowMeetingChat = "disabled"
	OnlineMeetingAllowMeetingChatEnabled  OnlineMeetingAllowMeetingChat = "enabled"
	OnlineMeetingAllowMeetingChatLimited  OnlineMeetingAllowMeetingChat = "limited"
)

func PossibleValuesForOnlineMeetingAllowMeetingChat() []string {
	return []string{
		string(OnlineMeetingAllowMeetingChatDisabled),
		string(OnlineMeetingAllowMeetingChatEnabled),
		string(OnlineMeetingAllowMeetingChatLimited),
	}
}

func (s *OnlineMeetingAllowMeetingChat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingAllowMeetingChat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingAllowMeetingChat(input string) (*OnlineMeetingAllowMeetingChat, error) {
	vals := map[string]OnlineMeetingAllowMeetingChat{
		"disabled": OnlineMeetingAllowMeetingChatDisabled,
		"enabled":  OnlineMeetingAllowMeetingChatEnabled,
		"limited":  OnlineMeetingAllowMeetingChatLimited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAllowMeetingChat(input)
	return &out, nil
}

type OnlineMeetingAllowedPresenters string

const (
	OnlineMeetingAllowedPresentersEveryone        OnlineMeetingAllowedPresenters = "everyone"
	OnlineMeetingAllowedPresentersOrganization    OnlineMeetingAllowedPresenters = "organization"
	OnlineMeetingAllowedPresentersOrganizer       OnlineMeetingAllowedPresenters = "organizer"
	OnlineMeetingAllowedPresentersRoleIsPresenter OnlineMeetingAllowedPresenters = "roleIsPresenter"
)

func PossibleValuesForOnlineMeetingAllowedPresenters() []string {
	return []string{
		string(OnlineMeetingAllowedPresentersEveryone),
		string(OnlineMeetingAllowedPresentersOrganization),
		string(OnlineMeetingAllowedPresentersOrganizer),
		string(OnlineMeetingAllowedPresentersRoleIsPresenter),
	}
}

func (s *OnlineMeetingAllowedPresenters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingAllowedPresenters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingAllowedPresenters(input string) (*OnlineMeetingAllowedPresenters, error) {
	vals := map[string]OnlineMeetingAllowedPresenters{
		"everyone":        OnlineMeetingAllowedPresentersEveryone,
		"organization":    OnlineMeetingAllowedPresentersOrganization,
		"organizer":       OnlineMeetingAllowedPresentersOrganizer,
		"roleispresenter": OnlineMeetingAllowedPresentersRoleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAllowedPresenters(input)
	return &out, nil
}

type OnlineMeetingShareMeetingChatHistoryDefault string

const (
	OnlineMeetingShareMeetingChatHistoryDefaultAll  OnlineMeetingShareMeetingChatHistoryDefault = "all"
	OnlineMeetingShareMeetingChatHistoryDefaultNone OnlineMeetingShareMeetingChatHistoryDefault = "none"
)

func PossibleValuesForOnlineMeetingShareMeetingChatHistoryDefault() []string {
	return []string{
		string(OnlineMeetingShareMeetingChatHistoryDefaultAll),
		string(OnlineMeetingShareMeetingChatHistoryDefaultNone),
	}
}

func (s *OnlineMeetingShareMeetingChatHistoryDefault) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingShareMeetingChatHistoryDefault(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingShareMeetingChatHistoryDefault(input string) (*OnlineMeetingShareMeetingChatHistoryDefault, error) {
	vals := map[string]OnlineMeetingShareMeetingChatHistoryDefault{
		"all":  OnlineMeetingShareMeetingChatHistoryDefaultAll,
		"none": OnlineMeetingShareMeetingChatHistoryDefaultNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingShareMeetingChatHistoryDefault(input)
	return &out, nil
}

type OpenShiftChangeRequestAssignedTo string

const (
	OpenShiftChangeRequestAssignedToManager   OpenShiftChangeRequestAssignedTo = "manager"
	OpenShiftChangeRequestAssignedToRecipient OpenShiftChangeRequestAssignedTo = "recipient"
	OpenShiftChangeRequestAssignedToSender    OpenShiftChangeRequestAssignedTo = "sender"
	OpenShiftChangeRequestAssignedToSystem    OpenShiftChangeRequestAssignedTo = "system"
)

func PossibleValuesForOpenShiftChangeRequestAssignedTo() []string {
	return []string{
		string(OpenShiftChangeRequestAssignedToManager),
		string(OpenShiftChangeRequestAssignedToRecipient),
		string(OpenShiftChangeRequestAssignedToSender),
		string(OpenShiftChangeRequestAssignedToSystem),
	}
}

func (s *OpenShiftChangeRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenShiftChangeRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenShiftChangeRequestAssignedTo(input string) (*OpenShiftChangeRequestAssignedTo, error) {
	vals := map[string]OpenShiftChangeRequestAssignedTo{
		"manager":   OpenShiftChangeRequestAssignedToManager,
		"recipient": OpenShiftChangeRequestAssignedToRecipient,
		"sender":    OpenShiftChangeRequestAssignedToSender,
		"system":    OpenShiftChangeRequestAssignedToSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftChangeRequestAssignedTo(input)
	return &out, nil
}

type OpenShiftChangeRequestState string

const (
	OpenShiftChangeRequestStateApproved OpenShiftChangeRequestState = "approved"
	OpenShiftChangeRequestStateDeclined OpenShiftChangeRequestState = "declined"
	OpenShiftChangeRequestStatePending  OpenShiftChangeRequestState = "pending"
)

func PossibleValuesForOpenShiftChangeRequestState() []string {
	return []string{
		string(OpenShiftChangeRequestStateApproved),
		string(OpenShiftChangeRequestStateDeclined),
		string(OpenShiftChangeRequestStatePending),
	}
}

func (s *OpenShiftChangeRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenShiftChangeRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenShiftChangeRequestState(input string) (*OpenShiftChangeRequestState, error) {
	vals := map[string]OpenShiftChangeRequestState{
		"approved": OpenShiftChangeRequestStateApproved,
		"declined": OpenShiftChangeRequestStateDeclined,
		"pending":  OpenShiftChangeRequestStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftChangeRequestState(input)
	return &out, nil
}

type OpenShiftItemTheme string

const (
	OpenShiftItemThemeBlue       OpenShiftItemTheme = "blue"
	OpenShiftItemThemeDarkBlue   OpenShiftItemTheme = "darkBlue"
	OpenShiftItemThemeDarkGreen  OpenShiftItemTheme = "darkGreen"
	OpenShiftItemThemeDarkPink   OpenShiftItemTheme = "darkPink"
	OpenShiftItemThemeDarkPurple OpenShiftItemTheme = "darkPurple"
	OpenShiftItemThemeDarkYellow OpenShiftItemTheme = "darkYellow"
	OpenShiftItemThemeGray       OpenShiftItemTheme = "gray"
	OpenShiftItemThemeGreen      OpenShiftItemTheme = "green"
	OpenShiftItemThemePink       OpenShiftItemTheme = "pink"
	OpenShiftItemThemePurple     OpenShiftItemTheme = "purple"
	OpenShiftItemThemeWhite      OpenShiftItemTheme = "white"
	OpenShiftItemThemeYellow     OpenShiftItemTheme = "yellow"
)

func PossibleValuesForOpenShiftItemTheme() []string {
	return []string{
		string(OpenShiftItemThemeBlue),
		string(OpenShiftItemThemeDarkBlue),
		string(OpenShiftItemThemeDarkGreen),
		string(OpenShiftItemThemeDarkPink),
		string(OpenShiftItemThemeDarkPurple),
		string(OpenShiftItemThemeDarkYellow),
		string(OpenShiftItemThemeGray),
		string(OpenShiftItemThemeGreen),
		string(OpenShiftItemThemePink),
		string(OpenShiftItemThemePurple),
		string(OpenShiftItemThemeWhite),
		string(OpenShiftItemThemeYellow),
	}
}

func (s *OpenShiftItemTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenShiftItemTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenShiftItemTheme(input string) (*OpenShiftItemTheme, error) {
	vals := map[string]OpenShiftItemTheme{
		"blue":       OpenShiftItemThemeBlue,
		"darkblue":   OpenShiftItemThemeDarkBlue,
		"darkgreen":  OpenShiftItemThemeDarkGreen,
		"darkpink":   OpenShiftItemThemeDarkPink,
		"darkpurple": OpenShiftItemThemeDarkPurple,
		"darkyellow": OpenShiftItemThemeDarkYellow,
		"gray":       OpenShiftItemThemeGray,
		"green":      OpenShiftItemThemeGreen,
		"pink":       OpenShiftItemThemePink,
		"purple":     OpenShiftItemThemePurple,
		"white":      OpenShiftItemThemeWhite,
		"yellow":     OpenShiftItemThemeYellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftItemTheme(input)
	return &out, nil
}

type OutlookCategoryColor string

const (
	OutlookCategoryColorNone     OutlookCategoryColor = "none"
	OutlookCategoryColorPreset0  OutlookCategoryColor = "preset0"
	OutlookCategoryColorPreset1  OutlookCategoryColor = "preset1"
	OutlookCategoryColorPreset10 OutlookCategoryColor = "preset10"
	OutlookCategoryColorPreset11 OutlookCategoryColor = "preset11"
	OutlookCategoryColorPreset12 OutlookCategoryColor = "preset12"
	OutlookCategoryColorPreset13 OutlookCategoryColor = "preset13"
	OutlookCategoryColorPreset14 OutlookCategoryColor = "preset14"
	OutlookCategoryColorPreset15 OutlookCategoryColor = "preset15"
	OutlookCategoryColorPreset16 OutlookCategoryColor = "preset16"
	OutlookCategoryColorPreset17 OutlookCategoryColor = "preset17"
	OutlookCategoryColorPreset18 OutlookCategoryColor = "preset18"
	OutlookCategoryColorPreset19 OutlookCategoryColor = "preset19"
	OutlookCategoryColorPreset2  OutlookCategoryColor = "preset2"
	OutlookCategoryColorPreset20 OutlookCategoryColor = "preset20"
	OutlookCategoryColorPreset21 OutlookCategoryColor = "preset21"
	OutlookCategoryColorPreset22 OutlookCategoryColor = "preset22"
	OutlookCategoryColorPreset23 OutlookCategoryColor = "preset23"
	OutlookCategoryColorPreset24 OutlookCategoryColor = "preset24"
	OutlookCategoryColorPreset3  OutlookCategoryColor = "preset3"
	OutlookCategoryColorPreset4  OutlookCategoryColor = "preset4"
	OutlookCategoryColorPreset5  OutlookCategoryColor = "preset5"
	OutlookCategoryColorPreset6  OutlookCategoryColor = "preset6"
	OutlookCategoryColorPreset7  OutlookCategoryColor = "preset7"
	OutlookCategoryColorPreset8  OutlookCategoryColor = "preset8"
	OutlookCategoryColorPreset9  OutlookCategoryColor = "preset9"
)

func PossibleValuesForOutlookCategoryColor() []string {
	return []string{
		string(OutlookCategoryColorNone),
		string(OutlookCategoryColorPreset0),
		string(OutlookCategoryColorPreset1),
		string(OutlookCategoryColorPreset10),
		string(OutlookCategoryColorPreset11),
		string(OutlookCategoryColorPreset12),
		string(OutlookCategoryColorPreset13),
		string(OutlookCategoryColorPreset14),
		string(OutlookCategoryColorPreset15),
		string(OutlookCategoryColorPreset16),
		string(OutlookCategoryColorPreset17),
		string(OutlookCategoryColorPreset18),
		string(OutlookCategoryColorPreset19),
		string(OutlookCategoryColorPreset2),
		string(OutlookCategoryColorPreset20),
		string(OutlookCategoryColorPreset21),
		string(OutlookCategoryColorPreset22),
		string(OutlookCategoryColorPreset23),
		string(OutlookCategoryColorPreset24),
		string(OutlookCategoryColorPreset3),
		string(OutlookCategoryColorPreset4),
		string(OutlookCategoryColorPreset5),
		string(OutlookCategoryColorPreset6),
		string(OutlookCategoryColorPreset7),
		string(OutlookCategoryColorPreset8),
		string(OutlookCategoryColorPreset9),
	}
}

func (s *OutlookCategoryColor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookCategoryColor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookCategoryColor(input string) (*OutlookCategoryColor, error) {
	vals := map[string]OutlookCategoryColor{
		"none":     OutlookCategoryColorNone,
		"preset0":  OutlookCategoryColorPreset0,
		"preset1":  OutlookCategoryColorPreset1,
		"preset10": OutlookCategoryColorPreset10,
		"preset11": OutlookCategoryColorPreset11,
		"preset12": OutlookCategoryColorPreset12,
		"preset13": OutlookCategoryColorPreset13,
		"preset14": OutlookCategoryColorPreset14,
		"preset15": OutlookCategoryColorPreset15,
		"preset16": OutlookCategoryColorPreset16,
		"preset17": OutlookCategoryColorPreset17,
		"preset18": OutlookCategoryColorPreset18,
		"preset19": OutlookCategoryColorPreset19,
		"preset2":  OutlookCategoryColorPreset2,
		"preset20": OutlookCategoryColorPreset20,
		"preset21": OutlookCategoryColorPreset21,
		"preset22": OutlookCategoryColorPreset22,
		"preset23": OutlookCategoryColorPreset23,
		"preset24": OutlookCategoryColorPreset24,
		"preset3":  OutlookCategoryColorPreset3,
		"preset4":  OutlookCategoryColorPreset4,
		"preset5":  OutlookCategoryColorPreset5,
		"preset6":  OutlookCategoryColorPreset6,
		"preset7":  OutlookCategoryColorPreset7,
		"preset8":  OutlookCategoryColorPreset8,
		"preset9":  OutlookCategoryColorPreset9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookCategoryColor(input)
	return &out, nil
}

type PhoneAuthenticationMethodPhoneType string

const (
	PhoneAuthenticationMethodPhoneTypeAlternateMobile PhoneAuthenticationMethodPhoneType = "alternateMobile"
	PhoneAuthenticationMethodPhoneTypeMobile          PhoneAuthenticationMethodPhoneType = "mobile"
	PhoneAuthenticationMethodPhoneTypeOffice          PhoneAuthenticationMethodPhoneType = "office"
)

func PossibleValuesForPhoneAuthenticationMethodPhoneType() []string {
	return []string{
		string(PhoneAuthenticationMethodPhoneTypeAlternateMobile),
		string(PhoneAuthenticationMethodPhoneTypeMobile),
		string(PhoneAuthenticationMethodPhoneTypeOffice),
	}
}

func (s *PhoneAuthenticationMethodPhoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhoneAuthenticationMethodPhoneType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhoneAuthenticationMethodPhoneType(input string) (*PhoneAuthenticationMethodPhoneType, error) {
	vals := map[string]PhoneAuthenticationMethodPhoneType{
		"alternatemobile": PhoneAuthenticationMethodPhoneTypeAlternateMobile,
		"mobile":          PhoneAuthenticationMethodPhoneTypeMobile,
		"office":          PhoneAuthenticationMethodPhoneTypeOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneAuthenticationMethodPhoneType(input)
	return &out, nil
}

type PhoneAuthenticationMethodSmsSignInState string

const (
	PhoneAuthenticationMethodSmsSignInStateNotAllowedByPolicy   PhoneAuthenticationMethodSmsSignInState = "notAllowedByPolicy"
	PhoneAuthenticationMethodSmsSignInStateNotConfigured        PhoneAuthenticationMethodSmsSignInState = "notConfigured"
	PhoneAuthenticationMethodSmsSignInStateNotEnabled           PhoneAuthenticationMethodSmsSignInState = "notEnabled"
	PhoneAuthenticationMethodSmsSignInStateNotSupported         PhoneAuthenticationMethodSmsSignInState = "notSupported"
	PhoneAuthenticationMethodSmsSignInStatePhoneNumberNotUnique PhoneAuthenticationMethodSmsSignInState = "phoneNumberNotUnique"
	PhoneAuthenticationMethodSmsSignInStateReady                PhoneAuthenticationMethodSmsSignInState = "ready"
)

func PossibleValuesForPhoneAuthenticationMethodSmsSignInState() []string {
	return []string{
		string(PhoneAuthenticationMethodSmsSignInStateNotAllowedByPolicy),
		string(PhoneAuthenticationMethodSmsSignInStateNotConfigured),
		string(PhoneAuthenticationMethodSmsSignInStateNotEnabled),
		string(PhoneAuthenticationMethodSmsSignInStateNotSupported),
		string(PhoneAuthenticationMethodSmsSignInStatePhoneNumberNotUnique),
		string(PhoneAuthenticationMethodSmsSignInStateReady),
	}
}

func (s *PhoneAuthenticationMethodSmsSignInState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhoneAuthenticationMethodSmsSignInState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhoneAuthenticationMethodSmsSignInState(input string) (*PhoneAuthenticationMethodSmsSignInState, error) {
	vals := map[string]PhoneAuthenticationMethodSmsSignInState{
		"notallowedbypolicy":   PhoneAuthenticationMethodSmsSignInStateNotAllowedByPolicy,
		"notconfigured":        PhoneAuthenticationMethodSmsSignInStateNotConfigured,
		"notenabled":           PhoneAuthenticationMethodSmsSignInStateNotEnabled,
		"notsupported":         PhoneAuthenticationMethodSmsSignInStateNotSupported,
		"phonenumbernotunique": PhoneAuthenticationMethodSmsSignInStatePhoneNumberNotUnique,
		"ready":                PhoneAuthenticationMethodSmsSignInStateReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneAuthenticationMethodSmsSignInState(input)
	return &out, nil
}

type PhoneType string

const (
	PhoneTypeAssistant   PhoneType = "assistant"
	PhoneTypeBusiness    PhoneType = "business"
	PhoneTypeBusinessFax PhoneType = "businessFax"
	PhoneTypeHome        PhoneType = "home"
	PhoneTypeHomeFax     PhoneType = "homeFax"
	PhoneTypeMobile      PhoneType = "mobile"
	PhoneTypeOther       PhoneType = "other"
	PhoneTypeOtherFax    PhoneType = "otherFax"
	PhoneTypePager       PhoneType = "pager"
	PhoneTypeRadio       PhoneType = "radio"
)

func PossibleValuesForPhoneType() []string {
	return []string{
		string(PhoneTypeAssistant),
		string(PhoneTypeBusiness),
		string(PhoneTypeBusinessFax),
		string(PhoneTypeHome),
		string(PhoneTypeHomeFax),
		string(PhoneTypeMobile),
		string(PhoneTypeOther),
		string(PhoneTypeOtherFax),
		string(PhoneTypePager),
		string(PhoneTypeRadio),
	}
}

func (s *PhoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhoneType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhoneType(input string) (*PhoneType, error) {
	vals := map[string]PhoneType{
		"assistant":   PhoneTypeAssistant,
		"business":    PhoneTypeBusiness,
		"businessfax": PhoneTypeBusinessFax,
		"home":        PhoneTypeHome,
		"homefax":     PhoneTypeHomeFax,
		"mobile":      PhoneTypeMobile,
		"other":       PhoneTypeOther,
		"otherfax":    PhoneTypeOtherFax,
		"pager":       PhoneTypePager,
		"radio":       PhoneTypeRadio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneType(input)
	return &out, nil
}

type PlannerPlanContainerType string

const (
	PlannerPlanContainerTypeGroup  PlannerPlanContainerType = "group"
	PlannerPlanContainerTypeRoster PlannerPlanContainerType = "roster"
)

func PossibleValuesForPlannerPlanContainerType() []string {
	return []string{
		string(PlannerPlanContainerTypeGroup),
		string(PlannerPlanContainerTypeRoster),
	}
}

func (s *PlannerPlanContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanContainerType(input string) (*PlannerPlanContainerType, error) {
	vals := map[string]PlannerPlanContainerType{
		"group":  PlannerPlanContainerTypeGroup,
		"roster": PlannerPlanContainerTypeRoster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContainerType(input)
	return &out, nil
}

type PlannerTaskDetailsPreviewType string

const (
	PlannerTaskDetailsPreviewTypeAutomatic   PlannerTaskDetailsPreviewType = "automatic"
	PlannerTaskDetailsPreviewTypeChecklist   PlannerTaskDetailsPreviewType = "checklist"
	PlannerTaskDetailsPreviewTypeDescription PlannerTaskDetailsPreviewType = "description"
	PlannerTaskDetailsPreviewTypeNoPreview   PlannerTaskDetailsPreviewType = "noPreview"
	PlannerTaskDetailsPreviewTypeReference   PlannerTaskDetailsPreviewType = "reference"
)

func PossibleValuesForPlannerTaskDetailsPreviewType() []string {
	return []string{
		string(PlannerTaskDetailsPreviewTypeAutomatic),
		string(PlannerTaskDetailsPreviewTypeChecklist),
		string(PlannerTaskDetailsPreviewTypeDescription),
		string(PlannerTaskDetailsPreviewTypeNoPreview),
		string(PlannerTaskDetailsPreviewTypeReference),
	}
}

func (s *PlannerTaskDetailsPreviewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskDetailsPreviewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskDetailsPreviewType(input string) (*PlannerTaskDetailsPreviewType, error) {
	vals := map[string]PlannerTaskDetailsPreviewType{
		"automatic":   PlannerTaskDetailsPreviewTypeAutomatic,
		"checklist":   PlannerTaskDetailsPreviewTypeChecklist,
		"description": PlannerTaskDetailsPreviewTypeDescription,
		"nopreview":   PlannerTaskDetailsPreviewTypeNoPreview,
		"reference":   PlannerTaskDetailsPreviewTypeReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskDetailsPreviewType(input)
	return &out, nil
}

type PlannerTaskPreviewType string

const (
	PlannerTaskPreviewTypeAutomatic   PlannerTaskPreviewType = "automatic"
	PlannerTaskPreviewTypeChecklist   PlannerTaskPreviewType = "checklist"
	PlannerTaskPreviewTypeDescription PlannerTaskPreviewType = "description"
	PlannerTaskPreviewTypeNoPreview   PlannerTaskPreviewType = "noPreview"
	PlannerTaskPreviewTypeReference   PlannerTaskPreviewType = "reference"
)

func PossibleValuesForPlannerTaskPreviewType() []string {
	return []string{
		string(PlannerTaskPreviewTypeAutomatic),
		string(PlannerTaskPreviewTypeChecklist),
		string(PlannerTaskPreviewTypeDescription),
		string(PlannerTaskPreviewTypeNoPreview),
		string(PlannerTaskPreviewTypeReference),
	}
}

func (s *PlannerTaskPreviewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskPreviewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskPreviewType(input string) (*PlannerTaskPreviewType, error) {
	vals := map[string]PlannerTaskPreviewType{
		"automatic":   PlannerTaskPreviewTypeAutomatic,
		"checklist":   PlannerTaskPreviewTypeChecklist,
		"description": PlannerTaskPreviewTypeDescription,
		"nopreview":   PlannerTaskPreviewTypeNoPreview,
		"reference":   PlannerTaskPreviewTypeReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskPreviewType(input)
	return &out, nil
}

type PrintJobConfigurationColorMode string

const (
	PrintJobConfigurationColorModeAuto          PrintJobConfigurationColorMode = "auto"
	PrintJobConfigurationColorModeBlackAndWhite PrintJobConfigurationColorMode = "blackAndWhite"
	PrintJobConfigurationColorModeColor         PrintJobConfigurationColorMode = "color"
	PrintJobConfigurationColorModeGrayscale     PrintJobConfigurationColorMode = "grayscale"
)

func PossibleValuesForPrintJobConfigurationColorMode() []string {
	return []string{
		string(PrintJobConfigurationColorModeAuto),
		string(PrintJobConfigurationColorModeBlackAndWhite),
		string(PrintJobConfigurationColorModeColor),
		string(PrintJobConfigurationColorModeGrayscale),
	}
}

func (s *PrintJobConfigurationColorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationColorMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationColorMode(input string) (*PrintJobConfigurationColorMode, error) {
	vals := map[string]PrintJobConfigurationColorMode{
		"auto":          PrintJobConfigurationColorModeAuto,
		"blackandwhite": PrintJobConfigurationColorModeBlackAndWhite,
		"color":         PrintJobConfigurationColorModeColor,
		"grayscale":     PrintJobConfigurationColorModeGrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationColorMode(input)
	return &out, nil
}

type PrintJobConfigurationDuplexMode string

const (
	PrintJobConfigurationDuplexModeFlipOnLongEdge  PrintJobConfigurationDuplexMode = "flipOnLongEdge"
	PrintJobConfigurationDuplexModeFlipOnShortEdge PrintJobConfigurationDuplexMode = "flipOnShortEdge"
	PrintJobConfigurationDuplexModeOneSided        PrintJobConfigurationDuplexMode = "oneSided"
)

func PossibleValuesForPrintJobConfigurationDuplexMode() []string {
	return []string{
		string(PrintJobConfigurationDuplexModeFlipOnLongEdge),
		string(PrintJobConfigurationDuplexModeFlipOnShortEdge),
		string(PrintJobConfigurationDuplexModeOneSided),
	}
}

func (s *PrintJobConfigurationDuplexMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationDuplexMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationDuplexMode(input string) (*PrintJobConfigurationDuplexMode, error) {
	vals := map[string]PrintJobConfigurationDuplexMode{
		"fliponlongedge":  PrintJobConfigurationDuplexModeFlipOnLongEdge,
		"fliponshortedge": PrintJobConfigurationDuplexModeFlipOnShortEdge,
		"onesided":        PrintJobConfigurationDuplexModeOneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationDuplexMode(input)
	return &out, nil
}

type PrintJobConfigurationFeedOrientation string

const (
	PrintJobConfigurationFeedOrientationLongEdgeFirst  PrintJobConfigurationFeedOrientation = "longEdgeFirst"
	PrintJobConfigurationFeedOrientationShortEdgeFirst PrintJobConfigurationFeedOrientation = "shortEdgeFirst"
)

func PossibleValuesForPrintJobConfigurationFeedOrientation() []string {
	return []string{
		string(PrintJobConfigurationFeedOrientationLongEdgeFirst),
		string(PrintJobConfigurationFeedOrientationShortEdgeFirst),
	}
}

func (s *PrintJobConfigurationFeedOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationFeedOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationFeedOrientation(input string) (*PrintJobConfigurationFeedOrientation, error) {
	vals := map[string]PrintJobConfigurationFeedOrientation{
		"longedgefirst":  PrintJobConfigurationFeedOrientationLongEdgeFirst,
		"shortedgefirst": PrintJobConfigurationFeedOrientationShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationFeedOrientation(input)
	return &out, nil
}

type PrintJobConfigurationFinishings string

const (
	PrintJobConfigurationFinishingsBind              PrintJobConfigurationFinishings = "bind"
	PrintJobConfigurationFinishingsCover             PrintJobConfigurationFinishings = "cover"
	PrintJobConfigurationFinishingsNone              PrintJobConfigurationFinishings = "none"
	PrintJobConfigurationFinishingsPunch             PrintJobConfigurationFinishings = "punch"
	PrintJobConfigurationFinishingsSaddleStitch      PrintJobConfigurationFinishings = "saddleStitch"
	PrintJobConfigurationFinishingsStaple            PrintJobConfigurationFinishings = "staple"
	PrintJobConfigurationFinishingsStapleBottomLeft  PrintJobConfigurationFinishings = "stapleBottomLeft"
	PrintJobConfigurationFinishingsStapleBottomRight PrintJobConfigurationFinishings = "stapleBottomRight"
	PrintJobConfigurationFinishingsStapleDualBottom  PrintJobConfigurationFinishings = "stapleDualBottom"
	PrintJobConfigurationFinishingsStapleDualLeft    PrintJobConfigurationFinishings = "stapleDualLeft"
	PrintJobConfigurationFinishingsStapleDualRight   PrintJobConfigurationFinishings = "stapleDualRight"
	PrintJobConfigurationFinishingsStapleDualTop     PrintJobConfigurationFinishings = "stapleDualTop"
	PrintJobConfigurationFinishingsStapleTopLeft     PrintJobConfigurationFinishings = "stapleTopLeft"
	PrintJobConfigurationFinishingsStapleTopRight    PrintJobConfigurationFinishings = "stapleTopRight"
	PrintJobConfigurationFinishingsStitchBottomEdge  PrintJobConfigurationFinishings = "stitchBottomEdge"
	PrintJobConfigurationFinishingsStitchEdge        PrintJobConfigurationFinishings = "stitchEdge"
	PrintJobConfigurationFinishingsStitchLeftEdge    PrintJobConfigurationFinishings = "stitchLeftEdge"
	PrintJobConfigurationFinishingsStitchRightEdge   PrintJobConfigurationFinishings = "stitchRightEdge"
	PrintJobConfigurationFinishingsStitchTopEdge     PrintJobConfigurationFinishings = "stitchTopEdge"
)

func PossibleValuesForPrintJobConfigurationFinishings() []string {
	return []string{
		string(PrintJobConfigurationFinishingsBind),
		string(PrintJobConfigurationFinishingsCover),
		string(PrintJobConfigurationFinishingsNone),
		string(PrintJobConfigurationFinishingsPunch),
		string(PrintJobConfigurationFinishingsSaddleStitch),
		string(PrintJobConfigurationFinishingsStaple),
		string(PrintJobConfigurationFinishingsStapleBottomLeft),
		string(PrintJobConfigurationFinishingsStapleBottomRight),
		string(PrintJobConfigurationFinishingsStapleDualBottom),
		string(PrintJobConfigurationFinishingsStapleDualLeft),
		string(PrintJobConfigurationFinishingsStapleDualRight),
		string(PrintJobConfigurationFinishingsStapleDualTop),
		string(PrintJobConfigurationFinishingsStapleTopLeft),
		string(PrintJobConfigurationFinishingsStapleTopRight),
		string(PrintJobConfigurationFinishingsStitchBottomEdge),
		string(PrintJobConfigurationFinishingsStitchEdge),
		string(PrintJobConfigurationFinishingsStitchLeftEdge),
		string(PrintJobConfigurationFinishingsStitchRightEdge),
		string(PrintJobConfigurationFinishingsStitchTopEdge),
	}
}

func (s *PrintJobConfigurationFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationFinishings(input string) (*PrintJobConfigurationFinishings, error) {
	vals := map[string]PrintJobConfigurationFinishings{
		"bind":              PrintJobConfigurationFinishingsBind,
		"cover":             PrintJobConfigurationFinishingsCover,
		"none":              PrintJobConfigurationFinishingsNone,
		"punch":             PrintJobConfigurationFinishingsPunch,
		"saddlestitch":      PrintJobConfigurationFinishingsSaddleStitch,
		"staple":            PrintJobConfigurationFinishingsStaple,
		"staplebottomleft":  PrintJobConfigurationFinishingsStapleBottomLeft,
		"staplebottomright": PrintJobConfigurationFinishingsStapleBottomRight,
		"stapledualbottom":  PrintJobConfigurationFinishingsStapleDualBottom,
		"stapledualleft":    PrintJobConfigurationFinishingsStapleDualLeft,
		"stapledualright":   PrintJobConfigurationFinishingsStapleDualRight,
		"stapledualtop":     PrintJobConfigurationFinishingsStapleDualTop,
		"stapletopleft":     PrintJobConfigurationFinishingsStapleTopLeft,
		"stapletopright":    PrintJobConfigurationFinishingsStapleTopRight,
		"stitchbottomedge":  PrintJobConfigurationFinishingsStitchBottomEdge,
		"stitchedge":        PrintJobConfigurationFinishingsStitchEdge,
		"stitchleftedge":    PrintJobConfigurationFinishingsStitchLeftEdge,
		"stitchrightedge":   PrintJobConfigurationFinishingsStitchRightEdge,
		"stitchtopedge":     PrintJobConfigurationFinishingsStitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationFinishings(input)
	return &out, nil
}

type PrintJobConfigurationMultipageLayout string

const (
	PrintJobConfigurationMultipageLayoutClockwiseFromBottomLeft         PrintJobConfigurationMultipageLayout = "clockwiseFromBottomLeft"
	PrintJobConfigurationMultipageLayoutClockwiseFromBottomRight        PrintJobConfigurationMultipageLayout = "clockwiseFromBottomRight"
	PrintJobConfigurationMultipageLayoutClockwiseFromTopLeft            PrintJobConfigurationMultipageLayout = "clockwiseFromTopLeft"
	PrintJobConfigurationMultipageLayoutClockwiseFromTopRight           PrintJobConfigurationMultipageLayout = "clockwiseFromTopRight"
	PrintJobConfigurationMultipageLayoutCounterclockwiseFromBottomLeft  PrintJobConfigurationMultipageLayout = "counterclockwiseFromBottomLeft"
	PrintJobConfigurationMultipageLayoutCounterclockwiseFromBottomRight PrintJobConfigurationMultipageLayout = "counterclockwiseFromBottomRight"
	PrintJobConfigurationMultipageLayoutCounterclockwiseFromTopLeft     PrintJobConfigurationMultipageLayout = "counterclockwiseFromTopLeft"
	PrintJobConfigurationMultipageLayoutCounterclockwiseFromTopRight    PrintJobConfigurationMultipageLayout = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrintJobConfigurationMultipageLayout() []string {
	return []string{
		string(PrintJobConfigurationMultipageLayoutClockwiseFromBottomLeft),
		string(PrintJobConfigurationMultipageLayoutClockwiseFromBottomRight),
		string(PrintJobConfigurationMultipageLayoutClockwiseFromTopLeft),
		string(PrintJobConfigurationMultipageLayoutClockwiseFromTopRight),
		string(PrintJobConfigurationMultipageLayoutCounterclockwiseFromBottomLeft),
		string(PrintJobConfigurationMultipageLayoutCounterclockwiseFromBottomRight),
		string(PrintJobConfigurationMultipageLayoutCounterclockwiseFromTopLeft),
		string(PrintJobConfigurationMultipageLayoutCounterclockwiseFromTopRight),
	}
}

func (s *PrintJobConfigurationMultipageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationMultipageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationMultipageLayout(input string) (*PrintJobConfigurationMultipageLayout, error) {
	vals := map[string]PrintJobConfigurationMultipageLayout{
		"clockwisefrombottomleft":         PrintJobConfigurationMultipageLayoutClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrintJobConfigurationMultipageLayoutClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrintJobConfigurationMultipageLayoutClockwiseFromTopLeft,
		"clockwisefromtopright":           PrintJobConfigurationMultipageLayoutClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrintJobConfigurationMultipageLayoutCounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrintJobConfigurationMultipageLayoutCounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrintJobConfigurationMultipageLayoutCounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrintJobConfigurationMultipageLayoutCounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationMultipageLayout(input)
	return &out, nil
}

type PrintJobConfigurationOrientation string

const (
	PrintJobConfigurationOrientationLandscape        PrintJobConfigurationOrientation = "landscape"
	PrintJobConfigurationOrientationPortrait         PrintJobConfigurationOrientation = "portrait"
	PrintJobConfigurationOrientationReverseLandscape PrintJobConfigurationOrientation = "reverseLandscape"
	PrintJobConfigurationOrientationReversePortrait  PrintJobConfigurationOrientation = "reversePortrait"
)

func PossibleValuesForPrintJobConfigurationOrientation() []string {
	return []string{
		string(PrintJobConfigurationOrientationLandscape),
		string(PrintJobConfigurationOrientationPortrait),
		string(PrintJobConfigurationOrientationReverseLandscape),
		string(PrintJobConfigurationOrientationReversePortrait),
	}
}

func (s *PrintJobConfigurationOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationOrientation(input string) (*PrintJobConfigurationOrientation, error) {
	vals := map[string]PrintJobConfigurationOrientation{
		"landscape":        PrintJobConfigurationOrientationLandscape,
		"portrait":         PrintJobConfigurationOrientationPortrait,
		"reverselandscape": PrintJobConfigurationOrientationReverseLandscape,
		"reverseportrait":  PrintJobConfigurationOrientationReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationOrientation(input)
	return &out, nil
}

type PrintJobConfigurationQuality string

const (
	PrintJobConfigurationQualityHigh   PrintJobConfigurationQuality = "high"
	PrintJobConfigurationQualityLow    PrintJobConfigurationQuality = "low"
	PrintJobConfigurationQualityMedium PrintJobConfigurationQuality = "medium"
)

func PossibleValuesForPrintJobConfigurationQuality() []string {
	return []string{
		string(PrintJobConfigurationQualityHigh),
		string(PrintJobConfigurationQualityLow),
		string(PrintJobConfigurationQualityMedium),
	}
}

func (s *PrintJobConfigurationQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationQuality(input string) (*PrintJobConfigurationQuality, error) {
	vals := map[string]PrintJobConfigurationQuality{
		"high":   PrintJobConfigurationQualityHigh,
		"low":    PrintJobConfigurationQualityLow,
		"medium": PrintJobConfigurationQualityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationQuality(input)
	return &out, nil
}

type PrintJobConfigurationScaling string

const (
	PrintJobConfigurationScalingAuto        PrintJobConfigurationScaling = "auto"
	PrintJobConfigurationScalingFill        PrintJobConfigurationScaling = "fill"
	PrintJobConfigurationScalingFit         PrintJobConfigurationScaling = "fit"
	PrintJobConfigurationScalingNone        PrintJobConfigurationScaling = "none"
	PrintJobConfigurationScalingShrinkToFit PrintJobConfigurationScaling = "shrinkToFit"
)

func PossibleValuesForPrintJobConfigurationScaling() []string {
	return []string{
		string(PrintJobConfigurationScalingAuto),
		string(PrintJobConfigurationScalingFill),
		string(PrintJobConfigurationScalingFit),
		string(PrintJobConfigurationScalingNone),
		string(PrintJobConfigurationScalingShrinkToFit),
	}
}

func (s *PrintJobConfigurationScaling) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationScaling(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationScaling(input string) (*PrintJobConfigurationScaling, error) {
	vals := map[string]PrintJobConfigurationScaling{
		"auto":        PrintJobConfigurationScalingAuto,
		"fill":        PrintJobConfigurationScalingFill,
		"fit":         PrintJobConfigurationScalingFit,
		"none":        PrintJobConfigurationScalingNone,
		"shrinktofit": PrintJobConfigurationScalingShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationScaling(input)
	return &out, nil
}

type PrintJobStatusDetails string

const (
	PrintJobStatusDetailsCompletedSuccessfully PrintJobStatusDetails = "completedSuccessfully"
	PrintJobStatusDetailsCompletedWithErrors   PrintJobStatusDetails = "completedWithErrors"
	PrintJobStatusDetailsCompletedWithWarnings PrintJobStatusDetails = "completedWithWarnings"
	PrintJobStatusDetailsInterpreting          PrintJobStatusDetails = "interpreting"
	PrintJobStatusDetailsReleaseWait           PrintJobStatusDetails = "releaseWait"
	PrintJobStatusDetailsTransforming          PrintJobStatusDetails = "transforming"
	PrintJobStatusDetailsUploadPending         PrintJobStatusDetails = "uploadPending"
)

func PossibleValuesForPrintJobStatusDetails() []string {
	return []string{
		string(PrintJobStatusDetailsCompletedSuccessfully),
		string(PrintJobStatusDetailsCompletedWithErrors),
		string(PrintJobStatusDetailsCompletedWithWarnings),
		string(PrintJobStatusDetailsInterpreting),
		string(PrintJobStatusDetailsReleaseWait),
		string(PrintJobStatusDetailsTransforming),
		string(PrintJobStatusDetailsUploadPending),
	}
}

func (s *PrintJobStatusDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobStatusDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobStatusDetails(input string) (*PrintJobStatusDetails, error) {
	vals := map[string]PrintJobStatusDetails{
		"completedsuccessfully": PrintJobStatusDetailsCompletedSuccessfully,
		"completedwitherrors":   PrintJobStatusDetailsCompletedWithErrors,
		"completedwithwarnings": PrintJobStatusDetailsCompletedWithWarnings,
		"interpreting":          PrintJobStatusDetailsInterpreting,
		"releasewait":           PrintJobStatusDetailsReleaseWait,
		"transforming":          PrintJobStatusDetailsTransforming,
		"uploadpending":         PrintJobStatusDetailsUploadPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusDetails(input)
	return &out, nil
}

type PrintJobStatusState string

const (
	PrintJobStatusStateAborted    PrintJobStatusState = "aborted"
	PrintJobStatusStateCanceled   PrintJobStatusState = "canceled"
	PrintJobStatusStateCompleted  PrintJobStatusState = "completed"
	PrintJobStatusStatePaused     PrintJobStatusState = "paused"
	PrintJobStatusStatePending    PrintJobStatusState = "pending"
	PrintJobStatusStateProcessing PrintJobStatusState = "processing"
	PrintJobStatusStateStopped    PrintJobStatusState = "stopped"
	PrintJobStatusStateUnknown    PrintJobStatusState = "unknown"
)

func PossibleValuesForPrintJobStatusState() []string {
	return []string{
		string(PrintJobStatusStateAborted),
		string(PrintJobStatusStateCanceled),
		string(PrintJobStatusStateCompleted),
		string(PrintJobStatusStatePaused),
		string(PrintJobStatusStatePending),
		string(PrintJobStatusStateProcessing),
		string(PrintJobStatusStateStopped),
		string(PrintJobStatusStateUnknown),
	}
}

func (s *PrintJobStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobStatusState(input string) (*PrintJobStatusState, error) {
	vals := map[string]PrintJobStatusState{
		"aborted":    PrintJobStatusStateAborted,
		"canceled":   PrintJobStatusStateCanceled,
		"completed":  PrintJobStatusStateCompleted,
		"paused":     PrintJobStatusStatePaused,
		"pending":    PrintJobStatusStatePending,
		"processing": PrintJobStatusStateProcessing,
		"stopped":    PrintJobStatusStateStopped,
		"unknown":    PrintJobStatusStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusState(input)
	return &out, nil
}

type PrintTaskStatusState string

const (
	PrintTaskStatusStateAborted    PrintTaskStatusState = "aborted"
	PrintTaskStatusStateCompleted  PrintTaskStatusState = "completed"
	PrintTaskStatusStatePending    PrintTaskStatusState = "pending"
	PrintTaskStatusStateProcessing PrintTaskStatusState = "processing"
)

func PossibleValuesForPrintTaskStatusState() []string {
	return []string{
		string(PrintTaskStatusStateAborted),
		string(PrintTaskStatusStateCompleted),
		string(PrintTaskStatusStatePending),
		string(PrintTaskStatusStateProcessing),
	}
}

func (s *PrintTaskStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintTaskStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintTaskStatusState(input string) (*PrintTaskStatusState, error) {
	vals := map[string]PrintTaskStatusState{
		"aborted":    PrintTaskStatusStateAborted,
		"completed":  PrintTaskStatusStateCompleted,
		"pending":    PrintTaskStatusStatePending,
		"processing": PrintTaskStatusStateProcessing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintTaskStatusState(input)
	return &out, nil
}

type PrintTaskTriggerEvent string

const (
	PrintTaskTriggerEventJobStarted PrintTaskTriggerEvent = "jobStarted"
)

func PossibleValuesForPrintTaskTriggerEvent() []string {
	return []string{
		string(PrintTaskTriggerEventJobStarted),
	}
}

func (s *PrintTaskTriggerEvent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintTaskTriggerEvent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintTaskTriggerEvent(input string) (*PrintTaskTriggerEvent, error) {
	vals := map[string]PrintTaskTriggerEvent{
		"jobstarted": PrintTaskTriggerEventJobStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintTaskTriggerEvent(input)
	return &out, nil
}

type PrinterCapabilitiesColorModes string

const (
	PrinterCapabilitiesColorModesAuto          PrinterCapabilitiesColorModes = "auto"
	PrinterCapabilitiesColorModesBlackAndWhite PrinterCapabilitiesColorModes = "blackAndWhite"
	PrinterCapabilitiesColorModesColor         PrinterCapabilitiesColorModes = "color"
	PrinterCapabilitiesColorModesGrayscale     PrinterCapabilitiesColorModes = "grayscale"
)

func PossibleValuesForPrinterCapabilitiesColorModes() []string {
	return []string{
		string(PrinterCapabilitiesColorModesAuto),
		string(PrinterCapabilitiesColorModesBlackAndWhite),
		string(PrinterCapabilitiesColorModesColor),
		string(PrinterCapabilitiesColorModesGrayscale),
	}
}

func (s *PrinterCapabilitiesColorModes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesColorModes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesColorModes(input string) (*PrinterCapabilitiesColorModes, error) {
	vals := map[string]PrinterCapabilitiesColorModes{
		"auto":          PrinterCapabilitiesColorModesAuto,
		"blackandwhite": PrinterCapabilitiesColorModesBlackAndWhite,
		"color":         PrinterCapabilitiesColorModesColor,
		"grayscale":     PrinterCapabilitiesColorModesGrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesColorModes(input)
	return &out, nil
}

type PrinterCapabilitiesDuplexModes string

const (
	PrinterCapabilitiesDuplexModesFlipOnLongEdge  PrinterCapabilitiesDuplexModes = "flipOnLongEdge"
	PrinterCapabilitiesDuplexModesFlipOnShortEdge PrinterCapabilitiesDuplexModes = "flipOnShortEdge"
	PrinterCapabilitiesDuplexModesOneSided        PrinterCapabilitiesDuplexModes = "oneSided"
)

func PossibleValuesForPrinterCapabilitiesDuplexModes() []string {
	return []string{
		string(PrinterCapabilitiesDuplexModesFlipOnLongEdge),
		string(PrinterCapabilitiesDuplexModesFlipOnShortEdge),
		string(PrinterCapabilitiesDuplexModesOneSided),
	}
}

func (s *PrinterCapabilitiesDuplexModes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesDuplexModes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesDuplexModes(input string) (*PrinterCapabilitiesDuplexModes, error) {
	vals := map[string]PrinterCapabilitiesDuplexModes{
		"fliponlongedge":  PrinterCapabilitiesDuplexModesFlipOnLongEdge,
		"fliponshortedge": PrinterCapabilitiesDuplexModesFlipOnShortEdge,
		"onesided":        PrinterCapabilitiesDuplexModesOneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesDuplexModes(input)
	return &out, nil
}

type PrinterCapabilitiesFeedOrientations string

const (
	PrinterCapabilitiesFeedOrientationsLongEdgeFirst  PrinterCapabilitiesFeedOrientations = "longEdgeFirst"
	PrinterCapabilitiesFeedOrientationsShortEdgeFirst PrinterCapabilitiesFeedOrientations = "shortEdgeFirst"
)

func PossibleValuesForPrinterCapabilitiesFeedOrientations() []string {
	return []string{
		string(PrinterCapabilitiesFeedOrientationsLongEdgeFirst),
		string(PrinterCapabilitiesFeedOrientationsShortEdgeFirst),
	}
}

func (s *PrinterCapabilitiesFeedOrientations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesFeedOrientations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesFeedOrientations(input string) (*PrinterCapabilitiesFeedOrientations, error) {
	vals := map[string]PrinterCapabilitiesFeedOrientations{
		"longedgefirst":  PrinterCapabilitiesFeedOrientationsLongEdgeFirst,
		"shortedgefirst": PrinterCapabilitiesFeedOrientationsShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFeedOrientations(input)
	return &out, nil
}

type PrinterCapabilitiesFinishings string

const (
	PrinterCapabilitiesFinishingsBind              PrinterCapabilitiesFinishings = "bind"
	PrinterCapabilitiesFinishingsCover             PrinterCapabilitiesFinishings = "cover"
	PrinterCapabilitiesFinishingsNone              PrinterCapabilitiesFinishings = "none"
	PrinterCapabilitiesFinishingsPunch             PrinterCapabilitiesFinishings = "punch"
	PrinterCapabilitiesFinishingsSaddleStitch      PrinterCapabilitiesFinishings = "saddleStitch"
	PrinterCapabilitiesFinishingsStaple            PrinterCapabilitiesFinishings = "staple"
	PrinterCapabilitiesFinishingsStapleBottomLeft  PrinterCapabilitiesFinishings = "stapleBottomLeft"
	PrinterCapabilitiesFinishingsStapleBottomRight PrinterCapabilitiesFinishings = "stapleBottomRight"
	PrinterCapabilitiesFinishingsStapleDualBottom  PrinterCapabilitiesFinishings = "stapleDualBottom"
	PrinterCapabilitiesFinishingsStapleDualLeft    PrinterCapabilitiesFinishings = "stapleDualLeft"
	PrinterCapabilitiesFinishingsStapleDualRight   PrinterCapabilitiesFinishings = "stapleDualRight"
	PrinterCapabilitiesFinishingsStapleDualTop     PrinterCapabilitiesFinishings = "stapleDualTop"
	PrinterCapabilitiesFinishingsStapleTopLeft     PrinterCapabilitiesFinishings = "stapleTopLeft"
	PrinterCapabilitiesFinishingsStapleTopRight    PrinterCapabilitiesFinishings = "stapleTopRight"
	PrinterCapabilitiesFinishingsStitchBottomEdge  PrinterCapabilitiesFinishings = "stitchBottomEdge"
	PrinterCapabilitiesFinishingsStitchEdge        PrinterCapabilitiesFinishings = "stitchEdge"
	PrinterCapabilitiesFinishingsStitchLeftEdge    PrinterCapabilitiesFinishings = "stitchLeftEdge"
	PrinterCapabilitiesFinishingsStitchRightEdge   PrinterCapabilitiesFinishings = "stitchRightEdge"
	PrinterCapabilitiesFinishingsStitchTopEdge     PrinterCapabilitiesFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterCapabilitiesFinishings() []string {
	return []string{
		string(PrinterCapabilitiesFinishingsBind),
		string(PrinterCapabilitiesFinishingsCover),
		string(PrinterCapabilitiesFinishingsNone),
		string(PrinterCapabilitiesFinishingsPunch),
		string(PrinterCapabilitiesFinishingsSaddleStitch),
		string(PrinterCapabilitiesFinishingsStaple),
		string(PrinterCapabilitiesFinishingsStapleBottomLeft),
		string(PrinterCapabilitiesFinishingsStapleBottomRight),
		string(PrinterCapabilitiesFinishingsStapleDualBottom),
		string(PrinterCapabilitiesFinishingsStapleDualLeft),
		string(PrinterCapabilitiesFinishingsStapleDualRight),
		string(PrinterCapabilitiesFinishingsStapleDualTop),
		string(PrinterCapabilitiesFinishingsStapleTopLeft),
		string(PrinterCapabilitiesFinishingsStapleTopRight),
		string(PrinterCapabilitiesFinishingsStitchBottomEdge),
		string(PrinterCapabilitiesFinishingsStitchEdge),
		string(PrinterCapabilitiesFinishingsStitchLeftEdge),
		string(PrinterCapabilitiesFinishingsStitchRightEdge),
		string(PrinterCapabilitiesFinishingsStitchTopEdge),
	}
}

func (s *PrinterCapabilitiesFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesFinishings(input string) (*PrinterCapabilitiesFinishings, error) {
	vals := map[string]PrinterCapabilitiesFinishings{
		"bind":              PrinterCapabilitiesFinishingsBind,
		"cover":             PrinterCapabilitiesFinishingsCover,
		"none":              PrinterCapabilitiesFinishingsNone,
		"punch":             PrinterCapabilitiesFinishingsPunch,
		"saddlestitch":      PrinterCapabilitiesFinishingsSaddleStitch,
		"staple":            PrinterCapabilitiesFinishingsStaple,
		"staplebottomleft":  PrinterCapabilitiesFinishingsStapleBottomLeft,
		"staplebottomright": PrinterCapabilitiesFinishingsStapleBottomRight,
		"stapledualbottom":  PrinterCapabilitiesFinishingsStapleDualBottom,
		"stapledualleft":    PrinterCapabilitiesFinishingsStapleDualLeft,
		"stapledualright":   PrinterCapabilitiesFinishingsStapleDualRight,
		"stapledualtop":     PrinterCapabilitiesFinishingsStapleDualTop,
		"stapletopleft":     PrinterCapabilitiesFinishingsStapleTopLeft,
		"stapletopright":    PrinterCapabilitiesFinishingsStapleTopRight,
		"stitchbottomedge":  PrinterCapabilitiesFinishingsStitchBottomEdge,
		"stitchedge":        PrinterCapabilitiesFinishingsStitchEdge,
		"stitchleftedge":    PrinterCapabilitiesFinishingsStitchLeftEdge,
		"stitchrightedge":   PrinterCapabilitiesFinishingsStitchRightEdge,
		"stitchtopedge":     PrinterCapabilitiesFinishingsStitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFinishings(input)
	return &out, nil
}

type PrinterCapabilitiesMultipageLayouts string

const (
	PrinterCapabilitiesMultipageLayoutsClockwiseFromBottomLeft         PrinterCapabilitiesMultipageLayouts = "clockwiseFromBottomLeft"
	PrinterCapabilitiesMultipageLayoutsClockwiseFromBottomRight        PrinterCapabilitiesMultipageLayouts = "clockwiseFromBottomRight"
	PrinterCapabilitiesMultipageLayoutsClockwiseFromTopLeft            PrinterCapabilitiesMultipageLayouts = "clockwiseFromTopLeft"
	PrinterCapabilitiesMultipageLayoutsClockwiseFromTopRight           PrinterCapabilitiesMultipageLayouts = "clockwiseFromTopRight"
	PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromBottomLeft  PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromBottomLeft"
	PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromBottomRight PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromBottomRight"
	PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromTopLeft     PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromTopLeft"
	PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromTopRight    PrinterCapabilitiesMultipageLayouts = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrinterCapabilitiesMultipageLayouts() []string {
	return []string{
		string(PrinterCapabilitiesMultipageLayoutsClockwiseFromBottomLeft),
		string(PrinterCapabilitiesMultipageLayoutsClockwiseFromBottomRight),
		string(PrinterCapabilitiesMultipageLayoutsClockwiseFromTopLeft),
		string(PrinterCapabilitiesMultipageLayoutsClockwiseFromTopRight),
		string(PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromBottomLeft),
		string(PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromBottomRight),
		string(PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromTopLeft),
		string(PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromTopRight),
	}
}

func (s *PrinterCapabilitiesMultipageLayouts) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesMultipageLayouts(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesMultipageLayouts(input string) (*PrinterCapabilitiesMultipageLayouts, error) {
	vals := map[string]PrinterCapabilitiesMultipageLayouts{
		"clockwisefrombottomleft":         PrinterCapabilitiesMultipageLayoutsClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterCapabilitiesMultipageLayoutsClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterCapabilitiesMultipageLayoutsClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterCapabilitiesMultipageLayoutsClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterCapabilitiesMultipageLayoutsCounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesMultipageLayouts(input)
	return &out, nil
}

type PrinterCapabilitiesOrientations string

const (
	PrinterCapabilitiesOrientationsLandscape        PrinterCapabilitiesOrientations = "landscape"
	PrinterCapabilitiesOrientationsPortrait         PrinterCapabilitiesOrientations = "portrait"
	PrinterCapabilitiesOrientationsReverseLandscape PrinterCapabilitiesOrientations = "reverseLandscape"
	PrinterCapabilitiesOrientationsReversePortrait  PrinterCapabilitiesOrientations = "reversePortrait"
)

func PossibleValuesForPrinterCapabilitiesOrientations() []string {
	return []string{
		string(PrinterCapabilitiesOrientationsLandscape),
		string(PrinterCapabilitiesOrientationsPortrait),
		string(PrinterCapabilitiesOrientationsReverseLandscape),
		string(PrinterCapabilitiesOrientationsReversePortrait),
	}
}

func (s *PrinterCapabilitiesOrientations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesOrientations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesOrientations(input string) (*PrinterCapabilitiesOrientations, error) {
	vals := map[string]PrinterCapabilitiesOrientations{
		"landscape":        PrinterCapabilitiesOrientationsLandscape,
		"portrait":         PrinterCapabilitiesOrientationsPortrait,
		"reverselandscape": PrinterCapabilitiesOrientationsReverseLandscape,
		"reverseportrait":  PrinterCapabilitiesOrientationsReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesOrientations(input)
	return &out, nil
}

type PrinterCapabilitiesQualities string

const (
	PrinterCapabilitiesQualitiesHigh   PrinterCapabilitiesQualities = "high"
	PrinterCapabilitiesQualitiesLow    PrinterCapabilitiesQualities = "low"
	PrinterCapabilitiesQualitiesMedium PrinterCapabilitiesQualities = "medium"
)

func PossibleValuesForPrinterCapabilitiesQualities() []string {
	return []string{
		string(PrinterCapabilitiesQualitiesHigh),
		string(PrinterCapabilitiesQualitiesLow),
		string(PrinterCapabilitiesQualitiesMedium),
	}
}

func (s *PrinterCapabilitiesQualities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesQualities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesQualities(input string) (*PrinterCapabilitiesQualities, error) {
	vals := map[string]PrinterCapabilitiesQualities{
		"high":   PrinterCapabilitiesQualitiesHigh,
		"low":    PrinterCapabilitiesQualitiesLow,
		"medium": PrinterCapabilitiesQualitiesMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesQualities(input)
	return &out, nil
}

type PrinterCapabilitiesScalings string

const (
	PrinterCapabilitiesScalingsAuto        PrinterCapabilitiesScalings = "auto"
	PrinterCapabilitiesScalingsFill        PrinterCapabilitiesScalings = "fill"
	PrinterCapabilitiesScalingsFit         PrinterCapabilitiesScalings = "fit"
	PrinterCapabilitiesScalingsNone        PrinterCapabilitiesScalings = "none"
	PrinterCapabilitiesScalingsShrinkToFit PrinterCapabilitiesScalings = "shrinkToFit"
)

func PossibleValuesForPrinterCapabilitiesScalings() []string {
	return []string{
		string(PrinterCapabilitiesScalingsAuto),
		string(PrinterCapabilitiesScalingsFill),
		string(PrinterCapabilitiesScalingsFit),
		string(PrinterCapabilitiesScalingsNone),
		string(PrinterCapabilitiesScalingsShrinkToFit),
	}
}

func (s *PrinterCapabilitiesScalings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesScalings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesScalings(input string) (*PrinterCapabilitiesScalings, error) {
	vals := map[string]PrinterCapabilitiesScalings{
		"auto":        PrinterCapabilitiesScalingsAuto,
		"fill":        PrinterCapabilitiesScalingsFill,
		"fit":         PrinterCapabilitiesScalingsFit,
		"none":        PrinterCapabilitiesScalingsNone,
		"shrinktofit": PrinterCapabilitiesScalingsShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesScalings(input)
	return &out, nil
}

type PrinterDefaultsColorMode string

const (
	PrinterDefaultsColorModeAuto          PrinterDefaultsColorMode = "auto"
	PrinterDefaultsColorModeBlackAndWhite PrinterDefaultsColorMode = "blackAndWhite"
	PrinterDefaultsColorModeColor         PrinterDefaultsColorMode = "color"
	PrinterDefaultsColorModeGrayscale     PrinterDefaultsColorMode = "grayscale"
)

func PossibleValuesForPrinterDefaultsColorMode() []string {
	return []string{
		string(PrinterDefaultsColorModeAuto),
		string(PrinterDefaultsColorModeBlackAndWhite),
		string(PrinterDefaultsColorModeColor),
		string(PrinterDefaultsColorModeGrayscale),
	}
}

func (s *PrinterDefaultsColorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsColorMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsColorMode(input string) (*PrinterDefaultsColorMode, error) {
	vals := map[string]PrinterDefaultsColorMode{
		"auto":          PrinterDefaultsColorModeAuto,
		"blackandwhite": PrinterDefaultsColorModeBlackAndWhite,
		"color":         PrinterDefaultsColorModeColor,
		"grayscale":     PrinterDefaultsColorModeGrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsColorMode(input)
	return &out, nil
}

type PrinterDefaultsDuplexMode string

const (
	PrinterDefaultsDuplexModeFlipOnLongEdge  PrinterDefaultsDuplexMode = "flipOnLongEdge"
	PrinterDefaultsDuplexModeFlipOnShortEdge PrinterDefaultsDuplexMode = "flipOnShortEdge"
	PrinterDefaultsDuplexModeOneSided        PrinterDefaultsDuplexMode = "oneSided"
)

func PossibleValuesForPrinterDefaultsDuplexMode() []string {
	return []string{
		string(PrinterDefaultsDuplexModeFlipOnLongEdge),
		string(PrinterDefaultsDuplexModeFlipOnShortEdge),
		string(PrinterDefaultsDuplexModeOneSided),
	}
}

func (s *PrinterDefaultsDuplexMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsDuplexMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsDuplexMode(input string) (*PrinterDefaultsDuplexMode, error) {
	vals := map[string]PrinterDefaultsDuplexMode{
		"fliponlongedge":  PrinterDefaultsDuplexModeFlipOnLongEdge,
		"fliponshortedge": PrinterDefaultsDuplexModeFlipOnShortEdge,
		"onesided":        PrinterDefaultsDuplexModeOneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsDuplexMode(input)
	return &out, nil
}

type PrinterDefaultsFinishings string

const (
	PrinterDefaultsFinishingsBind              PrinterDefaultsFinishings = "bind"
	PrinterDefaultsFinishingsCover             PrinterDefaultsFinishings = "cover"
	PrinterDefaultsFinishingsNone              PrinterDefaultsFinishings = "none"
	PrinterDefaultsFinishingsPunch             PrinterDefaultsFinishings = "punch"
	PrinterDefaultsFinishingsSaddleStitch      PrinterDefaultsFinishings = "saddleStitch"
	PrinterDefaultsFinishingsStaple            PrinterDefaultsFinishings = "staple"
	PrinterDefaultsFinishingsStapleBottomLeft  PrinterDefaultsFinishings = "stapleBottomLeft"
	PrinterDefaultsFinishingsStapleBottomRight PrinterDefaultsFinishings = "stapleBottomRight"
	PrinterDefaultsFinishingsStapleDualBottom  PrinterDefaultsFinishings = "stapleDualBottom"
	PrinterDefaultsFinishingsStapleDualLeft    PrinterDefaultsFinishings = "stapleDualLeft"
	PrinterDefaultsFinishingsStapleDualRight   PrinterDefaultsFinishings = "stapleDualRight"
	PrinterDefaultsFinishingsStapleDualTop     PrinterDefaultsFinishings = "stapleDualTop"
	PrinterDefaultsFinishingsStapleTopLeft     PrinterDefaultsFinishings = "stapleTopLeft"
	PrinterDefaultsFinishingsStapleTopRight    PrinterDefaultsFinishings = "stapleTopRight"
	PrinterDefaultsFinishingsStitchBottomEdge  PrinterDefaultsFinishings = "stitchBottomEdge"
	PrinterDefaultsFinishingsStitchEdge        PrinterDefaultsFinishings = "stitchEdge"
	PrinterDefaultsFinishingsStitchLeftEdge    PrinterDefaultsFinishings = "stitchLeftEdge"
	PrinterDefaultsFinishingsStitchRightEdge   PrinterDefaultsFinishings = "stitchRightEdge"
	PrinterDefaultsFinishingsStitchTopEdge     PrinterDefaultsFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterDefaultsFinishings() []string {
	return []string{
		string(PrinterDefaultsFinishingsBind),
		string(PrinterDefaultsFinishingsCover),
		string(PrinterDefaultsFinishingsNone),
		string(PrinterDefaultsFinishingsPunch),
		string(PrinterDefaultsFinishingsSaddleStitch),
		string(PrinterDefaultsFinishingsStaple),
		string(PrinterDefaultsFinishingsStapleBottomLeft),
		string(PrinterDefaultsFinishingsStapleBottomRight),
		string(PrinterDefaultsFinishingsStapleDualBottom),
		string(PrinterDefaultsFinishingsStapleDualLeft),
		string(PrinterDefaultsFinishingsStapleDualRight),
		string(PrinterDefaultsFinishingsStapleDualTop),
		string(PrinterDefaultsFinishingsStapleTopLeft),
		string(PrinterDefaultsFinishingsStapleTopRight),
		string(PrinterDefaultsFinishingsStitchBottomEdge),
		string(PrinterDefaultsFinishingsStitchEdge),
		string(PrinterDefaultsFinishingsStitchLeftEdge),
		string(PrinterDefaultsFinishingsStitchRightEdge),
		string(PrinterDefaultsFinishingsStitchTopEdge),
	}
}

func (s *PrinterDefaultsFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsFinishings(input string) (*PrinterDefaultsFinishings, error) {
	vals := map[string]PrinterDefaultsFinishings{
		"bind":              PrinterDefaultsFinishingsBind,
		"cover":             PrinterDefaultsFinishingsCover,
		"none":              PrinterDefaultsFinishingsNone,
		"punch":             PrinterDefaultsFinishingsPunch,
		"saddlestitch":      PrinterDefaultsFinishingsSaddleStitch,
		"staple":            PrinterDefaultsFinishingsStaple,
		"staplebottomleft":  PrinterDefaultsFinishingsStapleBottomLeft,
		"staplebottomright": PrinterDefaultsFinishingsStapleBottomRight,
		"stapledualbottom":  PrinterDefaultsFinishingsStapleDualBottom,
		"stapledualleft":    PrinterDefaultsFinishingsStapleDualLeft,
		"stapledualright":   PrinterDefaultsFinishingsStapleDualRight,
		"stapledualtop":     PrinterDefaultsFinishingsStapleDualTop,
		"stapletopleft":     PrinterDefaultsFinishingsStapleTopLeft,
		"stapletopright":    PrinterDefaultsFinishingsStapleTopRight,
		"stitchbottomedge":  PrinterDefaultsFinishingsStitchBottomEdge,
		"stitchedge":        PrinterDefaultsFinishingsStitchEdge,
		"stitchleftedge":    PrinterDefaultsFinishingsStitchLeftEdge,
		"stitchrightedge":   PrinterDefaultsFinishingsStitchRightEdge,
		"stitchtopedge":     PrinterDefaultsFinishingsStitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsFinishings(input)
	return &out, nil
}

type PrinterDefaultsMultipageLayout string

const (
	PrinterDefaultsMultipageLayoutClockwiseFromBottomLeft         PrinterDefaultsMultipageLayout = "clockwiseFromBottomLeft"
	PrinterDefaultsMultipageLayoutClockwiseFromBottomRight        PrinterDefaultsMultipageLayout = "clockwiseFromBottomRight"
	PrinterDefaultsMultipageLayoutClockwiseFromTopLeft            PrinterDefaultsMultipageLayout = "clockwiseFromTopLeft"
	PrinterDefaultsMultipageLayoutClockwiseFromTopRight           PrinterDefaultsMultipageLayout = "clockwiseFromTopRight"
	PrinterDefaultsMultipageLayoutCounterclockwiseFromBottomLeft  PrinterDefaultsMultipageLayout = "counterclockwiseFromBottomLeft"
	PrinterDefaultsMultipageLayoutCounterclockwiseFromBottomRight PrinterDefaultsMultipageLayout = "counterclockwiseFromBottomRight"
	PrinterDefaultsMultipageLayoutCounterclockwiseFromTopLeft     PrinterDefaultsMultipageLayout = "counterclockwiseFromTopLeft"
	PrinterDefaultsMultipageLayoutCounterclockwiseFromTopRight    PrinterDefaultsMultipageLayout = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrinterDefaultsMultipageLayout() []string {
	return []string{
		string(PrinterDefaultsMultipageLayoutClockwiseFromBottomLeft),
		string(PrinterDefaultsMultipageLayoutClockwiseFromBottomRight),
		string(PrinterDefaultsMultipageLayoutClockwiseFromTopLeft),
		string(PrinterDefaultsMultipageLayoutClockwiseFromTopRight),
		string(PrinterDefaultsMultipageLayoutCounterclockwiseFromBottomLeft),
		string(PrinterDefaultsMultipageLayoutCounterclockwiseFromBottomRight),
		string(PrinterDefaultsMultipageLayoutCounterclockwiseFromTopLeft),
		string(PrinterDefaultsMultipageLayoutCounterclockwiseFromTopRight),
	}
}

func (s *PrinterDefaultsMultipageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsMultipageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsMultipageLayout(input string) (*PrinterDefaultsMultipageLayout, error) {
	vals := map[string]PrinterDefaultsMultipageLayout{
		"clockwisefrombottomleft":         PrinterDefaultsMultipageLayoutClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDefaultsMultipageLayoutClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDefaultsMultipageLayoutClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDefaultsMultipageLayoutClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDefaultsMultipageLayoutCounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDefaultsMultipageLayoutCounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDefaultsMultipageLayoutCounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDefaultsMultipageLayoutCounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsMultipageLayout(input)
	return &out, nil
}

type PrinterDefaultsOrientation string

const (
	PrinterDefaultsOrientationLandscape        PrinterDefaultsOrientation = "landscape"
	PrinterDefaultsOrientationPortrait         PrinterDefaultsOrientation = "portrait"
	PrinterDefaultsOrientationReverseLandscape PrinterDefaultsOrientation = "reverseLandscape"
	PrinterDefaultsOrientationReversePortrait  PrinterDefaultsOrientation = "reversePortrait"
)

func PossibleValuesForPrinterDefaultsOrientation() []string {
	return []string{
		string(PrinterDefaultsOrientationLandscape),
		string(PrinterDefaultsOrientationPortrait),
		string(PrinterDefaultsOrientationReverseLandscape),
		string(PrinterDefaultsOrientationReversePortrait),
	}
}

func (s *PrinterDefaultsOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsOrientation(input string) (*PrinterDefaultsOrientation, error) {
	vals := map[string]PrinterDefaultsOrientation{
		"landscape":        PrinterDefaultsOrientationLandscape,
		"portrait":         PrinterDefaultsOrientationPortrait,
		"reverselandscape": PrinterDefaultsOrientationReverseLandscape,
		"reverseportrait":  PrinterDefaultsOrientationReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsOrientation(input)
	return &out, nil
}

type PrinterDefaultsQuality string

const (
	PrinterDefaultsQualityHigh   PrinterDefaultsQuality = "high"
	PrinterDefaultsQualityLow    PrinterDefaultsQuality = "low"
	PrinterDefaultsQualityMedium PrinterDefaultsQuality = "medium"
)

func PossibleValuesForPrinterDefaultsQuality() []string {
	return []string{
		string(PrinterDefaultsQualityHigh),
		string(PrinterDefaultsQualityLow),
		string(PrinterDefaultsQualityMedium),
	}
}

func (s *PrinterDefaultsQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsQuality(input string) (*PrinterDefaultsQuality, error) {
	vals := map[string]PrinterDefaultsQuality{
		"high":   PrinterDefaultsQualityHigh,
		"low":    PrinterDefaultsQualityLow,
		"medium": PrinterDefaultsQualityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsQuality(input)
	return &out, nil
}

type PrinterDefaultsScaling string

const (
	PrinterDefaultsScalingAuto        PrinterDefaultsScaling = "auto"
	PrinterDefaultsScalingFill        PrinterDefaultsScaling = "fill"
	PrinterDefaultsScalingFit         PrinterDefaultsScaling = "fit"
	PrinterDefaultsScalingNone        PrinterDefaultsScaling = "none"
	PrinterDefaultsScalingShrinkToFit PrinterDefaultsScaling = "shrinkToFit"
)

func PossibleValuesForPrinterDefaultsScaling() []string {
	return []string{
		string(PrinterDefaultsScalingAuto),
		string(PrinterDefaultsScalingFill),
		string(PrinterDefaultsScalingFit),
		string(PrinterDefaultsScalingNone),
		string(PrinterDefaultsScalingShrinkToFit),
	}
}

func (s *PrinterDefaultsScaling) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsScaling(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsScaling(input string) (*PrinterDefaultsScaling, error) {
	vals := map[string]PrinterDefaultsScaling{
		"auto":        PrinterDefaultsScalingAuto,
		"fill":        PrinterDefaultsScalingFill,
		"fit":         PrinterDefaultsScalingFit,
		"none":        PrinterDefaultsScalingNone,
		"shrinktofit": PrinterDefaultsScalingShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsScaling(input)
	return &out, nil
}

type PrinterStatusDetails string

const (
	PrinterStatusDetailsAlertRemovalOfBinaryChangeEntry           PrinterStatusDetails = "alertRemovalOfBinaryChangeEntry"
	PrinterStatusDetailsBanderAdded                               PrinterStatusDetails = "banderAdded"
	PrinterStatusDetailsBanderAlmostEmpty                         PrinterStatusDetails = "banderAlmostEmpty"
	PrinterStatusDetailsBanderAlmostFull                          PrinterStatusDetails = "banderAlmostFull"
	PrinterStatusDetailsBanderAtLimit                             PrinterStatusDetails = "banderAtLimit"
	PrinterStatusDetailsBanderClosed                              PrinterStatusDetails = "banderClosed"
	PrinterStatusDetailsBanderConfigurationChange                 PrinterStatusDetails = "banderConfigurationChange"
	PrinterStatusDetailsBanderCoverClosed                         PrinterStatusDetails = "banderCoverClosed"
	PrinterStatusDetailsBanderCoverOpen                           PrinterStatusDetails = "banderCoverOpen"
	PrinterStatusDetailsBanderEmpty                               PrinterStatusDetails = "banderEmpty"
	PrinterStatusDetailsBanderFull                                PrinterStatusDetails = "banderFull"
	PrinterStatusDetailsBanderInterlockClosed                     PrinterStatusDetails = "banderInterlockClosed"
	PrinterStatusDetailsBanderInterlockOpen                       PrinterStatusDetails = "banderInterlockOpen"
	PrinterStatusDetailsBanderJam                                 PrinterStatusDetails = "banderJam"
	PrinterStatusDetailsBanderLifeAlmostOver                      PrinterStatusDetails = "banderLifeAlmostOver"
	PrinterStatusDetailsBanderLifeOver                            PrinterStatusDetails = "banderLifeOver"
	PrinterStatusDetailsBanderMemoryExhausted                     PrinterStatusDetails = "banderMemoryExhausted"
	PrinterStatusDetailsBanderMissing                             PrinterStatusDetails = "banderMissing"
	PrinterStatusDetailsBanderMotorFailure                        PrinterStatusDetails = "banderMotorFailure"
	PrinterStatusDetailsBanderNearLimit                           PrinterStatusDetails = "banderNearLimit"
	PrinterStatusDetailsBanderOffline                             PrinterStatusDetails = "banderOffline"
	PrinterStatusDetailsBanderOpened                              PrinterStatusDetails = "banderOpened"
	PrinterStatusDetailsBanderOverTemperature                     PrinterStatusDetails = "banderOverTemperature"
	PrinterStatusDetailsBanderPowerSaver                          PrinterStatusDetails = "banderPowerSaver"
	PrinterStatusDetailsBanderRecoverableFailure                  PrinterStatusDetails = "banderRecoverableFailure"
	PrinterStatusDetailsBanderRecoverableStorage                  PrinterStatusDetails = "banderRecoverableStorage"
	PrinterStatusDetailsBanderRemoved                             PrinterStatusDetails = "banderRemoved"
	PrinterStatusDetailsBanderResourceAdded                       PrinterStatusDetails = "banderResourceAdded"
	PrinterStatusDetailsBanderResourceRemoved                     PrinterStatusDetails = "banderResourceRemoved"
	PrinterStatusDetailsBanderThermistorFailure                   PrinterStatusDetails = "banderThermistorFailure"
	PrinterStatusDetailsBanderTimingFailure                       PrinterStatusDetails = "banderTimingFailure"
	PrinterStatusDetailsBanderTurnedOff                           PrinterStatusDetails = "banderTurnedOff"
	PrinterStatusDetailsBanderTurnedOn                            PrinterStatusDetails = "banderTurnedOn"
	PrinterStatusDetailsBanderUnderTemperature                    PrinterStatusDetails = "banderUnderTemperature"
	PrinterStatusDetailsBanderUnrecoverableFailure                PrinterStatusDetails = "banderUnrecoverableFailure"
	PrinterStatusDetailsBanderUnrecoverableStorageError           PrinterStatusDetails = "banderUnrecoverableStorageError"
	PrinterStatusDetailsBanderWarmingUp                           PrinterStatusDetails = "banderWarmingUp"
	PrinterStatusDetailsBinderAdded                               PrinterStatusDetails = "binderAdded"
	PrinterStatusDetailsBinderAlmostEmpty                         PrinterStatusDetails = "binderAlmostEmpty"
	PrinterStatusDetailsBinderAlmostFull                          PrinterStatusDetails = "binderAlmostFull"
	PrinterStatusDetailsBinderAtLimit                             PrinterStatusDetails = "binderAtLimit"
	PrinterStatusDetailsBinderClosed                              PrinterStatusDetails = "binderClosed"
	PrinterStatusDetailsBinderConfigurationChange                 PrinterStatusDetails = "binderConfigurationChange"
	PrinterStatusDetailsBinderCoverClosed                         PrinterStatusDetails = "binderCoverClosed"
	PrinterStatusDetailsBinderCoverOpen                           PrinterStatusDetails = "binderCoverOpen"
	PrinterStatusDetailsBinderEmpty                               PrinterStatusDetails = "binderEmpty"
	PrinterStatusDetailsBinderFull                                PrinterStatusDetails = "binderFull"
	PrinterStatusDetailsBinderInterlockClosed                     PrinterStatusDetails = "binderInterlockClosed"
	PrinterStatusDetailsBinderInterlockOpen                       PrinterStatusDetails = "binderInterlockOpen"
	PrinterStatusDetailsBinderJam                                 PrinterStatusDetails = "binderJam"
	PrinterStatusDetailsBinderLifeAlmostOver                      PrinterStatusDetails = "binderLifeAlmostOver"
	PrinterStatusDetailsBinderLifeOver                            PrinterStatusDetails = "binderLifeOver"
	PrinterStatusDetailsBinderMemoryExhausted                     PrinterStatusDetails = "binderMemoryExhausted"
	PrinterStatusDetailsBinderMissing                             PrinterStatusDetails = "binderMissing"
	PrinterStatusDetailsBinderMotorFailure                        PrinterStatusDetails = "binderMotorFailure"
	PrinterStatusDetailsBinderNearLimit                           PrinterStatusDetails = "binderNearLimit"
	PrinterStatusDetailsBinderOffline                             PrinterStatusDetails = "binderOffline"
	PrinterStatusDetailsBinderOpened                              PrinterStatusDetails = "binderOpened"
	PrinterStatusDetailsBinderOverTemperature                     PrinterStatusDetails = "binderOverTemperature"
	PrinterStatusDetailsBinderPowerSaver                          PrinterStatusDetails = "binderPowerSaver"
	PrinterStatusDetailsBinderRecoverableFailure                  PrinterStatusDetails = "binderRecoverableFailure"
	PrinterStatusDetailsBinderRecoverableStorage                  PrinterStatusDetails = "binderRecoverableStorage"
	PrinterStatusDetailsBinderRemoved                             PrinterStatusDetails = "binderRemoved"
	PrinterStatusDetailsBinderResourceAdded                       PrinterStatusDetails = "binderResourceAdded"
	PrinterStatusDetailsBinderResourceRemoved                     PrinterStatusDetails = "binderResourceRemoved"
	PrinterStatusDetailsBinderThermistorFailure                   PrinterStatusDetails = "binderThermistorFailure"
	PrinterStatusDetailsBinderTimingFailure                       PrinterStatusDetails = "binderTimingFailure"
	PrinterStatusDetailsBinderTurnedOff                           PrinterStatusDetails = "binderTurnedOff"
	PrinterStatusDetailsBinderTurnedOn                            PrinterStatusDetails = "binderTurnedOn"
	PrinterStatusDetailsBinderUnderTemperature                    PrinterStatusDetails = "binderUnderTemperature"
	PrinterStatusDetailsBinderUnrecoverableFailure                PrinterStatusDetails = "binderUnrecoverableFailure"
	PrinterStatusDetailsBinderUnrecoverableStorageError           PrinterStatusDetails = "binderUnrecoverableStorageError"
	PrinterStatusDetailsBinderWarmingUp                           PrinterStatusDetails = "binderWarmingUp"
	PrinterStatusDetailsCameraFailure                             PrinterStatusDetails = "cameraFailure"
	PrinterStatusDetailsChamberCooling                            PrinterStatusDetails = "chamberCooling"
	PrinterStatusDetailsChamberFailure                            PrinterStatusDetails = "chamberFailure"
	PrinterStatusDetailsChamberHeating                            PrinterStatusDetails = "chamberHeating"
	PrinterStatusDetailsChamberTemperatureHigh                    PrinterStatusDetails = "chamberTemperatureHigh"
	PrinterStatusDetailsChamberTemperatureLow                     PrinterStatusDetails = "chamberTemperatureLow"
	PrinterStatusDetailsCleanerLifeAlmostOver                     PrinterStatusDetails = "cleanerLifeAlmostOver"
	PrinterStatusDetailsCleanerLifeOver                           PrinterStatusDetails = "cleanerLifeOver"
	PrinterStatusDetailsConfigurationChange                       PrinterStatusDetails = "configurationChange"
	PrinterStatusDetailsConnectingToDevice                        PrinterStatusDetails = "connectingToDevice"
	PrinterStatusDetailsCoverOpen                                 PrinterStatusDetails = "coverOpen"
	PrinterStatusDetailsDeactivated                               PrinterStatusDetails = "deactivated"
	PrinterStatusDetailsDeleted                                   PrinterStatusDetails = "deleted"
	PrinterStatusDetailsDeveloperEmpty                            PrinterStatusDetails = "developerEmpty"
	PrinterStatusDetailsDeveloperLow                              PrinterStatusDetails = "developerLow"
	PrinterStatusDetailsDieCutterAdded                            PrinterStatusDetails = "dieCutterAdded"
	PrinterStatusDetailsDieCutterAlmostEmpty                      PrinterStatusDetails = "dieCutterAlmostEmpty"
	PrinterStatusDetailsDieCutterAlmostFull                       PrinterStatusDetails = "dieCutterAlmostFull"
	PrinterStatusDetailsDieCutterAtLimit                          PrinterStatusDetails = "dieCutterAtLimit"
	PrinterStatusDetailsDieCutterClosed                           PrinterStatusDetails = "dieCutterClosed"
	PrinterStatusDetailsDieCutterConfigurationChange              PrinterStatusDetails = "dieCutterConfigurationChange"
	PrinterStatusDetailsDieCutterCoverClosed                      PrinterStatusDetails = "dieCutterCoverClosed"
	PrinterStatusDetailsDieCutterCoverOpen                        PrinterStatusDetails = "dieCutterCoverOpen"
	PrinterStatusDetailsDieCutterEmpty                            PrinterStatusDetails = "dieCutterEmpty"
	PrinterStatusDetailsDieCutterFull                             PrinterStatusDetails = "dieCutterFull"
	PrinterStatusDetailsDieCutterInterlockClosed                  PrinterStatusDetails = "dieCutterInterlockClosed"
	PrinterStatusDetailsDieCutterInterlockOpen                    PrinterStatusDetails = "dieCutterInterlockOpen"
	PrinterStatusDetailsDieCutterJam                              PrinterStatusDetails = "dieCutterJam"
	PrinterStatusDetailsDieCutterLifeAlmostOver                   PrinterStatusDetails = "dieCutterLifeAlmostOver"
	PrinterStatusDetailsDieCutterLifeOver                         PrinterStatusDetails = "dieCutterLifeOver"
	PrinterStatusDetailsDieCutterMemoryExhausted                  PrinterStatusDetails = "dieCutterMemoryExhausted"
	PrinterStatusDetailsDieCutterMissing                          PrinterStatusDetails = "dieCutterMissing"
	PrinterStatusDetailsDieCutterMotorFailure                     PrinterStatusDetails = "dieCutterMotorFailure"
	PrinterStatusDetailsDieCutterNearLimit                        PrinterStatusDetails = "dieCutterNearLimit"
	PrinterStatusDetailsDieCutterOffline                          PrinterStatusDetails = "dieCutterOffline"
	PrinterStatusDetailsDieCutterOpened                           PrinterStatusDetails = "dieCutterOpened"
	PrinterStatusDetailsDieCutterOverTemperature                  PrinterStatusDetails = "dieCutterOverTemperature"
	PrinterStatusDetailsDieCutterPowerSaver                       PrinterStatusDetails = "dieCutterPowerSaver"
	PrinterStatusDetailsDieCutterRecoverableFailure               PrinterStatusDetails = "dieCutterRecoverableFailure"
	PrinterStatusDetailsDieCutterRecoverableStorage               PrinterStatusDetails = "dieCutterRecoverableStorage"
	PrinterStatusDetailsDieCutterRemoved                          PrinterStatusDetails = "dieCutterRemoved"
	PrinterStatusDetailsDieCutterResourceAdded                    PrinterStatusDetails = "dieCutterResourceAdded"
	PrinterStatusDetailsDieCutterResourceRemoved                  PrinterStatusDetails = "dieCutterResourceRemoved"
	PrinterStatusDetailsDieCutterThermistorFailure                PrinterStatusDetails = "dieCutterThermistorFailure"
	PrinterStatusDetailsDieCutterTimingFailure                    PrinterStatusDetails = "dieCutterTimingFailure"
	PrinterStatusDetailsDieCutterTurnedOff                        PrinterStatusDetails = "dieCutterTurnedOff"
	PrinterStatusDetailsDieCutterTurnedOn                         PrinterStatusDetails = "dieCutterTurnedOn"
	PrinterStatusDetailsDieCutterUnderTemperature                 PrinterStatusDetails = "dieCutterUnderTemperature"
	PrinterStatusDetailsDieCutterUnrecoverableFailure             PrinterStatusDetails = "dieCutterUnrecoverableFailure"
	PrinterStatusDetailsDieCutterUnrecoverableStorageError        PrinterStatusDetails = "dieCutterUnrecoverableStorageError"
	PrinterStatusDetailsDieCutterWarmingUp                        PrinterStatusDetails = "dieCutterWarmingUp"
	PrinterStatusDetailsDoorOpen                                  PrinterStatusDetails = "doorOpen"
	PrinterStatusDetailsExtruderCooling                           PrinterStatusDetails = "extruderCooling"
	PrinterStatusDetailsExtruderFailure                           PrinterStatusDetails = "extruderFailure"
	PrinterStatusDetailsExtruderHeating                           PrinterStatusDetails = "extruderHeating"
	PrinterStatusDetailsExtruderJam                               PrinterStatusDetails = "extruderJam"
	PrinterStatusDetailsExtruderTemperatureHigh                   PrinterStatusDetails = "extruderTemperatureHigh"
	PrinterStatusDetailsExtruderTemperatureLow                    PrinterStatusDetails = "extruderTemperatureLow"
	PrinterStatusDetailsFanFailure                                PrinterStatusDetails = "fanFailure"
	PrinterStatusDetailsFaxModemLifeAlmostOver                    PrinterStatusDetails = "faxModemLifeAlmostOver"
	PrinterStatusDetailsFaxModemLifeOver                          PrinterStatusDetails = "faxModemLifeOver"
	PrinterStatusDetailsFaxModemMissing                           PrinterStatusDetails = "faxModemMissing"
	PrinterStatusDetailsFaxModemTurnedOff                         PrinterStatusDetails = "faxModemTurnedOff"
	PrinterStatusDetailsFaxModemTurnedOn                          PrinterStatusDetails = "faxModemTurnedOn"
	PrinterStatusDetailsFolderAdded                               PrinterStatusDetails = "folderAdded"
	PrinterStatusDetailsFolderAlmostEmpty                         PrinterStatusDetails = "folderAlmostEmpty"
	PrinterStatusDetailsFolderAlmostFull                          PrinterStatusDetails = "folderAlmostFull"
	PrinterStatusDetailsFolderAtLimit                             PrinterStatusDetails = "folderAtLimit"
	PrinterStatusDetailsFolderClosed                              PrinterStatusDetails = "folderClosed"
	PrinterStatusDetailsFolderConfigurationChange                 PrinterStatusDetails = "folderConfigurationChange"
	PrinterStatusDetailsFolderCoverClosed                         PrinterStatusDetails = "folderCoverClosed"
	PrinterStatusDetailsFolderCoverOpen                           PrinterStatusDetails = "folderCoverOpen"
	PrinterStatusDetailsFolderEmpty                               PrinterStatusDetails = "folderEmpty"
	PrinterStatusDetailsFolderFull                                PrinterStatusDetails = "folderFull"
	PrinterStatusDetailsFolderInterlockClosed                     PrinterStatusDetails = "folderInterlockClosed"
	PrinterStatusDetailsFolderInterlockOpen                       PrinterStatusDetails = "folderInterlockOpen"
	PrinterStatusDetailsFolderJam                                 PrinterStatusDetails = "folderJam"
	PrinterStatusDetailsFolderLifeAlmostOver                      PrinterStatusDetails = "folderLifeAlmostOver"
	PrinterStatusDetailsFolderLifeOver                            PrinterStatusDetails = "folderLifeOver"
	PrinterStatusDetailsFolderMemoryExhausted                     PrinterStatusDetails = "folderMemoryExhausted"
	PrinterStatusDetailsFolderMissing                             PrinterStatusDetails = "folderMissing"
	PrinterStatusDetailsFolderMotorFailure                        PrinterStatusDetails = "folderMotorFailure"
	PrinterStatusDetailsFolderNearLimit                           PrinterStatusDetails = "folderNearLimit"
	PrinterStatusDetailsFolderOffline                             PrinterStatusDetails = "folderOffline"
	PrinterStatusDetailsFolderOpened                              PrinterStatusDetails = "folderOpened"
	PrinterStatusDetailsFolderOverTemperature                     PrinterStatusDetails = "folderOverTemperature"
	PrinterStatusDetailsFolderPowerSaver                          PrinterStatusDetails = "folderPowerSaver"
	PrinterStatusDetailsFolderRecoverableFailure                  PrinterStatusDetails = "folderRecoverableFailure"
	PrinterStatusDetailsFolderRecoverableStorage                  PrinterStatusDetails = "folderRecoverableStorage"
	PrinterStatusDetailsFolderRemoved                             PrinterStatusDetails = "folderRemoved"
	PrinterStatusDetailsFolderResourceAdded                       PrinterStatusDetails = "folderResourceAdded"
	PrinterStatusDetailsFolderResourceRemoved                     PrinterStatusDetails = "folderResourceRemoved"
	PrinterStatusDetailsFolderThermistorFailure                   PrinterStatusDetails = "folderThermistorFailure"
	PrinterStatusDetailsFolderTimingFailure                       PrinterStatusDetails = "folderTimingFailure"
	PrinterStatusDetailsFolderTurnedOff                           PrinterStatusDetails = "folderTurnedOff"
	PrinterStatusDetailsFolderTurnedOn                            PrinterStatusDetails = "folderTurnedOn"
	PrinterStatusDetailsFolderUnderTemperature                    PrinterStatusDetails = "folderUnderTemperature"
	PrinterStatusDetailsFolderUnrecoverableFailure                PrinterStatusDetails = "folderUnrecoverableFailure"
	PrinterStatusDetailsFolderUnrecoverableStorageError           PrinterStatusDetails = "folderUnrecoverableStorageError"
	PrinterStatusDetailsFolderWarmingUp                           PrinterStatusDetails = "folderWarmingUp"
	PrinterStatusDetailsFuserOverTemp                             PrinterStatusDetails = "fuserOverTemp"
	PrinterStatusDetailsFuserUnderTemp                            PrinterStatusDetails = "fuserUnderTemp"
	PrinterStatusDetailsHibernate                                 PrinterStatusDetails = "hibernate"
	PrinterStatusDetailsHoldNewJobs                               PrinterStatusDetails = "holdNewJobs"
	PrinterStatusDetailsIdentifyPrinterRequested                  PrinterStatusDetails = "identifyPrinterRequested"
	PrinterStatusDetailsImprinterAdded                            PrinterStatusDetails = "imprinterAdded"
	PrinterStatusDetailsImprinterAlmostEmpty                      PrinterStatusDetails = "imprinterAlmostEmpty"
	PrinterStatusDetailsImprinterAlmostFull                       PrinterStatusDetails = "imprinterAlmostFull"
	PrinterStatusDetailsImprinterAtLimit                          PrinterStatusDetails = "imprinterAtLimit"
	PrinterStatusDetailsImprinterClosed                           PrinterStatusDetails = "imprinterClosed"
	PrinterStatusDetailsImprinterConfigurationChange              PrinterStatusDetails = "imprinterConfigurationChange"
	PrinterStatusDetailsImprinterCoverClosed                      PrinterStatusDetails = "imprinterCoverClosed"
	PrinterStatusDetailsImprinterCoverOpen                        PrinterStatusDetails = "imprinterCoverOpen"
	PrinterStatusDetailsImprinterEmpty                            PrinterStatusDetails = "imprinterEmpty"
	PrinterStatusDetailsImprinterFull                             PrinterStatusDetails = "imprinterFull"
	PrinterStatusDetailsImprinterInterlockClosed                  PrinterStatusDetails = "imprinterInterlockClosed"
	PrinterStatusDetailsImprinterInterlockOpen                    PrinterStatusDetails = "imprinterInterlockOpen"
	PrinterStatusDetailsImprinterJam                              PrinterStatusDetails = "imprinterJam"
	PrinterStatusDetailsImprinterLifeAlmostOver                   PrinterStatusDetails = "imprinterLifeAlmostOver"
	PrinterStatusDetailsImprinterLifeOver                         PrinterStatusDetails = "imprinterLifeOver"
	PrinterStatusDetailsImprinterMemoryExhausted                  PrinterStatusDetails = "imprinterMemoryExhausted"
	PrinterStatusDetailsImprinterMissing                          PrinterStatusDetails = "imprinterMissing"
	PrinterStatusDetailsImprinterMotorFailure                     PrinterStatusDetails = "imprinterMotorFailure"
	PrinterStatusDetailsImprinterNearLimit                        PrinterStatusDetails = "imprinterNearLimit"
	PrinterStatusDetailsImprinterOffline                          PrinterStatusDetails = "imprinterOffline"
	PrinterStatusDetailsImprinterOpened                           PrinterStatusDetails = "imprinterOpened"
	PrinterStatusDetailsImprinterOverTemperature                  PrinterStatusDetails = "imprinterOverTemperature"
	PrinterStatusDetailsImprinterPowerSaver                       PrinterStatusDetails = "imprinterPowerSaver"
	PrinterStatusDetailsImprinterRecoverableFailure               PrinterStatusDetails = "imprinterRecoverableFailure"
	PrinterStatusDetailsImprinterRecoverableStorage               PrinterStatusDetails = "imprinterRecoverableStorage"
	PrinterStatusDetailsImprinterRemoved                          PrinterStatusDetails = "imprinterRemoved"
	PrinterStatusDetailsImprinterResourceAdded                    PrinterStatusDetails = "imprinterResourceAdded"
	PrinterStatusDetailsImprinterResourceRemoved                  PrinterStatusDetails = "imprinterResourceRemoved"
	PrinterStatusDetailsImprinterThermistorFailure                PrinterStatusDetails = "imprinterThermistorFailure"
	PrinterStatusDetailsImprinterTimingFailure                    PrinterStatusDetails = "imprinterTimingFailure"
	PrinterStatusDetailsImprinterTurnedOff                        PrinterStatusDetails = "imprinterTurnedOff"
	PrinterStatusDetailsImprinterTurnedOn                         PrinterStatusDetails = "imprinterTurnedOn"
	PrinterStatusDetailsImprinterUnderTemperature                 PrinterStatusDetails = "imprinterUnderTemperature"
	PrinterStatusDetailsImprinterUnrecoverableFailure             PrinterStatusDetails = "imprinterUnrecoverableFailure"
	PrinterStatusDetailsImprinterUnrecoverableStorageError        PrinterStatusDetails = "imprinterUnrecoverableStorageError"
	PrinterStatusDetailsImprinterWarmingUp                        PrinterStatusDetails = "imprinterWarmingUp"
	PrinterStatusDetailsInputCannotFeedSizeSelected               PrinterStatusDetails = "inputCannotFeedSizeSelected"
	PrinterStatusDetailsInputManualInputRequest                   PrinterStatusDetails = "inputManualInputRequest"
	PrinterStatusDetailsInputMediaColorChange                     PrinterStatusDetails = "inputMediaColorChange"
	PrinterStatusDetailsInputMediaFormPartsChange                 PrinterStatusDetails = "inputMediaFormPartsChange"
	PrinterStatusDetailsInputMediaSizeChange                      PrinterStatusDetails = "inputMediaSizeChange"
	PrinterStatusDetailsInputMediaTrayFailure                     PrinterStatusDetails = "inputMediaTrayFailure"
	PrinterStatusDetailsInputMediaTrayFeedError                   PrinterStatusDetails = "inputMediaTrayFeedError"
	PrinterStatusDetailsInputMediaTrayJam                         PrinterStatusDetails = "inputMediaTrayJam"
	PrinterStatusDetailsInputMediaTypeChange                      PrinterStatusDetails = "inputMediaTypeChange"
	PrinterStatusDetailsInputMediaWeightChange                    PrinterStatusDetails = "inputMediaWeightChange"
	PrinterStatusDetailsInputPickRollerFailure                    PrinterStatusDetails = "inputPickRollerFailure"
	PrinterStatusDetailsInputPickRollerLifeOver                   PrinterStatusDetails = "inputPickRollerLifeOver"
	PrinterStatusDetailsInputPickRollerLifeWarn                   PrinterStatusDetails = "inputPickRollerLifeWarn"
	PrinterStatusDetailsInputPickRollerMissing                    PrinterStatusDetails = "inputPickRollerMissing"
	PrinterStatusDetailsInputTrayElevationFailure                 PrinterStatusDetails = "inputTrayElevationFailure"
	PrinterStatusDetailsInputTrayMissing                          PrinterStatusDetails = "inputTrayMissing"
	PrinterStatusDetailsInputTrayPositionFailure                  PrinterStatusDetails = "inputTrayPositionFailure"
	PrinterStatusDetailsInserterAdded                             PrinterStatusDetails = "inserterAdded"
	PrinterStatusDetailsInserterAlmostEmpty                       PrinterStatusDetails = "inserterAlmostEmpty"
	PrinterStatusDetailsInserterAlmostFull                        PrinterStatusDetails = "inserterAlmostFull"
	PrinterStatusDetailsInserterAtLimit                           PrinterStatusDetails = "inserterAtLimit"
	PrinterStatusDetailsInserterClosed                            PrinterStatusDetails = "inserterClosed"
	PrinterStatusDetailsInserterConfigurationChange               PrinterStatusDetails = "inserterConfigurationChange"
	PrinterStatusDetailsInserterCoverClosed                       PrinterStatusDetails = "inserterCoverClosed"
	PrinterStatusDetailsInserterCoverOpen                         PrinterStatusDetails = "inserterCoverOpen"
	PrinterStatusDetailsInserterEmpty                             PrinterStatusDetails = "inserterEmpty"
	PrinterStatusDetailsInserterFull                              PrinterStatusDetails = "inserterFull"
	PrinterStatusDetailsInserterInterlockClosed                   PrinterStatusDetails = "inserterInterlockClosed"
	PrinterStatusDetailsInserterInterlockOpen                     PrinterStatusDetails = "inserterInterlockOpen"
	PrinterStatusDetailsInserterJam                               PrinterStatusDetails = "inserterJam"
	PrinterStatusDetailsInserterLifeAlmostOver                    PrinterStatusDetails = "inserterLifeAlmostOver"
	PrinterStatusDetailsInserterLifeOver                          PrinterStatusDetails = "inserterLifeOver"
	PrinterStatusDetailsInserterMemoryExhausted                   PrinterStatusDetails = "inserterMemoryExhausted"
	PrinterStatusDetailsInserterMissing                           PrinterStatusDetails = "inserterMissing"
	PrinterStatusDetailsInserterMotorFailure                      PrinterStatusDetails = "inserterMotorFailure"
	PrinterStatusDetailsInserterNearLimit                         PrinterStatusDetails = "inserterNearLimit"
	PrinterStatusDetailsInserterOffline                           PrinterStatusDetails = "inserterOffline"
	PrinterStatusDetailsInserterOpened                            PrinterStatusDetails = "inserterOpened"
	PrinterStatusDetailsInserterOverTemperature                   PrinterStatusDetails = "inserterOverTemperature"
	PrinterStatusDetailsInserterPowerSaver                        PrinterStatusDetails = "inserterPowerSaver"
	PrinterStatusDetailsInserterRecoverableFailure                PrinterStatusDetails = "inserterRecoverableFailure"
	PrinterStatusDetailsInserterRecoverableStorage                PrinterStatusDetails = "inserterRecoverableStorage"
	PrinterStatusDetailsInserterRemoved                           PrinterStatusDetails = "inserterRemoved"
	PrinterStatusDetailsInserterResourceAdded                     PrinterStatusDetails = "inserterResourceAdded"
	PrinterStatusDetailsInserterResourceRemoved                   PrinterStatusDetails = "inserterResourceRemoved"
	PrinterStatusDetailsInserterThermistorFailure                 PrinterStatusDetails = "inserterThermistorFailure"
	PrinterStatusDetailsInserterTimingFailure                     PrinterStatusDetails = "inserterTimingFailure"
	PrinterStatusDetailsInserterTurnedOff                         PrinterStatusDetails = "inserterTurnedOff"
	PrinterStatusDetailsInserterTurnedOn                          PrinterStatusDetails = "inserterTurnedOn"
	PrinterStatusDetailsInserterUnderTemperature                  PrinterStatusDetails = "inserterUnderTemperature"
	PrinterStatusDetailsInserterUnrecoverableFailure              PrinterStatusDetails = "inserterUnrecoverableFailure"
	PrinterStatusDetailsInserterUnrecoverableStorageError         PrinterStatusDetails = "inserterUnrecoverableStorageError"
	PrinterStatusDetailsInserterWarmingUp                         PrinterStatusDetails = "inserterWarmingUp"
	PrinterStatusDetailsInterlockClosed                           PrinterStatusDetails = "interlockClosed"
	PrinterStatusDetailsInterlockOpen                             PrinterStatusDetails = "interlockOpen"
	PrinterStatusDetailsInterpreterCartridgeAdded                 PrinterStatusDetails = "interpreterCartridgeAdded"
	PrinterStatusDetailsInterpreterCartridgeDeleted               PrinterStatusDetails = "interpreterCartridgeDeleted"
	PrinterStatusDetailsInterpreterComplexPageEncountered         PrinterStatusDetails = "interpreterComplexPageEncountered"
	PrinterStatusDetailsInterpreterMemoryDecrease                 PrinterStatusDetails = "interpreterMemoryDecrease"
	PrinterStatusDetailsInterpreterMemoryIncrease                 PrinterStatusDetails = "interpreterMemoryIncrease"
	PrinterStatusDetailsInterpreterResourceAdded                  PrinterStatusDetails = "interpreterResourceAdded"
	PrinterStatusDetailsInterpreterResourceDeleted                PrinterStatusDetails = "interpreterResourceDeleted"
	PrinterStatusDetailsInterpreterResourceUnavailable            PrinterStatusDetails = "interpreterResourceUnavailable"
	PrinterStatusDetailsLampAtEol                                 PrinterStatusDetails = "lampAtEol"
	PrinterStatusDetailsLampFailure                               PrinterStatusDetails = "lampFailure"
	PrinterStatusDetailsLampNearEol                               PrinterStatusDetails = "lampNearEol"
	PrinterStatusDetailsLaserAtEol                                PrinterStatusDetails = "laserAtEol"
	PrinterStatusDetailsLaserFailure                              PrinterStatusDetails = "laserFailure"
	PrinterStatusDetailsLaserNearEol                              PrinterStatusDetails = "laserNearEol"
	PrinterStatusDetailsMakeEnvelopeAdded                         PrinterStatusDetails = "makeEnvelopeAdded"
	PrinterStatusDetailsMakeEnvelopeAlmostEmpty                   PrinterStatusDetails = "makeEnvelopeAlmostEmpty"
	PrinterStatusDetailsMakeEnvelopeAlmostFull                    PrinterStatusDetails = "makeEnvelopeAlmostFull"
	PrinterStatusDetailsMakeEnvelopeAtLimit                       PrinterStatusDetails = "makeEnvelopeAtLimit"
	PrinterStatusDetailsMakeEnvelopeClosed                        PrinterStatusDetails = "makeEnvelopeClosed"
	PrinterStatusDetailsMakeEnvelopeConfigurationChange           PrinterStatusDetails = "makeEnvelopeConfigurationChange"
	PrinterStatusDetailsMakeEnvelopeCoverClosed                   PrinterStatusDetails = "makeEnvelopeCoverClosed"
	PrinterStatusDetailsMakeEnvelopeCoverOpen                     PrinterStatusDetails = "makeEnvelopeCoverOpen"
	PrinterStatusDetailsMakeEnvelopeEmpty                         PrinterStatusDetails = "makeEnvelopeEmpty"
	PrinterStatusDetailsMakeEnvelopeFull                          PrinterStatusDetails = "makeEnvelopeFull"
	PrinterStatusDetailsMakeEnvelopeInterlockClosed               PrinterStatusDetails = "makeEnvelopeInterlockClosed"
	PrinterStatusDetailsMakeEnvelopeInterlockOpen                 PrinterStatusDetails = "makeEnvelopeInterlockOpen"
	PrinterStatusDetailsMakeEnvelopeJam                           PrinterStatusDetails = "makeEnvelopeJam"
	PrinterStatusDetailsMakeEnvelopeLifeAlmostOver                PrinterStatusDetails = "makeEnvelopeLifeAlmostOver"
	PrinterStatusDetailsMakeEnvelopeLifeOver                      PrinterStatusDetails = "makeEnvelopeLifeOver"
	PrinterStatusDetailsMakeEnvelopeMemoryExhausted               PrinterStatusDetails = "makeEnvelopeMemoryExhausted"
	PrinterStatusDetailsMakeEnvelopeMissing                       PrinterStatusDetails = "makeEnvelopeMissing"
	PrinterStatusDetailsMakeEnvelopeMotorFailure                  PrinterStatusDetails = "makeEnvelopeMotorFailure"
	PrinterStatusDetailsMakeEnvelopeNearLimit                     PrinterStatusDetails = "makeEnvelopeNearLimit"
	PrinterStatusDetailsMakeEnvelopeOffline                       PrinterStatusDetails = "makeEnvelopeOffline"
	PrinterStatusDetailsMakeEnvelopeOpened                        PrinterStatusDetails = "makeEnvelopeOpened"
	PrinterStatusDetailsMakeEnvelopeOverTemperature               PrinterStatusDetails = "makeEnvelopeOverTemperature"
	PrinterStatusDetailsMakeEnvelopePowerSaver                    PrinterStatusDetails = "makeEnvelopePowerSaver"
	PrinterStatusDetailsMakeEnvelopeRecoverableFailure            PrinterStatusDetails = "makeEnvelopeRecoverableFailure"
	PrinterStatusDetailsMakeEnvelopeRecoverableStorage            PrinterStatusDetails = "makeEnvelopeRecoverableStorage"
	PrinterStatusDetailsMakeEnvelopeRemoved                       PrinterStatusDetails = "makeEnvelopeRemoved"
	PrinterStatusDetailsMakeEnvelopeResourceAdded                 PrinterStatusDetails = "makeEnvelopeResourceAdded"
	PrinterStatusDetailsMakeEnvelopeResourceRemoved               PrinterStatusDetails = "makeEnvelopeResourceRemoved"
	PrinterStatusDetailsMakeEnvelopeThermistorFailure             PrinterStatusDetails = "makeEnvelopeThermistorFailure"
	PrinterStatusDetailsMakeEnvelopeTimingFailure                 PrinterStatusDetails = "makeEnvelopeTimingFailure"
	PrinterStatusDetailsMakeEnvelopeTurnedOff                     PrinterStatusDetails = "makeEnvelopeTurnedOff"
	PrinterStatusDetailsMakeEnvelopeTurnedOn                      PrinterStatusDetails = "makeEnvelopeTurnedOn"
	PrinterStatusDetailsMakeEnvelopeUnderTemperature              PrinterStatusDetails = "makeEnvelopeUnderTemperature"
	PrinterStatusDetailsMakeEnvelopeUnrecoverableFailure          PrinterStatusDetails = "makeEnvelopeUnrecoverableFailure"
	PrinterStatusDetailsMakeEnvelopeUnrecoverableStorageError     PrinterStatusDetails = "makeEnvelopeUnrecoverableStorageError"
	PrinterStatusDetailsMakeEnvelopeWarmingUp                     PrinterStatusDetails = "makeEnvelopeWarmingUp"
	PrinterStatusDetailsMarkerAdjustingPrintQuality               PrinterStatusDetails = "markerAdjustingPrintQuality"
	PrinterStatusDetailsMarkerCleanerMissing                      PrinterStatusDetails = "markerCleanerMissing"
	PrinterStatusDetailsMarkerDeveloperAlmostEmpty                PrinterStatusDetails = "markerDeveloperAlmostEmpty"
	PrinterStatusDetailsMarkerDeveloperEmpty                      PrinterStatusDetails = "markerDeveloperEmpty"
	PrinterStatusDetailsMarkerDeveloperMissing                    PrinterStatusDetails = "markerDeveloperMissing"
	PrinterStatusDetailsMarkerFuserMissing                        PrinterStatusDetails = "markerFuserMissing"
	PrinterStatusDetailsMarkerFuserThermistorFailure              PrinterStatusDetails = "markerFuserThermistorFailure"
	PrinterStatusDetailsMarkerFuserTimingFailure                  PrinterStatusDetails = "markerFuserTimingFailure"
	PrinterStatusDetailsMarkerInkAlmostEmpty                      PrinterStatusDetails = "markerInkAlmostEmpty"
	PrinterStatusDetailsMarkerInkEmpty                            PrinterStatusDetails = "markerInkEmpty"
	PrinterStatusDetailsMarkerInkMissing                          PrinterStatusDetails = "markerInkMissing"
	PrinterStatusDetailsMarkerOpcMissing                          PrinterStatusDetails = "markerOpcMissing"
	PrinterStatusDetailsMarkerPrintRibbonAlmostEmpty              PrinterStatusDetails = "markerPrintRibbonAlmostEmpty"
	PrinterStatusDetailsMarkerPrintRibbonEmpty                    PrinterStatusDetails = "markerPrintRibbonEmpty"
	PrinterStatusDetailsMarkerPrintRibbonMissing                  PrinterStatusDetails = "markerPrintRibbonMissing"
	PrinterStatusDetailsMarkerSupplyAlmostEmpty                   PrinterStatusDetails = "markerSupplyAlmostEmpty"
	PrinterStatusDetailsMarkerSupplyEmpty                         PrinterStatusDetails = "markerSupplyEmpty"
	PrinterStatusDetailsMarkerSupplyLow                           PrinterStatusDetails = "markerSupplyLow"
	PrinterStatusDetailsMarkerSupplyMissing                       PrinterStatusDetails = "markerSupplyMissing"
	PrinterStatusDetailsMarkerTonerCartridgeMissing               PrinterStatusDetails = "markerTonerCartridgeMissing"
	PrinterStatusDetailsMarkerTonerMissing                        PrinterStatusDetails = "markerTonerMissing"
	PrinterStatusDetailsMarkerWasteAlmostFull                     PrinterStatusDetails = "markerWasteAlmostFull"
	PrinterStatusDetailsMarkerWasteFull                           PrinterStatusDetails = "markerWasteFull"
	PrinterStatusDetailsMarkerWasteInkReceptacleAlmostFull        PrinterStatusDetails = "markerWasteInkReceptacleAlmostFull"
	PrinterStatusDetailsMarkerWasteInkReceptacleFull              PrinterStatusDetails = "markerWasteInkReceptacleFull"
	PrinterStatusDetailsMarkerWasteInkReceptacleMissing           PrinterStatusDetails = "markerWasteInkReceptacleMissing"
	PrinterStatusDetailsMarkerWasteMissing                        PrinterStatusDetails = "markerWasteMissing"
	PrinterStatusDetailsMarkerWasteTonerReceptacleAlmostFull      PrinterStatusDetails = "markerWasteTonerReceptacleAlmostFull"
	PrinterStatusDetailsMarkerWasteTonerReceptacleFull            PrinterStatusDetails = "markerWasteTonerReceptacleFull"
	PrinterStatusDetailsMarkerWasteTonerReceptacleMissing         PrinterStatusDetails = "markerWasteTonerReceptacleMissing"
	PrinterStatusDetailsMaterialEmpty                             PrinterStatusDetails = "materialEmpty"
	PrinterStatusDetailsMaterialLow                               PrinterStatusDetails = "materialLow"
	PrinterStatusDetailsMaterialNeeded                            PrinterStatusDetails = "materialNeeded"
	PrinterStatusDetailsMediaDrying                               PrinterStatusDetails = "mediaDrying"
	PrinterStatusDetailsMediaEmpty                                PrinterStatusDetails = "mediaEmpty"
	PrinterStatusDetailsMediaJam                                  PrinterStatusDetails = "mediaJam"
	PrinterStatusDetailsMediaLow                                  PrinterStatusDetails = "mediaLow"
	PrinterStatusDetailsMediaNeeded                               PrinterStatusDetails = "mediaNeeded"
	PrinterStatusDetailsMediaPathCannotDuplexMediaSelected        PrinterStatusDetails = "mediaPathCannotDuplexMediaSelected"
	PrinterStatusDetailsMediaPathFailure                          PrinterStatusDetails = "mediaPathFailure"
	PrinterStatusDetailsMediaPathInputEmpty                       PrinterStatusDetails = "mediaPathInputEmpty"
	PrinterStatusDetailsMediaPathInputFeedError                   PrinterStatusDetails = "mediaPathInputFeedError"
	PrinterStatusDetailsMediaPathInputJam                         PrinterStatusDetails = "mediaPathInputJam"
	PrinterStatusDetailsMediaPathInputRequest                     PrinterStatusDetails = "mediaPathInputRequest"
	PrinterStatusDetailsMediaPathJam                              PrinterStatusDetails = "mediaPathJam"
	PrinterStatusDetailsMediaPathMediaTrayAlmostFull              PrinterStatusDetails = "mediaPathMediaTrayAlmostFull"
	PrinterStatusDetailsMediaPathMediaTrayFull                    PrinterStatusDetails = "mediaPathMediaTrayFull"
	PrinterStatusDetailsMediaPathMediaTrayMissing                 PrinterStatusDetails = "mediaPathMediaTrayMissing"
	PrinterStatusDetailsMediaPathOutputFeedError                  PrinterStatusDetails = "mediaPathOutputFeedError"
	PrinterStatusDetailsMediaPathOutputFull                       PrinterStatusDetails = "mediaPathOutputFull"
	PrinterStatusDetailsMediaPathOutputJam                        PrinterStatusDetails = "mediaPathOutputJam"
	PrinterStatusDetailsMediaPathPickRollerFailure                PrinterStatusDetails = "mediaPathPickRollerFailure"
	PrinterStatusDetailsMediaPathPickRollerLifeOver               PrinterStatusDetails = "mediaPathPickRollerLifeOver"
	PrinterStatusDetailsMediaPathPickRollerLifeWarn               PrinterStatusDetails = "mediaPathPickRollerLifeWarn"
	PrinterStatusDetailsMediaPathPickRollerMissing                PrinterStatusDetails = "mediaPathPickRollerMissing"
	PrinterStatusDetailsMotorFailure                              PrinterStatusDetails = "motorFailure"
	PrinterStatusDetailsMovingToPaused                            PrinterStatusDetails = "movingToPaused"
	PrinterStatusDetailsNone                                      PrinterStatusDetails = "none"
	PrinterStatusDetailsOpticalPhotoConductorLifeOver             PrinterStatusDetails = "opticalPhotoConductorLifeOver"
	PrinterStatusDetailsOpticalPhotoConductorNearEndOfLife        PrinterStatusDetails = "opticalPhotoConductorNearEndOfLife"
	PrinterStatusDetailsOther                                     PrinterStatusDetails = "other"
	PrinterStatusDetailsOutputAreaAlmostFull                      PrinterStatusDetails = "outputAreaAlmostFull"
	PrinterStatusDetailsOutputAreaFull                            PrinterStatusDetails = "outputAreaFull"
	PrinterStatusDetailsOutputMailboxSelectFailure                PrinterStatusDetails = "outputMailboxSelectFailure"
	PrinterStatusDetailsOutputMediaTrayFailure                    PrinterStatusDetails = "outputMediaTrayFailure"
	PrinterStatusDetailsOutputMediaTrayFeedError                  PrinterStatusDetails = "outputMediaTrayFeedError"
	PrinterStatusDetailsOutputMediaTrayJam                        PrinterStatusDetails = "outputMediaTrayJam"
	PrinterStatusDetailsOutputTrayMissing                         PrinterStatusDetails = "outputTrayMissing"
	PrinterStatusDetailsPaused                                    PrinterStatusDetails = "paused"
	PrinterStatusDetailsPerforaterAdded                           PrinterStatusDetails = "perforaterAdded"
	PrinterStatusDetailsPerforaterAlmostEmpty                     PrinterStatusDetails = "perforaterAlmostEmpty"
	PrinterStatusDetailsPerforaterAlmostFull                      PrinterStatusDetails = "perforaterAlmostFull"
	PrinterStatusDetailsPerforaterAtLimit                         PrinterStatusDetails = "perforaterAtLimit"
	PrinterStatusDetailsPerforaterClosed                          PrinterStatusDetails = "perforaterClosed"
	PrinterStatusDetailsPerforaterConfigurationChange             PrinterStatusDetails = "perforaterConfigurationChange"
	PrinterStatusDetailsPerforaterCoverClosed                     PrinterStatusDetails = "perforaterCoverClosed"
	PrinterStatusDetailsPerforaterCoverOpen                       PrinterStatusDetails = "perforaterCoverOpen"
	PrinterStatusDetailsPerforaterEmpty                           PrinterStatusDetails = "perforaterEmpty"
	PrinterStatusDetailsPerforaterFull                            PrinterStatusDetails = "perforaterFull"
	PrinterStatusDetailsPerforaterInterlockClosed                 PrinterStatusDetails = "perforaterInterlockClosed"
	PrinterStatusDetailsPerforaterInterlockOpen                   PrinterStatusDetails = "perforaterInterlockOpen"
	PrinterStatusDetailsPerforaterJam                             PrinterStatusDetails = "perforaterJam"
	PrinterStatusDetailsPerforaterLifeAlmostOver                  PrinterStatusDetails = "perforaterLifeAlmostOver"
	PrinterStatusDetailsPerforaterLifeOver                        PrinterStatusDetails = "perforaterLifeOver"
	PrinterStatusDetailsPerforaterMemoryExhausted                 PrinterStatusDetails = "perforaterMemoryExhausted"
	PrinterStatusDetailsPerforaterMissing                         PrinterStatusDetails = "perforaterMissing"
	PrinterStatusDetailsPerforaterMotorFailure                    PrinterStatusDetails = "perforaterMotorFailure"
	PrinterStatusDetailsPerforaterNearLimit                       PrinterStatusDetails = "perforaterNearLimit"
	PrinterStatusDetailsPerforaterOffline                         PrinterStatusDetails = "perforaterOffline"
	PrinterStatusDetailsPerforaterOpened                          PrinterStatusDetails = "perforaterOpened"
	PrinterStatusDetailsPerforaterOverTemperature                 PrinterStatusDetails = "perforaterOverTemperature"
	PrinterStatusDetailsPerforaterPowerSaver                      PrinterStatusDetails = "perforaterPowerSaver"
	PrinterStatusDetailsPerforaterRecoverableFailure              PrinterStatusDetails = "perforaterRecoverableFailure"
	PrinterStatusDetailsPerforaterRecoverableStorage              PrinterStatusDetails = "perforaterRecoverableStorage"
	PrinterStatusDetailsPerforaterRemoved                         PrinterStatusDetails = "perforaterRemoved"
	PrinterStatusDetailsPerforaterResourceAdded                   PrinterStatusDetails = "perforaterResourceAdded"
	PrinterStatusDetailsPerforaterResourceRemoved                 PrinterStatusDetails = "perforaterResourceRemoved"
	PrinterStatusDetailsPerforaterThermistorFailure               PrinterStatusDetails = "perforaterThermistorFailure"
	PrinterStatusDetailsPerforaterTimingFailure                   PrinterStatusDetails = "perforaterTimingFailure"
	PrinterStatusDetailsPerforaterTurnedOff                       PrinterStatusDetails = "perforaterTurnedOff"
	PrinterStatusDetailsPerforaterTurnedOn                        PrinterStatusDetails = "perforaterTurnedOn"
	PrinterStatusDetailsPerforaterUnderTemperature                PrinterStatusDetails = "perforaterUnderTemperature"
	PrinterStatusDetailsPerforaterUnrecoverableFailure            PrinterStatusDetails = "perforaterUnrecoverableFailure"
	PrinterStatusDetailsPerforaterUnrecoverableStorageError       PrinterStatusDetails = "perforaterUnrecoverableStorageError"
	PrinterStatusDetailsPerforaterWarmingUp                       PrinterStatusDetails = "perforaterWarmingUp"
	PrinterStatusDetailsPlatformCooling                           PrinterStatusDetails = "platformCooling"
	PrinterStatusDetailsPlatformFailure                           PrinterStatusDetails = "platformFailure"
	PrinterStatusDetailsPlatformHeating                           PrinterStatusDetails = "platformHeating"
	PrinterStatusDetailsPlatformTemperatureHigh                   PrinterStatusDetails = "platformTemperatureHigh"
	PrinterStatusDetailsPlatformTemperatureLow                    PrinterStatusDetails = "platformTemperatureLow"
	PrinterStatusDetailsPowerDown                                 PrinterStatusDetails = "powerDown"
	PrinterStatusDetailsPowerUp                                   PrinterStatusDetails = "powerUp"
	PrinterStatusDetailsPrinterManualReset                        PrinterStatusDetails = "printerManualReset"
	PrinterStatusDetailsPrinterNmsReset                           PrinterStatusDetails = "printerNmsReset"
	PrinterStatusDetailsPrinterReadyToPrint                       PrinterStatusDetails = "printerReadyToPrint"
	PrinterStatusDetailsPuncherAdded                              PrinterStatusDetails = "puncherAdded"
	PrinterStatusDetailsPuncherAlmostEmpty                        PrinterStatusDetails = "puncherAlmostEmpty"
	PrinterStatusDetailsPuncherAlmostFull                         PrinterStatusDetails = "puncherAlmostFull"
	PrinterStatusDetailsPuncherAtLimit                            PrinterStatusDetails = "puncherAtLimit"
	PrinterStatusDetailsPuncherClosed                             PrinterStatusDetails = "puncherClosed"
	PrinterStatusDetailsPuncherConfigurationChange                PrinterStatusDetails = "puncherConfigurationChange"
	PrinterStatusDetailsPuncherCoverClosed                        PrinterStatusDetails = "puncherCoverClosed"
	PrinterStatusDetailsPuncherCoverOpen                          PrinterStatusDetails = "puncherCoverOpen"
	PrinterStatusDetailsPuncherEmpty                              PrinterStatusDetails = "puncherEmpty"
	PrinterStatusDetailsPuncherFull                               PrinterStatusDetails = "puncherFull"
	PrinterStatusDetailsPuncherInterlockClosed                    PrinterStatusDetails = "puncherInterlockClosed"
	PrinterStatusDetailsPuncherInterlockOpen                      PrinterStatusDetails = "puncherInterlockOpen"
	PrinterStatusDetailsPuncherJam                                PrinterStatusDetails = "puncherJam"
	PrinterStatusDetailsPuncherLifeAlmostOver                     PrinterStatusDetails = "puncherLifeAlmostOver"
	PrinterStatusDetailsPuncherLifeOver                           PrinterStatusDetails = "puncherLifeOver"
	PrinterStatusDetailsPuncherMemoryExhausted                    PrinterStatusDetails = "puncherMemoryExhausted"
	PrinterStatusDetailsPuncherMissing                            PrinterStatusDetails = "puncherMissing"
	PrinterStatusDetailsPuncherMotorFailure                       PrinterStatusDetails = "puncherMotorFailure"
	PrinterStatusDetailsPuncherNearLimit                          PrinterStatusDetails = "puncherNearLimit"
	PrinterStatusDetailsPuncherOffline                            PrinterStatusDetails = "puncherOffline"
	PrinterStatusDetailsPuncherOpened                             PrinterStatusDetails = "puncherOpened"
	PrinterStatusDetailsPuncherOverTemperature                    PrinterStatusDetails = "puncherOverTemperature"
	PrinterStatusDetailsPuncherPowerSaver                         PrinterStatusDetails = "puncherPowerSaver"
	PrinterStatusDetailsPuncherRecoverableFailure                 PrinterStatusDetails = "puncherRecoverableFailure"
	PrinterStatusDetailsPuncherRecoverableStorage                 PrinterStatusDetails = "puncherRecoverableStorage"
	PrinterStatusDetailsPuncherRemoved                            PrinterStatusDetails = "puncherRemoved"
	PrinterStatusDetailsPuncherResourceAdded                      PrinterStatusDetails = "puncherResourceAdded"
	PrinterStatusDetailsPuncherResourceRemoved                    PrinterStatusDetails = "puncherResourceRemoved"
	PrinterStatusDetailsPuncherThermistorFailure                  PrinterStatusDetails = "puncherThermistorFailure"
	PrinterStatusDetailsPuncherTimingFailure                      PrinterStatusDetails = "puncherTimingFailure"
	PrinterStatusDetailsPuncherTurnedOff                          PrinterStatusDetails = "puncherTurnedOff"
	PrinterStatusDetailsPuncherTurnedOn                           PrinterStatusDetails = "puncherTurnedOn"
	PrinterStatusDetailsPuncherUnderTemperature                   PrinterStatusDetails = "puncherUnderTemperature"
	PrinterStatusDetailsPuncherUnrecoverableFailure               PrinterStatusDetails = "puncherUnrecoverableFailure"
	PrinterStatusDetailsPuncherUnrecoverableStorageError          PrinterStatusDetails = "puncherUnrecoverableStorageError"
	PrinterStatusDetailsPuncherWarmingUp                          PrinterStatusDetails = "puncherWarmingUp"
	PrinterStatusDetailsResuming                                  PrinterStatusDetails = "resuming"
	PrinterStatusDetailsScanMediaPathFailure                      PrinterStatusDetails = "scanMediaPathFailure"
	PrinterStatusDetailsScanMediaPathInputEmpty                   PrinterStatusDetails = "scanMediaPathInputEmpty"
	PrinterStatusDetailsScanMediaPathInputFeedError               PrinterStatusDetails = "scanMediaPathInputFeedError"
	PrinterStatusDetailsScanMediaPathInputJam                     PrinterStatusDetails = "scanMediaPathInputJam"
	PrinterStatusDetailsScanMediaPathInputRequest                 PrinterStatusDetails = "scanMediaPathInputRequest"
	PrinterStatusDetailsScanMediaPathJam                          PrinterStatusDetails = "scanMediaPathJam"
	PrinterStatusDetailsScanMediaPathOutputFeedError              PrinterStatusDetails = "scanMediaPathOutputFeedError"
	PrinterStatusDetailsScanMediaPathOutputFull                   PrinterStatusDetails = "scanMediaPathOutputFull"
	PrinterStatusDetailsScanMediaPathOutputJam                    PrinterStatusDetails = "scanMediaPathOutputJam"
	PrinterStatusDetailsScanMediaPathPickRollerFailure            PrinterStatusDetails = "scanMediaPathPickRollerFailure"
	PrinterStatusDetailsScanMediaPathPickRollerLifeOver           PrinterStatusDetails = "scanMediaPathPickRollerLifeOver"
	PrinterStatusDetailsScanMediaPathPickRollerLifeWarn           PrinterStatusDetails = "scanMediaPathPickRollerLifeWarn"
	PrinterStatusDetailsScanMediaPathPickRollerMissing            PrinterStatusDetails = "scanMediaPathPickRollerMissing"
	PrinterStatusDetailsScanMediaPathTrayAlmostFull               PrinterStatusDetails = "scanMediaPathTrayAlmostFull"
	PrinterStatusDetailsScanMediaPathTrayFull                     PrinterStatusDetails = "scanMediaPathTrayFull"
	PrinterStatusDetailsScanMediaPathTrayMissing                  PrinterStatusDetails = "scanMediaPathTrayMissing"
	PrinterStatusDetailsScannerLightFailure                       PrinterStatusDetails = "scannerLightFailure"
	PrinterStatusDetailsScannerLightLifeAlmostOver                PrinterStatusDetails = "scannerLightLifeAlmostOver"
	PrinterStatusDetailsScannerLightLifeOver                      PrinterStatusDetails = "scannerLightLifeOver"
	PrinterStatusDetailsScannerLightMissing                       PrinterStatusDetails = "scannerLightMissing"
	PrinterStatusDetailsScannerSensorFailure                      PrinterStatusDetails = "scannerSensorFailure"
	PrinterStatusDetailsScannerSensorLifeAlmostOver               PrinterStatusDetails = "scannerSensorLifeAlmostOver"
	PrinterStatusDetailsScannerSensorLifeOver                     PrinterStatusDetails = "scannerSensorLifeOver"
	PrinterStatusDetailsScannerSensorMissing                      PrinterStatusDetails = "scannerSensorMissing"
	PrinterStatusDetailsSeparationCutterAdded                     PrinterStatusDetails = "separationCutterAdded"
	PrinterStatusDetailsSeparationCutterAlmostEmpty               PrinterStatusDetails = "separationCutterAlmostEmpty"
	PrinterStatusDetailsSeparationCutterAlmostFull                PrinterStatusDetails = "separationCutterAlmostFull"
	PrinterStatusDetailsSeparationCutterAtLimit                   PrinterStatusDetails = "separationCutterAtLimit"
	PrinterStatusDetailsSeparationCutterClosed                    PrinterStatusDetails = "separationCutterClosed"
	PrinterStatusDetailsSeparationCutterConfigurationChange       PrinterStatusDetails = "separationCutterConfigurationChange"
	PrinterStatusDetailsSeparationCutterCoverClosed               PrinterStatusDetails = "separationCutterCoverClosed"
	PrinterStatusDetailsSeparationCutterCoverOpen                 PrinterStatusDetails = "separationCutterCoverOpen"
	PrinterStatusDetailsSeparationCutterEmpty                     PrinterStatusDetails = "separationCutterEmpty"
	PrinterStatusDetailsSeparationCutterFull                      PrinterStatusDetails = "separationCutterFull"
	PrinterStatusDetailsSeparationCutterInterlockClosed           PrinterStatusDetails = "separationCutterInterlockClosed"
	PrinterStatusDetailsSeparationCutterInterlockOpen             PrinterStatusDetails = "separationCutterInterlockOpen"
	PrinterStatusDetailsSeparationCutterJam                       PrinterStatusDetails = "separationCutterJam"
	PrinterStatusDetailsSeparationCutterLifeAlmostOver            PrinterStatusDetails = "separationCutterLifeAlmostOver"
	PrinterStatusDetailsSeparationCutterLifeOver                  PrinterStatusDetails = "separationCutterLifeOver"
	PrinterStatusDetailsSeparationCutterMemoryExhausted           PrinterStatusDetails = "separationCutterMemoryExhausted"
	PrinterStatusDetailsSeparationCutterMissing                   PrinterStatusDetails = "separationCutterMissing"
	PrinterStatusDetailsSeparationCutterMotorFailure              PrinterStatusDetails = "separationCutterMotorFailure"
	PrinterStatusDetailsSeparationCutterNearLimit                 PrinterStatusDetails = "separationCutterNearLimit"
	PrinterStatusDetailsSeparationCutterOffline                   PrinterStatusDetails = "separationCutterOffline"
	PrinterStatusDetailsSeparationCutterOpened                    PrinterStatusDetails = "separationCutterOpened"
	PrinterStatusDetailsSeparationCutterOverTemperature           PrinterStatusDetails = "separationCutterOverTemperature"
	PrinterStatusDetailsSeparationCutterPowerSaver                PrinterStatusDetails = "separationCutterPowerSaver"
	PrinterStatusDetailsSeparationCutterRecoverableFailure        PrinterStatusDetails = "separationCutterRecoverableFailure"
	PrinterStatusDetailsSeparationCutterRecoverableStorage        PrinterStatusDetails = "separationCutterRecoverableStorage"
	PrinterStatusDetailsSeparationCutterRemoved                   PrinterStatusDetails = "separationCutterRemoved"
	PrinterStatusDetailsSeparationCutterResourceAdded             PrinterStatusDetails = "separationCutterResourceAdded"
	PrinterStatusDetailsSeparationCutterResourceRemoved           PrinterStatusDetails = "separationCutterResourceRemoved"
	PrinterStatusDetailsSeparationCutterThermistorFailure         PrinterStatusDetails = "separationCutterThermistorFailure"
	PrinterStatusDetailsSeparationCutterTimingFailure             PrinterStatusDetails = "separationCutterTimingFailure"
	PrinterStatusDetailsSeparationCutterTurnedOff                 PrinterStatusDetails = "separationCutterTurnedOff"
	PrinterStatusDetailsSeparationCutterTurnedOn                  PrinterStatusDetails = "separationCutterTurnedOn"
	PrinterStatusDetailsSeparationCutterUnderTemperature          PrinterStatusDetails = "separationCutterUnderTemperature"
	PrinterStatusDetailsSeparationCutterUnrecoverableFailure      PrinterStatusDetails = "separationCutterUnrecoverableFailure"
	PrinterStatusDetailsSeparationCutterUnrecoverableStorageError PrinterStatusDetails = "separationCutterUnrecoverableStorageError"
	PrinterStatusDetailsSeparationCutterWarmingUp                 PrinterStatusDetails = "separationCutterWarmingUp"
	PrinterStatusDetailsSheetRotatorAdded                         PrinterStatusDetails = "sheetRotatorAdded"
	PrinterStatusDetailsSheetRotatorAlmostEmpty                   PrinterStatusDetails = "sheetRotatorAlmostEmpty"
	PrinterStatusDetailsSheetRotatorAlmostFull                    PrinterStatusDetails = "sheetRotatorAlmostFull"
	PrinterStatusDetailsSheetRotatorAtLimit                       PrinterStatusDetails = "sheetRotatorAtLimit"
	PrinterStatusDetailsSheetRotatorClosed                        PrinterStatusDetails = "sheetRotatorClosed"
	PrinterStatusDetailsSheetRotatorConfigurationChange           PrinterStatusDetails = "sheetRotatorConfigurationChange"
	PrinterStatusDetailsSheetRotatorCoverClosed                   PrinterStatusDetails = "sheetRotatorCoverClosed"
	PrinterStatusDetailsSheetRotatorCoverOpen                     PrinterStatusDetails = "sheetRotatorCoverOpen"
	PrinterStatusDetailsSheetRotatorEmpty                         PrinterStatusDetails = "sheetRotatorEmpty"
	PrinterStatusDetailsSheetRotatorFull                          PrinterStatusDetails = "sheetRotatorFull"
	PrinterStatusDetailsSheetRotatorInterlockClosed               PrinterStatusDetails = "sheetRotatorInterlockClosed"
	PrinterStatusDetailsSheetRotatorInterlockOpen                 PrinterStatusDetails = "sheetRotatorInterlockOpen"
	PrinterStatusDetailsSheetRotatorJam                           PrinterStatusDetails = "sheetRotatorJam"
	PrinterStatusDetailsSheetRotatorLifeAlmostOver                PrinterStatusDetails = "sheetRotatorLifeAlmostOver"
	PrinterStatusDetailsSheetRotatorLifeOver                      PrinterStatusDetails = "sheetRotatorLifeOver"
	PrinterStatusDetailsSheetRotatorMemoryExhausted               PrinterStatusDetails = "sheetRotatorMemoryExhausted"
	PrinterStatusDetailsSheetRotatorMissing                       PrinterStatusDetails = "sheetRotatorMissing"
	PrinterStatusDetailsSheetRotatorMotorFailure                  PrinterStatusDetails = "sheetRotatorMotorFailure"
	PrinterStatusDetailsSheetRotatorNearLimit                     PrinterStatusDetails = "sheetRotatorNearLimit"
	PrinterStatusDetailsSheetRotatorOffline                       PrinterStatusDetails = "sheetRotatorOffline"
	PrinterStatusDetailsSheetRotatorOpened                        PrinterStatusDetails = "sheetRotatorOpened"
	PrinterStatusDetailsSheetRotatorOverTemperature               PrinterStatusDetails = "sheetRotatorOverTemperature"
	PrinterStatusDetailsSheetRotatorPowerSaver                    PrinterStatusDetails = "sheetRotatorPowerSaver"
	PrinterStatusDetailsSheetRotatorRecoverableFailure            PrinterStatusDetails = "sheetRotatorRecoverableFailure"
	PrinterStatusDetailsSheetRotatorRecoverableStorage            PrinterStatusDetails = "sheetRotatorRecoverableStorage"
	PrinterStatusDetailsSheetRotatorRemoved                       PrinterStatusDetails = "sheetRotatorRemoved"
	PrinterStatusDetailsSheetRotatorResourceAdded                 PrinterStatusDetails = "sheetRotatorResourceAdded"
	PrinterStatusDetailsSheetRotatorResourceRemoved               PrinterStatusDetails = "sheetRotatorResourceRemoved"
	PrinterStatusDetailsSheetRotatorThermistorFailure             PrinterStatusDetails = "sheetRotatorThermistorFailure"
	PrinterStatusDetailsSheetRotatorTimingFailure                 PrinterStatusDetails = "sheetRotatorTimingFailure"
	PrinterStatusDetailsSheetRotatorTurnedOff                     PrinterStatusDetails = "sheetRotatorTurnedOff"
	PrinterStatusDetailsSheetRotatorTurnedOn                      PrinterStatusDetails = "sheetRotatorTurnedOn"
	PrinterStatusDetailsSheetRotatorUnderTemperature              PrinterStatusDetails = "sheetRotatorUnderTemperature"
	PrinterStatusDetailsSheetRotatorUnrecoverableFailure          PrinterStatusDetails = "sheetRotatorUnrecoverableFailure"
	PrinterStatusDetailsSheetRotatorUnrecoverableStorageError     PrinterStatusDetails = "sheetRotatorUnrecoverableStorageError"
	PrinterStatusDetailsSheetRotatorWarmingUp                     PrinterStatusDetails = "sheetRotatorWarmingUp"
	PrinterStatusDetailsShutdown                                  PrinterStatusDetails = "shutdown"
	PrinterStatusDetailsSlitterAdded                              PrinterStatusDetails = "slitterAdded"
	PrinterStatusDetailsSlitterAlmostEmpty                        PrinterStatusDetails = "slitterAlmostEmpty"
	PrinterStatusDetailsSlitterAlmostFull                         PrinterStatusDetails = "slitterAlmostFull"
	PrinterStatusDetailsSlitterAtLimit                            PrinterStatusDetails = "slitterAtLimit"
	PrinterStatusDetailsSlitterClosed                             PrinterStatusDetails = "slitterClosed"
	PrinterStatusDetailsSlitterConfigurationChange                PrinterStatusDetails = "slitterConfigurationChange"
	PrinterStatusDetailsSlitterCoverClosed                        PrinterStatusDetails = "slitterCoverClosed"
	PrinterStatusDetailsSlitterCoverOpen                          PrinterStatusDetails = "slitterCoverOpen"
	PrinterStatusDetailsSlitterEmpty                              PrinterStatusDetails = "slitterEmpty"
	PrinterStatusDetailsSlitterFull                               PrinterStatusDetails = "slitterFull"
	PrinterStatusDetailsSlitterInterlockClosed                    PrinterStatusDetails = "slitterInterlockClosed"
	PrinterStatusDetailsSlitterInterlockOpen                      PrinterStatusDetails = "slitterInterlockOpen"
	PrinterStatusDetailsSlitterJam                                PrinterStatusDetails = "slitterJam"
	PrinterStatusDetailsSlitterLifeAlmostOver                     PrinterStatusDetails = "slitterLifeAlmostOver"
	PrinterStatusDetailsSlitterLifeOver                           PrinterStatusDetails = "slitterLifeOver"
	PrinterStatusDetailsSlitterMemoryExhausted                    PrinterStatusDetails = "slitterMemoryExhausted"
	PrinterStatusDetailsSlitterMissing                            PrinterStatusDetails = "slitterMissing"
	PrinterStatusDetailsSlitterMotorFailure                       PrinterStatusDetails = "slitterMotorFailure"
	PrinterStatusDetailsSlitterNearLimit                          PrinterStatusDetails = "slitterNearLimit"
	PrinterStatusDetailsSlitterOffline                            PrinterStatusDetails = "slitterOffline"
	PrinterStatusDetailsSlitterOpened                             PrinterStatusDetails = "slitterOpened"
	PrinterStatusDetailsSlitterOverTemperature                    PrinterStatusDetails = "slitterOverTemperature"
	PrinterStatusDetailsSlitterPowerSaver                         PrinterStatusDetails = "slitterPowerSaver"
	PrinterStatusDetailsSlitterRecoverableFailure                 PrinterStatusDetails = "slitterRecoverableFailure"
	PrinterStatusDetailsSlitterRecoverableStorage                 PrinterStatusDetails = "slitterRecoverableStorage"
	PrinterStatusDetailsSlitterRemoved                            PrinterStatusDetails = "slitterRemoved"
	PrinterStatusDetailsSlitterResourceAdded                      PrinterStatusDetails = "slitterResourceAdded"
	PrinterStatusDetailsSlitterResourceRemoved                    PrinterStatusDetails = "slitterResourceRemoved"
	PrinterStatusDetailsSlitterThermistorFailure                  PrinterStatusDetails = "slitterThermistorFailure"
	PrinterStatusDetailsSlitterTimingFailure                      PrinterStatusDetails = "slitterTimingFailure"
	PrinterStatusDetailsSlitterTurnedOff                          PrinterStatusDetails = "slitterTurnedOff"
	PrinterStatusDetailsSlitterTurnedOn                           PrinterStatusDetails = "slitterTurnedOn"
	PrinterStatusDetailsSlitterUnderTemperature                   PrinterStatusDetails = "slitterUnderTemperature"
	PrinterStatusDetailsSlitterUnrecoverableFailure               PrinterStatusDetails = "slitterUnrecoverableFailure"
	PrinterStatusDetailsSlitterUnrecoverableStorageError          PrinterStatusDetails = "slitterUnrecoverableStorageError"
	PrinterStatusDetailsSlitterWarmingUp                          PrinterStatusDetails = "slitterWarmingUp"
	PrinterStatusDetailsSpoolAreaFull                             PrinterStatusDetails = "spoolAreaFull"
	PrinterStatusDetailsStackerAdded                              PrinterStatusDetails = "stackerAdded"
	PrinterStatusDetailsStackerAlmostEmpty                        PrinterStatusDetails = "stackerAlmostEmpty"
	PrinterStatusDetailsStackerAlmostFull                         PrinterStatusDetails = "stackerAlmostFull"
	PrinterStatusDetailsStackerAtLimit                            PrinterStatusDetails = "stackerAtLimit"
	PrinterStatusDetailsStackerClosed                             PrinterStatusDetails = "stackerClosed"
	PrinterStatusDetailsStackerConfigurationChange                PrinterStatusDetails = "stackerConfigurationChange"
	PrinterStatusDetailsStackerCoverClosed                        PrinterStatusDetails = "stackerCoverClosed"
	PrinterStatusDetailsStackerCoverOpen                          PrinterStatusDetails = "stackerCoverOpen"
	PrinterStatusDetailsStackerEmpty                              PrinterStatusDetails = "stackerEmpty"
	PrinterStatusDetailsStackerFull                               PrinterStatusDetails = "stackerFull"
	PrinterStatusDetailsStackerInterlockClosed                    PrinterStatusDetails = "stackerInterlockClosed"
	PrinterStatusDetailsStackerInterlockOpen                      PrinterStatusDetails = "stackerInterlockOpen"
	PrinterStatusDetailsStackerJam                                PrinterStatusDetails = "stackerJam"
	PrinterStatusDetailsStackerLifeAlmostOver                     PrinterStatusDetails = "stackerLifeAlmostOver"
	PrinterStatusDetailsStackerLifeOver                           PrinterStatusDetails = "stackerLifeOver"
	PrinterStatusDetailsStackerMemoryExhausted                    PrinterStatusDetails = "stackerMemoryExhausted"
	PrinterStatusDetailsStackerMissing                            PrinterStatusDetails = "stackerMissing"
	PrinterStatusDetailsStackerMotorFailure                       PrinterStatusDetails = "stackerMotorFailure"
	PrinterStatusDetailsStackerNearLimit                          PrinterStatusDetails = "stackerNearLimit"
	PrinterStatusDetailsStackerOffline                            PrinterStatusDetails = "stackerOffline"
	PrinterStatusDetailsStackerOpened                             PrinterStatusDetails = "stackerOpened"
	PrinterStatusDetailsStackerOverTemperature                    PrinterStatusDetails = "stackerOverTemperature"
	PrinterStatusDetailsStackerPowerSaver                         PrinterStatusDetails = "stackerPowerSaver"
	PrinterStatusDetailsStackerRecoverableFailure                 PrinterStatusDetails = "stackerRecoverableFailure"
	PrinterStatusDetailsStackerRecoverableStorage                 PrinterStatusDetails = "stackerRecoverableStorage"
	PrinterStatusDetailsStackerRemoved                            PrinterStatusDetails = "stackerRemoved"
	PrinterStatusDetailsStackerResourceAdded                      PrinterStatusDetails = "stackerResourceAdded"
	PrinterStatusDetailsStackerResourceRemoved                    PrinterStatusDetails = "stackerResourceRemoved"
	PrinterStatusDetailsStackerThermistorFailure                  PrinterStatusDetails = "stackerThermistorFailure"
	PrinterStatusDetailsStackerTimingFailure                      PrinterStatusDetails = "stackerTimingFailure"
	PrinterStatusDetailsStackerTurnedOff                          PrinterStatusDetails = "stackerTurnedOff"
	PrinterStatusDetailsStackerTurnedOn                           PrinterStatusDetails = "stackerTurnedOn"
	PrinterStatusDetailsStackerUnderTemperature                   PrinterStatusDetails = "stackerUnderTemperature"
	PrinterStatusDetailsStackerUnrecoverableFailure               PrinterStatusDetails = "stackerUnrecoverableFailure"
	PrinterStatusDetailsStackerUnrecoverableStorageError          PrinterStatusDetails = "stackerUnrecoverableStorageError"
	PrinterStatusDetailsStackerWarmingUp                          PrinterStatusDetails = "stackerWarmingUp"
	PrinterStatusDetailsStandby                                   PrinterStatusDetails = "standby"
	PrinterStatusDetailsStaplerAdded                              PrinterStatusDetails = "staplerAdded"
	PrinterStatusDetailsStaplerAlmostEmpty                        PrinterStatusDetails = "staplerAlmostEmpty"
	PrinterStatusDetailsStaplerAlmostFull                         PrinterStatusDetails = "staplerAlmostFull"
	PrinterStatusDetailsStaplerAtLimit                            PrinterStatusDetails = "staplerAtLimit"
	PrinterStatusDetailsStaplerClosed                             PrinterStatusDetails = "staplerClosed"
	PrinterStatusDetailsStaplerConfigurationChange                PrinterStatusDetails = "staplerConfigurationChange"
	PrinterStatusDetailsStaplerCoverClosed                        PrinterStatusDetails = "staplerCoverClosed"
	PrinterStatusDetailsStaplerCoverOpen                          PrinterStatusDetails = "staplerCoverOpen"
	PrinterStatusDetailsStaplerEmpty                              PrinterStatusDetails = "staplerEmpty"
	PrinterStatusDetailsStaplerFull                               PrinterStatusDetails = "staplerFull"
	PrinterStatusDetailsStaplerInterlockClosed                    PrinterStatusDetails = "staplerInterlockClosed"
	PrinterStatusDetailsStaplerInterlockOpen                      PrinterStatusDetails = "staplerInterlockOpen"
	PrinterStatusDetailsStaplerJam                                PrinterStatusDetails = "staplerJam"
	PrinterStatusDetailsStaplerLifeAlmostOver                     PrinterStatusDetails = "staplerLifeAlmostOver"
	PrinterStatusDetailsStaplerLifeOver                           PrinterStatusDetails = "staplerLifeOver"
	PrinterStatusDetailsStaplerMemoryExhausted                    PrinterStatusDetails = "staplerMemoryExhausted"
	PrinterStatusDetailsStaplerMissing                            PrinterStatusDetails = "staplerMissing"
	PrinterStatusDetailsStaplerMotorFailure                       PrinterStatusDetails = "staplerMotorFailure"
	PrinterStatusDetailsStaplerNearLimit                          PrinterStatusDetails = "staplerNearLimit"
	PrinterStatusDetailsStaplerOffline                            PrinterStatusDetails = "staplerOffline"
	PrinterStatusDetailsStaplerOpened                             PrinterStatusDetails = "staplerOpened"
	PrinterStatusDetailsStaplerOverTemperature                    PrinterStatusDetails = "staplerOverTemperature"
	PrinterStatusDetailsStaplerPowerSaver                         PrinterStatusDetails = "staplerPowerSaver"
	PrinterStatusDetailsStaplerRecoverableFailure                 PrinterStatusDetails = "staplerRecoverableFailure"
	PrinterStatusDetailsStaplerRecoverableStorage                 PrinterStatusDetails = "staplerRecoverableStorage"
	PrinterStatusDetailsStaplerRemoved                            PrinterStatusDetails = "staplerRemoved"
	PrinterStatusDetailsStaplerResourceAdded                      PrinterStatusDetails = "staplerResourceAdded"
	PrinterStatusDetailsStaplerResourceRemoved                    PrinterStatusDetails = "staplerResourceRemoved"
	PrinterStatusDetailsStaplerThermistorFailure                  PrinterStatusDetails = "staplerThermistorFailure"
	PrinterStatusDetailsStaplerTimingFailure                      PrinterStatusDetails = "staplerTimingFailure"
	PrinterStatusDetailsStaplerTurnedOff                          PrinterStatusDetails = "staplerTurnedOff"
	PrinterStatusDetailsStaplerTurnedOn                           PrinterStatusDetails = "staplerTurnedOn"
	PrinterStatusDetailsStaplerUnderTemperature                   PrinterStatusDetails = "staplerUnderTemperature"
	PrinterStatusDetailsStaplerUnrecoverableFailure               PrinterStatusDetails = "staplerUnrecoverableFailure"
	PrinterStatusDetailsStaplerUnrecoverableStorageError          PrinterStatusDetails = "staplerUnrecoverableStorageError"
	PrinterStatusDetailsStaplerWarmingUp                          PrinterStatusDetails = "staplerWarmingUp"
	PrinterStatusDetailsStitcherAdded                             PrinterStatusDetails = "stitcherAdded"
	PrinterStatusDetailsStitcherAlmostEmpty                       PrinterStatusDetails = "stitcherAlmostEmpty"
	PrinterStatusDetailsStitcherAlmostFull                        PrinterStatusDetails = "stitcherAlmostFull"
	PrinterStatusDetailsStitcherAtLimit                           PrinterStatusDetails = "stitcherAtLimit"
	PrinterStatusDetailsStitcherClosed                            PrinterStatusDetails = "stitcherClosed"
	PrinterStatusDetailsStitcherConfigurationChange               PrinterStatusDetails = "stitcherConfigurationChange"
	PrinterStatusDetailsStitcherCoverClosed                       PrinterStatusDetails = "stitcherCoverClosed"
	PrinterStatusDetailsStitcherCoverOpen                         PrinterStatusDetails = "stitcherCoverOpen"
	PrinterStatusDetailsStitcherEmpty                             PrinterStatusDetails = "stitcherEmpty"
	PrinterStatusDetailsStitcherFull                              PrinterStatusDetails = "stitcherFull"
	PrinterStatusDetailsStitcherInterlockClosed                   PrinterStatusDetails = "stitcherInterlockClosed"
	PrinterStatusDetailsStitcherInterlockOpen                     PrinterStatusDetails = "stitcherInterlockOpen"
	PrinterStatusDetailsStitcherJam                               PrinterStatusDetails = "stitcherJam"
	PrinterStatusDetailsStitcherLifeAlmostOver                    PrinterStatusDetails = "stitcherLifeAlmostOver"
	PrinterStatusDetailsStitcherLifeOver                          PrinterStatusDetails = "stitcherLifeOver"
	PrinterStatusDetailsStitcherMemoryExhausted                   PrinterStatusDetails = "stitcherMemoryExhausted"
	PrinterStatusDetailsStitcherMissing                           PrinterStatusDetails = "stitcherMissing"
	PrinterStatusDetailsStitcherMotorFailure                      PrinterStatusDetails = "stitcherMotorFailure"
	PrinterStatusDetailsStitcherNearLimit                         PrinterStatusDetails = "stitcherNearLimit"
	PrinterStatusDetailsStitcherOffline                           PrinterStatusDetails = "stitcherOffline"
	PrinterStatusDetailsStitcherOpened                            PrinterStatusDetails = "stitcherOpened"
	PrinterStatusDetailsStitcherOverTemperature                   PrinterStatusDetails = "stitcherOverTemperature"
	PrinterStatusDetailsStitcherPowerSaver                        PrinterStatusDetails = "stitcherPowerSaver"
	PrinterStatusDetailsStitcherRecoverableFailure                PrinterStatusDetails = "stitcherRecoverableFailure"
	PrinterStatusDetailsStitcherRecoverableStorage                PrinterStatusDetails = "stitcherRecoverableStorage"
	PrinterStatusDetailsStitcherRemoved                           PrinterStatusDetails = "stitcherRemoved"
	PrinterStatusDetailsStitcherResourceAdded                     PrinterStatusDetails = "stitcherResourceAdded"
	PrinterStatusDetailsStitcherResourceRemoved                   PrinterStatusDetails = "stitcherResourceRemoved"
	PrinterStatusDetailsStitcherThermistorFailure                 PrinterStatusDetails = "stitcherThermistorFailure"
	PrinterStatusDetailsStitcherTimingFailure                     PrinterStatusDetails = "stitcherTimingFailure"
	PrinterStatusDetailsStitcherTurnedOff                         PrinterStatusDetails = "stitcherTurnedOff"
	PrinterStatusDetailsStitcherTurnedOn                          PrinterStatusDetails = "stitcherTurnedOn"
	PrinterStatusDetailsStitcherUnderTemperature                  PrinterStatusDetails = "stitcherUnderTemperature"
	PrinterStatusDetailsStitcherUnrecoverableFailure              PrinterStatusDetails = "stitcherUnrecoverableFailure"
	PrinterStatusDetailsStitcherUnrecoverableStorageError         PrinterStatusDetails = "stitcherUnrecoverableStorageError"
	PrinterStatusDetailsStitcherWarmingUp                         PrinterStatusDetails = "stitcherWarmingUp"
	PrinterStatusDetailsStoppedPartially                          PrinterStatusDetails = "stoppedPartially"
	PrinterStatusDetailsStopping                                  PrinterStatusDetails = "stopping"
	PrinterStatusDetailsSubunitAdded                              PrinterStatusDetails = "subunitAdded"
	PrinterStatusDetailsSubunitAlmostEmpty                        PrinterStatusDetails = "subunitAlmostEmpty"
	PrinterStatusDetailsSubunitAlmostFull                         PrinterStatusDetails = "subunitAlmostFull"
	PrinterStatusDetailsSubunitAtLimit                            PrinterStatusDetails = "subunitAtLimit"
	PrinterStatusDetailsSubunitClosed                             PrinterStatusDetails = "subunitClosed"
	PrinterStatusDetailsSubunitCoolingDown                        PrinterStatusDetails = "subunitCoolingDown"
	PrinterStatusDetailsSubunitEmpty                              PrinterStatusDetails = "subunitEmpty"
	PrinterStatusDetailsSubunitFull                               PrinterStatusDetails = "subunitFull"
	PrinterStatusDetailsSubunitLifeAlmostOver                     PrinterStatusDetails = "subunitLifeAlmostOver"
	PrinterStatusDetailsSubunitLifeOver                           PrinterStatusDetails = "subunitLifeOver"
	PrinterStatusDetailsSubunitMemoryExhausted                    PrinterStatusDetails = "subunitMemoryExhausted"
	PrinterStatusDetailsSubunitMissing                            PrinterStatusDetails = "subunitMissing"
	PrinterStatusDetailsSubunitMotorFailure                       PrinterStatusDetails = "subunitMotorFailure"
	PrinterStatusDetailsSubunitNearLimit                          PrinterStatusDetails = "subunitNearLimit"
	PrinterStatusDetailsSubunitOffline                            PrinterStatusDetails = "subunitOffline"
	PrinterStatusDetailsSubunitOpened                             PrinterStatusDetails = "subunitOpened"
	PrinterStatusDetailsSubunitOverTemperature                    PrinterStatusDetails = "subunitOverTemperature"
	PrinterStatusDetailsSubunitPowerSaver                         PrinterStatusDetails = "subunitPowerSaver"
	PrinterStatusDetailsSubunitRecoverableFailure                 PrinterStatusDetails = "subunitRecoverableFailure"
	PrinterStatusDetailsSubunitRecoverableStorage                 PrinterStatusDetails = "subunitRecoverableStorage"
	PrinterStatusDetailsSubunitRemoved                            PrinterStatusDetails = "subunitRemoved"
	PrinterStatusDetailsSubunitResourceAdded                      PrinterStatusDetails = "subunitResourceAdded"
	PrinterStatusDetailsSubunitResourceRemoved                    PrinterStatusDetails = "subunitResourceRemoved"
	PrinterStatusDetailsSubunitThermistorFailure                  PrinterStatusDetails = "subunitThermistorFailure"
	PrinterStatusDetailsSubunitTimingFailure                      PrinterStatusDetails = "subunitTimingFailure"
	PrinterStatusDetailsSubunitTurnedOff                          PrinterStatusDetails = "subunitTurnedOff"
	PrinterStatusDetailsSubunitTurnedOn                           PrinterStatusDetails = "subunitTurnedOn"
	PrinterStatusDetailsSubunitUnderTemperature                   PrinterStatusDetails = "subunitUnderTemperature"
	PrinterStatusDetailsSubunitUnrecoverableFailure               PrinterStatusDetails = "subunitUnrecoverableFailure"
	PrinterStatusDetailsSubunitUnrecoverableStorage               PrinterStatusDetails = "subunitUnrecoverableStorage"
	PrinterStatusDetailsSubunitWarmingUp                          PrinterStatusDetails = "subunitWarmingUp"
	PrinterStatusDetailsSuspend                                   PrinterStatusDetails = "suspend"
	PrinterStatusDetailsTesting                                   PrinterStatusDetails = "testing"
	PrinterStatusDetailsTimedOut                                  PrinterStatusDetails = "timedOut"
	PrinterStatusDetailsTonerEmpty                                PrinterStatusDetails = "tonerEmpty"
	PrinterStatusDetailsTonerLow                                  PrinterStatusDetails = "tonerLow"
	PrinterStatusDetailsTrimmerAdded                              PrinterStatusDetails = "trimmerAdded"
	PrinterStatusDetailsTrimmerAlmostEmpty                        PrinterStatusDetails = "trimmerAlmostEmpty"
	PrinterStatusDetailsTrimmerAlmostFull                         PrinterStatusDetails = "trimmerAlmostFull"
	PrinterStatusDetailsTrimmerAtLimit                            PrinterStatusDetails = "trimmerAtLimit"
	PrinterStatusDetailsTrimmerClosed                             PrinterStatusDetails = "trimmerClosed"
	PrinterStatusDetailsTrimmerConfigurationChange                PrinterStatusDetails = "trimmerConfigurationChange"
	PrinterStatusDetailsTrimmerCoverClosed                        PrinterStatusDetails = "trimmerCoverClosed"
	PrinterStatusDetailsTrimmerCoverOpen                          PrinterStatusDetails = "trimmerCoverOpen"
	PrinterStatusDetailsTrimmerEmpty                              PrinterStatusDetails = "trimmerEmpty"
	PrinterStatusDetailsTrimmerFull                               PrinterStatusDetails = "trimmerFull"
	PrinterStatusDetailsTrimmerInterlockClosed                    PrinterStatusDetails = "trimmerInterlockClosed"
	PrinterStatusDetailsTrimmerInterlockOpen                      PrinterStatusDetails = "trimmerInterlockOpen"
	PrinterStatusDetailsTrimmerJam                                PrinterStatusDetails = "trimmerJam"
	PrinterStatusDetailsTrimmerLifeAlmostOver                     PrinterStatusDetails = "trimmerLifeAlmostOver"
	PrinterStatusDetailsTrimmerLifeOver                           PrinterStatusDetails = "trimmerLifeOver"
	PrinterStatusDetailsTrimmerMemoryExhausted                    PrinterStatusDetails = "trimmerMemoryExhausted"
	PrinterStatusDetailsTrimmerMissing                            PrinterStatusDetails = "trimmerMissing"
	PrinterStatusDetailsTrimmerMotorFailure                       PrinterStatusDetails = "trimmerMotorFailure"
	PrinterStatusDetailsTrimmerNearLimit                          PrinterStatusDetails = "trimmerNearLimit"
	PrinterStatusDetailsTrimmerOffline                            PrinterStatusDetails = "trimmerOffline"
	PrinterStatusDetailsTrimmerOpened                             PrinterStatusDetails = "trimmerOpened"
	PrinterStatusDetailsTrimmerOverTemperature                    PrinterStatusDetails = "trimmerOverTemperature"
	PrinterStatusDetailsTrimmerPowerSaver                         PrinterStatusDetails = "trimmerPowerSaver"
	PrinterStatusDetailsTrimmerRecoverableFailure                 PrinterStatusDetails = "trimmerRecoverableFailure"
	PrinterStatusDetailsTrimmerRecoverableStorage                 PrinterStatusDetails = "trimmerRecoverableStorage"
	PrinterStatusDetailsTrimmerRemoved                            PrinterStatusDetails = "trimmerRemoved"
	PrinterStatusDetailsTrimmerResourceAdded                      PrinterStatusDetails = "trimmerResourceAdded"
	PrinterStatusDetailsTrimmerResourceRemoved                    PrinterStatusDetails = "trimmerResourceRemoved"
	PrinterStatusDetailsTrimmerThermistorFailure                  PrinterStatusDetails = "trimmerThermistorFailure"
	PrinterStatusDetailsTrimmerTimingFailure                      PrinterStatusDetails = "trimmerTimingFailure"
	PrinterStatusDetailsTrimmerTurnedOff                          PrinterStatusDetails = "trimmerTurnedOff"
	PrinterStatusDetailsTrimmerTurnedOn                           PrinterStatusDetails = "trimmerTurnedOn"
	PrinterStatusDetailsTrimmerUnderTemperature                   PrinterStatusDetails = "trimmerUnderTemperature"
	PrinterStatusDetailsTrimmerUnrecoverableFailure               PrinterStatusDetails = "trimmerUnrecoverableFailure"
	PrinterStatusDetailsTrimmerUnrecoverableStorageError          PrinterStatusDetails = "trimmerUnrecoverableStorageError"
	PrinterStatusDetailsTrimmerWarmingUp                          PrinterStatusDetails = "trimmerWarmingUp"
	PrinterStatusDetailsUnknown                                   PrinterStatusDetails = "unknown"
	PrinterStatusDetailsWrapperAdded                              PrinterStatusDetails = "wrapperAdded"
	PrinterStatusDetailsWrapperAlmostEmpty                        PrinterStatusDetails = "wrapperAlmostEmpty"
	PrinterStatusDetailsWrapperAlmostFull                         PrinterStatusDetails = "wrapperAlmostFull"
	PrinterStatusDetailsWrapperAtLimit                            PrinterStatusDetails = "wrapperAtLimit"
	PrinterStatusDetailsWrapperClosed                             PrinterStatusDetails = "wrapperClosed"
	PrinterStatusDetailsWrapperConfigurationChange                PrinterStatusDetails = "wrapperConfigurationChange"
	PrinterStatusDetailsWrapperCoverClosed                        PrinterStatusDetails = "wrapperCoverClosed"
	PrinterStatusDetailsWrapperCoverOpen                          PrinterStatusDetails = "wrapperCoverOpen"
	PrinterStatusDetailsWrapperEmpty                              PrinterStatusDetails = "wrapperEmpty"
	PrinterStatusDetailsWrapperFull                               PrinterStatusDetails = "wrapperFull"
	PrinterStatusDetailsWrapperInterlockClosed                    PrinterStatusDetails = "wrapperInterlockClosed"
	PrinterStatusDetailsWrapperInterlockOpen                      PrinterStatusDetails = "wrapperInterlockOpen"
	PrinterStatusDetailsWrapperJam                                PrinterStatusDetails = "wrapperJam"
	PrinterStatusDetailsWrapperLifeAlmostOver                     PrinterStatusDetails = "wrapperLifeAlmostOver"
	PrinterStatusDetailsWrapperLifeOver                           PrinterStatusDetails = "wrapperLifeOver"
	PrinterStatusDetailsWrapperMemoryExhausted                    PrinterStatusDetails = "wrapperMemoryExhausted"
	PrinterStatusDetailsWrapperMissing                            PrinterStatusDetails = "wrapperMissing"
	PrinterStatusDetailsWrapperMotorFailure                       PrinterStatusDetails = "wrapperMotorFailure"
	PrinterStatusDetailsWrapperNearLimit                          PrinterStatusDetails = "wrapperNearLimit"
	PrinterStatusDetailsWrapperOffline                            PrinterStatusDetails = "wrapperOffline"
	PrinterStatusDetailsWrapperOpened                             PrinterStatusDetails = "wrapperOpened"
	PrinterStatusDetailsWrapperOverTemperature                    PrinterStatusDetails = "wrapperOverTemperature"
	PrinterStatusDetailsWrapperPowerSaver                         PrinterStatusDetails = "wrapperPowerSaver"
	PrinterStatusDetailsWrapperRecoverableFailure                 PrinterStatusDetails = "wrapperRecoverableFailure"
	PrinterStatusDetailsWrapperRecoverableStorage                 PrinterStatusDetails = "wrapperRecoverableStorage"
	PrinterStatusDetailsWrapperRemoved                            PrinterStatusDetails = "wrapperRemoved"
	PrinterStatusDetailsWrapperResourceAdded                      PrinterStatusDetails = "wrapperResourceAdded"
	PrinterStatusDetailsWrapperResourceRemoved                    PrinterStatusDetails = "wrapperResourceRemoved"
	PrinterStatusDetailsWrapperThermistorFailure                  PrinterStatusDetails = "wrapperThermistorFailure"
	PrinterStatusDetailsWrapperTimingFailure                      PrinterStatusDetails = "wrapperTimingFailure"
	PrinterStatusDetailsWrapperTurnedOff                          PrinterStatusDetails = "wrapperTurnedOff"
	PrinterStatusDetailsWrapperTurnedOn                           PrinterStatusDetails = "wrapperTurnedOn"
	PrinterStatusDetailsWrapperUnderTemperature                   PrinterStatusDetails = "wrapperUnderTemperature"
	PrinterStatusDetailsWrapperUnrecoverableFailure               PrinterStatusDetails = "wrapperUnrecoverableFailure"
	PrinterStatusDetailsWrapperUnrecoverableStorageError          PrinterStatusDetails = "wrapperUnrecoverableStorageError"
	PrinterStatusDetailsWrapperWarmingUp                          PrinterStatusDetails = "wrapperWarmingUp"
)

func PossibleValuesForPrinterStatusDetails() []string {
	return []string{
		string(PrinterStatusDetailsAlertRemovalOfBinaryChangeEntry),
		string(PrinterStatusDetailsBanderAdded),
		string(PrinterStatusDetailsBanderAlmostEmpty),
		string(PrinterStatusDetailsBanderAlmostFull),
		string(PrinterStatusDetailsBanderAtLimit),
		string(PrinterStatusDetailsBanderClosed),
		string(PrinterStatusDetailsBanderConfigurationChange),
		string(PrinterStatusDetailsBanderCoverClosed),
		string(PrinterStatusDetailsBanderCoverOpen),
		string(PrinterStatusDetailsBanderEmpty),
		string(PrinterStatusDetailsBanderFull),
		string(PrinterStatusDetailsBanderInterlockClosed),
		string(PrinterStatusDetailsBanderInterlockOpen),
		string(PrinterStatusDetailsBanderJam),
		string(PrinterStatusDetailsBanderLifeAlmostOver),
		string(PrinterStatusDetailsBanderLifeOver),
		string(PrinterStatusDetailsBanderMemoryExhausted),
		string(PrinterStatusDetailsBanderMissing),
		string(PrinterStatusDetailsBanderMotorFailure),
		string(PrinterStatusDetailsBanderNearLimit),
		string(PrinterStatusDetailsBanderOffline),
		string(PrinterStatusDetailsBanderOpened),
		string(PrinterStatusDetailsBanderOverTemperature),
		string(PrinterStatusDetailsBanderPowerSaver),
		string(PrinterStatusDetailsBanderRecoverableFailure),
		string(PrinterStatusDetailsBanderRecoverableStorage),
		string(PrinterStatusDetailsBanderRemoved),
		string(PrinterStatusDetailsBanderResourceAdded),
		string(PrinterStatusDetailsBanderResourceRemoved),
		string(PrinterStatusDetailsBanderThermistorFailure),
		string(PrinterStatusDetailsBanderTimingFailure),
		string(PrinterStatusDetailsBanderTurnedOff),
		string(PrinterStatusDetailsBanderTurnedOn),
		string(PrinterStatusDetailsBanderUnderTemperature),
		string(PrinterStatusDetailsBanderUnrecoverableFailure),
		string(PrinterStatusDetailsBanderUnrecoverableStorageError),
		string(PrinterStatusDetailsBanderWarmingUp),
		string(PrinterStatusDetailsBinderAdded),
		string(PrinterStatusDetailsBinderAlmostEmpty),
		string(PrinterStatusDetailsBinderAlmostFull),
		string(PrinterStatusDetailsBinderAtLimit),
		string(PrinterStatusDetailsBinderClosed),
		string(PrinterStatusDetailsBinderConfigurationChange),
		string(PrinterStatusDetailsBinderCoverClosed),
		string(PrinterStatusDetailsBinderCoverOpen),
		string(PrinterStatusDetailsBinderEmpty),
		string(PrinterStatusDetailsBinderFull),
		string(PrinterStatusDetailsBinderInterlockClosed),
		string(PrinterStatusDetailsBinderInterlockOpen),
		string(PrinterStatusDetailsBinderJam),
		string(PrinterStatusDetailsBinderLifeAlmostOver),
		string(PrinterStatusDetailsBinderLifeOver),
		string(PrinterStatusDetailsBinderMemoryExhausted),
		string(PrinterStatusDetailsBinderMissing),
		string(PrinterStatusDetailsBinderMotorFailure),
		string(PrinterStatusDetailsBinderNearLimit),
		string(PrinterStatusDetailsBinderOffline),
		string(PrinterStatusDetailsBinderOpened),
		string(PrinterStatusDetailsBinderOverTemperature),
		string(PrinterStatusDetailsBinderPowerSaver),
		string(PrinterStatusDetailsBinderRecoverableFailure),
		string(PrinterStatusDetailsBinderRecoverableStorage),
		string(PrinterStatusDetailsBinderRemoved),
		string(PrinterStatusDetailsBinderResourceAdded),
		string(PrinterStatusDetailsBinderResourceRemoved),
		string(PrinterStatusDetailsBinderThermistorFailure),
		string(PrinterStatusDetailsBinderTimingFailure),
		string(PrinterStatusDetailsBinderTurnedOff),
		string(PrinterStatusDetailsBinderTurnedOn),
		string(PrinterStatusDetailsBinderUnderTemperature),
		string(PrinterStatusDetailsBinderUnrecoverableFailure),
		string(PrinterStatusDetailsBinderUnrecoverableStorageError),
		string(PrinterStatusDetailsBinderWarmingUp),
		string(PrinterStatusDetailsCameraFailure),
		string(PrinterStatusDetailsChamberCooling),
		string(PrinterStatusDetailsChamberFailure),
		string(PrinterStatusDetailsChamberHeating),
		string(PrinterStatusDetailsChamberTemperatureHigh),
		string(PrinterStatusDetailsChamberTemperatureLow),
		string(PrinterStatusDetailsCleanerLifeAlmostOver),
		string(PrinterStatusDetailsCleanerLifeOver),
		string(PrinterStatusDetailsConfigurationChange),
		string(PrinterStatusDetailsConnectingToDevice),
		string(PrinterStatusDetailsCoverOpen),
		string(PrinterStatusDetailsDeactivated),
		string(PrinterStatusDetailsDeleted),
		string(PrinterStatusDetailsDeveloperEmpty),
		string(PrinterStatusDetailsDeveloperLow),
		string(PrinterStatusDetailsDieCutterAdded),
		string(PrinterStatusDetailsDieCutterAlmostEmpty),
		string(PrinterStatusDetailsDieCutterAlmostFull),
		string(PrinterStatusDetailsDieCutterAtLimit),
		string(PrinterStatusDetailsDieCutterClosed),
		string(PrinterStatusDetailsDieCutterConfigurationChange),
		string(PrinterStatusDetailsDieCutterCoverClosed),
		string(PrinterStatusDetailsDieCutterCoverOpen),
		string(PrinterStatusDetailsDieCutterEmpty),
		string(PrinterStatusDetailsDieCutterFull),
		string(PrinterStatusDetailsDieCutterInterlockClosed),
		string(PrinterStatusDetailsDieCutterInterlockOpen),
		string(PrinterStatusDetailsDieCutterJam),
		string(PrinterStatusDetailsDieCutterLifeAlmostOver),
		string(PrinterStatusDetailsDieCutterLifeOver),
		string(PrinterStatusDetailsDieCutterMemoryExhausted),
		string(PrinterStatusDetailsDieCutterMissing),
		string(PrinterStatusDetailsDieCutterMotorFailure),
		string(PrinterStatusDetailsDieCutterNearLimit),
		string(PrinterStatusDetailsDieCutterOffline),
		string(PrinterStatusDetailsDieCutterOpened),
		string(PrinterStatusDetailsDieCutterOverTemperature),
		string(PrinterStatusDetailsDieCutterPowerSaver),
		string(PrinterStatusDetailsDieCutterRecoverableFailure),
		string(PrinterStatusDetailsDieCutterRecoverableStorage),
		string(PrinterStatusDetailsDieCutterRemoved),
		string(PrinterStatusDetailsDieCutterResourceAdded),
		string(PrinterStatusDetailsDieCutterResourceRemoved),
		string(PrinterStatusDetailsDieCutterThermistorFailure),
		string(PrinterStatusDetailsDieCutterTimingFailure),
		string(PrinterStatusDetailsDieCutterTurnedOff),
		string(PrinterStatusDetailsDieCutterTurnedOn),
		string(PrinterStatusDetailsDieCutterUnderTemperature),
		string(PrinterStatusDetailsDieCutterUnrecoverableFailure),
		string(PrinterStatusDetailsDieCutterUnrecoverableStorageError),
		string(PrinterStatusDetailsDieCutterWarmingUp),
		string(PrinterStatusDetailsDoorOpen),
		string(PrinterStatusDetailsExtruderCooling),
		string(PrinterStatusDetailsExtruderFailure),
		string(PrinterStatusDetailsExtruderHeating),
		string(PrinterStatusDetailsExtruderJam),
		string(PrinterStatusDetailsExtruderTemperatureHigh),
		string(PrinterStatusDetailsExtruderTemperatureLow),
		string(PrinterStatusDetailsFanFailure),
		string(PrinterStatusDetailsFaxModemLifeAlmostOver),
		string(PrinterStatusDetailsFaxModemLifeOver),
		string(PrinterStatusDetailsFaxModemMissing),
		string(PrinterStatusDetailsFaxModemTurnedOff),
		string(PrinterStatusDetailsFaxModemTurnedOn),
		string(PrinterStatusDetailsFolderAdded),
		string(PrinterStatusDetailsFolderAlmostEmpty),
		string(PrinterStatusDetailsFolderAlmostFull),
		string(PrinterStatusDetailsFolderAtLimit),
		string(PrinterStatusDetailsFolderClosed),
		string(PrinterStatusDetailsFolderConfigurationChange),
		string(PrinterStatusDetailsFolderCoverClosed),
		string(PrinterStatusDetailsFolderCoverOpen),
		string(PrinterStatusDetailsFolderEmpty),
		string(PrinterStatusDetailsFolderFull),
		string(PrinterStatusDetailsFolderInterlockClosed),
		string(PrinterStatusDetailsFolderInterlockOpen),
		string(PrinterStatusDetailsFolderJam),
		string(PrinterStatusDetailsFolderLifeAlmostOver),
		string(PrinterStatusDetailsFolderLifeOver),
		string(PrinterStatusDetailsFolderMemoryExhausted),
		string(PrinterStatusDetailsFolderMissing),
		string(PrinterStatusDetailsFolderMotorFailure),
		string(PrinterStatusDetailsFolderNearLimit),
		string(PrinterStatusDetailsFolderOffline),
		string(PrinterStatusDetailsFolderOpened),
		string(PrinterStatusDetailsFolderOverTemperature),
		string(PrinterStatusDetailsFolderPowerSaver),
		string(PrinterStatusDetailsFolderRecoverableFailure),
		string(PrinterStatusDetailsFolderRecoverableStorage),
		string(PrinterStatusDetailsFolderRemoved),
		string(PrinterStatusDetailsFolderResourceAdded),
		string(PrinterStatusDetailsFolderResourceRemoved),
		string(PrinterStatusDetailsFolderThermistorFailure),
		string(PrinterStatusDetailsFolderTimingFailure),
		string(PrinterStatusDetailsFolderTurnedOff),
		string(PrinterStatusDetailsFolderTurnedOn),
		string(PrinterStatusDetailsFolderUnderTemperature),
		string(PrinterStatusDetailsFolderUnrecoverableFailure),
		string(PrinterStatusDetailsFolderUnrecoverableStorageError),
		string(PrinterStatusDetailsFolderWarmingUp),
		string(PrinterStatusDetailsFuserOverTemp),
		string(PrinterStatusDetailsFuserUnderTemp),
		string(PrinterStatusDetailsHibernate),
		string(PrinterStatusDetailsHoldNewJobs),
		string(PrinterStatusDetailsIdentifyPrinterRequested),
		string(PrinterStatusDetailsImprinterAdded),
		string(PrinterStatusDetailsImprinterAlmostEmpty),
		string(PrinterStatusDetailsImprinterAlmostFull),
		string(PrinterStatusDetailsImprinterAtLimit),
		string(PrinterStatusDetailsImprinterClosed),
		string(PrinterStatusDetailsImprinterConfigurationChange),
		string(PrinterStatusDetailsImprinterCoverClosed),
		string(PrinterStatusDetailsImprinterCoverOpen),
		string(PrinterStatusDetailsImprinterEmpty),
		string(PrinterStatusDetailsImprinterFull),
		string(PrinterStatusDetailsImprinterInterlockClosed),
		string(PrinterStatusDetailsImprinterInterlockOpen),
		string(PrinterStatusDetailsImprinterJam),
		string(PrinterStatusDetailsImprinterLifeAlmostOver),
		string(PrinterStatusDetailsImprinterLifeOver),
		string(PrinterStatusDetailsImprinterMemoryExhausted),
		string(PrinterStatusDetailsImprinterMissing),
		string(PrinterStatusDetailsImprinterMotorFailure),
		string(PrinterStatusDetailsImprinterNearLimit),
		string(PrinterStatusDetailsImprinterOffline),
		string(PrinterStatusDetailsImprinterOpened),
		string(PrinterStatusDetailsImprinterOverTemperature),
		string(PrinterStatusDetailsImprinterPowerSaver),
		string(PrinterStatusDetailsImprinterRecoverableFailure),
		string(PrinterStatusDetailsImprinterRecoverableStorage),
		string(PrinterStatusDetailsImprinterRemoved),
		string(PrinterStatusDetailsImprinterResourceAdded),
		string(PrinterStatusDetailsImprinterResourceRemoved),
		string(PrinterStatusDetailsImprinterThermistorFailure),
		string(PrinterStatusDetailsImprinterTimingFailure),
		string(PrinterStatusDetailsImprinterTurnedOff),
		string(PrinterStatusDetailsImprinterTurnedOn),
		string(PrinterStatusDetailsImprinterUnderTemperature),
		string(PrinterStatusDetailsImprinterUnrecoverableFailure),
		string(PrinterStatusDetailsImprinterUnrecoverableStorageError),
		string(PrinterStatusDetailsImprinterWarmingUp),
		string(PrinterStatusDetailsInputCannotFeedSizeSelected),
		string(PrinterStatusDetailsInputManualInputRequest),
		string(PrinterStatusDetailsInputMediaColorChange),
		string(PrinterStatusDetailsInputMediaFormPartsChange),
		string(PrinterStatusDetailsInputMediaSizeChange),
		string(PrinterStatusDetailsInputMediaTrayFailure),
		string(PrinterStatusDetailsInputMediaTrayFeedError),
		string(PrinterStatusDetailsInputMediaTrayJam),
		string(PrinterStatusDetailsInputMediaTypeChange),
		string(PrinterStatusDetailsInputMediaWeightChange),
		string(PrinterStatusDetailsInputPickRollerFailure),
		string(PrinterStatusDetailsInputPickRollerLifeOver),
		string(PrinterStatusDetailsInputPickRollerLifeWarn),
		string(PrinterStatusDetailsInputPickRollerMissing),
		string(PrinterStatusDetailsInputTrayElevationFailure),
		string(PrinterStatusDetailsInputTrayMissing),
		string(PrinterStatusDetailsInputTrayPositionFailure),
		string(PrinterStatusDetailsInserterAdded),
		string(PrinterStatusDetailsInserterAlmostEmpty),
		string(PrinterStatusDetailsInserterAlmostFull),
		string(PrinterStatusDetailsInserterAtLimit),
		string(PrinterStatusDetailsInserterClosed),
		string(PrinterStatusDetailsInserterConfigurationChange),
		string(PrinterStatusDetailsInserterCoverClosed),
		string(PrinterStatusDetailsInserterCoverOpen),
		string(PrinterStatusDetailsInserterEmpty),
		string(PrinterStatusDetailsInserterFull),
		string(PrinterStatusDetailsInserterInterlockClosed),
		string(PrinterStatusDetailsInserterInterlockOpen),
		string(PrinterStatusDetailsInserterJam),
		string(PrinterStatusDetailsInserterLifeAlmostOver),
		string(PrinterStatusDetailsInserterLifeOver),
		string(PrinterStatusDetailsInserterMemoryExhausted),
		string(PrinterStatusDetailsInserterMissing),
		string(PrinterStatusDetailsInserterMotorFailure),
		string(PrinterStatusDetailsInserterNearLimit),
		string(PrinterStatusDetailsInserterOffline),
		string(PrinterStatusDetailsInserterOpened),
		string(PrinterStatusDetailsInserterOverTemperature),
		string(PrinterStatusDetailsInserterPowerSaver),
		string(PrinterStatusDetailsInserterRecoverableFailure),
		string(PrinterStatusDetailsInserterRecoverableStorage),
		string(PrinterStatusDetailsInserterRemoved),
		string(PrinterStatusDetailsInserterResourceAdded),
		string(PrinterStatusDetailsInserterResourceRemoved),
		string(PrinterStatusDetailsInserterThermistorFailure),
		string(PrinterStatusDetailsInserterTimingFailure),
		string(PrinterStatusDetailsInserterTurnedOff),
		string(PrinterStatusDetailsInserterTurnedOn),
		string(PrinterStatusDetailsInserterUnderTemperature),
		string(PrinterStatusDetailsInserterUnrecoverableFailure),
		string(PrinterStatusDetailsInserterUnrecoverableStorageError),
		string(PrinterStatusDetailsInserterWarmingUp),
		string(PrinterStatusDetailsInterlockClosed),
		string(PrinterStatusDetailsInterlockOpen),
		string(PrinterStatusDetailsInterpreterCartridgeAdded),
		string(PrinterStatusDetailsInterpreterCartridgeDeleted),
		string(PrinterStatusDetailsInterpreterComplexPageEncountered),
		string(PrinterStatusDetailsInterpreterMemoryDecrease),
		string(PrinterStatusDetailsInterpreterMemoryIncrease),
		string(PrinterStatusDetailsInterpreterResourceAdded),
		string(PrinterStatusDetailsInterpreterResourceDeleted),
		string(PrinterStatusDetailsInterpreterResourceUnavailable),
		string(PrinterStatusDetailsLampAtEol),
		string(PrinterStatusDetailsLampFailure),
		string(PrinterStatusDetailsLampNearEol),
		string(PrinterStatusDetailsLaserAtEol),
		string(PrinterStatusDetailsLaserFailure),
		string(PrinterStatusDetailsLaserNearEol),
		string(PrinterStatusDetailsMakeEnvelopeAdded),
		string(PrinterStatusDetailsMakeEnvelopeAlmostEmpty),
		string(PrinterStatusDetailsMakeEnvelopeAlmostFull),
		string(PrinterStatusDetailsMakeEnvelopeAtLimit),
		string(PrinterStatusDetailsMakeEnvelopeClosed),
		string(PrinterStatusDetailsMakeEnvelopeConfigurationChange),
		string(PrinterStatusDetailsMakeEnvelopeCoverClosed),
		string(PrinterStatusDetailsMakeEnvelopeCoverOpen),
		string(PrinterStatusDetailsMakeEnvelopeEmpty),
		string(PrinterStatusDetailsMakeEnvelopeFull),
		string(PrinterStatusDetailsMakeEnvelopeInterlockClosed),
		string(PrinterStatusDetailsMakeEnvelopeInterlockOpen),
		string(PrinterStatusDetailsMakeEnvelopeJam),
		string(PrinterStatusDetailsMakeEnvelopeLifeAlmostOver),
		string(PrinterStatusDetailsMakeEnvelopeLifeOver),
		string(PrinterStatusDetailsMakeEnvelopeMemoryExhausted),
		string(PrinterStatusDetailsMakeEnvelopeMissing),
		string(PrinterStatusDetailsMakeEnvelopeMotorFailure),
		string(PrinterStatusDetailsMakeEnvelopeNearLimit),
		string(PrinterStatusDetailsMakeEnvelopeOffline),
		string(PrinterStatusDetailsMakeEnvelopeOpened),
		string(PrinterStatusDetailsMakeEnvelopeOverTemperature),
		string(PrinterStatusDetailsMakeEnvelopePowerSaver),
		string(PrinterStatusDetailsMakeEnvelopeRecoverableFailure),
		string(PrinterStatusDetailsMakeEnvelopeRecoverableStorage),
		string(PrinterStatusDetailsMakeEnvelopeRemoved),
		string(PrinterStatusDetailsMakeEnvelopeResourceAdded),
		string(PrinterStatusDetailsMakeEnvelopeResourceRemoved),
		string(PrinterStatusDetailsMakeEnvelopeThermistorFailure),
		string(PrinterStatusDetailsMakeEnvelopeTimingFailure),
		string(PrinterStatusDetailsMakeEnvelopeTurnedOff),
		string(PrinterStatusDetailsMakeEnvelopeTurnedOn),
		string(PrinterStatusDetailsMakeEnvelopeUnderTemperature),
		string(PrinterStatusDetailsMakeEnvelopeUnrecoverableFailure),
		string(PrinterStatusDetailsMakeEnvelopeUnrecoverableStorageError),
		string(PrinterStatusDetailsMakeEnvelopeWarmingUp),
		string(PrinterStatusDetailsMarkerAdjustingPrintQuality),
		string(PrinterStatusDetailsMarkerCleanerMissing),
		string(PrinterStatusDetailsMarkerDeveloperAlmostEmpty),
		string(PrinterStatusDetailsMarkerDeveloperEmpty),
		string(PrinterStatusDetailsMarkerDeveloperMissing),
		string(PrinterStatusDetailsMarkerFuserMissing),
		string(PrinterStatusDetailsMarkerFuserThermistorFailure),
		string(PrinterStatusDetailsMarkerFuserTimingFailure),
		string(PrinterStatusDetailsMarkerInkAlmostEmpty),
		string(PrinterStatusDetailsMarkerInkEmpty),
		string(PrinterStatusDetailsMarkerInkMissing),
		string(PrinterStatusDetailsMarkerOpcMissing),
		string(PrinterStatusDetailsMarkerPrintRibbonAlmostEmpty),
		string(PrinterStatusDetailsMarkerPrintRibbonEmpty),
		string(PrinterStatusDetailsMarkerPrintRibbonMissing),
		string(PrinterStatusDetailsMarkerSupplyAlmostEmpty),
		string(PrinterStatusDetailsMarkerSupplyEmpty),
		string(PrinterStatusDetailsMarkerSupplyLow),
		string(PrinterStatusDetailsMarkerSupplyMissing),
		string(PrinterStatusDetailsMarkerTonerCartridgeMissing),
		string(PrinterStatusDetailsMarkerTonerMissing),
		string(PrinterStatusDetailsMarkerWasteAlmostFull),
		string(PrinterStatusDetailsMarkerWasteFull),
		string(PrinterStatusDetailsMarkerWasteInkReceptacleAlmostFull),
		string(PrinterStatusDetailsMarkerWasteInkReceptacleFull),
		string(PrinterStatusDetailsMarkerWasteInkReceptacleMissing),
		string(PrinterStatusDetailsMarkerWasteMissing),
		string(PrinterStatusDetailsMarkerWasteTonerReceptacleAlmostFull),
		string(PrinterStatusDetailsMarkerWasteTonerReceptacleFull),
		string(PrinterStatusDetailsMarkerWasteTonerReceptacleMissing),
		string(PrinterStatusDetailsMaterialEmpty),
		string(PrinterStatusDetailsMaterialLow),
		string(PrinterStatusDetailsMaterialNeeded),
		string(PrinterStatusDetailsMediaDrying),
		string(PrinterStatusDetailsMediaEmpty),
		string(PrinterStatusDetailsMediaJam),
		string(PrinterStatusDetailsMediaLow),
		string(PrinterStatusDetailsMediaNeeded),
		string(PrinterStatusDetailsMediaPathCannotDuplexMediaSelected),
		string(PrinterStatusDetailsMediaPathFailure),
		string(PrinterStatusDetailsMediaPathInputEmpty),
		string(PrinterStatusDetailsMediaPathInputFeedError),
		string(PrinterStatusDetailsMediaPathInputJam),
		string(PrinterStatusDetailsMediaPathInputRequest),
		string(PrinterStatusDetailsMediaPathJam),
		string(PrinterStatusDetailsMediaPathMediaTrayAlmostFull),
		string(PrinterStatusDetailsMediaPathMediaTrayFull),
		string(PrinterStatusDetailsMediaPathMediaTrayMissing),
		string(PrinterStatusDetailsMediaPathOutputFeedError),
		string(PrinterStatusDetailsMediaPathOutputFull),
		string(PrinterStatusDetailsMediaPathOutputJam),
		string(PrinterStatusDetailsMediaPathPickRollerFailure),
		string(PrinterStatusDetailsMediaPathPickRollerLifeOver),
		string(PrinterStatusDetailsMediaPathPickRollerLifeWarn),
		string(PrinterStatusDetailsMediaPathPickRollerMissing),
		string(PrinterStatusDetailsMotorFailure),
		string(PrinterStatusDetailsMovingToPaused),
		string(PrinterStatusDetailsNone),
		string(PrinterStatusDetailsOpticalPhotoConductorLifeOver),
		string(PrinterStatusDetailsOpticalPhotoConductorNearEndOfLife),
		string(PrinterStatusDetailsOther),
		string(PrinterStatusDetailsOutputAreaAlmostFull),
		string(PrinterStatusDetailsOutputAreaFull),
		string(PrinterStatusDetailsOutputMailboxSelectFailure),
		string(PrinterStatusDetailsOutputMediaTrayFailure),
		string(PrinterStatusDetailsOutputMediaTrayFeedError),
		string(PrinterStatusDetailsOutputMediaTrayJam),
		string(PrinterStatusDetailsOutputTrayMissing),
		string(PrinterStatusDetailsPaused),
		string(PrinterStatusDetailsPerforaterAdded),
		string(PrinterStatusDetailsPerforaterAlmostEmpty),
		string(PrinterStatusDetailsPerforaterAlmostFull),
		string(PrinterStatusDetailsPerforaterAtLimit),
		string(PrinterStatusDetailsPerforaterClosed),
		string(PrinterStatusDetailsPerforaterConfigurationChange),
		string(PrinterStatusDetailsPerforaterCoverClosed),
		string(PrinterStatusDetailsPerforaterCoverOpen),
		string(PrinterStatusDetailsPerforaterEmpty),
		string(PrinterStatusDetailsPerforaterFull),
		string(PrinterStatusDetailsPerforaterInterlockClosed),
		string(PrinterStatusDetailsPerforaterInterlockOpen),
		string(PrinterStatusDetailsPerforaterJam),
		string(PrinterStatusDetailsPerforaterLifeAlmostOver),
		string(PrinterStatusDetailsPerforaterLifeOver),
		string(PrinterStatusDetailsPerforaterMemoryExhausted),
		string(PrinterStatusDetailsPerforaterMissing),
		string(PrinterStatusDetailsPerforaterMotorFailure),
		string(PrinterStatusDetailsPerforaterNearLimit),
		string(PrinterStatusDetailsPerforaterOffline),
		string(PrinterStatusDetailsPerforaterOpened),
		string(PrinterStatusDetailsPerforaterOverTemperature),
		string(PrinterStatusDetailsPerforaterPowerSaver),
		string(PrinterStatusDetailsPerforaterRecoverableFailure),
		string(PrinterStatusDetailsPerforaterRecoverableStorage),
		string(PrinterStatusDetailsPerforaterRemoved),
		string(PrinterStatusDetailsPerforaterResourceAdded),
		string(PrinterStatusDetailsPerforaterResourceRemoved),
		string(PrinterStatusDetailsPerforaterThermistorFailure),
		string(PrinterStatusDetailsPerforaterTimingFailure),
		string(PrinterStatusDetailsPerforaterTurnedOff),
		string(PrinterStatusDetailsPerforaterTurnedOn),
		string(PrinterStatusDetailsPerforaterUnderTemperature),
		string(PrinterStatusDetailsPerforaterUnrecoverableFailure),
		string(PrinterStatusDetailsPerforaterUnrecoverableStorageError),
		string(PrinterStatusDetailsPerforaterWarmingUp),
		string(PrinterStatusDetailsPlatformCooling),
		string(PrinterStatusDetailsPlatformFailure),
		string(PrinterStatusDetailsPlatformHeating),
		string(PrinterStatusDetailsPlatformTemperatureHigh),
		string(PrinterStatusDetailsPlatformTemperatureLow),
		string(PrinterStatusDetailsPowerDown),
		string(PrinterStatusDetailsPowerUp),
		string(PrinterStatusDetailsPrinterManualReset),
		string(PrinterStatusDetailsPrinterNmsReset),
		string(PrinterStatusDetailsPrinterReadyToPrint),
		string(PrinterStatusDetailsPuncherAdded),
		string(PrinterStatusDetailsPuncherAlmostEmpty),
		string(PrinterStatusDetailsPuncherAlmostFull),
		string(PrinterStatusDetailsPuncherAtLimit),
		string(PrinterStatusDetailsPuncherClosed),
		string(PrinterStatusDetailsPuncherConfigurationChange),
		string(PrinterStatusDetailsPuncherCoverClosed),
		string(PrinterStatusDetailsPuncherCoverOpen),
		string(PrinterStatusDetailsPuncherEmpty),
		string(PrinterStatusDetailsPuncherFull),
		string(PrinterStatusDetailsPuncherInterlockClosed),
		string(PrinterStatusDetailsPuncherInterlockOpen),
		string(PrinterStatusDetailsPuncherJam),
		string(PrinterStatusDetailsPuncherLifeAlmostOver),
		string(PrinterStatusDetailsPuncherLifeOver),
		string(PrinterStatusDetailsPuncherMemoryExhausted),
		string(PrinterStatusDetailsPuncherMissing),
		string(PrinterStatusDetailsPuncherMotorFailure),
		string(PrinterStatusDetailsPuncherNearLimit),
		string(PrinterStatusDetailsPuncherOffline),
		string(PrinterStatusDetailsPuncherOpened),
		string(PrinterStatusDetailsPuncherOverTemperature),
		string(PrinterStatusDetailsPuncherPowerSaver),
		string(PrinterStatusDetailsPuncherRecoverableFailure),
		string(PrinterStatusDetailsPuncherRecoverableStorage),
		string(PrinterStatusDetailsPuncherRemoved),
		string(PrinterStatusDetailsPuncherResourceAdded),
		string(PrinterStatusDetailsPuncherResourceRemoved),
		string(PrinterStatusDetailsPuncherThermistorFailure),
		string(PrinterStatusDetailsPuncherTimingFailure),
		string(PrinterStatusDetailsPuncherTurnedOff),
		string(PrinterStatusDetailsPuncherTurnedOn),
		string(PrinterStatusDetailsPuncherUnderTemperature),
		string(PrinterStatusDetailsPuncherUnrecoverableFailure),
		string(PrinterStatusDetailsPuncherUnrecoverableStorageError),
		string(PrinterStatusDetailsPuncherWarmingUp),
		string(PrinterStatusDetailsResuming),
		string(PrinterStatusDetailsScanMediaPathFailure),
		string(PrinterStatusDetailsScanMediaPathInputEmpty),
		string(PrinterStatusDetailsScanMediaPathInputFeedError),
		string(PrinterStatusDetailsScanMediaPathInputJam),
		string(PrinterStatusDetailsScanMediaPathInputRequest),
		string(PrinterStatusDetailsScanMediaPathJam),
		string(PrinterStatusDetailsScanMediaPathOutputFeedError),
		string(PrinterStatusDetailsScanMediaPathOutputFull),
		string(PrinterStatusDetailsScanMediaPathOutputJam),
		string(PrinterStatusDetailsScanMediaPathPickRollerFailure),
		string(PrinterStatusDetailsScanMediaPathPickRollerLifeOver),
		string(PrinterStatusDetailsScanMediaPathPickRollerLifeWarn),
		string(PrinterStatusDetailsScanMediaPathPickRollerMissing),
		string(PrinterStatusDetailsScanMediaPathTrayAlmostFull),
		string(PrinterStatusDetailsScanMediaPathTrayFull),
		string(PrinterStatusDetailsScanMediaPathTrayMissing),
		string(PrinterStatusDetailsScannerLightFailure),
		string(PrinterStatusDetailsScannerLightLifeAlmostOver),
		string(PrinterStatusDetailsScannerLightLifeOver),
		string(PrinterStatusDetailsScannerLightMissing),
		string(PrinterStatusDetailsScannerSensorFailure),
		string(PrinterStatusDetailsScannerSensorLifeAlmostOver),
		string(PrinterStatusDetailsScannerSensorLifeOver),
		string(PrinterStatusDetailsScannerSensorMissing),
		string(PrinterStatusDetailsSeparationCutterAdded),
		string(PrinterStatusDetailsSeparationCutterAlmostEmpty),
		string(PrinterStatusDetailsSeparationCutterAlmostFull),
		string(PrinterStatusDetailsSeparationCutterAtLimit),
		string(PrinterStatusDetailsSeparationCutterClosed),
		string(PrinterStatusDetailsSeparationCutterConfigurationChange),
		string(PrinterStatusDetailsSeparationCutterCoverClosed),
		string(PrinterStatusDetailsSeparationCutterCoverOpen),
		string(PrinterStatusDetailsSeparationCutterEmpty),
		string(PrinterStatusDetailsSeparationCutterFull),
		string(PrinterStatusDetailsSeparationCutterInterlockClosed),
		string(PrinterStatusDetailsSeparationCutterInterlockOpen),
		string(PrinterStatusDetailsSeparationCutterJam),
		string(PrinterStatusDetailsSeparationCutterLifeAlmostOver),
		string(PrinterStatusDetailsSeparationCutterLifeOver),
		string(PrinterStatusDetailsSeparationCutterMemoryExhausted),
		string(PrinterStatusDetailsSeparationCutterMissing),
		string(PrinterStatusDetailsSeparationCutterMotorFailure),
		string(PrinterStatusDetailsSeparationCutterNearLimit),
		string(PrinterStatusDetailsSeparationCutterOffline),
		string(PrinterStatusDetailsSeparationCutterOpened),
		string(PrinterStatusDetailsSeparationCutterOverTemperature),
		string(PrinterStatusDetailsSeparationCutterPowerSaver),
		string(PrinterStatusDetailsSeparationCutterRecoverableFailure),
		string(PrinterStatusDetailsSeparationCutterRecoverableStorage),
		string(PrinterStatusDetailsSeparationCutterRemoved),
		string(PrinterStatusDetailsSeparationCutterResourceAdded),
		string(PrinterStatusDetailsSeparationCutterResourceRemoved),
		string(PrinterStatusDetailsSeparationCutterThermistorFailure),
		string(PrinterStatusDetailsSeparationCutterTimingFailure),
		string(PrinterStatusDetailsSeparationCutterTurnedOff),
		string(PrinterStatusDetailsSeparationCutterTurnedOn),
		string(PrinterStatusDetailsSeparationCutterUnderTemperature),
		string(PrinterStatusDetailsSeparationCutterUnrecoverableFailure),
		string(PrinterStatusDetailsSeparationCutterUnrecoverableStorageError),
		string(PrinterStatusDetailsSeparationCutterWarmingUp),
		string(PrinterStatusDetailsSheetRotatorAdded),
		string(PrinterStatusDetailsSheetRotatorAlmostEmpty),
		string(PrinterStatusDetailsSheetRotatorAlmostFull),
		string(PrinterStatusDetailsSheetRotatorAtLimit),
		string(PrinterStatusDetailsSheetRotatorClosed),
		string(PrinterStatusDetailsSheetRotatorConfigurationChange),
		string(PrinterStatusDetailsSheetRotatorCoverClosed),
		string(PrinterStatusDetailsSheetRotatorCoverOpen),
		string(PrinterStatusDetailsSheetRotatorEmpty),
		string(PrinterStatusDetailsSheetRotatorFull),
		string(PrinterStatusDetailsSheetRotatorInterlockClosed),
		string(PrinterStatusDetailsSheetRotatorInterlockOpen),
		string(PrinterStatusDetailsSheetRotatorJam),
		string(PrinterStatusDetailsSheetRotatorLifeAlmostOver),
		string(PrinterStatusDetailsSheetRotatorLifeOver),
		string(PrinterStatusDetailsSheetRotatorMemoryExhausted),
		string(PrinterStatusDetailsSheetRotatorMissing),
		string(PrinterStatusDetailsSheetRotatorMotorFailure),
		string(PrinterStatusDetailsSheetRotatorNearLimit),
		string(PrinterStatusDetailsSheetRotatorOffline),
		string(PrinterStatusDetailsSheetRotatorOpened),
		string(PrinterStatusDetailsSheetRotatorOverTemperature),
		string(PrinterStatusDetailsSheetRotatorPowerSaver),
		string(PrinterStatusDetailsSheetRotatorRecoverableFailure),
		string(PrinterStatusDetailsSheetRotatorRecoverableStorage),
		string(PrinterStatusDetailsSheetRotatorRemoved),
		string(PrinterStatusDetailsSheetRotatorResourceAdded),
		string(PrinterStatusDetailsSheetRotatorResourceRemoved),
		string(PrinterStatusDetailsSheetRotatorThermistorFailure),
		string(PrinterStatusDetailsSheetRotatorTimingFailure),
		string(PrinterStatusDetailsSheetRotatorTurnedOff),
		string(PrinterStatusDetailsSheetRotatorTurnedOn),
		string(PrinterStatusDetailsSheetRotatorUnderTemperature),
		string(PrinterStatusDetailsSheetRotatorUnrecoverableFailure),
		string(PrinterStatusDetailsSheetRotatorUnrecoverableStorageError),
		string(PrinterStatusDetailsSheetRotatorWarmingUp),
		string(PrinterStatusDetailsShutdown),
		string(PrinterStatusDetailsSlitterAdded),
		string(PrinterStatusDetailsSlitterAlmostEmpty),
		string(PrinterStatusDetailsSlitterAlmostFull),
		string(PrinterStatusDetailsSlitterAtLimit),
		string(PrinterStatusDetailsSlitterClosed),
		string(PrinterStatusDetailsSlitterConfigurationChange),
		string(PrinterStatusDetailsSlitterCoverClosed),
		string(PrinterStatusDetailsSlitterCoverOpen),
		string(PrinterStatusDetailsSlitterEmpty),
		string(PrinterStatusDetailsSlitterFull),
		string(PrinterStatusDetailsSlitterInterlockClosed),
		string(PrinterStatusDetailsSlitterInterlockOpen),
		string(PrinterStatusDetailsSlitterJam),
		string(PrinterStatusDetailsSlitterLifeAlmostOver),
		string(PrinterStatusDetailsSlitterLifeOver),
		string(PrinterStatusDetailsSlitterMemoryExhausted),
		string(PrinterStatusDetailsSlitterMissing),
		string(PrinterStatusDetailsSlitterMotorFailure),
		string(PrinterStatusDetailsSlitterNearLimit),
		string(PrinterStatusDetailsSlitterOffline),
		string(PrinterStatusDetailsSlitterOpened),
		string(PrinterStatusDetailsSlitterOverTemperature),
		string(PrinterStatusDetailsSlitterPowerSaver),
		string(PrinterStatusDetailsSlitterRecoverableFailure),
		string(PrinterStatusDetailsSlitterRecoverableStorage),
		string(PrinterStatusDetailsSlitterRemoved),
		string(PrinterStatusDetailsSlitterResourceAdded),
		string(PrinterStatusDetailsSlitterResourceRemoved),
		string(PrinterStatusDetailsSlitterThermistorFailure),
		string(PrinterStatusDetailsSlitterTimingFailure),
		string(PrinterStatusDetailsSlitterTurnedOff),
		string(PrinterStatusDetailsSlitterTurnedOn),
		string(PrinterStatusDetailsSlitterUnderTemperature),
		string(PrinterStatusDetailsSlitterUnrecoverableFailure),
		string(PrinterStatusDetailsSlitterUnrecoverableStorageError),
		string(PrinterStatusDetailsSlitterWarmingUp),
		string(PrinterStatusDetailsSpoolAreaFull),
		string(PrinterStatusDetailsStackerAdded),
		string(PrinterStatusDetailsStackerAlmostEmpty),
		string(PrinterStatusDetailsStackerAlmostFull),
		string(PrinterStatusDetailsStackerAtLimit),
		string(PrinterStatusDetailsStackerClosed),
		string(PrinterStatusDetailsStackerConfigurationChange),
		string(PrinterStatusDetailsStackerCoverClosed),
		string(PrinterStatusDetailsStackerCoverOpen),
		string(PrinterStatusDetailsStackerEmpty),
		string(PrinterStatusDetailsStackerFull),
		string(PrinterStatusDetailsStackerInterlockClosed),
		string(PrinterStatusDetailsStackerInterlockOpen),
		string(PrinterStatusDetailsStackerJam),
		string(PrinterStatusDetailsStackerLifeAlmostOver),
		string(PrinterStatusDetailsStackerLifeOver),
		string(PrinterStatusDetailsStackerMemoryExhausted),
		string(PrinterStatusDetailsStackerMissing),
		string(PrinterStatusDetailsStackerMotorFailure),
		string(PrinterStatusDetailsStackerNearLimit),
		string(PrinterStatusDetailsStackerOffline),
		string(PrinterStatusDetailsStackerOpened),
		string(PrinterStatusDetailsStackerOverTemperature),
		string(PrinterStatusDetailsStackerPowerSaver),
		string(PrinterStatusDetailsStackerRecoverableFailure),
		string(PrinterStatusDetailsStackerRecoverableStorage),
		string(PrinterStatusDetailsStackerRemoved),
		string(PrinterStatusDetailsStackerResourceAdded),
		string(PrinterStatusDetailsStackerResourceRemoved),
		string(PrinterStatusDetailsStackerThermistorFailure),
		string(PrinterStatusDetailsStackerTimingFailure),
		string(PrinterStatusDetailsStackerTurnedOff),
		string(PrinterStatusDetailsStackerTurnedOn),
		string(PrinterStatusDetailsStackerUnderTemperature),
		string(PrinterStatusDetailsStackerUnrecoverableFailure),
		string(PrinterStatusDetailsStackerUnrecoverableStorageError),
		string(PrinterStatusDetailsStackerWarmingUp),
		string(PrinterStatusDetailsStandby),
		string(PrinterStatusDetailsStaplerAdded),
		string(PrinterStatusDetailsStaplerAlmostEmpty),
		string(PrinterStatusDetailsStaplerAlmostFull),
		string(PrinterStatusDetailsStaplerAtLimit),
		string(PrinterStatusDetailsStaplerClosed),
		string(PrinterStatusDetailsStaplerConfigurationChange),
		string(PrinterStatusDetailsStaplerCoverClosed),
		string(PrinterStatusDetailsStaplerCoverOpen),
		string(PrinterStatusDetailsStaplerEmpty),
		string(PrinterStatusDetailsStaplerFull),
		string(PrinterStatusDetailsStaplerInterlockClosed),
		string(PrinterStatusDetailsStaplerInterlockOpen),
		string(PrinterStatusDetailsStaplerJam),
		string(PrinterStatusDetailsStaplerLifeAlmostOver),
		string(PrinterStatusDetailsStaplerLifeOver),
		string(PrinterStatusDetailsStaplerMemoryExhausted),
		string(PrinterStatusDetailsStaplerMissing),
		string(PrinterStatusDetailsStaplerMotorFailure),
		string(PrinterStatusDetailsStaplerNearLimit),
		string(PrinterStatusDetailsStaplerOffline),
		string(PrinterStatusDetailsStaplerOpened),
		string(PrinterStatusDetailsStaplerOverTemperature),
		string(PrinterStatusDetailsStaplerPowerSaver),
		string(PrinterStatusDetailsStaplerRecoverableFailure),
		string(PrinterStatusDetailsStaplerRecoverableStorage),
		string(PrinterStatusDetailsStaplerRemoved),
		string(PrinterStatusDetailsStaplerResourceAdded),
		string(PrinterStatusDetailsStaplerResourceRemoved),
		string(PrinterStatusDetailsStaplerThermistorFailure),
		string(PrinterStatusDetailsStaplerTimingFailure),
		string(PrinterStatusDetailsStaplerTurnedOff),
		string(PrinterStatusDetailsStaplerTurnedOn),
		string(PrinterStatusDetailsStaplerUnderTemperature),
		string(PrinterStatusDetailsStaplerUnrecoverableFailure),
		string(PrinterStatusDetailsStaplerUnrecoverableStorageError),
		string(PrinterStatusDetailsStaplerWarmingUp),
		string(PrinterStatusDetailsStitcherAdded),
		string(PrinterStatusDetailsStitcherAlmostEmpty),
		string(PrinterStatusDetailsStitcherAlmostFull),
		string(PrinterStatusDetailsStitcherAtLimit),
		string(PrinterStatusDetailsStitcherClosed),
		string(PrinterStatusDetailsStitcherConfigurationChange),
		string(PrinterStatusDetailsStitcherCoverClosed),
		string(PrinterStatusDetailsStitcherCoverOpen),
		string(PrinterStatusDetailsStitcherEmpty),
		string(PrinterStatusDetailsStitcherFull),
		string(PrinterStatusDetailsStitcherInterlockClosed),
		string(PrinterStatusDetailsStitcherInterlockOpen),
		string(PrinterStatusDetailsStitcherJam),
		string(PrinterStatusDetailsStitcherLifeAlmostOver),
		string(PrinterStatusDetailsStitcherLifeOver),
		string(PrinterStatusDetailsStitcherMemoryExhausted),
		string(PrinterStatusDetailsStitcherMissing),
		string(PrinterStatusDetailsStitcherMotorFailure),
		string(PrinterStatusDetailsStitcherNearLimit),
		string(PrinterStatusDetailsStitcherOffline),
		string(PrinterStatusDetailsStitcherOpened),
		string(PrinterStatusDetailsStitcherOverTemperature),
		string(PrinterStatusDetailsStitcherPowerSaver),
		string(PrinterStatusDetailsStitcherRecoverableFailure),
		string(PrinterStatusDetailsStitcherRecoverableStorage),
		string(PrinterStatusDetailsStitcherRemoved),
		string(PrinterStatusDetailsStitcherResourceAdded),
		string(PrinterStatusDetailsStitcherResourceRemoved),
		string(PrinterStatusDetailsStitcherThermistorFailure),
		string(PrinterStatusDetailsStitcherTimingFailure),
		string(PrinterStatusDetailsStitcherTurnedOff),
		string(PrinterStatusDetailsStitcherTurnedOn),
		string(PrinterStatusDetailsStitcherUnderTemperature),
		string(PrinterStatusDetailsStitcherUnrecoverableFailure),
		string(PrinterStatusDetailsStitcherUnrecoverableStorageError),
		string(PrinterStatusDetailsStitcherWarmingUp),
		string(PrinterStatusDetailsStoppedPartially),
		string(PrinterStatusDetailsStopping),
		string(PrinterStatusDetailsSubunitAdded),
		string(PrinterStatusDetailsSubunitAlmostEmpty),
		string(PrinterStatusDetailsSubunitAlmostFull),
		string(PrinterStatusDetailsSubunitAtLimit),
		string(PrinterStatusDetailsSubunitClosed),
		string(PrinterStatusDetailsSubunitCoolingDown),
		string(PrinterStatusDetailsSubunitEmpty),
		string(PrinterStatusDetailsSubunitFull),
		string(PrinterStatusDetailsSubunitLifeAlmostOver),
		string(PrinterStatusDetailsSubunitLifeOver),
		string(PrinterStatusDetailsSubunitMemoryExhausted),
		string(PrinterStatusDetailsSubunitMissing),
		string(PrinterStatusDetailsSubunitMotorFailure),
		string(PrinterStatusDetailsSubunitNearLimit),
		string(PrinterStatusDetailsSubunitOffline),
		string(PrinterStatusDetailsSubunitOpened),
		string(PrinterStatusDetailsSubunitOverTemperature),
		string(PrinterStatusDetailsSubunitPowerSaver),
		string(PrinterStatusDetailsSubunitRecoverableFailure),
		string(PrinterStatusDetailsSubunitRecoverableStorage),
		string(PrinterStatusDetailsSubunitRemoved),
		string(PrinterStatusDetailsSubunitResourceAdded),
		string(PrinterStatusDetailsSubunitResourceRemoved),
		string(PrinterStatusDetailsSubunitThermistorFailure),
		string(PrinterStatusDetailsSubunitTimingFailure),
		string(PrinterStatusDetailsSubunitTurnedOff),
		string(PrinterStatusDetailsSubunitTurnedOn),
		string(PrinterStatusDetailsSubunitUnderTemperature),
		string(PrinterStatusDetailsSubunitUnrecoverableFailure),
		string(PrinterStatusDetailsSubunitUnrecoverableStorage),
		string(PrinterStatusDetailsSubunitWarmingUp),
		string(PrinterStatusDetailsSuspend),
		string(PrinterStatusDetailsTesting),
		string(PrinterStatusDetailsTimedOut),
		string(PrinterStatusDetailsTonerEmpty),
		string(PrinterStatusDetailsTonerLow),
		string(PrinterStatusDetailsTrimmerAdded),
		string(PrinterStatusDetailsTrimmerAlmostEmpty),
		string(PrinterStatusDetailsTrimmerAlmostFull),
		string(PrinterStatusDetailsTrimmerAtLimit),
		string(PrinterStatusDetailsTrimmerClosed),
		string(PrinterStatusDetailsTrimmerConfigurationChange),
		string(PrinterStatusDetailsTrimmerCoverClosed),
		string(PrinterStatusDetailsTrimmerCoverOpen),
		string(PrinterStatusDetailsTrimmerEmpty),
		string(PrinterStatusDetailsTrimmerFull),
		string(PrinterStatusDetailsTrimmerInterlockClosed),
		string(PrinterStatusDetailsTrimmerInterlockOpen),
		string(PrinterStatusDetailsTrimmerJam),
		string(PrinterStatusDetailsTrimmerLifeAlmostOver),
		string(PrinterStatusDetailsTrimmerLifeOver),
		string(PrinterStatusDetailsTrimmerMemoryExhausted),
		string(PrinterStatusDetailsTrimmerMissing),
		string(PrinterStatusDetailsTrimmerMotorFailure),
		string(PrinterStatusDetailsTrimmerNearLimit),
		string(PrinterStatusDetailsTrimmerOffline),
		string(PrinterStatusDetailsTrimmerOpened),
		string(PrinterStatusDetailsTrimmerOverTemperature),
		string(PrinterStatusDetailsTrimmerPowerSaver),
		string(PrinterStatusDetailsTrimmerRecoverableFailure),
		string(PrinterStatusDetailsTrimmerRecoverableStorage),
		string(PrinterStatusDetailsTrimmerRemoved),
		string(PrinterStatusDetailsTrimmerResourceAdded),
		string(PrinterStatusDetailsTrimmerResourceRemoved),
		string(PrinterStatusDetailsTrimmerThermistorFailure),
		string(PrinterStatusDetailsTrimmerTimingFailure),
		string(PrinterStatusDetailsTrimmerTurnedOff),
		string(PrinterStatusDetailsTrimmerTurnedOn),
		string(PrinterStatusDetailsTrimmerUnderTemperature),
		string(PrinterStatusDetailsTrimmerUnrecoverableFailure),
		string(PrinterStatusDetailsTrimmerUnrecoverableStorageError),
		string(PrinterStatusDetailsTrimmerWarmingUp),
		string(PrinterStatusDetailsUnknown),
		string(PrinterStatusDetailsWrapperAdded),
		string(PrinterStatusDetailsWrapperAlmostEmpty),
		string(PrinterStatusDetailsWrapperAlmostFull),
		string(PrinterStatusDetailsWrapperAtLimit),
		string(PrinterStatusDetailsWrapperClosed),
		string(PrinterStatusDetailsWrapperConfigurationChange),
		string(PrinterStatusDetailsWrapperCoverClosed),
		string(PrinterStatusDetailsWrapperCoverOpen),
		string(PrinterStatusDetailsWrapperEmpty),
		string(PrinterStatusDetailsWrapperFull),
		string(PrinterStatusDetailsWrapperInterlockClosed),
		string(PrinterStatusDetailsWrapperInterlockOpen),
		string(PrinterStatusDetailsWrapperJam),
		string(PrinterStatusDetailsWrapperLifeAlmostOver),
		string(PrinterStatusDetailsWrapperLifeOver),
		string(PrinterStatusDetailsWrapperMemoryExhausted),
		string(PrinterStatusDetailsWrapperMissing),
		string(PrinterStatusDetailsWrapperMotorFailure),
		string(PrinterStatusDetailsWrapperNearLimit),
		string(PrinterStatusDetailsWrapperOffline),
		string(PrinterStatusDetailsWrapperOpened),
		string(PrinterStatusDetailsWrapperOverTemperature),
		string(PrinterStatusDetailsWrapperPowerSaver),
		string(PrinterStatusDetailsWrapperRecoverableFailure),
		string(PrinterStatusDetailsWrapperRecoverableStorage),
		string(PrinterStatusDetailsWrapperRemoved),
		string(PrinterStatusDetailsWrapperResourceAdded),
		string(PrinterStatusDetailsWrapperResourceRemoved),
		string(PrinterStatusDetailsWrapperThermistorFailure),
		string(PrinterStatusDetailsWrapperTimingFailure),
		string(PrinterStatusDetailsWrapperTurnedOff),
		string(PrinterStatusDetailsWrapperTurnedOn),
		string(PrinterStatusDetailsWrapperUnderTemperature),
		string(PrinterStatusDetailsWrapperUnrecoverableFailure),
		string(PrinterStatusDetailsWrapperUnrecoverableStorageError),
		string(PrinterStatusDetailsWrapperWarmingUp),
	}
}

func (s *PrinterStatusDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusDetails(input string) (*PrinterStatusDetails, error) {
	vals := map[string]PrinterStatusDetails{
		"alertremovalofbinarychangeentry":           PrinterStatusDetailsAlertRemovalOfBinaryChangeEntry,
		"banderadded":                               PrinterStatusDetailsBanderAdded,
		"banderalmostempty":                         PrinterStatusDetailsBanderAlmostEmpty,
		"banderalmostfull":                          PrinterStatusDetailsBanderAlmostFull,
		"banderatlimit":                             PrinterStatusDetailsBanderAtLimit,
		"banderclosed":                              PrinterStatusDetailsBanderClosed,
		"banderconfigurationchange":                 PrinterStatusDetailsBanderConfigurationChange,
		"bandercoverclosed":                         PrinterStatusDetailsBanderCoverClosed,
		"bandercoveropen":                           PrinterStatusDetailsBanderCoverOpen,
		"banderempty":                               PrinterStatusDetailsBanderEmpty,
		"banderfull":                                PrinterStatusDetailsBanderFull,
		"banderinterlockclosed":                     PrinterStatusDetailsBanderInterlockClosed,
		"banderinterlockopen":                       PrinterStatusDetailsBanderInterlockOpen,
		"banderjam":                                 PrinterStatusDetailsBanderJam,
		"banderlifealmostover":                      PrinterStatusDetailsBanderLifeAlmostOver,
		"banderlifeover":                            PrinterStatusDetailsBanderLifeOver,
		"bandermemoryexhausted":                     PrinterStatusDetailsBanderMemoryExhausted,
		"bandermissing":                             PrinterStatusDetailsBanderMissing,
		"bandermotorfailure":                        PrinterStatusDetailsBanderMotorFailure,
		"bandernearlimit":                           PrinterStatusDetailsBanderNearLimit,
		"banderoffline":                             PrinterStatusDetailsBanderOffline,
		"banderopened":                              PrinterStatusDetailsBanderOpened,
		"banderovertemperature":                     PrinterStatusDetailsBanderOverTemperature,
		"banderpowersaver":                          PrinterStatusDetailsBanderPowerSaver,
		"banderrecoverablefailure":                  PrinterStatusDetailsBanderRecoverableFailure,
		"banderrecoverablestorage":                  PrinterStatusDetailsBanderRecoverableStorage,
		"banderremoved":                             PrinterStatusDetailsBanderRemoved,
		"banderresourceadded":                       PrinterStatusDetailsBanderResourceAdded,
		"banderresourceremoved":                     PrinterStatusDetailsBanderResourceRemoved,
		"banderthermistorfailure":                   PrinterStatusDetailsBanderThermistorFailure,
		"bandertimingfailure":                       PrinterStatusDetailsBanderTimingFailure,
		"banderturnedoff":                           PrinterStatusDetailsBanderTurnedOff,
		"banderturnedon":                            PrinterStatusDetailsBanderTurnedOn,
		"banderundertemperature":                    PrinterStatusDetailsBanderUnderTemperature,
		"banderunrecoverablefailure":                PrinterStatusDetailsBanderUnrecoverableFailure,
		"banderunrecoverablestorageerror":           PrinterStatusDetailsBanderUnrecoverableStorageError,
		"banderwarmingup":                           PrinterStatusDetailsBanderWarmingUp,
		"binderadded":                               PrinterStatusDetailsBinderAdded,
		"binderalmostempty":                         PrinterStatusDetailsBinderAlmostEmpty,
		"binderalmostfull":                          PrinterStatusDetailsBinderAlmostFull,
		"binderatlimit":                             PrinterStatusDetailsBinderAtLimit,
		"binderclosed":                              PrinterStatusDetailsBinderClosed,
		"binderconfigurationchange":                 PrinterStatusDetailsBinderConfigurationChange,
		"bindercoverclosed":                         PrinterStatusDetailsBinderCoverClosed,
		"bindercoveropen":                           PrinterStatusDetailsBinderCoverOpen,
		"binderempty":                               PrinterStatusDetailsBinderEmpty,
		"binderfull":                                PrinterStatusDetailsBinderFull,
		"binderinterlockclosed":                     PrinterStatusDetailsBinderInterlockClosed,
		"binderinterlockopen":                       PrinterStatusDetailsBinderInterlockOpen,
		"binderjam":                                 PrinterStatusDetailsBinderJam,
		"binderlifealmostover":                      PrinterStatusDetailsBinderLifeAlmostOver,
		"binderlifeover":                            PrinterStatusDetailsBinderLifeOver,
		"bindermemoryexhausted":                     PrinterStatusDetailsBinderMemoryExhausted,
		"bindermissing":                             PrinterStatusDetailsBinderMissing,
		"bindermotorfailure":                        PrinterStatusDetailsBinderMotorFailure,
		"bindernearlimit":                           PrinterStatusDetailsBinderNearLimit,
		"binderoffline":                             PrinterStatusDetailsBinderOffline,
		"binderopened":                              PrinterStatusDetailsBinderOpened,
		"binderovertemperature":                     PrinterStatusDetailsBinderOverTemperature,
		"binderpowersaver":                          PrinterStatusDetailsBinderPowerSaver,
		"binderrecoverablefailure":                  PrinterStatusDetailsBinderRecoverableFailure,
		"binderrecoverablestorage":                  PrinterStatusDetailsBinderRecoverableStorage,
		"binderremoved":                             PrinterStatusDetailsBinderRemoved,
		"binderresourceadded":                       PrinterStatusDetailsBinderResourceAdded,
		"binderresourceremoved":                     PrinterStatusDetailsBinderResourceRemoved,
		"binderthermistorfailure":                   PrinterStatusDetailsBinderThermistorFailure,
		"bindertimingfailure":                       PrinterStatusDetailsBinderTimingFailure,
		"binderturnedoff":                           PrinterStatusDetailsBinderTurnedOff,
		"binderturnedon":                            PrinterStatusDetailsBinderTurnedOn,
		"binderundertemperature":                    PrinterStatusDetailsBinderUnderTemperature,
		"binderunrecoverablefailure":                PrinterStatusDetailsBinderUnrecoverableFailure,
		"binderunrecoverablestorageerror":           PrinterStatusDetailsBinderUnrecoverableStorageError,
		"binderwarmingup":                           PrinterStatusDetailsBinderWarmingUp,
		"camerafailure":                             PrinterStatusDetailsCameraFailure,
		"chambercooling":                            PrinterStatusDetailsChamberCooling,
		"chamberfailure":                            PrinterStatusDetailsChamberFailure,
		"chamberheating":                            PrinterStatusDetailsChamberHeating,
		"chambertemperaturehigh":                    PrinterStatusDetailsChamberTemperatureHigh,
		"chambertemperaturelow":                     PrinterStatusDetailsChamberTemperatureLow,
		"cleanerlifealmostover":                     PrinterStatusDetailsCleanerLifeAlmostOver,
		"cleanerlifeover":                           PrinterStatusDetailsCleanerLifeOver,
		"configurationchange":                       PrinterStatusDetailsConfigurationChange,
		"connectingtodevice":                        PrinterStatusDetailsConnectingToDevice,
		"coveropen":                                 PrinterStatusDetailsCoverOpen,
		"deactivated":                               PrinterStatusDetailsDeactivated,
		"deleted":                                   PrinterStatusDetailsDeleted,
		"developerempty":                            PrinterStatusDetailsDeveloperEmpty,
		"developerlow":                              PrinterStatusDetailsDeveloperLow,
		"diecutteradded":                            PrinterStatusDetailsDieCutterAdded,
		"diecutteralmostempty":                      PrinterStatusDetailsDieCutterAlmostEmpty,
		"diecutteralmostfull":                       PrinterStatusDetailsDieCutterAlmostFull,
		"diecutteratlimit":                          PrinterStatusDetailsDieCutterAtLimit,
		"diecutterclosed":                           PrinterStatusDetailsDieCutterClosed,
		"diecutterconfigurationchange":              PrinterStatusDetailsDieCutterConfigurationChange,
		"diecuttercoverclosed":                      PrinterStatusDetailsDieCutterCoverClosed,
		"diecuttercoveropen":                        PrinterStatusDetailsDieCutterCoverOpen,
		"diecutterempty":                            PrinterStatusDetailsDieCutterEmpty,
		"diecutterfull":                             PrinterStatusDetailsDieCutterFull,
		"diecutterinterlockclosed":                  PrinterStatusDetailsDieCutterInterlockClosed,
		"diecutterinterlockopen":                    PrinterStatusDetailsDieCutterInterlockOpen,
		"diecutterjam":                              PrinterStatusDetailsDieCutterJam,
		"diecutterlifealmostover":                   PrinterStatusDetailsDieCutterLifeAlmostOver,
		"diecutterlifeover":                         PrinterStatusDetailsDieCutterLifeOver,
		"diecuttermemoryexhausted":                  PrinterStatusDetailsDieCutterMemoryExhausted,
		"diecuttermissing":                          PrinterStatusDetailsDieCutterMissing,
		"diecuttermotorfailure":                     PrinterStatusDetailsDieCutterMotorFailure,
		"diecutternearlimit":                        PrinterStatusDetailsDieCutterNearLimit,
		"diecutteroffline":                          PrinterStatusDetailsDieCutterOffline,
		"diecutteropened":                           PrinterStatusDetailsDieCutterOpened,
		"diecutterovertemperature":                  PrinterStatusDetailsDieCutterOverTemperature,
		"diecutterpowersaver":                       PrinterStatusDetailsDieCutterPowerSaver,
		"diecutterrecoverablefailure":               PrinterStatusDetailsDieCutterRecoverableFailure,
		"diecutterrecoverablestorage":               PrinterStatusDetailsDieCutterRecoverableStorage,
		"diecutterremoved":                          PrinterStatusDetailsDieCutterRemoved,
		"diecutterresourceadded":                    PrinterStatusDetailsDieCutterResourceAdded,
		"diecutterresourceremoved":                  PrinterStatusDetailsDieCutterResourceRemoved,
		"diecutterthermistorfailure":                PrinterStatusDetailsDieCutterThermistorFailure,
		"diecuttertimingfailure":                    PrinterStatusDetailsDieCutterTimingFailure,
		"diecutterturnedoff":                        PrinterStatusDetailsDieCutterTurnedOff,
		"diecutterturnedon":                         PrinterStatusDetailsDieCutterTurnedOn,
		"diecutterundertemperature":                 PrinterStatusDetailsDieCutterUnderTemperature,
		"diecutterunrecoverablefailure":             PrinterStatusDetailsDieCutterUnrecoverableFailure,
		"diecutterunrecoverablestorageerror":        PrinterStatusDetailsDieCutterUnrecoverableStorageError,
		"diecutterwarmingup":                        PrinterStatusDetailsDieCutterWarmingUp,
		"dooropen":                                  PrinterStatusDetailsDoorOpen,
		"extrudercooling":                           PrinterStatusDetailsExtruderCooling,
		"extruderfailure":                           PrinterStatusDetailsExtruderFailure,
		"extruderheating":                           PrinterStatusDetailsExtruderHeating,
		"extruderjam":                               PrinterStatusDetailsExtruderJam,
		"extrudertemperaturehigh":                   PrinterStatusDetailsExtruderTemperatureHigh,
		"extrudertemperaturelow":                    PrinterStatusDetailsExtruderTemperatureLow,
		"fanfailure":                                PrinterStatusDetailsFanFailure,
		"faxmodemlifealmostover":                    PrinterStatusDetailsFaxModemLifeAlmostOver,
		"faxmodemlifeover":                          PrinterStatusDetailsFaxModemLifeOver,
		"faxmodemmissing":                           PrinterStatusDetailsFaxModemMissing,
		"faxmodemturnedoff":                         PrinterStatusDetailsFaxModemTurnedOff,
		"faxmodemturnedon":                          PrinterStatusDetailsFaxModemTurnedOn,
		"folderadded":                               PrinterStatusDetailsFolderAdded,
		"folderalmostempty":                         PrinterStatusDetailsFolderAlmostEmpty,
		"folderalmostfull":                          PrinterStatusDetailsFolderAlmostFull,
		"folderatlimit":                             PrinterStatusDetailsFolderAtLimit,
		"folderclosed":                              PrinterStatusDetailsFolderClosed,
		"folderconfigurationchange":                 PrinterStatusDetailsFolderConfigurationChange,
		"foldercoverclosed":                         PrinterStatusDetailsFolderCoverClosed,
		"foldercoveropen":                           PrinterStatusDetailsFolderCoverOpen,
		"folderempty":                               PrinterStatusDetailsFolderEmpty,
		"folderfull":                                PrinterStatusDetailsFolderFull,
		"folderinterlockclosed":                     PrinterStatusDetailsFolderInterlockClosed,
		"folderinterlockopen":                       PrinterStatusDetailsFolderInterlockOpen,
		"folderjam":                                 PrinterStatusDetailsFolderJam,
		"folderlifealmostover":                      PrinterStatusDetailsFolderLifeAlmostOver,
		"folderlifeover":                            PrinterStatusDetailsFolderLifeOver,
		"foldermemoryexhausted":                     PrinterStatusDetailsFolderMemoryExhausted,
		"foldermissing":                             PrinterStatusDetailsFolderMissing,
		"foldermotorfailure":                        PrinterStatusDetailsFolderMotorFailure,
		"foldernearlimit":                           PrinterStatusDetailsFolderNearLimit,
		"folderoffline":                             PrinterStatusDetailsFolderOffline,
		"folderopened":                              PrinterStatusDetailsFolderOpened,
		"folderovertemperature":                     PrinterStatusDetailsFolderOverTemperature,
		"folderpowersaver":                          PrinterStatusDetailsFolderPowerSaver,
		"folderrecoverablefailure":                  PrinterStatusDetailsFolderRecoverableFailure,
		"folderrecoverablestorage":                  PrinterStatusDetailsFolderRecoverableStorage,
		"folderremoved":                             PrinterStatusDetailsFolderRemoved,
		"folderresourceadded":                       PrinterStatusDetailsFolderResourceAdded,
		"folderresourceremoved":                     PrinterStatusDetailsFolderResourceRemoved,
		"folderthermistorfailure":                   PrinterStatusDetailsFolderThermistorFailure,
		"foldertimingfailure":                       PrinterStatusDetailsFolderTimingFailure,
		"folderturnedoff":                           PrinterStatusDetailsFolderTurnedOff,
		"folderturnedon":                            PrinterStatusDetailsFolderTurnedOn,
		"folderundertemperature":                    PrinterStatusDetailsFolderUnderTemperature,
		"folderunrecoverablefailure":                PrinterStatusDetailsFolderUnrecoverableFailure,
		"folderunrecoverablestorageerror":           PrinterStatusDetailsFolderUnrecoverableStorageError,
		"folderwarmingup":                           PrinterStatusDetailsFolderWarmingUp,
		"fuserovertemp":                             PrinterStatusDetailsFuserOverTemp,
		"fuserundertemp":                            PrinterStatusDetailsFuserUnderTemp,
		"hibernate":                                 PrinterStatusDetailsHibernate,
		"holdnewjobs":                               PrinterStatusDetailsHoldNewJobs,
		"identifyprinterrequested":                  PrinterStatusDetailsIdentifyPrinterRequested,
		"imprinteradded":                            PrinterStatusDetailsImprinterAdded,
		"imprinteralmostempty":                      PrinterStatusDetailsImprinterAlmostEmpty,
		"imprinteralmostfull":                       PrinterStatusDetailsImprinterAlmostFull,
		"imprinteratlimit":                          PrinterStatusDetailsImprinterAtLimit,
		"imprinterclosed":                           PrinterStatusDetailsImprinterClosed,
		"imprinterconfigurationchange":              PrinterStatusDetailsImprinterConfigurationChange,
		"imprintercoverclosed":                      PrinterStatusDetailsImprinterCoverClosed,
		"imprintercoveropen":                        PrinterStatusDetailsImprinterCoverOpen,
		"imprinterempty":                            PrinterStatusDetailsImprinterEmpty,
		"imprinterfull":                             PrinterStatusDetailsImprinterFull,
		"imprinterinterlockclosed":                  PrinterStatusDetailsImprinterInterlockClosed,
		"imprinterinterlockopen":                    PrinterStatusDetailsImprinterInterlockOpen,
		"imprinterjam":                              PrinterStatusDetailsImprinterJam,
		"imprinterlifealmostover":                   PrinterStatusDetailsImprinterLifeAlmostOver,
		"imprinterlifeover":                         PrinterStatusDetailsImprinterLifeOver,
		"imprintermemoryexhausted":                  PrinterStatusDetailsImprinterMemoryExhausted,
		"imprintermissing":                          PrinterStatusDetailsImprinterMissing,
		"imprintermotorfailure":                     PrinterStatusDetailsImprinterMotorFailure,
		"imprinternearlimit":                        PrinterStatusDetailsImprinterNearLimit,
		"imprinteroffline":                          PrinterStatusDetailsImprinterOffline,
		"imprinteropened":                           PrinterStatusDetailsImprinterOpened,
		"imprinterovertemperature":                  PrinterStatusDetailsImprinterOverTemperature,
		"imprinterpowersaver":                       PrinterStatusDetailsImprinterPowerSaver,
		"imprinterrecoverablefailure":               PrinterStatusDetailsImprinterRecoverableFailure,
		"imprinterrecoverablestorage":               PrinterStatusDetailsImprinterRecoverableStorage,
		"imprinterremoved":                          PrinterStatusDetailsImprinterRemoved,
		"imprinterresourceadded":                    PrinterStatusDetailsImprinterResourceAdded,
		"imprinterresourceremoved":                  PrinterStatusDetailsImprinterResourceRemoved,
		"imprinterthermistorfailure":                PrinterStatusDetailsImprinterThermistorFailure,
		"imprintertimingfailure":                    PrinterStatusDetailsImprinterTimingFailure,
		"imprinterturnedoff":                        PrinterStatusDetailsImprinterTurnedOff,
		"imprinterturnedon":                         PrinterStatusDetailsImprinterTurnedOn,
		"imprinterundertemperature":                 PrinterStatusDetailsImprinterUnderTemperature,
		"imprinterunrecoverablefailure":             PrinterStatusDetailsImprinterUnrecoverableFailure,
		"imprinterunrecoverablestorageerror":        PrinterStatusDetailsImprinterUnrecoverableStorageError,
		"imprinterwarmingup":                        PrinterStatusDetailsImprinterWarmingUp,
		"inputcannotfeedsizeselected":               PrinterStatusDetailsInputCannotFeedSizeSelected,
		"inputmanualinputrequest":                   PrinterStatusDetailsInputManualInputRequest,
		"inputmediacolorchange":                     PrinterStatusDetailsInputMediaColorChange,
		"inputmediaformpartschange":                 PrinterStatusDetailsInputMediaFormPartsChange,
		"inputmediasizechange":                      PrinterStatusDetailsInputMediaSizeChange,
		"inputmediatrayfailure":                     PrinterStatusDetailsInputMediaTrayFailure,
		"inputmediatrayfeederror":                   PrinterStatusDetailsInputMediaTrayFeedError,
		"inputmediatrayjam":                         PrinterStatusDetailsInputMediaTrayJam,
		"inputmediatypechange":                      PrinterStatusDetailsInputMediaTypeChange,
		"inputmediaweightchange":                    PrinterStatusDetailsInputMediaWeightChange,
		"inputpickrollerfailure":                    PrinterStatusDetailsInputPickRollerFailure,
		"inputpickrollerlifeover":                   PrinterStatusDetailsInputPickRollerLifeOver,
		"inputpickrollerlifewarn":                   PrinterStatusDetailsInputPickRollerLifeWarn,
		"inputpickrollermissing":                    PrinterStatusDetailsInputPickRollerMissing,
		"inputtrayelevationfailure":                 PrinterStatusDetailsInputTrayElevationFailure,
		"inputtraymissing":                          PrinterStatusDetailsInputTrayMissing,
		"inputtraypositionfailure":                  PrinterStatusDetailsInputTrayPositionFailure,
		"inserteradded":                             PrinterStatusDetailsInserterAdded,
		"inserteralmostempty":                       PrinterStatusDetailsInserterAlmostEmpty,
		"inserteralmostfull":                        PrinterStatusDetailsInserterAlmostFull,
		"inserteratlimit":                           PrinterStatusDetailsInserterAtLimit,
		"inserterclosed":                            PrinterStatusDetailsInserterClosed,
		"inserterconfigurationchange":               PrinterStatusDetailsInserterConfigurationChange,
		"insertercoverclosed":                       PrinterStatusDetailsInserterCoverClosed,
		"insertercoveropen":                         PrinterStatusDetailsInserterCoverOpen,
		"inserterempty":                             PrinterStatusDetailsInserterEmpty,
		"inserterfull":                              PrinterStatusDetailsInserterFull,
		"inserterinterlockclosed":                   PrinterStatusDetailsInserterInterlockClosed,
		"inserterinterlockopen":                     PrinterStatusDetailsInserterInterlockOpen,
		"inserterjam":                               PrinterStatusDetailsInserterJam,
		"inserterlifealmostover":                    PrinterStatusDetailsInserterLifeAlmostOver,
		"inserterlifeover":                          PrinterStatusDetailsInserterLifeOver,
		"insertermemoryexhausted":                   PrinterStatusDetailsInserterMemoryExhausted,
		"insertermissing":                           PrinterStatusDetailsInserterMissing,
		"insertermotorfailure":                      PrinterStatusDetailsInserterMotorFailure,
		"inserternearlimit":                         PrinterStatusDetailsInserterNearLimit,
		"inserteroffline":                           PrinterStatusDetailsInserterOffline,
		"inserteropened":                            PrinterStatusDetailsInserterOpened,
		"inserterovertemperature":                   PrinterStatusDetailsInserterOverTemperature,
		"inserterpowersaver":                        PrinterStatusDetailsInserterPowerSaver,
		"inserterrecoverablefailure":                PrinterStatusDetailsInserterRecoverableFailure,
		"inserterrecoverablestorage":                PrinterStatusDetailsInserterRecoverableStorage,
		"inserterremoved":                           PrinterStatusDetailsInserterRemoved,
		"inserterresourceadded":                     PrinterStatusDetailsInserterResourceAdded,
		"inserterresourceremoved":                   PrinterStatusDetailsInserterResourceRemoved,
		"inserterthermistorfailure":                 PrinterStatusDetailsInserterThermistorFailure,
		"insertertimingfailure":                     PrinterStatusDetailsInserterTimingFailure,
		"inserterturnedoff":                         PrinterStatusDetailsInserterTurnedOff,
		"inserterturnedon":                          PrinterStatusDetailsInserterTurnedOn,
		"inserterundertemperature":                  PrinterStatusDetailsInserterUnderTemperature,
		"inserterunrecoverablefailure":              PrinterStatusDetailsInserterUnrecoverableFailure,
		"inserterunrecoverablestorageerror":         PrinterStatusDetailsInserterUnrecoverableStorageError,
		"inserterwarmingup":                         PrinterStatusDetailsInserterWarmingUp,
		"interlockclosed":                           PrinterStatusDetailsInterlockClosed,
		"interlockopen":                             PrinterStatusDetailsInterlockOpen,
		"interpretercartridgeadded":                 PrinterStatusDetailsInterpreterCartridgeAdded,
		"interpretercartridgedeleted":               PrinterStatusDetailsInterpreterCartridgeDeleted,
		"interpretercomplexpageencountered":         PrinterStatusDetailsInterpreterComplexPageEncountered,
		"interpretermemorydecrease":                 PrinterStatusDetailsInterpreterMemoryDecrease,
		"interpretermemoryincrease":                 PrinterStatusDetailsInterpreterMemoryIncrease,
		"interpreterresourceadded":                  PrinterStatusDetailsInterpreterResourceAdded,
		"interpreterresourcedeleted":                PrinterStatusDetailsInterpreterResourceDeleted,
		"interpreterresourceunavailable":            PrinterStatusDetailsInterpreterResourceUnavailable,
		"lampateol":                                 PrinterStatusDetailsLampAtEol,
		"lampfailure":                               PrinterStatusDetailsLampFailure,
		"lampneareol":                               PrinterStatusDetailsLampNearEol,
		"laserateol":                                PrinterStatusDetailsLaserAtEol,
		"laserfailure":                              PrinterStatusDetailsLaserFailure,
		"laserneareol":                              PrinterStatusDetailsLaserNearEol,
		"makeenvelopeadded":                         PrinterStatusDetailsMakeEnvelopeAdded,
		"makeenvelopealmostempty":                   PrinterStatusDetailsMakeEnvelopeAlmostEmpty,
		"makeenvelopealmostfull":                    PrinterStatusDetailsMakeEnvelopeAlmostFull,
		"makeenvelopeatlimit":                       PrinterStatusDetailsMakeEnvelopeAtLimit,
		"makeenvelopeclosed":                        PrinterStatusDetailsMakeEnvelopeClosed,
		"makeenvelopeconfigurationchange":           PrinterStatusDetailsMakeEnvelopeConfigurationChange,
		"makeenvelopecoverclosed":                   PrinterStatusDetailsMakeEnvelopeCoverClosed,
		"makeenvelopecoveropen":                     PrinterStatusDetailsMakeEnvelopeCoverOpen,
		"makeenvelopeempty":                         PrinterStatusDetailsMakeEnvelopeEmpty,
		"makeenvelopefull":                          PrinterStatusDetailsMakeEnvelopeFull,
		"makeenvelopeinterlockclosed":               PrinterStatusDetailsMakeEnvelopeInterlockClosed,
		"makeenvelopeinterlockopen":                 PrinterStatusDetailsMakeEnvelopeInterlockOpen,
		"makeenvelopejam":                           PrinterStatusDetailsMakeEnvelopeJam,
		"makeenvelopelifealmostover":                PrinterStatusDetailsMakeEnvelopeLifeAlmostOver,
		"makeenvelopelifeover":                      PrinterStatusDetailsMakeEnvelopeLifeOver,
		"makeenvelopememoryexhausted":               PrinterStatusDetailsMakeEnvelopeMemoryExhausted,
		"makeenvelopemissing":                       PrinterStatusDetailsMakeEnvelopeMissing,
		"makeenvelopemotorfailure":                  PrinterStatusDetailsMakeEnvelopeMotorFailure,
		"makeenvelopenearlimit":                     PrinterStatusDetailsMakeEnvelopeNearLimit,
		"makeenvelopeoffline":                       PrinterStatusDetailsMakeEnvelopeOffline,
		"makeenvelopeopened":                        PrinterStatusDetailsMakeEnvelopeOpened,
		"makeenvelopeovertemperature":               PrinterStatusDetailsMakeEnvelopeOverTemperature,
		"makeenvelopepowersaver":                    PrinterStatusDetailsMakeEnvelopePowerSaver,
		"makeenveloperecoverablefailure":            PrinterStatusDetailsMakeEnvelopeRecoverableFailure,
		"makeenveloperecoverablestorage":            PrinterStatusDetailsMakeEnvelopeRecoverableStorage,
		"makeenveloperemoved":                       PrinterStatusDetailsMakeEnvelopeRemoved,
		"makeenveloperesourceadded":                 PrinterStatusDetailsMakeEnvelopeResourceAdded,
		"makeenveloperesourceremoved":               PrinterStatusDetailsMakeEnvelopeResourceRemoved,
		"makeenvelopethermistorfailure":             PrinterStatusDetailsMakeEnvelopeThermistorFailure,
		"makeenvelopetimingfailure":                 PrinterStatusDetailsMakeEnvelopeTimingFailure,
		"makeenvelopeturnedoff":                     PrinterStatusDetailsMakeEnvelopeTurnedOff,
		"makeenvelopeturnedon":                      PrinterStatusDetailsMakeEnvelopeTurnedOn,
		"makeenvelopeundertemperature":              PrinterStatusDetailsMakeEnvelopeUnderTemperature,
		"makeenvelopeunrecoverablefailure":          PrinterStatusDetailsMakeEnvelopeUnrecoverableFailure,
		"makeenvelopeunrecoverablestorageerror":     PrinterStatusDetailsMakeEnvelopeUnrecoverableStorageError,
		"makeenvelopewarmingup":                     PrinterStatusDetailsMakeEnvelopeWarmingUp,
		"markeradjustingprintquality":               PrinterStatusDetailsMarkerAdjustingPrintQuality,
		"markercleanermissing":                      PrinterStatusDetailsMarkerCleanerMissing,
		"markerdeveloperalmostempty":                PrinterStatusDetailsMarkerDeveloperAlmostEmpty,
		"markerdeveloperempty":                      PrinterStatusDetailsMarkerDeveloperEmpty,
		"markerdevelopermissing":                    PrinterStatusDetailsMarkerDeveloperMissing,
		"markerfusermissing":                        PrinterStatusDetailsMarkerFuserMissing,
		"markerfuserthermistorfailure":              PrinterStatusDetailsMarkerFuserThermistorFailure,
		"markerfusertimingfailure":                  PrinterStatusDetailsMarkerFuserTimingFailure,
		"markerinkalmostempty":                      PrinterStatusDetailsMarkerInkAlmostEmpty,
		"markerinkempty":                            PrinterStatusDetailsMarkerInkEmpty,
		"markerinkmissing":                          PrinterStatusDetailsMarkerInkMissing,
		"markeropcmissing":                          PrinterStatusDetailsMarkerOpcMissing,
		"markerprintribbonalmostempty":              PrinterStatusDetailsMarkerPrintRibbonAlmostEmpty,
		"markerprintribbonempty":                    PrinterStatusDetailsMarkerPrintRibbonEmpty,
		"markerprintribbonmissing":                  PrinterStatusDetailsMarkerPrintRibbonMissing,
		"markersupplyalmostempty":                   PrinterStatusDetailsMarkerSupplyAlmostEmpty,
		"markersupplyempty":                         PrinterStatusDetailsMarkerSupplyEmpty,
		"markersupplylow":                           PrinterStatusDetailsMarkerSupplyLow,
		"markersupplymissing":                       PrinterStatusDetailsMarkerSupplyMissing,
		"markertonercartridgemissing":               PrinterStatusDetailsMarkerTonerCartridgeMissing,
		"markertonermissing":                        PrinterStatusDetailsMarkerTonerMissing,
		"markerwastealmostfull":                     PrinterStatusDetailsMarkerWasteAlmostFull,
		"markerwastefull":                           PrinterStatusDetailsMarkerWasteFull,
		"markerwasteinkreceptaclealmostfull":        PrinterStatusDetailsMarkerWasteInkReceptacleAlmostFull,
		"markerwasteinkreceptaclefull":              PrinterStatusDetailsMarkerWasteInkReceptacleFull,
		"markerwasteinkreceptaclemissing":           PrinterStatusDetailsMarkerWasteInkReceptacleMissing,
		"markerwastemissing":                        PrinterStatusDetailsMarkerWasteMissing,
		"markerwastetonerreceptaclealmostfull":      PrinterStatusDetailsMarkerWasteTonerReceptacleAlmostFull,
		"markerwastetonerreceptaclefull":            PrinterStatusDetailsMarkerWasteTonerReceptacleFull,
		"markerwastetonerreceptaclemissing":         PrinterStatusDetailsMarkerWasteTonerReceptacleMissing,
		"materialempty":                             PrinterStatusDetailsMaterialEmpty,
		"materiallow":                               PrinterStatusDetailsMaterialLow,
		"materialneeded":                            PrinterStatusDetailsMaterialNeeded,
		"mediadrying":                               PrinterStatusDetailsMediaDrying,
		"mediaempty":                                PrinterStatusDetailsMediaEmpty,
		"mediajam":                                  PrinterStatusDetailsMediaJam,
		"medialow":                                  PrinterStatusDetailsMediaLow,
		"medianeeded":                               PrinterStatusDetailsMediaNeeded,
		"mediapathcannotduplexmediaselected":        PrinterStatusDetailsMediaPathCannotDuplexMediaSelected,
		"mediapathfailure":                          PrinterStatusDetailsMediaPathFailure,
		"mediapathinputempty":                       PrinterStatusDetailsMediaPathInputEmpty,
		"mediapathinputfeederror":                   PrinterStatusDetailsMediaPathInputFeedError,
		"mediapathinputjam":                         PrinterStatusDetailsMediaPathInputJam,
		"mediapathinputrequest":                     PrinterStatusDetailsMediaPathInputRequest,
		"mediapathjam":                              PrinterStatusDetailsMediaPathJam,
		"mediapathmediatrayalmostfull":              PrinterStatusDetailsMediaPathMediaTrayAlmostFull,
		"mediapathmediatrayfull":                    PrinterStatusDetailsMediaPathMediaTrayFull,
		"mediapathmediatraymissing":                 PrinterStatusDetailsMediaPathMediaTrayMissing,
		"mediapathoutputfeederror":                  PrinterStatusDetailsMediaPathOutputFeedError,
		"mediapathoutputfull":                       PrinterStatusDetailsMediaPathOutputFull,
		"mediapathoutputjam":                        PrinterStatusDetailsMediaPathOutputJam,
		"mediapathpickrollerfailure":                PrinterStatusDetailsMediaPathPickRollerFailure,
		"mediapathpickrollerlifeover":               PrinterStatusDetailsMediaPathPickRollerLifeOver,
		"mediapathpickrollerlifewarn":               PrinterStatusDetailsMediaPathPickRollerLifeWarn,
		"mediapathpickrollermissing":                PrinterStatusDetailsMediaPathPickRollerMissing,
		"motorfailure":                              PrinterStatusDetailsMotorFailure,
		"movingtopaused":                            PrinterStatusDetailsMovingToPaused,
		"none":                                      PrinterStatusDetailsNone,
		"opticalphotoconductorlifeover":             PrinterStatusDetailsOpticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife":        PrinterStatusDetailsOpticalPhotoConductorNearEndOfLife,
		"other":                                     PrinterStatusDetailsOther,
		"outputareaalmostfull":                      PrinterStatusDetailsOutputAreaAlmostFull,
		"outputareafull":                            PrinterStatusDetailsOutputAreaFull,
		"outputmailboxselectfailure":                PrinterStatusDetailsOutputMailboxSelectFailure,
		"outputmediatrayfailure":                    PrinterStatusDetailsOutputMediaTrayFailure,
		"outputmediatrayfeederror":                  PrinterStatusDetailsOutputMediaTrayFeedError,
		"outputmediatrayjam":                        PrinterStatusDetailsOutputMediaTrayJam,
		"outputtraymissing":                         PrinterStatusDetailsOutputTrayMissing,
		"paused":                                    PrinterStatusDetailsPaused,
		"perforateradded":                           PrinterStatusDetailsPerforaterAdded,
		"perforateralmostempty":                     PrinterStatusDetailsPerforaterAlmostEmpty,
		"perforateralmostfull":                      PrinterStatusDetailsPerforaterAlmostFull,
		"perforateratlimit":                         PrinterStatusDetailsPerforaterAtLimit,
		"perforaterclosed":                          PrinterStatusDetailsPerforaterClosed,
		"perforaterconfigurationchange":             PrinterStatusDetailsPerforaterConfigurationChange,
		"perforatercoverclosed":                     PrinterStatusDetailsPerforaterCoverClosed,
		"perforatercoveropen":                       PrinterStatusDetailsPerforaterCoverOpen,
		"perforaterempty":                           PrinterStatusDetailsPerforaterEmpty,
		"perforaterfull":                            PrinterStatusDetailsPerforaterFull,
		"perforaterinterlockclosed":                 PrinterStatusDetailsPerforaterInterlockClosed,
		"perforaterinterlockopen":                   PrinterStatusDetailsPerforaterInterlockOpen,
		"perforaterjam":                             PrinterStatusDetailsPerforaterJam,
		"perforaterlifealmostover":                  PrinterStatusDetailsPerforaterLifeAlmostOver,
		"perforaterlifeover":                        PrinterStatusDetailsPerforaterLifeOver,
		"perforatermemoryexhausted":                 PrinterStatusDetailsPerforaterMemoryExhausted,
		"perforatermissing":                         PrinterStatusDetailsPerforaterMissing,
		"perforatermotorfailure":                    PrinterStatusDetailsPerforaterMotorFailure,
		"perforaternearlimit":                       PrinterStatusDetailsPerforaterNearLimit,
		"perforateroffline":                         PrinterStatusDetailsPerforaterOffline,
		"perforateropened":                          PrinterStatusDetailsPerforaterOpened,
		"perforaterovertemperature":                 PrinterStatusDetailsPerforaterOverTemperature,
		"perforaterpowersaver":                      PrinterStatusDetailsPerforaterPowerSaver,
		"perforaterrecoverablefailure":              PrinterStatusDetailsPerforaterRecoverableFailure,
		"perforaterrecoverablestorage":              PrinterStatusDetailsPerforaterRecoverableStorage,
		"perforaterremoved":                         PrinterStatusDetailsPerforaterRemoved,
		"perforaterresourceadded":                   PrinterStatusDetailsPerforaterResourceAdded,
		"perforaterresourceremoved":                 PrinterStatusDetailsPerforaterResourceRemoved,
		"perforaterthermistorfailure":               PrinterStatusDetailsPerforaterThermistorFailure,
		"perforatertimingfailure":                   PrinterStatusDetailsPerforaterTimingFailure,
		"perforaterturnedoff":                       PrinterStatusDetailsPerforaterTurnedOff,
		"perforaterturnedon":                        PrinterStatusDetailsPerforaterTurnedOn,
		"perforaterundertemperature":                PrinterStatusDetailsPerforaterUnderTemperature,
		"perforaterunrecoverablefailure":            PrinterStatusDetailsPerforaterUnrecoverableFailure,
		"perforaterunrecoverablestorageerror":       PrinterStatusDetailsPerforaterUnrecoverableStorageError,
		"perforaterwarmingup":                       PrinterStatusDetailsPerforaterWarmingUp,
		"platformcooling":                           PrinterStatusDetailsPlatformCooling,
		"platformfailure":                           PrinterStatusDetailsPlatformFailure,
		"platformheating":                           PrinterStatusDetailsPlatformHeating,
		"platformtemperaturehigh":                   PrinterStatusDetailsPlatformTemperatureHigh,
		"platformtemperaturelow":                    PrinterStatusDetailsPlatformTemperatureLow,
		"powerdown":                                 PrinterStatusDetailsPowerDown,
		"powerup":                                   PrinterStatusDetailsPowerUp,
		"printermanualreset":                        PrinterStatusDetailsPrinterManualReset,
		"printernmsreset":                           PrinterStatusDetailsPrinterNmsReset,
		"printerreadytoprint":                       PrinterStatusDetailsPrinterReadyToPrint,
		"puncheradded":                              PrinterStatusDetailsPuncherAdded,
		"puncheralmostempty":                        PrinterStatusDetailsPuncherAlmostEmpty,
		"puncheralmostfull":                         PrinterStatusDetailsPuncherAlmostFull,
		"puncheratlimit":                            PrinterStatusDetailsPuncherAtLimit,
		"puncherclosed":                             PrinterStatusDetailsPuncherClosed,
		"puncherconfigurationchange":                PrinterStatusDetailsPuncherConfigurationChange,
		"punchercoverclosed":                        PrinterStatusDetailsPuncherCoverClosed,
		"punchercoveropen":                          PrinterStatusDetailsPuncherCoverOpen,
		"puncherempty":                              PrinterStatusDetailsPuncherEmpty,
		"puncherfull":                               PrinterStatusDetailsPuncherFull,
		"puncherinterlockclosed":                    PrinterStatusDetailsPuncherInterlockClosed,
		"puncherinterlockopen":                      PrinterStatusDetailsPuncherInterlockOpen,
		"puncherjam":                                PrinterStatusDetailsPuncherJam,
		"puncherlifealmostover":                     PrinterStatusDetailsPuncherLifeAlmostOver,
		"puncherlifeover":                           PrinterStatusDetailsPuncherLifeOver,
		"punchermemoryexhausted":                    PrinterStatusDetailsPuncherMemoryExhausted,
		"punchermissing":                            PrinterStatusDetailsPuncherMissing,
		"punchermotorfailure":                       PrinterStatusDetailsPuncherMotorFailure,
		"punchernearlimit":                          PrinterStatusDetailsPuncherNearLimit,
		"puncheroffline":                            PrinterStatusDetailsPuncherOffline,
		"puncheropened":                             PrinterStatusDetailsPuncherOpened,
		"puncherovertemperature":                    PrinterStatusDetailsPuncherOverTemperature,
		"puncherpowersaver":                         PrinterStatusDetailsPuncherPowerSaver,
		"puncherrecoverablefailure":                 PrinterStatusDetailsPuncherRecoverableFailure,
		"puncherrecoverablestorage":                 PrinterStatusDetailsPuncherRecoverableStorage,
		"puncherremoved":                            PrinterStatusDetailsPuncherRemoved,
		"puncherresourceadded":                      PrinterStatusDetailsPuncherResourceAdded,
		"puncherresourceremoved":                    PrinterStatusDetailsPuncherResourceRemoved,
		"puncherthermistorfailure":                  PrinterStatusDetailsPuncherThermistorFailure,
		"punchertimingfailure":                      PrinterStatusDetailsPuncherTimingFailure,
		"puncherturnedoff":                          PrinterStatusDetailsPuncherTurnedOff,
		"puncherturnedon":                           PrinterStatusDetailsPuncherTurnedOn,
		"puncherundertemperature":                   PrinterStatusDetailsPuncherUnderTemperature,
		"puncherunrecoverablefailure":               PrinterStatusDetailsPuncherUnrecoverableFailure,
		"puncherunrecoverablestorageerror":          PrinterStatusDetailsPuncherUnrecoverableStorageError,
		"puncherwarmingup":                          PrinterStatusDetailsPuncherWarmingUp,
		"resuming":                                  PrinterStatusDetailsResuming,
		"scanmediapathfailure":                      PrinterStatusDetailsScanMediaPathFailure,
		"scanmediapathinputempty":                   PrinterStatusDetailsScanMediaPathInputEmpty,
		"scanmediapathinputfeederror":               PrinterStatusDetailsScanMediaPathInputFeedError,
		"scanmediapathinputjam":                     PrinterStatusDetailsScanMediaPathInputJam,
		"scanmediapathinputrequest":                 PrinterStatusDetailsScanMediaPathInputRequest,
		"scanmediapathjam":                          PrinterStatusDetailsScanMediaPathJam,
		"scanmediapathoutputfeederror":              PrinterStatusDetailsScanMediaPathOutputFeedError,
		"scanmediapathoutputfull":                   PrinterStatusDetailsScanMediaPathOutputFull,
		"scanmediapathoutputjam":                    PrinterStatusDetailsScanMediaPathOutputJam,
		"scanmediapathpickrollerfailure":            PrinterStatusDetailsScanMediaPathPickRollerFailure,
		"scanmediapathpickrollerlifeover":           PrinterStatusDetailsScanMediaPathPickRollerLifeOver,
		"scanmediapathpickrollerlifewarn":           PrinterStatusDetailsScanMediaPathPickRollerLifeWarn,
		"scanmediapathpickrollermissing":            PrinterStatusDetailsScanMediaPathPickRollerMissing,
		"scanmediapathtrayalmostfull":               PrinterStatusDetailsScanMediaPathTrayAlmostFull,
		"scanmediapathtrayfull":                     PrinterStatusDetailsScanMediaPathTrayFull,
		"scanmediapathtraymissing":                  PrinterStatusDetailsScanMediaPathTrayMissing,
		"scannerlightfailure":                       PrinterStatusDetailsScannerLightFailure,
		"scannerlightlifealmostover":                PrinterStatusDetailsScannerLightLifeAlmostOver,
		"scannerlightlifeover":                      PrinterStatusDetailsScannerLightLifeOver,
		"scannerlightmissing":                       PrinterStatusDetailsScannerLightMissing,
		"scannersensorfailure":                      PrinterStatusDetailsScannerSensorFailure,
		"scannersensorlifealmostover":               PrinterStatusDetailsScannerSensorLifeAlmostOver,
		"scannersensorlifeover":                     PrinterStatusDetailsScannerSensorLifeOver,
		"scannersensormissing":                      PrinterStatusDetailsScannerSensorMissing,
		"separationcutteradded":                     PrinterStatusDetailsSeparationCutterAdded,
		"separationcutteralmostempty":               PrinterStatusDetailsSeparationCutterAlmostEmpty,
		"separationcutteralmostfull":                PrinterStatusDetailsSeparationCutterAlmostFull,
		"separationcutteratlimit":                   PrinterStatusDetailsSeparationCutterAtLimit,
		"separationcutterclosed":                    PrinterStatusDetailsSeparationCutterClosed,
		"separationcutterconfigurationchange":       PrinterStatusDetailsSeparationCutterConfigurationChange,
		"separationcuttercoverclosed":               PrinterStatusDetailsSeparationCutterCoverClosed,
		"separationcuttercoveropen":                 PrinterStatusDetailsSeparationCutterCoverOpen,
		"separationcutterempty":                     PrinterStatusDetailsSeparationCutterEmpty,
		"separationcutterfull":                      PrinterStatusDetailsSeparationCutterFull,
		"separationcutterinterlockclosed":           PrinterStatusDetailsSeparationCutterInterlockClosed,
		"separationcutterinterlockopen":             PrinterStatusDetailsSeparationCutterInterlockOpen,
		"separationcutterjam":                       PrinterStatusDetailsSeparationCutterJam,
		"separationcutterlifealmostover":            PrinterStatusDetailsSeparationCutterLifeAlmostOver,
		"separationcutterlifeover":                  PrinterStatusDetailsSeparationCutterLifeOver,
		"separationcuttermemoryexhausted":           PrinterStatusDetailsSeparationCutterMemoryExhausted,
		"separationcuttermissing":                   PrinterStatusDetailsSeparationCutterMissing,
		"separationcuttermotorfailure":              PrinterStatusDetailsSeparationCutterMotorFailure,
		"separationcutternearlimit":                 PrinterStatusDetailsSeparationCutterNearLimit,
		"separationcutteroffline":                   PrinterStatusDetailsSeparationCutterOffline,
		"separationcutteropened":                    PrinterStatusDetailsSeparationCutterOpened,
		"separationcutterovertemperature":           PrinterStatusDetailsSeparationCutterOverTemperature,
		"separationcutterpowersaver":                PrinterStatusDetailsSeparationCutterPowerSaver,
		"separationcutterrecoverablefailure":        PrinterStatusDetailsSeparationCutterRecoverableFailure,
		"separationcutterrecoverablestorage":        PrinterStatusDetailsSeparationCutterRecoverableStorage,
		"separationcutterremoved":                   PrinterStatusDetailsSeparationCutterRemoved,
		"separationcutterresourceadded":             PrinterStatusDetailsSeparationCutterResourceAdded,
		"separationcutterresourceremoved":           PrinterStatusDetailsSeparationCutterResourceRemoved,
		"separationcutterthermistorfailure":         PrinterStatusDetailsSeparationCutterThermistorFailure,
		"separationcuttertimingfailure":             PrinterStatusDetailsSeparationCutterTimingFailure,
		"separationcutterturnedoff":                 PrinterStatusDetailsSeparationCutterTurnedOff,
		"separationcutterturnedon":                  PrinterStatusDetailsSeparationCutterTurnedOn,
		"separationcutterundertemperature":          PrinterStatusDetailsSeparationCutterUnderTemperature,
		"separationcutterunrecoverablefailure":      PrinterStatusDetailsSeparationCutterUnrecoverableFailure,
		"separationcutterunrecoverablestorageerror": PrinterStatusDetailsSeparationCutterUnrecoverableStorageError,
		"separationcutterwarmingup":                 PrinterStatusDetailsSeparationCutterWarmingUp,
		"sheetrotatoradded":                         PrinterStatusDetailsSheetRotatorAdded,
		"sheetrotatoralmostempty":                   PrinterStatusDetailsSheetRotatorAlmostEmpty,
		"sheetrotatoralmostfull":                    PrinterStatusDetailsSheetRotatorAlmostFull,
		"sheetrotatoratlimit":                       PrinterStatusDetailsSheetRotatorAtLimit,
		"sheetrotatorclosed":                        PrinterStatusDetailsSheetRotatorClosed,
		"sheetrotatorconfigurationchange":           PrinterStatusDetailsSheetRotatorConfigurationChange,
		"sheetrotatorcoverclosed":                   PrinterStatusDetailsSheetRotatorCoverClosed,
		"sheetrotatorcoveropen":                     PrinterStatusDetailsSheetRotatorCoverOpen,
		"sheetrotatorempty":                         PrinterStatusDetailsSheetRotatorEmpty,
		"sheetrotatorfull":                          PrinterStatusDetailsSheetRotatorFull,
		"sheetrotatorinterlockclosed":               PrinterStatusDetailsSheetRotatorInterlockClosed,
		"sheetrotatorinterlockopen":                 PrinterStatusDetailsSheetRotatorInterlockOpen,
		"sheetrotatorjam":                           PrinterStatusDetailsSheetRotatorJam,
		"sheetrotatorlifealmostover":                PrinterStatusDetailsSheetRotatorLifeAlmostOver,
		"sheetrotatorlifeover":                      PrinterStatusDetailsSheetRotatorLifeOver,
		"sheetrotatormemoryexhausted":               PrinterStatusDetailsSheetRotatorMemoryExhausted,
		"sheetrotatormissing":                       PrinterStatusDetailsSheetRotatorMissing,
		"sheetrotatormotorfailure":                  PrinterStatusDetailsSheetRotatorMotorFailure,
		"sheetrotatornearlimit":                     PrinterStatusDetailsSheetRotatorNearLimit,
		"sheetrotatoroffline":                       PrinterStatusDetailsSheetRotatorOffline,
		"sheetrotatoropened":                        PrinterStatusDetailsSheetRotatorOpened,
		"sheetrotatorovertemperature":               PrinterStatusDetailsSheetRotatorOverTemperature,
		"sheetrotatorpowersaver":                    PrinterStatusDetailsSheetRotatorPowerSaver,
		"sheetrotatorrecoverablefailure":            PrinterStatusDetailsSheetRotatorRecoverableFailure,
		"sheetrotatorrecoverablestorage":            PrinterStatusDetailsSheetRotatorRecoverableStorage,
		"sheetrotatorremoved":                       PrinterStatusDetailsSheetRotatorRemoved,
		"sheetrotatorresourceadded":                 PrinterStatusDetailsSheetRotatorResourceAdded,
		"sheetrotatorresourceremoved":               PrinterStatusDetailsSheetRotatorResourceRemoved,
		"sheetrotatorthermistorfailure":             PrinterStatusDetailsSheetRotatorThermistorFailure,
		"sheetrotatortimingfailure":                 PrinterStatusDetailsSheetRotatorTimingFailure,
		"sheetrotatorturnedoff":                     PrinterStatusDetailsSheetRotatorTurnedOff,
		"sheetrotatorturnedon":                      PrinterStatusDetailsSheetRotatorTurnedOn,
		"sheetrotatorundertemperature":              PrinterStatusDetailsSheetRotatorUnderTemperature,
		"sheetrotatorunrecoverablefailure":          PrinterStatusDetailsSheetRotatorUnrecoverableFailure,
		"sheetrotatorunrecoverablestorageerror":     PrinterStatusDetailsSheetRotatorUnrecoverableStorageError,
		"sheetrotatorwarmingup":                     PrinterStatusDetailsSheetRotatorWarmingUp,
		"shutdown":                                  PrinterStatusDetailsShutdown,
		"slitteradded":                              PrinterStatusDetailsSlitterAdded,
		"slitteralmostempty":                        PrinterStatusDetailsSlitterAlmostEmpty,
		"slitteralmostfull":                         PrinterStatusDetailsSlitterAlmostFull,
		"slitteratlimit":                            PrinterStatusDetailsSlitterAtLimit,
		"slitterclosed":                             PrinterStatusDetailsSlitterClosed,
		"slitterconfigurationchange":                PrinterStatusDetailsSlitterConfigurationChange,
		"slittercoverclosed":                        PrinterStatusDetailsSlitterCoverClosed,
		"slittercoveropen":                          PrinterStatusDetailsSlitterCoverOpen,
		"slitterempty":                              PrinterStatusDetailsSlitterEmpty,
		"slitterfull":                               PrinterStatusDetailsSlitterFull,
		"slitterinterlockclosed":                    PrinterStatusDetailsSlitterInterlockClosed,
		"slitterinterlockopen":                      PrinterStatusDetailsSlitterInterlockOpen,
		"slitterjam":                                PrinterStatusDetailsSlitterJam,
		"slitterlifealmostover":                     PrinterStatusDetailsSlitterLifeAlmostOver,
		"slitterlifeover":                           PrinterStatusDetailsSlitterLifeOver,
		"slittermemoryexhausted":                    PrinterStatusDetailsSlitterMemoryExhausted,
		"slittermissing":                            PrinterStatusDetailsSlitterMissing,
		"slittermotorfailure":                       PrinterStatusDetailsSlitterMotorFailure,
		"slitternearlimit":                          PrinterStatusDetailsSlitterNearLimit,
		"slitteroffline":                            PrinterStatusDetailsSlitterOffline,
		"slitteropened":                             PrinterStatusDetailsSlitterOpened,
		"slitterovertemperature":                    PrinterStatusDetailsSlitterOverTemperature,
		"slitterpowersaver":                         PrinterStatusDetailsSlitterPowerSaver,
		"slitterrecoverablefailure":                 PrinterStatusDetailsSlitterRecoverableFailure,
		"slitterrecoverablestorage":                 PrinterStatusDetailsSlitterRecoverableStorage,
		"slitterremoved":                            PrinterStatusDetailsSlitterRemoved,
		"slitterresourceadded":                      PrinterStatusDetailsSlitterResourceAdded,
		"slitterresourceremoved":                    PrinterStatusDetailsSlitterResourceRemoved,
		"slitterthermistorfailure":                  PrinterStatusDetailsSlitterThermistorFailure,
		"slittertimingfailure":                      PrinterStatusDetailsSlitterTimingFailure,
		"slitterturnedoff":                          PrinterStatusDetailsSlitterTurnedOff,
		"slitterturnedon":                           PrinterStatusDetailsSlitterTurnedOn,
		"slitterundertemperature":                   PrinterStatusDetailsSlitterUnderTemperature,
		"slitterunrecoverablefailure":               PrinterStatusDetailsSlitterUnrecoverableFailure,
		"slitterunrecoverablestorageerror":          PrinterStatusDetailsSlitterUnrecoverableStorageError,
		"slitterwarmingup":                          PrinterStatusDetailsSlitterWarmingUp,
		"spoolareafull":                             PrinterStatusDetailsSpoolAreaFull,
		"stackeradded":                              PrinterStatusDetailsStackerAdded,
		"stackeralmostempty":                        PrinterStatusDetailsStackerAlmostEmpty,
		"stackeralmostfull":                         PrinterStatusDetailsStackerAlmostFull,
		"stackeratlimit":                            PrinterStatusDetailsStackerAtLimit,
		"stackerclosed":                             PrinterStatusDetailsStackerClosed,
		"stackerconfigurationchange":                PrinterStatusDetailsStackerConfigurationChange,
		"stackercoverclosed":                        PrinterStatusDetailsStackerCoverClosed,
		"stackercoveropen":                          PrinterStatusDetailsStackerCoverOpen,
		"stackerempty":                              PrinterStatusDetailsStackerEmpty,
		"stackerfull":                               PrinterStatusDetailsStackerFull,
		"stackerinterlockclosed":                    PrinterStatusDetailsStackerInterlockClosed,
		"stackerinterlockopen":                      PrinterStatusDetailsStackerInterlockOpen,
		"stackerjam":                                PrinterStatusDetailsStackerJam,
		"stackerlifealmostover":                     PrinterStatusDetailsStackerLifeAlmostOver,
		"stackerlifeover":                           PrinterStatusDetailsStackerLifeOver,
		"stackermemoryexhausted":                    PrinterStatusDetailsStackerMemoryExhausted,
		"stackermissing":                            PrinterStatusDetailsStackerMissing,
		"stackermotorfailure":                       PrinterStatusDetailsStackerMotorFailure,
		"stackernearlimit":                          PrinterStatusDetailsStackerNearLimit,
		"stackeroffline":                            PrinterStatusDetailsStackerOffline,
		"stackeropened":                             PrinterStatusDetailsStackerOpened,
		"stackerovertemperature":                    PrinterStatusDetailsStackerOverTemperature,
		"stackerpowersaver":                         PrinterStatusDetailsStackerPowerSaver,
		"stackerrecoverablefailure":                 PrinterStatusDetailsStackerRecoverableFailure,
		"stackerrecoverablestorage":                 PrinterStatusDetailsStackerRecoverableStorage,
		"stackerremoved":                            PrinterStatusDetailsStackerRemoved,
		"stackerresourceadded":                      PrinterStatusDetailsStackerResourceAdded,
		"stackerresourceremoved":                    PrinterStatusDetailsStackerResourceRemoved,
		"stackerthermistorfailure":                  PrinterStatusDetailsStackerThermistorFailure,
		"stackertimingfailure":                      PrinterStatusDetailsStackerTimingFailure,
		"stackerturnedoff":                          PrinterStatusDetailsStackerTurnedOff,
		"stackerturnedon":                           PrinterStatusDetailsStackerTurnedOn,
		"stackerundertemperature":                   PrinterStatusDetailsStackerUnderTemperature,
		"stackerunrecoverablefailure":               PrinterStatusDetailsStackerUnrecoverableFailure,
		"stackerunrecoverablestorageerror":          PrinterStatusDetailsStackerUnrecoverableStorageError,
		"stackerwarmingup":                          PrinterStatusDetailsStackerWarmingUp,
		"standby":                                   PrinterStatusDetailsStandby,
		"stapleradded":                              PrinterStatusDetailsStaplerAdded,
		"stapleralmostempty":                        PrinterStatusDetailsStaplerAlmostEmpty,
		"stapleralmostfull":                         PrinterStatusDetailsStaplerAlmostFull,
		"stapleratlimit":                            PrinterStatusDetailsStaplerAtLimit,
		"staplerclosed":                             PrinterStatusDetailsStaplerClosed,
		"staplerconfigurationchange":                PrinterStatusDetailsStaplerConfigurationChange,
		"staplercoverclosed":                        PrinterStatusDetailsStaplerCoverClosed,
		"staplercoveropen":                          PrinterStatusDetailsStaplerCoverOpen,
		"staplerempty":                              PrinterStatusDetailsStaplerEmpty,
		"staplerfull":                               PrinterStatusDetailsStaplerFull,
		"staplerinterlockclosed":                    PrinterStatusDetailsStaplerInterlockClosed,
		"staplerinterlockopen":                      PrinterStatusDetailsStaplerInterlockOpen,
		"staplerjam":                                PrinterStatusDetailsStaplerJam,
		"staplerlifealmostover":                     PrinterStatusDetailsStaplerLifeAlmostOver,
		"staplerlifeover":                           PrinterStatusDetailsStaplerLifeOver,
		"staplermemoryexhausted":                    PrinterStatusDetailsStaplerMemoryExhausted,
		"staplermissing":                            PrinterStatusDetailsStaplerMissing,
		"staplermotorfailure":                       PrinterStatusDetailsStaplerMotorFailure,
		"staplernearlimit":                          PrinterStatusDetailsStaplerNearLimit,
		"stapleroffline":                            PrinterStatusDetailsStaplerOffline,
		"stapleropened":                             PrinterStatusDetailsStaplerOpened,
		"staplerovertemperature":                    PrinterStatusDetailsStaplerOverTemperature,
		"staplerpowersaver":                         PrinterStatusDetailsStaplerPowerSaver,
		"staplerrecoverablefailure":                 PrinterStatusDetailsStaplerRecoverableFailure,
		"staplerrecoverablestorage":                 PrinterStatusDetailsStaplerRecoverableStorage,
		"staplerremoved":                            PrinterStatusDetailsStaplerRemoved,
		"staplerresourceadded":                      PrinterStatusDetailsStaplerResourceAdded,
		"staplerresourceremoved":                    PrinterStatusDetailsStaplerResourceRemoved,
		"staplerthermistorfailure":                  PrinterStatusDetailsStaplerThermistorFailure,
		"staplertimingfailure":                      PrinterStatusDetailsStaplerTimingFailure,
		"staplerturnedoff":                          PrinterStatusDetailsStaplerTurnedOff,
		"staplerturnedon":                           PrinterStatusDetailsStaplerTurnedOn,
		"staplerundertemperature":                   PrinterStatusDetailsStaplerUnderTemperature,
		"staplerunrecoverablefailure":               PrinterStatusDetailsStaplerUnrecoverableFailure,
		"staplerunrecoverablestorageerror":          PrinterStatusDetailsStaplerUnrecoverableStorageError,
		"staplerwarmingup":                          PrinterStatusDetailsStaplerWarmingUp,
		"stitcheradded":                             PrinterStatusDetailsStitcherAdded,
		"stitcheralmostempty":                       PrinterStatusDetailsStitcherAlmostEmpty,
		"stitcheralmostfull":                        PrinterStatusDetailsStitcherAlmostFull,
		"stitcheratlimit":                           PrinterStatusDetailsStitcherAtLimit,
		"stitcherclosed":                            PrinterStatusDetailsStitcherClosed,
		"stitcherconfigurationchange":               PrinterStatusDetailsStitcherConfigurationChange,
		"stitchercoverclosed":                       PrinterStatusDetailsStitcherCoverClosed,
		"stitchercoveropen":                         PrinterStatusDetailsStitcherCoverOpen,
		"stitcherempty":                             PrinterStatusDetailsStitcherEmpty,
		"stitcherfull":                              PrinterStatusDetailsStitcherFull,
		"stitcherinterlockclosed":                   PrinterStatusDetailsStitcherInterlockClosed,
		"stitcherinterlockopen":                     PrinterStatusDetailsStitcherInterlockOpen,
		"stitcherjam":                               PrinterStatusDetailsStitcherJam,
		"stitcherlifealmostover":                    PrinterStatusDetailsStitcherLifeAlmostOver,
		"stitcherlifeover":                          PrinterStatusDetailsStitcherLifeOver,
		"stitchermemoryexhausted":                   PrinterStatusDetailsStitcherMemoryExhausted,
		"stitchermissing":                           PrinterStatusDetailsStitcherMissing,
		"stitchermotorfailure":                      PrinterStatusDetailsStitcherMotorFailure,
		"stitchernearlimit":                         PrinterStatusDetailsStitcherNearLimit,
		"stitcheroffline":                           PrinterStatusDetailsStitcherOffline,
		"stitcheropened":                            PrinterStatusDetailsStitcherOpened,
		"stitcherovertemperature":                   PrinterStatusDetailsStitcherOverTemperature,
		"stitcherpowersaver":                        PrinterStatusDetailsStitcherPowerSaver,
		"stitcherrecoverablefailure":                PrinterStatusDetailsStitcherRecoverableFailure,
		"stitcherrecoverablestorage":                PrinterStatusDetailsStitcherRecoverableStorage,
		"stitcherremoved":                           PrinterStatusDetailsStitcherRemoved,
		"stitcherresourceadded":                     PrinterStatusDetailsStitcherResourceAdded,
		"stitcherresourceremoved":                   PrinterStatusDetailsStitcherResourceRemoved,
		"stitcherthermistorfailure":                 PrinterStatusDetailsStitcherThermistorFailure,
		"stitchertimingfailure":                     PrinterStatusDetailsStitcherTimingFailure,
		"stitcherturnedoff":                         PrinterStatusDetailsStitcherTurnedOff,
		"stitcherturnedon":                          PrinterStatusDetailsStitcherTurnedOn,
		"stitcherundertemperature":                  PrinterStatusDetailsStitcherUnderTemperature,
		"stitcherunrecoverablefailure":              PrinterStatusDetailsStitcherUnrecoverableFailure,
		"stitcherunrecoverablestorageerror":         PrinterStatusDetailsStitcherUnrecoverableStorageError,
		"stitcherwarmingup":                         PrinterStatusDetailsStitcherWarmingUp,
		"stoppedpartially":                          PrinterStatusDetailsStoppedPartially,
		"stopping":                                  PrinterStatusDetailsStopping,
		"subunitadded":                              PrinterStatusDetailsSubunitAdded,
		"subunitalmostempty":                        PrinterStatusDetailsSubunitAlmostEmpty,
		"subunitalmostfull":                         PrinterStatusDetailsSubunitAlmostFull,
		"subunitatlimit":                            PrinterStatusDetailsSubunitAtLimit,
		"subunitclosed":                             PrinterStatusDetailsSubunitClosed,
		"subunitcoolingdown":                        PrinterStatusDetailsSubunitCoolingDown,
		"subunitempty":                              PrinterStatusDetailsSubunitEmpty,
		"subunitfull":                               PrinterStatusDetailsSubunitFull,
		"subunitlifealmostover":                     PrinterStatusDetailsSubunitLifeAlmostOver,
		"subunitlifeover":                           PrinterStatusDetailsSubunitLifeOver,
		"subunitmemoryexhausted":                    PrinterStatusDetailsSubunitMemoryExhausted,
		"subunitmissing":                            PrinterStatusDetailsSubunitMissing,
		"subunitmotorfailure":                       PrinterStatusDetailsSubunitMotorFailure,
		"subunitnearlimit":                          PrinterStatusDetailsSubunitNearLimit,
		"subunitoffline":                            PrinterStatusDetailsSubunitOffline,
		"subunitopened":                             PrinterStatusDetailsSubunitOpened,
		"subunitovertemperature":                    PrinterStatusDetailsSubunitOverTemperature,
		"subunitpowersaver":                         PrinterStatusDetailsSubunitPowerSaver,
		"subunitrecoverablefailure":                 PrinterStatusDetailsSubunitRecoverableFailure,
		"subunitrecoverablestorage":                 PrinterStatusDetailsSubunitRecoverableStorage,
		"subunitremoved":                            PrinterStatusDetailsSubunitRemoved,
		"subunitresourceadded":                      PrinterStatusDetailsSubunitResourceAdded,
		"subunitresourceremoved":                    PrinterStatusDetailsSubunitResourceRemoved,
		"subunitthermistorfailure":                  PrinterStatusDetailsSubunitThermistorFailure,
		"subunittimingfailure":                      PrinterStatusDetailsSubunitTimingFailure,
		"subunitturnedoff":                          PrinterStatusDetailsSubunitTurnedOff,
		"subunitturnedon":                           PrinterStatusDetailsSubunitTurnedOn,
		"subunitundertemperature":                   PrinterStatusDetailsSubunitUnderTemperature,
		"subunitunrecoverablefailure":               PrinterStatusDetailsSubunitUnrecoverableFailure,
		"subunitunrecoverablestorage":               PrinterStatusDetailsSubunitUnrecoverableStorage,
		"subunitwarmingup":                          PrinterStatusDetailsSubunitWarmingUp,
		"suspend":                                   PrinterStatusDetailsSuspend,
		"testing":                                   PrinterStatusDetailsTesting,
		"timedout":                                  PrinterStatusDetailsTimedOut,
		"tonerempty":                                PrinterStatusDetailsTonerEmpty,
		"tonerlow":                                  PrinterStatusDetailsTonerLow,
		"trimmeradded":                              PrinterStatusDetailsTrimmerAdded,
		"trimmeralmostempty":                        PrinterStatusDetailsTrimmerAlmostEmpty,
		"trimmeralmostfull":                         PrinterStatusDetailsTrimmerAlmostFull,
		"trimmeratlimit":                            PrinterStatusDetailsTrimmerAtLimit,
		"trimmerclosed":                             PrinterStatusDetailsTrimmerClosed,
		"trimmerconfigurationchange":                PrinterStatusDetailsTrimmerConfigurationChange,
		"trimmercoverclosed":                        PrinterStatusDetailsTrimmerCoverClosed,
		"trimmercoveropen":                          PrinterStatusDetailsTrimmerCoverOpen,
		"trimmerempty":                              PrinterStatusDetailsTrimmerEmpty,
		"trimmerfull":                               PrinterStatusDetailsTrimmerFull,
		"trimmerinterlockclosed":                    PrinterStatusDetailsTrimmerInterlockClosed,
		"trimmerinterlockopen":                      PrinterStatusDetailsTrimmerInterlockOpen,
		"trimmerjam":                                PrinterStatusDetailsTrimmerJam,
		"trimmerlifealmostover":                     PrinterStatusDetailsTrimmerLifeAlmostOver,
		"trimmerlifeover":                           PrinterStatusDetailsTrimmerLifeOver,
		"trimmermemoryexhausted":                    PrinterStatusDetailsTrimmerMemoryExhausted,
		"trimmermissing":                            PrinterStatusDetailsTrimmerMissing,
		"trimmermotorfailure":                       PrinterStatusDetailsTrimmerMotorFailure,
		"trimmernearlimit":                          PrinterStatusDetailsTrimmerNearLimit,
		"trimmeroffline":                            PrinterStatusDetailsTrimmerOffline,
		"trimmeropened":                             PrinterStatusDetailsTrimmerOpened,
		"trimmerovertemperature":                    PrinterStatusDetailsTrimmerOverTemperature,
		"trimmerpowersaver":                         PrinterStatusDetailsTrimmerPowerSaver,
		"trimmerrecoverablefailure":                 PrinterStatusDetailsTrimmerRecoverableFailure,
		"trimmerrecoverablestorage":                 PrinterStatusDetailsTrimmerRecoverableStorage,
		"trimmerremoved":                            PrinterStatusDetailsTrimmerRemoved,
		"trimmerresourceadded":                      PrinterStatusDetailsTrimmerResourceAdded,
		"trimmerresourceremoved":                    PrinterStatusDetailsTrimmerResourceRemoved,
		"trimmerthermistorfailure":                  PrinterStatusDetailsTrimmerThermistorFailure,
		"trimmertimingfailure":                      PrinterStatusDetailsTrimmerTimingFailure,
		"trimmerturnedoff":                          PrinterStatusDetailsTrimmerTurnedOff,
		"trimmerturnedon":                           PrinterStatusDetailsTrimmerTurnedOn,
		"trimmerundertemperature":                   PrinterStatusDetailsTrimmerUnderTemperature,
		"trimmerunrecoverablefailure":               PrinterStatusDetailsTrimmerUnrecoverableFailure,
		"trimmerunrecoverablestorageerror":          PrinterStatusDetailsTrimmerUnrecoverableStorageError,
		"trimmerwarmingup":                          PrinterStatusDetailsTrimmerWarmingUp,
		"unknown":                                   PrinterStatusDetailsUnknown,
		"wrapperadded":                              PrinterStatusDetailsWrapperAdded,
		"wrapperalmostempty":                        PrinterStatusDetailsWrapperAlmostEmpty,
		"wrapperalmostfull":                         PrinterStatusDetailsWrapperAlmostFull,
		"wrapperatlimit":                            PrinterStatusDetailsWrapperAtLimit,
		"wrapperclosed":                             PrinterStatusDetailsWrapperClosed,
		"wrapperconfigurationchange":                PrinterStatusDetailsWrapperConfigurationChange,
		"wrappercoverclosed":                        PrinterStatusDetailsWrapperCoverClosed,
		"wrappercoveropen":                          PrinterStatusDetailsWrapperCoverOpen,
		"wrapperempty":                              PrinterStatusDetailsWrapperEmpty,
		"wrapperfull":                               PrinterStatusDetailsWrapperFull,
		"wrapperinterlockclosed":                    PrinterStatusDetailsWrapperInterlockClosed,
		"wrapperinterlockopen":                      PrinterStatusDetailsWrapperInterlockOpen,
		"wrapperjam":                                PrinterStatusDetailsWrapperJam,
		"wrapperlifealmostover":                     PrinterStatusDetailsWrapperLifeAlmostOver,
		"wrapperlifeover":                           PrinterStatusDetailsWrapperLifeOver,
		"wrappermemoryexhausted":                    PrinterStatusDetailsWrapperMemoryExhausted,
		"wrappermissing":                            PrinterStatusDetailsWrapperMissing,
		"wrappermotorfailure":                       PrinterStatusDetailsWrapperMotorFailure,
		"wrappernearlimit":                          PrinterStatusDetailsWrapperNearLimit,
		"wrapperoffline":                            PrinterStatusDetailsWrapperOffline,
		"wrapperopened":                             PrinterStatusDetailsWrapperOpened,
		"wrapperovertemperature":                    PrinterStatusDetailsWrapperOverTemperature,
		"wrapperpowersaver":                         PrinterStatusDetailsWrapperPowerSaver,
		"wrapperrecoverablefailure":                 PrinterStatusDetailsWrapperRecoverableFailure,
		"wrapperrecoverablestorage":                 PrinterStatusDetailsWrapperRecoverableStorage,
		"wrapperremoved":                            PrinterStatusDetailsWrapperRemoved,
		"wrapperresourceadded":                      PrinterStatusDetailsWrapperResourceAdded,
		"wrapperresourceremoved":                    PrinterStatusDetailsWrapperResourceRemoved,
		"wrapperthermistorfailure":                  PrinterStatusDetailsWrapperThermistorFailure,
		"wrappertimingfailure":                      PrinterStatusDetailsWrapperTimingFailure,
		"wrapperturnedoff":                          PrinterStatusDetailsWrapperTurnedOff,
		"wrapperturnedon":                           PrinterStatusDetailsWrapperTurnedOn,
		"wrapperundertemperature":                   PrinterStatusDetailsWrapperUnderTemperature,
		"wrapperunrecoverablefailure":               PrinterStatusDetailsWrapperUnrecoverableFailure,
		"wrapperunrecoverablestorageerror":          PrinterStatusDetailsWrapperUnrecoverableStorageError,
		"wrapperwarmingup":                          PrinterStatusDetailsWrapperWarmingUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusDetails(input)
	return &out, nil
}

type PrinterStatusState string

const (
	PrinterStatusStateIdle       PrinterStatusState = "idle"
	PrinterStatusStateProcessing PrinterStatusState = "processing"
	PrinterStatusStateStopped    PrinterStatusState = "stopped"
	PrinterStatusStateUnknown    PrinterStatusState = "unknown"
)

func PossibleValuesForPrinterStatusState() []string {
	return []string{
		string(PrinterStatusStateIdle),
		string(PrinterStatusStateProcessing),
		string(PrinterStatusStateStopped),
		string(PrinterStatusStateUnknown),
	}
}

func (s *PrinterStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusState(input string) (*PrinterStatusState, error) {
	vals := map[string]PrinterStatusState{
		"idle":       PrinterStatusStateIdle,
		"processing": PrinterStatusStateProcessing,
		"stopped":    PrinterStatusStateStopped,
		"unknown":    PrinterStatusStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusState(input)
	return &out, nil
}

type RecurrencePatternDaysOfWeek string

const (
	RecurrencePatternDaysOfWeekFriday    RecurrencePatternDaysOfWeek = "friday"
	RecurrencePatternDaysOfWeekMonday    RecurrencePatternDaysOfWeek = "monday"
	RecurrencePatternDaysOfWeekSaturday  RecurrencePatternDaysOfWeek = "saturday"
	RecurrencePatternDaysOfWeekSunday    RecurrencePatternDaysOfWeek = "sunday"
	RecurrencePatternDaysOfWeekThursday  RecurrencePatternDaysOfWeek = "thursday"
	RecurrencePatternDaysOfWeekTuesday   RecurrencePatternDaysOfWeek = "tuesday"
	RecurrencePatternDaysOfWeekWednesday RecurrencePatternDaysOfWeek = "wednesday"
)

func PossibleValuesForRecurrencePatternDaysOfWeek() []string {
	return []string{
		string(RecurrencePatternDaysOfWeekFriday),
		string(RecurrencePatternDaysOfWeekMonday),
		string(RecurrencePatternDaysOfWeekSaturday),
		string(RecurrencePatternDaysOfWeekSunday),
		string(RecurrencePatternDaysOfWeekThursday),
		string(RecurrencePatternDaysOfWeekTuesday),
		string(RecurrencePatternDaysOfWeekWednesday),
	}
}

func (s *RecurrencePatternDaysOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternDaysOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternDaysOfWeek(input string) (*RecurrencePatternDaysOfWeek, error) {
	vals := map[string]RecurrencePatternDaysOfWeek{
		"friday":    RecurrencePatternDaysOfWeekFriday,
		"monday":    RecurrencePatternDaysOfWeekMonday,
		"saturday":  RecurrencePatternDaysOfWeekSaturday,
		"sunday":    RecurrencePatternDaysOfWeekSunday,
		"thursday":  RecurrencePatternDaysOfWeekThursday,
		"tuesday":   RecurrencePatternDaysOfWeekTuesday,
		"wednesday": RecurrencePatternDaysOfWeekWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternDaysOfWeek(input)
	return &out, nil
}

type RecurrencePatternFirstDayOfWeek string

const (
	RecurrencePatternFirstDayOfWeekFriday    RecurrencePatternFirstDayOfWeek = "friday"
	RecurrencePatternFirstDayOfWeekMonday    RecurrencePatternFirstDayOfWeek = "monday"
	RecurrencePatternFirstDayOfWeekSaturday  RecurrencePatternFirstDayOfWeek = "saturday"
	RecurrencePatternFirstDayOfWeekSunday    RecurrencePatternFirstDayOfWeek = "sunday"
	RecurrencePatternFirstDayOfWeekThursday  RecurrencePatternFirstDayOfWeek = "thursday"
	RecurrencePatternFirstDayOfWeekTuesday   RecurrencePatternFirstDayOfWeek = "tuesday"
	RecurrencePatternFirstDayOfWeekWednesday RecurrencePatternFirstDayOfWeek = "wednesday"
)

func PossibleValuesForRecurrencePatternFirstDayOfWeek() []string {
	return []string{
		string(RecurrencePatternFirstDayOfWeekFriday),
		string(RecurrencePatternFirstDayOfWeekMonday),
		string(RecurrencePatternFirstDayOfWeekSaturday),
		string(RecurrencePatternFirstDayOfWeekSunday),
		string(RecurrencePatternFirstDayOfWeekThursday),
		string(RecurrencePatternFirstDayOfWeekTuesday),
		string(RecurrencePatternFirstDayOfWeekWednesday),
	}
}

func (s *RecurrencePatternFirstDayOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternFirstDayOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternFirstDayOfWeek(input string) (*RecurrencePatternFirstDayOfWeek, error) {
	vals := map[string]RecurrencePatternFirstDayOfWeek{
		"friday":    RecurrencePatternFirstDayOfWeekFriday,
		"monday":    RecurrencePatternFirstDayOfWeekMonday,
		"saturday":  RecurrencePatternFirstDayOfWeekSaturday,
		"sunday":    RecurrencePatternFirstDayOfWeekSunday,
		"thursday":  RecurrencePatternFirstDayOfWeekThursday,
		"tuesday":   RecurrencePatternFirstDayOfWeekTuesday,
		"wednesday": RecurrencePatternFirstDayOfWeekWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternFirstDayOfWeek(input)
	return &out, nil
}

type RecurrencePatternIndex string

const (
	RecurrencePatternIndexFirst  RecurrencePatternIndex = "first"
	RecurrencePatternIndexFourth RecurrencePatternIndex = "fourth"
	RecurrencePatternIndexLast   RecurrencePatternIndex = "last"
	RecurrencePatternIndexSecond RecurrencePatternIndex = "second"
	RecurrencePatternIndexThird  RecurrencePatternIndex = "third"
)

func PossibleValuesForRecurrencePatternIndex() []string {
	return []string{
		string(RecurrencePatternIndexFirst),
		string(RecurrencePatternIndexFourth),
		string(RecurrencePatternIndexLast),
		string(RecurrencePatternIndexSecond),
		string(RecurrencePatternIndexThird),
	}
}

func (s *RecurrencePatternIndex) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternIndex(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternIndex(input string) (*RecurrencePatternIndex, error) {
	vals := map[string]RecurrencePatternIndex{
		"first":  RecurrencePatternIndexFirst,
		"fourth": RecurrencePatternIndexFourth,
		"last":   RecurrencePatternIndexLast,
		"second": RecurrencePatternIndexSecond,
		"third":  RecurrencePatternIndexThird,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternIndex(input)
	return &out, nil
}

type RecurrencePatternType string

const (
	RecurrencePatternTypeAbsoluteMonthly RecurrencePatternType = "absoluteMonthly"
	RecurrencePatternTypeAbsoluteYearly  RecurrencePatternType = "absoluteYearly"
	RecurrencePatternTypeDaily           RecurrencePatternType = "daily"
	RecurrencePatternTypeRelativeMonthly RecurrencePatternType = "relativeMonthly"
	RecurrencePatternTypeRelativeYearly  RecurrencePatternType = "relativeYearly"
	RecurrencePatternTypeWeekly          RecurrencePatternType = "weekly"
)

func PossibleValuesForRecurrencePatternType() []string {
	return []string{
		string(RecurrencePatternTypeAbsoluteMonthly),
		string(RecurrencePatternTypeAbsoluteYearly),
		string(RecurrencePatternTypeDaily),
		string(RecurrencePatternTypeRelativeMonthly),
		string(RecurrencePatternTypeRelativeYearly),
		string(RecurrencePatternTypeWeekly),
	}
}

func (s *RecurrencePatternType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternType(input string) (*RecurrencePatternType, error) {
	vals := map[string]RecurrencePatternType{
		"absolutemonthly": RecurrencePatternTypeAbsoluteMonthly,
		"absoluteyearly":  RecurrencePatternTypeAbsoluteYearly,
		"daily":           RecurrencePatternTypeDaily,
		"relativemonthly": RecurrencePatternTypeRelativeMonthly,
		"relativeyearly":  RecurrencePatternTypeRelativeYearly,
		"weekly":          RecurrencePatternTypeWeekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternType(input)
	return &out, nil
}

type RecurrenceRangeType string

const (
	RecurrenceRangeTypeEndDate  RecurrenceRangeType = "endDate"
	RecurrenceRangeTypeNoEnd    RecurrenceRangeType = "noEnd"
	RecurrenceRangeTypeNumbered RecurrenceRangeType = "numbered"
)

func PossibleValuesForRecurrenceRangeType() []string {
	return []string{
		string(RecurrenceRangeTypeEndDate),
		string(RecurrenceRangeTypeNoEnd),
		string(RecurrenceRangeTypeNumbered),
	}
}

func (s *RecurrenceRangeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrenceRangeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrenceRangeType(input string) (*RecurrenceRangeType, error) {
	vals := map[string]RecurrenceRangeType{
		"enddate":  RecurrenceRangeTypeEndDate,
		"noend":    RecurrenceRangeTypeNoEnd,
		"numbered": RecurrenceRangeTypeNumbered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceRangeType(input)
	return &out, nil
}

type ResponseStatusResponse string

const (
	ResponseStatusResponseAccepted            ResponseStatusResponse = "accepted"
	ResponseStatusResponseDeclined            ResponseStatusResponse = "declined"
	ResponseStatusResponseNone                ResponseStatusResponse = "none"
	ResponseStatusResponseNotResponded        ResponseStatusResponse = "notResponded"
	ResponseStatusResponseOrganizer           ResponseStatusResponse = "organizer"
	ResponseStatusResponseTentativelyAccepted ResponseStatusResponse = "tentativelyAccepted"
)

func PossibleValuesForResponseStatusResponse() []string {
	return []string{
		string(ResponseStatusResponseAccepted),
		string(ResponseStatusResponseDeclined),
		string(ResponseStatusResponseNone),
		string(ResponseStatusResponseNotResponded),
		string(ResponseStatusResponseOrganizer),
		string(ResponseStatusResponseTentativelyAccepted),
	}
}

func (s *ResponseStatusResponse) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResponseStatusResponse(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResponseStatusResponse(input string) (*ResponseStatusResponse, error) {
	vals := map[string]ResponseStatusResponse{
		"accepted":            ResponseStatusResponseAccepted,
		"declined":            ResponseStatusResponseDeclined,
		"none":                ResponseStatusResponseNone,
		"notresponded":        ResponseStatusResponseNotResponded,
		"organizer":           ResponseStatusResponseOrganizer,
		"tentativelyaccepted": ResponseStatusResponseTentativelyAccepted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResponseStatusResponse(input)
	return &out, nil
}

type RetentionLabelSettingsBehaviorDuringRetentionPeriod string

const (
	RetentionLabelSettingsBehaviorDuringRetentionPeriodDoNotRetain              RetentionLabelSettingsBehaviorDuringRetentionPeriod = "doNotRetain"
	RetentionLabelSettingsBehaviorDuringRetentionPeriodRetain                   RetentionLabelSettingsBehaviorDuringRetentionPeriod = "retain"
	RetentionLabelSettingsBehaviorDuringRetentionPeriodRetainAsRecord           RetentionLabelSettingsBehaviorDuringRetentionPeriod = "retainAsRecord"
	RetentionLabelSettingsBehaviorDuringRetentionPeriodRetainAsRegulatoryRecord RetentionLabelSettingsBehaviorDuringRetentionPeriod = "retainAsRegulatoryRecord"
)

func PossibleValuesForRetentionLabelSettingsBehaviorDuringRetentionPeriod() []string {
	return []string{
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriodDoNotRetain),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriodRetain),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriodRetainAsRecord),
		string(RetentionLabelSettingsBehaviorDuringRetentionPeriodRetainAsRegulatoryRecord),
	}
}

func (s *RetentionLabelSettingsBehaviorDuringRetentionPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetentionLabelSettingsBehaviorDuringRetentionPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetentionLabelSettingsBehaviorDuringRetentionPeriod(input string) (*RetentionLabelSettingsBehaviorDuringRetentionPeriod, error) {
	vals := map[string]RetentionLabelSettingsBehaviorDuringRetentionPeriod{
		"donotretain":              RetentionLabelSettingsBehaviorDuringRetentionPeriodDoNotRetain,
		"retain":                   RetentionLabelSettingsBehaviorDuringRetentionPeriodRetain,
		"retainasrecord":           RetentionLabelSettingsBehaviorDuringRetentionPeriodRetainAsRecord,
		"retainasregulatoryrecord": RetentionLabelSettingsBehaviorDuringRetentionPeriodRetainAsRegulatoryRecord,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetentionLabelSettingsBehaviorDuringRetentionPeriod(input)
	return &out, nil
}

type RichLongRunningOperationStatus string

const (
	RichLongRunningOperationStatusFailed     RichLongRunningOperationStatus = "failed"
	RichLongRunningOperationStatusNotStarted RichLongRunningOperationStatus = "notStarted"
	RichLongRunningOperationStatusRunning    RichLongRunningOperationStatus = "running"
	RichLongRunningOperationStatusSucceeded  RichLongRunningOperationStatus = "succeeded"
)

func PossibleValuesForRichLongRunningOperationStatus() []string {
	return []string{
		string(RichLongRunningOperationStatusFailed),
		string(RichLongRunningOperationStatusNotStarted),
		string(RichLongRunningOperationStatusRunning),
		string(RichLongRunningOperationStatusSucceeded),
	}
}

func (s *RichLongRunningOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRichLongRunningOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRichLongRunningOperationStatus(input string) (*RichLongRunningOperationStatus, error) {
	vals := map[string]RichLongRunningOperationStatus{
		"failed":     RichLongRunningOperationStatusFailed,
		"notstarted": RichLongRunningOperationStatusNotStarted,
		"running":    RichLongRunningOperationStatusRunning,
		"succeeded":  RichLongRunningOperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RichLongRunningOperationStatus(input)
	return &out, nil
}

type ScheduleProvisionStatus string

const (
	ScheduleProvisionStatusCompleted  ScheduleProvisionStatus = "Completed"
	ScheduleProvisionStatusFailed     ScheduleProvisionStatus = "Failed"
	ScheduleProvisionStatusNotStarted ScheduleProvisionStatus = "NotStarted"
	ScheduleProvisionStatusRunning    ScheduleProvisionStatus = "Running"
)

func PossibleValuesForScheduleProvisionStatus() []string {
	return []string{
		string(ScheduleProvisionStatusCompleted),
		string(ScheduleProvisionStatusFailed),
		string(ScheduleProvisionStatusNotStarted),
		string(ScheduleProvisionStatusRunning),
	}
}

func (s *ScheduleProvisionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleProvisionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleProvisionStatus(input string) (*ScheduleProvisionStatus, error) {
	vals := map[string]ScheduleProvisionStatus{
		"completed":  ScheduleProvisionStatusCompleted,
		"failed":     ScheduleProvisionStatusFailed,
		"notstarted": ScheduleProvisionStatusNotStarted,
		"running":    ScheduleProvisionStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleProvisionStatus(input)
	return &out, nil
}

type ScoredEmailAddressSelectionLikelihood string

const (
	ScoredEmailAddressSelectionLikelihoodHigh         ScoredEmailAddressSelectionLikelihood = "high"
	ScoredEmailAddressSelectionLikelihoodNotSpecified ScoredEmailAddressSelectionLikelihood = "notSpecified"
)

func PossibleValuesForScoredEmailAddressSelectionLikelihood() []string {
	return []string{
		string(ScoredEmailAddressSelectionLikelihoodHigh),
		string(ScoredEmailAddressSelectionLikelihoodNotSpecified),
	}
}

func (s *ScoredEmailAddressSelectionLikelihood) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScoredEmailAddressSelectionLikelihood(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScoredEmailAddressSelectionLikelihood(input string) (*ScoredEmailAddressSelectionLikelihood, error) {
	vals := map[string]ScoredEmailAddressSelectionLikelihood{
		"high":         ScoredEmailAddressSelectionLikelihoodHigh,
		"notspecified": ScoredEmailAddressSelectionLikelihoodNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScoredEmailAddressSelectionLikelihood(input)
	return &out, nil
}

type SettingSourceSourceType string

const (
	SettingSourceSourceTypeDeviceConfiguration SettingSourceSourceType = "deviceConfiguration"
	SettingSourceSourceTypeDeviceIntent        SettingSourceSourceType = "deviceIntent"
)

func PossibleValuesForSettingSourceSourceType() []string {
	return []string{
		string(SettingSourceSourceTypeDeviceConfiguration),
		string(SettingSourceSourceTypeDeviceIntent),
	}
}

func (s *SettingSourceSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSettingSourceSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSettingSourceSourceType(input string) (*SettingSourceSourceType, error) {
	vals := map[string]SettingSourceSourceType{
		"deviceconfiguration": SettingSourceSourceTypeDeviceConfiguration,
		"deviceintent":        SettingSourceSourceTypeDeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingSourceSourceType(input)
	return &out, nil
}

type ShiftActivityTheme string

const (
	ShiftActivityThemeBlue       ShiftActivityTheme = "blue"
	ShiftActivityThemeDarkBlue   ShiftActivityTheme = "darkBlue"
	ShiftActivityThemeDarkGreen  ShiftActivityTheme = "darkGreen"
	ShiftActivityThemeDarkPink   ShiftActivityTheme = "darkPink"
	ShiftActivityThemeDarkPurple ShiftActivityTheme = "darkPurple"
	ShiftActivityThemeDarkYellow ShiftActivityTheme = "darkYellow"
	ShiftActivityThemeGray       ShiftActivityTheme = "gray"
	ShiftActivityThemeGreen      ShiftActivityTheme = "green"
	ShiftActivityThemePink       ShiftActivityTheme = "pink"
	ShiftActivityThemePurple     ShiftActivityTheme = "purple"
	ShiftActivityThemeWhite      ShiftActivityTheme = "white"
	ShiftActivityThemeYellow     ShiftActivityTheme = "yellow"
)

func PossibleValuesForShiftActivityTheme() []string {
	return []string{
		string(ShiftActivityThemeBlue),
		string(ShiftActivityThemeDarkBlue),
		string(ShiftActivityThemeDarkGreen),
		string(ShiftActivityThemeDarkPink),
		string(ShiftActivityThemeDarkPurple),
		string(ShiftActivityThemeDarkYellow),
		string(ShiftActivityThemeGray),
		string(ShiftActivityThemeGreen),
		string(ShiftActivityThemePink),
		string(ShiftActivityThemePurple),
		string(ShiftActivityThemeWhite),
		string(ShiftActivityThemeYellow),
	}
}

func (s *ShiftActivityTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShiftActivityTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShiftActivityTheme(input string) (*ShiftActivityTheme, error) {
	vals := map[string]ShiftActivityTheme{
		"blue":       ShiftActivityThemeBlue,
		"darkblue":   ShiftActivityThemeDarkBlue,
		"darkgreen":  ShiftActivityThemeDarkGreen,
		"darkpink":   ShiftActivityThemeDarkPink,
		"darkpurple": ShiftActivityThemeDarkPurple,
		"darkyellow": ShiftActivityThemeDarkYellow,
		"gray":       ShiftActivityThemeGray,
		"green":      ShiftActivityThemeGreen,
		"pink":       ShiftActivityThemePink,
		"purple":     ShiftActivityThemePurple,
		"white":      ShiftActivityThemeWhite,
		"yellow":     ShiftActivityThemeYellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShiftActivityTheme(input)
	return &out, nil
}

type ShiftItemTheme string

const (
	ShiftItemThemeBlue       ShiftItemTheme = "blue"
	ShiftItemThemeDarkBlue   ShiftItemTheme = "darkBlue"
	ShiftItemThemeDarkGreen  ShiftItemTheme = "darkGreen"
	ShiftItemThemeDarkPink   ShiftItemTheme = "darkPink"
	ShiftItemThemeDarkPurple ShiftItemTheme = "darkPurple"
	ShiftItemThemeDarkYellow ShiftItemTheme = "darkYellow"
	ShiftItemThemeGray       ShiftItemTheme = "gray"
	ShiftItemThemeGreen      ShiftItemTheme = "green"
	ShiftItemThemePink       ShiftItemTheme = "pink"
	ShiftItemThemePurple     ShiftItemTheme = "purple"
	ShiftItemThemeWhite      ShiftItemTheme = "white"
	ShiftItemThemeYellow     ShiftItemTheme = "yellow"
)

func PossibleValuesForShiftItemTheme() []string {
	return []string{
		string(ShiftItemThemeBlue),
		string(ShiftItemThemeDarkBlue),
		string(ShiftItemThemeDarkGreen),
		string(ShiftItemThemeDarkPink),
		string(ShiftItemThemeDarkPurple),
		string(ShiftItemThemeDarkYellow),
		string(ShiftItemThemeGray),
		string(ShiftItemThemeGreen),
		string(ShiftItemThemePink),
		string(ShiftItemThemePurple),
		string(ShiftItemThemeWhite),
		string(ShiftItemThemeYellow),
	}
}

func (s *ShiftItemTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShiftItemTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShiftItemTheme(input string) (*ShiftItemTheme, error) {
	vals := map[string]ShiftItemTheme{
		"blue":       ShiftItemThemeBlue,
		"darkblue":   ShiftItemThemeDarkBlue,
		"darkgreen":  ShiftItemThemeDarkGreen,
		"darkpink":   ShiftItemThemeDarkPink,
		"darkpurple": ShiftItemThemeDarkPurple,
		"darkyellow": ShiftItemThemeDarkYellow,
		"gray":       ShiftItemThemeGray,
		"green":      ShiftItemThemeGreen,
		"pink":       ShiftItemThemePink,
		"purple":     ShiftItemThemePurple,
		"white":      ShiftItemThemeWhite,
		"yellow":     ShiftItemThemeYellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShiftItemTheme(input)
	return &out, nil
}

type SwapShiftsChangeRequestAssignedTo string

const (
	SwapShiftsChangeRequestAssignedToManager   SwapShiftsChangeRequestAssignedTo = "manager"
	SwapShiftsChangeRequestAssignedToRecipient SwapShiftsChangeRequestAssignedTo = "recipient"
	SwapShiftsChangeRequestAssignedToSender    SwapShiftsChangeRequestAssignedTo = "sender"
	SwapShiftsChangeRequestAssignedToSystem    SwapShiftsChangeRequestAssignedTo = "system"
)

func PossibleValuesForSwapShiftsChangeRequestAssignedTo() []string {
	return []string{
		string(SwapShiftsChangeRequestAssignedToManager),
		string(SwapShiftsChangeRequestAssignedToRecipient),
		string(SwapShiftsChangeRequestAssignedToSender),
		string(SwapShiftsChangeRequestAssignedToSystem),
	}
}

func (s *SwapShiftsChangeRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSwapShiftsChangeRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSwapShiftsChangeRequestAssignedTo(input string) (*SwapShiftsChangeRequestAssignedTo, error) {
	vals := map[string]SwapShiftsChangeRequestAssignedTo{
		"manager":   SwapShiftsChangeRequestAssignedToManager,
		"recipient": SwapShiftsChangeRequestAssignedToRecipient,
		"sender":    SwapShiftsChangeRequestAssignedToSender,
		"system":    SwapShiftsChangeRequestAssignedToSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SwapShiftsChangeRequestAssignedTo(input)
	return &out, nil
}

type SwapShiftsChangeRequestState string

const (
	SwapShiftsChangeRequestStateApproved SwapShiftsChangeRequestState = "approved"
	SwapShiftsChangeRequestStateDeclined SwapShiftsChangeRequestState = "declined"
	SwapShiftsChangeRequestStatePending  SwapShiftsChangeRequestState = "pending"
)

func PossibleValuesForSwapShiftsChangeRequestState() []string {
	return []string{
		string(SwapShiftsChangeRequestStateApproved),
		string(SwapShiftsChangeRequestStateDeclined),
		string(SwapShiftsChangeRequestStatePending),
	}
}

func (s *SwapShiftsChangeRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSwapShiftsChangeRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSwapShiftsChangeRequestState(input string) (*SwapShiftsChangeRequestState, error) {
	vals := map[string]SwapShiftsChangeRequestState{
		"approved": SwapShiftsChangeRequestStateApproved,
		"declined": SwapShiftsChangeRequestStateDeclined,
		"pending":  SwapShiftsChangeRequestStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SwapShiftsChangeRequestState(input)
	return &out, nil
}

type TeamFunSettingsGiphyContentRating string

const (
	TeamFunSettingsGiphyContentRatingModerate TeamFunSettingsGiphyContentRating = "moderate"
	TeamFunSettingsGiphyContentRatingStrict   TeamFunSettingsGiphyContentRating = "strict"
)

func PossibleValuesForTeamFunSettingsGiphyContentRating() []string {
	return []string{
		string(TeamFunSettingsGiphyContentRatingModerate),
		string(TeamFunSettingsGiphyContentRatingStrict),
	}
}

func (s *TeamFunSettingsGiphyContentRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamFunSettingsGiphyContentRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamFunSettingsGiphyContentRating(input string) (*TeamFunSettingsGiphyContentRating, error) {
	vals := map[string]TeamFunSettingsGiphyContentRating{
		"moderate": TeamFunSettingsGiphyContentRatingModerate,
		"strict":   TeamFunSettingsGiphyContentRatingStrict,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamFunSettingsGiphyContentRating(input)
	return &out, nil
}

type TeamSpecialization string

const (
	TeamSpecializationEducationClass                         TeamSpecialization = "educationClass"
	TeamSpecializationEducationProfessionalLearningCommunity TeamSpecialization = "educationProfessionalLearningCommunity"
	TeamSpecializationEducationStaff                         TeamSpecialization = "educationStaff"
	TeamSpecializationEducationStandard                      TeamSpecialization = "educationStandard"
	TeamSpecializationHealthcareCareCoordination             TeamSpecialization = "healthcareCareCoordination"
	TeamSpecializationHealthcareStandard                     TeamSpecialization = "healthcareStandard"
	TeamSpecializationNone                                   TeamSpecialization = "none"
)

func PossibleValuesForTeamSpecialization() []string {
	return []string{
		string(TeamSpecializationEducationClass),
		string(TeamSpecializationEducationProfessionalLearningCommunity),
		string(TeamSpecializationEducationStaff),
		string(TeamSpecializationEducationStandard),
		string(TeamSpecializationHealthcareCareCoordination),
		string(TeamSpecializationHealthcareStandard),
		string(TeamSpecializationNone),
	}
}

func (s *TeamSpecialization) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamSpecialization(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamSpecialization(input string) (*TeamSpecialization, error) {
	vals := map[string]TeamSpecialization{
		"educationclass":                         TeamSpecializationEducationClass,
		"educationprofessionallearningcommunity": TeamSpecializationEducationProfessionalLearningCommunity,
		"educationstaff":                         TeamSpecializationEducationStaff,
		"educationstandard":                      TeamSpecializationEducationStandard,
		"healthcarecarecoordination":             TeamSpecializationHealthcareCareCoordination,
		"healthcarestandard":                     TeamSpecializationHealthcareStandard,
		"none":                                   TeamSpecializationNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamSpecialization(input)
	return &out, nil
}

type TeamVisibility string

const (
	TeamVisibilityHiddenMembership TeamVisibility = "hiddenMembership"
	TeamVisibilityPrivate          TeamVisibility = "private"
	TeamVisibilityPublic           TeamVisibility = "public"
)

func PossibleValuesForTeamVisibility() []string {
	return []string{
		string(TeamVisibilityHiddenMembership),
		string(TeamVisibilityPrivate),
		string(TeamVisibilityPublic),
	}
}

func (s *TeamVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamVisibility(input string) (*TeamVisibility, error) {
	vals := map[string]TeamVisibility{
		"hiddenmembership": TeamVisibilityHiddenMembership,
		"private":          TeamVisibilityPrivate,
		"public":           TeamVisibilityPublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamVisibility(input)
	return &out, nil
}

type TeamsAppDefinitionPublishingState string

const (
	TeamsAppDefinitionPublishingStatePublished TeamsAppDefinitionPublishingState = "published"
	TeamsAppDefinitionPublishingStateRejected  TeamsAppDefinitionPublishingState = "rejected"
	TeamsAppDefinitionPublishingStateSubmitted TeamsAppDefinitionPublishingState = "submitted"
)

func PossibleValuesForTeamsAppDefinitionPublishingState() []string {
	return []string{
		string(TeamsAppDefinitionPublishingStatePublished),
		string(TeamsAppDefinitionPublishingStateRejected),
		string(TeamsAppDefinitionPublishingStateSubmitted),
	}
}

func (s *TeamsAppDefinitionPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDefinitionPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDefinitionPublishingState(input string) (*TeamsAppDefinitionPublishingState, error) {
	vals := map[string]TeamsAppDefinitionPublishingState{
		"published": TeamsAppDefinitionPublishingStatePublished,
		"rejected":  TeamsAppDefinitionPublishingStateRejected,
		"submitted": TeamsAppDefinitionPublishingStateSubmitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDefinitionPublishingState(input)
	return &out, nil
}

type TeamsAppDistributionMethod string

const (
	TeamsAppDistributionMethodOrganization TeamsAppDistributionMethod = "organization"
	TeamsAppDistributionMethodSideloaded   TeamsAppDistributionMethod = "sideloaded"
	TeamsAppDistributionMethodStore        TeamsAppDistributionMethod = "store"
)

func PossibleValuesForTeamsAppDistributionMethod() []string {
	return []string{
		string(TeamsAppDistributionMethodOrganization),
		string(TeamsAppDistributionMethodSideloaded),
		string(TeamsAppDistributionMethodStore),
	}
}

func (s *TeamsAppDistributionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDistributionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDistributionMethod(input string) (*TeamsAppDistributionMethod, error) {
	vals := map[string]TeamsAppDistributionMethod{
		"organization": TeamsAppDistributionMethodOrganization,
		"sideloaded":   TeamsAppDistributionMethodSideloaded,
		"store":        TeamsAppDistributionMethodStore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDistributionMethod(input)
	return &out, nil
}

type TeamsAppResourceSpecificPermissionPermissionType string

const (
	TeamsAppResourceSpecificPermissionPermissionTypeApplication TeamsAppResourceSpecificPermissionPermissionType = "application"
	TeamsAppResourceSpecificPermissionPermissionTypeDelegated   TeamsAppResourceSpecificPermissionPermissionType = "delegated"
)

func PossibleValuesForTeamsAppResourceSpecificPermissionPermissionType() []string {
	return []string{
		string(TeamsAppResourceSpecificPermissionPermissionTypeApplication),
		string(TeamsAppResourceSpecificPermissionPermissionTypeDelegated),
	}
}

func (s *TeamsAppResourceSpecificPermissionPermissionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppResourceSpecificPermissionPermissionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppResourceSpecificPermissionPermissionType(input string) (*TeamsAppResourceSpecificPermissionPermissionType, error) {
	vals := map[string]TeamsAppResourceSpecificPermissionPermissionType{
		"application": TeamsAppResourceSpecificPermissionPermissionTypeApplication,
		"delegated":   TeamsAppResourceSpecificPermissionPermissionTypeDelegated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppResourceSpecificPermissionPermissionType(input)
	return &out, nil
}

type TeamsAsyncOperationOperationType string

const (
	TeamsAsyncOperationOperationTypeArchiveTeam   TeamsAsyncOperationOperationType = "archiveTeam"
	TeamsAsyncOperationOperationTypeCloneTeam     TeamsAsyncOperationOperationType = "cloneTeam"
	TeamsAsyncOperationOperationTypeCreateChannel TeamsAsyncOperationOperationType = "createChannel"
	TeamsAsyncOperationOperationTypeCreateTeam    TeamsAsyncOperationOperationType = "createTeam"
	TeamsAsyncOperationOperationTypeInvalid       TeamsAsyncOperationOperationType = "invalid"
	TeamsAsyncOperationOperationTypeTeamifyGroup  TeamsAsyncOperationOperationType = "teamifyGroup"
	TeamsAsyncOperationOperationTypeUnarchiveTeam TeamsAsyncOperationOperationType = "unarchiveTeam"
)

func PossibleValuesForTeamsAsyncOperationOperationType() []string {
	return []string{
		string(TeamsAsyncOperationOperationTypeArchiveTeam),
		string(TeamsAsyncOperationOperationTypeCloneTeam),
		string(TeamsAsyncOperationOperationTypeCreateChannel),
		string(TeamsAsyncOperationOperationTypeCreateTeam),
		string(TeamsAsyncOperationOperationTypeInvalid),
		string(TeamsAsyncOperationOperationTypeTeamifyGroup),
		string(TeamsAsyncOperationOperationTypeUnarchiveTeam),
	}
}

func (s *TeamsAsyncOperationOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAsyncOperationOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAsyncOperationOperationType(input string) (*TeamsAsyncOperationOperationType, error) {
	vals := map[string]TeamsAsyncOperationOperationType{
		"archiveteam":   TeamsAsyncOperationOperationTypeArchiveTeam,
		"cloneteam":     TeamsAsyncOperationOperationTypeCloneTeam,
		"createchannel": TeamsAsyncOperationOperationTypeCreateChannel,
		"createteam":    TeamsAsyncOperationOperationTypeCreateTeam,
		"invalid":       TeamsAsyncOperationOperationTypeInvalid,
		"teamifygroup":  TeamsAsyncOperationOperationTypeTeamifyGroup,
		"unarchiveteam": TeamsAsyncOperationOperationTypeUnarchiveTeam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAsyncOperationOperationType(input)
	return &out, nil
}

type TeamsAsyncOperationStatus string

const (
	TeamsAsyncOperationStatusFailed     TeamsAsyncOperationStatus = "failed"
	TeamsAsyncOperationStatusInProgress TeamsAsyncOperationStatus = "inProgress"
	TeamsAsyncOperationStatusInvalid    TeamsAsyncOperationStatus = "invalid"
	TeamsAsyncOperationStatusNotStarted TeamsAsyncOperationStatus = "notStarted"
	TeamsAsyncOperationStatusSucceeded  TeamsAsyncOperationStatus = "succeeded"
)

func PossibleValuesForTeamsAsyncOperationStatus() []string {
	return []string{
		string(TeamsAsyncOperationStatusFailed),
		string(TeamsAsyncOperationStatusInProgress),
		string(TeamsAsyncOperationStatusInvalid),
		string(TeamsAsyncOperationStatusNotStarted),
		string(TeamsAsyncOperationStatusSucceeded),
	}
}

func (s *TeamsAsyncOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAsyncOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAsyncOperationStatus(input string) (*TeamsAsyncOperationStatus, error) {
	vals := map[string]TeamsAsyncOperationStatus{
		"failed":     TeamsAsyncOperationStatusFailed,
		"inprogress": TeamsAsyncOperationStatusInProgress,
		"invalid":    TeamsAsyncOperationStatusInvalid,
		"notstarted": TeamsAsyncOperationStatusNotStarted,
		"succeeded":  TeamsAsyncOperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAsyncOperationStatus(input)
	return &out, nil
}

type TeamworkConversationIdentityConversationIdentityType string

const (
	TeamworkConversationIdentityConversationIdentityTypeChannel TeamworkConversationIdentityConversationIdentityType = "channel"
	TeamworkConversationIdentityConversationIdentityTypeChat    TeamworkConversationIdentityConversationIdentityType = "chat"
	TeamworkConversationIdentityConversationIdentityTypeTeam    TeamworkConversationIdentityConversationIdentityType = "team"
)

func PossibleValuesForTeamworkConversationIdentityConversationIdentityType() []string {
	return []string{
		string(TeamworkConversationIdentityConversationIdentityTypeChannel),
		string(TeamworkConversationIdentityConversationIdentityTypeChat),
		string(TeamworkConversationIdentityConversationIdentityTypeTeam),
	}
}

func (s *TeamworkConversationIdentityConversationIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkConversationIdentityConversationIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkConversationIdentityConversationIdentityType(input string) (*TeamworkConversationIdentityConversationIdentityType, error) {
	vals := map[string]TeamworkConversationIdentityConversationIdentityType{
		"channel": TeamworkConversationIdentityConversationIdentityTypeChannel,
		"chat":    TeamworkConversationIdentityConversationIdentityTypeChat,
		"team":    TeamworkConversationIdentityConversationIdentityTypeTeam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkConversationIdentityConversationIdentityType(input)
	return &out, nil
}

type TeamworkTagTagType string

const (
	TeamworkTagTagTypeStandard TeamworkTagTagType = "standard"
)

func PossibleValuesForTeamworkTagTagType() []string {
	return []string{
		string(TeamworkTagTagTypeStandard),
	}
}

func (s *TeamworkTagTagType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkTagTagType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkTagTagType(input string) (*TeamworkTagTagType, error) {
	vals := map[string]TeamworkTagTagType{
		"standard": TeamworkTagTagTypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkTagTagType(input)
	return &out, nil
}

type TeamworkUserIdentityUserIdentityType string

const (
	TeamworkUserIdentityUserIdentityTypeAadUser                      TeamworkUserIdentityUserIdentityType = "aadUser"
	TeamworkUserIdentityUserIdentityTypeAnonymousGuest               TeamworkUserIdentityUserIdentityType = "anonymousGuest"
	TeamworkUserIdentityUserIdentityTypeEmailUser                    TeamworkUserIdentityUserIdentityType = "emailUser"
	TeamworkUserIdentityUserIdentityTypeFederatedUser                TeamworkUserIdentityUserIdentityType = "federatedUser"
	TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser             TeamworkUserIdentityUserIdentityType = "onPremiseAadUser"
	TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser TeamworkUserIdentityUserIdentityType = "personalMicrosoftAccountUser"
	TeamworkUserIdentityUserIdentityTypePhoneUser                    TeamworkUserIdentityUserIdentityType = "phoneUser"
	TeamworkUserIdentityUserIdentityTypeSkypeUser                    TeamworkUserIdentityUserIdentityType = "skypeUser"
)

func PossibleValuesForTeamworkUserIdentityUserIdentityType() []string {
	return []string{
		string(TeamworkUserIdentityUserIdentityTypeAadUser),
		string(TeamworkUserIdentityUserIdentityTypeAnonymousGuest),
		string(TeamworkUserIdentityUserIdentityTypeEmailUser),
		string(TeamworkUserIdentityUserIdentityTypeFederatedUser),
		string(TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser),
		string(TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser),
		string(TeamworkUserIdentityUserIdentityTypePhoneUser),
		string(TeamworkUserIdentityUserIdentityTypeSkypeUser),
	}
}

func (s *TeamworkUserIdentityUserIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkUserIdentityUserIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkUserIdentityUserIdentityType(input string) (*TeamworkUserIdentityUserIdentityType, error) {
	vals := map[string]TeamworkUserIdentityUserIdentityType{
		"aaduser":                      TeamworkUserIdentityUserIdentityTypeAadUser,
		"anonymousguest":               TeamworkUserIdentityUserIdentityTypeAnonymousGuest,
		"emailuser":                    TeamworkUserIdentityUserIdentityTypeEmailUser,
		"federateduser":                TeamworkUserIdentityUserIdentityTypeFederatedUser,
		"onpremiseaaduser":             TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser,
		"personalmicrosoftaccountuser": TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser,
		"phoneuser":                    TeamworkUserIdentityUserIdentityTypePhoneUser,
		"skypeuser":                    TeamworkUserIdentityUserIdentityTypeSkypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkUserIdentityUserIdentityType(input)
	return &out, nil
}

type TermStoreGroupScope string

const (
	TermStoreGroupScopeGlobal         TermStoreGroupScope = "global"
	TermStoreGroupScopeSiteCollection TermStoreGroupScope = "siteCollection"
	TermStoreGroupScopeSystem         TermStoreGroupScope = "system"
)

func PossibleValuesForTermStoreGroupScope() []string {
	return []string{
		string(TermStoreGroupScopeGlobal),
		string(TermStoreGroupScopeSiteCollection),
		string(TermStoreGroupScopeSystem),
	}
}

func (s *TermStoreGroupScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTermStoreGroupScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTermStoreGroupScope(input string) (*TermStoreGroupScope, error) {
	vals := map[string]TermStoreGroupScope{
		"global":         TermStoreGroupScopeGlobal,
		"sitecollection": TermStoreGroupScopeSiteCollection,
		"system":         TermStoreGroupScopeSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TermStoreGroupScope(input)
	return &out, nil
}

type TermStoreRelationRelationship string

const (
	TermStoreRelationRelationshipPin   TermStoreRelationRelationship = "pin"
	TermStoreRelationRelationshipReuse TermStoreRelationRelationship = "reuse"
)

func PossibleValuesForTermStoreRelationRelationship() []string {
	return []string{
		string(TermStoreRelationRelationshipPin),
		string(TermStoreRelationRelationshipReuse),
	}
}

func (s *TermStoreRelationRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTermStoreRelationRelationship(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTermStoreRelationRelationship(input string) (*TermStoreRelationRelationship, error) {
	vals := map[string]TermStoreRelationRelationship{
		"pin":   TermStoreRelationRelationshipPin,
		"reuse": TermStoreRelationRelationshipReuse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TermStoreRelationRelationship(input)
	return &out, nil
}

type TimeOffItemTheme string

const (
	TimeOffItemThemeBlue       TimeOffItemTheme = "blue"
	TimeOffItemThemeDarkBlue   TimeOffItemTheme = "darkBlue"
	TimeOffItemThemeDarkGreen  TimeOffItemTheme = "darkGreen"
	TimeOffItemThemeDarkPink   TimeOffItemTheme = "darkPink"
	TimeOffItemThemeDarkPurple TimeOffItemTheme = "darkPurple"
	TimeOffItemThemeDarkYellow TimeOffItemTheme = "darkYellow"
	TimeOffItemThemeGray       TimeOffItemTheme = "gray"
	TimeOffItemThemeGreen      TimeOffItemTheme = "green"
	TimeOffItemThemePink       TimeOffItemTheme = "pink"
	TimeOffItemThemePurple     TimeOffItemTheme = "purple"
	TimeOffItemThemeWhite      TimeOffItemTheme = "white"
	TimeOffItemThemeYellow     TimeOffItemTheme = "yellow"
)

func PossibleValuesForTimeOffItemTheme() []string {
	return []string{
		string(TimeOffItemThemeBlue),
		string(TimeOffItemThemeDarkBlue),
		string(TimeOffItemThemeDarkGreen),
		string(TimeOffItemThemeDarkPink),
		string(TimeOffItemThemeDarkPurple),
		string(TimeOffItemThemeDarkYellow),
		string(TimeOffItemThemeGray),
		string(TimeOffItemThemeGreen),
		string(TimeOffItemThemePink),
		string(TimeOffItemThemePurple),
		string(TimeOffItemThemeWhite),
		string(TimeOffItemThemeYellow),
	}
}

func (s *TimeOffItemTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffItemTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffItemTheme(input string) (*TimeOffItemTheme, error) {
	vals := map[string]TimeOffItemTheme{
		"blue":       TimeOffItemThemeBlue,
		"darkblue":   TimeOffItemThemeDarkBlue,
		"darkgreen":  TimeOffItemThemeDarkGreen,
		"darkpink":   TimeOffItemThemeDarkPink,
		"darkpurple": TimeOffItemThemeDarkPurple,
		"darkyellow": TimeOffItemThemeDarkYellow,
		"gray":       TimeOffItemThemeGray,
		"green":      TimeOffItemThemeGreen,
		"pink":       TimeOffItemThemePink,
		"purple":     TimeOffItemThemePurple,
		"white":      TimeOffItemThemeWhite,
		"yellow":     TimeOffItemThemeYellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffItemTheme(input)
	return &out, nil
}

type TimeOffReasonIconType string

const (
	TimeOffReasonIconTypeCake        TimeOffReasonIconType = "cake"
	TimeOffReasonIconTypeCalendar    TimeOffReasonIconType = "calendar"
	TimeOffReasonIconTypeCar         TimeOffReasonIconType = "car"
	TimeOffReasonIconTypeClock       TimeOffReasonIconType = "clock"
	TimeOffReasonIconTypeCup         TimeOffReasonIconType = "cup"
	TimeOffReasonIconTypeDoctor      TimeOffReasonIconType = "doctor"
	TimeOffReasonIconTypeDog         TimeOffReasonIconType = "dog"
	TimeOffReasonIconTypeFirstAid    TimeOffReasonIconType = "firstAid"
	TimeOffReasonIconTypeGlobe       TimeOffReasonIconType = "globe"
	TimeOffReasonIconTypeJuryDuty    TimeOffReasonIconType = "juryDuty"
	TimeOffReasonIconTypeNone        TimeOffReasonIconType = "none"
	TimeOffReasonIconTypeNotWorking  TimeOffReasonIconType = "notWorking"
	TimeOffReasonIconTypePhone       TimeOffReasonIconType = "phone"
	TimeOffReasonIconTypePiggyBank   TimeOffReasonIconType = "piggyBank"
	TimeOffReasonIconTypePin         TimeOffReasonIconType = "pin"
	TimeOffReasonIconTypePlane       TimeOffReasonIconType = "plane"
	TimeOffReasonIconTypeRunning     TimeOffReasonIconType = "running"
	TimeOffReasonIconTypeSunny       TimeOffReasonIconType = "sunny"
	TimeOffReasonIconTypeTrafficCone TimeOffReasonIconType = "trafficCone"
	TimeOffReasonIconTypeUmbrella    TimeOffReasonIconType = "umbrella"
	TimeOffReasonIconTypeWeather     TimeOffReasonIconType = "weather"
)

func PossibleValuesForTimeOffReasonIconType() []string {
	return []string{
		string(TimeOffReasonIconTypeCake),
		string(TimeOffReasonIconTypeCalendar),
		string(TimeOffReasonIconTypeCar),
		string(TimeOffReasonIconTypeClock),
		string(TimeOffReasonIconTypeCup),
		string(TimeOffReasonIconTypeDoctor),
		string(TimeOffReasonIconTypeDog),
		string(TimeOffReasonIconTypeFirstAid),
		string(TimeOffReasonIconTypeGlobe),
		string(TimeOffReasonIconTypeJuryDuty),
		string(TimeOffReasonIconTypeNone),
		string(TimeOffReasonIconTypeNotWorking),
		string(TimeOffReasonIconTypePhone),
		string(TimeOffReasonIconTypePiggyBank),
		string(TimeOffReasonIconTypePin),
		string(TimeOffReasonIconTypePlane),
		string(TimeOffReasonIconTypeRunning),
		string(TimeOffReasonIconTypeSunny),
		string(TimeOffReasonIconTypeTrafficCone),
		string(TimeOffReasonIconTypeUmbrella),
		string(TimeOffReasonIconTypeWeather),
	}
}

func (s *TimeOffReasonIconType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffReasonIconType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffReasonIconType(input string) (*TimeOffReasonIconType, error) {
	vals := map[string]TimeOffReasonIconType{
		"cake":        TimeOffReasonIconTypeCake,
		"calendar":    TimeOffReasonIconTypeCalendar,
		"car":         TimeOffReasonIconTypeCar,
		"clock":       TimeOffReasonIconTypeClock,
		"cup":         TimeOffReasonIconTypeCup,
		"doctor":      TimeOffReasonIconTypeDoctor,
		"dog":         TimeOffReasonIconTypeDog,
		"firstaid":    TimeOffReasonIconTypeFirstAid,
		"globe":       TimeOffReasonIconTypeGlobe,
		"juryduty":    TimeOffReasonIconTypeJuryDuty,
		"none":        TimeOffReasonIconTypeNone,
		"notworking":  TimeOffReasonIconTypeNotWorking,
		"phone":       TimeOffReasonIconTypePhone,
		"piggybank":   TimeOffReasonIconTypePiggyBank,
		"pin":         TimeOffReasonIconTypePin,
		"plane":       TimeOffReasonIconTypePlane,
		"running":     TimeOffReasonIconTypeRunning,
		"sunny":       TimeOffReasonIconTypeSunny,
		"trafficcone": TimeOffReasonIconTypeTrafficCone,
		"umbrella":    TimeOffReasonIconTypeUmbrella,
		"weather":     TimeOffReasonIconTypeWeather,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffReasonIconType(input)
	return &out, nil
}

type TimeOffRequestAssignedTo string

const (
	TimeOffRequestAssignedToManager   TimeOffRequestAssignedTo = "manager"
	TimeOffRequestAssignedToRecipient TimeOffRequestAssignedTo = "recipient"
	TimeOffRequestAssignedToSender    TimeOffRequestAssignedTo = "sender"
	TimeOffRequestAssignedToSystem    TimeOffRequestAssignedTo = "system"
)

func PossibleValuesForTimeOffRequestAssignedTo() []string {
	return []string{
		string(TimeOffRequestAssignedToManager),
		string(TimeOffRequestAssignedToRecipient),
		string(TimeOffRequestAssignedToSender),
		string(TimeOffRequestAssignedToSystem),
	}
}

func (s *TimeOffRequestAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffRequestAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffRequestAssignedTo(input string) (*TimeOffRequestAssignedTo, error) {
	vals := map[string]TimeOffRequestAssignedTo{
		"manager":   TimeOffRequestAssignedToManager,
		"recipient": TimeOffRequestAssignedToRecipient,
		"sender":    TimeOffRequestAssignedToSender,
		"system":    TimeOffRequestAssignedToSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffRequestAssignedTo(input)
	return &out, nil
}

type TimeOffRequestState string

const (
	TimeOffRequestStateApproved TimeOffRequestState = "approved"
	TimeOffRequestStateDeclined TimeOffRequestState = "declined"
	TimeOffRequestStatePending  TimeOffRequestState = "pending"
)

func PossibleValuesForTimeOffRequestState() []string {
	return []string{
		string(TimeOffRequestStateApproved),
		string(TimeOffRequestStateDeclined),
		string(TimeOffRequestStatePending),
	}
}

func (s *TimeOffRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffRequestState(input string) (*TimeOffRequestState, error) {
	vals := map[string]TimeOffRequestState{
		"approved": TimeOffRequestStateApproved,
		"declined": TimeOffRequestStateDeclined,
		"pending":  TimeOffRequestStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffRequestState(input)
	return &out, nil
}

type TodoTaskImportance string

const (
	TodoTaskImportanceHigh   TodoTaskImportance = "high"
	TodoTaskImportanceLow    TodoTaskImportance = "low"
	TodoTaskImportanceNormal TodoTaskImportance = "normal"
)

func PossibleValuesForTodoTaskImportance() []string {
	return []string{
		string(TodoTaskImportanceHigh),
		string(TodoTaskImportanceLow),
		string(TodoTaskImportanceNormal),
	}
}

func (s *TodoTaskImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTodoTaskImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTodoTaskImportance(input string) (*TodoTaskImportance, error) {
	vals := map[string]TodoTaskImportance{
		"high":   TodoTaskImportanceHigh,
		"low":    TodoTaskImportanceLow,
		"normal": TodoTaskImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskImportance(input)
	return &out, nil
}

type TodoTaskListWellknownListName string

const (
	TodoTaskListWellknownListNameDefaultList   TodoTaskListWellknownListName = "defaultList"
	TodoTaskListWellknownListNameFlaggedEmails TodoTaskListWellknownListName = "flaggedEmails"
	TodoTaskListWellknownListNameNone          TodoTaskListWellknownListName = "none"
)

func PossibleValuesForTodoTaskListWellknownListName() []string {
	return []string{
		string(TodoTaskListWellknownListNameDefaultList),
		string(TodoTaskListWellknownListNameFlaggedEmails),
		string(TodoTaskListWellknownListNameNone),
	}
}

func (s *TodoTaskListWellknownListName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTodoTaskListWellknownListName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTodoTaskListWellknownListName(input string) (*TodoTaskListWellknownListName, error) {
	vals := map[string]TodoTaskListWellknownListName{
		"defaultlist":   TodoTaskListWellknownListNameDefaultList,
		"flaggedemails": TodoTaskListWellknownListNameFlaggedEmails,
		"none":          TodoTaskListWellknownListNameNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskListWellknownListName(input)
	return &out, nil
}

type TodoTaskStatus string

const (
	TodoTaskStatusCompleted       TodoTaskStatus = "completed"
	TodoTaskStatusDeferred        TodoTaskStatus = "deferred"
	TodoTaskStatusInProgress      TodoTaskStatus = "inProgress"
	TodoTaskStatusNotStarted      TodoTaskStatus = "notStarted"
	TodoTaskStatusWaitingOnOthers TodoTaskStatus = "waitingOnOthers"
)

func PossibleValuesForTodoTaskStatus() []string {
	return []string{
		string(TodoTaskStatusCompleted),
		string(TodoTaskStatusDeferred),
		string(TodoTaskStatusInProgress),
		string(TodoTaskStatusNotStarted),
		string(TodoTaskStatusWaitingOnOthers),
	}
}

func (s *TodoTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTodoTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTodoTaskStatus(input string) (*TodoTaskStatus, error) {
	vals := map[string]TodoTaskStatus{
		"completed":       TodoTaskStatusCompleted,
		"deferred":        TodoTaskStatusDeferred,
		"inprogress":      TodoTaskStatusInProgress,
		"notstarted":      TodoTaskStatusNotStarted,
		"waitingonothers": TodoTaskStatusWaitingOnOthers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskStatus(input)
	return &out, nil
}

type UserActivityStatus string

const (
	UserActivityStatusActive  UserActivityStatus = "active"
	UserActivityStatusDeleted UserActivityStatus = "deleted"
	UserActivityStatusIgnored UserActivityStatus = "ignored"
	UserActivityStatusUpdated UserActivityStatus = "updated"
)

func PossibleValuesForUserActivityStatus() []string {
	return []string{
		string(UserActivityStatusActive),
		string(UserActivityStatusDeleted),
		string(UserActivityStatusIgnored),
		string(UserActivityStatusUpdated),
	}
}

func (s *UserActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserActivityStatus(input string) (*UserActivityStatus, error) {
	vals := map[string]UserActivityStatus{
		"active":  UserActivityStatusActive,
		"deleted": UserActivityStatusDeleted,
		"ignored": UserActivityStatusIgnored,
		"updated": UserActivityStatusUpdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserActivityStatus(input)
	return &out, nil
}

type WebsiteType string

const (
	WebsiteTypeBlog    WebsiteType = "blog"
	WebsiteTypeHome    WebsiteType = "home"
	WebsiteTypeOther   WebsiteType = "other"
	WebsiteTypeProfile WebsiteType = "profile"
	WebsiteTypeWork    WebsiteType = "work"
)

func PossibleValuesForWebsiteType() []string {
	return []string{
		string(WebsiteTypeBlog),
		string(WebsiteTypeHome),
		string(WebsiteTypeOther),
		string(WebsiteTypeProfile),
		string(WebsiteTypeWork),
	}
}

func (s *WebsiteType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebsiteType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebsiteType(input string) (*WebsiteType, error) {
	vals := map[string]WebsiteType{
		"blog":    WebsiteTypeBlog,
		"home":    WebsiteTypeHome,
		"other":   WebsiteTypeOther,
		"profile": WebsiteTypeProfile,
		"work":    WebsiteTypeWork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebsiteType(input)
	return &out, nil
}

type WindowsDeviceMalwareStateCategory string

const (
	WindowsDeviceMalwareStateCategoryAdware                     WindowsDeviceMalwareStateCategory = "adware"
	WindowsDeviceMalwareStateCategoryAolExploit                 WindowsDeviceMalwareStateCategory = "aolExploit"
	WindowsDeviceMalwareStateCategoryBackdoor                   WindowsDeviceMalwareStateCategory = "backdoor"
	WindowsDeviceMalwareStateCategoryBehavior                   WindowsDeviceMalwareStateCategory = "behavior"
	WindowsDeviceMalwareStateCategoryBrowserModifier            WindowsDeviceMalwareStateCategory = "browserModifier"
	WindowsDeviceMalwareStateCategoryBrowserPlugin              WindowsDeviceMalwareStateCategory = "browserPlugin"
	WindowsDeviceMalwareStateCategoryCookie                     WindowsDeviceMalwareStateCategory = "cookie"
	WindowsDeviceMalwareStateCategoryDialer                     WindowsDeviceMalwareStateCategory = "dialer"
	WindowsDeviceMalwareStateCategoryEmailFlooder               WindowsDeviceMalwareStateCategory = "emailFlooder"
	WindowsDeviceMalwareStateCategoryEnterpriseUnwantedSoftware WindowsDeviceMalwareStateCategory = "enterpriseUnwantedSoftware"
	WindowsDeviceMalwareStateCategoryExploit                    WindowsDeviceMalwareStateCategory = "exploit"
	WindowsDeviceMalwareStateCategoryFilesharingProgram         WindowsDeviceMalwareStateCategory = "filesharingProgram"
	WindowsDeviceMalwareStateCategoryHipsRule                   WindowsDeviceMalwareStateCategory = "hipsRule"
	WindowsDeviceMalwareStateCategoryHostileActiveXControl      WindowsDeviceMalwareStateCategory = "hostileActiveXControl"
	WindowsDeviceMalwareStateCategoryIcqExploit                 WindowsDeviceMalwareStateCategory = "icqExploit"
	WindowsDeviceMalwareStateCategoryInvalid                    WindowsDeviceMalwareStateCategory = "invalid"
	WindowsDeviceMalwareStateCategoryJokeProgram                WindowsDeviceMalwareStateCategory = "jokeProgram"
	WindowsDeviceMalwareStateCategoryKeylogger                  WindowsDeviceMalwareStateCategory = "keylogger"
	WindowsDeviceMalwareStateCategoryKnown                      WindowsDeviceMalwareStateCategory = "known"
	WindowsDeviceMalwareStateCategoryMalwareCreationTool        WindowsDeviceMalwareStateCategory = "malwareCreationTool"
	WindowsDeviceMalwareStateCategoryMonitoringSoftware         WindowsDeviceMalwareStateCategory = "monitoringSoftware"
	WindowsDeviceMalwareStateCategoryNuker                      WindowsDeviceMalwareStateCategory = "nuker"
	WindowsDeviceMalwareStateCategoryPasswordStealer            WindowsDeviceMalwareStateCategory = "passwordStealer"
	WindowsDeviceMalwareStateCategoryPolicy                     WindowsDeviceMalwareStateCategory = "policy"
	WindowsDeviceMalwareStateCategoryPotentialUnwantedSoftware  WindowsDeviceMalwareStateCategory = "potentialUnwantedSoftware"
	WindowsDeviceMalwareStateCategoryRansom                     WindowsDeviceMalwareStateCategory = "ransom"
	WindowsDeviceMalwareStateCategoryRemoteAccessTrojan         WindowsDeviceMalwareStateCategory = "remoteAccessTrojan"
	WindowsDeviceMalwareStateCategoryRemoteControlSoftware      WindowsDeviceMalwareStateCategory = "remote_Control_Software"
	WindowsDeviceMalwareStateCategorySecurityDisabler           WindowsDeviceMalwareStateCategory = "securityDisabler"
	WindowsDeviceMalwareStateCategorySettingsModifier           WindowsDeviceMalwareStateCategory = "settingsModifier"
	WindowsDeviceMalwareStateCategorySoftwareBundler            WindowsDeviceMalwareStateCategory = "softwareBundler"
	WindowsDeviceMalwareStateCategorySpp                        WindowsDeviceMalwareStateCategory = "spp"
	WindowsDeviceMalwareStateCategorySpyware                    WindowsDeviceMalwareStateCategory = "spyware"
	WindowsDeviceMalwareStateCategoryStealthNotifier            WindowsDeviceMalwareStateCategory = "stealthNotifier"
	WindowsDeviceMalwareStateCategoryTool                       WindowsDeviceMalwareStateCategory = "tool"
	WindowsDeviceMalwareStateCategoryToolBar                    WindowsDeviceMalwareStateCategory = "toolBar"
	WindowsDeviceMalwareStateCategoryTrojan                     WindowsDeviceMalwareStateCategory = "trojan"
	WindowsDeviceMalwareStateCategoryTrojanDenialOfService      WindowsDeviceMalwareStateCategory = "trojanDenialOfService"
	WindowsDeviceMalwareStateCategoryTrojanDownloader           WindowsDeviceMalwareStateCategory = "trojanDownloader"
	WindowsDeviceMalwareStateCategoryTrojanDropper              WindowsDeviceMalwareStateCategory = "trojanDropper"
	WindowsDeviceMalwareStateCategoryTrojanFtp                  WindowsDeviceMalwareStateCategory = "trojanFtp"
	WindowsDeviceMalwareStateCategoryTrojanMassMailer           WindowsDeviceMalwareStateCategory = "trojanMassMailer"
	WindowsDeviceMalwareStateCategoryTrojanMonitoringSoftware   WindowsDeviceMalwareStateCategory = "trojanMonitoringSoftware"
	WindowsDeviceMalwareStateCategoryTrojanProxyServer          WindowsDeviceMalwareStateCategory = "trojanProxyServer"
	WindowsDeviceMalwareStateCategoryTrojanTelnet               WindowsDeviceMalwareStateCategory = "trojanTelnet"
	WindowsDeviceMalwareStateCategoryUnknown                    WindowsDeviceMalwareStateCategory = "unknown"
	WindowsDeviceMalwareStateCategoryVirus                      WindowsDeviceMalwareStateCategory = "virus"
	WindowsDeviceMalwareStateCategoryVulnerability              WindowsDeviceMalwareStateCategory = "vulnerability"
	WindowsDeviceMalwareStateCategoryWorm                       WindowsDeviceMalwareStateCategory = "worm"
)

func PossibleValuesForWindowsDeviceMalwareStateCategory() []string {
	return []string{
		string(WindowsDeviceMalwareStateCategoryAdware),
		string(WindowsDeviceMalwareStateCategoryAolExploit),
		string(WindowsDeviceMalwareStateCategoryBackdoor),
		string(WindowsDeviceMalwareStateCategoryBehavior),
		string(WindowsDeviceMalwareStateCategoryBrowserModifier),
		string(WindowsDeviceMalwareStateCategoryBrowserPlugin),
		string(WindowsDeviceMalwareStateCategoryCookie),
		string(WindowsDeviceMalwareStateCategoryDialer),
		string(WindowsDeviceMalwareStateCategoryEmailFlooder),
		string(WindowsDeviceMalwareStateCategoryEnterpriseUnwantedSoftware),
		string(WindowsDeviceMalwareStateCategoryExploit),
		string(WindowsDeviceMalwareStateCategoryFilesharingProgram),
		string(WindowsDeviceMalwareStateCategoryHipsRule),
		string(WindowsDeviceMalwareStateCategoryHostileActiveXControl),
		string(WindowsDeviceMalwareStateCategoryIcqExploit),
		string(WindowsDeviceMalwareStateCategoryInvalid),
		string(WindowsDeviceMalwareStateCategoryJokeProgram),
		string(WindowsDeviceMalwareStateCategoryKeylogger),
		string(WindowsDeviceMalwareStateCategoryKnown),
		string(WindowsDeviceMalwareStateCategoryMalwareCreationTool),
		string(WindowsDeviceMalwareStateCategoryMonitoringSoftware),
		string(WindowsDeviceMalwareStateCategoryNuker),
		string(WindowsDeviceMalwareStateCategoryPasswordStealer),
		string(WindowsDeviceMalwareStateCategoryPolicy),
		string(WindowsDeviceMalwareStateCategoryPotentialUnwantedSoftware),
		string(WindowsDeviceMalwareStateCategoryRansom),
		string(WindowsDeviceMalwareStateCategoryRemoteAccessTrojan),
		string(WindowsDeviceMalwareStateCategoryRemoteControlSoftware),
		string(WindowsDeviceMalwareStateCategorySecurityDisabler),
		string(WindowsDeviceMalwareStateCategorySettingsModifier),
		string(WindowsDeviceMalwareStateCategorySoftwareBundler),
		string(WindowsDeviceMalwareStateCategorySpp),
		string(WindowsDeviceMalwareStateCategorySpyware),
		string(WindowsDeviceMalwareStateCategoryStealthNotifier),
		string(WindowsDeviceMalwareStateCategoryTool),
		string(WindowsDeviceMalwareStateCategoryToolBar),
		string(WindowsDeviceMalwareStateCategoryTrojan),
		string(WindowsDeviceMalwareStateCategoryTrojanDenialOfService),
		string(WindowsDeviceMalwareStateCategoryTrojanDownloader),
		string(WindowsDeviceMalwareStateCategoryTrojanDropper),
		string(WindowsDeviceMalwareStateCategoryTrojanFtp),
		string(WindowsDeviceMalwareStateCategoryTrojanMassMailer),
		string(WindowsDeviceMalwareStateCategoryTrojanMonitoringSoftware),
		string(WindowsDeviceMalwareStateCategoryTrojanProxyServer),
		string(WindowsDeviceMalwareStateCategoryTrojanTelnet),
		string(WindowsDeviceMalwareStateCategoryUnknown),
		string(WindowsDeviceMalwareStateCategoryVirus),
		string(WindowsDeviceMalwareStateCategoryVulnerability),
		string(WindowsDeviceMalwareStateCategoryWorm),
	}
}

func (s *WindowsDeviceMalwareStateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceMalwareStateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceMalwareStateCategory(input string) (*WindowsDeviceMalwareStateCategory, error) {
	vals := map[string]WindowsDeviceMalwareStateCategory{
		"adware":                     WindowsDeviceMalwareStateCategoryAdware,
		"aolexploit":                 WindowsDeviceMalwareStateCategoryAolExploit,
		"backdoor":                   WindowsDeviceMalwareStateCategoryBackdoor,
		"behavior":                   WindowsDeviceMalwareStateCategoryBehavior,
		"browsermodifier":            WindowsDeviceMalwareStateCategoryBrowserModifier,
		"browserplugin":              WindowsDeviceMalwareStateCategoryBrowserPlugin,
		"cookie":                     WindowsDeviceMalwareStateCategoryCookie,
		"dialer":                     WindowsDeviceMalwareStateCategoryDialer,
		"emailflooder":               WindowsDeviceMalwareStateCategoryEmailFlooder,
		"enterpriseunwantedsoftware": WindowsDeviceMalwareStateCategoryEnterpriseUnwantedSoftware,
		"exploit":                    WindowsDeviceMalwareStateCategoryExploit,
		"filesharingprogram":         WindowsDeviceMalwareStateCategoryFilesharingProgram,
		"hipsrule":                   WindowsDeviceMalwareStateCategoryHipsRule,
		"hostileactivexcontrol":      WindowsDeviceMalwareStateCategoryHostileActiveXControl,
		"icqexploit":                 WindowsDeviceMalwareStateCategoryIcqExploit,
		"invalid":                    WindowsDeviceMalwareStateCategoryInvalid,
		"jokeprogram":                WindowsDeviceMalwareStateCategoryJokeProgram,
		"keylogger":                  WindowsDeviceMalwareStateCategoryKeylogger,
		"known":                      WindowsDeviceMalwareStateCategoryKnown,
		"malwarecreationtool":        WindowsDeviceMalwareStateCategoryMalwareCreationTool,
		"monitoringsoftware":         WindowsDeviceMalwareStateCategoryMonitoringSoftware,
		"nuker":                      WindowsDeviceMalwareStateCategoryNuker,
		"passwordstealer":            WindowsDeviceMalwareStateCategoryPasswordStealer,
		"policy":                     WindowsDeviceMalwareStateCategoryPolicy,
		"potentialunwantedsoftware":  WindowsDeviceMalwareStateCategoryPotentialUnwantedSoftware,
		"ransom":                     WindowsDeviceMalwareStateCategoryRansom,
		"remoteaccesstrojan":         WindowsDeviceMalwareStateCategoryRemoteAccessTrojan,
		"remote_control_software":    WindowsDeviceMalwareStateCategoryRemoteControlSoftware,
		"securitydisabler":           WindowsDeviceMalwareStateCategorySecurityDisabler,
		"settingsmodifier":           WindowsDeviceMalwareStateCategorySettingsModifier,
		"softwarebundler":            WindowsDeviceMalwareStateCategorySoftwareBundler,
		"spp":                        WindowsDeviceMalwareStateCategorySpp,
		"spyware":                    WindowsDeviceMalwareStateCategorySpyware,
		"stealthnotifier":            WindowsDeviceMalwareStateCategoryStealthNotifier,
		"tool":                       WindowsDeviceMalwareStateCategoryTool,
		"toolbar":                    WindowsDeviceMalwareStateCategoryToolBar,
		"trojan":                     WindowsDeviceMalwareStateCategoryTrojan,
		"trojandenialofservice":      WindowsDeviceMalwareStateCategoryTrojanDenialOfService,
		"trojandownloader":           WindowsDeviceMalwareStateCategoryTrojanDownloader,
		"trojandropper":              WindowsDeviceMalwareStateCategoryTrojanDropper,
		"trojanftp":                  WindowsDeviceMalwareStateCategoryTrojanFtp,
		"trojanmassmailer":           WindowsDeviceMalwareStateCategoryTrojanMassMailer,
		"trojanmonitoringsoftware":   WindowsDeviceMalwareStateCategoryTrojanMonitoringSoftware,
		"trojanproxyserver":          WindowsDeviceMalwareStateCategoryTrojanProxyServer,
		"trojantelnet":               WindowsDeviceMalwareStateCategoryTrojanTelnet,
		"unknown":                    WindowsDeviceMalwareStateCategoryUnknown,
		"virus":                      WindowsDeviceMalwareStateCategoryVirus,
		"vulnerability":              WindowsDeviceMalwareStateCategoryVulnerability,
		"worm":                       WindowsDeviceMalwareStateCategoryWorm,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceMalwareStateCategory(input)
	return &out, nil
}

type WindowsDeviceMalwareStateExecutionState string

const (
	WindowsDeviceMalwareStateExecutionStateAllowed    WindowsDeviceMalwareStateExecutionState = "allowed"
	WindowsDeviceMalwareStateExecutionStateBlocked    WindowsDeviceMalwareStateExecutionState = "blocked"
	WindowsDeviceMalwareStateExecutionStateNotRunning WindowsDeviceMalwareStateExecutionState = "notRunning"
	WindowsDeviceMalwareStateExecutionStateRunning    WindowsDeviceMalwareStateExecutionState = "running"
	WindowsDeviceMalwareStateExecutionStateUnknown    WindowsDeviceMalwareStateExecutionState = "unknown"
)

func PossibleValuesForWindowsDeviceMalwareStateExecutionState() []string {
	return []string{
		string(WindowsDeviceMalwareStateExecutionStateAllowed),
		string(WindowsDeviceMalwareStateExecutionStateBlocked),
		string(WindowsDeviceMalwareStateExecutionStateNotRunning),
		string(WindowsDeviceMalwareStateExecutionStateRunning),
		string(WindowsDeviceMalwareStateExecutionStateUnknown),
	}
}

func (s *WindowsDeviceMalwareStateExecutionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceMalwareStateExecutionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceMalwareStateExecutionState(input string) (*WindowsDeviceMalwareStateExecutionState, error) {
	vals := map[string]WindowsDeviceMalwareStateExecutionState{
		"allowed":    WindowsDeviceMalwareStateExecutionStateAllowed,
		"blocked":    WindowsDeviceMalwareStateExecutionStateBlocked,
		"notrunning": WindowsDeviceMalwareStateExecutionStateNotRunning,
		"running":    WindowsDeviceMalwareStateExecutionStateRunning,
		"unknown":    WindowsDeviceMalwareStateExecutionStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceMalwareStateExecutionState(input)
	return &out, nil
}

type WindowsDeviceMalwareStateSeverity string

const (
	WindowsDeviceMalwareStateSeverityHigh     WindowsDeviceMalwareStateSeverity = "high"
	WindowsDeviceMalwareStateSeverityLow      WindowsDeviceMalwareStateSeverity = "low"
	WindowsDeviceMalwareStateSeverityModerate WindowsDeviceMalwareStateSeverity = "moderate"
	WindowsDeviceMalwareStateSeveritySevere   WindowsDeviceMalwareStateSeverity = "severe"
	WindowsDeviceMalwareStateSeverityUnknown  WindowsDeviceMalwareStateSeverity = "unknown"
)

func PossibleValuesForWindowsDeviceMalwareStateSeverity() []string {
	return []string{
		string(WindowsDeviceMalwareStateSeverityHigh),
		string(WindowsDeviceMalwareStateSeverityLow),
		string(WindowsDeviceMalwareStateSeverityModerate),
		string(WindowsDeviceMalwareStateSeveritySevere),
		string(WindowsDeviceMalwareStateSeverityUnknown),
	}
}

func (s *WindowsDeviceMalwareStateSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceMalwareStateSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceMalwareStateSeverity(input string) (*WindowsDeviceMalwareStateSeverity, error) {
	vals := map[string]WindowsDeviceMalwareStateSeverity{
		"high":     WindowsDeviceMalwareStateSeverityHigh,
		"low":      WindowsDeviceMalwareStateSeverityLow,
		"moderate": WindowsDeviceMalwareStateSeverityModerate,
		"severe":   WindowsDeviceMalwareStateSeveritySevere,
		"unknown":  WindowsDeviceMalwareStateSeverityUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceMalwareStateSeverity(input)
	return &out, nil
}

type WindowsDeviceMalwareStateState string

const (
	WindowsDeviceMalwareStateStateAbandoned        WindowsDeviceMalwareStateState = "abandoned"
	WindowsDeviceMalwareStateStateAllowFailed      WindowsDeviceMalwareStateState = "allowFailed"
	WindowsDeviceMalwareStateStateAllowed          WindowsDeviceMalwareStateState = "allowed"
	WindowsDeviceMalwareStateStateBlockFailed      WindowsDeviceMalwareStateState = "blockFailed"
	WindowsDeviceMalwareStateStateBlocked          WindowsDeviceMalwareStateState = "blocked"
	WindowsDeviceMalwareStateStateCleanFailed      WindowsDeviceMalwareStateState = "cleanFailed"
	WindowsDeviceMalwareStateStateCleaned          WindowsDeviceMalwareStateState = "cleaned"
	WindowsDeviceMalwareStateStateDetected         WindowsDeviceMalwareStateState = "detected"
	WindowsDeviceMalwareStateStateQuarantineFailed WindowsDeviceMalwareStateState = "quarantineFailed"
	WindowsDeviceMalwareStateStateQuarantined      WindowsDeviceMalwareStateState = "quarantined"
	WindowsDeviceMalwareStateStateRemoveFailed     WindowsDeviceMalwareStateState = "removeFailed"
	WindowsDeviceMalwareStateStateRemoved          WindowsDeviceMalwareStateState = "removed"
	WindowsDeviceMalwareStateStateUnknown          WindowsDeviceMalwareStateState = "unknown"
)

func PossibleValuesForWindowsDeviceMalwareStateState() []string {
	return []string{
		string(WindowsDeviceMalwareStateStateAbandoned),
		string(WindowsDeviceMalwareStateStateAllowFailed),
		string(WindowsDeviceMalwareStateStateAllowed),
		string(WindowsDeviceMalwareStateStateBlockFailed),
		string(WindowsDeviceMalwareStateStateBlocked),
		string(WindowsDeviceMalwareStateStateCleanFailed),
		string(WindowsDeviceMalwareStateStateCleaned),
		string(WindowsDeviceMalwareStateStateDetected),
		string(WindowsDeviceMalwareStateStateQuarantineFailed),
		string(WindowsDeviceMalwareStateStateQuarantined),
		string(WindowsDeviceMalwareStateStateRemoveFailed),
		string(WindowsDeviceMalwareStateStateRemoved),
		string(WindowsDeviceMalwareStateStateUnknown),
	}
}

func (s *WindowsDeviceMalwareStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceMalwareStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceMalwareStateState(input string) (*WindowsDeviceMalwareStateState, error) {
	vals := map[string]WindowsDeviceMalwareStateState{
		"abandoned":        WindowsDeviceMalwareStateStateAbandoned,
		"allowfailed":      WindowsDeviceMalwareStateStateAllowFailed,
		"allowed":          WindowsDeviceMalwareStateStateAllowed,
		"blockfailed":      WindowsDeviceMalwareStateStateBlockFailed,
		"blocked":          WindowsDeviceMalwareStateStateBlocked,
		"cleanfailed":      WindowsDeviceMalwareStateStateCleanFailed,
		"cleaned":          WindowsDeviceMalwareStateStateCleaned,
		"detected":         WindowsDeviceMalwareStateStateDetected,
		"quarantinefailed": WindowsDeviceMalwareStateStateQuarantineFailed,
		"quarantined":      WindowsDeviceMalwareStateStateQuarantined,
		"removefailed":     WindowsDeviceMalwareStateStateRemoveFailed,
		"removed":          WindowsDeviceMalwareStateStateRemoved,
		"unknown":          WindowsDeviceMalwareStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceMalwareStateState(input)
	return &out, nil
}

type WindowsDeviceMalwareStateThreatState string

const (
	WindowsDeviceMalwareStateThreatStateActionFailed                      WindowsDeviceMalwareStateThreatState = "actionFailed"
	WindowsDeviceMalwareStateThreatStateActive                            WindowsDeviceMalwareStateThreatState = "active"
	WindowsDeviceMalwareStateThreatStateAllowed                           WindowsDeviceMalwareStateThreatState = "allowed"
	WindowsDeviceMalwareStateThreatStateCleaned                           WindowsDeviceMalwareStateThreatState = "cleaned"
	WindowsDeviceMalwareStateThreatStateFullScanRequired                  WindowsDeviceMalwareStateThreatState = "fullScanRequired"
	WindowsDeviceMalwareStateThreatStateManualStepsRequired               WindowsDeviceMalwareStateThreatState = "manualStepsRequired"
	WindowsDeviceMalwareStateThreatStateNoStatusCleared                   WindowsDeviceMalwareStateThreatState = "noStatusCleared"
	WindowsDeviceMalwareStateThreatStateQuarantined                       WindowsDeviceMalwareStateThreatState = "quarantined"
	WindowsDeviceMalwareStateThreatStateRebootRequired                    WindowsDeviceMalwareStateThreatState = "rebootRequired"
	WindowsDeviceMalwareStateThreatStateRemediatedWithNonCriticalFailures WindowsDeviceMalwareStateThreatState = "remediatedWithNonCriticalFailures"
	WindowsDeviceMalwareStateThreatStateRemoved                           WindowsDeviceMalwareStateThreatState = "removed"
)

func PossibleValuesForWindowsDeviceMalwareStateThreatState() []string {
	return []string{
		string(WindowsDeviceMalwareStateThreatStateActionFailed),
		string(WindowsDeviceMalwareStateThreatStateActive),
		string(WindowsDeviceMalwareStateThreatStateAllowed),
		string(WindowsDeviceMalwareStateThreatStateCleaned),
		string(WindowsDeviceMalwareStateThreatStateFullScanRequired),
		string(WindowsDeviceMalwareStateThreatStateManualStepsRequired),
		string(WindowsDeviceMalwareStateThreatStateNoStatusCleared),
		string(WindowsDeviceMalwareStateThreatStateQuarantined),
		string(WindowsDeviceMalwareStateThreatStateRebootRequired),
		string(WindowsDeviceMalwareStateThreatStateRemediatedWithNonCriticalFailures),
		string(WindowsDeviceMalwareStateThreatStateRemoved),
	}
}

func (s *WindowsDeviceMalwareStateThreatState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceMalwareStateThreatState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceMalwareStateThreatState(input string) (*WindowsDeviceMalwareStateThreatState, error) {
	vals := map[string]WindowsDeviceMalwareStateThreatState{
		"actionfailed":                      WindowsDeviceMalwareStateThreatStateActionFailed,
		"active":                            WindowsDeviceMalwareStateThreatStateActive,
		"allowed":                           WindowsDeviceMalwareStateThreatStateAllowed,
		"cleaned":                           WindowsDeviceMalwareStateThreatStateCleaned,
		"fullscanrequired":                  WindowsDeviceMalwareStateThreatStateFullScanRequired,
		"manualstepsrequired":               WindowsDeviceMalwareStateThreatStateManualStepsRequired,
		"nostatuscleared":                   WindowsDeviceMalwareStateThreatStateNoStatusCleared,
		"quarantined":                       WindowsDeviceMalwareStateThreatStateQuarantined,
		"rebootrequired":                    WindowsDeviceMalwareStateThreatStateRebootRequired,
		"remediatedwithnoncriticalfailures": WindowsDeviceMalwareStateThreatStateRemediatedWithNonCriticalFailures,
		"removed":                           WindowsDeviceMalwareStateThreatStateRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceMalwareStateThreatState(input)
	return &out, nil
}

type WindowsHelloForBusinessAuthenticationMethodKeyStrength string

const (
	WindowsHelloForBusinessAuthenticationMethodKeyStrengthNormal  WindowsHelloForBusinessAuthenticationMethodKeyStrength = "normal"
	WindowsHelloForBusinessAuthenticationMethodKeyStrengthUnknown WindowsHelloForBusinessAuthenticationMethodKeyStrength = "unknown"
	WindowsHelloForBusinessAuthenticationMethodKeyStrengthWeak    WindowsHelloForBusinessAuthenticationMethodKeyStrength = "weak"
)

func PossibleValuesForWindowsHelloForBusinessAuthenticationMethodKeyStrength() []string {
	return []string{
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrengthNormal),
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrengthUnknown),
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrengthWeak),
	}
}

func (s *WindowsHelloForBusinessAuthenticationMethodKeyStrength) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsHelloForBusinessAuthenticationMethodKeyStrength(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsHelloForBusinessAuthenticationMethodKeyStrength(input string) (*WindowsHelloForBusinessAuthenticationMethodKeyStrength, error) {
	vals := map[string]WindowsHelloForBusinessAuthenticationMethodKeyStrength{
		"normal":  WindowsHelloForBusinessAuthenticationMethodKeyStrengthNormal,
		"unknown": WindowsHelloForBusinessAuthenticationMethodKeyStrengthUnknown,
		"weak":    WindowsHelloForBusinessAuthenticationMethodKeyStrengthWeak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHelloForBusinessAuthenticationMethodKeyStrength(input)
	return &out, nil
}

type WindowsProtectionStateDeviceState string

const (
	WindowsProtectionStateDeviceStateClean              WindowsProtectionStateDeviceState = "clean"
	WindowsProtectionStateDeviceStateCritical           WindowsProtectionStateDeviceState = "critical"
	WindowsProtectionStateDeviceStateFullScanPending    WindowsProtectionStateDeviceState = "fullScanPending"
	WindowsProtectionStateDeviceStateManualStepsPending WindowsProtectionStateDeviceState = "manualStepsPending"
	WindowsProtectionStateDeviceStateOfflineScanPending WindowsProtectionStateDeviceState = "offlineScanPending"
	WindowsProtectionStateDeviceStateRebootPending      WindowsProtectionStateDeviceState = "rebootPending"
)

func PossibleValuesForWindowsProtectionStateDeviceState() []string {
	return []string{
		string(WindowsProtectionStateDeviceStateClean),
		string(WindowsProtectionStateDeviceStateCritical),
		string(WindowsProtectionStateDeviceStateFullScanPending),
		string(WindowsProtectionStateDeviceStateManualStepsPending),
		string(WindowsProtectionStateDeviceStateOfflineScanPending),
		string(WindowsProtectionStateDeviceStateRebootPending),
	}
}

func (s *WindowsProtectionStateDeviceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsProtectionStateDeviceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsProtectionStateDeviceState(input string) (*WindowsProtectionStateDeviceState, error) {
	vals := map[string]WindowsProtectionStateDeviceState{
		"clean":              WindowsProtectionStateDeviceStateClean,
		"critical":           WindowsProtectionStateDeviceStateCritical,
		"fullscanpending":    WindowsProtectionStateDeviceStateFullScanPending,
		"manualstepspending": WindowsProtectionStateDeviceStateManualStepsPending,
		"offlinescanpending": WindowsProtectionStateDeviceStateOfflineScanPending,
		"rebootpending":      WindowsProtectionStateDeviceStateRebootPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsProtectionStateDeviceState(input)
	return &out, nil
}

type WindowsProtectionStateProductStatus string

const (
	WindowsProtectionStateProductStatusAsSignaturesOutOfDate                           WindowsProtectionStateProductStatus = "asSignaturesOutOfDate"
	WindowsProtectionStateProductStatusAvSignaturesOutOfDate                           WindowsProtectionStateProductStatus = "avSignaturesOutOfDate"
	WindowsProtectionStateProductStatusNoFullScanHappenedForSpecifiedPeriod            WindowsProtectionStateProductStatus = "noFullScanHappenedForSpecifiedPeriod"
	WindowsProtectionStateProductStatusNoQuickScanHappenedForSpecifiedPeriod           WindowsProtectionStateProductStatus = "noQuickScanHappenedForSpecifiedPeriod"
	WindowsProtectionStateProductStatusNoStatus                                        WindowsProtectionStateProductStatus = "noStatus"
	WindowsProtectionStateProductStatusNoStatusFlagsSet                                WindowsProtectionStateProductStatus = "noStatusFlagsSet"
	WindowsProtectionStateProductStatusOfflineScanRequired                             WindowsProtectionStateProductStatus = "offlineScanRequired"
	WindowsProtectionStateProductStatusPendingFullScanDueToThreatAction                WindowsProtectionStateProductStatus = "pendingFullScanDueToThreatAction"
	WindowsProtectionStateProductStatusPendingManualStepsDueToThreatAction             WindowsProtectionStateProductStatus = "pendingManualStepsDueToThreatAction"
	WindowsProtectionStateProductStatusPendingRebootDueToThreatAction                  WindowsProtectionStateProductStatus = "pendingRebootDueToThreatAction"
	WindowsProtectionStateProductStatusPlatformAboutToBeOutdated                       WindowsProtectionStateProductStatus = "platformAboutToBeOutdated"
	WindowsProtectionStateProductStatusPlatformOutOfDate                               WindowsProtectionStateProductStatus = "platformOutOfDate"
	WindowsProtectionStateProductStatusPlatformUpdateInProgress                        WindowsProtectionStateProductStatus = "platformUpdateInProgress"
	WindowsProtectionStateProductStatusProductExpired                                  WindowsProtectionStateProductStatus = "productExpired"
	WindowsProtectionStateProductStatusProductRunningInEvaluationMode                  WindowsProtectionStateProductStatus = "productRunningInEvaluationMode"
	WindowsProtectionStateProductStatusProductRunningInNonGenuineMode                  WindowsProtectionStateProductStatus = "productRunningInNonGenuineMode"
	WindowsProtectionStateProductStatusSamplesPendingSubmission                        WindowsProtectionStateProductStatus = "samplesPendingSubmission"
	WindowsProtectionStateProductStatusServiceNotRunning                               WindowsProtectionStateProductStatus = "serviceNotRunning"
	WindowsProtectionStateProductStatusServiceShutdownAsPartOfSystemShutdown           WindowsProtectionStateProductStatus = "serviceShutdownAsPartOfSystemShutdown"
	WindowsProtectionStateProductStatusServiceStartedWithoutMalwareProtection          WindowsProtectionStateProductStatus = "serviceStartedWithoutMalwareProtection"
	WindowsProtectionStateProductStatusSignatureOrPlatformEndOfLifeIsPastOrIsImpending WindowsProtectionStateProductStatus = "signatureOrPlatformEndOfLifeIsPastOrIsImpending"
	WindowsProtectionStateProductStatusSystemInitiatedCleanInProgress                  WindowsProtectionStateProductStatus = "systemInitiatedCleanInProgress"
	WindowsProtectionStateProductStatusSystemInitiatedScanInProgress                   WindowsProtectionStateProductStatus = "systemInitiatedScanInProgress"
	WindowsProtectionStateProductStatusThreatRemediationFailedCritically               WindowsProtectionStateProductStatus = "threatRemediationFailedCritically"
	WindowsProtectionStateProductStatusThreatRemediationFailedNonCritically            WindowsProtectionStateProductStatus = "threatRemediationFailedNonCritically"
	WindowsProtectionStateProductStatusWindowsSModeSignaturesInUseOnNonWin10SInstall   WindowsProtectionStateProductStatus = "windowsSModeSignaturesInUseOnNonWin10SInstall"
)

func PossibleValuesForWindowsProtectionStateProductStatus() []string {
	return []string{
		string(WindowsProtectionStateProductStatusAsSignaturesOutOfDate),
		string(WindowsProtectionStateProductStatusAvSignaturesOutOfDate),
		string(WindowsProtectionStateProductStatusNoFullScanHappenedForSpecifiedPeriod),
		string(WindowsProtectionStateProductStatusNoQuickScanHappenedForSpecifiedPeriod),
		string(WindowsProtectionStateProductStatusNoStatus),
		string(WindowsProtectionStateProductStatusNoStatusFlagsSet),
		string(WindowsProtectionStateProductStatusOfflineScanRequired),
		string(WindowsProtectionStateProductStatusPendingFullScanDueToThreatAction),
		string(WindowsProtectionStateProductStatusPendingManualStepsDueToThreatAction),
		string(WindowsProtectionStateProductStatusPendingRebootDueToThreatAction),
		string(WindowsProtectionStateProductStatusPlatformAboutToBeOutdated),
		string(WindowsProtectionStateProductStatusPlatformOutOfDate),
		string(WindowsProtectionStateProductStatusPlatformUpdateInProgress),
		string(WindowsProtectionStateProductStatusProductExpired),
		string(WindowsProtectionStateProductStatusProductRunningInEvaluationMode),
		string(WindowsProtectionStateProductStatusProductRunningInNonGenuineMode),
		string(WindowsProtectionStateProductStatusSamplesPendingSubmission),
		string(WindowsProtectionStateProductStatusServiceNotRunning),
		string(WindowsProtectionStateProductStatusServiceShutdownAsPartOfSystemShutdown),
		string(WindowsProtectionStateProductStatusServiceStartedWithoutMalwareProtection),
		string(WindowsProtectionStateProductStatusSignatureOrPlatformEndOfLifeIsPastOrIsImpending),
		string(WindowsProtectionStateProductStatusSystemInitiatedCleanInProgress),
		string(WindowsProtectionStateProductStatusSystemInitiatedScanInProgress),
		string(WindowsProtectionStateProductStatusThreatRemediationFailedCritically),
		string(WindowsProtectionStateProductStatusThreatRemediationFailedNonCritically),
		string(WindowsProtectionStateProductStatusWindowsSModeSignaturesInUseOnNonWin10SInstall),
	}
}

func (s *WindowsProtectionStateProductStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsProtectionStateProductStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsProtectionStateProductStatus(input string) (*WindowsProtectionStateProductStatus, error) {
	vals := map[string]WindowsProtectionStateProductStatus{
		"assignaturesoutofdate":                           WindowsProtectionStateProductStatusAsSignaturesOutOfDate,
		"avsignaturesoutofdate":                           WindowsProtectionStateProductStatusAvSignaturesOutOfDate,
		"nofullscanhappenedforspecifiedperiod":            WindowsProtectionStateProductStatusNoFullScanHappenedForSpecifiedPeriod,
		"noquickscanhappenedforspecifiedperiod":           WindowsProtectionStateProductStatusNoQuickScanHappenedForSpecifiedPeriod,
		"nostatus":                                        WindowsProtectionStateProductStatusNoStatus,
		"nostatusflagsset":                                WindowsProtectionStateProductStatusNoStatusFlagsSet,
		"offlinescanrequired":                             WindowsProtectionStateProductStatusOfflineScanRequired,
		"pendingfullscanduetothreataction":                WindowsProtectionStateProductStatusPendingFullScanDueToThreatAction,
		"pendingmanualstepsduetothreataction":             WindowsProtectionStateProductStatusPendingManualStepsDueToThreatAction,
		"pendingrebootduetothreataction":                  WindowsProtectionStateProductStatusPendingRebootDueToThreatAction,
		"platformabouttobeoutdated":                       WindowsProtectionStateProductStatusPlatformAboutToBeOutdated,
		"platformoutofdate":                               WindowsProtectionStateProductStatusPlatformOutOfDate,
		"platformupdateinprogress":                        WindowsProtectionStateProductStatusPlatformUpdateInProgress,
		"productexpired":                                  WindowsProtectionStateProductStatusProductExpired,
		"productrunninginevaluationmode":                  WindowsProtectionStateProductStatusProductRunningInEvaluationMode,
		"productrunninginnongenuinemode":                  WindowsProtectionStateProductStatusProductRunningInNonGenuineMode,
		"samplespendingsubmission":                        WindowsProtectionStateProductStatusSamplesPendingSubmission,
		"servicenotrunning":                               WindowsProtectionStateProductStatusServiceNotRunning,
		"serviceshutdownaspartofsystemshutdown":           WindowsProtectionStateProductStatusServiceShutdownAsPartOfSystemShutdown,
		"servicestartedwithoutmalwareprotection":          WindowsProtectionStateProductStatusServiceStartedWithoutMalwareProtection,
		"signatureorplatformendoflifeispastorisimpending": WindowsProtectionStateProductStatusSignatureOrPlatformEndOfLifeIsPastOrIsImpending,
		"systeminitiatedcleaninprogress":                  WindowsProtectionStateProductStatusSystemInitiatedCleanInProgress,
		"systeminitiatedscaninprogress":                   WindowsProtectionStateProductStatusSystemInitiatedScanInProgress,
		"threatremediationfailedcritically":               WindowsProtectionStateProductStatusThreatRemediationFailedCritically,
		"threatremediationfailednoncritically":            WindowsProtectionStateProductStatusThreatRemediationFailedNonCritically,
		"windowssmodesignaturesinuseonnonwin10sinstall":   WindowsProtectionStateProductStatusWindowsSModeSignaturesInUseOnNonWin10SInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsProtectionStateProductStatus(input)
	return &out, nil
}

type WorkbookOperationStatus string

const (
	WorkbookOperationStatusFailed     WorkbookOperationStatus = "failed"
	WorkbookOperationStatusNotStarted WorkbookOperationStatus = "notStarted"
	WorkbookOperationStatusRunning    WorkbookOperationStatus = "running"
	WorkbookOperationStatusSucceeded  WorkbookOperationStatus = "succeeded"
)

func PossibleValuesForWorkbookOperationStatus() []string {
	return []string{
		string(WorkbookOperationStatusFailed),
		string(WorkbookOperationStatusNotStarted),
		string(WorkbookOperationStatusRunning),
		string(WorkbookOperationStatusSucceeded),
	}
}

func (s *WorkbookOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkbookOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkbookOperationStatus(input string) (*WorkbookOperationStatus, error) {
	vals := map[string]WorkbookOperationStatus{
		"failed":     WorkbookOperationStatusFailed,
		"notstarted": WorkbookOperationStatusNotStarted,
		"running":    WorkbookOperationStatusRunning,
		"succeeded":  WorkbookOperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkbookOperationStatus(input)
	return &out, nil
}

type WorkingHoursDaysOfWeek string

const (
	WorkingHoursDaysOfWeekFriday    WorkingHoursDaysOfWeek = "friday"
	WorkingHoursDaysOfWeekMonday    WorkingHoursDaysOfWeek = "monday"
	WorkingHoursDaysOfWeekSaturday  WorkingHoursDaysOfWeek = "saturday"
	WorkingHoursDaysOfWeekSunday    WorkingHoursDaysOfWeek = "sunday"
	WorkingHoursDaysOfWeekThursday  WorkingHoursDaysOfWeek = "thursday"
	WorkingHoursDaysOfWeekTuesday   WorkingHoursDaysOfWeek = "tuesday"
	WorkingHoursDaysOfWeekWednesday WorkingHoursDaysOfWeek = "wednesday"
)

func PossibleValuesForWorkingHoursDaysOfWeek() []string {
	return []string{
		string(WorkingHoursDaysOfWeekFriday),
		string(WorkingHoursDaysOfWeekMonday),
		string(WorkingHoursDaysOfWeekSaturday),
		string(WorkingHoursDaysOfWeekSunday),
		string(WorkingHoursDaysOfWeekThursday),
		string(WorkingHoursDaysOfWeekTuesday),
		string(WorkingHoursDaysOfWeekWednesday),
	}
}

func (s *WorkingHoursDaysOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkingHoursDaysOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkingHoursDaysOfWeek(input string) (*WorkingHoursDaysOfWeek, error) {
	vals := map[string]WorkingHoursDaysOfWeek{
		"friday":    WorkingHoursDaysOfWeekFriday,
		"monday":    WorkingHoursDaysOfWeekMonday,
		"saturday":  WorkingHoursDaysOfWeekSaturday,
		"sunday":    WorkingHoursDaysOfWeekSunday,
		"thursday":  WorkingHoursDaysOfWeekThursday,
		"tuesday":   WorkingHoursDaysOfWeekTuesday,
		"wednesday": WorkingHoursDaysOfWeekWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkingHoursDaysOfWeek(input)
	return &out, nil
}
