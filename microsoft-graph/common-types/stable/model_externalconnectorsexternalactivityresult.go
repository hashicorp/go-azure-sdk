package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalActivityResult struct {
	Error         *PublicError                                  `json:"error,omitempty"`
	Id            *string                                       `json:"id,omitempty"`
	ODataType     *string                                       `json:"@odata.type,omitempty"`
	PerformedBy   *ExternalConnectorsIdentity                   `json:"performedBy,omitempty"`
	StartDateTime *string                                       `json:"startDateTime,omitempty"`
	Type          *ExternalConnectorsExternalActivityResultType `json:"type,omitempty"`
}
