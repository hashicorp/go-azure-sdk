package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateConnectorHealthMetricValue struct {
	DateTime     *string `json:"dateTime,omitempty"`
	FailureCount *int64  `json:"failureCount,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	SuccessCount *int64  `json:"successCount,omitempty"`
}
