package applicationsynchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSynchronizationClient struct {
	Client *msgraph.Client
}

func NewApplicationSynchronizationClientWithBaseURI(api sdkEnv.Api) (*ApplicationSynchronizationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationsynchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationSynchronizationClient: %+v", err)
	}

	return &ApplicationSynchronizationClient{
		Client: client,
	}, nil
}
