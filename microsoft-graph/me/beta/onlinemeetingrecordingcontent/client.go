package onlinemeetingrecordingcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRecordingContentClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingRecordingContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingRecordingContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingrecordingcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingRecordingContentClient: %+v", err)
	}

	return &OnlineMeetingRecordingContentClient{
		Client: client,
	}, nil
}
