package usereventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventInstanceAttachmentClient: %+v", err)
	}

	return &UserEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
