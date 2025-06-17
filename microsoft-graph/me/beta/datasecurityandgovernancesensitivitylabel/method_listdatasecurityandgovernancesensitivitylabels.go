package datasecurityandgovernancesensitivitylabel

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

type ListDataSecurityAndGovernanceSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SensitivityLabel
}

type ListDataSecurityAndGovernanceSensitivityLabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SensitivityLabel
}

type ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListDataSecurityAndGovernanceSensitivityLabelsOperationOptions() ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions {
	return ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions{}
}

func (o ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDataSecurityAndGovernanceSensitivityLabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDataSecurityAndGovernanceSensitivityLabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDataSecurityAndGovernanceSensitivityLabels - Get sensitivityLabels from me
func (c DataSecurityAndGovernanceSensitivityLabelClient) ListDataSecurityAndGovernanceSensitivityLabels(ctx context.Context, options ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions) (result ListDataSecurityAndGovernanceSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDataSecurityAndGovernanceSensitivityLabelsCustomPager{},
		Path:          "/me/dataSecurityAndGovernance/sensitivityLabels",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.SensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDataSecurityAndGovernanceSensitivityLabelsComplete retrieves all the results into a single object
func (c DataSecurityAndGovernanceSensitivityLabelClient) ListDataSecurityAndGovernanceSensitivityLabelsComplete(ctx context.Context, options ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions) (ListDataSecurityAndGovernanceSensitivityLabelsCompleteResult, error) {
	return c.ListDataSecurityAndGovernanceSensitivityLabelsCompleteMatchingPredicate(ctx, options, SensitivityLabelOperationPredicate{})
}

// ListDataSecurityAndGovernanceSensitivityLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataSecurityAndGovernanceSensitivityLabelClient) ListDataSecurityAndGovernanceSensitivityLabelsCompleteMatchingPredicate(ctx context.Context, options ListDataSecurityAndGovernanceSensitivityLabelsOperationOptions, predicate SensitivityLabelOperationPredicate) (result ListDataSecurityAndGovernanceSensitivityLabelsCompleteResult, err error) {
	items := make([]beta.SensitivityLabel, 0)

	resp, err := c.ListDataSecurityAndGovernanceSensitivityLabels(ctx, options)
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

	result = ListDataSecurityAndGovernanceSensitivityLabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
