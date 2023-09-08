package usereventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventAttachmentClient: %+v", err)
	}

	return &UserEventAttachmentClient{
		Client: client,
	}, nil
}
