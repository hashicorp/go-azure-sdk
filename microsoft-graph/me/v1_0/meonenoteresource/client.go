package meonenoteresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteResourceClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteResourceClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenoteresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteResourceClient: %+v", err)
	}

	return &MeOnenoteResourceClient{
		Client: client,
	}, nil
}
