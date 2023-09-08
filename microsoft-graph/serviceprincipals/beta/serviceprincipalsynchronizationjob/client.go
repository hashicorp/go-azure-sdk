package serviceprincipalsynchronizationjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSynchronizationJobClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSynchronizationJobClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalSynchronizationJobClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalsynchronizationjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSynchronizationJobClient: %+v", err)
	}

	return &ServicePrincipalSynchronizationJobClient{
		Client: client,
	}, nil
}
