package analytic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyticClient struct {
	Client *msgraph.Client
}

func NewAnalyticClientWithBaseURI(sdkApi sdkEnv.Api) (*AnalyticClient, error) {
	client, err := msgraph.NewClient(sdkApi, "analytic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AnalyticClient: %+v", err)
	}

	return &AnalyticClient{
		Client: client,
	}, nil
}
