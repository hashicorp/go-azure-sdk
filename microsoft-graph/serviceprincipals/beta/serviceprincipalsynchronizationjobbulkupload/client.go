package serviceprincipalsynchronizationjobbulkupload

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSynchronizationJobBulkUploadClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSynchronizationJobBulkUploadClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalSynchronizationJobBulkUploadClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalsynchronizationjobbulkupload", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSynchronizationJobBulkUploadClient: %+v", err)
	}

	return &ServicePrincipalSynchronizationJobBulkUploadClient{
		Client: client,
	}, nil
}
