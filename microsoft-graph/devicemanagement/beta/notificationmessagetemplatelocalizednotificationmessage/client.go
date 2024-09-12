package notificationmessagetemplatelocalizednotificationmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationMessageTemplateLocalizedNotificationMessageClient struct {
	Client *msgraph.Client
}

func NewNotificationMessageTemplateLocalizedNotificationMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*NotificationMessageTemplateLocalizedNotificationMessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "notificationmessagetemplatelocalizednotificationmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NotificationMessageTemplateLocalizedNotificationMessageClient: %+v", err)
	}

	return &NotificationMessageTemplateLocalizedNotificationMessageClient{
		Client: client,
	}, nil
}
