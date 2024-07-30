package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipOutlierInsight struct {
	Container              *DirectoryObject                              `json:"container,omitempty"`
	ContainerId            *string                                       `json:"containerId,omitempty"`
	Id                     *string                                       `json:"id,omitempty"`
	InsightCreatedDateTime *string                                       `json:"insightCreatedDateTime,omitempty"`
	LastModifiedBy         *User                                         `json:"lastModifiedBy,omitempty"`
	Member                 *DirectoryObject                              `json:"member,omitempty"`
	MemberId               *string                                       `json:"memberId,omitempty"`
	ODataType              *string                                       `json:"@odata.type,omitempty"`
	OutlierContainerType   *MembershipOutlierInsightOutlierContainerType `json:"outlierContainerType,omitempty"`
	OutlierMemberType      *MembershipOutlierInsightOutlierMemberType    `json:"outlierMemberType,omitempty"`
}
