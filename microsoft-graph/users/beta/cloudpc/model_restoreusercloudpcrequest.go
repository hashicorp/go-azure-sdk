package cloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreUserCloudPCRequest struct {
	CloudPCSnapshotId *string `json:"cloudPcSnapshotId,omitempty"`
}
