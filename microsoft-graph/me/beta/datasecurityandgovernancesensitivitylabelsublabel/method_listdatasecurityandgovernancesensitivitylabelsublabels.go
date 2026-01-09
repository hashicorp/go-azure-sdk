package datasecurityandgovernancesensitivitylabelsublabel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SensitivityLabel
}

type ListDataSecurityAndGovernanceSensitivityLabelSublabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SensitivityLabel
}

type ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions struct {
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

func DefaultListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions() ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions {
	return ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions{}
}

func (o ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions) ToOData() *odata.Query {
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

func (o ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDataSecurityAndGovernanceSensitivityLabelSublabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDataSecurityAndGovernanceSensitivityLabelSublabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDataSecurityAndGovernanceSensitivityLabelSublabels - Get sublabels from me
func (c DataSecurityAndGovernanceSensitivityLabelSublabelClient) ListDataSecurityAndGovernanceSensitivityLabelSublabels(ctx context.Context, id beta.MeDataSecurityAndGovernanceSensitivityLabelId, options ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions) (result ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDataSecurityAndGovernanceSensitivityLabelSublabelsCustomPager{},
		Path:          fmt.Sprintf("%s/sublabels", id.ID()),
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

// ListDataSecurityAndGovernanceSensitivityLabelSublabelsComplete retrieves all the results into a single object
func (c DataSecurityAndGovernanceSensitivityLabelSublabelClient) ListDataSecurityAndGovernanceSensitivityLabelSublabelsComplete(ctx context.Context, id beta.MeDataSecurityAndGovernanceSensitivityLabelId, options ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions) (ListDataSecurityAndGovernanceSensitivityLabelSublabelsCompleteResult, error) {
	return c.ListDataSecurityAndGovernanceSensitivityLabelSublabelsCompleteMatchingPredicate(ctx, id, options, SensitivityLabelOperationPredicate{})
}

// ListDataSecurityAndGovernanceSensitivityLabelSublabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataSecurityAndGovernanceSensitivityLabelSublabelClient) ListDataSecurityAndGovernanceSensitivityLabelSublabelsCompleteMatchingPredicate(ctx context.Context, id beta.MeDataSecurityAndGovernanceSensitivityLabelId, options ListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions, predicate SensitivityLabelOperationPredicate) (result ListDataSecurityAndGovernanceSensitivityLabelSublabelsCompleteResult, err error) {
	items := make([]beta.SensitivityLabel, 0)

	resp, err := c.ListDataSecurityAndGovernanceSensitivityLabelSublabels(ctx, id, options)
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

	result = ListDataSecurityAndGovernanceSensitivityLabelSublabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
