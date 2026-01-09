package onlinemeetingtranscript

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingTranscriptClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingTranscriptClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingTranscriptClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingtranscript", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingTranscriptClient: %+v", err)
	}

	return &OnlineMeetingTranscriptClient{
		Client: client,
	}, nil
}
