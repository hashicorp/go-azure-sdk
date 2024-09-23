package profilewebaccount

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileWebAccountClient struct {
	Client *msgraph.Client
}

func NewProfileWebAccountClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileWebAccountClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilewebaccount", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileWebAccountClient: %+v", err)
	}

	return &ProfileWebAccountClient{
		Client: client,
	}, nil
}
