package mecontact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactClient struct {
	Client *msgraph.Client
}

func NewMeContactClientWithBaseURI(api sdkEnv.Api) (*MeContactClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactClient: %+v", err)
	}

	return &MeContactClient{
		Client: client,
	}, nil
}
