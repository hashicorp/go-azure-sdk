package driverootlastmodifiedbyuser

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDriveRootLastModifiedByUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.User
}

type GetDriveRootLastModifiedByUserOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDriveRootLastModifiedByUserOperationOptions() GetDriveRootLastModifiedByUserOperationOptions {
	return GetDriveRootLastModifiedByUserOperationOptions{}
}

func (o GetDriveRootLastModifiedByUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveRootLastModifiedByUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDriveRootLastModifiedByUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveRootLastModifiedByUser - Get lastModifiedByUser from users. Identity of the user who last modified the item.
// Read-only.
func (c DriveRootLastModifiedByUserClient) GetDriveRootLastModifiedByUser(ctx context.Context, id stable.UserIdDriveId, options GetDriveRootLastModifiedByUserOperationOptions) (result GetDriveRootLastModifiedByUserOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/root/lastModifiedByUser", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model stable.User
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
