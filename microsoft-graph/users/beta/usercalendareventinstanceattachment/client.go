package usercalendareventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventInstanceAttachmentClient: %+v", err)
	}

	return &UserCalendarEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
