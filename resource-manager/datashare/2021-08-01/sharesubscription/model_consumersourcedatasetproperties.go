package sharesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerSourceDataSetProperties struct {
	DataSetId       *string      `json:"dataSetId,omitempty"`
	DataSetLocation *string      `json:"dataSetLocation,omitempty"`
	DataSetName     *string      `json:"dataSetName,omitempty"`
	DataSetPath     *string      `json:"dataSetPath,omitempty"`
	DataSetType     *DataSetType `json:"dataSetType,omitempty"`
}
