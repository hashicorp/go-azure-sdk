package meonlinemeetingtranscriptmetadatacontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingTranscriptMetadataContentClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingTranscriptMetadataContentClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingTranscriptMetadataContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingtranscriptmetadatacontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingTranscriptMetadataContentClient: %+v", err)
	}

	return &MeOnlineMeetingTranscriptMetadataContentClient{
		Client: client,
	}, nil
}
