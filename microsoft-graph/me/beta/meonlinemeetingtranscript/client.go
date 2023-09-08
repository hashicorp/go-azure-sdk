package meonlinemeetingtranscript

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingTranscriptClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingTranscriptClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingTranscriptClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingtranscript", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingTranscriptClient: %+v", err)
	}

	return &MeOnlineMeetingTranscriptClient{
		Client: client,
	}, nil
}
