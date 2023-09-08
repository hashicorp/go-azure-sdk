package mesponsor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSponsorClient struct {
	Client *msgraph.Client
}

func NewMeSponsorClientWithBaseURI(api sdkEnv.Api) (*MeSponsorClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesponsor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSponsorClient: %+v", err)
	}

	return &MeSponsorClient{
		Client: client,
	}, nil
}
