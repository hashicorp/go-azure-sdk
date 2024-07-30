package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTask struct {
	ActiveChecklistItemCount  *int64                                `json:"activeChecklistItemCount,omitempty"`
	AppliedCategories         *PlannerAppliedCategories             `json:"appliedCategories,omitempty"`
	AssignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`
	AssigneePriority          *string                               `json:"assigneePriority,omitempty"`
	Assignments               *PlannerAssignments                   `json:"assignments,omitempty"`
	BucketId                  *string                               `json:"bucketId,omitempty"`
	BucketTaskBoardFormat     *PlannerBucketTaskBoardTaskFormat     `json:"bucketTaskBoardFormat,omitempty"`
	ChecklistItemCount        *int64                                `json:"checklistItemCount,omitempty"`
	CompletedBy               *IdentitySet                          `json:"completedBy,omitempty"`
	CompletedDateTime         *string                               `json:"completedDateTime,omitempty"`
	ConversationThreadId      *string                               `json:"conversationThreadId,omitempty"`
	CreatedBy                 *IdentitySet                          `json:"createdBy,omitempty"`
	CreatedDateTime           *string                               `json:"createdDateTime,omitempty"`
	Details                   *PlannerTaskDetails                   `json:"details,omitempty"`
	DueDateTime               *string                               `json:"dueDateTime,omitempty"`
	HasDescription            *bool                                 `json:"hasDescription,omitempty"`
	Id                        *string                               `json:"id,omitempty"`
	ODataType                 *string                               `json:"@odata.type,omitempty"`
	OrderHint                 *string                               `json:"orderHint,omitempty"`
	PercentComplete           *int64                                `json:"percentComplete,omitempty"`
	PlanId                    *string                               `json:"planId,omitempty"`
	PreviewType               *PlannerTaskPreviewType               `json:"previewType,omitempty"`
	Priority                  *int64                                `json:"priority,omitempty"`
	ProgressTaskBoardFormat   *PlannerProgressTaskBoardTaskFormat   `json:"progressTaskBoardFormat,omitempty"`
	ReferenceCount            *int64                                `json:"referenceCount,omitempty"`
	StartDateTime             *string                               `json:"startDateTime,omitempty"`
	Title                     *string                               `json:"title,omitempty"`
}
