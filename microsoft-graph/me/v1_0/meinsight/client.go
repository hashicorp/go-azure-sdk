package meinsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInsightClient struct {
	Client *msgraph.Client
}

func NewMeInsightClientWithBaseURI(api sdkEnv.Api) (*MeInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInsightClient: %+v", err)
	}

	return &MeInsightClient{
		Client: client,
	}, nil
}
