package hardwarepasswordinfo

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

type ListHardwarePasswordInfosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HardwarePasswordInfo
}

type ListHardwarePasswordInfosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HardwarePasswordInfo
}

type ListHardwarePasswordInfosOperationOptions struct {
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

func DefaultListHardwarePasswordInfosOperationOptions() ListHardwarePasswordInfosOperationOptions {
	return ListHardwarePasswordInfosOperationOptions{}
}

func (o ListHardwarePasswordInfosOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHardwarePasswordInfosOperationOptions) ToOData() *odata.Query {
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

func (o ListHardwarePasswordInfosOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHardwarePasswordInfosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHardwarePasswordInfosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHardwarePasswordInfos - Get hardwarePasswordInfo from deviceManagement. Intune will provide customer the ability
// to configure BIOS configuration settings on the enrolled Windows 10 and Windows 11 Microsoft Entra joined devices.
// Starting from June, 2024, customers should start using hardwarePasswordDetail resource type - Microsoft Graph beta |
// Microsoft Learn. HardwarePasswordInfo will be marked as deprecated with Intune Release 2409
func (c HardwarePasswordInfoClient) ListHardwarePasswordInfos(ctx context.Context, options ListHardwarePasswordInfosOperationOptions) (result ListHardwarePasswordInfosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHardwarePasswordInfosCustomPager{},
		Path:          "/deviceManagement/hardwarePasswordInfo",
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
		Values *[]beta.HardwarePasswordInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHardwarePasswordInfosComplete retrieves all the results into a single object
func (c HardwarePasswordInfoClient) ListHardwarePasswordInfosComplete(ctx context.Context, options ListHardwarePasswordInfosOperationOptions) (ListHardwarePasswordInfosCompleteResult, error) {
	return c.ListHardwarePasswordInfosCompleteMatchingPredicate(ctx, options, HardwarePasswordInfoOperationPredicate{})
}

// ListHardwarePasswordInfosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HardwarePasswordInfoClient) ListHardwarePasswordInfosCompleteMatchingPredicate(ctx context.Context, options ListHardwarePasswordInfosOperationOptions, predicate HardwarePasswordInfoOperationPredicate) (result ListHardwarePasswordInfosCompleteResult, err error) {
	items := make([]beta.HardwarePasswordInfo, 0)

	resp, err := c.ListHardwarePasswordInfos(ctx, options)
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

	result = ListHardwarePasswordInfosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
