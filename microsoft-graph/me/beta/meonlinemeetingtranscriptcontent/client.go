package meonlinemeetingtranscriptcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingTranscriptContentClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingTranscriptContentClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingTranscriptContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingtranscriptcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingTranscriptContentClient: %+v", err)
	}

	return &MeOnlineMeetingTranscriptContentClient{
		Client: client,
	}, nil
}
