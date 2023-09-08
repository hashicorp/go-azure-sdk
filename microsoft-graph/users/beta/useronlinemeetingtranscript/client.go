package useronlinemeetingtranscript

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingTranscriptClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingTranscriptClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingTranscriptClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingtranscript", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingTranscriptClient: %+v", err)
	}

	return &UserOnlineMeetingTranscriptClient{
		Client: client,
	}, nil
}
