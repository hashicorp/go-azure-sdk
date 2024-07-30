package people

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeopleClient struct {
	Client *msgraph.Client
}

func NewPeopleClientWithBaseURI(sdkApi sdkEnv.Api) (*PeopleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "people", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PeopleClient: %+v", err)
	}

	return &PeopleClient{
		Client: client,
	}, nil
}
