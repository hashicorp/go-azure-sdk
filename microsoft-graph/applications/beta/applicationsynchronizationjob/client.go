package applicationsynchronizationjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSynchronizationJobClient struct {
	Client *msgraph.Client
}

func NewApplicationSynchronizationJobClientWithBaseURI(api sdkEnv.Api) (*ApplicationSynchronizationJobClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationsynchronizationjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationSynchronizationJobClient: %+v", err)
	}

	return &ApplicationSynchronizationJobClient{
		Client: client,
	}, nil
}
