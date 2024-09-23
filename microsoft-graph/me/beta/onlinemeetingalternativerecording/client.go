package onlinemeetingalternativerecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAlternativeRecordingClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingAlternativeRecordingClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingAlternativeRecordingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingalternativerecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingAlternativeRecordingClient: %+v", err)
	}

	return &OnlineMeetingAlternativeRecordingClient{
		Client: client,
	}, nil
}
