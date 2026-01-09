package hardwarepassworddetail

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

type ListHardwarePasswordDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HardwarePasswordDetail
}

type ListHardwarePasswordDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HardwarePasswordDetail
}

type ListHardwarePasswordDetailsOperationOptions struct {
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

func DefaultListHardwarePasswordDetailsOperationOptions() ListHardwarePasswordDetailsOperationOptions {
	return ListHardwarePasswordDetailsOperationOptions{}
}

func (o ListHardwarePasswordDetailsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHardwarePasswordDetailsOperationOptions) ToOData() *odata.Query {
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

func (o ListHardwarePasswordDetailsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHardwarePasswordDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHardwarePasswordDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHardwarePasswordDetails - Get hardwarePasswordDetails from deviceManagement. Device BIOS password information for
// devices with managed BIOS and firmware configuration, which provides device serial number, list of previous
// passwords, and current password.
func (c HardwarePasswordDetailClient) ListHardwarePasswordDetails(ctx context.Context, options ListHardwarePasswordDetailsOperationOptions) (result ListHardwarePasswordDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHardwarePasswordDetailsCustomPager{},
		Path:          "/deviceManagement/hardwarePasswordDetails",
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
		Values *[]beta.HardwarePasswordDetail `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHardwarePasswordDetailsComplete retrieves all the results into a single object
func (c HardwarePasswordDetailClient) ListHardwarePasswordDetailsComplete(ctx context.Context, options ListHardwarePasswordDetailsOperationOptions) (ListHardwarePasswordDetailsCompleteResult, error) {
	return c.ListHardwarePasswordDetailsCompleteMatchingPredicate(ctx, options, HardwarePasswordDetailOperationPredicate{})
}

// ListHardwarePasswordDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HardwarePasswordDetailClient) ListHardwarePasswordDetailsCompleteMatchingPredicate(ctx context.Context, options ListHardwarePasswordDetailsOperationOptions, predicate HardwarePasswordDetailOperationPredicate) (result ListHardwarePasswordDetailsCompleteResult, err error) {
	items := make([]beta.HardwarePasswordDetail, 0)

	resp, err := c.ListHardwarePasswordDetails(ctx, options)
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

	result = ListHardwarePasswordDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
