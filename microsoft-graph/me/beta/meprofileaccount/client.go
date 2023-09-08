package meprofileaccount

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileAccountClient struct {
	Client *msgraph.Client
}

func NewMeProfileAccountClientWithBaseURI(api sdkEnv.Api) (*MeProfileAccountClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileaccount", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileAccountClient: %+v", err)
	}

	return &MeProfileAccountClient{
		Client: client,
	}, nil
}
