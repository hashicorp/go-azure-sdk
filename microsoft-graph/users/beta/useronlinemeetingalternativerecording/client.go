package useronlinemeetingalternativerecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingAlternativeRecordingClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingAlternativeRecordingClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingAlternativeRecordingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingalternativerecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingAlternativeRecordingClient: %+v", err)
	}

	return &UserOnlineMeetingAlternativeRecordingClient{
		Client: client,
	}, nil
}
