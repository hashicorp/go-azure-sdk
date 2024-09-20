package driveitemlistitemversion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDriveItemListItemVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.ListItemVersion
}

type CreateDriveItemListItemVersionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDriveItemListItemVersionOperationOptions() CreateDriveItemListItemVersionOperationOptions {
	return CreateDriveItemListItemVersionOperationOptions{}
}

func (o CreateDriveItemListItemVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveItemListItemVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveItemListItemVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveItemListItemVersion - Create new navigation property to versions for groups
func (c DriveItemListItemVersionClient) CreateDriveItemListItemVersion(ctx context.Context, id stable.GroupIdDriveIdItemId, input stable.ListItemVersion, options CreateDriveItemListItemVersionOperationOptions) (result CreateDriveItemListItemVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/listItem/versions", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalListItemVersionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
