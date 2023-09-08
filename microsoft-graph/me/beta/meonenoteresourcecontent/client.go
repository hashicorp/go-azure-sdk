package meonenoteresourcecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteResourceContentClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteResourceContentClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteResourceContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenoteresourcecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteResourceContentClient: %+v", err)
	}

	return &MeOnenoteResourceContentClient{
		Client: client,
	}, nil
}
