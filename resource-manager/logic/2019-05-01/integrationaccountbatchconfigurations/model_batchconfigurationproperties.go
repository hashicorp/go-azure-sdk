package integrationaccountbatchconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchConfigurationProperties struct {
	BatchGroupName  string               `json:"batchGroupName"`
	ChangedTime     *string              `json:"changedTime,omitempty"`
	CreatedTime     *string              `json:"createdTime,omitempty"`
	Metadata        *interface{}         `json:"metadata,omitempty"`
	ReleaseCriteria BatchReleaseCriteria `json:"releaseCriteria"`
}
