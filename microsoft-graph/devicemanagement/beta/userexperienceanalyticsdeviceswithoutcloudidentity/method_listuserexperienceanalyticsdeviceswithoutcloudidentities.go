package userexperienceanalyticsdeviceswithoutcloudidentity

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

type ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsDeviceWithoutCloudIdentity
}

type ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsDeviceWithoutCloudIdentity
}

type ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDevicesWithoutCloudIdentities ...
func (c UserExperienceAnalyticsDevicesWithoutCloudIdentityClient) ListUserExperienceAnalyticsDevicesWithoutCloudIdentities(ctx context.Context) (result ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsDevicesWithoutCloudIdentity",
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
		Values *[]beta.UserExperienceAnalyticsDeviceWithoutCloudIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDevicesWithoutCloudIdentityClient) ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesComplete(ctx context.Context) (ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsDeviceWithoutCloudIdentityOperationPredicate{})
}

// ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDevicesWithoutCloudIdentityClient) ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsDeviceWithoutCloudIdentityOperationPredicate) (result ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsDeviceWithoutCloudIdentity, 0)

	resp, err := c.ListUserExperienceAnalyticsDevicesWithoutCloudIdentities(ctx)
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

	result = ListUserExperienceAnalyticsDevicesWithoutCloudIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
