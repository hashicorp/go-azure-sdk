package onlinemeetingaiinsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAiInsightClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingAiInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingAiInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingaiinsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingAiInsightClient: %+v", err)
	}

	return &OnlineMeetingAiInsightClient{
		Client: client,
	}, nil
}
