package grouppolicymigrationreportgrouppolicysettingmapping

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

type ListGroupPolicyMigrationReportGroupPolicySettingMappingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicySettingMapping
}

type ListGroupPolicyMigrationReportGroupPolicySettingMappingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicySettingMapping
}

type ListGroupPolicyMigrationReportGroupPolicySettingMappingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyMigrationReportGroupPolicySettingMappingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyMigrationReportGroupPolicySettingMappings ...
func (c GroupPolicyMigrationReportGroupPolicySettingMappingClient) ListGroupPolicyMigrationReportGroupPolicySettingMappings(ctx context.Context, id DeviceManagementGroupPolicyMigrationReportId) (result ListGroupPolicyMigrationReportGroupPolicySettingMappingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyMigrationReportGroupPolicySettingMappingsCustomPager{},
		Path:       fmt.Sprintf("%s/groupPolicySettingMappings", id.ID()),
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
		Values *[]beta.GroupPolicySettingMapping `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyMigrationReportGroupPolicySettingMappingsComplete retrieves all the results into a single object
func (c GroupPolicyMigrationReportGroupPolicySettingMappingClient) ListGroupPolicyMigrationReportGroupPolicySettingMappingsComplete(ctx context.Context, id DeviceManagementGroupPolicyMigrationReportId) (ListGroupPolicyMigrationReportGroupPolicySettingMappingsCompleteResult, error) {
	return c.ListGroupPolicyMigrationReportGroupPolicySettingMappingsCompleteMatchingPredicate(ctx, id, GroupPolicySettingMappingOperationPredicate{})
}

// ListGroupPolicyMigrationReportGroupPolicySettingMappingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyMigrationReportGroupPolicySettingMappingClient) ListGroupPolicyMigrationReportGroupPolicySettingMappingsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyMigrationReportId, predicate GroupPolicySettingMappingOperationPredicate) (result ListGroupPolicyMigrationReportGroupPolicySettingMappingsCompleteResult, err error) {
	items := make([]beta.GroupPolicySettingMapping, 0)

	resp, err := c.ListGroupPolicyMigrationReportGroupPolicySettingMappings(ctx, id)
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

	result = ListGroupPolicyMigrationReportGroupPolicySettingMappingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
