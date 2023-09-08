package meanalytic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAnalyticClient struct {
	Client *msgraph.Client
}

func NewMeAnalyticClientWithBaseURI(api sdkEnv.Api) (*MeAnalyticClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meanalytic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAnalyticClient: %+v", err)
	}

	return &MeAnalyticClient{
		Client: client,
	}, nil
}
