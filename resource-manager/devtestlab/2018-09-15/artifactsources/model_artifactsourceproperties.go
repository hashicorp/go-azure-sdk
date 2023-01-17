package artifactsources

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactSourceProperties struct {
	ArmTemplateFolderPath *string            `json:"armTemplateFolderPath,omitempty"`
	BranchRef             *string            `json:"branchRef,omitempty"`
	CreatedDate           *string            `json:"createdDate,omitempty"`
	DisplayName           *string            `json:"displayName,omitempty"`
	FolderPath            *string            `json:"folderPath,omitempty"`
	ProvisioningState     *string            `json:"provisioningState,omitempty"`
	SecurityToken         *string            `json:"securityToken,omitempty"`
	SourceType            *SourceControlType `json:"sourceType,omitempty"`
	Status                *EnableStatus      `json:"status,omitempty"`
	UniqueIdentifier      *string            `json:"uniqueIdentifier,omitempty"`
	Uri                   *string            `json:"uri,omitempty"`
}

func (o *ArtifactSourceProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ArtifactSourceProperties) SetCreatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDate = &formatted
}
