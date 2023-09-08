package meonlinemeetingalternativerecording

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingAlternativeRecordingClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingAlternativeRecordingClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingAlternativeRecordingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingalternativerecording", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingAlternativeRecordingClient: %+v", err)
	}

	return &MeOnlineMeetingAlternativeRecordingClient{
		Client: client,
	}, nil
}
