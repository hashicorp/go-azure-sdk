package connectordefinitions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorDefinitionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataConnectorDefinition
}

type DataConnectorDefinitionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DataConnectorDefinition
}

type DataConnectorDefinitionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DataConnectorDefinitionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DataConnectorDefinitionsList ...
func (c ConnectorDefinitionsClient) DataConnectorDefinitionsList(ctx context.Context, id WorkspaceId) (result DataConnectorDefinitionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DataConnectorDefinitionsListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/dataConnectorDefinitions", id.ID()),
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

	temp := make([]DataConnectorDefinition, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalDataConnectorDefinitionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for DataConnectorDefinition (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// DataConnectorDefinitionsListComplete retrieves all the results into a single object
func (c ConnectorDefinitionsClient) DataConnectorDefinitionsListComplete(ctx context.Context, id WorkspaceId) (DataConnectorDefinitionsListCompleteResult, error) {
	return c.DataConnectorDefinitionsListCompleteMatchingPredicate(ctx, id, DataConnectorDefinitionOperationPredicate{})
}

// DataConnectorDefinitionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConnectorDefinitionsClient) DataConnectorDefinitionsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate DataConnectorDefinitionOperationPredicate) (result DataConnectorDefinitionsListCompleteResult, err error) {
	items := make([]DataConnectorDefinition, 0)

	resp, err := c.DataConnectorDefinitionsList(ctx, id)
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

	result = DataConnectorDefinitionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
