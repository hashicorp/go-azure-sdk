package memanager

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagerClient struct {
	Client *msgraph.Client
}

func NewMeManagerClientWithBaseURI(api sdkEnv.Api) (*MeManagerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanager", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagerClient: %+v", err)
	}

	return &MeManagerClient{
		Client: client,
	}, nil
}
