package manageddatabasesensitivitylabels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCurrentByDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SensitivityLabel
}

type ListCurrentByDatabaseCompleteResult struct {
	Items []SensitivityLabel
}

type ListCurrentByDatabaseOperationOptions struct {
	Filter *string
}

func DefaultListCurrentByDatabaseOperationOptions() ListCurrentByDatabaseOperationOptions {
	return ListCurrentByDatabaseOperationOptions{}
}

func (o ListCurrentByDatabaseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCurrentByDatabaseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListCurrentByDatabaseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListCurrentByDatabase ...
func (c ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabase(ctx context.Context, id ManagedInstanceDatabaseId, options ListCurrentByDatabaseOperationOptions) (result ListCurrentByDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/currentSensitivityLabels", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]SensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCurrentByDatabaseComplete retrieves all the results into a single object
func (c ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabaseComplete(ctx context.Context, id ManagedInstanceDatabaseId, options ListCurrentByDatabaseOperationOptions) (ListCurrentByDatabaseCompleteResult, error) {
	return c.ListCurrentByDatabaseCompleteMatchingPredicate(ctx, id, options, SensitivityLabelOperationPredicate{})
}

// ListCurrentByDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabaseCompleteMatchingPredicate(ctx context.Context, id ManagedInstanceDatabaseId, options ListCurrentByDatabaseOperationOptions, predicate SensitivityLabelOperationPredicate) (result ListCurrentByDatabaseCompleteResult, err error) {
	items := make([]SensitivityLabel, 0)

	resp, err := c.ListCurrentByDatabase(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListCurrentByDatabaseCompleteResult{
		Items: items,
	}
	return
}
