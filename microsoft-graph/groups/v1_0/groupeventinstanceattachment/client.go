package groupeventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventInstanceAttachmentClient: %+v", err)
	}

	return &GroupEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
