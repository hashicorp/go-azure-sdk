package applicationsynchronizationjobbulkupload

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSynchronizationJobBulkUploadClient struct {
	Client *msgraph.Client
}

func NewApplicationSynchronizationJobBulkUploadClientWithBaseURI(api sdkEnv.Api) (*ApplicationSynchronizationJobBulkUploadClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationsynchronizationjobbulkupload", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationSynchronizationJobBulkUploadClient: %+v", err)
	}

	return &ApplicationSynchronizationJobBulkUploadClient{
		Client: client,
	}, nil
}
