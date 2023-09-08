package meonlinemeeting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeeting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingClient: %+v", err)
	}

	return &MeOnlineMeetingClient{
		Client: client,
	}, nil
}
