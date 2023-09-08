package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanConfigurationLocalization struct {
	Buckets     *[]PlannerPlanConfigurationBucketLocalization `json:"buckets,omitempty"`
	Id          *string                                       `json:"id,omitempty"`
	LanguageTag *string                                       `json:"languageTag,omitempty"`
	ODataType   *string                                       `json:"@odata.type,omitempty"`
	PlanTitle   *string                                       `json:"planTitle,omitempty"`
}
