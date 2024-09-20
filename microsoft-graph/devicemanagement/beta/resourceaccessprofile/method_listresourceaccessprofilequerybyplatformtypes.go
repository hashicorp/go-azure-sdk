package resourceaccessprofile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListResourceAccessProfileQueryByPlatformTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileBase
}

type ListResourceAccessProfileQueryByPlatformTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileBase
}

type ListResourceAccessProfileQueryByPlatformTypesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListResourceAccessProfileQueryByPlatformTypesOperationOptions() ListResourceAccessProfileQueryByPlatformTypesOperationOptions {
	return ListResourceAccessProfileQueryByPlatformTypesOperationOptions{}
}

func (o ListResourceAccessProfileQueryByPlatformTypesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListResourceAccessProfileQueryByPlatformTypesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListResourceAccessProfileQueryByPlatformTypesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListResourceAccessProfileQueryByPlatformTypesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListResourceAccessProfileQueryByPlatformTypesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResourceAccessProfileQueryByPlatformTypes - Invoke action queryByPlatformType
func (c ResourceAccessProfileClient) ListResourceAccessProfileQueryByPlatformTypes(ctx context.Context, input ListResourceAccessProfileQueryByPlatformTypesRequest, options ListResourceAccessProfileQueryByPlatformTypesOperationOptions) (result ListResourceAccessProfileQueryByPlatformTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListResourceAccessProfileQueryByPlatformTypesCustomPager{},
		Path:          "/deviceManagement/resourceAccessProfiles/queryByPlatformType",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DeviceManagementResourceAccessProfileBase, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDeviceManagementResourceAccessProfileBaseImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DeviceManagementResourceAccessProfileBase (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListResourceAccessProfileQueryByPlatformTypesComplete retrieves all the results into a single object
func (c ResourceAccessProfileClient) ListResourceAccessProfileQueryByPlatformTypesComplete(ctx context.Context, input ListResourceAccessProfileQueryByPlatformTypesRequest, options ListResourceAccessProfileQueryByPlatformTypesOperationOptions) (ListResourceAccessProfileQueryByPlatformTypesCompleteResult, error) {
	return c.ListResourceAccessProfileQueryByPlatformTypesCompleteMatchingPredicate(ctx, input, options, DeviceManagementResourceAccessProfileBaseOperationPredicate{})
}

// ListResourceAccessProfileQueryByPlatformTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileClient) ListResourceAccessProfileQueryByPlatformTypesCompleteMatchingPredicate(ctx context.Context, input ListResourceAccessProfileQueryByPlatformTypesRequest, options ListResourceAccessProfileQueryByPlatformTypesOperationOptions, predicate DeviceManagementResourceAccessProfileBaseOperationPredicate) (result ListResourceAccessProfileQueryByPlatformTypesCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileBase, 0)

	resp, err := c.ListResourceAccessProfileQueryByPlatformTypes(ctx, input, options)
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

	result = ListResourceAccessProfileQueryByPlatformTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
