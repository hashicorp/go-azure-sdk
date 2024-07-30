package notificationmessagetemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationMessageTemplateClient struct {
	Client *msgraph.Client
}

func NewNotificationMessageTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*NotificationMessageTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "notificationmessagetemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NotificationMessageTemplateClient: %+v", err)
	}

	return &NotificationMessageTemplateClient{
		Client: client,
	}, nil
}
