package containerappspatches

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatchDetails struct {
	DetectionStatus     DetectionStatus       `json:"detectionStatus"`
	LastDetectionTime   string                `json:"lastDetectionTime"`
	NewImageName        *string               `json:"newImageName,omitempty"`
	NewLayer            *PatchDetailsNewLayer `json:"newLayer,omitempty"`
	OldLayer            *PatchDetailsOldLayer `json:"oldLayer,omitempty"`
	PatchType           *PatchType            `json:"patchType,omitempty"`
	TargetContainerName string                `json:"targetContainerName"`
	TargetImage         string                `json:"targetImage"`
}

func (o *PatchDetails) GetLastDetectionTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastDetectionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PatchDetails) SetLastDetectionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastDetectionTime = formatted
}
