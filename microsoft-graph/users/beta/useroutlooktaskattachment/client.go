package useroutlooktaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskAttachmentClient: %+v", err)
	}

	return &UserOutlookTaskAttachmentClient{
		Client: client,
	}, nil
}
