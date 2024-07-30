package site

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

type ActivityStatisticsActivity string

const (
	ActivityStatisticsActivityCall    ActivityStatisticsActivity = "Call"
	ActivityStatisticsActivityChat    ActivityStatisticsActivity = "Chat"
	ActivityStatisticsActivityEmail   ActivityStatisticsActivity = "Email"
	ActivityStatisticsActivityFocus   ActivityStatisticsActivity = "Focus"
	ActivityStatisticsActivityMeeting ActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForActivityStatisticsActivity() []string {
	return []string{
		string(ActivityStatisticsActivityCall),
		string(ActivityStatisticsActivityChat),
		string(ActivityStatisticsActivityEmail),
		string(ActivityStatisticsActivityFocus),
		string(ActivityStatisticsActivityMeeting),
	}
}

func (s *ActivityStatisticsActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActivityStatisticsActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActivityStatisticsActivity(input string) (*ActivityStatisticsActivity, error) {
	vals := map[string]ActivityStatisticsActivity{
		"call":    ActivityStatisticsActivityCall,
		"chat":    ActivityStatisticsActivityChat,
		"email":   ActivityStatisticsActivityEmail,
		"focus":   ActivityStatisticsActivityFocus,
		"meeting": ActivityStatisticsActivityMeeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityStatisticsActivity(input)
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

type AppLogCollectionRequestStatus string

const (
	AppLogCollectionRequestStatusCompleted AppLogCollectionRequestStatus = "completed"
	AppLogCollectionRequestStatusFailed    AppLogCollectionRequestStatus = "failed"
	AppLogCollectionRequestStatusPending   AppLogCollectionRequestStatus = "pending"
)

func PossibleValuesForAppLogCollectionRequestStatus() []string {
	return []string{
		string(AppLogCollectionRequestStatusCompleted),
		string(AppLogCollectionRequestStatusFailed),
		string(AppLogCollectionRequestStatusPending),
	}
}

func (s *AppLogCollectionRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppLogCollectionRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppLogCollectionRequestStatus(input string) (*AppLogCollectionRequestStatus, error) {
	vals := map[string]AppLogCollectionRequestStatus{
		"completed": AppLogCollectionRequestStatusCompleted,
		"failed":    AppLogCollectionRequestStatusFailed,
		"pending":   AppLogCollectionRequestStatusPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLogCollectionRequestStatus(input)
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

type AttributeDefinitionMetadataEntryKey string

const (
	AttributeDefinitionMetadataEntryKeyBaseAttributeName       AttributeDefinitionMetadataEntryKey = "BaseAttributeName"
	AttributeDefinitionMetadataEntryKeyComplexObjectDefinition AttributeDefinitionMetadataEntryKey = "ComplexObjectDefinition"
	AttributeDefinitionMetadataEntryKeyIsContainer             AttributeDefinitionMetadataEntryKey = "IsContainer"
	AttributeDefinitionMetadataEntryKeyIsCustomerDefined       AttributeDefinitionMetadataEntryKey = "IsCustomerDefined"
	AttributeDefinitionMetadataEntryKeyIsDomainQualified       AttributeDefinitionMetadataEntryKey = "IsDomainQualified"
	AttributeDefinitionMetadataEntryKeyLinkPropertyNames       AttributeDefinitionMetadataEntryKey = "LinkPropertyNames"
	AttributeDefinitionMetadataEntryKeyLinkTypeName            AttributeDefinitionMetadataEntryKey = "LinkTypeName"
	AttributeDefinitionMetadataEntryKeyMaximumLength           AttributeDefinitionMetadataEntryKey = "MaximumLength"
	AttributeDefinitionMetadataEntryKeyReferencedProperty      AttributeDefinitionMetadataEntryKey = "ReferencedProperty"
)

func PossibleValuesForAttributeDefinitionMetadataEntryKey() []string {
	return []string{
		string(AttributeDefinitionMetadataEntryKeyBaseAttributeName),
		string(AttributeDefinitionMetadataEntryKeyComplexObjectDefinition),
		string(AttributeDefinitionMetadataEntryKeyIsContainer),
		string(AttributeDefinitionMetadataEntryKeyIsCustomerDefined),
		string(AttributeDefinitionMetadataEntryKeyIsDomainQualified),
		string(AttributeDefinitionMetadataEntryKeyLinkPropertyNames),
		string(AttributeDefinitionMetadataEntryKeyLinkTypeName),
		string(AttributeDefinitionMetadataEntryKeyMaximumLength),
		string(AttributeDefinitionMetadataEntryKeyReferencedProperty),
	}
}

func (s *AttributeDefinitionMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionMetadataEntryKey(input string) (*AttributeDefinitionMetadataEntryKey, error) {
	vals := map[string]AttributeDefinitionMetadataEntryKey{
		"baseattributename":       AttributeDefinitionMetadataEntryKeyBaseAttributeName,
		"complexobjectdefinition": AttributeDefinitionMetadataEntryKeyComplexObjectDefinition,
		"iscontainer":             AttributeDefinitionMetadataEntryKeyIsContainer,
		"iscustomerdefined":       AttributeDefinitionMetadataEntryKeyIsCustomerDefined,
		"isdomainqualified":       AttributeDefinitionMetadataEntryKeyIsDomainQualified,
		"linkpropertynames":       AttributeDefinitionMetadataEntryKeyLinkPropertyNames,
		"linktypename":            AttributeDefinitionMetadataEntryKeyLinkTypeName,
		"maximumlength":           AttributeDefinitionMetadataEntryKeyMaximumLength,
		"referencedproperty":      AttributeDefinitionMetadataEntryKeyReferencedProperty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMetadataEntryKey(input)
	return &out, nil
}

type AttributeDefinitionMutability string

const (
	AttributeDefinitionMutabilityImmutable AttributeDefinitionMutability = "Immutable"
	AttributeDefinitionMutabilityReadOnly  AttributeDefinitionMutability = "ReadOnly"
	AttributeDefinitionMutabilityReadWrite AttributeDefinitionMutability = "ReadWrite"
	AttributeDefinitionMutabilityWriteOnly AttributeDefinitionMutability = "WriteOnly"
)

func PossibleValuesForAttributeDefinitionMutability() []string {
	return []string{
		string(AttributeDefinitionMutabilityImmutable),
		string(AttributeDefinitionMutabilityReadOnly),
		string(AttributeDefinitionMutabilityReadWrite),
		string(AttributeDefinitionMutabilityWriteOnly),
	}
}

func (s *AttributeDefinitionMutability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionMutability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionMutability(input string) (*AttributeDefinitionMutability, error) {
	vals := map[string]AttributeDefinitionMutability{
		"immutable": AttributeDefinitionMutabilityImmutable,
		"readonly":  AttributeDefinitionMutabilityReadOnly,
		"readwrite": AttributeDefinitionMutabilityReadWrite,
		"writeonly": AttributeDefinitionMutabilityWriteOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMutability(input)
	return &out, nil
}

type AttributeDefinitionType string

const (
	AttributeDefinitionTypeBinary    AttributeDefinitionType = "Binary"
	AttributeDefinitionTypeBoolean   AttributeDefinitionType = "Boolean"
	AttributeDefinitionTypeDateTime  AttributeDefinitionType = "DateTime"
	AttributeDefinitionTypeInteger   AttributeDefinitionType = "Integer"
	AttributeDefinitionTypeReference AttributeDefinitionType = "Reference"
	AttributeDefinitionTypeString    AttributeDefinitionType = "String"
)

func PossibleValuesForAttributeDefinitionType() []string {
	return []string{
		string(AttributeDefinitionTypeBinary),
		string(AttributeDefinitionTypeBoolean),
		string(AttributeDefinitionTypeDateTime),
		string(AttributeDefinitionTypeInteger),
		string(AttributeDefinitionTypeReference),
		string(AttributeDefinitionTypeString),
	}
}

func (s *AttributeDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionType(input string) (*AttributeDefinitionType, error) {
	vals := map[string]AttributeDefinitionType{
		"binary":    AttributeDefinitionTypeBinary,
		"boolean":   AttributeDefinitionTypeBoolean,
		"datetime":  AttributeDefinitionTypeDateTime,
		"integer":   AttributeDefinitionTypeInteger,
		"reference": AttributeDefinitionTypeReference,
		"string":    AttributeDefinitionTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionType(input)
	return &out, nil
}

type AttributeMappingFlowBehavior string

const (
	AttributeMappingFlowBehaviorFlowAlways      AttributeMappingFlowBehavior = "FlowAlways"
	AttributeMappingFlowBehaviorFlowWhenChanged AttributeMappingFlowBehavior = "FlowWhenChanged"
)

func PossibleValuesForAttributeMappingFlowBehavior() []string {
	return []string{
		string(AttributeMappingFlowBehaviorFlowAlways),
		string(AttributeMappingFlowBehaviorFlowWhenChanged),
	}
}

func (s *AttributeMappingFlowBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeMappingFlowBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeMappingFlowBehavior(input string) (*AttributeMappingFlowBehavior, error) {
	vals := map[string]AttributeMappingFlowBehavior{
		"flowalways":      AttributeMappingFlowBehaviorFlowAlways,
		"flowwhenchanged": AttributeMappingFlowBehaviorFlowWhenChanged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingFlowBehavior(input)
	return &out, nil
}

type AttributeMappingFlowType string

const (
	AttributeMappingFlowTypeAlways            AttributeMappingFlowType = "Always"
	AttributeMappingFlowTypeAttributeAddOnly  AttributeMappingFlowType = "AttributeAddOnly"
	AttributeMappingFlowTypeMultiValueAddOnly AttributeMappingFlowType = "MultiValueAddOnly"
	AttributeMappingFlowTypeObjectAddOnly     AttributeMappingFlowType = "ObjectAddOnly"
	AttributeMappingFlowTypeValueAddOnly      AttributeMappingFlowType = "ValueAddOnly"
)

func PossibleValuesForAttributeMappingFlowType() []string {
	return []string{
		string(AttributeMappingFlowTypeAlways),
		string(AttributeMappingFlowTypeAttributeAddOnly),
		string(AttributeMappingFlowTypeMultiValueAddOnly),
		string(AttributeMappingFlowTypeObjectAddOnly),
		string(AttributeMappingFlowTypeValueAddOnly),
	}
}

func (s *AttributeMappingFlowType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeMappingFlowType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeMappingFlowType(input string) (*AttributeMappingFlowType, error) {
	vals := map[string]AttributeMappingFlowType{
		"always":            AttributeMappingFlowTypeAlways,
		"attributeaddonly":  AttributeMappingFlowTypeAttributeAddOnly,
		"multivalueaddonly": AttributeMappingFlowTypeMultiValueAddOnly,
		"objectaddonly":     AttributeMappingFlowTypeObjectAddOnly,
		"valueaddonly":      AttributeMappingFlowTypeValueAddOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingFlowType(input)
	return &out, nil
}

type AttributeMappingSourceType string

const (
	AttributeMappingSourceTypeAttribute AttributeMappingSourceType = "Attribute"
	AttributeMappingSourceTypeConstant  AttributeMappingSourceType = "Constant"
	AttributeMappingSourceTypeFunction  AttributeMappingSourceType = "Function"
)

func PossibleValuesForAttributeMappingSourceType() []string {
	return []string{
		string(AttributeMappingSourceTypeAttribute),
		string(AttributeMappingSourceTypeConstant),
		string(AttributeMappingSourceTypeFunction),
	}
}

func (s *AttributeMappingSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeMappingSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeMappingSourceType(input string) (*AttributeMappingSourceType, error) {
	vals := map[string]AttributeMappingSourceType{
		"attribute": AttributeMappingSourceTypeAttribute,
		"constant":  AttributeMappingSourceTypeConstant,
		"function":  AttributeMappingSourceTypeFunction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingSourceType(input)
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

type BaseSitePagePageLayout string

const (
	BaseSitePagePageLayoutArticle           BaseSitePagePageLayout = "article"
	BaseSitePagePageLayoutHome              BaseSitePagePageLayout = "home"
	BaseSitePagePageLayoutMicrosoftReserved BaseSitePagePageLayout = "microsoftReserved"
)

func PossibleValuesForBaseSitePagePageLayout() []string {
	return []string{
		string(BaseSitePagePageLayoutArticle),
		string(BaseSitePagePageLayoutHome),
		string(BaseSitePagePageLayoutMicrosoftReserved),
	}
}

func (s *BaseSitePagePageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBaseSitePagePageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBaseSitePagePageLayout(input string) (*BaseSitePagePageLayout, error) {
	vals := map[string]BaseSitePagePageLayout{
		"article":           BaseSitePagePageLayoutArticle,
		"home":              BaseSitePagePageLayoutHome,
		"microsoftreserved": BaseSitePagePageLayoutMicrosoftReserved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BaseSitePagePageLayout(input)
	return &out, nil
}

type BitlockerRecoveryKeyVolumeType string

const (
	BitlockerRecoveryKeyVolumeTypeFixedDataVolume       BitlockerRecoveryKeyVolumeType = "fixedDataVolume"
	BitlockerRecoveryKeyVolumeTypeOperatingSystemVolume BitlockerRecoveryKeyVolumeType = "operatingSystemVolume"
	BitlockerRecoveryKeyVolumeTypeRemovableDataVolume   BitlockerRecoveryKeyVolumeType = "removableDataVolume"
)

func PossibleValuesForBitlockerRecoveryKeyVolumeType() []string {
	return []string{
		string(BitlockerRecoveryKeyVolumeTypeFixedDataVolume),
		string(BitlockerRecoveryKeyVolumeTypeOperatingSystemVolume),
		string(BitlockerRecoveryKeyVolumeTypeRemovableDataVolume),
	}
}

func (s *BitlockerRecoveryKeyVolumeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitlockerRecoveryKeyVolumeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitlockerRecoveryKeyVolumeType(input string) (*BitlockerRecoveryKeyVolumeType, error) {
	vals := map[string]BitlockerRecoveryKeyVolumeType{
		"fixeddatavolume":       BitlockerRecoveryKeyVolumeTypeFixedDataVolume,
		"operatingsystemvolume": BitlockerRecoveryKeyVolumeTypeOperatingSystemVolume,
		"removabledatavolume":   BitlockerRecoveryKeyVolumeTypeRemovableDataVolume,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitlockerRecoveryKeyVolumeType(input)
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

type ChannelModerationSettingsReplyRestriction string

const (
	ChannelModerationSettingsReplyRestrictionAuthorAndModerators ChannelModerationSettingsReplyRestriction = "authorAndModerators"
	ChannelModerationSettingsReplyRestrictionEveryone            ChannelModerationSettingsReplyRestriction = "everyone"
)

func PossibleValuesForChannelModerationSettingsReplyRestriction() []string {
	return []string{
		string(ChannelModerationSettingsReplyRestrictionAuthorAndModerators),
		string(ChannelModerationSettingsReplyRestrictionEveryone),
	}
}

func (s *ChannelModerationSettingsReplyRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChannelModerationSettingsReplyRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChannelModerationSettingsReplyRestriction(input string) (*ChannelModerationSettingsReplyRestriction, error) {
	vals := map[string]ChannelModerationSettingsReplyRestriction{
		"authorandmoderators": ChannelModerationSettingsReplyRestrictionAuthorAndModerators,
		"everyone":            ChannelModerationSettingsReplyRestrictionEveryone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelModerationSettingsReplyRestriction(input)
	return &out, nil
}

type ChannelModerationSettingsUserNewMessageRestriction string

const (
	ChannelModerationSettingsUserNewMessageRestrictionEveryone             ChannelModerationSettingsUserNewMessageRestriction = "everyone"
	ChannelModerationSettingsUserNewMessageRestrictionEveryoneExceptGuests ChannelModerationSettingsUserNewMessageRestriction = "everyoneExceptGuests"
	ChannelModerationSettingsUserNewMessageRestrictionModerators           ChannelModerationSettingsUserNewMessageRestriction = "moderators"
)

func PossibleValuesForChannelModerationSettingsUserNewMessageRestriction() []string {
	return []string{
		string(ChannelModerationSettingsUserNewMessageRestrictionEveryone),
		string(ChannelModerationSettingsUserNewMessageRestrictionEveryoneExceptGuests),
		string(ChannelModerationSettingsUserNewMessageRestrictionModerators),
	}
}

func (s *ChannelModerationSettingsUserNewMessageRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChannelModerationSettingsUserNewMessageRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChannelModerationSettingsUserNewMessageRestriction(input string) (*ChannelModerationSettingsUserNewMessageRestriction, error) {
	vals := map[string]ChannelModerationSettingsUserNewMessageRestriction{
		"everyone":             ChannelModerationSettingsUserNewMessageRestrictionEveryone,
		"everyoneexceptguests": ChannelModerationSettingsUserNewMessageRestrictionEveryoneExceptGuests,
		"moderators":           ChannelModerationSettingsUserNewMessageRestrictionModerators,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChannelModerationSettingsUserNewMessageRestriction(input)
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

type CloudPCConnectivityResultStatus string

const (
	CloudPCConnectivityResultStatusAvailable            CloudPCConnectivityResultStatus = "available"
	CloudPCConnectivityResultStatusAvailableWithWarning CloudPCConnectivityResultStatus = "availableWithWarning"
	CloudPCConnectivityResultStatusUnavailable          CloudPCConnectivityResultStatus = "unavailable"
	CloudPCConnectivityResultStatusUnknown              CloudPCConnectivityResultStatus = "unknown"
)

func PossibleValuesForCloudPCConnectivityResultStatus() []string {
	return []string{
		string(CloudPCConnectivityResultStatusAvailable),
		string(CloudPCConnectivityResultStatusAvailableWithWarning),
		string(CloudPCConnectivityResultStatusUnavailable),
		string(CloudPCConnectivityResultStatusUnknown),
	}
}

func (s *CloudPCConnectivityResultStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectivityResultStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectivityResultStatus(input string) (*CloudPCConnectivityResultStatus, error) {
	vals := map[string]CloudPCConnectivityResultStatus{
		"available":            CloudPCConnectivityResultStatusAvailable,
		"availablewithwarning": CloudPCConnectivityResultStatusAvailableWithWarning,
		"unavailable":          CloudPCConnectivityResultStatusUnavailable,
		"unknown":              CloudPCConnectivityResultStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityResultStatus(input)
	return &out, nil
}

type CloudPCDiskEncryptionState string

const (
	CloudPCDiskEncryptionStateEncryptedUsingCustomerManagedKey CloudPCDiskEncryptionState = "encryptedUsingCustomerManagedKey"
	CloudPCDiskEncryptionStateEncryptedUsingPlatformManagedKey CloudPCDiskEncryptionState = "encryptedUsingPlatformManagedKey"
	CloudPCDiskEncryptionStateNotAvailable                     CloudPCDiskEncryptionState = "notAvailable"
	CloudPCDiskEncryptionStateNotEncrypted                     CloudPCDiskEncryptionState = "notEncrypted"
)

func PossibleValuesForCloudPCDiskEncryptionState() []string {
	return []string{
		string(CloudPCDiskEncryptionStateEncryptedUsingCustomerManagedKey),
		string(CloudPCDiskEncryptionStateEncryptedUsingPlatformManagedKey),
		string(CloudPCDiskEncryptionStateNotAvailable),
		string(CloudPCDiskEncryptionStateNotEncrypted),
	}
}

func (s *CloudPCDiskEncryptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDiskEncryptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDiskEncryptionState(input string) (*CloudPCDiskEncryptionState, error) {
	vals := map[string]CloudPCDiskEncryptionState{
		"encryptedusingcustomermanagedkey": CloudPCDiskEncryptionStateEncryptedUsingCustomerManagedKey,
		"encryptedusingplatformmanagedkey": CloudPCDiskEncryptionStateEncryptedUsingPlatformManagedKey,
		"notavailable":                     CloudPCDiskEncryptionStateNotAvailable,
		"notencrypted":                     CloudPCDiskEncryptionStateNotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDiskEncryptionState(input)
	return &out, nil
}

type CloudPCHealthCheckItemResult string

const (
	CloudPCHealthCheckItemResultFailure CloudPCHealthCheckItemResult = "failure"
	CloudPCHealthCheckItemResultSuccess CloudPCHealthCheckItemResult = "success"
	CloudPCHealthCheckItemResultUnknown CloudPCHealthCheckItemResult = "unknown"
)

func PossibleValuesForCloudPCHealthCheckItemResult() []string {
	return []string{
		string(CloudPCHealthCheckItemResultFailure),
		string(CloudPCHealthCheckItemResultSuccess),
		string(CloudPCHealthCheckItemResultUnknown),
	}
}

func (s *CloudPCHealthCheckItemResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCHealthCheckItemResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCHealthCheckItemResult(input string) (*CloudPCHealthCheckItemResult, error) {
	vals := map[string]CloudPCHealthCheckItemResult{
		"failure": CloudPCHealthCheckItemResultFailure,
		"success": CloudPCHealthCheckItemResultSuccess,
		"unknown": CloudPCHealthCheckItemResultUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCHealthCheckItemResult(input)
	return &out, nil
}

type CloudPCOsVersion string

const (
	CloudPCOsVersionWindows10 CloudPCOsVersion = "windows10"
	CloudPCOsVersionWindows11 CloudPCOsVersion = "windows11"
)

func PossibleValuesForCloudPCOsVersion() []string {
	return []string{
		string(CloudPCOsVersionWindows10),
		string(CloudPCOsVersionWindows11),
	}
}

func (s *CloudPCOsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOsVersion(input string) (*CloudPCOsVersion, error) {
	vals := map[string]CloudPCOsVersion{
		"windows10": CloudPCOsVersionWindows10,
		"windows11": CloudPCOsVersionWindows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOsVersion(input)
	return &out, nil
}

type CloudPCPartnerAgentInstallResultInstallStatus string

const (
	CloudPCPartnerAgentInstallResultInstallStatusInstallFailed   CloudPCPartnerAgentInstallResultInstallStatus = "installFailed"
	CloudPCPartnerAgentInstallResultInstallStatusInstalled       CloudPCPartnerAgentInstallResultInstallStatus = "installed"
	CloudPCPartnerAgentInstallResultInstallStatusInstalling      CloudPCPartnerAgentInstallResultInstallStatus = "installing"
	CloudPCPartnerAgentInstallResultInstallStatusLicensed        CloudPCPartnerAgentInstallResultInstallStatus = "licensed"
	CloudPCPartnerAgentInstallResultInstallStatusUninstallFailed CloudPCPartnerAgentInstallResultInstallStatus = "uninstallFailed"
	CloudPCPartnerAgentInstallResultInstallStatusUninstalling    CloudPCPartnerAgentInstallResultInstallStatus = "uninstalling"
)

func PossibleValuesForCloudPCPartnerAgentInstallResultInstallStatus() []string {
	return []string{
		string(CloudPCPartnerAgentInstallResultInstallStatusInstallFailed),
		string(CloudPCPartnerAgentInstallResultInstallStatusInstalled),
		string(CloudPCPartnerAgentInstallResultInstallStatusInstalling),
		string(CloudPCPartnerAgentInstallResultInstallStatusLicensed),
		string(CloudPCPartnerAgentInstallResultInstallStatusUninstallFailed),
		string(CloudPCPartnerAgentInstallResultInstallStatusUninstalling),
	}
}

func (s *CloudPCPartnerAgentInstallResultInstallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPartnerAgentInstallResultInstallStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPartnerAgentInstallResultInstallStatus(input string) (*CloudPCPartnerAgentInstallResultInstallStatus, error) {
	vals := map[string]CloudPCPartnerAgentInstallResultInstallStatus{
		"installfailed":   CloudPCPartnerAgentInstallResultInstallStatusInstallFailed,
		"installed":       CloudPCPartnerAgentInstallResultInstallStatusInstalled,
		"installing":      CloudPCPartnerAgentInstallResultInstallStatusInstalling,
		"licensed":        CloudPCPartnerAgentInstallResultInstallStatusLicensed,
		"uninstallfailed": CloudPCPartnerAgentInstallResultInstallStatusUninstallFailed,
		"uninstalling":    CloudPCPartnerAgentInstallResultInstallStatusUninstalling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentInstallResultInstallStatus(input)
	return &out, nil
}

type CloudPCPartnerAgentInstallResultPartnerAgentName string

const (
	CloudPCPartnerAgentInstallResultPartnerAgentNameCitrix CloudPCPartnerAgentInstallResultPartnerAgentName = "citrix"
	CloudPCPartnerAgentInstallResultPartnerAgentNameHp     CloudPCPartnerAgentInstallResultPartnerAgentName = "hp"
	CloudPCPartnerAgentInstallResultPartnerAgentNameVMware CloudPCPartnerAgentInstallResultPartnerAgentName = "vMware"
)

func PossibleValuesForCloudPCPartnerAgentInstallResultPartnerAgentName() []string {
	return []string{
		string(CloudPCPartnerAgentInstallResultPartnerAgentNameCitrix),
		string(CloudPCPartnerAgentInstallResultPartnerAgentNameHp),
		string(CloudPCPartnerAgentInstallResultPartnerAgentNameVMware),
	}
}

func (s *CloudPCPartnerAgentInstallResultPartnerAgentName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPartnerAgentInstallResultPartnerAgentName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPartnerAgentInstallResultPartnerAgentName(input string) (*CloudPCPartnerAgentInstallResultPartnerAgentName, error) {
	vals := map[string]CloudPCPartnerAgentInstallResultPartnerAgentName{
		"citrix": CloudPCPartnerAgentInstallResultPartnerAgentNameCitrix,
		"hp":     CloudPCPartnerAgentInstallResultPartnerAgentNameHp,
		"vmware": CloudPCPartnerAgentInstallResultPartnerAgentNameVMware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentInstallResultPartnerAgentName(input)
	return &out, nil
}

type CloudPCPowerState string

const (
	CloudPCPowerStatePoweredOff CloudPCPowerState = "poweredOff"
	CloudPCPowerStateRunning    CloudPCPowerState = "running"
)

func PossibleValuesForCloudPCPowerState() []string {
	return []string{
		string(CloudPCPowerStatePoweredOff),
		string(CloudPCPowerStateRunning),
	}
}

func (s *CloudPCPowerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPowerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPowerState(input string) (*CloudPCPowerState, error) {
	vals := map[string]CloudPCPowerState{
		"poweredoff": CloudPCPowerStatePoweredOff,
		"running":    CloudPCPowerStateRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPowerState(input)
	return &out, nil
}

type CloudPCProvisioningType string

const (
	CloudPCProvisioningTypeDedicated CloudPCProvisioningType = "dedicated"
	CloudPCProvisioningTypeShared    CloudPCProvisioningType = "shared"
)

func PossibleValuesForCloudPCProvisioningType() []string {
	return []string{
		string(CloudPCProvisioningTypeDedicated),
		string(CloudPCProvisioningTypeShared),
	}
}

func (s *CloudPCProvisioningType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCProvisioningType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCProvisioningType(input string) (*CloudPCProvisioningType, error) {
	vals := map[string]CloudPCProvisioningType{
		"dedicated": CloudPCProvisioningTypeDedicated,
		"shared":    CloudPCProvisioningTypeShared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningType(input)
	return &out, nil
}

type CloudPCRemoteActionResultActionState string

const (
	CloudPCRemoteActionResultActionStateActive       CloudPCRemoteActionResultActionState = "active"
	CloudPCRemoteActionResultActionStateCanceled     CloudPCRemoteActionResultActionState = "canceled"
	CloudPCRemoteActionResultActionStateDone         CloudPCRemoteActionResultActionState = "done"
	CloudPCRemoteActionResultActionStateFailed       CloudPCRemoteActionResultActionState = "failed"
	CloudPCRemoteActionResultActionStateNone         CloudPCRemoteActionResultActionState = "none"
	CloudPCRemoteActionResultActionStateNotSupported CloudPCRemoteActionResultActionState = "notSupported"
	CloudPCRemoteActionResultActionStatePending      CloudPCRemoteActionResultActionState = "pending"
)

func PossibleValuesForCloudPCRemoteActionResultActionState() []string {
	return []string{
		string(CloudPCRemoteActionResultActionStateActive),
		string(CloudPCRemoteActionResultActionStateCanceled),
		string(CloudPCRemoteActionResultActionStateDone),
		string(CloudPCRemoteActionResultActionStateFailed),
		string(CloudPCRemoteActionResultActionStateNone),
		string(CloudPCRemoteActionResultActionStateNotSupported),
		string(CloudPCRemoteActionResultActionStatePending),
	}
}

func (s *CloudPCRemoteActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRemoteActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRemoteActionResultActionState(input string) (*CloudPCRemoteActionResultActionState, error) {
	vals := map[string]CloudPCRemoteActionResultActionState{
		"active":       CloudPCRemoteActionResultActionStateActive,
		"canceled":     CloudPCRemoteActionResultActionStateCanceled,
		"done":         CloudPCRemoteActionResultActionStateDone,
		"failed":       CloudPCRemoteActionResultActionStateFailed,
		"none":         CloudPCRemoteActionResultActionStateNone,
		"notsupported": CloudPCRemoteActionResultActionStateNotSupported,
		"pending":      CloudPCRemoteActionResultActionStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionResultActionState(input)
	return &out, nil
}

type CloudPCServicePlanType string

const (
	CloudPCServicePlanTypeBusiness   CloudPCServicePlanType = "business"
	CloudPCServicePlanTypeEnterprise CloudPCServicePlanType = "enterprise"
)

func PossibleValuesForCloudPCServicePlanType() []string {
	return []string{
		string(CloudPCServicePlanTypeBusiness),
		string(CloudPCServicePlanTypeEnterprise),
	}
}

func (s *CloudPCServicePlanType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCServicePlanType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCServicePlanType(input string) (*CloudPCServicePlanType, error) {
	vals := map[string]CloudPCServicePlanType{
		"business":   CloudPCServicePlanTypeBusiness,
		"enterprise": CloudPCServicePlanTypeEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCServicePlanType(input)
	return &out, nil
}

type CloudPCStatus string

const (
	CloudPCStatusDeprovisioning          CloudPCStatus = "deprovisioning"
	CloudPCStatusFailed                  CloudPCStatus = "failed"
	CloudPCStatusInGracePeriod           CloudPCStatus = "inGracePeriod"
	CloudPCStatusMovingRegion            CloudPCStatus = "movingRegion"
	CloudPCStatusNotProvisioned          CloudPCStatus = "notProvisioned"
	CloudPCStatusPendingProvision        CloudPCStatus = "pendingProvision"
	CloudPCStatusProvisioned             CloudPCStatus = "provisioned"
	CloudPCStatusProvisionedWithWarnings CloudPCStatus = "provisionedWithWarnings"
	CloudPCStatusProvisioning            CloudPCStatus = "provisioning"
	CloudPCStatusResizePendingLicense    CloudPCStatus = "resizePendingLicense"
	CloudPCStatusResizing                CloudPCStatus = "resizing"
	CloudPCStatusRestoring               CloudPCStatus = "restoring"
	CloudPCStatusUpdatingSingleSignOn    CloudPCStatus = "updatingSingleSignOn"
)

func PossibleValuesForCloudPCStatus() []string {
	return []string{
		string(CloudPCStatusDeprovisioning),
		string(CloudPCStatusFailed),
		string(CloudPCStatusInGracePeriod),
		string(CloudPCStatusMovingRegion),
		string(CloudPCStatusNotProvisioned),
		string(CloudPCStatusPendingProvision),
		string(CloudPCStatusProvisioned),
		string(CloudPCStatusProvisionedWithWarnings),
		string(CloudPCStatusProvisioning),
		string(CloudPCStatusResizePendingLicense),
		string(CloudPCStatusResizing),
		string(CloudPCStatusRestoring),
		string(CloudPCStatusUpdatingSingleSignOn),
	}
}

func (s *CloudPCStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCStatus(input string) (*CloudPCStatus, error) {
	vals := map[string]CloudPCStatus{
		"deprovisioning":          CloudPCStatusDeprovisioning,
		"failed":                  CloudPCStatusFailed,
		"ingraceperiod":           CloudPCStatusInGracePeriod,
		"movingregion":            CloudPCStatusMovingRegion,
		"notprovisioned":          CloudPCStatusNotProvisioned,
		"pendingprovision":        CloudPCStatusPendingProvision,
		"provisioned":             CloudPCStatusProvisioned,
		"provisionedwithwarnings": CloudPCStatusProvisionedWithWarnings,
		"provisioning":            CloudPCStatusProvisioning,
		"resizependinglicense":    CloudPCStatusResizePendingLicense,
		"resizing":                CloudPCStatusResizing,
		"restoring":               CloudPCStatusRestoring,
		"updatingsinglesignon":    CloudPCStatusUpdatingSingleSignOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCStatus(input)
	return &out, nil
}

type CloudPCUserAccountType string

const (
	CloudPCUserAccountTypeAdministrator CloudPCUserAccountType = "administrator"
	CloudPCUserAccountTypeStandardUser  CloudPCUserAccountType = "standardUser"
)

func PossibleValuesForCloudPCUserAccountType() []string {
	return []string{
		string(CloudPCUserAccountTypeAdministrator),
		string(CloudPCUserAccountTypeStandardUser),
	}
}

func (s *CloudPCUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCUserAccountType(input string) (*CloudPCUserAccountType, error) {
	vals := map[string]CloudPCUserAccountType{
		"administrator": CloudPCUserAccountTypeAdministrator,
		"standarduser":  CloudPCUserAccountTypeStandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCUserAccountType(input)
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

type CommunicationsIdentitySetEndpointType string

const (
	CommunicationsIdentitySetEndpointTypeDefault                   CommunicationsIdentitySetEndpointType = "default"
	CommunicationsIdentitySetEndpointTypeSkypeForBusiness          CommunicationsIdentitySetEndpointType = "skypeForBusiness"
	CommunicationsIdentitySetEndpointTypeSkypeForBusinessVoipPhone CommunicationsIdentitySetEndpointType = "skypeForBusinessVoipPhone"
	CommunicationsIdentitySetEndpointTypeVoicemail                 CommunicationsIdentitySetEndpointType = "voicemail"
)

func PossibleValuesForCommunicationsIdentitySetEndpointType() []string {
	return []string{
		string(CommunicationsIdentitySetEndpointTypeDefault),
		string(CommunicationsIdentitySetEndpointTypeSkypeForBusiness),
		string(CommunicationsIdentitySetEndpointTypeSkypeForBusinessVoipPhone),
		string(CommunicationsIdentitySetEndpointTypeVoicemail),
	}
}

func (s *CommunicationsIdentitySetEndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommunicationsIdentitySetEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommunicationsIdentitySetEndpointType(input string) (*CommunicationsIdentitySetEndpointType, error) {
	vals := map[string]CommunicationsIdentitySetEndpointType{
		"default":                   CommunicationsIdentitySetEndpointTypeDefault,
		"skypeforbusiness":          CommunicationsIdentitySetEndpointTypeSkypeForBusiness,
		"skypeforbusinessvoipphone": CommunicationsIdentitySetEndpointTypeSkypeForBusinessVoipPhone,
		"voicemail":                 CommunicationsIdentitySetEndpointTypeVoicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommunicationsIdentitySetEndpointType(input)
	return &out, nil
}

type ConfigurationManagerClientHealthStateState string

const (
	ConfigurationManagerClientHealthStateStateCommunicationError ConfigurationManagerClientHealthStateState = "communicationError"
	ConfigurationManagerClientHealthStateStateHealthy            ConfigurationManagerClientHealthStateState = "healthy"
	ConfigurationManagerClientHealthStateStateInstallFailed      ConfigurationManagerClientHealthStateState = "installFailed"
	ConfigurationManagerClientHealthStateStateInstalled          ConfigurationManagerClientHealthStateState = "installed"
	ConfigurationManagerClientHealthStateStateUnknown            ConfigurationManagerClientHealthStateState = "unknown"
	ConfigurationManagerClientHealthStateStateUpdateFailed       ConfigurationManagerClientHealthStateState = "updateFailed"
)

func PossibleValuesForConfigurationManagerClientHealthStateState() []string {
	return []string{
		string(ConfigurationManagerClientHealthStateStateCommunicationError),
		string(ConfigurationManagerClientHealthStateStateHealthy),
		string(ConfigurationManagerClientHealthStateStateInstallFailed),
		string(ConfigurationManagerClientHealthStateStateInstalled),
		string(ConfigurationManagerClientHealthStateStateUnknown),
		string(ConfigurationManagerClientHealthStateStateUpdateFailed),
	}
}

func (s *ConfigurationManagerClientHealthStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerClientHealthStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerClientHealthStateState(input string) (*ConfigurationManagerClientHealthStateState, error) {
	vals := map[string]ConfigurationManagerClientHealthStateState{
		"communicationerror": ConfigurationManagerClientHealthStateStateCommunicationError,
		"healthy":            ConfigurationManagerClientHealthStateStateHealthy,
		"installfailed":      ConfigurationManagerClientHealthStateStateInstallFailed,
		"installed":          ConfigurationManagerClientHealthStateStateInstalled,
		"unknown":            ConfigurationManagerClientHealthStateStateUnknown,
		"updatefailed":       ConfigurationManagerClientHealthStateStateUpdateFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerClientHealthStateState(input)
	return &out, nil
}

type DecisionItemPrincipalResourceMembershipMembershipType string

const (
	DecisionItemPrincipalResourceMembershipMembershipTypeDirect   DecisionItemPrincipalResourceMembershipMembershipType = "direct"
	DecisionItemPrincipalResourceMembershipMembershipTypeIndirect DecisionItemPrincipalResourceMembershipMembershipType = "indirect"
)

func PossibleValuesForDecisionItemPrincipalResourceMembershipMembershipType() []string {
	return []string{
		string(DecisionItemPrincipalResourceMembershipMembershipTypeDirect),
		string(DecisionItemPrincipalResourceMembershipMembershipTypeIndirect),
	}
}

func (s *DecisionItemPrincipalResourceMembershipMembershipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDecisionItemPrincipalResourceMembershipMembershipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDecisionItemPrincipalResourceMembershipMembershipType(input string) (*DecisionItemPrincipalResourceMembershipMembershipType, error) {
	vals := map[string]DecisionItemPrincipalResourceMembershipMembershipType{
		"direct":   DecisionItemPrincipalResourceMembershipMembershipTypeDirect,
		"indirect": DecisionItemPrincipalResourceMembershipMembershipTypeIndirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DecisionItemPrincipalResourceMembershipMembershipType(input)
	return &out, nil
}

type DelegatedPermissionClassificationClassification string

const (
	DelegatedPermissionClassificationClassificationHigh   DelegatedPermissionClassificationClassification = "high"
	DelegatedPermissionClassificationClassificationLow    DelegatedPermissionClassificationClassification = "low"
	DelegatedPermissionClassificationClassificationMedium DelegatedPermissionClassificationClassification = "medium"
)

func PossibleValuesForDelegatedPermissionClassificationClassification() []string {
	return []string{
		string(DelegatedPermissionClassificationClassificationHigh),
		string(DelegatedPermissionClassificationClassificationLow),
		string(DelegatedPermissionClassificationClassificationMedium),
	}
}

func (s *DelegatedPermissionClassificationClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedPermissionClassificationClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedPermissionClassificationClassification(input string) (*DelegatedPermissionClassificationClassification, error) {
	vals := map[string]DelegatedPermissionClassificationClassification{
		"high":   DelegatedPermissionClassificationClassificationHigh,
		"low":    DelegatedPermissionClassificationClassificationLow,
		"medium": DelegatedPermissionClassificationClassificationMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedPermissionClassificationClassification(input)
	return &out, nil
}

type DetectedAppPlatform string

const (
	DetectedAppPlatformAndroidDedicatedAndFullyManaged DetectedAppPlatform = "androidDedicatedAndFullyManaged"
	DetectedAppPlatformAndroidDeviceAdministrator      DetectedAppPlatform = "androidDeviceAdministrator"
	DetectedAppPlatformAndroidOSP                      DetectedAppPlatform = "androidOSP"
	DetectedAppPlatformAndroidWorkProfile              DetectedAppPlatform = "androidWorkProfile"
	DetectedAppPlatformChromeOS                        DetectedAppPlatform = "chromeOS"
	DetectedAppPlatformIos                             DetectedAppPlatform = "ios"
	DetectedAppPlatformMacOS                           DetectedAppPlatform = "macOS"
	DetectedAppPlatformUnknown                         DetectedAppPlatform = "unknown"
	DetectedAppPlatformWindows                         DetectedAppPlatform = "windows"
	DetectedAppPlatformWindowsHolographic              DetectedAppPlatform = "windowsHolographic"
	DetectedAppPlatformWindowsMobile                   DetectedAppPlatform = "windowsMobile"
)

func PossibleValuesForDetectedAppPlatform() []string {
	return []string{
		string(DetectedAppPlatformAndroidDedicatedAndFullyManaged),
		string(DetectedAppPlatformAndroidDeviceAdministrator),
		string(DetectedAppPlatformAndroidOSP),
		string(DetectedAppPlatformAndroidWorkProfile),
		string(DetectedAppPlatformChromeOS),
		string(DetectedAppPlatformIos),
		string(DetectedAppPlatformMacOS),
		string(DetectedAppPlatformUnknown),
		string(DetectedAppPlatformWindows),
		string(DetectedAppPlatformWindowsHolographic),
		string(DetectedAppPlatformWindowsMobile),
	}
}

func (s *DetectedAppPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDetectedAppPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDetectedAppPlatform(input string) (*DetectedAppPlatform, error) {
	vals := map[string]DetectedAppPlatform{
		"androiddedicatedandfullymanaged": DetectedAppPlatformAndroidDedicatedAndFullyManaged,
		"androiddeviceadministrator":      DetectedAppPlatformAndroidDeviceAdministrator,
		"androidosp":                      DetectedAppPlatformAndroidOSP,
		"androidworkprofile":              DetectedAppPlatformAndroidWorkProfile,
		"chromeos":                        DetectedAppPlatformChromeOS,
		"ios":                             DetectedAppPlatformIos,
		"macos":                           DetectedAppPlatformMacOS,
		"unknown":                         DetectedAppPlatformUnknown,
		"windows":                         DetectedAppPlatformWindows,
		"windowsholographic":              DetectedAppPlatformWindowsHolographic,
		"windowsmobile":                   DetectedAppPlatformWindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedAppPlatform(input)
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

type DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone),
	}
}

func (s *DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude,
		"include": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude,
		"none":    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
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
	DeviceCompliancePolicyStatePlatformTypeAll                DeviceCompliancePolicyStatePlatformType = "all"
	DeviceCompliancePolicyStatePlatformTypeAndroid            DeviceCompliancePolicyStatePlatformType = "android"
	DeviceCompliancePolicyStatePlatformTypeAndroidAOSP        DeviceCompliancePolicyStatePlatformType = "androidAOSP"
	DeviceCompliancePolicyStatePlatformTypeAndroidForWork     DeviceCompliancePolicyStatePlatformType = "androidForWork"
	DeviceCompliancePolicyStatePlatformTypeAndroidWorkProfile DeviceCompliancePolicyStatePlatformType = "androidWorkProfile"
	DeviceCompliancePolicyStatePlatformTypeIOS                DeviceCompliancePolicyStatePlatformType = "iOS"
	DeviceCompliancePolicyStatePlatformTypeMacOS              DeviceCompliancePolicyStatePlatformType = "macOS"
	DeviceCompliancePolicyStatePlatformTypeWindows10AndLater  DeviceCompliancePolicyStatePlatformType = "windows10AndLater"
	DeviceCompliancePolicyStatePlatformTypeWindows10XProfile  DeviceCompliancePolicyStatePlatformType = "windows10XProfile"
	DeviceCompliancePolicyStatePlatformTypeWindows81AndLater  DeviceCompliancePolicyStatePlatformType = "windows81AndLater"
	DeviceCompliancePolicyStatePlatformTypeWindowsPhone81     DeviceCompliancePolicyStatePlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceCompliancePolicyStatePlatformType() []string {
	return []string{
		string(DeviceCompliancePolicyStatePlatformTypeAll),
		string(DeviceCompliancePolicyStatePlatformTypeAndroid),
		string(DeviceCompliancePolicyStatePlatformTypeAndroidAOSP),
		string(DeviceCompliancePolicyStatePlatformTypeAndroidForWork),
		string(DeviceCompliancePolicyStatePlatformTypeAndroidWorkProfile),
		string(DeviceCompliancePolicyStatePlatformTypeIOS),
		string(DeviceCompliancePolicyStatePlatformTypeMacOS),
		string(DeviceCompliancePolicyStatePlatformTypeWindows10AndLater),
		string(DeviceCompliancePolicyStatePlatformTypeWindows10XProfile),
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
		"all":                DeviceCompliancePolicyStatePlatformTypeAll,
		"android":            DeviceCompliancePolicyStatePlatformTypeAndroid,
		"androidaosp":        DeviceCompliancePolicyStatePlatformTypeAndroidAOSP,
		"androidforwork":     DeviceCompliancePolicyStatePlatformTypeAndroidForWork,
		"androidworkprofile": DeviceCompliancePolicyStatePlatformTypeAndroidWorkProfile,
		"ios":                DeviceCompliancePolicyStatePlatformTypeIOS,
		"macos":              DeviceCompliancePolicyStatePlatformTypeMacOS,
		"windows10andlater":  DeviceCompliancePolicyStatePlatformTypeWindows10AndLater,
		"windows10xprofile":  DeviceCompliancePolicyStatePlatformTypeWindows10XProfile,
		"windows81andlater":  DeviceCompliancePolicyStatePlatformTypeWindows81AndLater,
		"windowsphone81":     DeviceCompliancePolicyStatePlatformTypeWindowsPhone81,
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
	DeviceConfigurationStatePlatformTypeAll                DeviceConfigurationStatePlatformType = "all"
	DeviceConfigurationStatePlatformTypeAndroid            DeviceConfigurationStatePlatformType = "android"
	DeviceConfigurationStatePlatformTypeAndroidAOSP        DeviceConfigurationStatePlatformType = "androidAOSP"
	DeviceConfigurationStatePlatformTypeAndroidForWork     DeviceConfigurationStatePlatformType = "androidForWork"
	DeviceConfigurationStatePlatformTypeAndroidWorkProfile DeviceConfigurationStatePlatformType = "androidWorkProfile"
	DeviceConfigurationStatePlatformTypeIOS                DeviceConfigurationStatePlatformType = "iOS"
	DeviceConfigurationStatePlatformTypeMacOS              DeviceConfigurationStatePlatformType = "macOS"
	DeviceConfigurationStatePlatformTypeWindows10AndLater  DeviceConfigurationStatePlatformType = "windows10AndLater"
	DeviceConfigurationStatePlatformTypeWindows10XProfile  DeviceConfigurationStatePlatformType = "windows10XProfile"
	DeviceConfigurationStatePlatformTypeWindows81AndLater  DeviceConfigurationStatePlatformType = "windows81AndLater"
	DeviceConfigurationStatePlatformTypeWindowsPhone81     DeviceConfigurationStatePlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceConfigurationStatePlatformType() []string {
	return []string{
		string(DeviceConfigurationStatePlatformTypeAll),
		string(DeviceConfigurationStatePlatformTypeAndroid),
		string(DeviceConfigurationStatePlatformTypeAndroidAOSP),
		string(DeviceConfigurationStatePlatformTypeAndroidForWork),
		string(DeviceConfigurationStatePlatformTypeAndroidWorkProfile),
		string(DeviceConfigurationStatePlatformTypeIOS),
		string(DeviceConfigurationStatePlatformTypeMacOS),
		string(DeviceConfigurationStatePlatformTypeWindows10AndLater),
		string(DeviceConfigurationStatePlatformTypeWindows10XProfile),
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
		"all":                DeviceConfigurationStatePlatformTypeAll,
		"android":            DeviceConfigurationStatePlatformTypeAndroid,
		"androidaosp":        DeviceConfigurationStatePlatformTypeAndroidAOSP,
		"androidforwork":     DeviceConfigurationStatePlatformTypeAndroidForWork,
		"androidworkprofile": DeviceConfigurationStatePlatformTypeAndroidWorkProfile,
		"ios":                DeviceConfigurationStatePlatformTypeIOS,
		"macos":              DeviceConfigurationStatePlatformTypeMacOS,
		"windows10andlater":  DeviceConfigurationStatePlatformTypeWindows10AndLater,
		"windows10xprofile":  DeviceConfigurationStatePlatformTypeWindows10XProfile,
		"windows81andlater":  DeviceConfigurationStatePlatformTypeWindows81AndLater,
		"windowsphone81":     DeviceConfigurationStatePlatformTypeWindowsPhone81,
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

type DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultLimit                                          DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultPlatformRestrictions                           DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindowsHelloForBusiness                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDeviceComanagementAuthorityConfiguration              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeEnrollmentNotificationsConfiguration                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeLimit                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypePlatformRestrictions                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeSinglePlatformRestriction                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeUnknown                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindowsHelloForBusiness                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultLimit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultPlatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeEnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeLimit),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypePlatformRestrictions),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeSinglePlatformRestriction),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeUnknown),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeDeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeEnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeLimit,
		"platformrestrictions":                                  DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypePlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeSinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeUnknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationTypeWindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentConfigurationDeviceEnrollmentConfigurationType(input)
	return &out, nil
}

type DeviceHealthAttestationStateFirmwareProtection string

const (
	DeviceHealthAttestationStateFirmwareProtectionDisabled                       DeviceHealthAttestationStateFirmwareProtection = "disabled"
	DeviceHealthAttestationStateFirmwareProtectionFirmwareAttackSurfaceReduction DeviceHealthAttestationStateFirmwareProtection = "firmwareAttackSurfaceReduction"
	DeviceHealthAttestationStateFirmwareProtectionNotApplicable                  DeviceHealthAttestationStateFirmwareProtection = "notApplicable"
	DeviceHealthAttestationStateFirmwareProtectionSystemGuardSecureLaunch        DeviceHealthAttestationStateFirmwareProtection = "systemGuardSecureLaunch"
)

func PossibleValuesForDeviceHealthAttestationStateFirmwareProtection() []string {
	return []string{
		string(DeviceHealthAttestationStateFirmwareProtectionDisabled),
		string(DeviceHealthAttestationStateFirmwareProtectionFirmwareAttackSurfaceReduction),
		string(DeviceHealthAttestationStateFirmwareProtectionNotApplicable),
		string(DeviceHealthAttestationStateFirmwareProtectionSystemGuardSecureLaunch),
	}
}

func (s *DeviceHealthAttestationStateFirmwareProtection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateFirmwareProtection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateFirmwareProtection(input string) (*DeviceHealthAttestationStateFirmwareProtection, error) {
	vals := map[string]DeviceHealthAttestationStateFirmwareProtection{
		"disabled":                       DeviceHealthAttestationStateFirmwareProtectionDisabled,
		"firmwareattacksurfacereduction": DeviceHealthAttestationStateFirmwareProtectionFirmwareAttackSurfaceReduction,
		"notapplicable":                  DeviceHealthAttestationStateFirmwareProtectionNotApplicable,
		"systemguardsecurelaunch":        DeviceHealthAttestationStateFirmwareProtectionSystemGuardSecureLaunch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateFirmwareProtection(input)
	return &out, nil
}

type DeviceHealthAttestationStateMemoryAccessProtection string

const (
	DeviceHealthAttestationStateMemoryAccessProtectionDisabled      DeviceHealthAttestationStateMemoryAccessProtection = "disabled"
	DeviceHealthAttestationStateMemoryAccessProtectionEnabled       DeviceHealthAttestationStateMemoryAccessProtection = "enabled"
	DeviceHealthAttestationStateMemoryAccessProtectionNotApplicable DeviceHealthAttestationStateMemoryAccessProtection = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateMemoryAccessProtection() []string {
	return []string{
		string(DeviceHealthAttestationStateMemoryAccessProtectionDisabled),
		string(DeviceHealthAttestationStateMemoryAccessProtectionEnabled),
		string(DeviceHealthAttestationStateMemoryAccessProtectionNotApplicable),
	}
}

func (s *DeviceHealthAttestationStateMemoryAccessProtection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateMemoryAccessProtection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateMemoryAccessProtection(input string) (*DeviceHealthAttestationStateMemoryAccessProtection, error) {
	vals := map[string]DeviceHealthAttestationStateMemoryAccessProtection{
		"disabled":      DeviceHealthAttestationStateMemoryAccessProtectionDisabled,
		"enabled":       DeviceHealthAttestationStateMemoryAccessProtectionEnabled,
		"notapplicable": DeviceHealthAttestationStateMemoryAccessProtectionNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateMemoryAccessProtection(input)
	return &out, nil
}

type DeviceHealthAttestationStateMemoryIntegrityProtection string

const (
	DeviceHealthAttestationStateMemoryIntegrityProtectionDisabled      DeviceHealthAttestationStateMemoryIntegrityProtection = "disabled"
	DeviceHealthAttestationStateMemoryIntegrityProtectionEnabled       DeviceHealthAttestationStateMemoryIntegrityProtection = "enabled"
	DeviceHealthAttestationStateMemoryIntegrityProtectionNotApplicable DeviceHealthAttestationStateMemoryIntegrityProtection = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateMemoryIntegrityProtection() []string {
	return []string{
		string(DeviceHealthAttestationStateMemoryIntegrityProtectionDisabled),
		string(DeviceHealthAttestationStateMemoryIntegrityProtectionEnabled),
		string(DeviceHealthAttestationStateMemoryIntegrityProtectionNotApplicable),
	}
}

func (s *DeviceHealthAttestationStateMemoryIntegrityProtection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateMemoryIntegrityProtection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateMemoryIntegrityProtection(input string) (*DeviceHealthAttestationStateMemoryIntegrityProtection, error) {
	vals := map[string]DeviceHealthAttestationStateMemoryIntegrityProtection{
		"disabled":      DeviceHealthAttestationStateMemoryIntegrityProtectionDisabled,
		"enabled":       DeviceHealthAttestationStateMemoryIntegrityProtectionEnabled,
		"notapplicable": DeviceHealthAttestationStateMemoryIntegrityProtectionNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateMemoryIntegrityProtection(input)
	return &out, nil
}

type DeviceHealthAttestationStateSecuredCorePC string

const (
	DeviceHealthAttestationStateSecuredCorePCDisabled      DeviceHealthAttestationStateSecuredCorePC = "disabled"
	DeviceHealthAttestationStateSecuredCorePCEnabled       DeviceHealthAttestationStateSecuredCorePC = "enabled"
	DeviceHealthAttestationStateSecuredCorePCNotApplicable DeviceHealthAttestationStateSecuredCorePC = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateSecuredCorePC() []string {
	return []string{
		string(DeviceHealthAttestationStateSecuredCorePCDisabled),
		string(DeviceHealthAttestationStateSecuredCorePCEnabled),
		string(DeviceHealthAttestationStateSecuredCorePCNotApplicable),
	}
}

func (s *DeviceHealthAttestationStateSecuredCorePC) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateSecuredCorePC(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateSecuredCorePC(input string) (*DeviceHealthAttestationStateSecuredCorePC, error) {
	vals := map[string]DeviceHealthAttestationStateSecuredCorePC{
		"disabled":      DeviceHealthAttestationStateSecuredCorePCDisabled,
		"enabled":       DeviceHealthAttestationStateSecuredCorePCEnabled,
		"notapplicable": DeviceHealthAttestationStateSecuredCorePCNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateSecuredCorePC(input)
	return &out, nil
}

type DeviceHealthAttestationStateSystemManagementMode string

const (
	DeviceHealthAttestationStateSystemManagementModeLevel1        DeviceHealthAttestationStateSystemManagementMode = "level1"
	DeviceHealthAttestationStateSystemManagementModeLevel2        DeviceHealthAttestationStateSystemManagementMode = "level2"
	DeviceHealthAttestationStateSystemManagementModeLevel3        DeviceHealthAttestationStateSystemManagementMode = "level3"
	DeviceHealthAttestationStateSystemManagementModeNotApplicable DeviceHealthAttestationStateSystemManagementMode = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateSystemManagementMode() []string {
	return []string{
		string(DeviceHealthAttestationStateSystemManagementModeLevel1),
		string(DeviceHealthAttestationStateSystemManagementModeLevel2),
		string(DeviceHealthAttestationStateSystemManagementModeLevel3),
		string(DeviceHealthAttestationStateSystemManagementModeNotApplicable),
	}
}

func (s *DeviceHealthAttestationStateSystemManagementMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateSystemManagementMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateSystemManagementMode(input string) (*DeviceHealthAttestationStateSystemManagementMode, error) {
	vals := map[string]DeviceHealthAttestationStateSystemManagementMode{
		"level1":        DeviceHealthAttestationStateSystemManagementModeLevel1,
		"level2":        DeviceHealthAttestationStateSystemManagementModeLevel2,
		"level3":        DeviceHealthAttestationStateSystemManagementModeLevel3,
		"notapplicable": DeviceHealthAttestationStateSystemManagementModeNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateSystemManagementMode(input)
	return &out, nil
}

type DeviceHealthAttestationStateVirtualizationBasedSecurity string

const (
	DeviceHealthAttestationStateVirtualizationBasedSecurityDisabled      DeviceHealthAttestationStateVirtualizationBasedSecurity = "disabled"
	DeviceHealthAttestationStateVirtualizationBasedSecurityEnabled       DeviceHealthAttestationStateVirtualizationBasedSecurity = "enabled"
	DeviceHealthAttestationStateVirtualizationBasedSecurityNotApplicable DeviceHealthAttestationStateVirtualizationBasedSecurity = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateVirtualizationBasedSecurity() []string {
	return []string{
		string(DeviceHealthAttestationStateVirtualizationBasedSecurityDisabled),
		string(DeviceHealthAttestationStateVirtualizationBasedSecurityEnabled),
		string(DeviceHealthAttestationStateVirtualizationBasedSecurityNotApplicable),
	}
}

func (s *DeviceHealthAttestationStateVirtualizationBasedSecurity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateVirtualizationBasedSecurity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateVirtualizationBasedSecurity(input string) (*DeviceHealthAttestationStateVirtualizationBasedSecurity, error) {
	vals := map[string]DeviceHealthAttestationStateVirtualizationBasedSecurity{
		"disabled":      DeviceHealthAttestationStateVirtualizationBasedSecurityDisabled,
		"enabled":       DeviceHealthAttestationStateVirtualizationBasedSecurityEnabled,
		"notapplicable": DeviceHealthAttestationStateVirtualizationBasedSecurityNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateVirtualizationBasedSecurity(input)
	return &out, nil
}

type DeviceHealthScriptPolicyStateDetectionState string

const (
	DeviceHealthScriptPolicyStateDetectionStateFail          DeviceHealthScriptPolicyStateDetectionState = "fail"
	DeviceHealthScriptPolicyStateDetectionStateNotApplicable DeviceHealthScriptPolicyStateDetectionState = "notApplicable"
	DeviceHealthScriptPolicyStateDetectionStatePending       DeviceHealthScriptPolicyStateDetectionState = "pending"
	DeviceHealthScriptPolicyStateDetectionStateScriptError   DeviceHealthScriptPolicyStateDetectionState = "scriptError"
	DeviceHealthScriptPolicyStateDetectionStateSuccess       DeviceHealthScriptPolicyStateDetectionState = "success"
	DeviceHealthScriptPolicyStateDetectionStateUnknown       DeviceHealthScriptPolicyStateDetectionState = "unknown"
)

func PossibleValuesForDeviceHealthScriptPolicyStateDetectionState() []string {
	return []string{
		string(DeviceHealthScriptPolicyStateDetectionStateFail),
		string(DeviceHealthScriptPolicyStateDetectionStateNotApplicable),
		string(DeviceHealthScriptPolicyStateDetectionStatePending),
		string(DeviceHealthScriptPolicyStateDetectionStateScriptError),
		string(DeviceHealthScriptPolicyStateDetectionStateSuccess),
		string(DeviceHealthScriptPolicyStateDetectionStateUnknown),
	}
}

func (s *DeviceHealthScriptPolicyStateDetectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptPolicyStateDetectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptPolicyStateDetectionState(input string) (*DeviceHealthScriptPolicyStateDetectionState, error) {
	vals := map[string]DeviceHealthScriptPolicyStateDetectionState{
		"fail":          DeviceHealthScriptPolicyStateDetectionStateFail,
		"notapplicable": DeviceHealthScriptPolicyStateDetectionStateNotApplicable,
		"pending":       DeviceHealthScriptPolicyStateDetectionStatePending,
		"scripterror":   DeviceHealthScriptPolicyStateDetectionStateScriptError,
		"success":       DeviceHealthScriptPolicyStateDetectionStateSuccess,
		"unknown":       DeviceHealthScriptPolicyStateDetectionStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptPolicyStateDetectionState(input)
	return &out, nil
}

type DeviceHealthScriptPolicyStateRemediationState string

const (
	DeviceHealthScriptPolicyStateRemediationStateRemediationFailed DeviceHealthScriptPolicyStateRemediationState = "remediationFailed"
	DeviceHealthScriptPolicyStateRemediationStateScriptError       DeviceHealthScriptPolicyStateRemediationState = "scriptError"
	DeviceHealthScriptPolicyStateRemediationStateSkipped           DeviceHealthScriptPolicyStateRemediationState = "skipped"
	DeviceHealthScriptPolicyStateRemediationStateSuccess           DeviceHealthScriptPolicyStateRemediationState = "success"
	DeviceHealthScriptPolicyStateRemediationStateUnknown           DeviceHealthScriptPolicyStateRemediationState = "unknown"
)

func PossibleValuesForDeviceHealthScriptPolicyStateRemediationState() []string {
	return []string{
		string(DeviceHealthScriptPolicyStateRemediationStateRemediationFailed),
		string(DeviceHealthScriptPolicyStateRemediationStateScriptError),
		string(DeviceHealthScriptPolicyStateRemediationStateSkipped),
		string(DeviceHealthScriptPolicyStateRemediationStateSuccess),
		string(DeviceHealthScriptPolicyStateRemediationStateUnknown),
	}
}

func (s *DeviceHealthScriptPolicyStateRemediationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptPolicyStateRemediationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptPolicyStateRemediationState(input string) (*DeviceHealthScriptPolicyStateRemediationState, error) {
	vals := map[string]DeviceHealthScriptPolicyStateRemediationState{
		"remediationfailed": DeviceHealthScriptPolicyStateRemediationStateRemediationFailed,
		"scripterror":       DeviceHealthScriptPolicyStateRemediationStateScriptError,
		"skipped":           DeviceHealthScriptPolicyStateRemediationStateSkipped,
		"success":           DeviceHealthScriptPolicyStateRemediationStateSuccess,
		"unknown":           DeviceHealthScriptPolicyStateRemediationStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptPolicyStateRemediationState(input)
	return &out, nil
}

type DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus string

const (
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusIncompleteData DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "incompleteData"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusNotSupported   DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "notSupported"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusTrusted        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "trusted"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusUnTrusted      DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "unTrusted"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusUnknown        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "unknown"
)

func PossibleValuesForDeviceIdentityAttestationDetailDeviceIdentityAttestationStatus() []string {
	return []string{
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusIncompleteData),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusNotSupported),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusTrusted),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusUnTrusted),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusUnknown),
	}
}

func (s *DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceIdentityAttestationDetailDeviceIdentityAttestationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceIdentityAttestationDetailDeviceIdentityAttestationStatus(input string) (*DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus, error) {
	vals := map[string]DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus{
		"incompletedata": DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusIncompleteData,
		"notsupported":   DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusNotSupported,
		"trusted":        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusTrusted,
		"untrusted":      DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusUnTrusted,
		"unknown":        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus(input)
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

type DirectoryDefinitionDiscoverabilities string

const (
	DirectoryDefinitionDiscoverabilitiesAttributeDataTypes  DirectoryDefinitionDiscoverabilities = "AttributeDataTypes"
	DirectoryDefinitionDiscoverabilitiesAttributeNames      DirectoryDefinitionDiscoverabilities = "AttributeNames"
	DirectoryDefinitionDiscoverabilitiesAttributeReadOnly   DirectoryDefinitionDiscoverabilities = "AttributeReadOnly"
	DirectoryDefinitionDiscoverabilitiesNone                DirectoryDefinitionDiscoverabilities = "None"
	DirectoryDefinitionDiscoverabilitiesReferenceAttributes DirectoryDefinitionDiscoverabilities = "ReferenceAttributes"
)

func PossibleValuesForDirectoryDefinitionDiscoverabilities() []string {
	return []string{
		string(DirectoryDefinitionDiscoverabilitiesAttributeDataTypes),
		string(DirectoryDefinitionDiscoverabilitiesAttributeNames),
		string(DirectoryDefinitionDiscoverabilitiesAttributeReadOnly),
		string(DirectoryDefinitionDiscoverabilitiesNone),
		string(DirectoryDefinitionDiscoverabilitiesReferenceAttributes),
	}
}

func (s *DirectoryDefinitionDiscoverabilities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDirectoryDefinitionDiscoverabilities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDirectoryDefinitionDiscoverabilities(input string) (*DirectoryDefinitionDiscoverabilities, error) {
	vals := map[string]DirectoryDefinitionDiscoverabilities{
		"attributedatatypes":  DirectoryDefinitionDiscoverabilitiesAttributeDataTypes,
		"attributenames":      DirectoryDefinitionDiscoverabilitiesAttributeNames,
		"attributereadonly":   DirectoryDefinitionDiscoverabilitiesAttributeReadOnly,
		"none":                DirectoryDefinitionDiscoverabilitiesNone,
		"referenceattributes": DirectoryDefinitionDiscoverabilitiesReferenceAttributes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DirectoryDefinitionDiscoverabilities(input)
	return &out, nil
}

type DriveItemSourceApplication string

const (
	DriveItemSourceApplicationOffice     DriveItemSourceApplication = "office"
	DriveItemSourceApplicationOneDrive   DriveItemSourceApplication = "oneDrive"
	DriveItemSourceApplicationPowerPoint DriveItemSourceApplication = "powerPoint"
	DriveItemSourceApplicationSharePoint DriveItemSourceApplication = "sharePoint"
	DriveItemSourceApplicationStream     DriveItemSourceApplication = "stream"
	DriveItemSourceApplicationTeams      DriveItemSourceApplication = "teams"
	DriveItemSourceApplicationYammer     DriveItemSourceApplication = "yammer"
)

func PossibleValuesForDriveItemSourceApplication() []string {
	return []string{
		string(DriveItemSourceApplicationOffice),
		string(DriveItemSourceApplicationOneDrive),
		string(DriveItemSourceApplicationPowerPoint),
		string(DriveItemSourceApplicationSharePoint),
		string(DriveItemSourceApplicationStream),
		string(DriveItemSourceApplicationTeams),
		string(DriveItemSourceApplicationYammer),
	}
}

func (s *DriveItemSourceApplication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDriveItemSourceApplication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDriveItemSourceApplication(input string) (*DriveItemSourceApplication, error) {
	vals := map[string]DriveItemSourceApplication{
		"office":     DriveItemSourceApplicationOffice,
		"onedrive":   DriveItemSourceApplicationOneDrive,
		"powerpoint": DriveItemSourceApplicationPowerPoint,
		"sharepoint": DriveItemSourceApplicationSharePoint,
		"stream":     DriveItemSourceApplicationStream,
		"teams":      DriveItemSourceApplicationTeams,
		"yammer":     DriveItemSourceApplicationYammer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DriveItemSourceApplication(input)
	return &out, nil
}

type EducationalActivityAllowedAudiences string

const (
	EducationalActivityAllowedAudiencesContacts               EducationalActivityAllowedAudiences = "contacts"
	EducationalActivityAllowedAudiencesEveryone               EducationalActivityAllowedAudiences = "everyone"
	EducationalActivityAllowedAudiencesFamily                 EducationalActivityAllowedAudiences = "family"
	EducationalActivityAllowedAudiencesFederatedOrganizations EducationalActivityAllowedAudiences = "federatedOrganizations"
	EducationalActivityAllowedAudiencesGroupMembers           EducationalActivityAllowedAudiences = "groupMembers"
	EducationalActivityAllowedAudiencesMe                     EducationalActivityAllowedAudiences = "me"
	EducationalActivityAllowedAudiencesOrganization           EducationalActivityAllowedAudiences = "organization"
)

func PossibleValuesForEducationalActivityAllowedAudiences() []string {
	return []string{
		string(EducationalActivityAllowedAudiencesContacts),
		string(EducationalActivityAllowedAudiencesEveryone),
		string(EducationalActivityAllowedAudiencesFamily),
		string(EducationalActivityAllowedAudiencesFederatedOrganizations),
		string(EducationalActivityAllowedAudiencesGroupMembers),
		string(EducationalActivityAllowedAudiencesMe),
		string(EducationalActivityAllowedAudiencesOrganization),
	}
}

func (s *EducationalActivityAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationalActivityAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationalActivityAllowedAudiences(input string) (*EducationalActivityAllowedAudiences, error) {
	vals := map[string]EducationalActivityAllowedAudiences{
		"contacts":               EducationalActivityAllowedAudiencesContacts,
		"everyone":               EducationalActivityAllowedAudiencesEveryone,
		"family":                 EducationalActivityAllowedAudiencesFamily,
		"federatedorganizations": EducationalActivityAllowedAudiencesFederatedOrganizations,
		"groupmembers":           EducationalActivityAllowedAudiencesGroupMembers,
		"me":                     EducationalActivityAllowedAudiencesMe,
		"organization":           EducationalActivityAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationalActivityAllowedAudiences(input)
	return &out, nil
}

type EnrollmentConfigurationAssignmentSource string

const (
	EnrollmentConfigurationAssignmentSourceDirect     EnrollmentConfigurationAssignmentSource = "direct"
	EnrollmentConfigurationAssignmentSourcePolicySets EnrollmentConfigurationAssignmentSource = "policySets"
)

func PossibleValuesForEnrollmentConfigurationAssignmentSource() []string {
	return []string{
		string(EnrollmentConfigurationAssignmentSourceDirect),
		string(EnrollmentConfigurationAssignmentSourcePolicySets),
	}
}

func (s *EnrollmentConfigurationAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentConfigurationAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentConfigurationAssignmentSource(input string) (*EnrollmentConfigurationAssignmentSource, error) {
	vals := map[string]EnrollmentConfigurationAssignmentSource{
		"direct":     EnrollmentConfigurationAssignmentSourceDirect,
		"policysets": EnrollmentConfigurationAssignmentSourcePolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentConfigurationAssignmentSource(input)
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

type GroupAccessType string

const (
	GroupAccessTypeNone    GroupAccessType = "none"
	GroupAccessTypePrivate GroupAccessType = "private"
	GroupAccessTypePublic  GroupAccessType = "public"
	GroupAccessTypeSecret  GroupAccessType = "secret"
)

func PossibleValuesForGroupAccessType() []string {
	return []string{
		string(GroupAccessTypeNone),
		string(GroupAccessTypePrivate),
		string(GroupAccessTypePublic),
		string(GroupAccessTypeSecret),
	}
}

func (s *GroupAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupAccessType(input string) (*GroupAccessType, error) {
	vals := map[string]GroupAccessType{
		"none":    GroupAccessTypeNone,
		"private": GroupAccessTypePrivate,
		"public":  GroupAccessTypePublic,
		"secret":  GroupAccessTypeSecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupAccessType(input)
	return &out, nil
}

type HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState string

const (
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateNotConfigured                         HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "notConfigured"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateNotLicensed                           HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "notLicensed"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateRebootRequired                        HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "rebootRequired"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateRunning                               HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "running"
	HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateVirtualizationBasedSecurityNotRunning HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState = "virtualizationBasedSecurityNotRunning"
)

func PossibleValuesForHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState() []string {
	return []string{
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateNotConfigured),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateNotLicensed),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateRebootRequired),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateRunning),
		string(HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateVirtualizationBasedSecurityNotRunning),
	}
}

func (s *HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(input string) (*HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState, error) {
	vals := map[string]HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState{
		"notconfigured":                         HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateNotConfigured,
		"notlicensed":                           HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateNotLicensed,
		"rebootrequired":                        HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateRebootRequired,
		"running":                               HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateRunning,
		"virtualizationbasedsecuritynotrunning": HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardStateVirtualizationBasedSecurityNotRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState(input)
	return &out, nil
}

type HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState string

const (
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateDmaProtectionRequired        HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "dmaProtectionRequired"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateHyperVNotAvailable           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "hyperVNotAvailable"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateHyperVNotSupportedForGuestVM HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "hyperVNotSupportedForGuestVM"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateMeetHardwareRequirements     HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "meetHardwareRequirements"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateSecureBootRequired           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "secureBootRequired"
)

func PossibleValuesForHardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState() []string {
	return []string{
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateDmaProtectionRequired),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateHyperVNotAvailable),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateHyperVNotSupportedForGuestVM),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateMeetHardwareRequirements),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateSecureBootRequired),
	}
}

func (s *HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(input string) (*HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState, error) {
	vals := map[string]HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState{
		"dmaprotectionrequired":        HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateDmaProtectionRequired,
		"hypervnotavailable":           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateHyperVNotAvailable,
		"hypervnotsupportedforguestvm": HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateHyperVNotSupportedForGuestVM,
		"meethardwarerequirements":     HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateMeetHardwareRequirements,
		"securebootrequired":           HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementStateSecureBootRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(input)
	return &out, nil
}

type HardwareInformationDeviceGuardVirtualizationBasedSecurityState string

const (
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateDoesNotMeetHardwareRequirements HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "doesNotMeetHardwareRequirements"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateNotConfigured                   HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "notConfigured"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateNotLicensed                     HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "notLicensed"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateOther                           HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "other"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRebootRequired                  HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "rebootRequired"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRequire64BitArchitecture        HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "require64BitArchitecture"
	HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRunning                         HardwareInformationDeviceGuardVirtualizationBasedSecurityState = "running"
)

func PossibleValuesForHardwareInformationDeviceGuardVirtualizationBasedSecurityState() []string {
	return []string{
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateDoesNotMeetHardwareRequirements),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateNotConfigured),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateNotLicensed),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateOther),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRebootRequired),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRequire64BitArchitecture),
		string(HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRunning),
	}
}

func (s *HardwareInformationDeviceGuardVirtualizationBasedSecurityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceGuardVirtualizationBasedSecurityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceGuardVirtualizationBasedSecurityState(input string) (*HardwareInformationDeviceGuardVirtualizationBasedSecurityState, error) {
	vals := map[string]HardwareInformationDeviceGuardVirtualizationBasedSecurityState{
		"doesnotmeethardwarerequirements": HardwareInformationDeviceGuardVirtualizationBasedSecurityStateDoesNotMeetHardwareRequirements,
		"notconfigured":                   HardwareInformationDeviceGuardVirtualizationBasedSecurityStateNotConfigured,
		"notlicensed":                     HardwareInformationDeviceGuardVirtualizationBasedSecurityStateNotLicensed,
		"other":                           HardwareInformationDeviceGuardVirtualizationBasedSecurityStateOther,
		"rebootrequired":                  HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRebootRequired,
		"require64bitarchitecture":        HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRequire64BitArchitecture,
		"running":                         HardwareInformationDeviceGuardVirtualizationBasedSecurityStateRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceGuardVirtualizationBasedSecurityState(input)
	return &out, nil
}

type HardwareInformationDeviceLicensingStatus string

const (
	HardwareInformationDeviceLicensingStatusAcquiringDeviceLicense                HardwareInformationDeviceLicensingStatus = "acquiringDeviceLicense"
	HardwareInformationDeviceLicensingStatusDeviceIdentityVerificationFailed      HardwareInformationDeviceLicensingStatus = "deviceIdentityVerificationFailed"
	HardwareInformationDeviceLicensingStatusDeviceIsNotAzureActiveDirectoryJoined HardwareInformationDeviceLicensingStatus = "deviceIsNotAzureActiveDirectoryJoined"
	HardwareInformationDeviceLicensingStatusDeviceLicenseRefreshFailed            HardwareInformationDeviceLicensingStatus = "deviceLicenseRefreshFailed"
	HardwareInformationDeviceLicensingStatusDeviceLicenseRefreshSucceed           HardwareInformationDeviceLicensingStatus = "deviceLicenseRefreshSucceed"
	HardwareInformationDeviceLicensingStatusDeviceLicenseRemoveFailed             HardwareInformationDeviceLicensingStatus = "deviceLicenseRemoveFailed"
	HardwareInformationDeviceLicensingStatusDeviceLicenseRemoveSucceed            HardwareInformationDeviceLicensingStatus = "deviceLicenseRemoveSucceed"
	HardwareInformationDeviceLicensingStatusLicenseRefreshPending                 HardwareInformationDeviceLicensingStatus = "licenseRefreshPending"
	HardwareInformationDeviceLicensingStatusLicenseRefreshStarted                 HardwareInformationDeviceLicensingStatus = "licenseRefreshStarted"
	HardwareInformationDeviceLicensingStatusMicrosoftAccountVerificationFailed    HardwareInformationDeviceLicensingStatus = "microsoftAccountVerificationFailed"
	HardwareInformationDeviceLicensingStatusRefreshingDeviceLicense               HardwareInformationDeviceLicensingStatus = "refreshingDeviceLicense"
	HardwareInformationDeviceLicensingStatusRemovingDeviceLicense                 HardwareInformationDeviceLicensingStatus = "removingDeviceLicense"
	HardwareInformationDeviceLicensingStatusUnknown                               HardwareInformationDeviceLicensingStatus = "unknown"
	HardwareInformationDeviceLicensingStatusVerifyingMicrosoftAccountIdentity     HardwareInformationDeviceLicensingStatus = "verifyingMicrosoftAccountIdentity"
	HardwareInformationDeviceLicensingStatusVerifyingMicrosoftDeviceIdentity      HardwareInformationDeviceLicensingStatus = "verifyingMicrosoftDeviceIdentity"
)

func PossibleValuesForHardwareInformationDeviceLicensingStatus() []string {
	return []string{
		string(HardwareInformationDeviceLicensingStatusAcquiringDeviceLicense),
		string(HardwareInformationDeviceLicensingStatusDeviceIdentityVerificationFailed),
		string(HardwareInformationDeviceLicensingStatusDeviceIsNotAzureActiveDirectoryJoined),
		string(HardwareInformationDeviceLicensingStatusDeviceLicenseRefreshFailed),
		string(HardwareInformationDeviceLicensingStatusDeviceLicenseRefreshSucceed),
		string(HardwareInformationDeviceLicensingStatusDeviceLicenseRemoveFailed),
		string(HardwareInformationDeviceLicensingStatusDeviceLicenseRemoveSucceed),
		string(HardwareInformationDeviceLicensingStatusLicenseRefreshPending),
		string(HardwareInformationDeviceLicensingStatusLicenseRefreshStarted),
		string(HardwareInformationDeviceLicensingStatusMicrosoftAccountVerificationFailed),
		string(HardwareInformationDeviceLicensingStatusRefreshingDeviceLicense),
		string(HardwareInformationDeviceLicensingStatusRemovingDeviceLicense),
		string(HardwareInformationDeviceLicensingStatusUnknown),
		string(HardwareInformationDeviceLicensingStatusVerifyingMicrosoftAccountIdentity),
		string(HardwareInformationDeviceLicensingStatusVerifyingMicrosoftDeviceIdentity),
	}
}

func (s *HardwareInformationDeviceLicensingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareInformationDeviceLicensingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareInformationDeviceLicensingStatus(input string) (*HardwareInformationDeviceLicensingStatus, error) {
	vals := map[string]HardwareInformationDeviceLicensingStatus{
		"acquiringdevicelicense":                HardwareInformationDeviceLicensingStatusAcquiringDeviceLicense,
		"deviceidentityverificationfailed":      HardwareInformationDeviceLicensingStatusDeviceIdentityVerificationFailed,
		"deviceisnotazureactivedirectoryjoined": HardwareInformationDeviceLicensingStatusDeviceIsNotAzureActiveDirectoryJoined,
		"devicelicenserefreshfailed":            HardwareInformationDeviceLicensingStatusDeviceLicenseRefreshFailed,
		"devicelicenserefreshsucceed":           HardwareInformationDeviceLicensingStatusDeviceLicenseRefreshSucceed,
		"devicelicenseremovefailed":             HardwareInformationDeviceLicensingStatusDeviceLicenseRemoveFailed,
		"devicelicenseremovesucceed":            HardwareInformationDeviceLicensingStatusDeviceLicenseRemoveSucceed,
		"licenserefreshpending":                 HardwareInformationDeviceLicensingStatusLicenseRefreshPending,
		"licenserefreshstarted":                 HardwareInformationDeviceLicensingStatusLicenseRefreshStarted,
		"microsoftaccountverificationfailed":    HardwareInformationDeviceLicensingStatusMicrosoftAccountVerificationFailed,
		"refreshingdevicelicense":               HardwareInformationDeviceLicensingStatusRefreshingDeviceLicense,
		"removingdevicelicense":                 HardwareInformationDeviceLicensingStatusRemovingDeviceLicense,
		"unknown":                               HardwareInformationDeviceLicensingStatusUnknown,
		"verifyingmicrosoftaccountidentity":     HardwareInformationDeviceLicensingStatusVerifyingMicrosoftAccountIdentity,
		"verifyingmicrosoftdeviceidentity":      HardwareInformationDeviceLicensingStatusVerifyingMicrosoftDeviceIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareInformationDeviceLicensingStatus(input)
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

type ItemAddressAllowedAudiences string

const (
	ItemAddressAllowedAudiencesContacts               ItemAddressAllowedAudiences = "contacts"
	ItemAddressAllowedAudiencesEveryone               ItemAddressAllowedAudiences = "everyone"
	ItemAddressAllowedAudiencesFamily                 ItemAddressAllowedAudiences = "family"
	ItemAddressAllowedAudiencesFederatedOrganizations ItemAddressAllowedAudiences = "federatedOrganizations"
	ItemAddressAllowedAudiencesGroupMembers           ItemAddressAllowedAudiences = "groupMembers"
	ItemAddressAllowedAudiencesMe                     ItemAddressAllowedAudiences = "me"
	ItemAddressAllowedAudiencesOrganization           ItemAddressAllowedAudiences = "organization"
)

func PossibleValuesForItemAddressAllowedAudiences() []string {
	return []string{
		string(ItemAddressAllowedAudiencesContacts),
		string(ItemAddressAllowedAudiencesEveryone),
		string(ItemAddressAllowedAudiencesFamily),
		string(ItemAddressAllowedAudiencesFederatedOrganizations),
		string(ItemAddressAllowedAudiencesGroupMembers),
		string(ItemAddressAllowedAudiencesMe),
		string(ItemAddressAllowedAudiencesOrganization),
	}
}

func (s *ItemAddressAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemAddressAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemAddressAllowedAudiences(input string) (*ItemAddressAllowedAudiences, error) {
	vals := map[string]ItemAddressAllowedAudiences{
		"contacts":               ItemAddressAllowedAudiencesContacts,
		"everyone":               ItemAddressAllowedAudiencesEveryone,
		"family":                 ItemAddressAllowedAudiencesFamily,
		"federatedorganizations": ItemAddressAllowedAudiencesFederatedOrganizations,
		"groupmembers":           ItemAddressAllowedAudiencesGroupMembers,
		"me":                     ItemAddressAllowedAudiencesMe,
		"organization":           ItemAddressAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemAddressAllowedAudiences(input)
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

type ItemEmailAllowedAudiences string

const (
	ItemEmailAllowedAudiencesContacts               ItemEmailAllowedAudiences = "contacts"
	ItemEmailAllowedAudiencesEveryone               ItemEmailAllowedAudiences = "everyone"
	ItemEmailAllowedAudiencesFamily                 ItemEmailAllowedAudiences = "family"
	ItemEmailAllowedAudiencesFederatedOrganizations ItemEmailAllowedAudiences = "federatedOrganizations"
	ItemEmailAllowedAudiencesGroupMembers           ItemEmailAllowedAudiences = "groupMembers"
	ItemEmailAllowedAudiencesMe                     ItemEmailAllowedAudiences = "me"
	ItemEmailAllowedAudiencesOrganization           ItemEmailAllowedAudiences = "organization"
)

func PossibleValuesForItemEmailAllowedAudiences() []string {
	return []string{
		string(ItemEmailAllowedAudiencesContacts),
		string(ItemEmailAllowedAudiencesEveryone),
		string(ItemEmailAllowedAudiencesFamily),
		string(ItemEmailAllowedAudiencesFederatedOrganizations),
		string(ItemEmailAllowedAudiencesGroupMembers),
		string(ItemEmailAllowedAudiencesMe),
		string(ItemEmailAllowedAudiencesOrganization),
	}
}

func (s *ItemEmailAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemEmailAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemEmailAllowedAudiences(input string) (*ItemEmailAllowedAudiences, error) {
	vals := map[string]ItemEmailAllowedAudiences{
		"contacts":               ItemEmailAllowedAudiencesContacts,
		"everyone":               ItemEmailAllowedAudiencesEveryone,
		"family":                 ItemEmailAllowedAudiencesFamily,
		"federatedorganizations": ItemEmailAllowedAudiencesFederatedOrganizations,
		"groupmembers":           ItemEmailAllowedAudiencesGroupMembers,
		"me":                     ItemEmailAllowedAudiencesMe,
		"organization":           ItemEmailAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemEmailAllowedAudiences(input)
	return &out, nil
}

type ItemEmailType string

const (
	ItemEmailTypeMain     ItemEmailType = "main"
	ItemEmailTypeOther    ItemEmailType = "other"
	ItemEmailTypePersonal ItemEmailType = "personal"
	ItemEmailTypeUnknown  ItemEmailType = "unknown"
	ItemEmailTypeWork     ItemEmailType = "work"
)

func PossibleValuesForItemEmailType() []string {
	return []string{
		string(ItemEmailTypeMain),
		string(ItemEmailTypeOther),
		string(ItemEmailTypePersonal),
		string(ItemEmailTypeUnknown),
		string(ItemEmailTypeWork),
	}
}

func (s *ItemEmailType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemEmailType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemEmailType(input string) (*ItemEmailType, error) {
	vals := map[string]ItemEmailType{
		"main":     ItemEmailTypeMain,
		"other":    ItemEmailTypeOther,
		"personal": ItemEmailTypePersonal,
		"unknown":  ItemEmailTypeUnknown,
		"work":     ItemEmailTypeWork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemEmailType(input)
	return &out, nil
}

type ItemPatentAllowedAudiences string

const (
	ItemPatentAllowedAudiencesContacts               ItemPatentAllowedAudiences = "contacts"
	ItemPatentAllowedAudiencesEveryone               ItemPatentAllowedAudiences = "everyone"
	ItemPatentAllowedAudiencesFamily                 ItemPatentAllowedAudiences = "family"
	ItemPatentAllowedAudiencesFederatedOrganizations ItemPatentAllowedAudiences = "federatedOrganizations"
	ItemPatentAllowedAudiencesGroupMembers           ItemPatentAllowedAudiences = "groupMembers"
	ItemPatentAllowedAudiencesMe                     ItemPatentAllowedAudiences = "me"
	ItemPatentAllowedAudiencesOrganization           ItemPatentAllowedAudiences = "organization"
)

func PossibleValuesForItemPatentAllowedAudiences() []string {
	return []string{
		string(ItemPatentAllowedAudiencesContacts),
		string(ItemPatentAllowedAudiencesEveryone),
		string(ItemPatentAllowedAudiencesFamily),
		string(ItemPatentAllowedAudiencesFederatedOrganizations),
		string(ItemPatentAllowedAudiencesGroupMembers),
		string(ItemPatentAllowedAudiencesMe),
		string(ItemPatentAllowedAudiencesOrganization),
	}
}

func (s *ItemPatentAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPatentAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPatentAllowedAudiences(input string) (*ItemPatentAllowedAudiences, error) {
	vals := map[string]ItemPatentAllowedAudiences{
		"contacts":               ItemPatentAllowedAudiencesContacts,
		"everyone":               ItemPatentAllowedAudiencesEveryone,
		"family":                 ItemPatentAllowedAudiencesFamily,
		"federatedorganizations": ItemPatentAllowedAudiencesFederatedOrganizations,
		"groupmembers":           ItemPatentAllowedAudiencesGroupMembers,
		"me":                     ItemPatentAllowedAudiencesMe,
		"organization":           ItemPatentAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPatentAllowedAudiences(input)
	return &out, nil
}

type ItemPhoneAllowedAudiences string

const (
	ItemPhoneAllowedAudiencesContacts               ItemPhoneAllowedAudiences = "contacts"
	ItemPhoneAllowedAudiencesEveryone               ItemPhoneAllowedAudiences = "everyone"
	ItemPhoneAllowedAudiencesFamily                 ItemPhoneAllowedAudiences = "family"
	ItemPhoneAllowedAudiencesFederatedOrganizations ItemPhoneAllowedAudiences = "federatedOrganizations"
	ItemPhoneAllowedAudiencesGroupMembers           ItemPhoneAllowedAudiences = "groupMembers"
	ItemPhoneAllowedAudiencesMe                     ItemPhoneAllowedAudiences = "me"
	ItemPhoneAllowedAudiencesOrganization           ItemPhoneAllowedAudiences = "organization"
)

func PossibleValuesForItemPhoneAllowedAudiences() []string {
	return []string{
		string(ItemPhoneAllowedAudiencesContacts),
		string(ItemPhoneAllowedAudiencesEveryone),
		string(ItemPhoneAllowedAudiencesFamily),
		string(ItemPhoneAllowedAudiencesFederatedOrganizations),
		string(ItemPhoneAllowedAudiencesGroupMembers),
		string(ItemPhoneAllowedAudiencesMe),
		string(ItemPhoneAllowedAudiencesOrganization),
	}
}

func (s *ItemPhoneAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPhoneAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPhoneAllowedAudiences(input string) (*ItemPhoneAllowedAudiences, error) {
	vals := map[string]ItemPhoneAllowedAudiences{
		"contacts":               ItemPhoneAllowedAudiencesContacts,
		"everyone":               ItemPhoneAllowedAudiencesEveryone,
		"family":                 ItemPhoneAllowedAudiencesFamily,
		"federatedorganizations": ItemPhoneAllowedAudiencesFederatedOrganizations,
		"groupmembers":           ItemPhoneAllowedAudiencesGroupMembers,
		"me":                     ItemPhoneAllowedAudiencesMe,
		"organization":           ItemPhoneAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPhoneAllowedAudiences(input)
	return &out, nil
}

type ItemPhoneType string

const (
	ItemPhoneTypeAssistant   ItemPhoneType = "assistant"
	ItemPhoneTypeBusiness    ItemPhoneType = "business"
	ItemPhoneTypeBusinessFax ItemPhoneType = "businessFax"
	ItemPhoneTypeHome        ItemPhoneType = "home"
	ItemPhoneTypeHomeFax     ItemPhoneType = "homeFax"
	ItemPhoneTypeMobile      ItemPhoneType = "mobile"
	ItemPhoneTypeOther       ItemPhoneType = "other"
	ItemPhoneTypeOtherFax    ItemPhoneType = "otherFax"
	ItemPhoneTypePager       ItemPhoneType = "pager"
	ItemPhoneTypeRadio       ItemPhoneType = "radio"
)

func PossibleValuesForItemPhoneType() []string {
	return []string{
		string(ItemPhoneTypeAssistant),
		string(ItemPhoneTypeBusiness),
		string(ItemPhoneTypeBusinessFax),
		string(ItemPhoneTypeHome),
		string(ItemPhoneTypeHomeFax),
		string(ItemPhoneTypeMobile),
		string(ItemPhoneTypeOther),
		string(ItemPhoneTypeOtherFax),
		string(ItemPhoneTypePager),
		string(ItemPhoneTypeRadio),
	}
}

func (s *ItemPhoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPhoneType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPhoneType(input string) (*ItemPhoneType, error) {
	vals := map[string]ItemPhoneType{
		"assistant":   ItemPhoneTypeAssistant,
		"business":    ItemPhoneTypeBusiness,
		"businessfax": ItemPhoneTypeBusinessFax,
		"home":        ItemPhoneTypeHome,
		"homefax":     ItemPhoneTypeHomeFax,
		"mobile":      ItemPhoneTypeMobile,
		"other":       ItemPhoneTypeOther,
		"otherfax":    ItemPhoneTypeOtherFax,
		"pager":       ItemPhoneTypePager,
		"radio":       ItemPhoneTypeRadio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPhoneType(input)
	return &out, nil
}

type ItemPublicationAllowedAudiences string

const (
	ItemPublicationAllowedAudiencesContacts               ItemPublicationAllowedAudiences = "contacts"
	ItemPublicationAllowedAudiencesEveryone               ItemPublicationAllowedAudiences = "everyone"
	ItemPublicationAllowedAudiencesFamily                 ItemPublicationAllowedAudiences = "family"
	ItemPublicationAllowedAudiencesFederatedOrganizations ItemPublicationAllowedAudiences = "federatedOrganizations"
	ItemPublicationAllowedAudiencesGroupMembers           ItemPublicationAllowedAudiences = "groupMembers"
	ItemPublicationAllowedAudiencesMe                     ItemPublicationAllowedAudiences = "me"
	ItemPublicationAllowedAudiencesOrganization           ItemPublicationAllowedAudiences = "organization"
)

func PossibleValuesForItemPublicationAllowedAudiences() []string {
	return []string{
		string(ItemPublicationAllowedAudiencesContacts),
		string(ItemPublicationAllowedAudiencesEveryone),
		string(ItemPublicationAllowedAudiencesFamily),
		string(ItemPublicationAllowedAudiencesFederatedOrganizations),
		string(ItemPublicationAllowedAudiencesGroupMembers),
		string(ItemPublicationAllowedAudiencesMe),
		string(ItemPublicationAllowedAudiencesOrganization),
	}
}

func (s *ItemPublicationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPublicationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPublicationAllowedAudiences(input string) (*ItemPublicationAllowedAudiences, error) {
	vals := map[string]ItemPublicationAllowedAudiences{
		"contacts":               ItemPublicationAllowedAudiencesContacts,
		"everyone":               ItemPublicationAllowedAudiencesEveryone,
		"family":                 ItemPublicationAllowedAudiencesFamily,
		"federatedorganizations": ItemPublicationAllowedAudiencesFederatedOrganizations,
		"groupmembers":           ItemPublicationAllowedAudiencesGroupMembers,
		"me":                     ItemPublicationAllowedAudiencesMe,
		"organization":           ItemPublicationAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPublicationAllowedAudiences(input)
	return &out, nil
}

type KeyCredentialConfigurationRestrictionType string

const (
	KeyCredentialConfigurationRestrictionTypeAsymmetricKeyLifetime       KeyCredentialConfigurationRestrictionType = "asymmetricKeyLifetime"
	KeyCredentialConfigurationRestrictionTypeTrustedCertificateAuthority KeyCredentialConfigurationRestrictionType = "trustedCertificateAuthority"
)

func PossibleValuesForKeyCredentialConfigurationRestrictionType() []string {
	return []string{
		string(KeyCredentialConfigurationRestrictionTypeAsymmetricKeyLifetime),
		string(KeyCredentialConfigurationRestrictionTypeTrustedCertificateAuthority),
	}
}

func (s *KeyCredentialConfigurationRestrictionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKeyCredentialConfigurationRestrictionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKeyCredentialConfigurationRestrictionType(input string) (*KeyCredentialConfigurationRestrictionType, error) {
	vals := map[string]KeyCredentialConfigurationRestrictionType{
		"asymmetrickeylifetime":       KeyCredentialConfigurationRestrictionTypeAsymmetricKeyLifetime,
		"trustedcertificateauthority": KeyCredentialConfigurationRestrictionTypeTrustedCertificateAuthority,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyCredentialConfigurationRestrictionType(input)
	return &out, nil
}

type LanguageProficiencyAllowedAudiences string

const (
	LanguageProficiencyAllowedAudiencesContacts               LanguageProficiencyAllowedAudiences = "contacts"
	LanguageProficiencyAllowedAudiencesEveryone               LanguageProficiencyAllowedAudiences = "everyone"
	LanguageProficiencyAllowedAudiencesFamily                 LanguageProficiencyAllowedAudiences = "family"
	LanguageProficiencyAllowedAudiencesFederatedOrganizations LanguageProficiencyAllowedAudiences = "federatedOrganizations"
	LanguageProficiencyAllowedAudiencesGroupMembers           LanguageProficiencyAllowedAudiences = "groupMembers"
	LanguageProficiencyAllowedAudiencesMe                     LanguageProficiencyAllowedAudiences = "me"
	LanguageProficiencyAllowedAudiencesOrganization           LanguageProficiencyAllowedAudiences = "organization"
)

func PossibleValuesForLanguageProficiencyAllowedAudiences() []string {
	return []string{
		string(LanguageProficiencyAllowedAudiencesContacts),
		string(LanguageProficiencyAllowedAudiencesEveryone),
		string(LanguageProficiencyAllowedAudiencesFamily),
		string(LanguageProficiencyAllowedAudiencesFederatedOrganizations),
		string(LanguageProficiencyAllowedAudiencesGroupMembers),
		string(LanguageProficiencyAllowedAudiencesMe),
		string(LanguageProficiencyAllowedAudiencesOrganization),
	}
}

func (s *LanguageProficiencyAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyAllowedAudiences(input string) (*LanguageProficiencyAllowedAudiences, error) {
	vals := map[string]LanguageProficiencyAllowedAudiences{
		"contacts":               LanguageProficiencyAllowedAudiencesContacts,
		"everyone":               LanguageProficiencyAllowedAudiencesEveryone,
		"family":                 LanguageProficiencyAllowedAudiencesFamily,
		"federatedorganizations": LanguageProficiencyAllowedAudiencesFederatedOrganizations,
		"groupmembers":           LanguageProficiencyAllowedAudiencesGroupMembers,
		"me":                     LanguageProficiencyAllowedAudiencesMe,
		"organization":           LanguageProficiencyAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyAllowedAudiences(input)
	return &out, nil
}

type LanguageProficiencyProficiency string

const (
	LanguageProficiencyProficiencyConversational      LanguageProficiencyProficiency = "conversational"
	LanguageProficiencyProficiencyElementary          LanguageProficiencyProficiency = "elementary"
	LanguageProficiencyProficiencyFullProfessional    LanguageProficiencyProficiency = "fullProfessional"
	LanguageProficiencyProficiencyLimitedWorking      LanguageProficiencyProficiency = "limitedWorking"
	LanguageProficiencyProficiencyNativeOrBilingual   LanguageProficiencyProficiency = "nativeOrBilingual"
	LanguageProficiencyProficiencyProfessionalWorking LanguageProficiencyProficiency = "professionalWorking"
)

func PossibleValuesForLanguageProficiencyProficiency() []string {
	return []string{
		string(LanguageProficiencyProficiencyConversational),
		string(LanguageProficiencyProficiencyElementary),
		string(LanguageProficiencyProficiencyFullProfessional),
		string(LanguageProficiencyProficiencyLimitedWorking),
		string(LanguageProficiencyProficiencyNativeOrBilingual),
		string(LanguageProficiencyProficiencyProfessionalWorking),
	}
}

func (s *LanguageProficiencyProficiency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyProficiency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyProficiency(input string) (*LanguageProficiencyProficiency, error) {
	vals := map[string]LanguageProficiencyProficiency{
		"conversational":      LanguageProficiencyProficiencyConversational,
		"elementary":          LanguageProficiencyProficiencyElementary,
		"fullprofessional":    LanguageProficiencyProficiencyFullProfessional,
		"limitedworking":      LanguageProficiencyProficiencyLimitedWorking,
		"nativeorbilingual":   LanguageProficiencyProficiencyNativeOrBilingual,
		"professionalworking": LanguageProficiencyProficiencyProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyProficiency(input)
	return &out, nil
}

type LanguageProficiencyReading string

const (
	LanguageProficiencyReadingConversational      LanguageProficiencyReading = "conversational"
	LanguageProficiencyReadingElementary          LanguageProficiencyReading = "elementary"
	LanguageProficiencyReadingFullProfessional    LanguageProficiencyReading = "fullProfessional"
	LanguageProficiencyReadingLimitedWorking      LanguageProficiencyReading = "limitedWorking"
	LanguageProficiencyReadingNativeOrBilingual   LanguageProficiencyReading = "nativeOrBilingual"
	LanguageProficiencyReadingProfessionalWorking LanguageProficiencyReading = "professionalWorking"
)

func PossibleValuesForLanguageProficiencyReading() []string {
	return []string{
		string(LanguageProficiencyReadingConversational),
		string(LanguageProficiencyReadingElementary),
		string(LanguageProficiencyReadingFullProfessional),
		string(LanguageProficiencyReadingLimitedWorking),
		string(LanguageProficiencyReadingNativeOrBilingual),
		string(LanguageProficiencyReadingProfessionalWorking),
	}
}

func (s *LanguageProficiencyReading) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyReading(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyReading(input string) (*LanguageProficiencyReading, error) {
	vals := map[string]LanguageProficiencyReading{
		"conversational":      LanguageProficiencyReadingConversational,
		"elementary":          LanguageProficiencyReadingElementary,
		"fullprofessional":    LanguageProficiencyReadingFullProfessional,
		"limitedworking":      LanguageProficiencyReadingLimitedWorking,
		"nativeorbilingual":   LanguageProficiencyReadingNativeOrBilingual,
		"professionalworking": LanguageProficiencyReadingProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyReading(input)
	return &out, nil
}

type LanguageProficiencySpoken string

const (
	LanguageProficiencySpokenConversational      LanguageProficiencySpoken = "conversational"
	LanguageProficiencySpokenElementary          LanguageProficiencySpoken = "elementary"
	LanguageProficiencySpokenFullProfessional    LanguageProficiencySpoken = "fullProfessional"
	LanguageProficiencySpokenLimitedWorking      LanguageProficiencySpoken = "limitedWorking"
	LanguageProficiencySpokenNativeOrBilingual   LanguageProficiencySpoken = "nativeOrBilingual"
	LanguageProficiencySpokenProfessionalWorking LanguageProficiencySpoken = "professionalWorking"
)

func PossibleValuesForLanguageProficiencySpoken() []string {
	return []string{
		string(LanguageProficiencySpokenConversational),
		string(LanguageProficiencySpokenElementary),
		string(LanguageProficiencySpokenFullProfessional),
		string(LanguageProficiencySpokenLimitedWorking),
		string(LanguageProficiencySpokenNativeOrBilingual),
		string(LanguageProficiencySpokenProfessionalWorking),
	}
}

func (s *LanguageProficiencySpoken) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencySpoken(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencySpoken(input string) (*LanguageProficiencySpoken, error) {
	vals := map[string]LanguageProficiencySpoken{
		"conversational":      LanguageProficiencySpokenConversational,
		"elementary":          LanguageProficiencySpokenElementary,
		"fullprofessional":    LanguageProficiencySpokenFullProfessional,
		"limitedworking":      LanguageProficiencySpokenLimitedWorking,
		"nativeorbilingual":   LanguageProficiencySpokenNativeOrBilingual,
		"professionalworking": LanguageProficiencySpokenProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencySpoken(input)
	return &out, nil
}

type LanguageProficiencyWritten string

const (
	LanguageProficiencyWrittenConversational      LanguageProficiencyWritten = "conversational"
	LanguageProficiencyWrittenElementary          LanguageProficiencyWritten = "elementary"
	LanguageProficiencyWrittenFullProfessional    LanguageProficiencyWritten = "fullProfessional"
	LanguageProficiencyWrittenLimitedWorking      LanguageProficiencyWritten = "limitedWorking"
	LanguageProficiencyWrittenNativeOrBilingual   LanguageProficiencyWritten = "nativeOrBilingual"
	LanguageProficiencyWrittenProfessionalWorking LanguageProficiencyWritten = "professionalWorking"
)

func PossibleValuesForLanguageProficiencyWritten() []string {
	return []string{
		string(LanguageProficiencyWrittenConversational),
		string(LanguageProficiencyWrittenElementary),
		string(LanguageProficiencyWrittenFullProfessional),
		string(LanguageProficiencyWrittenLimitedWorking),
		string(LanguageProficiencyWrittenNativeOrBilingual),
		string(LanguageProficiencyWrittenProfessionalWorking),
	}
}

func (s *LanguageProficiencyWritten) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyWritten(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyWritten(input string) (*LanguageProficiencyWritten, error) {
	vals := map[string]LanguageProficiencyWritten{
		"conversational":      LanguageProficiencyWrittenConversational,
		"elementary":          LanguageProficiencyWrittenElementary,
		"fullprofessional":    LanguageProficiencyWrittenFullProfessional,
		"limitedworking":      LanguageProficiencyWrittenLimitedWorking,
		"nativeorbilingual":   LanguageProficiencyWrittenNativeOrBilingual,
		"professionalworking": LanguageProficiencyWrittenProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyWritten(input)
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
	MailboxSettingsUserPurposeUnknown   MailboxSettingsUserPurpose = "unknown"
	MailboxSettingsUserPurposeUser      MailboxSettingsUserPurpose = "user"
)

func PossibleValuesForMailboxSettingsUserPurpose() []string {
	return []string{
		string(MailboxSettingsUserPurposeEquipment),
		string(MailboxSettingsUserPurposeLinked),
		string(MailboxSettingsUserPurposeOthers),
		string(MailboxSettingsUserPurposeRoom),
		string(MailboxSettingsUserPurposeShared),
		string(MailboxSettingsUserPurposeUnknown),
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
		"unknown":   MailboxSettingsUserPurposeUnknown,
		"user":      MailboxSettingsUserPurposeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsUserPurpose(input)
	return &out, nil
}

type MailboxSettingsUserPurposeV2 string

const (
	MailboxSettingsUserPurposeV2Equipment MailboxSettingsUserPurposeV2 = "equipment"
	MailboxSettingsUserPurposeV2Linked    MailboxSettingsUserPurposeV2 = "linked"
	MailboxSettingsUserPurposeV2Others    MailboxSettingsUserPurposeV2 = "others"
	MailboxSettingsUserPurposeV2Room      MailboxSettingsUserPurposeV2 = "room"
	MailboxSettingsUserPurposeV2Shared    MailboxSettingsUserPurposeV2 = "shared"
	MailboxSettingsUserPurposeV2Unknown   MailboxSettingsUserPurposeV2 = "unknown"
	MailboxSettingsUserPurposeV2User      MailboxSettingsUserPurposeV2 = "user"
)

func PossibleValuesForMailboxSettingsUserPurposeV2() []string {
	return []string{
		string(MailboxSettingsUserPurposeV2Equipment),
		string(MailboxSettingsUserPurposeV2Linked),
		string(MailboxSettingsUserPurposeV2Others),
		string(MailboxSettingsUserPurposeV2Room),
		string(MailboxSettingsUserPurposeV2Shared),
		string(MailboxSettingsUserPurposeV2Unknown),
		string(MailboxSettingsUserPurposeV2User),
	}
}

func (s *MailboxSettingsUserPurposeV2) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxSettingsUserPurposeV2(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxSettingsUserPurposeV2(input string) (*MailboxSettingsUserPurposeV2, error) {
	vals := map[string]MailboxSettingsUserPurposeV2{
		"equipment": MailboxSettingsUserPurposeV2Equipment,
		"linked":    MailboxSettingsUserPurposeV2Linked,
		"others":    MailboxSettingsUserPurposeV2Others,
		"room":      MailboxSettingsUserPurposeV2Room,
		"shared":    MailboxSettingsUserPurposeV2Shared,
		"unknown":   MailboxSettingsUserPurposeV2Unknown,
		"user":      MailboxSettingsUserPurposeV2User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsUserPurposeV2(input)
	return &out, nil
}

type ManagedAppRegistrationFlaggedReasons string

const (
	ManagedAppRegistrationFlaggedReasonsAndroidBootloaderUnlocked ManagedAppRegistrationFlaggedReasons = "androidBootloaderUnlocked"
	ManagedAppRegistrationFlaggedReasonsAndroidFactoryRomModified ManagedAppRegistrationFlaggedReasons = "androidFactoryRomModified"
	ManagedAppRegistrationFlaggedReasonsNone                      ManagedAppRegistrationFlaggedReasons = "none"
	ManagedAppRegistrationFlaggedReasonsRootedDevice              ManagedAppRegistrationFlaggedReasons = "rootedDevice"
)

func PossibleValuesForManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(ManagedAppRegistrationFlaggedReasonsAndroidBootloaderUnlocked),
		string(ManagedAppRegistrationFlaggedReasonsAndroidFactoryRomModified),
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
		"androidbootloaderunlocked": ManagedAppRegistrationFlaggedReasonsAndroidBootloaderUnlocked,
		"androidfactoryrommodified": ManagedAppRegistrationFlaggedReasonsAndroidFactoryRomModified,
		"none":                      ManagedAppRegistrationFlaggedReasonsNone,
		"rooteddevice":              ManagedAppRegistrationFlaggedReasonsRootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}

type ManagedDeviceChassisType string

const (
	ManagedDeviceChassisTypeDesktop          ManagedDeviceChassisType = "desktop"
	ManagedDeviceChassisTypeEnterpriseServer ManagedDeviceChassisType = "enterpriseServer"
	ManagedDeviceChassisTypeLaptop           ManagedDeviceChassisType = "laptop"
	ManagedDeviceChassisTypeMobileOther      ManagedDeviceChassisType = "mobileOther"
	ManagedDeviceChassisTypeMobileUnknown    ManagedDeviceChassisType = "mobileUnknown"
	ManagedDeviceChassisTypePhone            ManagedDeviceChassisType = "phone"
	ManagedDeviceChassisTypeTablet           ManagedDeviceChassisType = "tablet"
	ManagedDeviceChassisTypeUnknown          ManagedDeviceChassisType = "unknown"
	ManagedDeviceChassisTypeWorksWorkstation ManagedDeviceChassisType = "worksWorkstation"
)

func PossibleValuesForManagedDeviceChassisType() []string {
	return []string{
		string(ManagedDeviceChassisTypeDesktop),
		string(ManagedDeviceChassisTypeEnterpriseServer),
		string(ManagedDeviceChassisTypeLaptop),
		string(ManagedDeviceChassisTypeMobileOther),
		string(ManagedDeviceChassisTypeMobileUnknown),
		string(ManagedDeviceChassisTypePhone),
		string(ManagedDeviceChassisTypeTablet),
		string(ManagedDeviceChassisTypeUnknown),
		string(ManagedDeviceChassisTypeWorksWorkstation),
	}
}

func (s *ManagedDeviceChassisType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceChassisType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceChassisType(input string) (*ManagedDeviceChassisType, error) {
	vals := map[string]ManagedDeviceChassisType{
		"desktop":          ManagedDeviceChassisTypeDesktop,
		"enterpriseserver": ManagedDeviceChassisTypeEnterpriseServer,
		"laptop":           ManagedDeviceChassisTypeLaptop,
		"mobileother":      ManagedDeviceChassisTypeMobileOther,
		"mobileunknown":    ManagedDeviceChassisTypeMobileUnknown,
		"phone":            ManagedDeviceChassisTypePhone,
		"tablet":           ManagedDeviceChassisTypeTablet,
		"unknown":          ManagedDeviceChassisTypeUnknown,
		"worksworkstation": ManagedDeviceChassisTypeWorksWorkstation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceChassisType(input)
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
	ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseCorporateWorkProfile ManagedDeviceDeviceEnrollmentType = "androidEnterpriseCorporateWorkProfile"
	ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseDedicatedDevice      ManagedDeviceDeviceEnrollmentType = "androidEnterpriseDedicatedDevice"
	ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseFullyManaged         ManagedDeviceDeviceEnrollmentType = "androidEnterpriseFullyManaged"
	ManagedDeviceDeviceEnrollmentTypeAppleBulkWithUser                     ManagedDeviceDeviceEnrollmentType = "appleBulkWithUser"
	ManagedDeviceDeviceEnrollmentTypeAppleBulkWithoutUser                  ManagedDeviceDeviceEnrollmentType = "appleBulkWithoutUser"
	ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollment                   ManagedDeviceDeviceEnrollmentType = "appleUserEnrollment"
	ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollmentWithServiceAccount ManagedDeviceDeviceEnrollmentType = "appleUserEnrollmentWithServiceAccount"
	ManagedDeviceDeviceEnrollmentTypeAzureAdJoinUsingAzureVmExtension      ManagedDeviceDeviceEnrollmentType = "azureAdJoinUsingAzureVmExtension"
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
		string(ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseCorporateWorkProfile),
		string(ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseDedicatedDevice),
		string(ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseFullyManaged),
		string(ManagedDeviceDeviceEnrollmentTypeAppleBulkWithUser),
		string(ManagedDeviceDeviceEnrollmentTypeAppleBulkWithoutUser),
		string(ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollment),
		string(ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollmentWithServiceAccount),
		string(ManagedDeviceDeviceEnrollmentTypeAzureAdJoinUsingAzureVmExtension),
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
		"androidenterprisecorporateworkprofile": ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseCorporateWorkProfile,
		"androidenterprisededicateddevice":      ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseDedicatedDevice,
		"androidenterprisefullymanaged":         ManagedDeviceDeviceEnrollmentTypeAndroidEnterpriseFullyManaged,
		"applebulkwithuser":                     ManagedDeviceDeviceEnrollmentTypeAppleBulkWithUser,
		"applebulkwithoutuser":                  ManagedDeviceDeviceEnrollmentTypeAppleBulkWithoutUser,
		"appleuserenrollment":                   ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": ManagedDeviceDeviceEnrollmentTypeAppleUserEnrollmentWithServiceAccount,
		"azureadjoinusingazurevmextension":      ManagedDeviceDeviceEnrollmentTypeAzureAdJoinUsingAzureVmExtension,
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

type ManagedDeviceDeviceType string

const (
	ManagedDeviceDeviceTypeAndroid           ManagedDeviceDeviceType = "android"
	ManagedDeviceDeviceTypeAndroidEnterprise ManagedDeviceDeviceType = "androidEnterprise"
	ManagedDeviceDeviceTypeAndroidForWork    ManagedDeviceDeviceType = "androidForWork"
	ManagedDeviceDeviceTypeAndroidnGMS       ManagedDeviceDeviceType = "androidnGMS"
	ManagedDeviceDeviceTypeBlackberry        ManagedDeviceDeviceType = "blackberry"
	ManagedDeviceDeviceTypeChromeOS          ManagedDeviceDeviceType = "chromeOS"
	ManagedDeviceDeviceTypeCloudPC           ManagedDeviceDeviceType = "cloudPC"
	ManagedDeviceDeviceTypeDesktop           ManagedDeviceDeviceType = "desktop"
	ManagedDeviceDeviceTypeHoloLens          ManagedDeviceDeviceType = "holoLens"
	ManagedDeviceDeviceTypeIPad              ManagedDeviceDeviceType = "iPad"
	ManagedDeviceDeviceTypeIPhone            ManagedDeviceDeviceType = "iPhone"
	ManagedDeviceDeviceTypeIPod              ManagedDeviceDeviceType = "iPod"
	ManagedDeviceDeviceTypeISocConsumer      ManagedDeviceDeviceType = "iSocConsumer"
	ManagedDeviceDeviceTypeLinux             ManagedDeviceDeviceType = "linux"
	ManagedDeviceDeviceTypeMac               ManagedDeviceDeviceType = "mac"
	ManagedDeviceDeviceTypeMacMDM            ManagedDeviceDeviceType = "macMDM"
	ManagedDeviceDeviceTypeNokia             ManagedDeviceDeviceType = "nokia"
	ManagedDeviceDeviceTypePalm              ManagedDeviceDeviceType = "palm"
	ManagedDeviceDeviceTypeSurfaceHub        ManagedDeviceDeviceType = "surfaceHub"
	ManagedDeviceDeviceTypeUnix              ManagedDeviceDeviceType = "unix"
	ManagedDeviceDeviceTypeUnknown           ManagedDeviceDeviceType = "unknown"
	ManagedDeviceDeviceTypeWinCE             ManagedDeviceDeviceType = "winCE"
	ManagedDeviceDeviceTypeWinEmbedded       ManagedDeviceDeviceType = "winEmbedded"
	ManagedDeviceDeviceTypeWinMO6            ManagedDeviceDeviceType = "winMO6"
	ManagedDeviceDeviceTypeWindows10x        ManagedDeviceDeviceType = "windows10x"
	ManagedDeviceDeviceTypeWindowsPhone      ManagedDeviceDeviceType = "windowsPhone"
	ManagedDeviceDeviceTypeWindowsRT         ManagedDeviceDeviceType = "windowsRT"
)

func PossibleValuesForManagedDeviceDeviceType() []string {
	return []string{
		string(ManagedDeviceDeviceTypeAndroid),
		string(ManagedDeviceDeviceTypeAndroidEnterprise),
		string(ManagedDeviceDeviceTypeAndroidForWork),
		string(ManagedDeviceDeviceTypeAndroidnGMS),
		string(ManagedDeviceDeviceTypeBlackberry),
		string(ManagedDeviceDeviceTypeChromeOS),
		string(ManagedDeviceDeviceTypeCloudPC),
		string(ManagedDeviceDeviceTypeDesktop),
		string(ManagedDeviceDeviceTypeHoloLens),
		string(ManagedDeviceDeviceTypeIPad),
		string(ManagedDeviceDeviceTypeIPhone),
		string(ManagedDeviceDeviceTypeIPod),
		string(ManagedDeviceDeviceTypeISocConsumer),
		string(ManagedDeviceDeviceTypeLinux),
		string(ManagedDeviceDeviceTypeMac),
		string(ManagedDeviceDeviceTypeMacMDM),
		string(ManagedDeviceDeviceTypeNokia),
		string(ManagedDeviceDeviceTypePalm),
		string(ManagedDeviceDeviceTypeSurfaceHub),
		string(ManagedDeviceDeviceTypeUnix),
		string(ManagedDeviceDeviceTypeUnknown),
		string(ManagedDeviceDeviceTypeWinCE),
		string(ManagedDeviceDeviceTypeWinEmbedded),
		string(ManagedDeviceDeviceTypeWinMO6),
		string(ManagedDeviceDeviceTypeWindows10x),
		string(ManagedDeviceDeviceTypeWindowsPhone),
		string(ManagedDeviceDeviceTypeWindowsRT),
	}
}

func (s *ManagedDeviceDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceDeviceType(input string) (*ManagedDeviceDeviceType, error) {
	vals := map[string]ManagedDeviceDeviceType{
		"android":           ManagedDeviceDeviceTypeAndroid,
		"androidenterprise": ManagedDeviceDeviceTypeAndroidEnterprise,
		"androidforwork":    ManagedDeviceDeviceTypeAndroidForWork,
		"androidngms":       ManagedDeviceDeviceTypeAndroidnGMS,
		"blackberry":        ManagedDeviceDeviceTypeBlackberry,
		"chromeos":          ManagedDeviceDeviceTypeChromeOS,
		"cloudpc":           ManagedDeviceDeviceTypeCloudPC,
		"desktop":           ManagedDeviceDeviceTypeDesktop,
		"hololens":          ManagedDeviceDeviceTypeHoloLens,
		"ipad":              ManagedDeviceDeviceTypeIPad,
		"iphone":            ManagedDeviceDeviceTypeIPhone,
		"ipod":              ManagedDeviceDeviceTypeIPod,
		"isocconsumer":      ManagedDeviceDeviceTypeISocConsumer,
		"linux":             ManagedDeviceDeviceTypeLinux,
		"mac":               ManagedDeviceDeviceTypeMac,
		"macmdm":            ManagedDeviceDeviceTypeMacMDM,
		"nokia":             ManagedDeviceDeviceTypeNokia,
		"palm":              ManagedDeviceDeviceTypePalm,
		"surfacehub":        ManagedDeviceDeviceTypeSurfaceHub,
		"unix":              ManagedDeviceDeviceTypeUnix,
		"unknown":           ManagedDeviceDeviceTypeUnknown,
		"wince":             ManagedDeviceDeviceTypeWinCE,
		"winembedded":       ManagedDeviceDeviceTypeWinEmbedded,
		"winmo6":            ManagedDeviceDeviceTypeWinMO6,
		"windows10x":        ManagedDeviceDeviceTypeWindows10x,
		"windowsphone":      ManagedDeviceDeviceTypeWindowsPhone,
		"windowsrt":         ManagedDeviceDeviceTypeWindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceType(input)
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

type ManagedDeviceJoinType string

const (
	ManagedDeviceJoinTypeAzureADJoined       ManagedDeviceJoinType = "azureADJoined"
	ManagedDeviceJoinTypeAzureADRegistered   ManagedDeviceJoinType = "azureADRegistered"
	ManagedDeviceJoinTypeHybridAzureADJoined ManagedDeviceJoinType = "hybridAzureADJoined"
	ManagedDeviceJoinTypeUnknown             ManagedDeviceJoinType = "unknown"
)

func PossibleValuesForManagedDeviceJoinType() []string {
	return []string{
		string(ManagedDeviceJoinTypeAzureADJoined),
		string(ManagedDeviceJoinTypeAzureADRegistered),
		string(ManagedDeviceJoinTypeHybridAzureADJoined),
		string(ManagedDeviceJoinTypeUnknown),
	}
}

func (s *ManagedDeviceJoinType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceJoinType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceJoinType(input string) (*ManagedDeviceJoinType, error) {
	vals := map[string]ManagedDeviceJoinType{
		"azureadjoined":       ManagedDeviceJoinTypeAzureADJoined,
		"azureadregistered":   ManagedDeviceJoinTypeAzureADRegistered,
		"hybridazureadjoined": ManagedDeviceJoinTypeHybridAzureADJoined,
		"unknown":             ManagedDeviceJoinTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceJoinType(input)
	return &out, nil
}

type ManagedDeviceLostModeState string

const (
	ManagedDeviceLostModeStateDisabled ManagedDeviceLostModeState = "disabled"
	ManagedDeviceLostModeStateEnabled  ManagedDeviceLostModeState = "enabled"
)

func PossibleValuesForManagedDeviceLostModeState() []string {
	return []string{
		string(ManagedDeviceLostModeStateDisabled),
		string(ManagedDeviceLostModeStateEnabled),
	}
}

func (s *ManagedDeviceLostModeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceLostModeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceLostModeState(input string) (*ManagedDeviceLostModeState, error) {
	vals := map[string]ManagedDeviceLostModeState{
		"disabled": ManagedDeviceLostModeStateDisabled,
		"enabled":  ManagedDeviceLostModeStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceLostModeState(input)
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
	ManagedDeviceManagementAgentIntuneAosp                        ManagedDeviceManagementAgent = "intuneAosp"
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
		string(ManagedDeviceManagementAgentIntuneAosp),
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
		"intuneaosp":                        ManagedDeviceManagementAgentIntuneAosp,
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

type ManagedDeviceManagementFeatures string

const (
	ManagedDeviceManagementFeaturesMicrosoftManagedDesktop ManagedDeviceManagementFeatures = "microsoftManagedDesktop"
	ManagedDeviceManagementFeaturesNone                    ManagedDeviceManagementFeatures = "none"
)

func PossibleValuesForManagedDeviceManagementFeatures() []string {
	return []string{
		string(ManagedDeviceManagementFeaturesMicrosoftManagedDesktop),
		string(ManagedDeviceManagementFeaturesNone),
	}
}

func (s *ManagedDeviceManagementFeatures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceManagementFeatures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceManagementFeatures(input string) (*ManagedDeviceManagementFeatures, error) {
	vals := map[string]ManagedDeviceManagementFeatures{
		"microsoftmanageddesktop": ManagedDeviceManagementFeaturesMicrosoftManagedDesktop,
		"none":                    ManagedDeviceManagementFeaturesNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementFeatures(input)
	return &out, nil
}

type ManagedDeviceManagementState string

const (
	ManagedDeviceManagementStateDeletePending  ManagedDeviceManagementState = "deletePending"
	ManagedDeviceManagementStateDiscovered     ManagedDeviceManagementState = "discovered"
	ManagedDeviceManagementStateManaged        ManagedDeviceManagementState = "managed"
	ManagedDeviceManagementStateRetireCanceled ManagedDeviceManagementState = "retireCanceled"
	ManagedDeviceManagementStateRetireFailed   ManagedDeviceManagementState = "retireFailed"
	ManagedDeviceManagementStateRetireIssued   ManagedDeviceManagementState = "retireIssued"
	ManagedDeviceManagementStateRetirePending  ManagedDeviceManagementState = "retirePending"
	ManagedDeviceManagementStateUnhealthy      ManagedDeviceManagementState = "unhealthy"
	ManagedDeviceManagementStateWipeCanceled   ManagedDeviceManagementState = "wipeCanceled"
	ManagedDeviceManagementStateWipeFailed     ManagedDeviceManagementState = "wipeFailed"
	ManagedDeviceManagementStateWipeIssued     ManagedDeviceManagementState = "wipeIssued"
	ManagedDeviceManagementStateWipePending    ManagedDeviceManagementState = "wipePending"
)

func PossibleValuesForManagedDeviceManagementState() []string {
	return []string{
		string(ManagedDeviceManagementStateDeletePending),
		string(ManagedDeviceManagementStateDiscovered),
		string(ManagedDeviceManagementStateManaged),
		string(ManagedDeviceManagementStateRetireCanceled),
		string(ManagedDeviceManagementStateRetireFailed),
		string(ManagedDeviceManagementStateRetireIssued),
		string(ManagedDeviceManagementStateRetirePending),
		string(ManagedDeviceManagementStateUnhealthy),
		string(ManagedDeviceManagementStateWipeCanceled),
		string(ManagedDeviceManagementStateWipeFailed),
		string(ManagedDeviceManagementStateWipeIssued),
		string(ManagedDeviceManagementStateWipePending),
	}
}

func (s *ManagedDeviceManagementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceManagementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceManagementState(input string) (*ManagedDeviceManagementState, error) {
	vals := map[string]ManagedDeviceManagementState{
		"deletepending":  ManagedDeviceManagementStateDeletePending,
		"discovered":     ManagedDeviceManagementStateDiscovered,
		"managed":        ManagedDeviceManagementStateManaged,
		"retirecanceled": ManagedDeviceManagementStateRetireCanceled,
		"retirefailed":   ManagedDeviceManagementStateRetireFailed,
		"retireissued":   ManagedDeviceManagementStateRetireIssued,
		"retirepending":  ManagedDeviceManagementStateRetirePending,
		"unhealthy":      ManagedDeviceManagementStateUnhealthy,
		"wipecanceled":   ManagedDeviceManagementStateWipeCanceled,
		"wipefailed":     ManagedDeviceManagementStateWipeFailed,
		"wipeissued":     ManagedDeviceManagementStateWipeIssued,
		"wipepending":    ManagedDeviceManagementStateWipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementState(input)
	return &out, nil
}

type ManagedDeviceMobileAppConfigurationSettingStateState string

const (
	ManagedDeviceMobileAppConfigurationSettingStateStateCompliant     ManagedDeviceMobileAppConfigurationSettingStateState = "compliant"
	ManagedDeviceMobileAppConfigurationSettingStateStateConflict      ManagedDeviceMobileAppConfigurationSettingStateState = "conflict"
	ManagedDeviceMobileAppConfigurationSettingStateStateError         ManagedDeviceMobileAppConfigurationSettingStateState = "error"
	ManagedDeviceMobileAppConfigurationSettingStateStateNonCompliant  ManagedDeviceMobileAppConfigurationSettingStateState = "nonCompliant"
	ManagedDeviceMobileAppConfigurationSettingStateStateNotApplicable ManagedDeviceMobileAppConfigurationSettingStateState = "notApplicable"
	ManagedDeviceMobileAppConfigurationSettingStateStateNotAssigned   ManagedDeviceMobileAppConfigurationSettingStateState = "notAssigned"
	ManagedDeviceMobileAppConfigurationSettingStateStateRemediated    ManagedDeviceMobileAppConfigurationSettingStateState = "remediated"
	ManagedDeviceMobileAppConfigurationSettingStateStateUnknown       ManagedDeviceMobileAppConfigurationSettingStateState = "unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationSettingStateState() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationSettingStateStateCompliant),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateConflict),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateError),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateNonCompliant),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateNotApplicable),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateNotAssigned),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateRemediated),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateUnknown),
	}
}

func (s *ManagedDeviceMobileAppConfigurationSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationSettingStateState(input string) (*ManagedDeviceMobileAppConfigurationSettingStateState, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationSettingStateState{
		"compliant":     ManagedDeviceMobileAppConfigurationSettingStateStateCompliant,
		"conflict":      ManagedDeviceMobileAppConfigurationSettingStateStateConflict,
		"error":         ManagedDeviceMobileAppConfigurationSettingStateStateError,
		"noncompliant":  ManagedDeviceMobileAppConfigurationSettingStateStateNonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationSettingStateStateNotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationSettingStateStateNotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationSettingStateStateRemediated,
		"unknown":       ManagedDeviceMobileAppConfigurationSettingStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationSettingStateState(input)
	return &out, nil
}

type ManagedDeviceMobileAppConfigurationStatePlatformType string

const (
	ManagedDeviceMobileAppConfigurationStatePlatformTypeAll                ManagedDeviceMobileAppConfigurationStatePlatformType = "all"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroid            ManagedDeviceMobileAppConfigurationStatePlatformType = "android"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidAOSP        ManagedDeviceMobileAppConfigurationStatePlatformType = "androidAOSP"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidForWork     ManagedDeviceMobileAppConfigurationStatePlatformType = "androidForWork"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidWorkProfile ManagedDeviceMobileAppConfigurationStatePlatformType = "androidWorkProfile"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeIOS                ManagedDeviceMobileAppConfigurationStatePlatformType = "iOS"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeMacOS              ManagedDeviceMobileAppConfigurationStatePlatformType = "macOS"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows10AndLater  ManagedDeviceMobileAppConfigurationStatePlatformType = "windows10AndLater"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows10XProfile  ManagedDeviceMobileAppConfigurationStatePlatformType = "windows10XProfile"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows81AndLater  ManagedDeviceMobileAppConfigurationStatePlatformType = "windows81AndLater"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeWindowsPhone81     ManagedDeviceMobileAppConfigurationStatePlatformType = "windowsPhone81"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationStatePlatformType() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeAll),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroid),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidAOSP),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidForWork),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidWorkProfile),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeIOS),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeMacOS),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows10AndLater),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows10XProfile),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows81AndLater),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeWindowsPhone81),
	}
}

func (s *ManagedDeviceMobileAppConfigurationStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationStatePlatformType(input string) (*ManagedDeviceMobileAppConfigurationStatePlatformType, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationStatePlatformType{
		"all":                ManagedDeviceMobileAppConfigurationStatePlatformTypeAll,
		"android":            ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroid,
		"androidaosp":        ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidAOSP,
		"androidforwork":     ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidForWork,
		"androidworkprofile": ManagedDeviceMobileAppConfigurationStatePlatformTypeAndroidWorkProfile,
		"ios":                ManagedDeviceMobileAppConfigurationStatePlatformTypeIOS,
		"macos":              ManagedDeviceMobileAppConfigurationStatePlatformTypeMacOS,
		"windows10andlater":  ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows10AndLater,
		"windows10xprofile":  ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows10XProfile,
		"windows81andlater":  ManagedDeviceMobileAppConfigurationStatePlatformTypeWindows81AndLater,
		"windowsphone81":     ManagedDeviceMobileAppConfigurationStatePlatformTypeWindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationStatePlatformType(input)
	return &out, nil
}

type ManagedDeviceMobileAppConfigurationStateState string

const (
	ManagedDeviceMobileAppConfigurationStateStateCompliant     ManagedDeviceMobileAppConfigurationStateState = "compliant"
	ManagedDeviceMobileAppConfigurationStateStateConflict      ManagedDeviceMobileAppConfigurationStateState = "conflict"
	ManagedDeviceMobileAppConfigurationStateStateError         ManagedDeviceMobileAppConfigurationStateState = "error"
	ManagedDeviceMobileAppConfigurationStateStateNonCompliant  ManagedDeviceMobileAppConfigurationStateState = "nonCompliant"
	ManagedDeviceMobileAppConfigurationStateStateNotApplicable ManagedDeviceMobileAppConfigurationStateState = "notApplicable"
	ManagedDeviceMobileAppConfigurationStateStateNotAssigned   ManagedDeviceMobileAppConfigurationStateState = "notAssigned"
	ManagedDeviceMobileAppConfigurationStateStateRemediated    ManagedDeviceMobileAppConfigurationStateState = "remediated"
	ManagedDeviceMobileAppConfigurationStateStateUnknown       ManagedDeviceMobileAppConfigurationStateState = "unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationStateState() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationStateStateCompliant),
		string(ManagedDeviceMobileAppConfigurationStateStateConflict),
		string(ManagedDeviceMobileAppConfigurationStateStateError),
		string(ManagedDeviceMobileAppConfigurationStateStateNonCompliant),
		string(ManagedDeviceMobileAppConfigurationStateStateNotApplicable),
		string(ManagedDeviceMobileAppConfigurationStateStateNotAssigned),
		string(ManagedDeviceMobileAppConfigurationStateStateRemediated),
		string(ManagedDeviceMobileAppConfigurationStateStateUnknown),
	}
}

func (s *ManagedDeviceMobileAppConfigurationStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationStateState(input string) (*ManagedDeviceMobileAppConfigurationStateState, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationStateState{
		"compliant":     ManagedDeviceMobileAppConfigurationStateStateCompliant,
		"conflict":      ManagedDeviceMobileAppConfigurationStateStateConflict,
		"error":         ManagedDeviceMobileAppConfigurationStateStateError,
		"noncompliant":  ManagedDeviceMobileAppConfigurationStateStateNonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationStateStateNotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationStateStateNotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationStateStateRemediated,
		"unknown":       ManagedDeviceMobileAppConfigurationStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationStateState(input)
	return &out, nil
}

type ManagedDeviceOwnerType string

const (
	ManagedDeviceOwnerTypeCompany  ManagedDeviceOwnerType = "company"
	ManagedDeviceOwnerTypePersonal ManagedDeviceOwnerType = "personal"
	ManagedDeviceOwnerTypeUnknown  ManagedDeviceOwnerType = "unknown"
)

func PossibleValuesForManagedDeviceOwnerType() []string {
	return []string{
		string(ManagedDeviceOwnerTypeCompany),
		string(ManagedDeviceOwnerTypePersonal),
		string(ManagedDeviceOwnerTypeUnknown),
	}
}

func (s *ManagedDeviceOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceOwnerType(input string) (*ManagedDeviceOwnerType, error) {
	vals := map[string]ManagedDeviceOwnerType{
		"company":  ManagedDeviceOwnerTypeCompany,
		"personal": ManagedDeviceOwnerTypePersonal,
		"unknown":  ManagedDeviceOwnerTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceOwnerType(input)
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

type ManagedDeviceProcessorArchitecture string

const (
	ManagedDeviceProcessorArchitectureArM64   ManagedDeviceProcessorArchitecture = "arM64"
	ManagedDeviceProcessorArchitectureArm     ManagedDeviceProcessorArchitecture = "arm"
	ManagedDeviceProcessorArchitectureUnknown ManagedDeviceProcessorArchitecture = "unknown"
	ManagedDeviceProcessorArchitectureX64     ManagedDeviceProcessorArchitecture = "x64"
	ManagedDeviceProcessorArchitectureX86     ManagedDeviceProcessorArchitecture = "x86"
)

func PossibleValuesForManagedDeviceProcessorArchitecture() []string {
	return []string{
		string(ManagedDeviceProcessorArchitectureArM64),
		string(ManagedDeviceProcessorArchitectureArm),
		string(ManagedDeviceProcessorArchitectureUnknown),
		string(ManagedDeviceProcessorArchitectureX64),
		string(ManagedDeviceProcessorArchitectureX86),
	}
}

func (s *ManagedDeviceProcessorArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceProcessorArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceProcessorArchitecture(input string) (*ManagedDeviceProcessorArchitecture, error) {
	vals := map[string]ManagedDeviceProcessorArchitecture{
		"arm64":   ManagedDeviceProcessorArchitectureArM64,
		"arm":     ManagedDeviceProcessorArchitectureArm,
		"unknown": ManagedDeviceProcessorArchitectureUnknown,
		"x64":     ManagedDeviceProcessorArchitectureX64,
		"x86":     ManagedDeviceProcessorArchitectureX86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceProcessorArchitecture(input)
	return &out, nil
}

type MediaSourceContentCategory string

const (
	MediaSourceContentCategoryLiveStream      MediaSourceContentCategory = "liveStream"
	MediaSourceContentCategoryMeeting         MediaSourceContentCategory = "meeting"
	MediaSourceContentCategoryPresentation    MediaSourceContentCategory = "presentation"
	MediaSourceContentCategoryScreenRecording MediaSourceContentCategory = "screenRecording"
)

func PossibleValuesForMediaSourceContentCategory() []string {
	return []string{
		string(MediaSourceContentCategoryLiveStream),
		string(MediaSourceContentCategoryMeeting),
		string(MediaSourceContentCategoryPresentation),
		string(MediaSourceContentCategoryScreenRecording),
	}
}

func (s *MediaSourceContentCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaSourceContentCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaSourceContentCategory(input string) (*MediaSourceContentCategory, error) {
	vals := map[string]MediaSourceContentCategory{
		"livestream":      MediaSourceContentCategoryLiveStream,
		"meeting":         MediaSourceContentCategoryMeeting,
		"presentation":    MediaSourceContentCategoryPresentation,
		"screenrecording": MediaSourceContentCategoryScreenRecording,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaSourceContentCategory(input)
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

type MeetingRegistrationAllowedRegistrant string

const (
	MeetingRegistrationAllowedRegistrantEveryone     MeetingRegistrationAllowedRegistrant = "everyone"
	MeetingRegistrationAllowedRegistrantOrganization MeetingRegistrationAllowedRegistrant = "organization"
)

func PossibleValuesForMeetingRegistrationAllowedRegistrant() []string {
	return []string{
		string(MeetingRegistrationAllowedRegistrantEveryone),
		string(MeetingRegistrationAllowedRegistrantOrganization),
	}
}

func (s *MeetingRegistrationAllowedRegistrant) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingRegistrationAllowedRegistrant(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingRegistrationAllowedRegistrant(input string) (*MeetingRegistrationAllowedRegistrant, error) {
	vals := map[string]MeetingRegistrationAllowedRegistrant{
		"everyone":     MeetingRegistrationAllowedRegistrantEveryone,
		"organization": MeetingRegistrationAllowedRegistrantOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationAllowedRegistrant(input)
	return &out, nil
}

type MeetingRegistrationQuestionAnswerInputType string

const (
	MeetingRegistrationQuestionAnswerInputTypeRadioButton MeetingRegistrationQuestionAnswerInputType = "radioButton"
	MeetingRegistrationQuestionAnswerInputTypeText        MeetingRegistrationQuestionAnswerInputType = "text"
)

func PossibleValuesForMeetingRegistrationQuestionAnswerInputType() []string {
	return []string{
		string(MeetingRegistrationQuestionAnswerInputTypeRadioButton),
		string(MeetingRegistrationQuestionAnswerInputTypeText),
	}
}

func (s *MeetingRegistrationQuestionAnswerInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingRegistrationQuestionAnswerInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingRegistrationQuestionAnswerInputType(input string) (*MeetingRegistrationQuestionAnswerInputType, error) {
	vals := map[string]MeetingRegistrationQuestionAnswerInputType{
		"radiobutton": MeetingRegistrationQuestionAnswerInputTypeRadioButton,
		"text":        MeetingRegistrationQuestionAnswerInputTypeText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationQuestionAnswerInputType(input)
	return &out, nil
}

type MembershipRuleProcessingStatusStatus string

const (
	MembershipRuleProcessingStatusStatusFailed                 MembershipRuleProcessingStatusStatus = "Failed"
	MembershipRuleProcessingStatusStatusNotStarted             MembershipRuleProcessingStatusStatus = "NotStarted"
	MembershipRuleProcessingStatusStatusRunning                MembershipRuleProcessingStatusStatus = "Running"
	MembershipRuleProcessingStatusStatusSucceeded              MembershipRuleProcessingStatusStatus = "Succeeded"
	MembershipRuleProcessingStatusStatusUnsupportedFutureValue MembershipRuleProcessingStatusStatus = "UnsupportedFutureValue"
)

func PossibleValuesForMembershipRuleProcessingStatusStatus() []string {
	return []string{
		string(MembershipRuleProcessingStatusStatusFailed),
		string(MembershipRuleProcessingStatusStatusNotStarted),
		string(MembershipRuleProcessingStatusStatusRunning),
		string(MembershipRuleProcessingStatusStatusSucceeded),
		string(MembershipRuleProcessingStatusStatusUnsupportedFutureValue),
	}
}

func (s *MembershipRuleProcessingStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMembershipRuleProcessingStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMembershipRuleProcessingStatusStatus(input string) (*MembershipRuleProcessingStatusStatus, error) {
	vals := map[string]MembershipRuleProcessingStatusStatus{
		"failed":                 MembershipRuleProcessingStatusStatusFailed,
		"notstarted":             MembershipRuleProcessingStatusStatusNotStarted,
		"running":                MembershipRuleProcessingStatusStatusRunning,
		"succeeded":              MembershipRuleProcessingStatusStatusSucceeded,
		"unsupportedfuturevalue": MembershipRuleProcessingStatusStatusUnsupportedFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MembershipRuleProcessingStatusStatus(input)
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

type MicrosoftAuthenticatorAuthenticationMethodClientAppName string

const (
	MicrosoftAuthenticatorAuthenticationMethodClientAppNameMicrosoftAuthenticator MicrosoftAuthenticatorAuthenticationMethodClientAppName = "microsoftAuthenticator"
	MicrosoftAuthenticatorAuthenticationMethodClientAppNameOutlookMobile          MicrosoftAuthenticatorAuthenticationMethodClientAppName = "outlookMobile"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodClientAppName() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodClientAppNameMicrosoftAuthenticator),
		string(MicrosoftAuthenticatorAuthenticationMethodClientAppNameOutlookMobile),
	}
}

func (s *MicrosoftAuthenticatorAuthenticationMethodClientAppName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftAuthenticatorAuthenticationMethodClientAppName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftAuthenticatorAuthenticationMethodClientAppName(input string) (*MicrosoftAuthenticatorAuthenticationMethodClientAppName, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodClientAppName{
		"microsoftauthenticator": MicrosoftAuthenticatorAuthenticationMethodClientAppNameMicrosoftAuthenticator,
		"outlookmobile":          MicrosoftAuthenticatorAuthenticationMethodClientAppNameOutlookMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodClientAppName(input)
	return &out, nil
}

type MobileAppIntentAndStateDetailInstallState string

const (
	MobileAppIntentAndStateDetailInstallStateFailed          MobileAppIntentAndStateDetailInstallState = "failed"
	MobileAppIntentAndStateDetailInstallStateInstalled       MobileAppIntentAndStateDetailInstallState = "installed"
	MobileAppIntentAndStateDetailInstallStateNotApplicable   MobileAppIntentAndStateDetailInstallState = "notApplicable"
	MobileAppIntentAndStateDetailInstallStateNotInstalled    MobileAppIntentAndStateDetailInstallState = "notInstalled"
	MobileAppIntentAndStateDetailInstallStatePendingInstall  MobileAppIntentAndStateDetailInstallState = "pendingInstall"
	MobileAppIntentAndStateDetailInstallStateUninstallFailed MobileAppIntentAndStateDetailInstallState = "uninstallFailed"
	MobileAppIntentAndStateDetailInstallStateUnknown         MobileAppIntentAndStateDetailInstallState = "unknown"
)

func PossibleValuesForMobileAppIntentAndStateDetailInstallState() []string {
	return []string{
		string(MobileAppIntentAndStateDetailInstallStateFailed),
		string(MobileAppIntentAndStateDetailInstallStateInstalled),
		string(MobileAppIntentAndStateDetailInstallStateNotApplicable),
		string(MobileAppIntentAndStateDetailInstallStateNotInstalled),
		string(MobileAppIntentAndStateDetailInstallStatePendingInstall),
		string(MobileAppIntentAndStateDetailInstallStateUninstallFailed),
		string(MobileAppIntentAndStateDetailInstallStateUnknown),
	}
}

func (s *MobileAppIntentAndStateDetailInstallState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppIntentAndStateDetailInstallState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppIntentAndStateDetailInstallState(input string) (*MobileAppIntentAndStateDetailInstallState, error) {
	vals := map[string]MobileAppIntentAndStateDetailInstallState{
		"failed":          MobileAppIntentAndStateDetailInstallStateFailed,
		"installed":       MobileAppIntentAndStateDetailInstallStateInstalled,
		"notapplicable":   MobileAppIntentAndStateDetailInstallStateNotApplicable,
		"notinstalled":    MobileAppIntentAndStateDetailInstallStateNotInstalled,
		"pendinginstall":  MobileAppIntentAndStateDetailInstallStatePendingInstall,
		"uninstallfailed": MobileAppIntentAndStateDetailInstallStateUninstallFailed,
		"unknown":         MobileAppIntentAndStateDetailInstallStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppIntentAndStateDetailInstallState(input)
	return &out, nil
}

type MobileAppIntentAndStateDetailMobileAppIntent string

const (
	MobileAppIntentAndStateDetailMobileAppIntentAvailable                         MobileAppIntentAndStateDetailMobileAppIntent = "available"
	MobileAppIntentAndStateDetailMobileAppIntentAvailableInstallWithoutEnrollment MobileAppIntentAndStateDetailMobileAppIntent = "availableInstallWithoutEnrollment"
	MobileAppIntentAndStateDetailMobileAppIntentExclude                           MobileAppIntentAndStateDetailMobileAppIntent = "exclude"
	MobileAppIntentAndStateDetailMobileAppIntentNotAvailable                      MobileAppIntentAndStateDetailMobileAppIntent = "notAvailable"
	MobileAppIntentAndStateDetailMobileAppIntentRequiredAndAvailableInstall       MobileAppIntentAndStateDetailMobileAppIntent = "requiredAndAvailableInstall"
	MobileAppIntentAndStateDetailMobileAppIntentRequiredInstall                   MobileAppIntentAndStateDetailMobileAppIntent = "requiredInstall"
	MobileAppIntentAndStateDetailMobileAppIntentRequiredUninstall                 MobileAppIntentAndStateDetailMobileAppIntent = "requiredUninstall"
)

func PossibleValuesForMobileAppIntentAndStateDetailMobileAppIntent() []string {
	return []string{
		string(MobileAppIntentAndStateDetailMobileAppIntentAvailable),
		string(MobileAppIntentAndStateDetailMobileAppIntentAvailableInstallWithoutEnrollment),
		string(MobileAppIntentAndStateDetailMobileAppIntentExclude),
		string(MobileAppIntentAndStateDetailMobileAppIntentNotAvailable),
		string(MobileAppIntentAndStateDetailMobileAppIntentRequiredAndAvailableInstall),
		string(MobileAppIntentAndStateDetailMobileAppIntentRequiredInstall),
		string(MobileAppIntentAndStateDetailMobileAppIntentRequiredUninstall),
	}
}

func (s *MobileAppIntentAndStateDetailMobileAppIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppIntentAndStateDetailMobileAppIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppIntentAndStateDetailMobileAppIntent(input string) (*MobileAppIntentAndStateDetailMobileAppIntent, error) {
	vals := map[string]MobileAppIntentAndStateDetailMobileAppIntent{
		"available":                         MobileAppIntentAndStateDetailMobileAppIntentAvailable,
		"availableinstallwithoutenrollment": MobileAppIntentAndStateDetailMobileAppIntentAvailableInstallWithoutEnrollment,
		"exclude":                           MobileAppIntentAndStateDetailMobileAppIntentExclude,
		"notavailable":                      MobileAppIntentAndStateDetailMobileAppIntentNotAvailable,
		"requiredandavailableinstall":       MobileAppIntentAndStateDetailMobileAppIntentRequiredAndAvailableInstall,
		"requiredinstall":                   MobileAppIntentAndStateDetailMobileAppIntentRequiredInstall,
		"requireduninstall":                 MobileAppIntentAndStateDetailMobileAppIntentRequiredUninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppIntentAndStateDetailMobileAppIntent(input)
	return &out, nil
}

type MobileAppSupportedDeviceTypeType string

const (
	MobileAppSupportedDeviceTypeTypeAndroid           MobileAppSupportedDeviceTypeType = "android"
	MobileAppSupportedDeviceTypeTypeAndroidEnterprise MobileAppSupportedDeviceTypeType = "androidEnterprise"
	MobileAppSupportedDeviceTypeTypeAndroidForWork    MobileAppSupportedDeviceTypeType = "androidForWork"
	MobileAppSupportedDeviceTypeTypeAndroidnGMS       MobileAppSupportedDeviceTypeType = "androidnGMS"
	MobileAppSupportedDeviceTypeTypeBlackberry        MobileAppSupportedDeviceTypeType = "blackberry"
	MobileAppSupportedDeviceTypeTypeChromeOS          MobileAppSupportedDeviceTypeType = "chromeOS"
	MobileAppSupportedDeviceTypeTypeCloudPC           MobileAppSupportedDeviceTypeType = "cloudPC"
	MobileAppSupportedDeviceTypeTypeDesktop           MobileAppSupportedDeviceTypeType = "desktop"
	MobileAppSupportedDeviceTypeTypeHoloLens          MobileAppSupportedDeviceTypeType = "holoLens"
	MobileAppSupportedDeviceTypeTypeIPad              MobileAppSupportedDeviceTypeType = "iPad"
	MobileAppSupportedDeviceTypeTypeIPhone            MobileAppSupportedDeviceTypeType = "iPhone"
	MobileAppSupportedDeviceTypeTypeIPod              MobileAppSupportedDeviceTypeType = "iPod"
	MobileAppSupportedDeviceTypeTypeISocConsumer      MobileAppSupportedDeviceTypeType = "iSocConsumer"
	MobileAppSupportedDeviceTypeTypeLinux             MobileAppSupportedDeviceTypeType = "linux"
	MobileAppSupportedDeviceTypeTypeMac               MobileAppSupportedDeviceTypeType = "mac"
	MobileAppSupportedDeviceTypeTypeMacMDM            MobileAppSupportedDeviceTypeType = "macMDM"
	MobileAppSupportedDeviceTypeTypeNokia             MobileAppSupportedDeviceTypeType = "nokia"
	MobileAppSupportedDeviceTypeTypePalm              MobileAppSupportedDeviceTypeType = "palm"
	MobileAppSupportedDeviceTypeTypeSurfaceHub        MobileAppSupportedDeviceTypeType = "surfaceHub"
	MobileAppSupportedDeviceTypeTypeUnix              MobileAppSupportedDeviceTypeType = "unix"
	MobileAppSupportedDeviceTypeTypeUnknown           MobileAppSupportedDeviceTypeType = "unknown"
	MobileAppSupportedDeviceTypeTypeWinCE             MobileAppSupportedDeviceTypeType = "winCE"
	MobileAppSupportedDeviceTypeTypeWinEmbedded       MobileAppSupportedDeviceTypeType = "winEmbedded"
	MobileAppSupportedDeviceTypeTypeWinMO6            MobileAppSupportedDeviceTypeType = "winMO6"
	MobileAppSupportedDeviceTypeTypeWindows10x        MobileAppSupportedDeviceTypeType = "windows10x"
	MobileAppSupportedDeviceTypeTypeWindowsPhone      MobileAppSupportedDeviceTypeType = "windowsPhone"
	MobileAppSupportedDeviceTypeTypeWindowsRT         MobileAppSupportedDeviceTypeType = "windowsRT"
)

func PossibleValuesForMobileAppSupportedDeviceTypeType() []string {
	return []string{
		string(MobileAppSupportedDeviceTypeTypeAndroid),
		string(MobileAppSupportedDeviceTypeTypeAndroidEnterprise),
		string(MobileAppSupportedDeviceTypeTypeAndroidForWork),
		string(MobileAppSupportedDeviceTypeTypeAndroidnGMS),
		string(MobileAppSupportedDeviceTypeTypeBlackberry),
		string(MobileAppSupportedDeviceTypeTypeChromeOS),
		string(MobileAppSupportedDeviceTypeTypeCloudPC),
		string(MobileAppSupportedDeviceTypeTypeDesktop),
		string(MobileAppSupportedDeviceTypeTypeHoloLens),
		string(MobileAppSupportedDeviceTypeTypeIPad),
		string(MobileAppSupportedDeviceTypeTypeIPhone),
		string(MobileAppSupportedDeviceTypeTypeIPod),
		string(MobileAppSupportedDeviceTypeTypeISocConsumer),
		string(MobileAppSupportedDeviceTypeTypeLinux),
		string(MobileAppSupportedDeviceTypeTypeMac),
		string(MobileAppSupportedDeviceTypeTypeMacMDM),
		string(MobileAppSupportedDeviceTypeTypeNokia),
		string(MobileAppSupportedDeviceTypeTypePalm),
		string(MobileAppSupportedDeviceTypeTypeSurfaceHub),
		string(MobileAppSupportedDeviceTypeTypeUnix),
		string(MobileAppSupportedDeviceTypeTypeUnknown),
		string(MobileAppSupportedDeviceTypeTypeWinCE),
		string(MobileAppSupportedDeviceTypeTypeWinEmbedded),
		string(MobileAppSupportedDeviceTypeTypeWinMO6),
		string(MobileAppSupportedDeviceTypeTypeWindows10x),
		string(MobileAppSupportedDeviceTypeTypeWindowsPhone),
		string(MobileAppSupportedDeviceTypeTypeWindowsRT),
	}
}

func (s *MobileAppSupportedDeviceTypeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppSupportedDeviceTypeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppSupportedDeviceTypeType(input string) (*MobileAppSupportedDeviceTypeType, error) {
	vals := map[string]MobileAppSupportedDeviceTypeType{
		"android":           MobileAppSupportedDeviceTypeTypeAndroid,
		"androidenterprise": MobileAppSupportedDeviceTypeTypeAndroidEnterprise,
		"androidforwork":    MobileAppSupportedDeviceTypeTypeAndroidForWork,
		"androidngms":       MobileAppSupportedDeviceTypeTypeAndroidnGMS,
		"blackberry":        MobileAppSupportedDeviceTypeTypeBlackberry,
		"chromeos":          MobileAppSupportedDeviceTypeTypeChromeOS,
		"cloudpc":           MobileAppSupportedDeviceTypeTypeCloudPC,
		"desktop":           MobileAppSupportedDeviceTypeTypeDesktop,
		"hololens":          MobileAppSupportedDeviceTypeTypeHoloLens,
		"ipad":              MobileAppSupportedDeviceTypeTypeIPad,
		"iphone":            MobileAppSupportedDeviceTypeTypeIPhone,
		"ipod":              MobileAppSupportedDeviceTypeTypeIPod,
		"isocconsumer":      MobileAppSupportedDeviceTypeTypeISocConsumer,
		"linux":             MobileAppSupportedDeviceTypeTypeLinux,
		"mac":               MobileAppSupportedDeviceTypeTypeMac,
		"macmdm":            MobileAppSupportedDeviceTypeTypeMacMDM,
		"nokia":             MobileAppSupportedDeviceTypeTypeNokia,
		"palm":              MobileAppSupportedDeviceTypeTypePalm,
		"surfacehub":        MobileAppSupportedDeviceTypeTypeSurfaceHub,
		"unix":              MobileAppSupportedDeviceTypeTypeUnix,
		"unknown":           MobileAppSupportedDeviceTypeTypeUnknown,
		"wince":             MobileAppSupportedDeviceTypeTypeWinCE,
		"winembedded":       MobileAppSupportedDeviceTypeTypeWinEmbedded,
		"winmo6":            MobileAppSupportedDeviceTypeTypeWinMO6,
		"windows10x":        MobileAppSupportedDeviceTypeTypeWindows10x,
		"windowsphone":      MobileAppSupportedDeviceTypeTypeWindowsPhone,
		"windowsrt":         MobileAppSupportedDeviceTypeTypeWindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupportedDeviceTypeType(input)
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

type NotificationPriority string

const (
	NotificationPriorityHigh NotificationPriority = "High"
	NotificationPriorityLow  NotificationPriority = "Low"
	NotificationPriorityNone NotificationPriority = "None"
)

func PossibleValuesForNotificationPriority() []string {
	return []string{
		string(NotificationPriorityHigh),
		string(NotificationPriorityLow),
		string(NotificationPriorityNone),
	}
}

func (s *NotificationPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotificationPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotificationPriority(input string) (*NotificationPriority, error) {
	vals := map[string]NotificationPriority{
		"high": NotificationPriorityHigh,
		"low":  NotificationPriorityLow,
		"none": NotificationPriorityNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationPriority(input)
	return &out, nil
}

type ObjectDefinitionMetadataEntryKey string

const (
	ObjectDefinitionMetadataEntryKeyBaseObjectName               ObjectDefinitionMetadataEntryKey = "BaseObjectName"
	ObjectDefinitionMetadataEntryKeyConnectorDataStorageRequired ObjectDefinitionMetadataEntryKey = "ConnectorDataStorageRequired"
	ObjectDefinitionMetadataEntryKeyExtensions                   ObjectDefinitionMetadataEntryKey = "Extensions"
	ObjectDefinitionMetadataEntryKeyIsSoftDeletionSupported      ObjectDefinitionMetadataEntryKey = "IsSoftDeletionSupported"
	ObjectDefinitionMetadataEntryKeyIsSynchronizeAllSupported    ObjectDefinitionMetadataEntryKey = "IsSynchronizeAllSupported"
	ObjectDefinitionMetadataEntryKeyPropertyNameAccountEnabled   ObjectDefinitionMetadataEntryKey = "PropertyNameAccountEnabled"
	ObjectDefinitionMetadataEntryKeyPropertyNameSoftDeleted      ObjectDefinitionMetadataEntryKey = "PropertyNameSoftDeleted"
)

func PossibleValuesForObjectDefinitionMetadataEntryKey() []string {
	return []string{
		string(ObjectDefinitionMetadataEntryKeyBaseObjectName),
		string(ObjectDefinitionMetadataEntryKeyConnectorDataStorageRequired),
		string(ObjectDefinitionMetadataEntryKeyExtensions),
		string(ObjectDefinitionMetadataEntryKeyIsSoftDeletionSupported),
		string(ObjectDefinitionMetadataEntryKeyIsSynchronizeAllSupported),
		string(ObjectDefinitionMetadataEntryKeyPropertyNameAccountEnabled),
		string(ObjectDefinitionMetadataEntryKeyPropertyNameSoftDeleted),
	}
}

func (s *ObjectDefinitionMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectDefinitionMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectDefinitionMetadataEntryKey(input string) (*ObjectDefinitionMetadataEntryKey, error) {
	vals := map[string]ObjectDefinitionMetadataEntryKey{
		"baseobjectname":               ObjectDefinitionMetadataEntryKeyBaseObjectName,
		"connectordatastoragerequired": ObjectDefinitionMetadataEntryKeyConnectorDataStorageRequired,
		"extensions":                   ObjectDefinitionMetadataEntryKeyExtensions,
		"issoftdeletionsupported":      ObjectDefinitionMetadataEntryKeyIsSoftDeletionSupported,
		"issynchronizeallsupported":    ObjectDefinitionMetadataEntryKeyIsSynchronizeAllSupported,
		"propertynameaccountenabled":   ObjectDefinitionMetadataEntryKeyPropertyNameAccountEnabled,
		"propertynamesoftdeleted":      ObjectDefinitionMetadataEntryKeyPropertyNameSoftDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectDefinitionMetadataEntryKey(input)
	return &out, nil
}

type ObjectMappingFlowTypes string

const (
	ObjectMappingFlowTypesAdd    ObjectMappingFlowTypes = "Add"
	ObjectMappingFlowTypesDelete ObjectMappingFlowTypes = "Delete"
	ObjectMappingFlowTypesNone   ObjectMappingFlowTypes = "None"
	ObjectMappingFlowTypesUpdate ObjectMappingFlowTypes = "Update"
)

func PossibleValuesForObjectMappingFlowTypes() []string {
	return []string{
		string(ObjectMappingFlowTypesAdd),
		string(ObjectMappingFlowTypesDelete),
		string(ObjectMappingFlowTypesNone),
		string(ObjectMappingFlowTypesUpdate),
	}
}

func (s *ObjectMappingFlowTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectMappingFlowTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectMappingFlowTypes(input string) (*ObjectMappingFlowTypes, error) {
	vals := map[string]ObjectMappingFlowTypes{
		"add":    ObjectMappingFlowTypesAdd,
		"delete": ObjectMappingFlowTypesDelete,
		"none":   ObjectMappingFlowTypesNone,
		"update": ObjectMappingFlowTypesUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectMappingFlowTypes(input)
	return &out, nil
}

type ObjectMappingMetadataEntryKey string

const (
	ObjectMappingMetadataEntryKeyDisableMonitoringForChanges ObjectMappingMetadataEntryKey = "DisableMonitoringForChanges"
	ObjectMappingMetadataEntryKeyDisposition                 ObjectMappingMetadataEntryKey = "Disposition"
	ObjectMappingMetadataEntryKeyEscrowBehavior              ObjectMappingMetadataEntryKey = "EscrowBehavior"
	ObjectMappingMetadataEntryKeyExcludeFromReporting        ObjectMappingMetadataEntryKey = "ExcludeFromReporting"
	ObjectMappingMetadataEntryKeyIsCustomerDefined           ObjectMappingMetadataEntryKey = "IsCustomerDefined"
	ObjectMappingMetadataEntryKeyOriginalJoiningProperty     ObjectMappingMetadataEntryKey = "OriginalJoiningProperty"
	ObjectMappingMetadataEntryKeyUnsynchronized              ObjectMappingMetadataEntryKey = "Unsynchronized"
)

func PossibleValuesForObjectMappingMetadataEntryKey() []string {
	return []string{
		string(ObjectMappingMetadataEntryKeyDisableMonitoringForChanges),
		string(ObjectMappingMetadataEntryKeyDisposition),
		string(ObjectMappingMetadataEntryKeyEscrowBehavior),
		string(ObjectMappingMetadataEntryKeyExcludeFromReporting),
		string(ObjectMappingMetadataEntryKeyIsCustomerDefined),
		string(ObjectMappingMetadataEntryKeyOriginalJoiningProperty),
		string(ObjectMappingMetadataEntryKeyUnsynchronized),
	}
}

func (s *ObjectMappingMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectMappingMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectMappingMetadataEntryKey(input string) (*ObjectMappingMetadataEntryKey, error) {
	vals := map[string]ObjectMappingMetadataEntryKey{
		"disablemonitoringforchanges": ObjectMappingMetadataEntryKeyDisableMonitoringForChanges,
		"disposition":                 ObjectMappingMetadataEntryKeyDisposition,
		"escrowbehavior":              ObjectMappingMetadataEntryKeyEscrowBehavior,
		"excludefromreporting":        ObjectMappingMetadataEntryKeyExcludeFromReporting,
		"iscustomerdefined":           ObjectMappingMetadataEntryKeyIsCustomerDefined,
		"originaljoiningproperty":     ObjectMappingMetadataEntryKeyOriginalJoiningProperty,
		"unsynchronized":              ObjectMappingMetadataEntryKeyUnsynchronized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectMappingMetadataEntryKey(input)
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

type OnlineMeetingAnonymizeIdentityForRoles string

const (
	OnlineMeetingAnonymizeIdentityForRolesAttendee    OnlineMeetingAnonymizeIdentityForRoles = "attendee"
	OnlineMeetingAnonymizeIdentityForRolesCoorganizer OnlineMeetingAnonymizeIdentityForRoles = "coorganizer"
	OnlineMeetingAnonymizeIdentityForRolesPresenter   OnlineMeetingAnonymizeIdentityForRoles = "presenter"
	OnlineMeetingAnonymizeIdentityForRolesProducer    OnlineMeetingAnonymizeIdentityForRoles = "producer"
)

func PossibleValuesForOnlineMeetingAnonymizeIdentityForRoles() []string {
	return []string{
		string(OnlineMeetingAnonymizeIdentityForRolesAttendee),
		string(OnlineMeetingAnonymizeIdentityForRolesCoorganizer),
		string(OnlineMeetingAnonymizeIdentityForRolesPresenter),
		string(OnlineMeetingAnonymizeIdentityForRolesProducer),
	}
}

func (s *OnlineMeetingAnonymizeIdentityForRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingAnonymizeIdentityForRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingAnonymizeIdentityForRoles(input string) (*OnlineMeetingAnonymizeIdentityForRoles, error) {
	vals := map[string]OnlineMeetingAnonymizeIdentityForRoles{
		"attendee":    OnlineMeetingAnonymizeIdentityForRolesAttendee,
		"coorganizer": OnlineMeetingAnonymizeIdentityForRolesCoorganizer,
		"presenter":   OnlineMeetingAnonymizeIdentityForRolesPresenter,
		"producer":    OnlineMeetingAnonymizeIdentityForRolesProducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAnonymizeIdentityForRoles(input)
	return &out, nil
}

type OnlineMeetingCapabilities string

const (
	OnlineMeetingCapabilitiesQuestionAndAnswer OnlineMeetingCapabilities = "questionAndAnswer"
)

func PossibleValuesForOnlineMeetingCapabilities() []string {
	return []string{
		string(OnlineMeetingCapabilitiesQuestionAndAnswer),
	}
}

func (s *OnlineMeetingCapabilities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingCapabilities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingCapabilities(input string) (*OnlineMeetingCapabilities, error) {
	vals := map[string]OnlineMeetingCapabilities{
		"questionandanswer": OnlineMeetingCapabilitiesQuestionAndAnswer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingCapabilities(input)
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

type OutlookTaskImportance string

const (
	OutlookTaskImportanceHigh   OutlookTaskImportance = "high"
	OutlookTaskImportanceLow    OutlookTaskImportance = "low"
	OutlookTaskImportanceNormal OutlookTaskImportance = "normal"
)

func PossibleValuesForOutlookTaskImportance() []string {
	return []string{
		string(OutlookTaskImportanceHigh),
		string(OutlookTaskImportanceLow),
		string(OutlookTaskImportanceNormal),
	}
}

func (s *OutlookTaskImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookTaskImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookTaskImportance(input string) (*OutlookTaskImportance, error) {
	vals := map[string]OutlookTaskImportance{
		"high":   OutlookTaskImportanceHigh,
		"low":    OutlookTaskImportanceLow,
		"normal": OutlookTaskImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskImportance(input)
	return &out, nil
}

type OutlookTaskSensitivity string

const (
	OutlookTaskSensitivityConfidential OutlookTaskSensitivity = "confidential"
	OutlookTaskSensitivityNormal       OutlookTaskSensitivity = "normal"
	OutlookTaskSensitivityPersonal     OutlookTaskSensitivity = "personal"
	OutlookTaskSensitivityPrivate      OutlookTaskSensitivity = "private"
)

func PossibleValuesForOutlookTaskSensitivity() []string {
	return []string{
		string(OutlookTaskSensitivityConfidential),
		string(OutlookTaskSensitivityNormal),
		string(OutlookTaskSensitivityPersonal),
		string(OutlookTaskSensitivityPrivate),
	}
}

func (s *OutlookTaskSensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookTaskSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookTaskSensitivity(input string) (*OutlookTaskSensitivity, error) {
	vals := map[string]OutlookTaskSensitivity{
		"confidential": OutlookTaskSensitivityConfidential,
		"normal":       OutlookTaskSensitivityNormal,
		"personal":     OutlookTaskSensitivityPersonal,
		"private":      OutlookTaskSensitivityPrivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskSensitivity(input)
	return &out, nil
}

type OutlookTaskStatus string

const (
	OutlookTaskStatusCompleted       OutlookTaskStatus = "completed"
	OutlookTaskStatusDeferred        OutlookTaskStatus = "deferred"
	OutlookTaskStatusInProgress      OutlookTaskStatus = "inProgress"
	OutlookTaskStatusNotStarted      OutlookTaskStatus = "notStarted"
	OutlookTaskStatusWaitingOnOthers OutlookTaskStatus = "waitingOnOthers"
)

func PossibleValuesForOutlookTaskStatus() []string {
	return []string{
		string(OutlookTaskStatusCompleted),
		string(OutlookTaskStatusDeferred),
		string(OutlookTaskStatusInProgress),
		string(OutlookTaskStatusNotStarted),
		string(OutlookTaskStatusWaitingOnOthers),
	}
}

func (s *OutlookTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookTaskStatus(input string) (*OutlookTaskStatus, error) {
	vals := map[string]OutlookTaskStatus{
		"completed":       OutlookTaskStatusCompleted,
		"deferred":        OutlookTaskStatusDeferred,
		"inprogress":      OutlookTaskStatusInProgress,
		"notstarted":      OutlookTaskStatusNotStarted,
		"waitingonothers": OutlookTaskStatusWaitingOnOthers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskStatus(input)
	return &out, nil
}

type PasswordCredentialConfigurationRestrictionType string

const (
	PasswordCredentialConfigurationRestrictionTypeCustomPasswordAddition PasswordCredentialConfigurationRestrictionType = "customPasswordAddition"
	PasswordCredentialConfigurationRestrictionTypePasswordAddition       PasswordCredentialConfigurationRestrictionType = "passwordAddition"
	PasswordCredentialConfigurationRestrictionTypePasswordLifetime       PasswordCredentialConfigurationRestrictionType = "passwordLifetime"
	PasswordCredentialConfigurationRestrictionTypeSymmetricKeyAddition   PasswordCredentialConfigurationRestrictionType = "symmetricKeyAddition"
	PasswordCredentialConfigurationRestrictionTypeSymmetricKeyLifetime   PasswordCredentialConfigurationRestrictionType = "symmetricKeyLifetime"
)

func PossibleValuesForPasswordCredentialConfigurationRestrictionType() []string {
	return []string{
		string(PasswordCredentialConfigurationRestrictionTypeCustomPasswordAddition),
		string(PasswordCredentialConfigurationRestrictionTypePasswordAddition),
		string(PasswordCredentialConfigurationRestrictionTypePasswordLifetime),
		string(PasswordCredentialConfigurationRestrictionTypeSymmetricKeyAddition),
		string(PasswordCredentialConfigurationRestrictionTypeSymmetricKeyLifetime),
	}
}

func (s *PasswordCredentialConfigurationRestrictionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePasswordCredentialConfigurationRestrictionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePasswordCredentialConfigurationRestrictionType(input string) (*PasswordCredentialConfigurationRestrictionType, error) {
	vals := map[string]PasswordCredentialConfigurationRestrictionType{
		"custompasswordaddition": PasswordCredentialConfigurationRestrictionTypeCustomPasswordAddition,
		"passwordaddition":       PasswordCredentialConfigurationRestrictionTypePasswordAddition,
		"passwordlifetime":       PasswordCredentialConfigurationRestrictionTypePasswordLifetime,
		"symmetrickeyaddition":   PasswordCredentialConfigurationRestrictionTypeSymmetricKeyAddition,
		"symmetrickeylifetime":   PasswordCredentialConfigurationRestrictionTypeSymmetricKeyLifetime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PasswordCredentialConfigurationRestrictionType(input)
	return &out, nil
}

type PersonAnnotationAllowedAudiences string

const (
	PersonAnnotationAllowedAudiencesContacts               PersonAnnotationAllowedAudiences = "contacts"
	PersonAnnotationAllowedAudiencesEveryone               PersonAnnotationAllowedAudiences = "everyone"
	PersonAnnotationAllowedAudiencesFamily                 PersonAnnotationAllowedAudiences = "family"
	PersonAnnotationAllowedAudiencesFederatedOrganizations PersonAnnotationAllowedAudiences = "federatedOrganizations"
	PersonAnnotationAllowedAudiencesGroupMembers           PersonAnnotationAllowedAudiences = "groupMembers"
	PersonAnnotationAllowedAudiencesMe                     PersonAnnotationAllowedAudiences = "me"
	PersonAnnotationAllowedAudiencesOrganization           PersonAnnotationAllowedAudiences = "organization"
)

func PossibleValuesForPersonAnnotationAllowedAudiences() []string {
	return []string{
		string(PersonAnnotationAllowedAudiencesContacts),
		string(PersonAnnotationAllowedAudiencesEveryone),
		string(PersonAnnotationAllowedAudiencesFamily),
		string(PersonAnnotationAllowedAudiencesFederatedOrganizations),
		string(PersonAnnotationAllowedAudiencesGroupMembers),
		string(PersonAnnotationAllowedAudiencesMe),
		string(PersonAnnotationAllowedAudiencesOrganization),
	}
}

func (s *PersonAnnotationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAnnotationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAnnotationAllowedAudiences(input string) (*PersonAnnotationAllowedAudiences, error) {
	vals := map[string]PersonAnnotationAllowedAudiences{
		"contacts":               PersonAnnotationAllowedAudiencesContacts,
		"everyone":               PersonAnnotationAllowedAudiencesEveryone,
		"family":                 PersonAnnotationAllowedAudiencesFamily,
		"federatedorganizations": PersonAnnotationAllowedAudiencesFederatedOrganizations,
		"groupmembers":           PersonAnnotationAllowedAudiencesGroupMembers,
		"me":                     PersonAnnotationAllowedAudiencesMe,
		"organization":           PersonAnnotationAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnotationAllowedAudiences(input)
	return &out, nil
}

type PersonAnnualEventAllowedAudiences string

const (
	PersonAnnualEventAllowedAudiencesContacts               PersonAnnualEventAllowedAudiences = "contacts"
	PersonAnnualEventAllowedAudiencesEveryone               PersonAnnualEventAllowedAudiences = "everyone"
	PersonAnnualEventAllowedAudiencesFamily                 PersonAnnualEventAllowedAudiences = "family"
	PersonAnnualEventAllowedAudiencesFederatedOrganizations PersonAnnualEventAllowedAudiences = "federatedOrganizations"
	PersonAnnualEventAllowedAudiencesGroupMembers           PersonAnnualEventAllowedAudiences = "groupMembers"
	PersonAnnualEventAllowedAudiencesMe                     PersonAnnualEventAllowedAudiences = "me"
	PersonAnnualEventAllowedAudiencesOrganization           PersonAnnualEventAllowedAudiences = "organization"
)

func PossibleValuesForPersonAnnualEventAllowedAudiences() []string {
	return []string{
		string(PersonAnnualEventAllowedAudiencesContacts),
		string(PersonAnnualEventAllowedAudiencesEveryone),
		string(PersonAnnualEventAllowedAudiencesFamily),
		string(PersonAnnualEventAllowedAudiencesFederatedOrganizations),
		string(PersonAnnualEventAllowedAudiencesGroupMembers),
		string(PersonAnnualEventAllowedAudiencesMe),
		string(PersonAnnualEventAllowedAudiencesOrganization),
	}
}

func (s *PersonAnnualEventAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAnnualEventAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAnnualEventAllowedAudiences(input string) (*PersonAnnualEventAllowedAudiences, error) {
	vals := map[string]PersonAnnualEventAllowedAudiences{
		"contacts":               PersonAnnualEventAllowedAudiencesContacts,
		"everyone":               PersonAnnualEventAllowedAudiencesEveryone,
		"family":                 PersonAnnualEventAllowedAudiencesFamily,
		"federatedorganizations": PersonAnnualEventAllowedAudiencesFederatedOrganizations,
		"groupmembers":           PersonAnnualEventAllowedAudiencesGroupMembers,
		"me":                     PersonAnnualEventAllowedAudiencesMe,
		"organization":           PersonAnnualEventAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnualEventAllowedAudiences(input)
	return &out, nil
}

type PersonAnnualEventType string

const (
	PersonAnnualEventTypeBirthday PersonAnnualEventType = "birthday"
	PersonAnnualEventTypeOther    PersonAnnualEventType = "other"
	PersonAnnualEventTypeWedding  PersonAnnualEventType = "wedding"
	PersonAnnualEventTypeWork     PersonAnnualEventType = "work"
)

func PossibleValuesForPersonAnnualEventType() []string {
	return []string{
		string(PersonAnnualEventTypeBirthday),
		string(PersonAnnualEventTypeOther),
		string(PersonAnnualEventTypeWedding),
		string(PersonAnnualEventTypeWork),
	}
}

func (s *PersonAnnualEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAnnualEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAnnualEventType(input string) (*PersonAnnualEventType, error) {
	vals := map[string]PersonAnnualEventType{
		"birthday": PersonAnnualEventTypeBirthday,
		"other":    PersonAnnualEventTypeOther,
		"wedding":  PersonAnnualEventTypeWedding,
		"work":     PersonAnnualEventTypeWork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnualEventType(input)
	return &out, nil
}

type PersonAwardAllowedAudiences string

const (
	PersonAwardAllowedAudiencesContacts               PersonAwardAllowedAudiences = "contacts"
	PersonAwardAllowedAudiencesEveryone               PersonAwardAllowedAudiences = "everyone"
	PersonAwardAllowedAudiencesFamily                 PersonAwardAllowedAudiences = "family"
	PersonAwardAllowedAudiencesFederatedOrganizations PersonAwardAllowedAudiences = "federatedOrganizations"
	PersonAwardAllowedAudiencesGroupMembers           PersonAwardAllowedAudiences = "groupMembers"
	PersonAwardAllowedAudiencesMe                     PersonAwardAllowedAudiences = "me"
	PersonAwardAllowedAudiencesOrganization           PersonAwardAllowedAudiences = "organization"
)

func PossibleValuesForPersonAwardAllowedAudiences() []string {
	return []string{
		string(PersonAwardAllowedAudiencesContacts),
		string(PersonAwardAllowedAudiencesEveryone),
		string(PersonAwardAllowedAudiencesFamily),
		string(PersonAwardAllowedAudiencesFederatedOrganizations),
		string(PersonAwardAllowedAudiencesGroupMembers),
		string(PersonAwardAllowedAudiencesMe),
		string(PersonAwardAllowedAudiencesOrganization),
	}
}

func (s *PersonAwardAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAwardAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAwardAllowedAudiences(input string) (*PersonAwardAllowedAudiences, error) {
	vals := map[string]PersonAwardAllowedAudiences{
		"contacts":               PersonAwardAllowedAudiencesContacts,
		"everyone":               PersonAwardAllowedAudiencesEveryone,
		"family":                 PersonAwardAllowedAudiencesFamily,
		"federatedorganizations": PersonAwardAllowedAudiencesFederatedOrganizations,
		"groupmembers":           PersonAwardAllowedAudiencesGroupMembers,
		"me":                     PersonAwardAllowedAudiencesMe,
		"organization":           PersonAwardAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAwardAllowedAudiences(input)
	return &out, nil
}

type PersonCertificationAllowedAudiences string

const (
	PersonCertificationAllowedAudiencesContacts               PersonCertificationAllowedAudiences = "contacts"
	PersonCertificationAllowedAudiencesEveryone               PersonCertificationAllowedAudiences = "everyone"
	PersonCertificationAllowedAudiencesFamily                 PersonCertificationAllowedAudiences = "family"
	PersonCertificationAllowedAudiencesFederatedOrganizations PersonCertificationAllowedAudiences = "federatedOrganizations"
	PersonCertificationAllowedAudiencesGroupMembers           PersonCertificationAllowedAudiences = "groupMembers"
	PersonCertificationAllowedAudiencesMe                     PersonCertificationAllowedAudiences = "me"
	PersonCertificationAllowedAudiencesOrganization           PersonCertificationAllowedAudiences = "organization"
)

func PossibleValuesForPersonCertificationAllowedAudiences() []string {
	return []string{
		string(PersonCertificationAllowedAudiencesContacts),
		string(PersonCertificationAllowedAudiencesEveryone),
		string(PersonCertificationAllowedAudiencesFamily),
		string(PersonCertificationAllowedAudiencesFederatedOrganizations),
		string(PersonCertificationAllowedAudiencesGroupMembers),
		string(PersonCertificationAllowedAudiencesMe),
		string(PersonCertificationAllowedAudiencesOrganization),
	}
}

func (s *PersonCertificationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonCertificationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonCertificationAllowedAudiences(input string) (*PersonCertificationAllowedAudiences, error) {
	vals := map[string]PersonCertificationAllowedAudiences{
		"contacts":               PersonCertificationAllowedAudiencesContacts,
		"everyone":               PersonCertificationAllowedAudiencesEveryone,
		"family":                 PersonCertificationAllowedAudiencesFamily,
		"federatedorganizations": PersonCertificationAllowedAudiencesFederatedOrganizations,
		"groupmembers":           PersonCertificationAllowedAudiencesGroupMembers,
		"me":                     PersonCertificationAllowedAudiencesMe,
		"organization":           PersonCertificationAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonCertificationAllowedAudiences(input)
	return &out, nil
}

type PersonInterestAllowedAudiences string

const (
	PersonInterestAllowedAudiencesContacts               PersonInterestAllowedAudiences = "contacts"
	PersonInterestAllowedAudiencesEveryone               PersonInterestAllowedAudiences = "everyone"
	PersonInterestAllowedAudiencesFamily                 PersonInterestAllowedAudiences = "family"
	PersonInterestAllowedAudiencesFederatedOrganizations PersonInterestAllowedAudiences = "federatedOrganizations"
	PersonInterestAllowedAudiencesGroupMembers           PersonInterestAllowedAudiences = "groupMembers"
	PersonInterestAllowedAudiencesMe                     PersonInterestAllowedAudiences = "me"
	PersonInterestAllowedAudiencesOrganization           PersonInterestAllowedAudiences = "organization"
)

func PossibleValuesForPersonInterestAllowedAudiences() []string {
	return []string{
		string(PersonInterestAllowedAudiencesContacts),
		string(PersonInterestAllowedAudiencesEveryone),
		string(PersonInterestAllowedAudiencesFamily),
		string(PersonInterestAllowedAudiencesFederatedOrganizations),
		string(PersonInterestAllowedAudiencesGroupMembers),
		string(PersonInterestAllowedAudiencesMe),
		string(PersonInterestAllowedAudiencesOrganization),
	}
}

func (s *PersonInterestAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonInterestAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonInterestAllowedAudiences(input string) (*PersonInterestAllowedAudiences, error) {
	vals := map[string]PersonInterestAllowedAudiences{
		"contacts":               PersonInterestAllowedAudiencesContacts,
		"everyone":               PersonInterestAllowedAudiencesEveryone,
		"family":                 PersonInterestAllowedAudiencesFamily,
		"federatedorganizations": PersonInterestAllowedAudiencesFederatedOrganizations,
		"groupmembers":           PersonInterestAllowedAudiencesGroupMembers,
		"me":                     PersonInterestAllowedAudiencesMe,
		"organization":           PersonInterestAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonInterestAllowedAudiences(input)
	return &out, nil
}

type PersonNameAllowedAudiences string

const (
	PersonNameAllowedAudiencesContacts               PersonNameAllowedAudiences = "contacts"
	PersonNameAllowedAudiencesEveryone               PersonNameAllowedAudiences = "everyone"
	PersonNameAllowedAudiencesFamily                 PersonNameAllowedAudiences = "family"
	PersonNameAllowedAudiencesFederatedOrganizations PersonNameAllowedAudiences = "federatedOrganizations"
	PersonNameAllowedAudiencesGroupMembers           PersonNameAllowedAudiences = "groupMembers"
	PersonNameAllowedAudiencesMe                     PersonNameAllowedAudiences = "me"
	PersonNameAllowedAudiencesOrganization           PersonNameAllowedAudiences = "organization"
)

func PossibleValuesForPersonNameAllowedAudiences() []string {
	return []string{
		string(PersonNameAllowedAudiencesContacts),
		string(PersonNameAllowedAudiencesEveryone),
		string(PersonNameAllowedAudiencesFamily),
		string(PersonNameAllowedAudiencesFederatedOrganizations),
		string(PersonNameAllowedAudiencesGroupMembers),
		string(PersonNameAllowedAudiencesMe),
		string(PersonNameAllowedAudiencesOrganization),
	}
}

func (s *PersonNameAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonNameAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonNameAllowedAudiences(input string) (*PersonNameAllowedAudiences, error) {
	vals := map[string]PersonNameAllowedAudiences{
		"contacts":               PersonNameAllowedAudiencesContacts,
		"everyone":               PersonNameAllowedAudiencesEveryone,
		"family":                 PersonNameAllowedAudiencesFamily,
		"federatedorganizations": PersonNameAllowedAudiencesFederatedOrganizations,
		"groupmembers":           PersonNameAllowedAudiencesGroupMembers,
		"me":                     PersonNameAllowedAudiencesMe,
		"organization":           PersonNameAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonNameAllowedAudiences(input)
	return &out, nil
}

type PersonWebsiteAllowedAudiences string

const (
	PersonWebsiteAllowedAudiencesContacts               PersonWebsiteAllowedAudiences = "contacts"
	PersonWebsiteAllowedAudiencesEveryone               PersonWebsiteAllowedAudiences = "everyone"
	PersonWebsiteAllowedAudiencesFamily                 PersonWebsiteAllowedAudiences = "family"
	PersonWebsiteAllowedAudiencesFederatedOrganizations PersonWebsiteAllowedAudiences = "federatedOrganizations"
	PersonWebsiteAllowedAudiencesGroupMembers           PersonWebsiteAllowedAudiences = "groupMembers"
	PersonWebsiteAllowedAudiencesMe                     PersonWebsiteAllowedAudiences = "me"
	PersonWebsiteAllowedAudiencesOrganization           PersonWebsiteAllowedAudiences = "organization"
)

func PossibleValuesForPersonWebsiteAllowedAudiences() []string {
	return []string{
		string(PersonWebsiteAllowedAudiencesContacts),
		string(PersonWebsiteAllowedAudiencesEveryone),
		string(PersonWebsiteAllowedAudiencesFamily),
		string(PersonWebsiteAllowedAudiencesFederatedOrganizations),
		string(PersonWebsiteAllowedAudiencesGroupMembers),
		string(PersonWebsiteAllowedAudiencesMe),
		string(PersonWebsiteAllowedAudiencesOrganization),
	}
}

func (s *PersonWebsiteAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonWebsiteAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonWebsiteAllowedAudiences(input string) (*PersonWebsiteAllowedAudiences, error) {
	vals := map[string]PersonWebsiteAllowedAudiences{
		"contacts":               PersonWebsiteAllowedAudiencesContacts,
		"everyone":               PersonWebsiteAllowedAudiencesEveryone,
		"family":                 PersonWebsiteAllowedAudiencesFamily,
		"federatedorganizations": PersonWebsiteAllowedAudiencesFederatedOrganizations,
		"groupmembers":           PersonWebsiteAllowedAudiencesGroupMembers,
		"me":                     PersonWebsiteAllowedAudiencesMe,
		"organization":           PersonWebsiteAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonWebsiteAllowedAudiences(input)
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

type PlannerBucketCreationCreationSourceKind string

const (
	PlannerBucketCreationCreationSourceKindExternal    PlannerBucketCreationCreationSourceKind = "external"
	PlannerBucketCreationCreationSourceKindNone        PlannerBucketCreationCreationSourceKind = "none"
	PlannerBucketCreationCreationSourceKindPublication PlannerBucketCreationCreationSourceKind = "publication"
)

func PossibleValuesForPlannerBucketCreationCreationSourceKind() []string {
	return []string{
		string(PlannerBucketCreationCreationSourceKindExternal),
		string(PlannerBucketCreationCreationSourceKindNone),
		string(PlannerBucketCreationCreationSourceKindPublication),
	}
}

func (s *PlannerBucketCreationCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerBucketCreationCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerBucketCreationCreationSourceKind(input string) (*PlannerBucketCreationCreationSourceKind, error) {
	vals := map[string]PlannerBucketCreationCreationSourceKind{
		"external":    PlannerBucketCreationCreationSourceKindExternal,
		"none":        PlannerBucketCreationCreationSourceKindNone,
		"publication": PlannerBucketCreationCreationSourceKindPublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerBucketCreationCreationSourceKind(input)
	return &out, nil
}

type PlannerPlanContainerType string

const (
	PlannerPlanContainerTypeDriveItem PlannerPlanContainerType = "driveItem"
	PlannerPlanContainerTypeGroup     PlannerPlanContainerType = "group"
	PlannerPlanContainerTypeProject   PlannerPlanContainerType = "project"
	PlannerPlanContainerTypeRoster    PlannerPlanContainerType = "roster"
	PlannerPlanContainerTypeUser      PlannerPlanContainerType = "user"
)

func PossibleValuesForPlannerPlanContainerType() []string {
	return []string{
		string(PlannerPlanContainerTypeDriveItem),
		string(PlannerPlanContainerTypeGroup),
		string(PlannerPlanContainerTypeProject),
		string(PlannerPlanContainerTypeRoster),
		string(PlannerPlanContainerTypeUser),
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
		"driveitem": PlannerPlanContainerTypeDriveItem,
		"group":     PlannerPlanContainerTypeGroup,
		"project":   PlannerPlanContainerTypeProject,
		"roster":    PlannerPlanContainerTypeRoster,
		"user":      PlannerPlanContainerTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContainerType(input)
	return &out, nil
}

type PlannerPlanCreationCreationSourceKind string

const (
	PlannerPlanCreationCreationSourceKindExternal    PlannerPlanCreationCreationSourceKind = "external"
	PlannerPlanCreationCreationSourceKindNone        PlannerPlanCreationCreationSourceKind = "none"
	PlannerPlanCreationCreationSourceKindPublication PlannerPlanCreationCreationSourceKind = "publication"
)

func PossibleValuesForPlannerPlanCreationCreationSourceKind() []string {
	return []string{
		string(PlannerPlanCreationCreationSourceKindExternal),
		string(PlannerPlanCreationCreationSourceKindNone),
		string(PlannerPlanCreationCreationSourceKindPublication),
	}
}

func (s *PlannerPlanCreationCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanCreationCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanCreationCreationSourceKind(input string) (*PlannerPlanCreationCreationSourceKind, error) {
	vals := map[string]PlannerPlanCreationCreationSourceKind{
		"external":    PlannerPlanCreationCreationSourceKindExternal,
		"none":        PlannerPlanCreationCreationSourceKindNone,
		"publication": PlannerPlanCreationCreationSourceKindPublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanCreationCreationSourceKind(input)
	return &out, nil
}

type PlannerSharedWithContainerAccessLevel string

const (
	PlannerSharedWithContainerAccessLevelFullAccess      PlannerSharedWithContainerAccessLevel = "fullAccess"
	PlannerSharedWithContainerAccessLevelReadAccess      PlannerSharedWithContainerAccessLevel = "readAccess"
	PlannerSharedWithContainerAccessLevelReadWriteAccess PlannerSharedWithContainerAccessLevel = "readWriteAccess"
)

func PossibleValuesForPlannerSharedWithContainerAccessLevel() []string {
	return []string{
		string(PlannerSharedWithContainerAccessLevelFullAccess),
		string(PlannerSharedWithContainerAccessLevelReadAccess),
		string(PlannerSharedWithContainerAccessLevelReadWriteAccess),
	}
}

func (s *PlannerSharedWithContainerAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerSharedWithContainerAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerSharedWithContainerAccessLevel(input string) (*PlannerSharedWithContainerAccessLevel, error) {
	vals := map[string]PlannerSharedWithContainerAccessLevel{
		"fullaccess":      PlannerSharedWithContainerAccessLevelFullAccess,
		"readaccess":      PlannerSharedWithContainerAccessLevelReadAccess,
		"readwriteaccess": PlannerSharedWithContainerAccessLevelReadWriteAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerSharedWithContainerAccessLevel(input)
	return &out, nil
}

type PlannerSharedWithContainerType string

const (
	PlannerSharedWithContainerTypeDriveItem PlannerSharedWithContainerType = "driveItem"
	PlannerSharedWithContainerTypeGroup     PlannerSharedWithContainerType = "group"
	PlannerSharedWithContainerTypeProject   PlannerSharedWithContainerType = "project"
	PlannerSharedWithContainerTypeRoster    PlannerSharedWithContainerType = "roster"
	PlannerSharedWithContainerTypeUser      PlannerSharedWithContainerType = "user"
)

func PossibleValuesForPlannerSharedWithContainerType() []string {
	return []string{
		string(PlannerSharedWithContainerTypeDriveItem),
		string(PlannerSharedWithContainerTypeGroup),
		string(PlannerSharedWithContainerTypeProject),
		string(PlannerSharedWithContainerTypeRoster),
		string(PlannerSharedWithContainerTypeUser),
	}
}

func (s *PlannerSharedWithContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerSharedWithContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerSharedWithContainerType(input string) (*PlannerSharedWithContainerType, error) {
	vals := map[string]PlannerSharedWithContainerType{
		"driveitem": PlannerSharedWithContainerTypeDriveItem,
		"group":     PlannerSharedWithContainerTypeGroup,
		"project":   PlannerSharedWithContainerTypeProject,
		"roster":    PlannerSharedWithContainerTypeRoster,
		"user":      PlannerSharedWithContainerTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerSharedWithContainerType(input)
	return &out, nil
}

type PlannerTaskCreationCreationSourceKind string

const (
	PlannerTaskCreationCreationSourceKindExternal    PlannerTaskCreationCreationSourceKind = "external"
	PlannerTaskCreationCreationSourceKindNone        PlannerTaskCreationCreationSourceKind = "none"
	PlannerTaskCreationCreationSourceKindPublication PlannerTaskCreationCreationSourceKind = "publication"
)

func PossibleValuesForPlannerTaskCreationCreationSourceKind() []string {
	return []string{
		string(PlannerTaskCreationCreationSourceKindExternal),
		string(PlannerTaskCreationCreationSourceKindNone),
		string(PlannerTaskCreationCreationSourceKindPublication),
	}
}

func (s *PlannerTaskCreationCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskCreationCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskCreationCreationSourceKind(input string) (*PlannerTaskCreationCreationSourceKind, error) {
	vals := map[string]PlannerTaskCreationCreationSourceKind{
		"external":    PlannerTaskCreationCreationSourceKindExternal,
		"none":        PlannerTaskCreationCreationSourceKindNone,
		"publication": PlannerTaskCreationCreationSourceKindPublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskCreationCreationSourceKind(input)
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

type PlannerTaskSpecifiedCompletionRequirements string

const (
	PlannerTaskSpecifiedCompletionRequirementsChecklistCompletion PlannerTaskSpecifiedCompletionRequirements = "checklistCompletion"
	PlannerTaskSpecifiedCompletionRequirementsNone                PlannerTaskSpecifiedCompletionRequirements = "none"
)

func PossibleValuesForPlannerTaskSpecifiedCompletionRequirements() []string {
	return []string{
		string(PlannerTaskSpecifiedCompletionRequirementsChecklistCompletion),
		string(PlannerTaskSpecifiedCompletionRequirementsNone),
	}
}

func (s *PlannerTaskSpecifiedCompletionRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskSpecifiedCompletionRequirements(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskSpecifiedCompletionRequirements(input string) (*PlannerTaskSpecifiedCompletionRequirements, error) {
	vals := map[string]PlannerTaskSpecifiedCompletionRequirements{
		"checklistcompletion": PlannerTaskSpecifiedCompletionRequirementsChecklistCompletion,
		"none":                PlannerTaskSpecifiedCompletionRequirementsNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskSpecifiedCompletionRequirements(input)
	return &out, nil
}

type PlannerTeamsPublicationInfoCreationSourceKind string

const (
	PlannerTeamsPublicationInfoCreationSourceKindExternal    PlannerTeamsPublicationInfoCreationSourceKind = "external"
	PlannerTeamsPublicationInfoCreationSourceKindNone        PlannerTeamsPublicationInfoCreationSourceKind = "none"
	PlannerTeamsPublicationInfoCreationSourceKindPublication PlannerTeamsPublicationInfoCreationSourceKind = "publication"
)

func PossibleValuesForPlannerTeamsPublicationInfoCreationSourceKind() []string {
	return []string{
		string(PlannerTeamsPublicationInfoCreationSourceKindExternal),
		string(PlannerTeamsPublicationInfoCreationSourceKindNone),
		string(PlannerTeamsPublicationInfoCreationSourceKindPublication),
	}
}

func (s *PlannerTeamsPublicationInfoCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTeamsPublicationInfoCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTeamsPublicationInfoCreationSourceKind(input string) (*PlannerTeamsPublicationInfoCreationSourceKind, error) {
	vals := map[string]PlannerTeamsPublicationInfoCreationSourceKind{
		"external":    PlannerTeamsPublicationInfoCreationSourceKindExternal,
		"none":        PlannerTeamsPublicationInfoCreationSourceKindNone,
		"publication": PlannerTeamsPublicationInfoCreationSourceKindPublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTeamsPublicationInfoCreationSourceKind(input)
	return &out, nil
}

type PlatformCredentialAuthenticationMethodKeyStrength string

const (
	PlatformCredentialAuthenticationMethodKeyStrengthNormal  PlatformCredentialAuthenticationMethodKeyStrength = "normal"
	PlatformCredentialAuthenticationMethodKeyStrengthUnknown PlatformCredentialAuthenticationMethodKeyStrength = "unknown"
	PlatformCredentialAuthenticationMethodKeyStrengthWeak    PlatformCredentialAuthenticationMethodKeyStrength = "weak"
)

func PossibleValuesForPlatformCredentialAuthenticationMethodKeyStrength() []string {
	return []string{
		string(PlatformCredentialAuthenticationMethodKeyStrengthNormal),
		string(PlatformCredentialAuthenticationMethodKeyStrengthUnknown),
		string(PlatformCredentialAuthenticationMethodKeyStrengthWeak),
	}
}

func (s *PlatformCredentialAuthenticationMethodKeyStrength) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlatformCredentialAuthenticationMethodKeyStrength(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlatformCredentialAuthenticationMethodKeyStrength(input string) (*PlatformCredentialAuthenticationMethodKeyStrength, error) {
	vals := map[string]PlatformCredentialAuthenticationMethodKeyStrength{
		"normal":  PlatformCredentialAuthenticationMethodKeyStrengthNormal,
		"unknown": PlatformCredentialAuthenticationMethodKeyStrengthUnknown,
		"weak":    PlatformCredentialAuthenticationMethodKeyStrengthWeak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformCredentialAuthenticationMethodKeyStrength(input)
	return &out, nil
}

type PlatformCredentialAuthenticationMethodPlatform string

const (
	PlatformCredentialAuthenticationMethodPlatformAndroid PlatformCredentialAuthenticationMethodPlatform = "android"
	PlatformCredentialAuthenticationMethodPlatformIOS     PlatformCredentialAuthenticationMethodPlatform = "iOS"
	PlatformCredentialAuthenticationMethodPlatformLinux   PlatformCredentialAuthenticationMethodPlatform = "linux"
	PlatformCredentialAuthenticationMethodPlatformMacOS   PlatformCredentialAuthenticationMethodPlatform = "macOS"
	PlatformCredentialAuthenticationMethodPlatformUnknown PlatformCredentialAuthenticationMethodPlatform = "unknown"
	PlatformCredentialAuthenticationMethodPlatformWindows PlatformCredentialAuthenticationMethodPlatform = "windows"
)

func PossibleValuesForPlatformCredentialAuthenticationMethodPlatform() []string {
	return []string{
		string(PlatformCredentialAuthenticationMethodPlatformAndroid),
		string(PlatformCredentialAuthenticationMethodPlatformIOS),
		string(PlatformCredentialAuthenticationMethodPlatformLinux),
		string(PlatformCredentialAuthenticationMethodPlatformMacOS),
		string(PlatformCredentialAuthenticationMethodPlatformUnknown),
		string(PlatformCredentialAuthenticationMethodPlatformWindows),
	}
}

func (s *PlatformCredentialAuthenticationMethodPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlatformCredentialAuthenticationMethodPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlatformCredentialAuthenticationMethodPlatform(input string) (*PlatformCredentialAuthenticationMethodPlatform, error) {
	vals := map[string]PlatformCredentialAuthenticationMethodPlatform{
		"android": PlatformCredentialAuthenticationMethodPlatformAndroid,
		"ios":     PlatformCredentialAuthenticationMethodPlatformIOS,
		"linux":   PlatformCredentialAuthenticationMethodPlatformLinux,
		"macos":   PlatformCredentialAuthenticationMethodPlatformMacOS,
		"unknown": PlatformCredentialAuthenticationMethodPlatformUnknown,
		"windows": PlatformCredentialAuthenticationMethodPlatformWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformCredentialAuthenticationMethodPlatform(input)
	return &out, nil
}

type PostImportance string

const (
	PostImportanceHigh   PostImportance = "high"
	PostImportanceLow    PostImportance = "low"
	PostImportanceNormal PostImportance = "normal"
)

func PossibleValuesForPostImportance() []string {
	return []string{
		string(PostImportanceHigh),
		string(PostImportanceLow),
		string(PostImportanceNormal),
	}
}

func (s *PostImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePostImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePostImportance(input string) (*PostImportance, error) {
	vals := map[string]PostImportance{
		"high":   PostImportanceHigh,
		"low":    PostImportanceLow,
		"normal": PostImportanceNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PostImportance(input)
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

type PrintJobStatusProcessingState string

const (
	PrintJobStatusProcessingStateAborted    PrintJobStatusProcessingState = "aborted"
	PrintJobStatusProcessingStateCanceled   PrintJobStatusProcessingState = "canceled"
	PrintJobStatusProcessingStateCompleted  PrintJobStatusProcessingState = "completed"
	PrintJobStatusProcessingStatePaused     PrintJobStatusProcessingState = "paused"
	PrintJobStatusProcessingStatePending    PrintJobStatusProcessingState = "pending"
	PrintJobStatusProcessingStateProcessing PrintJobStatusProcessingState = "processing"
	PrintJobStatusProcessingStateStopped    PrintJobStatusProcessingState = "stopped"
	PrintJobStatusProcessingStateUnknown    PrintJobStatusProcessingState = "unknown"
)

func PossibleValuesForPrintJobStatusProcessingState() []string {
	return []string{
		string(PrintJobStatusProcessingStateAborted),
		string(PrintJobStatusProcessingStateCanceled),
		string(PrintJobStatusProcessingStateCompleted),
		string(PrintJobStatusProcessingStatePaused),
		string(PrintJobStatusProcessingStatePending),
		string(PrintJobStatusProcessingStateProcessing),
		string(PrintJobStatusProcessingStateStopped),
		string(PrintJobStatusProcessingStateUnknown),
	}
}

func (s *PrintJobStatusProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobStatusProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobStatusProcessingState(input string) (*PrintJobStatusProcessingState, error) {
	vals := map[string]PrintJobStatusProcessingState{
		"aborted":    PrintJobStatusProcessingStateAborted,
		"canceled":   PrintJobStatusProcessingStateCanceled,
		"completed":  PrintJobStatusProcessingStateCompleted,
		"paused":     PrintJobStatusProcessingStatePaused,
		"pending":    PrintJobStatusProcessingStatePending,
		"processing": PrintJobStatusProcessingStateProcessing,
		"stopped":    PrintJobStatusProcessingStateStopped,
		"unknown":    PrintJobStatusProcessingStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusProcessingState(input)
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

type PrinterCapabilitiesFeedDirections string

const (
	PrinterCapabilitiesFeedDirectionsLongEdgeFirst  PrinterCapabilitiesFeedDirections = "longEdgeFirst"
	PrinterCapabilitiesFeedDirectionsShortEdgeFirst PrinterCapabilitiesFeedDirections = "shortEdgeFirst"
)

func PossibleValuesForPrinterCapabilitiesFeedDirections() []string {
	return []string{
		string(PrinterCapabilitiesFeedDirectionsLongEdgeFirst),
		string(PrinterCapabilitiesFeedDirectionsShortEdgeFirst),
	}
}

func (s *PrinterCapabilitiesFeedDirections) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesFeedDirections(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesFeedDirections(input string) (*PrinterCapabilitiesFeedDirections, error) {
	vals := map[string]PrinterCapabilitiesFeedDirections{
		"longedgefirst":  PrinterCapabilitiesFeedDirectionsLongEdgeFirst,
		"shortedgefirst": PrinterCapabilitiesFeedDirectionsShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFeedDirections(input)
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

type PrinterCapabilitiesSupportedColorConfigurations string

const (
	PrinterCapabilitiesSupportedColorConfigurationsAuto          PrinterCapabilitiesSupportedColorConfigurations = "auto"
	PrinterCapabilitiesSupportedColorConfigurationsBlackAndWhite PrinterCapabilitiesSupportedColorConfigurations = "blackAndWhite"
	PrinterCapabilitiesSupportedColorConfigurationsColor         PrinterCapabilitiesSupportedColorConfigurations = "color"
	PrinterCapabilitiesSupportedColorConfigurationsGrayscale     PrinterCapabilitiesSupportedColorConfigurations = "grayscale"
)

func PossibleValuesForPrinterCapabilitiesSupportedColorConfigurations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedColorConfigurationsAuto),
		string(PrinterCapabilitiesSupportedColorConfigurationsBlackAndWhite),
		string(PrinterCapabilitiesSupportedColorConfigurationsColor),
		string(PrinterCapabilitiesSupportedColorConfigurationsGrayscale),
	}
}

func (s *PrinterCapabilitiesSupportedColorConfigurations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedColorConfigurations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedColorConfigurations(input string) (*PrinterCapabilitiesSupportedColorConfigurations, error) {
	vals := map[string]PrinterCapabilitiesSupportedColorConfigurations{
		"auto":          PrinterCapabilitiesSupportedColorConfigurationsAuto,
		"blackandwhite": PrinterCapabilitiesSupportedColorConfigurationsBlackAndWhite,
		"color":         PrinterCapabilitiesSupportedColorConfigurationsColor,
		"grayscale":     PrinterCapabilitiesSupportedColorConfigurationsGrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedColorConfigurations(input)
	return &out, nil
}

type PrinterCapabilitiesSupportedDuplexConfigurations string

const (
	PrinterCapabilitiesSupportedDuplexConfigurationsOneSided          PrinterCapabilitiesSupportedDuplexConfigurations = "oneSided"
	PrinterCapabilitiesSupportedDuplexConfigurationsTwoSidedLongEdge  PrinterCapabilitiesSupportedDuplexConfigurations = "twoSidedLongEdge"
	PrinterCapabilitiesSupportedDuplexConfigurationsTwoSidedShortEdge PrinterCapabilitiesSupportedDuplexConfigurations = "twoSidedShortEdge"
)

func PossibleValuesForPrinterCapabilitiesSupportedDuplexConfigurations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedDuplexConfigurationsOneSided),
		string(PrinterCapabilitiesSupportedDuplexConfigurationsTwoSidedLongEdge),
		string(PrinterCapabilitiesSupportedDuplexConfigurationsTwoSidedShortEdge),
	}
}

func (s *PrinterCapabilitiesSupportedDuplexConfigurations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedDuplexConfigurations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedDuplexConfigurations(input string) (*PrinterCapabilitiesSupportedDuplexConfigurations, error) {
	vals := map[string]PrinterCapabilitiesSupportedDuplexConfigurations{
		"onesided":          PrinterCapabilitiesSupportedDuplexConfigurationsOneSided,
		"twosidedlongedge":  PrinterCapabilitiesSupportedDuplexConfigurationsTwoSidedLongEdge,
		"twosidedshortedge": PrinterCapabilitiesSupportedDuplexConfigurationsTwoSidedShortEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedDuplexConfigurations(input)
	return &out, nil
}

type PrinterCapabilitiesSupportedFinishings string

const (
	PrinterCapabilitiesSupportedFinishingsBind              PrinterCapabilitiesSupportedFinishings = "bind"
	PrinterCapabilitiesSupportedFinishingsCover             PrinterCapabilitiesSupportedFinishings = "cover"
	PrinterCapabilitiesSupportedFinishingsNone              PrinterCapabilitiesSupportedFinishings = "none"
	PrinterCapabilitiesSupportedFinishingsPunch             PrinterCapabilitiesSupportedFinishings = "punch"
	PrinterCapabilitiesSupportedFinishingsSaddleStitch      PrinterCapabilitiesSupportedFinishings = "saddleStitch"
	PrinterCapabilitiesSupportedFinishingsStaple            PrinterCapabilitiesSupportedFinishings = "staple"
	PrinterCapabilitiesSupportedFinishingsStapleBottomLeft  PrinterCapabilitiesSupportedFinishings = "stapleBottomLeft"
	PrinterCapabilitiesSupportedFinishingsStapleBottomRight PrinterCapabilitiesSupportedFinishings = "stapleBottomRight"
	PrinterCapabilitiesSupportedFinishingsStapleDualBottom  PrinterCapabilitiesSupportedFinishings = "stapleDualBottom"
	PrinterCapabilitiesSupportedFinishingsStapleDualLeft    PrinterCapabilitiesSupportedFinishings = "stapleDualLeft"
	PrinterCapabilitiesSupportedFinishingsStapleDualRight   PrinterCapabilitiesSupportedFinishings = "stapleDualRight"
	PrinterCapabilitiesSupportedFinishingsStapleDualTop     PrinterCapabilitiesSupportedFinishings = "stapleDualTop"
	PrinterCapabilitiesSupportedFinishingsStapleTopLeft     PrinterCapabilitiesSupportedFinishings = "stapleTopLeft"
	PrinterCapabilitiesSupportedFinishingsStapleTopRight    PrinterCapabilitiesSupportedFinishings = "stapleTopRight"
	PrinterCapabilitiesSupportedFinishingsStitchBottomEdge  PrinterCapabilitiesSupportedFinishings = "stitchBottomEdge"
	PrinterCapabilitiesSupportedFinishingsStitchEdge        PrinterCapabilitiesSupportedFinishings = "stitchEdge"
	PrinterCapabilitiesSupportedFinishingsStitchLeftEdge    PrinterCapabilitiesSupportedFinishings = "stitchLeftEdge"
	PrinterCapabilitiesSupportedFinishingsStitchRightEdge   PrinterCapabilitiesSupportedFinishings = "stitchRightEdge"
	PrinterCapabilitiesSupportedFinishingsStitchTopEdge     PrinterCapabilitiesSupportedFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterCapabilitiesSupportedFinishings() []string {
	return []string{
		string(PrinterCapabilitiesSupportedFinishingsBind),
		string(PrinterCapabilitiesSupportedFinishingsCover),
		string(PrinterCapabilitiesSupportedFinishingsNone),
		string(PrinterCapabilitiesSupportedFinishingsPunch),
		string(PrinterCapabilitiesSupportedFinishingsSaddleStitch),
		string(PrinterCapabilitiesSupportedFinishingsStaple),
		string(PrinterCapabilitiesSupportedFinishingsStapleBottomLeft),
		string(PrinterCapabilitiesSupportedFinishingsStapleBottomRight),
		string(PrinterCapabilitiesSupportedFinishingsStapleDualBottom),
		string(PrinterCapabilitiesSupportedFinishingsStapleDualLeft),
		string(PrinterCapabilitiesSupportedFinishingsStapleDualRight),
		string(PrinterCapabilitiesSupportedFinishingsStapleDualTop),
		string(PrinterCapabilitiesSupportedFinishingsStapleTopLeft),
		string(PrinterCapabilitiesSupportedFinishingsStapleTopRight),
		string(PrinterCapabilitiesSupportedFinishingsStitchBottomEdge),
		string(PrinterCapabilitiesSupportedFinishingsStitchEdge),
		string(PrinterCapabilitiesSupportedFinishingsStitchLeftEdge),
		string(PrinterCapabilitiesSupportedFinishingsStitchRightEdge),
		string(PrinterCapabilitiesSupportedFinishingsStitchTopEdge),
	}
}

func (s *PrinterCapabilitiesSupportedFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedFinishings(input string) (*PrinterCapabilitiesSupportedFinishings, error) {
	vals := map[string]PrinterCapabilitiesSupportedFinishings{
		"bind":              PrinterCapabilitiesSupportedFinishingsBind,
		"cover":             PrinterCapabilitiesSupportedFinishingsCover,
		"none":              PrinterCapabilitiesSupportedFinishingsNone,
		"punch":             PrinterCapabilitiesSupportedFinishingsPunch,
		"saddlestitch":      PrinterCapabilitiesSupportedFinishingsSaddleStitch,
		"staple":            PrinterCapabilitiesSupportedFinishingsStaple,
		"staplebottomleft":  PrinterCapabilitiesSupportedFinishingsStapleBottomLeft,
		"staplebottomright": PrinterCapabilitiesSupportedFinishingsStapleBottomRight,
		"stapledualbottom":  PrinterCapabilitiesSupportedFinishingsStapleDualBottom,
		"stapledualleft":    PrinterCapabilitiesSupportedFinishingsStapleDualLeft,
		"stapledualright":   PrinterCapabilitiesSupportedFinishingsStapleDualRight,
		"stapledualtop":     PrinterCapabilitiesSupportedFinishingsStapleDualTop,
		"stapletopleft":     PrinterCapabilitiesSupportedFinishingsStapleTopLeft,
		"stapletopright":    PrinterCapabilitiesSupportedFinishingsStapleTopRight,
		"stitchbottomedge":  PrinterCapabilitiesSupportedFinishingsStitchBottomEdge,
		"stitchedge":        PrinterCapabilitiesSupportedFinishingsStitchEdge,
		"stitchleftedge":    PrinterCapabilitiesSupportedFinishingsStitchLeftEdge,
		"stitchrightedge":   PrinterCapabilitiesSupportedFinishingsStitchRightEdge,
		"stitchtopedge":     PrinterCapabilitiesSupportedFinishingsStitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedFinishings(input)
	return &out, nil
}

type PrinterCapabilitiesSupportedMediaTypes string

const (
	PrinterCapabilitiesSupportedMediaTypesContinuous      PrinterCapabilitiesSupportedMediaTypes = "continuous"
	PrinterCapabilitiesSupportedMediaTypesContinuousLong  PrinterCapabilitiesSupportedMediaTypes = "continuousLong"
	PrinterCapabilitiesSupportedMediaTypesContinuousShort PrinterCapabilitiesSupportedMediaTypes = "continuousShort"
	PrinterCapabilitiesSupportedMediaTypesEnvelope        PrinterCapabilitiesSupportedMediaTypes = "envelope"
	PrinterCapabilitiesSupportedMediaTypesEnvelopePlain   PrinterCapabilitiesSupportedMediaTypes = "envelopePlain"
	PrinterCapabilitiesSupportedMediaTypesEnvelopeWindow  PrinterCapabilitiesSupportedMediaTypes = "envelopeWindow"
	PrinterCapabilitiesSupportedMediaTypesLabels          PrinterCapabilitiesSupportedMediaTypes = "labels"
	PrinterCapabilitiesSupportedMediaTypesMultiLayer      PrinterCapabilitiesSupportedMediaTypes = "multiLayer"
	PrinterCapabilitiesSupportedMediaTypesMultiPartForm   PrinterCapabilitiesSupportedMediaTypes = "multiPartForm"
	PrinterCapabilitiesSupportedMediaTypesScreen          PrinterCapabilitiesSupportedMediaTypes = "screen"
	PrinterCapabilitiesSupportedMediaTypesScreenPaged     PrinterCapabilitiesSupportedMediaTypes = "screenPaged"
	PrinterCapabilitiesSupportedMediaTypesStationery      PrinterCapabilitiesSupportedMediaTypes = "stationery"
	PrinterCapabilitiesSupportedMediaTypesTransparency    PrinterCapabilitiesSupportedMediaTypes = "transparency"
)

func PossibleValuesForPrinterCapabilitiesSupportedMediaTypes() []string {
	return []string{
		string(PrinterCapabilitiesSupportedMediaTypesContinuous),
		string(PrinterCapabilitiesSupportedMediaTypesContinuousLong),
		string(PrinterCapabilitiesSupportedMediaTypesContinuousShort),
		string(PrinterCapabilitiesSupportedMediaTypesEnvelope),
		string(PrinterCapabilitiesSupportedMediaTypesEnvelopePlain),
		string(PrinterCapabilitiesSupportedMediaTypesEnvelopeWindow),
		string(PrinterCapabilitiesSupportedMediaTypesLabels),
		string(PrinterCapabilitiesSupportedMediaTypesMultiLayer),
		string(PrinterCapabilitiesSupportedMediaTypesMultiPartForm),
		string(PrinterCapabilitiesSupportedMediaTypesScreen),
		string(PrinterCapabilitiesSupportedMediaTypesScreenPaged),
		string(PrinterCapabilitiesSupportedMediaTypesStationery),
		string(PrinterCapabilitiesSupportedMediaTypesTransparency),
	}
}

func (s *PrinterCapabilitiesSupportedMediaTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedMediaTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedMediaTypes(input string) (*PrinterCapabilitiesSupportedMediaTypes, error) {
	vals := map[string]PrinterCapabilitiesSupportedMediaTypes{
		"continuous":      PrinterCapabilitiesSupportedMediaTypesContinuous,
		"continuouslong":  PrinterCapabilitiesSupportedMediaTypesContinuousLong,
		"continuousshort": PrinterCapabilitiesSupportedMediaTypesContinuousShort,
		"envelope":        PrinterCapabilitiesSupportedMediaTypesEnvelope,
		"envelopeplain":   PrinterCapabilitiesSupportedMediaTypesEnvelopePlain,
		"envelopewindow":  PrinterCapabilitiesSupportedMediaTypesEnvelopeWindow,
		"labels":          PrinterCapabilitiesSupportedMediaTypesLabels,
		"multilayer":      PrinterCapabilitiesSupportedMediaTypesMultiLayer,
		"multipartform":   PrinterCapabilitiesSupportedMediaTypesMultiPartForm,
		"screen":          PrinterCapabilitiesSupportedMediaTypesScreen,
		"screenpaged":     PrinterCapabilitiesSupportedMediaTypesScreenPaged,
		"stationery":      PrinterCapabilitiesSupportedMediaTypesStationery,
		"transparency":    PrinterCapabilitiesSupportedMediaTypesTransparency,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedMediaTypes(input)
	return &out, nil
}

type PrinterCapabilitiesSupportedOrientations string

const (
	PrinterCapabilitiesSupportedOrientationsLandscape        PrinterCapabilitiesSupportedOrientations = "landscape"
	PrinterCapabilitiesSupportedOrientationsPortrait         PrinterCapabilitiesSupportedOrientations = "portrait"
	PrinterCapabilitiesSupportedOrientationsReverseLandscape PrinterCapabilitiesSupportedOrientations = "reverseLandscape"
	PrinterCapabilitiesSupportedOrientationsReversePortrait  PrinterCapabilitiesSupportedOrientations = "reversePortrait"
)

func PossibleValuesForPrinterCapabilitiesSupportedOrientations() []string {
	return []string{
		string(PrinterCapabilitiesSupportedOrientationsLandscape),
		string(PrinterCapabilitiesSupportedOrientationsPortrait),
		string(PrinterCapabilitiesSupportedOrientationsReverseLandscape),
		string(PrinterCapabilitiesSupportedOrientationsReversePortrait),
	}
}

func (s *PrinterCapabilitiesSupportedOrientations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedOrientations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedOrientations(input string) (*PrinterCapabilitiesSupportedOrientations, error) {
	vals := map[string]PrinterCapabilitiesSupportedOrientations{
		"landscape":        PrinterCapabilitiesSupportedOrientationsLandscape,
		"portrait":         PrinterCapabilitiesSupportedOrientationsPortrait,
		"reverselandscape": PrinterCapabilitiesSupportedOrientationsReverseLandscape,
		"reverseportrait":  PrinterCapabilitiesSupportedOrientationsReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedOrientations(input)
	return &out, nil
}

type PrinterCapabilitiesSupportedPresentationDirections string

const (
	PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromBottomLeft         PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromBottomLeft"
	PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromBottomRight        PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromBottomRight"
	PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromTopLeft            PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromTopLeft"
	PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromTopRight           PrinterCapabilitiesSupportedPresentationDirections = "clockwiseFromTopRight"
	PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromBottomLeft  PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromBottomLeft"
	PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromBottomRight PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromBottomRight"
	PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromTopLeft     PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromTopLeft"
	PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromTopRight    PrinterCapabilitiesSupportedPresentationDirections = "counterClockwiseFromTopRight"
)

func PossibleValuesForPrinterCapabilitiesSupportedPresentationDirections() []string {
	return []string{
		string(PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromBottomLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromBottomRight),
		string(PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromTopLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromTopRight),
		string(PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromBottomLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromBottomRight),
		string(PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromTopLeft),
		string(PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromTopRight),
	}
}

func (s *PrinterCapabilitiesSupportedPresentationDirections) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedPresentationDirections(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedPresentationDirections(input string) (*PrinterCapabilitiesSupportedPresentationDirections, error) {
	vals := map[string]PrinterCapabilitiesSupportedPresentationDirections{
		"clockwisefrombottomleft":         PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterCapabilitiesSupportedPresentationDirectionsClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterCapabilitiesSupportedPresentationDirectionsCounterClockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedPresentationDirections(input)
	return &out, nil
}

type PrinterCapabilitiesSupportedPrintQualities string

const (
	PrinterCapabilitiesSupportedPrintQualitiesHigh   PrinterCapabilitiesSupportedPrintQualities = "high"
	PrinterCapabilitiesSupportedPrintQualitiesLow    PrinterCapabilitiesSupportedPrintQualities = "low"
	PrinterCapabilitiesSupportedPrintQualitiesMedium PrinterCapabilitiesSupportedPrintQualities = "medium"
)

func PossibleValuesForPrinterCapabilitiesSupportedPrintQualities() []string {
	return []string{
		string(PrinterCapabilitiesSupportedPrintQualitiesHigh),
		string(PrinterCapabilitiesSupportedPrintQualitiesLow),
		string(PrinterCapabilitiesSupportedPrintQualitiesMedium),
	}
}

func (s *PrinterCapabilitiesSupportedPrintQualities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedPrintQualities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedPrintQualities(input string) (*PrinterCapabilitiesSupportedPrintQualities, error) {
	vals := map[string]PrinterCapabilitiesSupportedPrintQualities{
		"high":   PrinterCapabilitiesSupportedPrintQualitiesHigh,
		"low":    PrinterCapabilitiesSupportedPrintQualitiesLow,
		"medium": PrinterCapabilitiesSupportedPrintQualitiesMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedPrintQualities(input)
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

type PrinterDefaultsDuplexConfiguration string

const (
	PrinterDefaultsDuplexConfigurationOneSided          PrinterDefaultsDuplexConfiguration = "oneSided"
	PrinterDefaultsDuplexConfigurationTwoSidedLongEdge  PrinterDefaultsDuplexConfiguration = "twoSidedLongEdge"
	PrinterDefaultsDuplexConfigurationTwoSidedShortEdge PrinterDefaultsDuplexConfiguration = "twoSidedShortEdge"
)

func PossibleValuesForPrinterDefaultsDuplexConfiguration() []string {
	return []string{
		string(PrinterDefaultsDuplexConfigurationOneSided),
		string(PrinterDefaultsDuplexConfigurationTwoSidedLongEdge),
		string(PrinterDefaultsDuplexConfigurationTwoSidedShortEdge),
	}
}

func (s *PrinterDefaultsDuplexConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsDuplexConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsDuplexConfiguration(input string) (*PrinterDefaultsDuplexConfiguration, error) {
	vals := map[string]PrinterDefaultsDuplexConfiguration{
		"onesided":          PrinterDefaultsDuplexConfigurationOneSided,
		"twosidedlongedge":  PrinterDefaultsDuplexConfigurationTwoSidedLongEdge,
		"twosidedshortedge": PrinterDefaultsDuplexConfigurationTwoSidedShortEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsDuplexConfiguration(input)
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

type PrinterDefaultsPresentationDirection string

const (
	PrinterDefaultsPresentationDirectionClockwiseFromBottomLeft         PrinterDefaultsPresentationDirection = "clockwiseFromBottomLeft"
	PrinterDefaultsPresentationDirectionClockwiseFromBottomRight        PrinterDefaultsPresentationDirection = "clockwiseFromBottomRight"
	PrinterDefaultsPresentationDirectionClockwiseFromTopLeft            PrinterDefaultsPresentationDirection = "clockwiseFromTopLeft"
	PrinterDefaultsPresentationDirectionClockwiseFromTopRight           PrinterDefaultsPresentationDirection = "clockwiseFromTopRight"
	PrinterDefaultsPresentationDirectionCounterClockwiseFromBottomLeft  PrinterDefaultsPresentationDirection = "counterClockwiseFromBottomLeft"
	PrinterDefaultsPresentationDirectionCounterClockwiseFromBottomRight PrinterDefaultsPresentationDirection = "counterClockwiseFromBottomRight"
	PrinterDefaultsPresentationDirectionCounterClockwiseFromTopLeft     PrinterDefaultsPresentationDirection = "counterClockwiseFromTopLeft"
	PrinterDefaultsPresentationDirectionCounterClockwiseFromTopRight    PrinterDefaultsPresentationDirection = "counterClockwiseFromTopRight"
)

func PossibleValuesForPrinterDefaultsPresentationDirection() []string {
	return []string{
		string(PrinterDefaultsPresentationDirectionClockwiseFromBottomLeft),
		string(PrinterDefaultsPresentationDirectionClockwiseFromBottomRight),
		string(PrinterDefaultsPresentationDirectionClockwiseFromTopLeft),
		string(PrinterDefaultsPresentationDirectionClockwiseFromTopRight),
		string(PrinterDefaultsPresentationDirectionCounterClockwiseFromBottomLeft),
		string(PrinterDefaultsPresentationDirectionCounterClockwiseFromBottomRight),
		string(PrinterDefaultsPresentationDirectionCounterClockwiseFromTopLeft),
		string(PrinterDefaultsPresentationDirectionCounterClockwiseFromTopRight),
	}
}

func (s *PrinterDefaultsPresentationDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsPresentationDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsPresentationDirection(input string) (*PrinterDefaultsPresentationDirection, error) {
	vals := map[string]PrinterDefaultsPresentationDirection{
		"clockwisefrombottomleft":         PrinterDefaultsPresentationDirectionClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDefaultsPresentationDirectionClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDefaultsPresentationDirectionClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDefaultsPresentationDirectionClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDefaultsPresentationDirectionCounterClockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDefaultsPresentationDirectionCounterClockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDefaultsPresentationDirectionCounterClockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDefaultsPresentationDirectionCounterClockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPresentationDirection(input)
	return &out, nil
}

type PrinterDefaultsPrintColorConfiguration string

const (
	PrinterDefaultsPrintColorConfigurationAuto          PrinterDefaultsPrintColorConfiguration = "auto"
	PrinterDefaultsPrintColorConfigurationBlackAndWhite PrinterDefaultsPrintColorConfiguration = "blackAndWhite"
	PrinterDefaultsPrintColorConfigurationColor         PrinterDefaultsPrintColorConfiguration = "color"
	PrinterDefaultsPrintColorConfigurationGrayscale     PrinterDefaultsPrintColorConfiguration = "grayscale"
)

func PossibleValuesForPrinterDefaultsPrintColorConfiguration() []string {
	return []string{
		string(PrinterDefaultsPrintColorConfigurationAuto),
		string(PrinterDefaultsPrintColorConfigurationBlackAndWhite),
		string(PrinterDefaultsPrintColorConfigurationColor),
		string(PrinterDefaultsPrintColorConfigurationGrayscale),
	}
}

func (s *PrinterDefaultsPrintColorConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsPrintColorConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsPrintColorConfiguration(input string) (*PrinterDefaultsPrintColorConfiguration, error) {
	vals := map[string]PrinterDefaultsPrintColorConfiguration{
		"auto":          PrinterDefaultsPrintColorConfigurationAuto,
		"blackandwhite": PrinterDefaultsPrintColorConfigurationBlackAndWhite,
		"color":         PrinterDefaultsPrintColorConfigurationColor,
		"grayscale":     PrinterDefaultsPrintColorConfigurationGrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPrintColorConfiguration(input)
	return &out, nil
}

type PrinterDefaultsPrintQuality string

const (
	PrinterDefaultsPrintQualityHigh   PrinterDefaultsPrintQuality = "high"
	PrinterDefaultsPrintQualityLow    PrinterDefaultsPrintQuality = "low"
	PrinterDefaultsPrintQualityMedium PrinterDefaultsPrintQuality = "medium"
)

func PossibleValuesForPrinterDefaultsPrintQuality() []string {
	return []string{
		string(PrinterDefaultsPrintQualityHigh),
		string(PrinterDefaultsPrintQualityLow),
		string(PrinterDefaultsPrintQualityMedium),
	}
}

func (s *PrinterDefaultsPrintQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsPrintQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsPrintQuality(input string) (*PrinterDefaultsPrintQuality, error) {
	vals := map[string]PrinterDefaultsPrintQuality{
		"high":   PrinterDefaultsPrintQualityHigh,
		"low":    PrinterDefaultsPrintQualityLow,
		"medium": PrinterDefaultsPrintQualityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsPrintQuality(input)
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

type PrinterDocumentConfigurationColorMode string

const (
	PrinterDocumentConfigurationColorModeAuto          PrinterDocumentConfigurationColorMode = "auto"
	PrinterDocumentConfigurationColorModeBlackAndWhite PrinterDocumentConfigurationColorMode = "blackAndWhite"
	PrinterDocumentConfigurationColorModeColor         PrinterDocumentConfigurationColorMode = "color"
	PrinterDocumentConfigurationColorModeGrayscale     PrinterDocumentConfigurationColorMode = "grayscale"
)

func PossibleValuesForPrinterDocumentConfigurationColorMode() []string {
	return []string{
		string(PrinterDocumentConfigurationColorModeAuto),
		string(PrinterDocumentConfigurationColorModeBlackAndWhite),
		string(PrinterDocumentConfigurationColorModeColor),
		string(PrinterDocumentConfigurationColorModeGrayscale),
	}
}

func (s *PrinterDocumentConfigurationColorMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationColorMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationColorMode(input string) (*PrinterDocumentConfigurationColorMode, error) {
	vals := map[string]PrinterDocumentConfigurationColorMode{
		"auto":          PrinterDocumentConfigurationColorModeAuto,
		"blackandwhite": PrinterDocumentConfigurationColorModeBlackAndWhite,
		"color":         PrinterDocumentConfigurationColorModeColor,
		"grayscale":     PrinterDocumentConfigurationColorModeGrayscale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationColorMode(input)
	return &out, nil
}

type PrinterDocumentConfigurationDuplexMode string

const (
	PrinterDocumentConfigurationDuplexModeFlipOnLongEdge  PrinterDocumentConfigurationDuplexMode = "flipOnLongEdge"
	PrinterDocumentConfigurationDuplexModeFlipOnShortEdge PrinterDocumentConfigurationDuplexMode = "flipOnShortEdge"
	PrinterDocumentConfigurationDuplexModeOneSided        PrinterDocumentConfigurationDuplexMode = "oneSided"
)

func PossibleValuesForPrinterDocumentConfigurationDuplexMode() []string {
	return []string{
		string(PrinterDocumentConfigurationDuplexModeFlipOnLongEdge),
		string(PrinterDocumentConfigurationDuplexModeFlipOnShortEdge),
		string(PrinterDocumentConfigurationDuplexModeOneSided),
	}
}

func (s *PrinterDocumentConfigurationDuplexMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationDuplexMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationDuplexMode(input string) (*PrinterDocumentConfigurationDuplexMode, error) {
	vals := map[string]PrinterDocumentConfigurationDuplexMode{
		"fliponlongedge":  PrinterDocumentConfigurationDuplexModeFlipOnLongEdge,
		"fliponshortedge": PrinterDocumentConfigurationDuplexModeFlipOnShortEdge,
		"onesided":        PrinterDocumentConfigurationDuplexModeOneSided,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationDuplexMode(input)
	return &out, nil
}

type PrinterDocumentConfigurationFeedDirection string

const (
	PrinterDocumentConfigurationFeedDirectionLongEdgeFirst  PrinterDocumentConfigurationFeedDirection = "longEdgeFirst"
	PrinterDocumentConfigurationFeedDirectionShortEdgeFirst PrinterDocumentConfigurationFeedDirection = "shortEdgeFirst"
)

func PossibleValuesForPrinterDocumentConfigurationFeedDirection() []string {
	return []string{
		string(PrinterDocumentConfigurationFeedDirectionLongEdgeFirst),
		string(PrinterDocumentConfigurationFeedDirectionShortEdgeFirst),
	}
}

func (s *PrinterDocumentConfigurationFeedDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationFeedDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationFeedDirection(input string) (*PrinterDocumentConfigurationFeedDirection, error) {
	vals := map[string]PrinterDocumentConfigurationFeedDirection{
		"longedgefirst":  PrinterDocumentConfigurationFeedDirectionLongEdgeFirst,
		"shortedgefirst": PrinterDocumentConfigurationFeedDirectionShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFeedDirection(input)
	return &out, nil
}

type PrinterDocumentConfigurationFeedOrientation string

const (
	PrinterDocumentConfigurationFeedOrientationLongEdgeFirst  PrinterDocumentConfigurationFeedOrientation = "longEdgeFirst"
	PrinterDocumentConfigurationFeedOrientationShortEdgeFirst PrinterDocumentConfigurationFeedOrientation = "shortEdgeFirst"
)

func PossibleValuesForPrinterDocumentConfigurationFeedOrientation() []string {
	return []string{
		string(PrinterDocumentConfigurationFeedOrientationLongEdgeFirst),
		string(PrinterDocumentConfigurationFeedOrientationShortEdgeFirst),
	}
}

func (s *PrinterDocumentConfigurationFeedOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationFeedOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationFeedOrientation(input string) (*PrinterDocumentConfigurationFeedOrientation, error) {
	vals := map[string]PrinterDocumentConfigurationFeedOrientation{
		"longedgefirst":  PrinterDocumentConfigurationFeedOrientationLongEdgeFirst,
		"shortedgefirst": PrinterDocumentConfigurationFeedOrientationShortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFeedOrientation(input)
	return &out, nil
}

type PrinterDocumentConfigurationFinishings string

const (
	PrinterDocumentConfigurationFinishingsBind              PrinterDocumentConfigurationFinishings = "bind"
	PrinterDocumentConfigurationFinishingsCover             PrinterDocumentConfigurationFinishings = "cover"
	PrinterDocumentConfigurationFinishingsNone              PrinterDocumentConfigurationFinishings = "none"
	PrinterDocumentConfigurationFinishingsPunch             PrinterDocumentConfigurationFinishings = "punch"
	PrinterDocumentConfigurationFinishingsSaddleStitch      PrinterDocumentConfigurationFinishings = "saddleStitch"
	PrinterDocumentConfigurationFinishingsStaple            PrinterDocumentConfigurationFinishings = "staple"
	PrinterDocumentConfigurationFinishingsStapleBottomLeft  PrinterDocumentConfigurationFinishings = "stapleBottomLeft"
	PrinterDocumentConfigurationFinishingsStapleBottomRight PrinterDocumentConfigurationFinishings = "stapleBottomRight"
	PrinterDocumentConfigurationFinishingsStapleDualBottom  PrinterDocumentConfigurationFinishings = "stapleDualBottom"
	PrinterDocumentConfigurationFinishingsStapleDualLeft    PrinterDocumentConfigurationFinishings = "stapleDualLeft"
	PrinterDocumentConfigurationFinishingsStapleDualRight   PrinterDocumentConfigurationFinishings = "stapleDualRight"
	PrinterDocumentConfigurationFinishingsStapleDualTop     PrinterDocumentConfigurationFinishings = "stapleDualTop"
	PrinterDocumentConfigurationFinishingsStapleTopLeft     PrinterDocumentConfigurationFinishings = "stapleTopLeft"
	PrinterDocumentConfigurationFinishingsStapleTopRight    PrinterDocumentConfigurationFinishings = "stapleTopRight"
	PrinterDocumentConfigurationFinishingsStitchBottomEdge  PrinterDocumentConfigurationFinishings = "stitchBottomEdge"
	PrinterDocumentConfigurationFinishingsStitchEdge        PrinterDocumentConfigurationFinishings = "stitchEdge"
	PrinterDocumentConfigurationFinishingsStitchLeftEdge    PrinterDocumentConfigurationFinishings = "stitchLeftEdge"
	PrinterDocumentConfigurationFinishingsStitchRightEdge   PrinterDocumentConfigurationFinishings = "stitchRightEdge"
	PrinterDocumentConfigurationFinishingsStitchTopEdge     PrinterDocumentConfigurationFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterDocumentConfigurationFinishings() []string {
	return []string{
		string(PrinterDocumentConfigurationFinishingsBind),
		string(PrinterDocumentConfigurationFinishingsCover),
		string(PrinterDocumentConfigurationFinishingsNone),
		string(PrinterDocumentConfigurationFinishingsPunch),
		string(PrinterDocumentConfigurationFinishingsSaddleStitch),
		string(PrinterDocumentConfigurationFinishingsStaple),
		string(PrinterDocumentConfigurationFinishingsStapleBottomLeft),
		string(PrinterDocumentConfigurationFinishingsStapleBottomRight),
		string(PrinterDocumentConfigurationFinishingsStapleDualBottom),
		string(PrinterDocumentConfigurationFinishingsStapleDualLeft),
		string(PrinterDocumentConfigurationFinishingsStapleDualRight),
		string(PrinterDocumentConfigurationFinishingsStapleDualTop),
		string(PrinterDocumentConfigurationFinishingsStapleTopLeft),
		string(PrinterDocumentConfigurationFinishingsStapleTopRight),
		string(PrinterDocumentConfigurationFinishingsStitchBottomEdge),
		string(PrinterDocumentConfigurationFinishingsStitchEdge),
		string(PrinterDocumentConfigurationFinishingsStitchLeftEdge),
		string(PrinterDocumentConfigurationFinishingsStitchRightEdge),
		string(PrinterDocumentConfigurationFinishingsStitchTopEdge),
	}
}

func (s *PrinterDocumentConfigurationFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationFinishings(input string) (*PrinterDocumentConfigurationFinishings, error) {
	vals := map[string]PrinterDocumentConfigurationFinishings{
		"bind":              PrinterDocumentConfigurationFinishingsBind,
		"cover":             PrinterDocumentConfigurationFinishingsCover,
		"none":              PrinterDocumentConfigurationFinishingsNone,
		"punch":             PrinterDocumentConfigurationFinishingsPunch,
		"saddlestitch":      PrinterDocumentConfigurationFinishingsSaddleStitch,
		"staple":            PrinterDocumentConfigurationFinishingsStaple,
		"staplebottomleft":  PrinterDocumentConfigurationFinishingsStapleBottomLeft,
		"staplebottomright": PrinterDocumentConfigurationFinishingsStapleBottomRight,
		"stapledualbottom":  PrinterDocumentConfigurationFinishingsStapleDualBottom,
		"stapledualleft":    PrinterDocumentConfigurationFinishingsStapleDualLeft,
		"stapledualright":   PrinterDocumentConfigurationFinishingsStapleDualRight,
		"stapledualtop":     PrinterDocumentConfigurationFinishingsStapleDualTop,
		"stapletopleft":     PrinterDocumentConfigurationFinishingsStapleTopLeft,
		"stapletopright":    PrinterDocumentConfigurationFinishingsStapleTopRight,
		"stitchbottomedge":  PrinterDocumentConfigurationFinishingsStitchBottomEdge,
		"stitchedge":        PrinterDocumentConfigurationFinishingsStitchEdge,
		"stitchleftedge":    PrinterDocumentConfigurationFinishingsStitchLeftEdge,
		"stitchrightedge":   PrinterDocumentConfigurationFinishingsStitchRightEdge,
		"stitchtopedge":     PrinterDocumentConfigurationFinishingsStitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFinishings(input)
	return &out, nil
}

type PrinterDocumentConfigurationMultipageLayout string

const (
	PrinterDocumentConfigurationMultipageLayoutClockwiseFromBottomLeft         PrinterDocumentConfigurationMultipageLayout = "clockwiseFromBottomLeft"
	PrinterDocumentConfigurationMultipageLayoutClockwiseFromBottomRight        PrinterDocumentConfigurationMultipageLayout = "clockwiseFromBottomRight"
	PrinterDocumentConfigurationMultipageLayoutClockwiseFromTopLeft            PrinterDocumentConfigurationMultipageLayout = "clockwiseFromTopLeft"
	PrinterDocumentConfigurationMultipageLayoutClockwiseFromTopRight           PrinterDocumentConfigurationMultipageLayout = "clockwiseFromTopRight"
	PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromBottomLeft  PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromBottomLeft"
	PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromBottomRight PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromBottomRight"
	PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromTopLeft     PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromTopLeft"
	PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromTopRight    PrinterDocumentConfigurationMultipageLayout = "counterclockwiseFromTopRight"
)

func PossibleValuesForPrinterDocumentConfigurationMultipageLayout() []string {
	return []string{
		string(PrinterDocumentConfigurationMultipageLayoutClockwiseFromBottomLeft),
		string(PrinterDocumentConfigurationMultipageLayoutClockwiseFromBottomRight),
		string(PrinterDocumentConfigurationMultipageLayoutClockwiseFromTopLeft),
		string(PrinterDocumentConfigurationMultipageLayoutClockwiseFromTopRight),
		string(PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromBottomLeft),
		string(PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromBottomRight),
		string(PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromTopLeft),
		string(PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromTopRight),
	}
}

func (s *PrinterDocumentConfigurationMultipageLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationMultipageLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationMultipageLayout(input string) (*PrinterDocumentConfigurationMultipageLayout, error) {
	vals := map[string]PrinterDocumentConfigurationMultipageLayout{
		"clockwisefrombottomleft":         PrinterDocumentConfigurationMultipageLayoutClockwiseFromBottomLeft,
		"clockwisefrombottomright":        PrinterDocumentConfigurationMultipageLayoutClockwiseFromBottomRight,
		"clockwisefromtopleft":            PrinterDocumentConfigurationMultipageLayoutClockwiseFromTopLeft,
		"clockwisefromtopright":           PrinterDocumentConfigurationMultipageLayoutClockwiseFromTopRight,
		"counterclockwisefrombottomleft":  PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromBottomLeft,
		"counterclockwisefrombottomright": PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromBottomRight,
		"counterclockwisefromtopleft":     PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromTopLeft,
		"counterclockwisefromtopright":    PrinterDocumentConfigurationMultipageLayoutCounterclockwiseFromTopRight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationMultipageLayout(input)
	return &out, nil
}

type PrinterDocumentConfigurationOrientation string

const (
	PrinterDocumentConfigurationOrientationLandscape        PrinterDocumentConfigurationOrientation = "landscape"
	PrinterDocumentConfigurationOrientationPortrait         PrinterDocumentConfigurationOrientation = "portrait"
	PrinterDocumentConfigurationOrientationReverseLandscape PrinterDocumentConfigurationOrientation = "reverseLandscape"
	PrinterDocumentConfigurationOrientationReversePortrait  PrinterDocumentConfigurationOrientation = "reversePortrait"
)

func PossibleValuesForPrinterDocumentConfigurationOrientation() []string {
	return []string{
		string(PrinterDocumentConfigurationOrientationLandscape),
		string(PrinterDocumentConfigurationOrientationPortrait),
		string(PrinterDocumentConfigurationOrientationReverseLandscape),
		string(PrinterDocumentConfigurationOrientationReversePortrait),
	}
}

func (s *PrinterDocumentConfigurationOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationOrientation(input string) (*PrinterDocumentConfigurationOrientation, error) {
	vals := map[string]PrinterDocumentConfigurationOrientation{
		"landscape":        PrinterDocumentConfigurationOrientationLandscape,
		"portrait":         PrinterDocumentConfigurationOrientationPortrait,
		"reverselandscape": PrinterDocumentConfigurationOrientationReverseLandscape,
		"reverseportrait":  PrinterDocumentConfigurationOrientationReversePortrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationOrientation(input)
	return &out, nil
}

type PrinterDocumentConfigurationQuality string

const (
	PrinterDocumentConfigurationQualityHigh   PrinterDocumentConfigurationQuality = "high"
	PrinterDocumentConfigurationQualityLow    PrinterDocumentConfigurationQuality = "low"
	PrinterDocumentConfigurationQualityMedium PrinterDocumentConfigurationQuality = "medium"
)

func PossibleValuesForPrinterDocumentConfigurationQuality() []string {
	return []string{
		string(PrinterDocumentConfigurationQualityHigh),
		string(PrinterDocumentConfigurationQualityLow),
		string(PrinterDocumentConfigurationQualityMedium),
	}
}

func (s *PrinterDocumentConfigurationQuality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationQuality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationQuality(input string) (*PrinterDocumentConfigurationQuality, error) {
	vals := map[string]PrinterDocumentConfigurationQuality{
		"high":   PrinterDocumentConfigurationQualityHigh,
		"low":    PrinterDocumentConfigurationQualityLow,
		"medium": PrinterDocumentConfigurationQualityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationQuality(input)
	return &out, nil
}

type PrinterDocumentConfigurationScaling string

const (
	PrinterDocumentConfigurationScalingAuto        PrinterDocumentConfigurationScaling = "auto"
	PrinterDocumentConfigurationScalingFill        PrinterDocumentConfigurationScaling = "fill"
	PrinterDocumentConfigurationScalingFit         PrinterDocumentConfigurationScaling = "fit"
	PrinterDocumentConfigurationScalingNone        PrinterDocumentConfigurationScaling = "none"
	PrinterDocumentConfigurationScalingShrinkToFit PrinterDocumentConfigurationScaling = "shrinkToFit"
)

func PossibleValuesForPrinterDocumentConfigurationScaling() []string {
	return []string{
		string(PrinterDocumentConfigurationScalingAuto),
		string(PrinterDocumentConfigurationScalingFill),
		string(PrinterDocumentConfigurationScalingFit),
		string(PrinterDocumentConfigurationScalingNone),
		string(PrinterDocumentConfigurationScalingShrinkToFit),
	}
}

func (s *PrinterDocumentConfigurationScaling) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationScaling(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationScaling(input string) (*PrinterDocumentConfigurationScaling, error) {
	vals := map[string]PrinterDocumentConfigurationScaling{
		"auto":        PrinterDocumentConfigurationScalingAuto,
		"fill":        PrinterDocumentConfigurationScalingFill,
		"fit":         PrinterDocumentConfigurationScalingFit,
		"none":        PrinterDocumentConfigurationScalingNone,
		"shrinktofit": PrinterDocumentConfigurationScalingShrinkToFit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationScaling(input)
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

type PrinterStatusProcessingState string

const (
	PrinterStatusProcessingStateIdle       PrinterStatusProcessingState = "idle"
	PrinterStatusProcessingStateProcessing PrinterStatusProcessingState = "processing"
	PrinterStatusProcessingStateStopped    PrinterStatusProcessingState = "stopped"
	PrinterStatusProcessingStateUnknown    PrinterStatusProcessingState = "unknown"
)

func PossibleValuesForPrinterStatusProcessingState() []string {
	return []string{
		string(PrinterStatusProcessingStateIdle),
		string(PrinterStatusProcessingStateProcessing),
		string(PrinterStatusProcessingStateStopped),
		string(PrinterStatusProcessingStateUnknown),
	}
}

func (s *PrinterStatusProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusProcessingState(input string) (*PrinterStatusProcessingState, error) {
	vals := map[string]PrinterStatusProcessingState{
		"idle":       PrinterStatusProcessingStateIdle,
		"processing": PrinterStatusProcessingStateProcessing,
		"stopped":    PrinterStatusProcessingStateStopped,
		"unknown":    PrinterStatusProcessingStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusProcessingState(input)
	return &out, nil
}

type PrinterStatusProcessingStateReasons string

const (
	PrinterStatusProcessingStateReasonsConnectingToDevice                 PrinterStatusProcessingStateReasons = "connectingToDevice"
	PrinterStatusProcessingStateReasonsCoverOpen                          PrinterStatusProcessingStateReasons = "coverOpen"
	PrinterStatusProcessingStateReasonsDeveloperEmpty                     PrinterStatusProcessingStateReasons = "developerEmpty"
	PrinterStatusProcessingStateReasonsDeveloperLow                       PrinterStatusProcessingStateReasons = "developerLow"
	PrinterStatusProcessingStateReasonsDoorOpen                           PrinterStatusProcessingStateReasons = "doorOpen"
	PrinterStatusProcessingStateReasonsFuserOverTemp                      PrinterStatusProcessingStateReasons = "fuserOverTemp"
	PrinterStatusProcessingStateReasonsFuserUnderTemp                     PrinterStatusProcessingStateReasons = "fuserUnderTemp"
	PrinterStatusProcessingStateReasonsInputTrayMissing                   PrinterStatusProcessingStateReasons = "inputTrayMissing"
	PrinterStatusProcessingStateReasonsInterlockOpen                      PrinterStatusProcessingStateReasons = "interlockOpen"
	PrinterStatusProcessingStateReasonsInterpreterResourceUnavailable     PrinterStatusProcessingStateReasons = "interpreterResourceUnavailable"
	PrinterStatusProcessingStateReasonsMarkerSupplyEmpty                  PrinterStatusProcessingStateReasons = "markerSupplyEmpty"
	PrinterStatusProcessingStateReasonsMarkerSupplyLow                    PrinterStatusProcessingStateReasons = "markerSupplyLow"
	PrinterStatusProcessingStateReasonsMarkerWasteAlmostFull              PrinterStatusProcessingStateReasons = "markerWasteAlmostFull"
	PrinterStatusProcessingStateReasonsMarkerWasteFull                    PrinterStatusProcessingStateReasons = "markerWasteFull"
	PrinterStatusProcessingStateReasonsMediaEmpty                         PrinterStatusProcessingStateReasons = "mediaEmpty"
	PrinterStatusProcessingStateReasonsMediaJam                           PrinterStatusProcessingStateReasons = "mediaJam"
	PrinterStatusProcessingStateReasonsMediaLow                           PrinterStatusProcessingStateReasons = "mediaLow"
	PrinterStatusProcessingStateReasonsMediaNeeded                        PrinterStatusProcessingStateReasons = "mediaNeeded"
	PrinterStatusProcessingStateReasonsMovingToPaused                     PrinterStatusProcessingStateReasons = "movingToPaused"
	PrinterStatusProcessingStateReasonsNone                               PrinterStatusProcessingStateReasons = "none"
	PrinterStatusProcessingStateReasonsOpticalPhotoConductorLifeOver      PrinterStatusProcessingStateReasons = "opticalPhotoConductorLifeOver"
	PrinterStatusProcessingStateReasonsOpticalPhotoConductorNearEndOfLife PrinterStatusProcessingStateReasons = "opticalPhotoConductorNearEndOfLife"
	PrinterStatusProcessingStateReasonsOther                              PrinterStatusProcessingStateReasons = "other"
	PrinterStatusProcessingStateReasonsOutputAreaAlmostFull               PrinterStatusProcessingStateReasons = "outputAreaAlmostFull"
	PrinterStatusProcessingStateReasonsOutputAreaFull                     PrinterStatusProcessingStateReasons = "outputAreaFull"
	PrinterStatusProcessingStateReasonsOutputTrayMissing                  PrinterStatusProcessingStateReasons = "outputTrayMissing"
	PrinterStatusProcessingStateReasonsPaused                             PrinterStatusProcessingStateReasons = "paused"
	PrinterStatusProcessingStateReasonsShutdown                           PrinterStatusProcessingStateReasons = "shutdown"
	PrinterStatusProcessingStateReasonsSpoolAreaFull                      PrinterStatusProcessingStateReasons = "spoolAreaFull"
	PrinterStatusProcessingStateReasonsStoppedPartially                   PrinterStatusProcessingStateReasons = "stoppedPartially"
	PrinterStatusProcessingStateReasonsStopping                           PrinterStatusProcessingStateReasons = "stopping"
	PrinterStatusProcessingStateReasonsTimedOut                           PrinterStatusProcessingStateReasons = "timedOut"
	PrinterStatusProcessingStateReasonsTonerEmpty                         PrinterStatusProcessingStateReasons = "tonerEmpty"
	PrinterStatusProcessingStateReasonsTonerLow                           PrinterStatusProcessingStateReasons = "tonerLow"
)

func PossibleValuesForPrinterStatusProcessingStateReasons() []string {
	return []string{
		string(PrinterStatusProcessingStateReasonsConnectingToDevice),
		string(PrinterStatusProcessingStateReasonsCoverOpen),
		string(PrinterStatusProcessingStateReasonsDeveloperEmpty),
		string(PrinterStatusProcessingStateReasonsDeveloperLow),
		string(PrinterStatusProcessingStateReasonsDoorOpen),
		string(PrinterStatusProcessingStateReasonsFuserOverTemp),
		string(PrinterStatusProcessingStateReasonsFuserUnderTemp),
		string(PrinterStatusProcessingStateReasonsInputTrayMissing),
		string(PrinterStatusProcessingStateReasonsInterlockOpen),
		string(PrinterStatusProcessingStateReasonsInterpreterResourceUnavailable),
		string(PrinterStatusProcessingStateReasonsMarkerSupplyEmpty),
		string(PrinterStatusProcessingStateReasonsMarkerSupplyLow),
		string(PrinterStatusProcessingStateReasonsMarkerWasteAlmostFull),
		string(PrinterStatusProcessingStateReasonsMarkerWasteFull),
		string(PrinterStatusProcessingStateReasonsMediaEmpty),
		string(PrinterStatusProcessingStateReasonsMediaJam),
		string(PrinterStatusProcessingStateReasonsMediaLow),
		string(PrinterStatusProcessingStateReasonsMediaNeeded),
		string(PrinterStatusProcessingStateReasonsMovingToPaused),
		string(PrinterStatusProcessingStateReasonsNone),
		string(PrinterStatusProcessingStateReasonsOpticalPhotoConductorLifeOver),
		string(PrinterStatusProcessingStateReasonsOpticalPhotoConductorNearEndOfLife),
		string(PrinterStatusProcessingStateReasonsOther),
		string(PrinterStatusProcessingStateReasonsOutputAreaAlmostFull),
		string(PrinterStatusProcessingStateReasonsOutputAreaFull),
		string(PrinterStatusProcessingStateReasonsOutputTrayMissing),
		string(PrinterStatusProcessingStateReasonsPaused),
		string(PrinterStatusProcessingStateReasonsShutdown),
		string(PrinterStatusProcessingStateReasonsSpoolAreaFull),
		string(PrinterStatusProcessingStateReasonsStoppedPartially),
		string(PrinterStatusProcessingStateReasonsStopping),
		string(PrinterStatusProcessingStateReasonsTimedOut),
		string(PrinterStatusProcessingStateReasonsTonerEmpty),
		string(PrinterStatusProcessingStateReasonsTonerLow),
	}
}

func (s *PrinterStatusProcessingStateReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusProcessingStateReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusProcessingStateReasons(input string) (*PrinterStatusProcessingStateReasons, error) {
	vals := map[string]PrinterStatusProcessingStateReasons{
		"connectingtodevice":                 PrinterStatusProcessingStateReasonsConnectingToDevice,
		"coveropen":                          PrinterStatusProcessingStateReasonsCoverOpen,
		"developerempty":                     PrinterStatusProcessingStateReasonsDeveloperEmpty,
		"developerlow":                       PrinterStatusProcessingStateReasonsDeveloperLow,
		"dooropen":                           PrinterStatusProcessingStateReasonsDoorOpen,
		"fuserovertemp":                      PrinterStatusProcessingStateReasonsFuserOverTemp,
		"fuserundertemp":                     PrinterStatusProcessingStateReasonsFuserUnderTemp,
		"inputtraymissing":                   PrinterStatusProcessingStateReasonsInputTrayMissing,
		"interlockopen":                      PrinterStatusProcessingStateReasonsInterlockOpen,
		"interpreterresourceunavailable":     PrinterStatusProcessingStateReasonsInterpreterResourceUnavailable,
		"markersupplyempty":                  PrinterStatusProcessingStateReasonsMarkerSupplyEmpty,
		"markersupplylow":                    PrinterStatusProcessingStateReasonsMarkerSupplyLow,
		"markerwastealmostfull":              PrinterStatusProcessingStateReasonsMarkerWasteAlmostFull,
		"markerwastefull":                    PrinterStatusProcessingStateReasonsMarkerWasteFull,
		"mediaempty":                         PrinterStatusProcessingStateReasonsMediaEmpty,
		"mediajam":                           PrinterStatusProcessingStateReasonsMediaJam,
		"medialow":                           PrinterStatusProcessingStateReasonsMediaLow,
		"medianeeded":                        PrinterStatusProcessingStateReasonsMediaNeeded,
		"movingtopaused":                     PrinterStatusProcessingStateReasonsMovingToPaused,
		"none":                               PrinterStatusProcessingStateReasonsNone,
		"opticalphotoconductorlifeover":      PrinterStatusProcessingStateReasonsOpticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife": PrinterStatusProcessingStateReasonsOpticalPhotoConductorNearEndOfLife,
		"other":                              PrinterStatusProcessingStateReasonsOther,
		"outputareaalmostfull":               PrinterStatusProcessingStateReasonsOutputAreaAlmostFull,
		"outputareafull":                     PrinterStatusProcessingStateReasonsOutputAreaFull,
		"outputtraymissing":                  PrinterStatusProcessingStateReasonsOutputTrayMissing,
		"paused":                             PrinterStatusProcessingStateReasonsPaused,
		"shutdown":                           PrinterStatusProcessingStateReasonsShutdown,
		"spoolareafull":                      PrinterStatusProcessingStateReasonsSpoolAreaFull,
		"stoppedpartially":                   PrinterStatusProcessingStateReasonsStoppedPartially,
		"stopping":                           PrinterStatusProcessingStateReasonsStopping,
		"timedout":                           PrinterStatusProcessingStateReasonsTimedOut,
		"tonerempty":                         PrinterStatusProcessingStateReasonsTonerEmpty,
		"tonerlow":                           PrinterStatusProcessingStateReasonsTonerLow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusProcessingStateReasons(input)
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

type ProjectParticipationAllowedAudiences string

const (
	ProjectParticipationAllowedAudiencesContacts               ProjectParticipationAllowedAudiences = "contacts"
	ProjectParticipationAllowedAudiencesEveryone               ProjectParticipationAllowedAudiences = "everyone"
	ProjectParticipationAllowedAudiencesFamily                 ProjectParticipationAllowedAudiences = "family"
	ProjectParticipationAllowedAudiencesFederatedOrganizations ProjectParticipationAllowedAudiences = "federatedOrganizations"
	ProjectParticipationAllowedAudiencesGroupMembers           ProjectParticipationAllowedAudiences = "groupMembers"
	ProjectParticipationAllowedAudiencesMe                     ProjectParticipationAllowedAudiences = "me"
	ProjectParticipationAllowedAudiencesOrganization           ProjectParticipationAllowedAudiences = "organization"
)

func PossibleValuesForProjectParticipationAllowedAudiences() []string {
	return []string{
		string(ProjectParticipationAllowedAudiencesContacts),
		string(ProjectParticipationAllowedAudiencesEveryone),
		string(ProjectParticipationAllowedAudiencesFamily),
		string(ProjectParticipationAllowedAudiencesFederatedOrganizations),
		string(ProjectParticipationAllowedAudiencesGroupMembers),
		string(ProjectParticipationAllowedAudiencesMe),
		string(ProjectParticipationAllowedAudiencesOrganization),
	}
}

func (s *ProjectParticipationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProjectParticipationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProjectParticipationAllowedAudiences(input string) (*ProjectParticipationAllowedAudiences, error) {
	vals := map[string]ProjectParticipationAllowedAudiences{
		"contacts":               ProjectParticipationAllowedAudiencesContacts,
		"everyone":               ProjectParticipationAllowedAudiencesEveryone,
		"family":                 ProjectParticipationAllowedAudiencesFamily,
		"federatedorganizations": ProjectParticipationAllowedAudiencesFederatedOrganizations,
		"groupmembers":           ProjectParticipationAllowedAudiencesGroupMembers,
		"me":                     ProjectParticipationAllowedAudiencesMe,
		"organization":           ProjectParticipationAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectParticipationAllowedAudiences(input)
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

type RelatedPersonRelationship string

const (
	RelatedPersonRelationshipAlternateContact RelatedPersonRelationship = "alternateContact"
	RelatedPersonRelationshipAssistant        RelatedPersonRelationship = "assistant"
	RelatedPersonRelationshipChild            RelatedPersonRelationship = "child"
	RelatedPersonRelationshipColleague        RelatedPersonRelationship = "colleague"
	RelatedPersonRelationshipDirectReport     RelatedPersonRelationship = "directReport"
	RelatedPersonRelationshipDotLineManager   RelatedPersonRelationship = "dotLineManager"
	RelatedPersonRelationshipDotLineReport    RelatedPersonRelationship = "dotLineReport"
	RelatedPersonRelationshipEmergencyContact RelatedPersonRelationship = "emergencyContact"
	RelatedPersonRelationshipFriend           RelatedPersonRelationship = "friend"
	RelatedPersonRelationshipManager          RelatedPersonRelationship = "manager"
	RelatedPersonRelationshipOther            RelatedPersonRelationship = "other"
	RelatedPersonRelationshipParent           RelatedPersonRelationship = "parent"
	RelatedPersonRelationshipSibling          RelatedPersonRelationship = "sibling"
	RelatedPersonRelationshipSponsor          RelatedPersonRelationship = "sponsor"
	RelatedPersonRelationshipSpouse           RelatedPersonRelationship = "spouse"
)

func PossibleValuesForRelatedPersonRelationship() []string {
	return []string{
		string(RelatedPersonRelationshipAlternateContact),
		string(RelatedPersonRelationshipAssistant),
		string(RelatedPersonRelationshipChild),
		string(RelatedPersonRelationshipColleague),
		string(RelatedPersonRelationshipDirectReport),
		string(RelatedPersonRelationshipDotLineManager),
		string(RelatedPersonRelationshipDotLineReport),
		string(RelatedPersonRelationshipEmergencyContact),
		string(RelatedPersonRelationshipFriend),
		string(RelatedPersonRelationshipManager),
		string(RelatedPersonRelationshipOther),
		string(RelatedPersonRelationshipParent),
		string(RelatedPersonRelationshipSibling),
		string(RelatedPersonRelationshipSponsor),
		string(RelatedPersonRelationshipSpouse),
	}
}

func (s *RelatedPersonRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRelatedPersonRelationship(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRelatedPersonRelationship(input string) (*RelatedPersonRelationship, error) {
	vals := map[string]RelatedPersonRelationship{
		"alternatecontact": RelatedPersonRelationshipAlternateContact,
		"assistant":        RelatedPersonRelationshipAssistant,
		"child":            RelatedPersonRelationshipChild,
		"colleague":        RelatedPersonRelationshipColleague,
		"directreport":     RelatedPersonRelationshipDirectReport,
		"dotlinemanager":   RelatedPersonRelationshipDotLineManager,
		"dotlinereport":    RelatedPersonRelationshipDotLineReport,
		"emergencycontact": RelatedPersonRelationshipEmergencyContact,
		"friend":           RelatedPersonRelationshipFriend,
		"manager":          RelatedPersonRelationshipManager,
		"other":            RelatedPersonRelationshipOther,
		"parent":           RelatedPersonRelationshipParent,
		"sibling":          RelatedPersonRelationshipSibling,
		"sponsor":          RelatedPersonRelationshipSponsor,
		"spouse":           RelatedPersonRelationshipSpouse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelatedPersonRelationship(input)
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

type SecurityBaselineContributingPolicySourceType string

const (
	SecurityBaselineContributingPolicySourceTypeDeviceConfiguration SecurityBaselineContributingPolicySourceType = "deviceConfiguration"
	SecurityBaselineContributingPolicySourceTypeDeviceIntent        SecurityBaselineContributingPolicySourceType = "deviceIntent"
)

func PossibleValuesForSecurityBaselineContributingPolicySourceType() []string {
	return []string{
		string(SecurityBaselineContributingPolicySourceTypeDeviceConfiguration),
		string(SecurityBaselineContributingPolicySourceTypeDeviceIntent),
	}
}

func (s *SecurityBaselineContributingPolicySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineContributingPolicySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineContributingPolicySourceType(input string) (*SecurityBaselineContributingPolicySourceType, error) {
	vals := map[string]SecurityBaselineContributingPolicySourceType{
		"deviceconfiguration": SecurityBaselineContributingPolicySourceTypeDeviceConfiguration,
		"deviceintent":        SecurityBaselineContributingPolicySourceTypeDeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineContributingPolicySourceType(input)
	return &out, nil
}

type SecurityBaselineSettingStateState string

const (
	SecurityBaselineSettingStateStateConflict      SecurityBaselineSettingStateState = "conflict"
	SecurityBaselineSettingStateStateError         SecurityBaselineSettingStateState = "error"
	SecurityBaselineSettingStateStateNotApplicable SecurityBaselineSettingStateState = "notApplicable"
	SecurityBaselineSettingStateStateNotSecure     SecurityBaselineSettingStateState = "notSecure"
	SecurityBaselineSettingStateStateSecure        SecurityBaselineSettingStateState = "secure"
	SecurityBaselineSettingStateStateUnknown       SecurityBaselineSettingStateState = "unknown"
)

func PossibleValuesForSecurityBaselineSettingStateState() []string {
	return []string{
		string(SecurityBaselineSettingStateStateConflict),
		string(SecurityBaselineSettingStateStateError),
		string(SecurityBaselineSettingStateStateNotApplicable),
		string(SecurityBaselineSettingStateStateNotSecure),
		string(SecurityBaselineSettingStateStateSecure),
		string(SecurityBaselineSettingStateStateUnknown),
	}
}

func (s *SecurityBaselineSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineSettingStateState(input string) (*SecurityBaselineSettingStateState, error) {
	vals := map[string]SecurityBaselineSettingStateState{
		"conflict":      SecurityBaselineSettingStateStateConflict,
		"error":         SecurityBaselineSettingStateStateError,
		"notapplicable": SecurityBaselineSettingStateStateNotApplicable,
		"notsecure":     SecurityBaselineSettingStateStateNotSecure,
		"secure":        SecurityBaselineSettingStateStateSecure,
		"unknown":       SecurityBaselineSettingStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineSettingStateState(input)
	return &out, nil
}

type SecurityBaselineStateState string

const (
	SecurityBaselineStateStateConflict      SecurityBaselineStateState = "conflict"
	SecurityBaselineStateStateError         SecurityBaselineStateState = "error"
	SecurityBaselineStateStateNotApplicable SecurityBaselineStateState = "notApplicable"
	SecurityBaselineStateStateNotSecure     SecurityBaselineStateState = "notSecure"
	SecurityBaselineStateStateSecure        SecurityBaselineStateState = "secure"
	SecurityBaselineStateStateUnknown       SecurityBaselineStateState = "unknown"
)

func PossibleValuesForSecurityBaselineStateState() []string {
	return []string{
		string(SecurityBaselineStateStateConflict),
		string(SecurityBaselineStateStateError),
		string(SecurityBaselineStateStateNotApplicable),
		string(SecurityBaselineStateStateNotSecure),
		string(SecurityBaselineStateStateSecure),
		string(SecurityBaselineStateStateUnknown),
	}
}

func (s *SecurityBaselineStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineStateState(input string) (*SecurityBaselineStateState, error) {
	vals := map[string]SecurityBaselineStateState{
		"conflict":      SecurityBaselineStateStateConflict,
		"error":         SecurityBaselineStateStateError,
		"notapplicable": SecurityBaselineStateStateNotApplicable,
		"notsecure":     SecurityBaselineStateStateNotSecure,
		"secure":        SecurityBaselineStateStateSecure,
		"unknown":       SecurityBaselineStateStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineStateState(input)
	return &out, nil
}

type SensitivityLabelApplicableTo string

const (
	SensitivityLabelApplicableToEmail        SensitivityLabelApplicableTo = "email"
	SensitivityLabelApplicableToSite         SensitivityLabelApplicableTo = "site"
	SensitivityLabelApplicableToTeamwork     SensitivityLabelApplicableTo = "teamwork"
	SensitivityLabelApplicableToUnifiedGroup SensitivityLabelApplicableTo = "unifiedGroup"
)

func PossibleValuesForSensitivityLabelApplicableTo() []string {
	return []string{
		string(SensitivityLabelApplicableToEmail),
		string(SensitivityLabelApplicableToSite),
		string(SensitivityLabelApplicableToTeamwork),
		string(SensitivityLabelApplicableToUnifiedGroup),
	}
}

func (s *SensitivityLabelApplicableTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelApplicableTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelApplicableTo(input string) (*SensitivityLabelApplicableTo, error) {
	vals := map[string]SensitivityLabelApplicableTo{
		"email":        SensitivityLabelApplicableToEmail,
		"site":         SensitivityLabelApplicableToSite,
		"teamwork":     SensitivityLabelApplicableToTeamwork,
		"unifiedgroup": SensitivityLabelApplicableToUnifiedGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelApplicableTo(input)
	return &out, nil
}

type SensitivityLabelApplicationMode string

const (
	SensitivityLabelApplicationModeAutomatic   SensitivityLabelApplicationMode = "automatic"
	SensitivityLabelApplicationModeManual      SensitivityLabelApplicationMode = "manual"
	SensitivityLabelApplicationModeRecommended SensitivityLabelApplicationMode = "recommended"
)

func PossibleValuesForSensitivityLabelApplicationMode() []string {
	return []string{
		string(SensitivityLabelApplicationModeAutomatic),
		string(SensitivityLabelApplicationModeManual),
		string(SensitivityLabelApplicationModeRecommended),
	}
}

func (s *SensitivityLabelApplicationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelApplicationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelApplicationMode(input string) (*SensitivityLabelApplicationMode, error) {
	vals := map[string]SensitivityLabelApplicationMode{
		"automatic":   SensitivityLabelApplicationModeAutomatic,
		"manual":      SensitivityLabelApplicationModeManual,
		"recommended": SensitivityLabelApplicationModeRecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelApplicationMode(input)
	return &out, nil
}

type SensitivityPolicySettingsApplicableTo string

const (
	SensitivityPolicySettingsApplicableToEmail        SensitivityPolicySettingsApplicableTo = "email"
	SensitivityPolicySettingsApplicableToSite         SensitivityPolicySettingsApplicableTo = "site"
	SensitivityPolicySettingsApplicableToTeamwork     SensitivityPolicySettingsApplicableTo = "teamwork"
	SensitivityPolicySettingsApplicableToUnifiedGroup SensitivityPolicySettingsApplicableTo = "unifiedGroup"
)

func PossibleValuesForSensitivityPolicySettingsApplicableTo() []string {
	return []string{
		string(SensitivityPolicySettingsApplicableToEmail),
		string(SensitivityPolicySettingsApplicableToSite),
		string(SensitivityPolicySettingsApplicableToTeamwork),
		string(SensitivityPolicySettingsApplicableToUnifiedGroup),
	}
}

func (s *SensitivityPolicySettingsApplicableTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityPolicySettingsApplicableTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityPolicySettingsApplicableTo(input string) (*SensitivityPolicySettingsApplicableTo, error) {
	vals := map[string]SensitivityPolicySettingsApplicableTo{
		"email":        SensitivityPolicySettingsApplicableToEmail,
		"site":         SensitivityPolicySettingsApplicableToSite,
		"teamwork":     SensitivityPolicySettingsApplicableToTeamwork,
		"unifiedgroup": SensitivityPolicySettingsApplicableToUnifiedGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityPolicySettingsApplicableTo(input)
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

type SignInPreferencesUserPreferredMethodForSecondaryAuthentication string

const (
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationOath                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "oath"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationPush                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "push"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationSms                  SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "sms"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceAlternateMobile SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "voiceAlternateMobile"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceMobile          SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "voiceMobile"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceOffice          SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "voiceOffice"
)

func PossibleValuesForSignInPreferencesUserPreferredMethodForSecondaryAuthentication() []string {
	return []string{
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationOath),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationPush),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationSms),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceAlternateMobile),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceMobile),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceOffice),
	}
}

func (s *SignInPreferencesUserPreferredMethodForSecondaryAuthentication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInPreferencesUserPreferredMethodForSecondaryAuthentication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInPreferencesUserPreferredMethodForSecondaryAuthentication(input string) (*SignInPreferencesUserPreferredMethodForSecondaryAuthentication, error) {
	vals := map[string]SignInPreferencesUserPreferredMethodForSecondaryAuthentication{
		"oath":                 SignInPreferencesUserPreferredMethodForSecondaryAuthenticationOath,
		"push":                 SignInPreferencesUserPreferredMethodForSecondaryAuthenticationPush,
		"sms":                  SignInPreferencesUserPreferredMethodForSecondaryAuthenticationSms,
		"voicealternatemobile": SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceAlternateMobile,
		"voicemobile":          SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceMobile,
		"voiceoffice":          SignInPreferencesUserPreferredMethodForSecondaryAuthenticationVoiceOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInPreferencesUserPreferredMethodForSecondaryAuthentication(input)
	return &out, nil
}

type SkillProficiencyAllowedAudiences string

const (
	SkillProficiencyAllowedAudiencesContacts               SkillProficiencyAllowedAudiences = "contacts"
	SkillProficiencyAllowedAudiencesEveryone               SkillProficiencyAllowedAudiences = "everyone"
	SkillProficiencyAllowedAudiencesFamily                 SkillProficiencyAllowedAudiences = "family"
	SkillProficiencyAllowedAudiencesFederatedOrganizations SkillProficiencyAllowedAudiences = "federatedOrganizations"
	SkillProficiencyAllowedAudiencesGroupMembers           SkillProficiencyAllowedAudiences = "groupMembers"
	SkillProficiencyAllowedAudiencesMe                     SkillProficiencyAllowedAudiences = "me"
	SkillProficiencyAllowedAudiencesOrganization           SkillProficiencyAllowedAudiences = "organization"
)

func PossibleValuesForSkillProficiencyAllowedAudiences() []string {
	return []string{
		string(SkillProficiencyAllowedAudiencesContacts),
		string(SkillProficiencyAllowedAudiencesEveryone),
		string(SkillProficiencyAllowedAudiencesFamily),
		string(SkillProficiencyAllowedAudiencesFederatedOrganizations),
		string(SkillProficiencyAllowedAudiencesGroupMembers),
		string(SkillProficiencyAllowedAudiencesMe),
		string(SkillProficiencyAllowedAudiencesOrganization),
	}
}

func (s *SkillProficiencyAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkillProficiencyAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkillProficiencyAllowedAudiences(input string) (*SkillProficiencyAllowedAudiences, error) {
	vals := map[string]SkillProficiencyAllowedAudiences{
		"contacts":               SkillProficiencyAllowedAudiencesContacts,
		"everyone":               SkillProficiencyAllowedAudiencesEveryone,
		"family":                 SkillProficiencyAllowedAudiencesFamily,
		"federatedorganizations": SkillProficiencyAllowedAudiencesFederatedOrganizations,
		"groupmembers":           SkillProficiencyAllowedAudiencesGroupMembers,
		"me":                     SkillProficiencyAllowedAudiencesMe,
		"organization":           SkillProficiencyAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkillProficiencyAllowedAudiences(input)
	return &out, nil
}

type SkillProficiencyProficiency string

const (
	SkillProficiencyProficiencyAdvancedProfessional SkillProficiencyProficiency = "advancedProfessional"
	SkillProficiencyProficiencyElementary           SkillProficiencyProficiency = "elementary"
	SkillProficiencyProficiencyExpert               SkillProficiencyProficiency = "expert"
	SkillProficiencyProficiencyGeneralProfessional  SkillProficiencyProficiency = "generalProfessional"
	SkillProficiencyProficiencyLimitedWorking       SkillProficiencyProficiency = "limitedWorking"
)

func PossibleValuesForSkillProficiencyProficiency() []string {
	return []string{
		string(SkillProficiencyProficiencyAdvancedProfessional),
		string(SkillProficiencyProficiencyElementary),
		string(SkillProficiencyProficiencyExpert),
		string(SkillProficiencyProficiencyGeneralProfessional),
		string(SkillProficiencyProficiencyLimitedWorking),
	}
}

func (s *SkillProficiencyProficiency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkillProficiencyProficiency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkillProficiencyProficiency(input string) (*SkillProficiencyProficiency, error) {
	vals := map[string]SkillProficiencyProficiency{
		"advancedprofessional": SkillProficiencyProficiencyAdvancedProfessional,
		"elementary":           SkillProficiencyProficiencyElementary,
		"expert":               SkillProficiencyProficiencyExpert,
		"generalprofessional":  SkillProficiencyProficiencyGeneralProfessional,
		"limitedworking":       SkillProficiencyProficiencyLimitedWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkillProficiencyProficiency(input)
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

type SynchronizationMetadataEntryKey string

const (
	SynchronizationMetadataEntryKeyConfigurationFields                      SynchronizationMetadataEntryKey = "configurationFields"
	SynchronizationMetadataEntryKeyGalleryApplicationIdentifier             SynchronizationMetadataEntryKey = "galleryApplicationIdentifier"
	SynchronizationMetadataEntryKeyGalleryApplicationKey                    SynchronizationMetadataEntryKey = "galleryApplicationKey"
	SynchronizationMetadataEntryKeyIsOAuthEnabled                           SynchronizationMetadataEntryKey = "isOAuthEnabled"
	SynchronizationMetadataEntryKeyIsSynchronizationAgentAssignmentRequired SynchronizationMetadataEntryKey = "IsSynchronizationAgentAssignmentRequired"
	SynchronizationMetadataEntryKeyIsSynchronizationAgentRequired           SynchronizationMetadataEntryKey = "isSynchronizationAgentRequired"
	SynchronizationMetadataEntryKeyIsSynchronizationInPreview               SynchronizationMetadataEntryKey = "isSynchronizationInPreview"
	SynchronizationMetadataEntryKeyOAuthSettings                            SynchronizationMetadataEntryKey = "oAuthSettings"
	SynchronizationMetadataEntryKeySynchronizationLearnMoreIbizaFwLink      SynchronizationMetadataEntryKey = "synchronizationLearnMoreIbizaFwLink"
)

func PossibleValuesForSynchronizationMetadataEntryKey() []string {
	return []string{
		string(SynchronizationMetadataEntryKeyConfigurationFields),
		string(SynchronizationMetadataEntryKeyGalleryApplicationIdentifier),
		string(SynchronizationMetadataEntryKeyGalleryApplicationKey),
		string(SynchronizationMetadataEntryKeyIsOAuthEnabled),
		string(SynchronizationMetadataEntryKeyIsSynchronizationAgentAssignmentRequired),
		string(SynchronizationMetadataEntryKeyIsSynchronizationAgentRequired),
		string(SynchronizationMetadataEntryKeyIsSynchronizationInPreview),
		string(SynchronizationMetadataEntryKeyOAuthSettings),
		string(SynchronizationMetadataEntryKeySynchronizationLearnMoreIbizaFwLink),
	}
}

func (s *SynchronizationMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationMetadataEntryKey(input string) (*SynchronizationMetadataEntryKey, error) {
	vals := map[string]SynchronizationMetadataEntryKey{
		"configurationfields":                      SynchronizationMetadataEntryKeyConfigurationFields,
		"galleryapplicationidentifier":             SynchronizationMetadataEntryKeyGalleryApplicationIdentifier,
		"galleryapplicationkey":                    SynchronizationMetadataEntryKeyGalleryApplicationKey,
		"isoauthenabled":                           SynchronizationMetadataEntryKeyIsOAuthEnabled,
		"issynchronizationagentassignmentrequired": SynchronizationMetadataEntryKeyIsSynchronizationAgentAssignmentRequired,
		"issynchronizationagentrequired":           SynchronizationMetadataEntryKeyIsSynchronizationAgentRequired,
		"issynchronizationinpreview":               SynchronizationMetadataEntryKeyIsSynchronizationInPreview,
		"oauthsettings":                            SynchronizationMetadataEntryKeyOAuthSettings,
		"synchronizationlearnmoreibizafwlink":      SynchronizationMetadataEntryKeySynchronizationLearnMoreIbizaFwLink,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationMetadataEntryKey(input)
	return &out, nil
}

type SynchronizationQuarantineReason string

const (
	SynchronizationQuarantineReasonEncounteredBaseEscrowThreshold       SynchronizationQuarantineReason = "EncounteredBaseEscrowThreshold"
	SynchronizationQuarantineReasonEncounteredEscrowProportionThreshold SynchronizationQuarantineReason = "EncounteredEscrowProportionThreshold"
	SynchronizationQuarantineReasonEncounteredQuarantineException       SynchronizationQuarantineReason = "EncounteredQuarantineException"
	SynchronizationQuarantineReasonEncounteredTotalEscrowThreshold      SynchronizationQuarantineReason = "EncounteredTotalEscrowThreshold"
	SynchronizationQuarantineReasonIngestionInterrupted                 SynchronizationQuarantineReason = "IngestionInterrupted"
	SynchronizationQuarantineReasonQuarantinedOnDemand                  SynchronizationQuarantineReason = "QuarantinedOnDemand"
	SynchronizationQuarantineReasonTooManyDeletes                       SynchronizationQuarantineReason = "TooManyDeletes"
	SynchronizationQuarantineReasonUnknown                              SynchronizationQuarantineReason = "Unknown"
)

func PossibleValuesForSynchronizationQuarantineReason() []string {
	return []string{
		string(SynchronizationQuarantineReasonEncounteredBaseEscrowThreshold),
		string(SynchronizationQuarantineReasonEncounteredEscrowProportionThreshold),
		string(SynchronizationQuarantineReasonEncounteredQuarantineException),
		string(SynchronizationQuarantineReasonEncounteredTotalEscrowThreshold),
		string(SynchronizationQuarantineReasonIngestionInterrupted),
		string(SynchronizationQuarantineReasonQuarantinedOnDemand),
		string(SynchronizationQuarantineReasonTooManyDeletes),
		string(SynchronizationQuarantineReasonUnknown),
	}
}

func (s *SynchronizationQuarantineReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationQuarantineReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationQuarantineReason(input string) (*SynchronizationQuarantineReason, error) {
	vals := map[string]SynchronizationQuarantineReason{
		"encounteredbaseescrowthreshold":       SynchronizationQuarantineReasonEncounteredBaseEscrowThreshold,
		"encounteredescrowproportionthreshold": SynchronizationQuarantineReasonEncounteredEscrowProportionThreshold,
		"encounteredquarantineexception":       SynchronizationQuarantineReasonEncounteredQuarantineException,
		"encounteredtotalescrowthreshold":      SynchronizationQuarantineReasonEncounteredTotalEscrowThreshold,
		"ingestioninterrupted":                 SynchronizationQuarantineReasonIngestionInterrupted,
		"quarantinedondemand":                  SynchronizationQuarantineReasonQuarantinedOnDemand,
		"toomanydeletes":                       SynchronizationQuarantineReasonTooManyDeletes,
		"unknown":                              SynchronizationQuarantineReasonUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationQuarantineReason(input)
	return &out, nil
}

type SynchronizationScheduleState string

const (
	SynchronizationScheduleStateActive   SynchronizationScheduleState = "Active"
	SynchronizationScheduleStateDisabled SynchronizationScheduleState = "Disabled"
	SynchronizationScheduleStatePaused   SynchronizationScheduleState = "Paused"
)

func PossibleValuesForSynchronizationScheduleState() []string {
	return []string{
		string(SynchronizationScheduleStateActive),
		string(SynchronizationScheduleStateDisabled),
		string(SynchronizationScheduleStatePaused),
	}
}

func (s *SynchronizationScheduleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationScheduleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationScheduleState(input string) (*SynchronizationScheduleState, error) {
	vals := map[string]SynchronizationScheduleState{
		"active":   SynchronizationScheduleStateActive,
		"disabled": SynchronizationScheduleStateDisabled,
		"paused":   SynchronizationScheduleStatePaused,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationScheduleState(input)
	return &out, nil
}

type SynchronizationSecretKeyStringValuePairKey string

const (
	SynchronizationSecretKeyStringValuePairKeyAppKey                          SynchronizationSecretKeyStringValuePairKey = "AppKey"
	SynchronizationSecretKeyStringValuePairKeyApplicationTemplateIdentifier   SynchronizationSecretKeyStringValuePairKey = "ApplicationTemplateIdentifier"
	SynchronizationSecretKeyStringValuePairKeyAuthenticationType              SynchronizationSecretKeyStringValuePairKey = "AuthenticationType"
	SynchronizationSecretKeyStringValuePairKeyBaseAddress                     SynchronizationSecretKeyStringValuePairKey = "BaseAddress"
	SynchronizationSecretKeyStringValuePairKeyClientIdentifier                SynchronizationSecretKeyStringValuePairKey = "ClientIdentifier"
	SynchronizationSecretKeyStringValuePairKeyClientSecret                    SynchronizationSecretKeyStringValuePairKey = "ClientSecret"
	SynchronizationSecretKeyStringValuePairKeyCompanyId                       SynchronizationSecretKeyStringValuePairKey = "CompanyId"
	SynchronizationSecretKeyStringValuePairKeyConnectionString                SynchronizationSecretKeyStringValuePairKey = "ConnectionString"
	SynchronizationSecretKeyStringValuePairKeyConsumerKey                     SynchronizationSecretKeyStringValuePairKey = "ConsumerKey"
	SynchronizationSecretKeyStringValuePairKeyConsumerSecret                  SynchronizationSecretKeyStringValuePairKey = "ConsumerSecret"
	SynchronizationSecretKeyStringValuePairKeyDomain                          SynchronizationSecretKeyStringValuePairKey = "Domain"
	SynchronizationSecretKeyStringValuePairKeyEnforceDomain                   SynchronizationSecretKeyStringValuePairKey = "EnforceDomain"
	SynchronizationSecretKeyStringValuePairKeyHardDeletesEnabled              SynchronizationSecretKeyStringValuePairKey = "HardDeletesEnabled"
	SynchronizationSecretKeyStringValuePairKeyInstanceName                    SynchronizationSecretKeyStringValuePairKey = "InstanceName"
	SynchronizationSecretKeyStringValuePairKeyNone                            SynchronizationSecretKeyStringValuePairKey = "None"
	SynchronizationSecretKeyStringValuePairKeyOauth2AccessToken               SynchronizationSecretKeyStringValuePairKey = "Oauth2AccessToken"
	SynchronizationSecretKeyStringValuePairKeyOauth2AccessTokenCreationTime   SynchronizationSecretKeyStringValuePairKey = "Oauth2AccessTokenCreationTime"
	SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationCode         SynchronizationSecretKeyStringValuePairKey = "Oauth2AuthorizationCode"
	SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationUri          SynchronizationSecretKeyStringValuePairKey = "Oauth2AuthorizationUri"
	SynchronizationSecretKeyStringValuePairKeyOauth2ClientId                  SynchronizationSecretKeyStringValuePairKey = "Oauth2ClientId"
	SynchronizationSecretKeyStringValuePairKeyOauth2ClientSecret              SynchronizationSecretKeyStringValuePairKey = "Oauth2ClientSecret"
	SynchronizationSecretKeyStringValuePairKeyOauth2RedirectUri               SynchronizationSecretKeyStringValuePairKey = "Oauth2RedirectUri"
	SynchronizationSecretKeyStringValuePairKeyOauth2RefreshToken              SynchronizationSecretKeyStringValuePairKey = "Oauth2RefreshToken"
	SynchronizationSecretKeyStringValuePairKeyOauth2TokenExchangeUri          SynchronizationSecretKeyStringValuePairKey = "Oauth2TokenExchangeUri"
	SynchronizationSecretKeyStringValuePairKeyPassword                        SynchronizationSecretKeyStringValuePairKey = "Password"
	SynchronizationSecretKeyStringValuePairKeyPerformInboundEntitlementGrants SynchronizationSecretKeyStringValuePairKey = "PerformInboundEntitlementGrants"
	SynchronizationSecretKeyStringValuePairKeySandbox                         SynchronizationSecretKeyStringValuePairKey = "Sandbox"
	SynchronizationSecretKeyStringValuePairKeySandboxName                     SynchronizationSecretKeyStringValuePairKey = "SandboxName"
	SynchronizationSecretKeyStringValuePairKeySecretToken                     SynchronizationSecretKeyStringValuePairKey = "SecretToken"
	SynchronizationSecretKeyStringValuePairKeyServer                          SynchronizationSecretKeyStringValuePairKey = "Server"
	SynchronizationSecretKeyStringValuePairKeySingleSignOnType                SynchronizationSecretKeyStringValuePairKey = "SingleSignOnType"
	SynchronizationSecretKeyStringValuePairKeySkipOutOfScopeDeletions         SynchronizationSecretKeyStringValuePairKey = "SkipOutOfScopeDeletions"
	SynchronizationSecretKeyStringValuePairKeySyncAgentADContainer            SynchronizationSecretKeyStringValuePairKey = "SyncAgentADContainer"
	SynchronizationSecretKeyStringValuePairKeySyncAgentCompatibilityKey       SynchronizationSecretKeyStringValuePairKey = "SyncAgentCompatibilityKey"
	SynchronizationSecretKeyStringValuePairKeySyncAll                         SynchronizationSecretKeyStringValuePairKey = "SyncAll"
	SynchronizationSecretKeyStringValuePairKeySyncNotificationSettings        SynchronizationSecretKeyStringValuePairKey = "SyncNotificationSettings"
	SynchronizationSecretKeyStringValuePairKeySynchronizationSchedule         SynchronizationSecretKeyStringValuePairKey = "SynchronizationSchedule"
	SynchronizationSecretKeyStringValuePairKeySystemOfRecord                  SynchronizationSecretKeyStringValuePairKey = "SystemOfRecord"
	SynchronizationSecretKeyStringValuePairKeyTestReferences                  SynchronizationSecretKeyStringValuePairKey = "TestReferences"
	SynchronizationSecretKeyStringValuePairKeyTokenExpiration                 SynchronizationSecretKeyStringValuePairKey = "TokenExpiration"
	SynchronizationSecretKeyStringValuePairKeyTokenKey                        SynchronizationSecretKeyStringValuePairKey = "TokenKey"
	SynchronizationSecretKeyStringValuePairKeyUpdateKeyOnSoftDelete           SynchronizationSecretKeyStringValuePairKey = "UpdateKeyOnSoftDelete"
	SynchronizationSecretKeyStringValuePairKeyUrl                             SynchronizationSecretKeyStringValuePairKey = "Url"
	SynchronizationSecretKeyStringValuePairKeyUserName                        SynchronizationSecretKeyStringValuePairKey = "UserName"
	SynchronizationSecretKeyStringValuePairKeyValidateDomain                  SynchronizationSecretKeyStringValuePairKey = "ValidateDomain"
)

func PossibleValuesForSynchronizationSecretKeyStringValuePairKey() []string {
	return []string{
		string(SynchronizationSecretKeyStringValuePairKeyAppKey),
		string(SynchronizationSecretKeyStringValuePairKeyApplicationTemplateIdentifier),
		string(SynchronizationSecretKeyStringValuePairKeyAuthenticationType),
		string(SynchronizationSecretKeyStringValuePairKeyBaseAddress),
		string(SynchronizationSecretKeyStringValuePairKeyClientIdentifier),
		string(SynchronizationSecretKeyStringValuePairKeyClientSecret),
		string(SynchronizationSecretKeyStringValuePairKeyCompanyId),
		string(SynchronizationSecretKeyStringValuePairKeyConnectionString),
		string(SynchronizationSecretKeyStringValuePairKeyConsumerKey),
		string(SynchronizationSecretKeyStringValuePairKeyConsumerSecret),
		string(SynchronizationSecretKeyStringValuePairKeyDomain),
		string(SynchronizationSecretKeyStringValuePairKeyEnforceDomain),
		string(SynchronizationSecretKeyStringValuePairKeyHardDeletesEnabled),
		string(SynchronizationSecretKeyStringValuePairKeyInstanceName),
		string(SynchronizationSecretKeyStringValuePairKeyNone),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AccessToken),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AccessTokenCreationTime),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationCode),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationUri),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2ClientId),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2ClientSecret),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2RedirectUri),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2RefreshToken),
		string(SynchronizationSecretKeyStringValuePairKeyOauth2TokenExchangeUri),
		string(SynchronizationSecretKeyStringValuePairKeyPassword),
		string(SynchronizationSecretKeyStringValuePairKeyPerformInboundEntitlementGrants),
		string(SynchronizationSecretKeyStringValuePairKeySandbox),
		string(SynchronizationSecretKeyStringValuePairKeySandboxName),
		string(SynchronizationSecretKeyStringValuePairKeySecretToken),
		string(SynchronizationSecretKeyStringValuePairKeyServer),
		string(SynchronizationSecretKeyStringValuePairKeySingleSignOnType),
		string(SynchronizationSecretKeyStringValuePairKeySkipOutOfScopeDeletions),
		string(SynchronizationSecretKeyStringValuePairKeySyncAgentADContainer),
		string(SynchronizationSecretKeyStringValuePairKeySyncAgentCompatibilityKey),
		string(SynchronizationSecretKeyStringValuePairKeySyncAll),
		string(SynchronizationSecretKeyStringValuePairKeySyncNotificationSettings),
		string(SynchronizationSecretKeyStringValuePairKeySynchronizationSchedule),
		string(SynchronizationSecretKeyStringValuePairKeySystemOfRecord),
		string(SynchronizationSecretKeyStringValuePairKeyTestReferences),
		string(SynchronizationSecretKeyStringValuePairKeyTokenExpiration),
		string(SynchronizationSecretKeyStringValuePairKeyTokenKey),
		string(SynchronizationSecretKeyStringValuePairKeyUpdateKeyOnSoftDelete),
		string(SynchronizationSecretKeyStringValuePairKeyUrl),
		string(SynchronizationSecretKeyStringValuePairKeyUserName),
		string(SynchronizationSecretKeyStringValuePairKeyValidateDomain),
	}
}

func (s *SynchronizationSecretKeyStringValuePairKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationSecretKeyStringValuePairKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationSecretKeyStringValuePairKey(input string) (*SynchronizationSecretKeyStringValuePairKey, error) {
	vals := map[string]SynchronizationSecretKeyStringValuePairKey{
		"appkey":                          SynchronizationSecretKeyStringValuePairKeyAppKey,
		"applicationtemplateidentifier":   SynchronizationSecretKeyStringValuePairKeyApplicationTemplateIdentifier,
		"authenticationtype":              SynchronizationSecretKeyStringValuePairKeyAuthenticationType,
		"baseaddress":                     SynchronizationSecretKeyStringValuePairKeyBaseAddress,
		"clientidentifier":                SynchronizationSecretKeyStringValuePairKeyClientIdentifier,
		"clientsecret":                    SynchronizationSecretKeyStringValuePairKeyClientSecret,
		"companyid":                       SynchronizationSecretKeyStringValuePairKeyCompanyId,
		"connectionstring":                SynchronizationSecretKeyStringValuePairKeyConnectionString,
		"consumerkey":                     SynchronizationSecretKeyStringValuePairKeyConsumerKey,
		"consumersecret":                  SynchronizationSecretKeyStringValuePairKeyConsumerSecret,
		"domain":                          SynchronizationSecretKeyStringValuePairKeyDomain,
		"enforcedomain":                   SynchronizationSecretKeyStringValuePairKeyEnforceDomain,
		"harddeletesenabled":              SynchronizationSecretKeyStringValuePairKeyHardDeletesEnabled,
		"instancename":                    SynchronizationSecretKeyStringValuePairKeyInstanceName,
		"none":                            SynchronizationSecretKeyStringValuePairKeyNone,
		"oauth2accesstoken":               SynchronizationSecretKeyStringValuePairKeyOauth2AccessToken,
		"oauth2accesstokencreationtime":   SynchronizationSecretKeyStringValuePairKeyOauth2AccessTokenCreationTime,
		"oauth2authorizationcode":         SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationCode,
		"oauth2authorizationuri":          SynchronizationSecretKeyStringValuePairKeyOauth2AuthorizationUri,
		"oauth2clientid":                  SynchronizationSecretKeyStringValuePairKeyOauth2ClientId,
		"oauth2clientsecret":              SynchronizationSecretKeyStringValuePairKeyOauth2ClientSecret,
		"oauth2redirecturi":               SynchronizationSecretKeyStringValuePairKeyOauth2RedirectUri,
		"oauth2refreshtoken":              SynchronizationSecretKeyStringValuePairKeyOauth2RefreshToken,
		"oauth2tokenexchangeuri":          SynchronizationSecretKeyStringValuePairKeyOauth2TokenExchangeUri,
		"password":                        SynchronizationSecretKeyStringValuePairKeyPassword,
		"performinboundentitlementgrants": SynchronizationSecretKeyStringValuePairKeyPerformInboundEntitlementGrants,
		"sandbox":                         SynchronizationSecretKeyStringValuePairKeySandbox,
		"sandboxname":                     SynchronizationSecretKeyStringValuePairKeySandboxName,
		"secrettoken":                     SynchronizationSecretKeyStringValuePairKeySecretToken,
		"server":                          SynchronizationSecretKeyStringValuePairKeyServer,
		"singlesignontype":                SynchronizationSecretKeyStringValuePairKeySingleSignOnType,
		"skipoutofscopedeletions":         SynchronizationSecretKeyStringValuePairKeySkipOutOfScopeDeletions,
		"syncagentadcontainer":            SynchronizationSecretKeyStringValuePairKeySyncAgentADContainer,
		"syncagentcompatibilitykey":       SynchronizationSecretKeyStringValuePairKeySyncAgentCompatibilityKey,
		"syncall":                         SynchronizationSecretKeyStringValuePairKeySyncAll,
		"syncnotificationsettings":        SynchronizationSecretKeyStringValuePairKeySyncNotificationSettings,
		"synchronizationschedule":         SynchronizationSecretKeyStringValuePairKeySynchronizationSchedule,
		"systemofrecord":                  SynchronizationSecretKeyStringValuePairKeySystemOfRecord,
		"testreferences":                  SynchronizationSecretKeyStringValuePairKeyTestReferences,
		"tokenexpiration":                 SynchronizationSecretKeyStringValuePairKeyTokenExpiration,
		"tokenkey":                        SynchronizationSecretKeyStringValuePairKeyTokenKey,
		"updatekeyonsoftdelete":           SynchronizationSecretKeyStringValuePairKeyUpdateKeyOnSoftDelete,
		"url":                             SynchronizationSecretKeyStringValuePairKeyUrl,
		"username":                        SynchronizationSecretKeyStringValuePairKeyUserName,
		"validatedomain":                  SynchronizationSecretKeyStringValuePairKeyValidateDomain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationSecretKeyStringValuePairKey(input)
	return &out, nil
}

type SynchronizationStatusCode string

const (
	SynchronizationStatusCodeActive        SynchronizationStatusCode = "Active"
	SynchronizationStatusCodeNotConfigured SynchronizationStatusCode = "NotConfigured"
	SynchronizationStatusCodeNotRun        SynchronizationStatusCode = "NotRun"
	SynchronizationStatusCodePaused        SynchronizationStatusCode = "Paused"
	SynchronizationStatusCodeQuarantine    SynchronizationStatusCode = "Quarantine"
)

func PossibleValuesForSynchronizationStatusCode() []string {
	return []string{
		string(SynchronizationStatusCodeActive),
		string(SynchronizationStatusCodeNotConfigured),
		string(SynchronizationStatusCodeNotRun),
		string(SynchronizationStatusCodePaused),
		string(SynchronizationStatusCodeQuarantine),
	}
}

func (s *SynchronizationStatusCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationStatusCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationStatusCode(input string) (*SynchronizationStatusCode, error) {
	vals := map[string]SynchronizationStatusCode{
		"active":        SynchronizationStatusCodeActive,
		"notconfigured": SynchronizationStatusCodeNotConfigured,
		"notrun":        SynchronizationStatusCodeNotRun,
		"paused":        SynchronizationStatusCodePaused,
		"quarantine":    SynchronizationStatusCodeQuarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationStatusCode(input)
	return &out, nil
}

type SynchronizationTaskExecutionState string

const (
	SynchronizationTaskExecutionStateEntryLevelErrors SynchronizationTaskExecutionState = "EntryLevelErrors"
	SynchronizationTaskExecutionStateFailed           SynchronizationTaskExecutionState = "Failed"
	SynchronizationTaskExecutionStateSucceeded        SynchronizationTaskExecutionState = "Succeeded"
)

func PossibleValuesForSynchronizationTaskExecutionState() []string {
	return []string{
		string(SynchronizationTaskExecutionStateEntryLevelErrors),
		string(SynchronizationTaskExecutionStateFailed),
		string(SynchronizationTaskExecutionStateSucceeded),
	}
}

func (s *SynchronizationTaskExecutionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationTaskExecutionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationTaskExecutionState(input string) (*SynchronizationTaskExecutionState, error) {
	vals := map[string]SynchronizationTaskExecutionState{
		"entrylevelerrors": SynchronizationTaskExecutionStateEntryLevelErrors,
		"failed":           SynchronizationTaskExecutionStateFailed,
		"succeeded":        SynchronizationTaskExecutionStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationTaskExecutionState(input)
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

type TeamTemplateDefinitionAudience string

const (
	TeamTemplateDefinitionAudienceOrganization TeamTemplateDefinitionAudience = "organization"
	TeamTemplateDefinitionAudiencePublic       TeamTemplateDefinitionAudience = "public"
	TeamTemplateDefinitionAudienceUser         TeamTemplateDefinitionAudience = "user"
)

func PossibleValuesForTeamTemplateDefinitionAudience() []string {
	return []string{
		string(TeamTemplateDefinitionAudienceOrganization),
		string(TeamTemplateDefinitionAudiencePublic),
		string(TeamTemplateDefinitionAudienceUser),
	}
}

func (s *TeamTemplateDefinitionAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamTemplateDefinitionAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamTemplateDefinitionAudience(input string) (*TeamTemplateDefinitionAudience, error) {
	vals := map[string]TeamTemplateDefinitionAudience{
		"organization": TeamTemplateDefinitionAudienceOrganization,
		"public":       TeamTemplateDefinitionAudiencePublic,
		"user":         TeamTemplateDefinitionAudienceUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamTemplateDefinitionAudience(input)
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

type TeamsAppDefinitionAllowedInstallationScopes string

const (
	TeamsAppDefinitionAllowedInstallationScopesGroupChat TeamsAppDefinitionAllowedInstallationScopes = "groupChat"
	TeamsAppDefinitionAllowedInstallationScopesPersonal  TeamsAppDefinitionAllowedInstallationScopes = "personal"
	TeamsAppDefinitionAllowedInstallationScopesTeam      TeamsAppDefinitionAllowedInstallationScopes = "team"
)

func PossibleValuesForTeamsAppDefinitionAllowedInstallationScopes() []string {
	return []string{
		string(TeamsAppDefinitionAllowedInstallationScopesGroupChat),
		string(TeamsAppDefinitionAllowedInstallationScopesPersonal),
		string(TeamsAppDefinitionAllowedInstallationScopesTeam),
	}
}

func (s *TeamsAppDefinitionAllowedInstallationScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDefinitionAllowedInstallationScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDefinitionAllowedInstallationScopes(input string) (*TeamsAppDefinitionAllowedInstallationScopes, error) {
	vals := map[string]TeamsAppDefinitionAllowedInstallationScopes{
		"groupchat": TeamsAppDefinitionAllowedInstallationScopesGroupChat,
		"personal":  TeamsAppDefinitionAllowedInstallationScopesPersonal,
		"team":      TeamsAppDefinitionAllowedInstallationScopesTeam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDefinitionAllowedInstallationScopes(input)
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
	TeamsAsyncOperationOperationTypeCreateChat    TeamsAsyncOperationOperationType = "createChat"
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
		string(TeamsAsyncOperationOperationTypeCreateChat),
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
		"createchat":    TeamsAsyncOperationOperationTypeCreateChat,
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
	TeamworkUserIdentityUserIdentityTypeAadUser                        TeamworkUserIdentityUserIdentityType = "aadUser"
	TeamworkUserIdentityUserIdentityTypeAnonymousGuest                 TeamworkUserIdentityUserIdentityType = "anonymousGuest"
	TeamworkUserIdentityUserIdentityTypeAzureCommunicationServicesUser TeamworkUserIdentityUserIdentityType = "azureCommunicationServicesUser"
	TeamworkUserIdentityUserIdentityTypeEmailUser                      TeamworkUserIdentityUserIdentityType = "emailUser"
	TeamworkUserIdentityUserIdentityTypeFederatedUser                  TeamworkUserIdentityUserIdentityType = "federatedUser"
	TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser               TeamworkUserIdentityUserIdentityType = "onPremiseAadUser"
	TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser   TeamworkUserIdentityUserIdentityType = "personalMicrosoftAccountUser"
	TeamworkUserIdentityUserIdentityTypePhoneUser                      TeamworkUserIdentityUserIdentityType = "phoneUser"
	TeamworkUserIdentityUserIdentityTypeSkypeUser                      TeamworkUserIdentityUserIdentityType = "skypeUser"
)

func PossibleValuesForTeamworkUserIdentityUserIdentityType() []string {
	return []string{
		string(TeamworkUserIdentityUserIdentityTypeAadUser),
		string(TeamworkUserIdentityUserIdentityTypeAnonymousGuest),
		string(TeamworkUserIdentityUserIdentityTypeAzureCommunicationServicesUser),
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
		"aaduser":                        TeamworkUserIdentityUserIdentityTypeAadUser,
		"anonymousguest":                 TeamworkUserIdentityUserIdentityTypeAnonymousGuest,
		"azurecommunicationservicesuser": TeamworkUserIdentityUserIdentityTypeAzureCommunicationServicesUser,
		"emailuser":                      TeamworkUserIdentityUserIdentityTypeEmailUser,
		"federateduser":                  TeamworkUserIdentityUserIdentityTypeFederatedUser,
		"onpremiseaaduser":               TeamworkUserIdentityUserIdentityTypeOnPremiseAadUser,
		"personalmicrosoftaccountuser":   TeamworkUserIdentityUserIdentityTypePersonalMicrosoftAccountUser,
		"phoneuser":                      TeamworkUserIdentityUserIdentityTypePhoneUser,
		"skypeuser":                      TeamworkUserIdentityUserIdentityTypeSkypeUser,
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

type ThreatAssessmentRequestCategory string

const (
	ThreatAssessmentRequestCategoryMalware   ThreatAssessmentRequestCategory = "malware"
	ThreatAssessmentRequestCategoryPhishing  ThreatAssessmentRequestCategory = "phishing"
	ThreatAssessmentRequestCategorySpam      ThreatAssessmentRequestCategory = "spam"
	ThreatAssessmentRequestCategoryUndefined ThreatAssessmentRequestCategory = "undefined"
)

func PossibleValuesForThreatAssessmentRequestCategory() []string {
	return []string{
		string(ThreatAssessmentRequestCategoryMalware),
		string(ThreatAssessmentRequestCategoryPhishing),
		string(ThreatAssessmentRequestCategorySpam),
		string(ThreatAssessmentRequestCategoryUndefined),
	}
}

func (s *ThreatAssessmentRequestCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestCategory(input string) (*ThreatAssessmentRequestCategory, error) {
	vals := map[string]ThreatAssessmentRequestCategory{
		"malware":   ThreatAssessmentRequestCategoryMalware,
		"phishing":  ThreatAssessmentRequestCategoryPhishing,
		"spam":      ThreatAssessmentRequestCategorySpam,
		"undefined": ThreatAssessmentRequestCategoryUndefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestCategory(input)
	return &out, nil
}

type ThreatAssessmentRequestContentType string

const (
	ThreatAssessmentRequestContentTypeFile ThreatAssessmentRequestContentType = "file"
	ThreatAssessmentRequestContentTypeMail ThreatAssessmentRequestContentType = "mail"
	ThreatAssessmentRequestContentTypeUrl  ThreatAssessmentRequestContentType = "url"
)

func PossibleValuesForThreatAssessmentRequestContentType() []string {
	return []string{
		string(ThreatAssessmentRequestContentTypeFile),
		string(ThreatAssessmentRequestContentTypeMail),
		string(ThreatAssessmentRequestContentTypeUrl),
	}
}

func (s *ThreatAssessmentRequestContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestContentType(input string) (*ThreatAssessmentRequestContentType, error) {
	vals := map[string]ThreatAssessmentRequestContentType{
		"file": ThreatAssessmentRequestContentTypeFile,
		"mail": ThreatAssessmentRequestContentTypeMail,
		"url":  ThreatAssessmentRequestContentTypeUrl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestContentType(input)
	return &out, nil
}

type ThreatAssessmentRequestExpectedAssessment string

const (
	ThreatAssessmentRequestExpectedAssessmentBlock   ThreatAssessmentRequestExpectedAssessment = "block"
	ThreatAssessmentRequestExpectedAssessmentUnblock ThreatAssessmentRequestExpectedAssessment = "unblock"
)

func PossibleValuesForThreatAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(ThreatAssessmentRequestExpectedAssessmentBlock),
		string(ThreatAssessmentRequestExpectedAssessmentUnblock),
	}
}

func (s *ThreatAssessmentRequestExpectedAssessment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestExpectedAssessment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestExpectedAssessment(input string) (*ThreatAssessmentRequestExpectedAssessment, error) {
	vals := map[string]ThreatAssessmentRequestExpectedAssessment{
		"block":   ThreatAssessmentRequestExpectedAssessmentBlock,
		"unblock": ThreatAssessmentRequestExpectedAssessmentUnblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestExpectedAssessment(input)
	return &out, nil
}

type ThreatAssessmentRequestRequestSource string

const (
	ThreatAssessmentRequestRequestSourceAdministrator ThreatAssessmentRequestRequestSource = "administrator"
	ThreatAssessmentRequestRequestSourceUndefined     ThreatAssessmentRequestRequestSource = "undefined"
	ThreatAssessmentRequestRequestSourceUser          ThreatAssessmentRequestRequestSource = "user"
)

func PossibleValuesForThreatAssessmentRequestRequestSource() []string {
	return []string{
		string(ThreatAssessmentRequestRequestSourceAdministrator),
		string(ThreatAssessmentRequestRequestSourceUndefined),
		string(ThreatAssessmentRequestRequestSourceUser),
	}
}

func (s *ThreatAssessmentRequestRequestSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestRequestSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestRequestSource(input string) (*ThreatAssessmentRequestRequestSource, error) {
	vals := map[string]ThreatAssessmentRequestRequestSource{
		"administrator": ThreatAssessmentRequestRequestSourceAdministrator,
		"undefined":     ThreatAssessmentRequestRequestSourceUndefined,
		"user":          ThreatAssessmentRequestRequestSourceUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestRequestSource(input)
	return &out, nil
}

type ThreatAssessmentRequestStatus string

const (
	ThreatAssessmentRequestStatusCompleted ThreatAssessmentRequestStatus = "completed"
	ThreatAssessmentRequestStatusPending   ThreatAssessmentRequestStatus = "pending"
)

func PossibleValuesForThreatAssessmentRequestStatus() []string {
	return []string{
		string(ThreatAssessmentRequestStatusCompleted),
		string(ThreatAssessmentRequestStatusPending),
	}
}

func (s *ThreatAssessmentRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestStatus(input string) (*ThreatAssessmentRequestStatus, error) {
	vals := map[string]ThreatAssessmentRequestStatus{
		"completed": ThreatAssessmentRequestStatusCompleted,
		"pending":   ThreatAssessmentRequestStatusPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestStatus(input)
	return &out, nil
}

type ThreatAssessmentResultResultType string

const (
	ThreatAssessmentResultResultTypeCheckPolicy ThreatAssessmentResultResultType = "checkPolicy"
	ThreatAssessmentResultResultTypeRescan      ThreatAssessmentResultResultType = "rescan"
)

func PossibleValuesForThreatAssessmentResultResultType() []string {
	return []string{
		string(ThreatAssessmentResultResultTypeCheckPolicy),
		string(ThreatAssessmentResultResultTypeRescan),
	}
}

func (s *ThreatAssessmentResultResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentResultResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentResultResultType(input string) (*ThreatAssessmentResultResultType, error) {
	vals := map[string]ThreatAssessmentResultResultType{
		"checkpolicy": ThreatAssessmentResultResultTypeCheckPolicy,
		"rescan":      ThreatAssessmentResultResultTypeRescan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentResultResultType(input)
	return &out, nil
}

type TimeCardConfirmedBy string

const (
	TimeCardConfirmedByManager TimeCardConfirmedBy = "manager"
	TimeCardConfirmedByNone    TimeCardConfirmedBy = "none"
	TimeCardConfirmedByUser    TimeCardConfirmedBy = "user"
)

func PossibleValuesForTimeCardConfirmedBy() []string {
	return []string{
		string(TimeCardConfirmedByManager),
		string(TimeCardConfirmedByNone),
		string(TimeCardConfirmedByUser),
	}
}

func (s *TimeCardConfirmedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeCardConfirmedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeCardConfirmedBy(input string) (*TimeCardConfirmedBy, error) {
	vals := map[string]TimeCardConfirmedBy{
		"manager": TimeCardConfirmedByManager,
		"none":    TimeCardConfirmedByNone,
		"user":    TimeCardConfirmedByUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeCardConfirmedBy(input)
	return &out, nil
}

type TimeCardState string

const (
	TimeCardStateClockedIn  TimeCardState = "clockedIn"
	TimeCardStateClockedOut TimeCardState = "clockedOut"
	TimeCardStateOnBreak    TimeCardState = "onBreak"
)

func PossibleValuesForTimeCardState() []string {
	return []string{
		string(TimeCardStateClockedIn),
		string(TimeCardStateClockedOut),
		string(TimeCardStateOnBreak),
	}
}

func (s *TimeCardState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeCardState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeCardState(input string) (*TimeCardState, error) {
	vals := map[string]TimeCardState{
		"clockedin":  TimeCardStateClockedIn,
		"clockedout": TimeCardStateClockedOut,
		"onbreak":    TimeCardStateOnBreak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeCardState(input)
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

type TranslationLanguageOverrideTranslationBehavior string

const (
	TranslationLanguageOverrideTranslationBehaviorAsk TranslationLanguageOverrideTranslationBehavior = "Ask"
	TranslationLanguageOverrideTranslationBehaviorNo  TranslationLanguageOverrideTranslationBehavior = "No"
	TranslationLanguageOverrideTranslationBehaviorYes TranslationLanguageOverrideTranslationBehavior = "Yes"
)

func PossibleValuesForTranslationLanguageOverrideTranslationBehavior() []string {
	return []string{
		string(TranslationLanguageOverrideTranslationBehaviorAsk),
		string(TranslationLanguageOverrideTranslationBehaviorNo),
		string(TranslationLanguageOverrideTranslationBehaviorYes),
	}
}

func (s *TranslationLanguageOverrideTranslationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTranslationLanguageOverrideTranslationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTranslationLanguageOverrideTranslationBehavior(input string) (*TranslationLanguageOverrideTranslationBehavior, error) {
	vals := map[string]TranslationLanguageOverrideTranslationBehavior{
		"ask": TranslationLanguageOverrideTranslationBehaviorAsk,
		"no":  TranslationLanguageOverrideTranslationBehaviorNo,
		"yes": TranslationLanguageOverrideTranslationBehaviorYes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TranslationLanguageOverrideTranslationBehavior(input)
	return &out, nil
}

type TranslationPreferencesTranslationBehavior string

const (
	TranslationPreferencesTranslationBehaviorAsk TranslationPreferencesTranslationBehavior = "Ask"
	TranslationPreferencesTranslationBehaviorNo  TranslationPreferencesTranslationBehavior = "No"
	TranslationPreferencesTranslationBehaviorYes TranslationPreferencesTranslationBehavior = "Yes"
)

func PossibleValuesForTranslationPreferencesTranslationBehavior() []string {
	return []string{
		string(TranslationPreferencesTranslationBehaviorAsk),
		string(TranslationPreferencesTranslationBehaviorNo),
		string(TranslationPreferencesTranslationBehaviorYes),
	}
}

func (s *TranslationPreferencesTranslationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTranslationPreferencesTranslationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTranslationPreferencesTranslationBehavior(input string) (*TranslationPreferencesTranslationBehavior, error) {
	vals := map[string]TranslationPreferencesTranslationBehavior{
		"ask": TranslationPreferencesTranslationBehaviorAsk,
		"no":  TranslationPreferencesTranslationBehaviorNo,
		"yes": TranslationPreferencesTranslationBehaviorYes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TranslationPreferencesTranslationBehavior(input)
	return &out, nil
}

type TypedEmailAddressType string

const (
	TypedEmailAddressTypeMain     TypedEmailAddressType = "main"
	TypedEmailAddressTypeOther    TypedEmailAddressType = "other"
	TypedEmailAddressTypePersonal TypedEmailAddressType = "personal"
	TypedEmailAddressTypeUnknown  TypedEmailAddressType = "unknown"
	TypedEmailAddressTypeWork     TypedEmailAddressType = "work"
)

func PossibleValuesForTypedEmailAddressType() []string {
	return []string{
		string(TypedEmailAddressTypeMain),
		string(TypedEmailAddressTypeOther),
		string(TypedEmailAddressTypePersonal),
		string(TypedEmailAddressTypeUnknown),
		string(TypedEmailAddressTypeWork),
	}
}

func (s *TypedEmailAddressType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTypedEmailAddressType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTypedEmailAddressType(input string) (*TypedEmailAddressType, error) {
	vals := map[string]TypedEmailAddressType{
		"main":     TypedEmailAddressTypeMain,
		"other":    TypedEmailAddressTypeOther,
		"personal": TypedEmailAddressTypePersonal,
		"unknown":  TypedEmailAddressTypeUnknown,
		"work":     TypedEmailAddressTypeWork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TypedEmailAddressType(input)
	return &out, nil
}

type UsageRightState string

const (
	UsageRightStateActive    UsageRightState = "active"
	UsageRightStateInactive  UsageRightState = "inactive"
	UsageRightStateSuspended UsageRightState = "suspended"
	UsageRightStateWarning   UsageRightState = "warning"
)

func PossibleValuesForUsageRightState() []string {
	return []string{
		string(UsageRightStateActive),
		string(UsageRightStateInactive),
		string(UsageRightStateSuspended),
		string(UsageRightStateWarning),
	}
}

func (s *UsageRightState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsageRightState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUsageRightState(input string) (*UsageRightState, error) {
	vals := map[string]UsageRightState{
		"active":    UsageRightStateActive,
		"inactive":  UsageRightStateInactive,
		"suspended": UsageRightStateSuspended,
		"warning":   UsageRightStateWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageRightState(input)
	return &out, nil
}

type UserAccountInformationAllowedAudiences string

const (
	UserAccountInformationAllowedAudiencesContacts               UserAccountInformationAllowedAudiences = "contacts"
	UserAccountInformationAllowedAudiencesEveryone               UserAccountInformationAllowedAudiences = "everyone"
	UserAccountInformationAllowedAudiencesFamily                 UserAccountInformationAllowedAudiences = "family"
	UserAccountInformationAllowedAudiencesFederatedOrganizations UserAccountInformationAllowedAudiences = "federatedOrganizations"
	UserAccountInformationAllowedAudiencesGroupMembers           UserAccountInformationAllowedAudiences = "groupMembers"
	UserAccountInformationAllowedAudiencesMe                     UserAccountInformationAllowedAudiences = "me"
	UserAccountInformationAllowedAudiencesOrganization           UserAccountInformationAllowedAudiences = "organization"
)

func PossibleValuesForUserAccountInformationAllowedAudiences() []string {
	return []string{
		string(UserAccountInformationAllowedAudiencesContacts),
		string(UserAccountInformationAllowedAudiencesEveryone),
		string(UserAccountInformationAllowedAudiencesFamily),
		string(UserAccountInformationAllowedAudiencesFederatedOrganizations),
		string(UserAccountInformationAllowedAudiencesGroupMembers),
		string(UserAccountInformationAllowedAudiencesMe),
		string(UserAccountInformationAllowedAudiencesOrganization),
	}
}

func (s *UserAccountInformationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserAccountInformationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserAccountInformationAllowedAudiences(input string) (*UserAccountInformationAllowedAudiences, error) {
	vals := map[string]UserAccountInformationAllowedAudiences{
		"contacts":               UserAccountInformationAllowedAudiencesContacts,
		"everyone":               UserAccountInformationAllowedAudiencesEveryone,
		"family":                 UserAccountInformationAllowedAudiencesFamily,
		"federatedorganizations": UserAccountInformationAllowedAudiencesFederatedOrganizations,
		"groupmembers":           UserAccountInformationAllowedAudiencesGroupMembers,
		"me":                     UserAccountInformationAllowedAudiencesMe,
		"organization":           UserAccountInformationAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserAccountInformationAllowedAudiences(input)
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

type VirtualEventRegistrationStatus string

const (
	VirtualEventRegistrationStatusCanceled            VirtualEventRegistrationStatus = "canceled"
	VirtualEventRegistrationStatusPendingApproval     VirtualEventRegistrationStatus = "pendingApproval"
	VirtualEventRegistrationStatusRegistered          VirtualEventRegistrationStatus = "registered"
	VirtualEventRegistrationStatusRejectedByOrganizer VirtualEventRegistrationStatus = "rejectedByOrganizer"
	VirtualEventRegistrationStatusWaitlisted          VirtualEventRegistrationStatus = "waitlisted"
)

func PossibleValuesForVirtualEventRegistrationStatus() []string {
	return []string{
		string(VirtualEventRegistrationStatusCanceled),
		string(VirtualEventRegistrationStatusPendingApproval),
		string(VirtualEventRegistrationStatusRegistered),
		string(VirtualEventRegistrationStatusRejectedByOrganizer),
		string(VirtualEventRegistrationStatusWaitlisted),
	}
}

func (s *VirtualEventRegistrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventRegistrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventRegistrationStatus(input string) (*VirtualEventRegistrationStatus, error) {
	vals := map[string]VirtualEventRegistrationStatus{
		"canceled":            VirtualEventRegistrationStatusCanceled,
		"pendingapproval":     VirtualEventRegistrationStatusPendingApproval,
		"registered":          VirtualEventRegistrationStatusRegistered,
		"rejectedbyorganizer": VirtualEventRegistrationStatusRejectedByOrganizer,
		"waitlisted":          VirtualEventRegistrationStatusWaitlisted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventRegistrationStatus(input)
	return &out, nil
}

type VirtualEventSessionAllowMeetingChat string

const (
	VirtualEventSessionAllowMeetingChatDisabled VirtualEventSessionAllowMeetingChat = "disabled"
	VirtualEventSessionAllowMeetingChatEnabled  VirtualEventSessionAllowMeetingChat = "enabled"
	VirtualEventSessionAllowMeetingChatLimited  VirtualEventSessionAllowMeetingChat = "limited"
)

func PossibleValuesForVirtualEventSessionAllowMeetingChat() []string {
	return []string{
		string(VirtualEventSessionAllowMeetingChatDisabled),
		string(VirtualEventSessionAllowMeetingChatEnabled),
		string(VirtualEventSessionAllowMeetingChatLimited),
	}
}

func (s *VirtualEventSessionAllowMeetingChat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionAllowMeetingChat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionAllowMeetingChat(input string) (*VirtualEventSessionAllowMeetingChat, error) {
	vals := map[string]VirtualEventSessionAllowMeetingChat{
		"disabled": VirtualEventSessionAllowMeetingChatDisabled,
		"enabled":  VirtualEventSessionAllowMeetingChatEnabled,
		"limited":  VirtualEventSessionAllowMeetingChatLimited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAllowMeetingChat(input)
	return &out, nil
}

type VirtualEventSessionAllowedPresenters string

const (
	VirtualEventSessionAllowedPresentersEveryone        VirtualEventSessionAllowedPresenters = "everyone"
	VirtualEventSessionAllowedPresentersOrganization    VirtualEventSessionAllowedPresenters = "organization"
	VirtualEventSessionAllowedPresentersOrganizer       VirtualEventSessionAllowedPresenters = "organizer"
	VirtualEventSessionAllowedPresentersRoleIsPresenter VirtualEventSessionAllowedPresenters = "roleIsPresenter"
)

func PossibleValuesForVirtualEventSessionAllowedPresenters() []string {
	return []string{
		string(VirtualEventSessionAllowedPresentersEveryone),
		string(VirtualEventSessionAllowedPresentersOrganization),
		string(VirtualEventSessionAllowedPresentersOrganizer),
		string(VirtualEventSessionAllowedPresentersRoleIsPresenter),
	}
}

func (s *VirtualEventSessionAllowedPresenters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionAllowedPresenters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionAllowedPresenters(input string) (*VirtualEventSessionAllowedPresenters, error) {
	vals := map[string]VirtualEventSessionAllowedPresenters{
		"everyone":        VirtualEventSessionAllowedPresentersEveryone,
		"organization":    VirtualEventSessionAllowedPresentersOrganization,
		"organizer":       VirtualEventSessionAllowedPresentersOrganizer,
		"roleispresenter": VirtualEventSessionAllowedPresentersRoleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAllowedPresenters(input)
	return &out, nil
}

type VirtualEventSessionAnonymizeIdentityForRoles string

const (
	VirtualEventSessionAnonymizeIdentityForRolesAttendee    VirtualEventSessionAnonymizeIdentityForRoles = "attendee"
	VirtualEventSessionAnonymizeIdentityForRolesCoorganizer VirtualEventSessionAnonymizeIdentityForRoles = "coorganizer"
	VirtualEventSessionAnonymizeIdentityForRolesPresenter   VirtualEventSessionAnonymizeIdentityForRoles = "presenter"
	VirtualEventSessionAnonymizeIdentityForRolesProducer    VirtualEventSessionAnonymizeIdentityForRoles = "producer"
)

func PossibleValuesForVirtualEventSessionAnonymizeIdentityForRoles() []string {
	return []string{
		string(VirtualEventSessionAnonymizeIdentityForRolesAttendee),
		string(VirtualEventSessionAnonymizeIdentityForRolesCoorganizer),
		string(VirtualEventSessionAnonymizeIdentityForRolesPresenter),
		string(VirtualEventSessionAnonymizeIdentityForRolesProducer),
	}
}

func (s *VirtualEventSessionAnonymizeIdentityForRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionAnonymizeIdentityForRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionAnonymizeIdentityForRoles(input string) (*VirtualEventSessionAnonymizeIdentityForRoles, error) {
	vals := map[string]VirtualEventSessionAnonymizeIdentityForRoles{
		"attendee":    VirtualEventSessionAnonymizeIdentityForRolesAttendee,
		"coorganizer": VirtualEventSessionAnonymizeIdentityForRolesCoorganizer,
		"presenter":   VirtualEventSessionAnonymizeIdentityForRolesPresenter,
		"producer":    VirtualEventSessionAnonymizeIdentityForRolesProducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionAnonymizeIdentityForRoles(input)
	return &out, nil
}

type VirtualEventSessionShareMeetingChatHistoryDefault string

const (
	VirtualEventSessionShareMeetingChatHistoryDefaultAll  VirtualEventSessionShareMeetingChatHistoryDefault = "all"
	VirtualEventSessionShareMeetingChatHistoryDefaultNone VirtualEventSessionShareMeetingChatHistoryDefault = "none"
)

func PossibleValuesForVirtualEventSessionShareMeetingChatHistoryDefault() []string {
	return []string{
		string(VirtualEventSessionShareMeetingChatHistoryDefaultAll),
		string(VirtualEventSessionShareMeetingChatHistoryDefaultNone),
	}
}

func (s *VirtualEventSessionShareMeetingChatHistoryDefault) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventSessionShareMeetingChatHistoryDefault(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventSessionShareMeetingChatHistoryDefault(input string) (*VirtualEventSessionShareMeetingChatHistoryDefault, error) {
	vals := map[string]VirtualEventSessionShareMeetingChatHistoryDefault{
		"all":  VirtualEventSessionShareMeetingChatHistoryDefaultAll,
		"none": VirtualEventSessionShareMeetingChatHistoryDefaultNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventSessionShareMeetingChatHistoryDefault(input)
	return &out, nil
}

type VirtualEventWebinarAudience string

const (
	VirtualEventWebinarAudienceEveryone     VirtualEventWebinarAudience = "everyone"
	VirtualEventWebinarAudienceOrganization VirtualEventWebinarAudience = "organization"
)

func PossibleValuesForVirtualEventWebinarAudience() []string {
	return []string{
		string(VirtualEventWebinarAudienceEveryone),
		string(VirtualEventWebinarAudienceOrganization),
	}
}

func (s *VirtualEventWebinarAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventWebinarAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventWebinarAudience(input string) (*VirtualEventWebinarAudience, error) {
	vals := map[string]VirtualEventWebinarAudience{
		"everyone":     VirtualEventWebinarAudienceEveryone,
		"organization": VirtualEventWebinarAudienceOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventWebinarAudience(input)
	return &out, nil
}

type VirtualEventWebinarStatus string

const (
	VirtualEventWebinarStatusCanceled  VirtualEventWebinarStatus = "canceled"
	VirtualEventWebinarStatusDraft     VirtualEventWebinarStatus = "draft"
	VirtualEventWebinarStatusPublished VirtualEventWebinarStatus = "published"
)

func PossibleValuesForVirtualEventWebinarStatus() []string {
	return []string{
		string(VirtualEventWebinarStatusCanceled),
		string(VirtualEventWebinarStatusDraft),
		string(VirtualEventWebinarStatusPublished),
	}
}

func (s *VirtualEventWebinarStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventWebinarStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventWebinarStatus(input string) (*VirtualEventWebinarStatus, error) {
	vals := map[string]VirtualEventWebinarStatus{
		"canceled":  VirtualEventWebinarStatusCanceled,
		"draft":     VirtualEventWebinarStatusDraft,
		"published": VirtualEventWebinarStatusPublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventWebinarStatus(input)
	return &out, nil
}

type WebAccountAllowedAudiences string

const (
	WebAccountAllowedAudiencesContacts               WebAccountAllowedAudiences = "contacts"
	WebAccountAllowedAudiencesEveryone               WebAccountAllowedAudiences = "everyone"
	WebAccountAllowedAudiencesFamily                 WebAccountAllowedAudiences = "family"
	WebAccountAllowedAudiencesFederatedOrganizations WebAccountAllowedAudiences = "federatedOrganizations"
	WebAccountAllowedAudiencesGroupMembers           WebAccountAllowedAudiences = "groupMembers"
	WebAccountAllowedAudiencesMe                     WebAccountAllowedAudiences = "me"
	WebAccountAllowedAudiencesOrganization           WebAccountAllowedAudiences = "organization"
)

func PossibleValuesForWebAccountAllowedAudiences() []string {
	return []string{
		string(WebAccountAllowedAudiencesContacts),
		string(WebAccountAllowedAudiencesEveryone),
		string(WebAccountAllowedAudiencesFamily),
		string(WebAccountAllowedAudiencesFederatedOrganizations),
		string(WebAccountAllowedAudiencesGroupMembers),
		string(WebAccountAllowedAudiencesMe),
		string(WebAccountAllowedAudiencesOrganization),
	}
}

func (s *WebAccountAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebAccountAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebAccountAllowedAudiences(input string) (*WebAccountAllowedAudiences, error) {
	vals := map[string]WebAccountAllowedAudiences{
		"contacts":               WebAccountAllowedAudiencesContacts,
		"everyone":               WebAccountAllowedAudiencesEveryone,
		"family":                 WebAccountAllowedAudiencesFamily,
		"federatedorganizations": WebAccountAllowedAudiencesFederatedOrganizations,
		"groupmembers":           WebAccountAllowedAudiencesGroupMembers,
		"me":                     WebAccountAllowedAudiencesMe,
		"organization":           WebAccountAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebAccountAllowedAudiences(input)
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

type WorkPositionAllowedAudiences string

const (
	WorkPositionAllowedAudiencesContacts               WorkPositionAllowedAudiences = "contacts"
	WorkPositionAllowedAudiencesEveryone               WorkPositionAllowedAudiences = "everyone"
	WorkPositionAllowedAudiencesFamily                 WorkPositionAllowedAudiences = "family"
	WorkPositionAllowedAudiencesFederatedOrganizations WorkPositionAllowedAudiences = "federatedOrganizations"
	WorkPositionAllowedAudiencesGroupMembers           WorkPositionAllowedAudiences = "groupMembers"
	WorkPositionAllowedAudiencesMe                     WorkPositionAllowedAudiences = "me"
	WorkPositionAllowedAudiencesOrganization           WorkPositionAllowedAudiences = "organization"
)

func PossibleValuesForWorkPositionAllowedAudiences() []string {
	return []string{
		string(WorkPositionAllowedAudiencesContacts),
		string(WorkPositionAllowedAudiencesEveryone),
		string(WorkPositionAllowedAudiencesFamily),
		string(WorkPositionAllowedAudiencesFederatedOrganizations),
		string(WorkPositionAllowedAudiencesGroupMembers),
		string(WorkPositionAllowedAudiencesMe),
		string(WorkPositionAllowedAudiencesOrganization),
	}
}

func (s *WorkPositionAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkPositionAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkPositionAllowedAudiences(input string) (*WorkPositionAllowedAudiences, error) {
	vals := map[string]WorkPositionAllowedAudiences{
		"contacts":               WorkPositionAllowedAudiencesContacts,
		"everyone":               WorkPositionAllowedAudiencesEveryone,
		"family":                 WorkPositionAllowedAudiencesFamily,
		"federatedorganizations": WorkPositionAllowedAudiencesFederatedOrganizations,
		"groupmembers":           WorkPositionAllowedAudiencesGroupMembers,
		"me":                     WorkPositionAllowedAudiencesMe,
		"organization":           WorkPositionAllowedAudiencesOrganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkPositionAllowedAudiences(input)
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
