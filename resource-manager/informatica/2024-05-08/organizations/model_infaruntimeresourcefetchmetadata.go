package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InfaRuntimeResourceFetchMetaData struct {
	CreatedBy                  string                              `json:"createdBy"`
	CreatedTime                string                              `json:"createdTime"`
	Description                *string                             `json:"description,omitempty"`
	Id                         string                              `json:"id"`
	Name                       string                              `json:"name"`
	ServerlessConfigProperties InfaServerlessFetchConfigProperties `json:"serverlessConfigProperties"`
	Status                     string                              `json:"status"`
	StatusLocalized            string                              `json:"statusLocalized"`
	StatusMessage              string                              `json:"statusMessage"`
	Type                       RuntimeType                         `json:"type"`
	UpdatedBy                  string                              `json:"updatedBy"`
	UpdatedTime                string                              `json:"updatedTime"`
}
