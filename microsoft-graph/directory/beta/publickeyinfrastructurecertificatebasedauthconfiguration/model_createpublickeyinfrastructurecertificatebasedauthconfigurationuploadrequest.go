package publickeyinfrastructurecertificatebasedauthconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePublicKeyInfrastructureCertificateBasedAuthConfigurationUploadRequest struct {
	Sha256FileHash *string `json:"sha256FileHash,omitempty"`
	UploadUrl      *string `json:"uploadUrl,omitempty"`
}
