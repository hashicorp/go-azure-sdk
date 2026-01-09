package onlinemeetingbroadcastrecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingBroadcastRecordingClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingBroadcastRecordingClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingBroadcastRecordingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingbroadcastrecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingBroadcastRecordingClient: %+v", err)
	}

	return &OnlineMeetingBroadcastRecordingClient{
		Client: client,
	}, nil
}
