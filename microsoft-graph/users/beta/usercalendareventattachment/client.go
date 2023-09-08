package usercalendareventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventAttachmentClient: %+v", err)
	}

	return &UserCalendarEventAttachmentClient{
		Client: client,
	}, nil
}
