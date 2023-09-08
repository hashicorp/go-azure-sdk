package meonlinemeetingrecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingRecordingClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingRecordingClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingRecordingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingrecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingRecordingClient: %+v", err)
	}

	return &MeOnlineMeetingRecordingClient{
		Client: client,
	}, nil
}
