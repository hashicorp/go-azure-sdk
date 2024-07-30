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

type ListResourceAccessProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileBase
}

type ListResourceAccessProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileBase
}

type ListResourceAccessProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListResourceAccessProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResourceAccessProfiles ...
func (c ResourceAccessProfileClient) ListResourceAccessProfiles(ctx context.Context) (result ListResourceAccessProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListResourceAccessProfilesCustomPager{},
		Path:       "/deviceManagement/resourceAccessProfiles",
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

// ListResourceAccessProfilesComplete retrieves all the results into a single object
func (c ResourceAccessProfileClient) ListResourceAccessProfilesComplete(ctx context.Context) (ListResourceAccessProfilesCompleteResult, error) {
	return c.ListResourceAccessProfilesCompleteMatchingPredicate(ctx, DeviceManagementResourceAccessProfileBaseOperationPredicate{})
}

// ListResourceAccessProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileClient) ListResourceAccessProfilesCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementResourceAccessProfileBaseOperationPredicate) (result ListResourceAccessProfilesCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileBase, 0)

	resp, err := c.ListResourceAccessProfiles(ctx)
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

	result = ListResourceAccessProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
