package meevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventClient struct {
	Client *msgraph.Client
}

func NewMeEventClientWithBaseURI(api sdkEnv.Api) (*MeEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventClient: %+v", err)
	}

	return &MeEventClient{
		Client: client,
	}, nil
}
