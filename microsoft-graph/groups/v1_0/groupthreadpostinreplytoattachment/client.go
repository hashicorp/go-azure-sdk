package groupthreadpostinreplytoattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostInReplyToAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostInReplyToAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostInReplyToAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpostinreplytoattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostInReplyToAttachmentClient: %+v", err)
	}

	return &GroupThreadPostInReplyToAttachmentClient{
		Client: client,
	}, nil
}
