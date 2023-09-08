package meprofilename

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileNameClient struct {
	Client *msgraph.Client
}

func NewMeProfileNameClientWithBaseURI(api sdkEnv.Api) (*MeProfileNameClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofilename", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileNameClient: %+v", err)
	}

	return &MeProfileNameClient{
		Client: client,
	}, nil
}
