package containerappsrevisionreplicas

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type GetReplicaOperationResponse struct {
	HttpResponse *http.Response
	Model        *Replica
}

// GetReplica ...
func (c ContainerAppsRevisionReplicasClient) GetReplica(ctx context.Context, id ReplicaId) (result GetReplicaOperationResponse, err error) {
	req, err := c.preparerForGetReplica(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsrevisionreplicas.ContainerAppsRevisionReplicasClient", "GetReplica", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsrevisionreplicas.ContainerAppsRevisionReplicasClient", "GetReplica", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetReplica(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsrevisionreplicas.ContainerAppsRevisionReplicasClient", "GetReplica", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetReplica prepares the GetReplica request.
func (c ContainerAppsRevisionReplicasClient) preparerForGetReplica(ctx context.Context, id ReplicaId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetReplica handles the response to the GetReplica request. The method always
// closes the http.Response Body.
func (c ContainerAppsRevisionReplicasClient) responderForGetReplica(resp *http.Response) (result GetReplicaOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
