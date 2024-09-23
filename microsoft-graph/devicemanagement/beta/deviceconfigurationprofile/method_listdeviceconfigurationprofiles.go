package deviceconfigurationprofile

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

type ListDeviceConfigurationProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationProfile
}

type ListDeviceConfigurationProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationProfile
}

type ListDeviceConfigurationProfilesOperationOptions struct {
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

func DefaultListDeviceConfigurationProfilesOperationOptions() ListDeviceConfigurationProfilesOperationOptions {
	return ListDeviceConfigurationProfilesOperationOptions{}
}

func (o ListDeviceConfigurationProfilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceConfigurationProfilesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceConfigurationProfilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceConfigurationProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationProfiles - Get deviceConfigurationProfiles from deviceManagement. Profile Id of the object.
func (c DeviceConfigurationProfileClient) ListDeviceConfigurationProfiles(ctx context.Context, options ListDeviceConfigurationProfilesOperationOptions) (result ListDeviceConfigurationProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceConfigurationProfilesCustomPager{},
		Path:          "/deviceManagement/deviceConfigurationProfiles",
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
		Values *[]beta.DeviceConfigurationProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationProfilesComplete retrieves all the results into a single object
func (c DeviceConfigurationProfileClient) ListDeviceConfigurationProfilesComplete(ctx context.Context, options ListDeviceConfigurationProfilesOperationOptions) (ListDeviceConfigurationProfilesCompleteResult, error) {
	return c.ListDeviceConfigurationProfilesCompleteMatchingPredicate(ctx, options, DeviceConfigurationProfileOperationPredicate{})
}

// ListDeviceConfigurationProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationProfileClient) ListDeviceConfigurationProfilesCompleteMatchingPredicate(ctx context.Context, options ListDeviceConfigurationProfilesOperationOptions, predicate DeviceConfigurationProfileOperationPredicate) (result ListDeviceConfigurationProfilesCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationProfile, 0)

	resp, err := c.ListDeviceConfigurationProfiles(ctx, options)
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

	result = ListDeviceConfigurationProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
