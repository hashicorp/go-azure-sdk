package groupthreadpostattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpostattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostAttachmentClient: %+v", err)
	}

	return &GroupThreadPostAttachmentClient{
		Client: client,
	}, nil
}
