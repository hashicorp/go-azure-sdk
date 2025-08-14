package databasemigrations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoToCosmosDbvCoreMongoGetForScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatabaseMigrationCosmosDbMongo
}

type MongoToCosmosDbvCoreMongoGetForScopeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DatabaseMigrationCosmosDbMongo
}

type MongoToCosmosDbvCoreMongoGetForScopeCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *MongoToCosmosDbvCoreMongoGetForScopeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// MongoToCosmosDbvCoreMongoGetForScope ...
func (c DatabaseMigrationsClient) MongoToCosmosDbvCoreMongoGetForScope(ctx context.Context, id MongoClusterId) (result MongoToCosmosDbvCoreMongoGetForScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &MongoToCosmosDbvCoreMongoGetForScopeCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.DataMigration/databaseMigrations", id.ID()),
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
		Values *[]DatabaseMigrationCosmosDbMongo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MongoToCosmosDbvCoreMongoGetForScopeComplete retrieves all the results into a single object
func (c DatabaseMigrationsClient) MongoToCosmosDbvCoreMongoGetForScopeComplete(ctx context.Context, id MongoClusterId) (MongoToCosmosDbvCoreMongoGetForScopeCompleteResult, error) {
	return c.MongoToCosmosDbvCoreMongoGetForScopeCompleteMatchingPredicate(ctx, id, DatabaseMigrationCosmosDbMongoOperationPredicate{})
}

// MongoToCosmosDbvCoreMongoGetForScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatabaseMigrationsClient) MongoToCosmosDbvCoreMongoGetForScopeCompleteMatchingPredicate(ctx context.Context, id MongoClusterId, predicate DatabaseMigrationCosmosDbMongoOperationPredicate) (result MongoToCosmosDbvCoreMongoGetForScopeCompleteResult, err error) {
	items := make([]DatabaseMigrationCosmosDbMongo, 0)

	resp, err := c.MongoToCosmosDbvCoreMongoGetForScope(ctx, id)
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

	result = MongoToCosmosDbvCoreMongoGetForScopeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
