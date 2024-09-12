package drivelistcreatedbyuser

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

type GetDriveListCreatedByUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.User
}

type GetDriveListCreatedByUserOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDriveListCreatedByUserOperationOptions() GetDriveListCreatedByUserOperationOptions {
	return GetDriveListCreatedByUserOperationOptions{}
}

func (o GetDriveListCreatedByUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveListCreatedByUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDriveListCreatedByUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveListCreatedByUser - Get createdByUser from me
func (c DriveListCreatedByUserClient) GetDriveListCreatedByUser(ctx context.Context, id beta.MeDriveId, options GetDriveListCreatedByUserOperationOptions) (result GetDriveListCreatedByUserOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/list/createdByUser", id.ID()),
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

	var model beta.User
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
