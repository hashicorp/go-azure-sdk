package datalakestoreaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataLakeStoreAccountInformation struct {
	Id         *string                                    `json:"id,omitempty"`
	Name       *string                                    `json:"name,omitempty"`
	Properties *DataLakeStoreAccountInformationProperties `json:"properties"`
	Type       *string                                    `json:"type,omitempty"`
}
