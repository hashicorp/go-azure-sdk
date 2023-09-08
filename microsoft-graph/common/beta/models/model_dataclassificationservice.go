package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataClassificationService struct {
	ClassifyFileJobs        *[]JobResponseBase       `json:"classifyFileJobs,omitempty"`
	ClassifyTextJobs        *[]JobResponseBase       `json:"classifyTextJobs,omitempty"`
	EvaluateDlpPoliciesJobs *[]JobResponseBase       `json:"evaluateDlpPoliciesJobs,omitempty"`
	EvaluateLabelJobs       *[]JobResponseBase       `json:"evaluateLabelJobs,omitempty"`
	ExactMatchDataStores    *[]ExactMatchDataStore   `json:"exactMatchDataStores,omitempty"`
	ExactMatchUploadAgents  *[]ExactMatchUploadAgent `json:"exactMatchUploadAgents,omitempty"`
	Id                      *string                  `json:"id,omitempty"`
	Jobs                    *[]JobResponseBase       `json:"jobs,omitempty"`
	ODataType               *string                  `json:"@odata.type,omitempty"`
	SensitiveTypes          *[]SensitiveType         `json:"sensitiveTypes,omitempty"`
	SensitivityLabels       *[]SensitivityLabel      `json:"sensitivityLabels,omitempty"`
}
