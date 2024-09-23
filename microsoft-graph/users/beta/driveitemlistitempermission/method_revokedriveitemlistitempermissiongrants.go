package driveitemlistitempermission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevokeDriveItemListItemPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Permission
}

type RevokeDriveItemListItemPermissionGrantsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRevokeDriveItemListItemPermissionGrantsOperationOptions() RevokeDriveItemListItemPermissionGrantsOperationOptions {
	return RevokeDriveItemListItemPermissionGrantsOperationOptions{}
}

func (o RevokeDriveItemListItemPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RevokeDriveItemListItemPermissionGrantsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RevokeDriveItemListItemPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RevokeDriveItemListItemPermissionGrants - Invoke action revokeGrants. Revoke access to a listItem or driveItem
// granted via a sharing link by removing the specified recipient from the link.
func (c DriveItemListItemPermissionClient) RevokeDriveItemListItemPermissionGrants(ctx context.Context, id beta.UserIdDriveIdItemIdListItemPermissionId, input RevokeDriveItemListItemPermissionGrantsRequest, options RevokeDriveItemListItemPermissionGrantsOperationOptions) (result RevokeDriveItemListItemPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/revokeGrants", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.Permission
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
