package synchronizationjobbulkupload

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobBulkUploadClient struct {
	Client *msgraph.Client
}

func NewSynchronizationJobBulkUploadClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationJobBulkUploadClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationjobbulkupload", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationJobBulkUploadClient: %+v", err)
	}

	return &SynchronizationJobBulkUploadClient{
		Client: client,
	}, nil
}
