package profile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileClient struct {
	Client *msgraph.Client
}

func NewProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileClient: %+v", err)
	}

	return &ProfileClient{
		Client: client,
	}, nil
}
