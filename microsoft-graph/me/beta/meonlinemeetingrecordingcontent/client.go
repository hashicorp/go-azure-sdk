package meonlinemeetingrecordingcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingRecordingContentClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingRecordingContentClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingRecordingContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingrecordingcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingRecordingContentClient: %+v", err)
	}

	return &MeOnlineMeetingRecordingContentClient{
		Client: client,
	}, nil
}
