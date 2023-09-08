package meeventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventInstanceAttachmentClient: %+v", err)
	}

	return &MeEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
