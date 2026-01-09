package onlinemeetingtranscriptcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingTranscriptContentClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingTranscriptContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingTranscriptContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingtranscriptcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingTranscriptContentClient: %+v", err)
	}

	return &OnlineMeetingTranscriptContentClient{
		Client: client,
	}, nil
}
