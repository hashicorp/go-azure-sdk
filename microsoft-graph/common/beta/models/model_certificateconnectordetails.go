package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateConnectorDetails struct {
	ConnectorName       *string `json:"connectorName,omitempty"`
	ConnectorVersion    *string `json:"connectorVersion,omitempty"`
	EnrollmentDateTime  *string `json:"enrollmentDateTime,omitempty"`
	Id                  *string `json:"id,omitempty"`
	LastCheckinDateTime *string `json:"lastCheckinDateTime,omitempty"`
	MachineName         *string `json:"machineName,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
}
