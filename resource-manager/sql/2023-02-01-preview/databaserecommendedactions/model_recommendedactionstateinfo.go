package databaserecommendedactions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionStateInfo struct {
	ActionInitiatedBy *RecommendedActionInitiatedBy `json:"actionInitiatedBy,omitempty"`
	CurrentValue      RecommendedActionCurrentState `json:"currentValue"`
	LastModified      *string                       `json:"lastModified,omitempty"`
}

func (o *RecommendedActionStateInfo) GetLastModifiedAsTime() (*time.Time, error) {
	if o.LastModified == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModified, "2006-01-02T15:04:05Z07:00")
}
