package profileaccount

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileAccountClient struct {
	Client *msgraph.Client
}

func NewProfileAccountClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileAccountClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileaccount", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileAccountClient: %+v", err)
	}

	return &ProfileAccountClient{
		Client: client,
	}, nil
}
