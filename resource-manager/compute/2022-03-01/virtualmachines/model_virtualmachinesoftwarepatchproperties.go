package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineSoftwarePatchProperties struct {
	ActivityId           *string                     `json:"activityId,omitempty"`
	AssessmentState      *PatchAssessmentState       `json:"assessmentState,omitempty"`
	Classifications      *[]string                   `json:"classifications,omitempty"`
	KbId                 *string                     `json:"kbId,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	PatchId              *string                     `json:"patchId,omitempty"`
	PublishedDate        *string                     `json:"publishedDate,omitempty"`
	RebootBehavior       *VMGuestPatchRebootBehavior `json:"rebootBehavior,omitempty"`
	Version              *string                     `json:"version,omitempty"`
}
