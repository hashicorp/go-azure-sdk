package person

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonClient struct {
	Client *msgraph.Client
}

func NewPersonClientWithBaseURI(sdkApi sdkEnv.Api) (*PersonClient, error) {
	client, err := msgraph.NewClient(sdkApi, "person", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PersonClient: %+v", err)
	}

	return &PersonClient{
		Client: client,
	}, nil
}
