package meanalyticactivitystatistic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAnalyticActivityStatisticClient struct {
	Client *msgraph.Client
}

func NewMeAnalyticActivityStatisticClientWithBaseURI(api sdkEnv.Api) (*MeAnalyticActivityStatisticClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meanalyticactivitystatistic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAnalyticActivityStatisticClient: %+v", err)
	}

	return &MeAnalyticActivityStatisticClient{
		Client: client,
	}, nil
}
