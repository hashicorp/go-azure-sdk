package applicationsynchronizationsecret

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSynchronizationSecretClient struct {
	Client *msgraph.Client
}

func NewApplicationSynchronizationSecretClientWithBaseURI(api sdkEnv.Api) (*ApplicationSynchronizationSecretClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationsynchronizationsecret", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationSynchronizationSecretClient: %+v", err)
	}

	return &ApplicationSynchronizationSecretClient{
		Client: client,
	}, nil
}
