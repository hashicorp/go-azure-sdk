package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DCRConfiguration struct {
	DataCollectionEndpoint        string `json:"dataCollectionEndpoint"`
	DataCollectionRuleImmutableId string `json:"dataCollectionRuleImmutableId"`
	StreamName                    string `json:"streamName"`
}
