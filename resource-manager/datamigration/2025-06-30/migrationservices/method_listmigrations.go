package migrationservices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMigrationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseMigrationBase
}

type ListMigrationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DatabaseMigrationBase
}

type ListMigrationsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListMigrationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMigrations ...
func (c MigrationServicesClient) ListMigrations(ctx context.Context, id MigrationServiceId) (result ListMigrationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMigrationsCustomPager{},
		Path:       fmt.Sprintf("%s/listMigrations", id.ID()),
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
		Values *[]DatabaseMigrationBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMigrationsComplete retrieves all the results into a single object
func (c MigrationServicesClient) ListMigrationsComplete(ctx context.Context, id MigrationServiceId) (ListMigrationsCompleteResult, error) {
	return c.ListMigrationsCompleteMatchingPredicate(ctx, id, DatabaseMigrationBaseOperationPredicate{})
}

// ListMigrationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MigrationServicesClient) ListMigrationsCompleteMatchingPredicate(ctx context.Context, id MigrationServiceId, predicate DatabaseMigrationBaseOperationPredicate) (result ListMigrationsCompleteResult, err error) {
	items := make([]DatabaseMigrationBase, 0)

	resp, err := c.ListMigrations(ctx, id)
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

	result = ListMigrationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
