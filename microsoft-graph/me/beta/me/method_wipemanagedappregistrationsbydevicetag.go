package me

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WipeManagedAppRegistrationsByDeviceTagOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// WipeManagedAppRegistrationsByDeviceTag - Invoke action wipeManagedAppRegistrationsByDeviceTag. Issues a wipe
// operation on an app registration with specified device tag.
func (c MeClient) WipeManagedAppRegistrationsByDeviceTag(ctx context.Context, input WipeManagedAppRegistrationsByDeviceTagRequest) (result WipeManagedAppRegistrationsByDeviceTagOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       "/me/wipeManagedAppRegistrationsByDeviceTag",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
