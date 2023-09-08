package mepeople

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePeopleClient struct {
	Client *msgraph.Client
}

func NewMePeopleClientWithBaseURI(api sdkEnv.Api) (*MePeopleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mepeople", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePeopleClient: %+v", err)
	}

	return &MePeopleClient{
		Client: client,
	}, nil
}
