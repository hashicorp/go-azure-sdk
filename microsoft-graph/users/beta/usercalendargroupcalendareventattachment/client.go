package usercalendargroupcalendareventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventAttachmentClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventAttachmentClient{
		Client: client,
	}, nil
}
