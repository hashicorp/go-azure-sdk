package resourceaccessprofile

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

type ListResourceAccessProfileQueryByPlatformTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileBase
}

type ListResourceAccessProfileQueryByPlatformTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileBase
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

// ListResourceAccessProfileQueryByPlatformTypes ...
func (c ResourceAccessProfileClient) ListResourceAccessProfileQueryByPlatformTypes(ctx context.Context, input ListResourceAccessProfileQueryByPlatformTypesRequest) (result ListResourceAccessProfileQueryByPlatformTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListResourceAccessProfileQueryByPlatformTypesCustomPager{},
		Path:       "/deviceManagement/resourceAccessProfiles/queryByPlatformType",
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
		Values *[]beta.DeviceManagementResourceAccessProfileBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListResourceAccessProfileQueryByPlatformTypesComplete retrieves all the results into a single object
func (c ResourceAccessProfileClient) ListResourceAccessProfileQueryByPlatformTypesComplete(ctx context.Context, input ListResourceAccessProfileQueryByPlatformTypesRequest) (ListResourceAccessProfileQueryByPlatformTypesCompleteResult, error) {
	return c.ListResourceAccessProfileQueryByPlatformTypesCompleteMatchingPredicate(ctx, input, DeviceManagementResourceAccessProfileBaseOperationPredicate{})
}

// ListResourceAccessProfileQueryByPlatformTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileClient) ListResourceAccessProfileQueryByPlatformTypesCompleteMatchingPredicate(ctx context.Context, input ListResourceAccessProfileQueryByPlatformTypesRequest, predicate DeviceManagementResourceAccessProfileBaseOperationPredicate) (result ListResourceAccessProfileQueryByPlatformTypesCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileBase, 0)

	resp, err := c.ListResourceAccessProfileQueryByPlatformTypes(ctx, input)
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
