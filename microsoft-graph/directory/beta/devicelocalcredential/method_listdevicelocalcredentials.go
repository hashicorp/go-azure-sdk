package devicelocalcredential

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

type ListDeviceLocalCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceLocalCredentialInfo
}

type ListDeviceLocalCredentialsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceLocalCredentialInfo
}

type ListDeviceLocalCredentialsOperationOptions struct {
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

func DefaultListDeviceLocalCredentialsOperationOptions() ListDeviceLocalCredentialsOperationOptions {
	return ListDeviceLocalCredentialsOperationOptions{}
}

func (o ListDeviceLocalCredentialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceLocalCredentialsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceLocalCredentialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceLocalCredentialsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceLocalCredentialsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceLocalCredentials - List deviceLocalCredentialInfo. Get a list of the deviceLocalCredentialInfo objects and
// their properties excluding the credentials.
func (c DeviceLocalCredentialClient) ListDeviceLocalCredentials(ctx context.Context, options ListDeviceLocalCredentialsOperationOptions) (result ListDeviceLocalCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceLocalCredentialsCustomPager{},
		Path:          "/directory/deviceLocalCredentials",
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
		Values *[]beta.DeviceLocalCredentialInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceLocalCredentialsComplete retrieves all the results into a single object
func (c DeviceLocalCredentialClient) ListDeviceLocalCredentialsComplete(ctx context.Context, options ListDeviceLocalCredentialsOperationOptions) (ListDeviceLocalCredentialsCompleteResult, error) {
	return c.ListDeviceLocalCredentialsCompleteMatchingPredicate(ctx, options, DeviceLocalCredentialInfoOperationPredicate{})
}

// ListDeviceLocalCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceLocalCredentialClient) ListDeviceLocalCredentialsCompleteMatchingPredicate(ctx context.Context, options ListDeviceLocalCredentialsOperationOptions, predicate DeviceLocalCredentialInfoOperationPredicate) (result ListDeviceLocalCredentialsCompleteResult, err error) {
	items := make([]beta.DeviceLocalCredentialInfo, 0)

	resp, err := c.ListDeviceLocalCredentials(ctx, options)
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

	result = ListDeviceLocalCredentialsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
