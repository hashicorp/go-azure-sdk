package userexperienceanalyticsremoteconnection

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

type ListUserExperienceAnalyticsRemoteConnectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsRemoteConnection
}

type ListUserExperienceAnalyticsRemoteConnectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsRemoteConnection
}

type ListUserExperienceAnalyticsRemoteConnectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsRemoteConnectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsRemoteConnections ...
func (c UserExperienceAnalyticsRemoteConnectionClient) ListUserExperienceAnalyticsRemoteConnections(ctx context.Context) (result ListUserExperienceAnalyticsRemoteConnectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsRemoteConnectionsCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsRemoteConnection",
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
		Values *[]beta.UserExperienceAnalyticsRemoteConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsRemoteConnectionsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsRemoteConnectionClient) ListUserExperienceAnalyticsRemoteConnectionsComplete(ctx context.Context) (ListUserExperienceAnalyticsRemoteConnectionsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsRemoteConnectionsCompleteMatchingPredicate(ctx, UserExperienceAnalyticsRemoteConnectionOperationPredicate{})
}

// ListUserExperienceAnalyticsRemoteConnectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsRemoteConnectionClient) ListUserExperienceAnalyticsRemoteConnectionsCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsRemoteConnectionOperationPredicate) (result ListUserExperienceAnalyticsRemoteConnectionsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsRemoteConnection, 0)

	resp, err := c.ListUserExperienceAnalyticsRemoteConnections(ctx)
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

	result = ListUserExperienceAnalyticsRemoteConnectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
