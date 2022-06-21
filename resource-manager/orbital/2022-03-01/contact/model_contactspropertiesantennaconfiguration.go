package contact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactsPropertiesAntennaConfiguration struct {
	DestinationIp *string   `json:"destinationIp,omitempty"`
	SourceIps     *[]string `json:"sourceIps,omitempty"`
}
