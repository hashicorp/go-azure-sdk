package userexperienceanalyticsdevicestartupprocess

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserExperienceAnalyticsDeviceStartupProcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsDeviceStartupProcess
}

type ListUserExperienceAnalyticsDeviceStartupProcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsDeviceStartupProcess
}

type ListUserExperienceAnalyticsDeviceStartupProcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceStartupProcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceStartupProces ...
func (c UserExperienceAnalyticsDeviceStartupProcessClient) ListUserExperienceAnalyticsDeviceStartupProces(ctx context.Context) (result ListUserExperienceAnalyticsDeviceStartupProcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserExperienceAnalyticsDeviceStartupProcesCustomPager{},
		Path:       "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses",
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
		Values *[]stable.UserExperienceAnalyticsDeviceStartupProcess `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceStartupProcesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceStartupProcessClient) ListUserExperienceAnalyticsDeviceStartupProcesComplete(ctx context.Context) (ListUserExperienceAnalyticsDeviceStartupProcesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceStartupProcesCompleteMatchingPredicate(ctx, UserExperienceAnalyticsDeviceStartupProcessOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceStartupProcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceStartupProcessClient) ListUserExperienceAnalyticsDeviceStartupProcesCompleteMatchingPredicate(ctx context.Context, predicate UserExperienceAnalyticsDeviceStartupProcessOperationPredicate) (result ListUserExperienceAnalyticsDeviceStartupProcesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsDeviceStartupProcess, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceStartupProces(ctx)
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

	result = ListUserExperienceAnalyticsDeviceStartupProcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
