package labels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]LabelListResult
}

type GetLabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []LabelListResult
}

type GetLabelsOperationOptions struct {
	AcceptDatetime *string
	After          *string
	Name           *string
	Select         *string
	SyncToken      *string
}

func DefaultGetLabelsOperationOptions() GetLabelsOperationOptions {
	return GetLabelsOperationOptions{}
}

func (o GetLabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.AcceptDatetime != nil {
		out.Append("Accept-Datetime", fmt.Sprintf("%v", *o.AcceptDatetime))
	}
	if o.SyncToken != nil {
		out.Append("Sync-Token", fmt.Sprintf("%v", *o.SyncToken))
	}
	return &out
}

func (o GetLabelsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetLabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.After != nil {
		out.Append("After", fmt.Sprintf("%v", *o.After))
	}
	if o.Name != nil {
		out.Append("name", fmt.Sprintf("%v", *o.Name))
	}
	if o.Select != nil {
		out.Append("$Select", fmt.Sprintf("%v", *o.Select))
	}
	return &out
}

type GetLabelsCustomPager struct {
	NextLink *odata.Link `json:"@nextLink"`
}

func (p *GetLabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetLabels ...
func (c LabelsClient) GetLabels(ctx context.Context, options GetLabelsOperationOptions) (result GetLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/vnd.microsoft.appconfig.labelset+json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GetLabelsCustomPager{},
		Path:          "/labels",
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
		Values *[]LabelListResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetLabelsComplete retrieves all the results into a single object
func (c LabelsClient) GetLabelsComplete(ctx context.Context, options GetLabelsOperationOptions) (GetLabelsCompleteResult, error) {
	return c.GetLabelsCompleteMatchingPredicate(ctx, options, LabelListResultOperationPredicate{})
}

// GetLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LabelsClient) GetLabelsCompleteMatchingPredicate(ctx context.Context, options GetLabelsOperationOptions, predicate LabelListResultOperationPredicate) (result GetLabelsCompleteResult, err error) {
	items := make([]LabelListResult, 0)

	resp, err := c.GetLabels(ctx, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = GetLabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
