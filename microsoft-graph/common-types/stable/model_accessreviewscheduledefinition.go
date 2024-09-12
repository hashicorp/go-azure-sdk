package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewScheduleDefinition{}

type AccessReviewScheduleDefinition struct {
	// Defines the list of additional users or group members to be notified of the access review progress.
	AdditionalNotificationRecipients *[]AccessReviewNotificationRecipientItem `json:"additionalNotificationRecipients,omitempty"`

	// User who created this review. Read-only.
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`

	// Timestamp when the access review series was created. Supports $select. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description provided by review creators to provide more context of the review to admins. Supports $select.
	DescriptionForAdmins nullable.Type[string] `json:"descriptionForAdmins,omitempty"`

	// Description provided by review creators to provide more context of the review to reviewers. Reviewers see this
	// description in the email sent to them requesting their review. Email notifications support up to 256 characters.
	// Supports $select.
	DescriptionForReviewers nullable.Type[string] `json:"descriptionForReviewers,omitempty"`

	// Name of the access review series. Supports $select and $orderby. Required on create.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers are
	// notified to take action if no users are found from the list of reviewers specified. This could occur when either the
	// group owner is specified as the reviewer but the group owner doesn't exist, or manager is specified as reviewer but a
	// user's manager doesn't exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select. NOTE: The
	// value of this property will be ignored if fallback reviewers are assigned through the stageSettings property.
	FallbackReviewers *[]AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`

	// This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines
	// which Microsoft 365 groups are reviewed. Each group becomes a unique accessReviewInstance of the access review
	// series. For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring
	// instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
	InstanceEnumerationScope AccessReviewScope `json:"instanceEnumerationScope"`

	// If the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review
	// that doesn't recur will have exactly one instance. Instances also represent each unique resource under review in the
	// accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource has a unique
	// instance for each recurrence.
	Instances *[]AccessReviewInstance `json:"instances,omitempty"`

	// Timestamp when the access review series was last modified. Supports $select. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// This collection of access review scopes is used to define who are the reviewers. The reviewers property is only
	// updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of
	// options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
	// NOTE: The value of this property will be ignored if reviewers are assigned through the stageSettings property.
	Reviewers *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`

	// Defines the entities whose access is reviewed. For supported scopes, see accessReviewScope. Required on create.
	// Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope
	// of your access review definition using the Microsoft Graph API.
	Scope AccessReviewScope `json:"scope"`

	// The settings for an access review series, see type definition below. Supports $select. Required on create.
	Settings *AccessReviewScheduleSettings `json:"settings,omitempty"`

	// Required only for a multi-stage access review to define the stages and their settings. You can break down each review
	// instance into up to three sequential stages, where each stage can have a different set of reviewers, fallback
	// reviewers, and settings. Stages are created sequentially based on the dependsOn property. Optional. When this
	// property is defined, its settings are used instead of the corresponding settings in the
	// accessReviewScheduleDefinition object and its settings, reviewers, and fallbackReviewers properties.
	StageSettings *[]AccessReviewStageSettings `json:"stageSettings,omitempty"`

	// This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted,
	// Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter
	// (eq only). Read-only.
	Status nullable.Type[string] `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AccessReviewScheduleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewScheduleDefinition{}

func (s AccessReviewScheduleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewScheduleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewScheduleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewScheduleDefinition: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "status")
	decoded["@odata.type"] = "#microsoft.graph.accessReviewScheduleDefinition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewScheduleDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReviewScheduleDefinition{}

func (s *AccessReviewScheduleDefinition) UnmarshalJSON(bytes []byte) error {
	type alias AccessReviewScheduleDefinition
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AccessReviewScheduleDefinition: %+v", err)
	}

	s.AdditionalNotificationRecipients = decoded.AdditionalNotificationRecipients
	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DescriptionForAdmins = decoded.DescriptionForAdmins
	s.DescriptionForReviewers = decoded.DescriptionForReviewers
	s.DisplayName = decoded.DisplayName
	s.FallbackReviewers = decoded.FallbackReviewers
	s.Id = decoded.Id
	s.Instances = decoded.Instances
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Reviewers = decoded.Reviewers
	s.Settings = decoded.Settings
	s.StageSettings = decoded.StageSettings
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewScheduleDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["instanceEnumerationScope"]; ok {
		impl, err := UnmarshalAccessReviewScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InstanceEnumerationScope' for 'AccessReviewScheduleDefinition': %+v", err)
		}
		s.InstanceEnumerationScope = impl
	}

	if v, ok := temp["scope"]; ok {
		impl, err := UnmarshalAccessReviewScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Scope' for 'AccessReviewScheduleDefinition': %+v", err)
		}
		s.Scope = impl
	}
	return nil
}
