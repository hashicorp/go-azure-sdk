package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionDownloadDetails struct {
	AppLogDecryptionAlgorithm *AppLogCollectionDownloadDetailsAppLogDecryptionAlgorithm `json:"appLogDecryptionAlgorithm,omitempty"`
	DecryptionKey             *string                                                   `json:"decryptionKey,omitempty"`
	DownloadUrl               *string                                                   `json:"downloadUrl,omitempty"`
	ODataType                 *string                                                   `json:"@odata.type,omitempty"`
}
