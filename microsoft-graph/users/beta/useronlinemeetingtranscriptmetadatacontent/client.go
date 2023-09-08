package useronlinemeetingtranscriptmetadatacontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingTranscriptMetadataContentClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingTranscriptMetadataContentClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingTranscriptMetadataContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingtranscriptmetadatacontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingTranscriptMetadataContentClient: %+v", err)
	}

	return &UserOnlineMeetingTranscriptMetadataContentClient{
		Client: client,
	}, nil
}
