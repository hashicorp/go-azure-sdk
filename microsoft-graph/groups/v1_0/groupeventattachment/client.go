package groupeventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventAttachmentClient: %+v", err)
	}

	return &GroupEventAttachmentClient{
		Client: client,
	}, nil
}
