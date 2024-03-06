package restorables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableMongodbCollectionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RestorableMongodbCollectionsListResult
}

type RestorableMongodbCollectionsListOperationOptions struct {
	RestorableMongodbDatabaseRid *string
}

func DefaultRestorableMongodbCollectionsListOperationOptions() RestorableMongodbCollectionsListOperationOptions {
	return RestorableMongodbCollectionsListOperationOptions{}
}

func (o RestorableMongodbCollectionsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestorableMongodbCollectionsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o RestorableMongodbCollectionsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RestorableMongodbDatabaseRid != nil {
		out.Append("restorableMongodbDatabaseRid", fmt.Sprintf("%v", *o.RestorableMongodbDatabaseRid))
	}
	return &out
}

// RestorableMongodbCollectionsList ...
func (c RestorablesClient) RestorableMongodbCollectionsList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableMongodbCollectionsListOperationOptions) (result RestorableMongodbCollectionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/restorableMongodbCollections", id.ID()),
		OptionsObject: options,
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

	var model RestorableMongodbCollectionsListResult
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
