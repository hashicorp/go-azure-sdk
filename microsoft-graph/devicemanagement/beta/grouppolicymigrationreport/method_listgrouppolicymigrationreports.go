package grouppolicymigrationreport

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

type ListGroupPolicyMigrationReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyMigrationReport
}

type ListGroupPolicyMigrationReportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyMigrationReport
}

type ListGroupPolicyMigrationReportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyMigrationReportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyMigrationReports ...
func (c GroupPolicyMigrationReportClient) ListGroupPolicyMigrationReports(ctx context.Context) (result ListGroupPolicyMigrationReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyMigrationReportsCustomPager{},
		Path:       "/deviceManagement/groupPolicyMigrationReports",
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
		Values *[]beta.GroupPolicyMigrationReport `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyMigrationReportsComplete retrieves all the results into a single object
func (c GroupPolicyMigrationReportClient) ListGroupPolicyMigrationReportsComplete(ctx context.Context) (ListGroupPolicyMigrationReportsCompleteResult, error) {
	return c.ListGroupPolicyMigrationReportsCompleteMatchingPredicate(ctx, GroupPolicyMigrationReportOperationPredicate{})
}

// ListGroupPolicyMigrationReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyMigrationReportClient) ListGroupPolicyMigrationReportsCompleteMatchingPredicate(ctx context.Context, predicate GroupPolicyMigrationReportOperationPredicate) (result ListGroupPolicyMigrationReportsCompleteResult, err error) {
	items := make([]beta.GroupPolicyMigrationReport, 0)

	resp, err := c.ListGroupPolicyMigrationReports(ctx)
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

	result = ListGroupPolicyMigrationReportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
