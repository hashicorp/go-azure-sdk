package user

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeBaseType string

const (
	AttendeeBaseTypeOptional AttendeeBaseType = "optional"
	AttendeeBaseTypeRequired AttendeeBaseType = "required"
	AttendeeBaseTypeResource AttendeeBaseType = "resource"
)

func PossibleValuesForAttendeeBaseType() []string {
	return []string{
		string(AttendeeBaseTypeOptional),
		string(AttendeeBaseTypeRequired),
		string(AttendeeBaseTypeResource),
	}
}

func (s *AttendeeBaseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttendeeBaseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttendeeBaseType(input string) (*AttendeeBaseType, error) {
	vals := map[string]AttendeeBaseType{
		"optional": AttendeeBaseTypeOptional,
		"required": AttendeeBaseTypeRequired,
		"resource": AttendeeBaseTypeResource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeBaseType(input)
	return &out, nil
}

type CreateTranslateExchangeIdRequestSourceIdType string

const (
	CreateTranslateExchangeIdRequestSourceIdTypeEntryId              CreateTranslateExchangeIdRequestSourceIdType = "entryId"
	CreateTranslateExchangeIdRequestSourceIdTypeEwsId                CreateTranslateExchangeIdRequestSourceIdType = "ewsId"
	CreateTranslateExchangeIdRequestSourceIdTypeImmutableEntryId     CreateTranslateExchangeIdRequestSourceIdType = "immutableEntryId"
	CreateTranslateExchangeIdRequestSourceIdTypeRestId               CreateTranslateExchangeIdRequestSourceIdType = "restId"
	CreateTranslateExchangeIdRequestSourceIdTypeRestImmutableEntryId CreateTranslateExchangeIdRequestSourceIdType = "restImmutableEntryId"
)

func PossibleValuesForCreateTranslateExchangeIdRequestSourceIdType() []string {
	return []string{
		string(CreateTranslateExchangeIdRequestSourceIdTypeEntryId),
		string(CreateTranslateExchangeIdRequestSourceIdTypeEwsId),
		string(CreateTranslateExchangeIdRequestSourceIdTypeImmutableEntryId),
		string(CreateTranslateExchangeIdRequestSourceIdTypeRestId),
		string(CreateTranslateExchangeIdRequestSourceIdTypeRestImmutableEntryId),
	}
}

func (s *CreateTranslateExchangeIdRequestSourceIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateTranslateExchangeIdRequestSourceIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateTranslateExchangeIdRequestSourceIdType(input string) (*CreateTranslateExchangeIdRequestSourceIdType, error) {
	vals := map[string]CreateTranslateExchangeIdRequestSourceIdType{
		"entryid":              CreateTranslateExchangeIdRequestSourceIdTypeEntryId,
		"ewsid":                CreateTranslateExchangeIdRequestSourceIdTypeEwsId,
		"immutableentryid":     CreateTranslateExchangeIdRequestSourceIdTypeImmutableEntryId,
		"restid":               CreateTranslateExchangeIdRequestSourceIdTypeRestId,
		"restimmutableentryid": CreateTranslateExchangeIdRequestSourceIdTypeRestImmutableEntryId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateTranslateExchangeIdRequestSourceIdType(input)
	return &out, nil
}

type CreateTranslateExchangeIdRequestTargetIdType string

const (
	CreateTranslateExchangeIdRequestTargetIdTypeEntryId              CreateTranslateExchangeIdRequestTargetIdType = "entryId"
	CreateTranslateExchangeIdRequestTargetIdTypeEwsId                CreateTranslateExchangeIdRequestTargetIdType = "ewsId"
	CreateTranslateExchangeIdRequestTargetIdTypeImmutableEntryId     CreateTranslateExchangeIdRequestTargetIdType = "immutableEntryId"
	CreateTranslateExchangeIdRequestTargetIdTypeRestId               CreateTranslateExchangeIdRequestTargetIdType = "restId"
	CreateTranslateExchangeIdRequestTargetIdTypeRestImmutableEntryId CreateTranslateExchangeIdRequestTargetIdType = "restImmutableEntryId"
)

func PossibleValuesForCreateTranslateExchangeIdRequestTargetIdType() []string {
	return []string{
		string(CreateTranslateExchangeIdRequestTargetIdTypeEntryId),
		string(CreateTranslateExchangeIdRequestTargetIdTypeEwsId),
		string(CreateTranslateExchangeIdRequestTargetIdTypeImmutableEntryId),
		string(CreateTranslateExchangeIdRequestTargetIdTypeRestId),
		string(CreateTranslateExchangeIdRequestTargetIdTypeRestImmutableEntryId),
	}
}

func (s *CreateTranslateExchangeIdRequestTargetIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateTranslateExchangeIdRequestTargetIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateTranslateExchangeIdRequestTargetIdType(input string) (*CreateTranslateExchangeIdRequestTargetIdType, error) {
	vals := map[string]CreateTranslateExchangeIdRequestTargetIdType{
		"entryid":              CreateTranslateExchangeIdRequestTargetIdTypeEntryId,
		"ewsid":                CreateTranslateExchangeIdRequestTargetIdTypeEwsId,
		"immutableentryid":     CreateTranslateExchangeIdRequestTargetIdTypeImmutableEntryId,
		"restid":               CreateTranslateExchangeIdRequestTargetIdTypeRestId,
		"restimmutableentryid": CreateTranslateExchangeIdRequestTargetIdTypeRestImmutableEntryId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateTranslateExchangeIdRequestTargetIdType(input)
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

type GetUserMailTipRequestMailTipsOptions string

const (
	GetUserMailTipRequestMailTipsOptionsAutomaticReplies     GetUserMailTipRequestMailTipsOptions = "automaticReplies"
	GetUserMailTipRequestMailTipsOptionsCustomMailTip        GetUserMailTipRequestMailTipsOptions = "customMailTip"
	GetUserMailTipRequestMailTipsOptionsDeliveryRestriction  GetUserMailTipRequestMailTipsOptions = "deliveryRestriction"
	GetUserMailTipRequestMailTipsOptionsExternalMemberCount  GetUserMailTipRequestMailTipsOptions = "externalMemberCount"
	GetUserMailTipRequestMailTipsOptionsMailboxFullStatus    GetUserMailTipRequestMailTipsOptions = "mailboxFullStatus"
	GetUserMailTipRequestMailTipsOptionsMaxMessageSize       GetUserMailTipRequestMailTipsOptions = "maxMessageSize"
	GetUserMailTipRequestMailTipsOptionsModerationStatus     GetUserMailTipRequestMailTipsOptions = "moderationStatus"
	GetUserMailTipRequestMailTipsOptionsRecipientScope       GetUserMailTipRequestMailTipsOptions = "recipientScope"
	GetUserMailTipRequestMailTipsOptionsRecipientSuggestions GetUserMailTipRequestMailTipsOptions = "recipientSuggestions"
	GetUserMailTipRequestMailTipsOptionsTotalMemberCount     GetUserMailTipRequestMailTipsOptions = "totalMemberCount"
)

func PossibleValuesForGetUserMailTipRequestMailTipsOptions() []string {
	return []string{
		string(GetUserMailTipRequestMailTipsOptionsAutomaticReplies),
		string(GetUserMailTipRequestMailTipsOptionsCustomMailTip),
		string(GetUserMailTipRequestMailTipsOptionsDeliveryRestriction),
		string(GetUserMailTipRequestMailTipsOptionsExternalMemberCount),
		string(GetUserMailTipRequestMailTipsOptionsMailboxFullStatus),
		string(GetUserMailTipRequestMailTipsOptionsMaxMessageSize),
		string(GetUserMailTipRequestMailTipsOptionsModerationStatus),
		string(GetUserMailTipRequestMailTipsOptionsRecipientScope),
		string(GetUserMailTipRequestMailTipsOptionsRecipientSuggestions),
		string(GetUserMailTipRequestMailTipsOptionsTotalMemberCount),
	}
}

func (s *GetUserMailTipRequestMailTipsOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGetUserMailTipRequestMailTipsOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGetUserMailTipRequestMailTipsOptions(input string) (*GetUserMailTipRequestMailTipsOptions, error) {
	vals := map[string]GetUserMailTipRequestMailTipsOptions{
		"automaticreplies":     GetUserMailTipRequestMailTipsOptionsAutomaticReplies,
		"custommailtip":        GetUserMailTipRequestMailTipsOptionsCustomMailTip,
		"deliveryrestriction":  GetUserMailTipRequestMailTipsOptionsDeliveryRestriction,
		"externalmembercount":  GetUserMailTipRequestMailTipsOptionsExternalMemberCount,
		"mailboxfullstatus":    GetUserMailTipRequestMailTipsOptionsMailboxFullStatus,
		"maxmessagesize":       GetUserMailTipRequestMailTipsOptionsMaxMessageSize,
		"moderationstatus":     GetUserMailTipRequestMailTipsOptionsModerationStatus,
		"recipientscope":       GetUserMailTipRequestMailTipsOptionsRecipientScope,
		"recipientsuggestions": GetUserMailTipRequestMailTipsOptionsRecipientSuggestions,
		"totalmembercount":     GetUserMailTipRequestMailTipsOptionsTotalMemberCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GetUserMailTipRequestMailTipsOptions(input)
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

type LocationConstraintItemLocationType string

const (
	LocationConstraintItemLocationTypeBusinessAddress LocationConstraintItemLocationType = "businessAddress"
	LocationConstraintItemLocationTypeConferenceRoom  LocationConstraintItemLocationType = "conferenceRoom"
	LocationConstraintItemLocationTypeDefault         LocationConstraintItemLocationType = "default"
	LocationConstraintItemLocationTypeGeoCoordinates  LocationConstraintItemLocationType = "geoCoordinates"
	LocationConstraintItemLocationTypeHomeAddress     LocationConstraintItemLocationType = "homeAddress"
	LocationConstraintItemLocationTypeHotel           LocationConstraintItemLocationType = "hotel"
	LocationConstraintItemLocationTypeLocalBusiness   LocationConstraintItemLocationType = "localBusiness"
	LocationConstraintItemLocationTypePostalAddress   LocationConstraintItemLocationType = "postalAddress"
	LocationConstraintItemLocationTypeRestaurant      LocationConstraintItemLocationType = "restaurant"
	LocationConstraintItemLocationTypeStreetAddress   LocationConstraintItemLocationType = "streetAddress"
)

func PossibleValuesForLocationConstraintItemLocationType() []string {
	return []string{
		string(LocationConstraintItemLocationTypeBusinessAddress),
		string(LocationConstraintItemLocationTypeConferenceRoom),
		string(LocationConstraintItemLocationTypeDefault),
		string(LocationConstraintItemLocationTypeGeoCoordinates),
		string(LocationConstraintItemLocationTypeHomeAddress),
		string(LocationConstraintItemLocationTypeHotel),
		string(LocationConstraintItemLocationTypeLocalBusiness),
		string(LocationConstraintItemLocationTypePostalAddress),
		string(LocationConstraintItemLocationTypeRestaurant),
		string(LocationConstraintItemLocationTypeStreetAddress),
	}
}

func (s *LocationConstraintItemLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationConstraintItemLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationConstraintItemLocationType(input string) (*LocationConstraintItemLocationType, error) {
	vals := map[string]LocationConstraintItemLocationType{
		"businessaddress": LocationConstraintItemLocationTypeBusinessAddress,
		"conferenceroom":  LocationConstraintItemLocationTypeConferenceRoom,
		"default":         LocationConstraintItemLocationTypeDefault,
		"geocoordinates":  LocationConstraintItemLocationTypeGeoCoordinates,
		"homeaddress":     LocationConstraintItemLocationTypeHomeAddress,
		"hotel":           LocationConstraintItemLocationTypeHotel,
		"localbusiness":   LocationConstraintItemLocationTypeLocalBusiness,
		"postaladdress":   LocationConstraintItemLocationTypePostalAddress,
		"restaurant":      LocationConstraintItemLocationTypeRestaurant,
		"streetaddress":   LocationConstraintItemLocationTypeStreetAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationConstraintItemLocationType(input)
	return &out, nil
}

type LocationConstraintItemUniqueIdType string

const (
	LocationConstraintItemUniqueIdTypeBing          LocationConstraintItemUniqueIdType = "bing"
	LocationConstraintItemUniqueIdTypeDirectory     LocationConstraintItemUniqueIdType = "directory"
	LocationConstraintItemUniqueIdTypeLocationStore LocationConstraintItemUniqueIdType = "locationStore"
	LocationConstraintItemUniqueIdTypePrivate       LocationConstraintItemUniqueIdType = "private"
	LocationConstraintItemUniqueIdTypeUnknown       LocationConstraintItemUniqueIdType = "unknown"
)

func PossibleValuesForLocationConstraintItemUniqueIdType() []string {
	return []string{
		string(LocationConstraintItemUniqueIdTypeBing),
		string(LocationConstraintItemUniqueIdTypeDirectory),
		string(LocationConstraintItemUniqueIdTypeLocationStore),
		string(LocationConstraintItemUniqueIdTypePrivate),
		string(LocationConstraintItemUniqueIdTypeUnknown),
	}
}

func (s *LocationConstraintItemUniqueIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationConstraintItemUniqueIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationConstraintItemUniqueIdType(input string) (*LocationConstraintItemUniqueIdType, error) {
	vals := map[string]LocationConstraintItemUniqueIdType{
		"bing":          LocationConstraintItemUniqueIdTypeBing,
		"directory":     LocationConstraintItemUniqueIdTypeDirectory,
		"locationstore": LocationConstraintItemUniqueIdTypeLocationStore,
		"private":       LocationConstraintItemUniqueIdTypePrivate,
		"unknown":       LocationConstraintItemUniqueIdTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationConstraintItemUniqueIdType(input)
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

type PhysicalAddressType string

const (
	PhysicalAddressTypeBusiness PhysicalAddressType = "business"
	PhysicalAddressTypeHome     PhysicalAddressType = "home"
	PhysicalAddressTypeOther    PhysicalAddressType = "other"
	PhysicalAddressTypeUnknown  PhysicalAddressType = "unknown"
)

func PossibleValuesForPhysicalAddressType() []string {
	return []string{
		string(PhysicalAddressTypeBusiness),
		string(PhysicalAddressTypeHome),
		string(PhysicalAddressTypeOther),
		string(PhysicalAddressTypeUnknown),
	}
}

func (s *PhysicalAddressType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhysicalAddressType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhysicalAddressType(input string) (*PhysicalAddressType, error) {
	vals := map[string]PhysicalAddressType{
		"business": PhysicalAddressTypeBusiness,
		"home":     PhysicalAddressTypeHome,
		"other":    PhysicalAddressTypeOther,
		"unknown":  PhysicalAddressTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhysicalAddressType(input)
	return &out, nil
}

type TimeConstraintActivityDomain string

const (
	TimeConstraintActivityDomainPersonal     TimeConstraintActivityDomain = "personal"
	TimeConstraintActivityDomainUnknown      TimeConstraintActivityDomain = "unknown"
	TimeConstraintActivityDomainUnrestricted TimeConstraintActivityDomain = "unrestricted"
	TimeConstraintActivityDomainWork         TimeConstraintActivityDomain = "work"
)

func PossibleValuesForTimeConstraintActivityDomain() []string {
	return []string{
		string(TimeConstraintActivityDomainPersonal),
		string(TimeConstraintActivityDomainUnknown),
		string(TimeConstraintActivityDomainUnrestricted),
		string(TimeConstraintActivityDomainWork),
	}
}

func (s *TimeConstraintActivityDomain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeConstraintActivityDomain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeConstraintActivityDomain(input string) (*TimeConstraintActivityDomain, error) {
	vals := map[string]TimeConstraintActivityDomain{
		"personal":     TimeConstraintActivityDomainPersonal,
		"unknown":      TimeConstraintActivityDomainUnknown,
		"unrestricted": TimeConstraintActivityDomainUnrestricted,
		"work":         TimeConstraintActivityDomainWork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeConstraintActivityDomain(input)
	return &out, nil
}
