package useronlinemeetingrecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingRecordingClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingRecordingClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingRecordingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingrecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingRecordingClient: %+v", err)
	}

	return &UserOnlineMeetingRecordingClient{
		Client: client,
	}, nil
}
