package meonlinemeetingbroadcastrecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingBroadcastRecordingClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingBroadcastRecordingClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingBroadcastRecordingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingbroadcastrecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingBroadcastRecordingClient: %+v", err)
	}

	return &MeOnlineMeetingBroadcastRecordingClient{
		Client: client,
	}, nil
}
