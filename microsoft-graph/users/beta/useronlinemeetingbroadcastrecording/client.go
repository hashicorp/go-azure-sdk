package useronlinemeetingbroadcastrecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingBroadcastRecordingClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingBroadcastRecordingClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingBroadcastRecordingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingbroadcastrecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingBroadcastRecordingClient: %+v", err)
	}

	return &UserOnlineMeetingBroadcastRecordingClient{
		Client: client,
	}, nil
}
