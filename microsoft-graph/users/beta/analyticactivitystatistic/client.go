package analyticactivitystatistic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyticActivityStatisticClient struct {
	Client *msgraph.Client
}

func NewAnalyticActivityStatisticClientWithBaseURI(sdkApi sdkEnv.Api) (*AnalyticActivityStatisticClient, error) {
	client, err := msgraph.NewClient(sdkApi, "analyticactivitystatistic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AnalyticActivityStatisticClient: %+v", err)
	}

	return &AnalyticActivityStatisticClient{
		Client: client,
	}, nil
}
