package useronlinemeetingtranscriptcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingTranscriptContentClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingTranscriptContentClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingTranscriptContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingtranscriptcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingTranscriptContentClient: %+v", err)
	}

	return &UserOnlineMeetingTranscriptContentClient{
		Client: client,
	}, nil
}
