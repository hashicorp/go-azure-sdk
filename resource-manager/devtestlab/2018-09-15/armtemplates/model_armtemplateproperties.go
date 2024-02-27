package armtemplates

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmTemplateProperties struct {
	Contents                 *interface{}               `json:"contents,omitempty"`
	CreatedDate              *string                    `json:"createdDate,omitempty"`
	Description              *string                    `json:"description,omitempty"`
	DisplayName              *string                    `json:"displayName,omitempty"`
	Enabled                  *bool                      `json:"enabled,omitempty"`
	Icon                     *string                    `json:"icon,omitempty"`
	ParametersValueFilesInfo *[]ParametersValueFileInfo `json:"parametersValueFilesInfo,omitempty"`
	Publisher                *string                    `json:"publisher,omitempty"`
}

func (o *ArmTemplateProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}
