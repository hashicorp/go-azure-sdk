package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActivityTimeSet struct {
	LastRecordedDateTime *string `json:"lastRecordedDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	ObservedDateTime     *string `json:"observedDateTime,omitempty"`
	RecordedDateTime     *string `json:"recordedDateTime,omitempty"`
}
