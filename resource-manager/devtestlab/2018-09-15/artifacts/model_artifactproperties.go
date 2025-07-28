package artifacts

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactProperties struct {
	CreatedDate  *string                 `json:"createdDate,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	FilePath     *string                 `json:"filePath,omitempty"`
	Icon         *string                 `json:"icon,omitempty"`
	Parameters   *map[string]interface{} `json:"parameters,omitempty"`
	Publisher    *string                 `json:"publisher,omitempty"`
	TargetOsType *string                 `json:"targetOsType,omitempty"`
	Title        *string                 `json:"title,omitempty"`
}

func (o *ArtifactProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ArtifactProperties) SetCreatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDate = &formatted
}
