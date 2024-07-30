package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationInfo struct {
	CertificateUserIds *[]string `json:"certificateUserIds,omitempty"`
	ODataType          *string   `json:"@odata.type,omitempty"`
}
