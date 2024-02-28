package securescore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ControlsListBySecureScoreOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SecureScoreControlDetails
}

type ControlsListBySecureScoreCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SecureScoreControlDetails
}

type ControlsListBySecureScoreOperationOptions struct {
	Expand *ExpandControlsEnum
}

func DefaultControlsListBySecureScoreOperationOptions() ControlsListBySecureScoreOperationOptions {
	return ControlsListBySecureScoreOperationOptions{}
}

func (o ControlsListBySecureScoreOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ControlsListBySecureScoreOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ControlsListBySecureScoreOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	return &out
}

// ControlsListBySecureScore ...
func (c SecureScoreClient) ControlsListBySecureScore(ctx context.Context, id SecureScoreId, options ControlsListBySecureScoreOperationOptions) (result ControlsListBySecureScoreOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/secureScoreControls", id.ID()),
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
		Values *[]SecureScoreControlDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ControlsListBySecureScoreComplete retrieves all the results into a single object
func (c SecureScoreClient) ControlsListBySecureScoreComplete(ctx context.Context, id SecureScoreId, options ControlsListBySecureScoreOperationOptions) (ControlsListBySecureScoreCompleteResult, error) {
	return c.ControlsListBySecureScoreCompleteMatchingPredicate(ctx, id, options, SecureScoreControlDetailsOperationPredicate{})
}

// ControlsListBySecureScoreCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecureScoreClient) ControlsListBySecureScoreCompleteMatchingPredicate(ctx context.Context, id SecureScoreId, options ControlsListBySecureScoreOperationOptions, predicate SecureScoreControlDetailsOperationPredicate) (result ControlsListBySecureScoreCompleteResult, err error) {
	items := make([]SecureScoreControlDetails, 0)

	resp, err := c.ControlsListBySecureScore(ctx, id, options)
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

	result = ControlsListBySecureScoreCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
