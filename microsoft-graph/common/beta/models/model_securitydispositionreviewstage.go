package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDispositionReviewStage struct {
	Id                      *string   `json:"id,omitempty"`
	Name                    *string   `json:"name,omitempty"`
	ODataType               *string   `json:"@odata.type,omitempty"`
	ReviewersEmailAddresses *[]string `json:"reviewersEmailAddresses,omitempty"`
	StageNumber             *int64    `json:"stageNumber,omitempty"`
}
