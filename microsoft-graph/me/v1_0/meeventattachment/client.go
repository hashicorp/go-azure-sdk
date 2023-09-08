package meeventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventAttachmentClient: %+v", err)
	}

	return &MeEventAttachmentClient{
		Client: client,
	}, nil
}
