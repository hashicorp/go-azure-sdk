package applicationowner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationOwnerClient struct {
	Client *msgraph.Client
}

func NewApplicationOwnerClientWithBaseURI(api sdkEnv.Api) (*ApplicationOwnerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationOwnerClient: %+v", err)
	}

	return &ApplicationOwnerClient{
		Client: client,
	}, nil
}
