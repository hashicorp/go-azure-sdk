package onlinemeetingtranscriptmetadatacontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingTranscriptMetadataContentClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingTranscriptMetadataContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingTranscriptMetadataContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingtranscriptmetadatacontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingTranscriptMetadataContentClient: %+v", err)
	}

	return &OnlineMeetingTranscriptMetadataContentClient{
		Client: client,
	}, nil
}
