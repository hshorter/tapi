// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/damianoneill/tapi/restapi/operations"
	"github.com/damianoneill/tapi/restapi/operations/tapi_common"
	"github.com/damianoneill/tapi/restapi/operations/tapi_connectivity"
	"github.com/damianoneill/tapi/restapi/operations/tapi_notification"
	"github.com/damianoneill/tapi/restapi/operations/tapi_oam"
	"github.com/damianoneill/tapi/restapi/operations/tapi_path_computation"
	"github.com/damianoneill/tapi/restapi/operations/tapi_topology"
	"github.com/damianoneill/tapi/restapi/operations/tapi_virtual_network"
)

//go:generate swagger generate server --target .. --name Tapi --spec ../tapi-2.1.0.yml

func configureFlags(api *operations.TapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TapiCommonGetDataContextHandler = tapi_common.GetDataContextHandlerFunc(func(params tapi_common.GetDataContextParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContext has not yet been implemented")
	})
	api.TapiCommonGetDataContextServiceInterfacePointUUIDHandler = tapi_common.GetDataContextServiceInterfacePointUUIDHandlerFunc(func(params tapi_common.GetDataContextServiceInterfacePointUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextServiceInterfacePointUUID has not yet been implemented")
	})
	api.TapiPathComputationPostOperationsComputeP2PPathHandler = tapi_path_computation.PostOperationsComputeP2PPathHandlerFunc(func(params tapi_path_computation.PostOperationsComputeP2PPathParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_path_computation.PostOperationsComputeP2PPath has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsCreateConnectivityServiceHandler = tapi_connectivity.PostOperationsCreateConnectivityServiceHandlerFunc(func(params tapi_connectivity.PostOperationsCreateConnectivityServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsCreateConnectivityService has not yet been implemented")
	})
	api.TapiNotificationPostOperationsCreateNotificationSubscriptionServiceHandler = tapi_notification.PostOperationsCreateNotificationSubscriptionServiceHandlerFunc(func(params tapi_notification.PostOperationsCreateNotificationSubscriptionServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_notification.PostOperationsCreateNotificationSubscriptionService has not yet been implemented")
	})
	api.TapiOamPostOperationsCreateOamJobHandler = tapi_oam.PostOperationsCreateOamJobHandlerFunc(func(params tapi_oam.PostOperationsCreateOamJobParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsCreateOamJob has not yet been implemented")
	})
	api.TapiOamPostOperationsCreateOamServiceHandler = tapi_oam.PostOperationsCreateOamServiceHandlerFunc(func(params tapi_oam.PostOperationsCreateOamServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsCreateOamService has not yet been implemented")
	})
	api.TapiOamPostOperationsCreateOamServiceEndPointHandler = tapi_oam.PostOperationsCreateOamServiceEndPointHandlerFunc(func(params tapi_oam.PostOperationsCreateOamServiceEndPointParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsCreateOamServiceEndPoint has not yet been implemented")
	})
	api.TapiVirtualNetworkPostOperationsCreateVirtualNetworkServiceHandler = tapi_virtual_network.PostOperationsCreateVirtualNetworkServiceHandlerFunc(func(params tapi_virtual_network.PostOperationsCreateVirtualNetworkServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_virtual_network.PostOperationsCreateVirtualNetworkService has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsDeleteConnectivityServiceHandler = tapi_connectivity.PostOperationsDeleteConnectivityServiceHandlerFunc(func(params tapi_connectivity.PostOperationsDeleteConnectivityServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsDeleteConnectivityService has not yet been implemented")
	})
	api.TapiNotificationPostOperationsDeleteNotificationSubscriptionServiceHandler = tapi_notification.PostOperationsDeleteNotificationSubscriptionServiceHandlerFunc(func(params tapi_notification.PostOperationsDeleteNotificationSubscriptionServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_notification.PostOperationsDeleteNotificationSubscriptionService has not yet been implemented")
	})
	api.TapiOamPostOperationsDeleteOamJobHandler = tapi_oam.PostOperationsDeleteOamJobHandlerFunc(func(params tapi_oam.PostOperationsDeleteOamJobParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsDeleteOamJob has not yet been implemented")
	})
	api.TapiOamPostOperationsDeleteOamServiceHandler = tapi_oam.PostOperationsDeleteOamServiceHandlerFunc(func(params tapi_oam.PostOperationsDeleteOamServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsDeleteOamService has not yet been implemented")
	})
	api.TapiOamPostOperationsDeleteOamServiceEndPointHandler = tapi_oam.PostOperationsDeleteOamServiceEndPointHandlerFunc(func(params tapi_oam.PostOperationsDeleteOamServiceEndPointParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsDeleteOamServiceEndPoint has not yet been implemented")
	})
	api.TapiPathComputationPostOperationsDeleteP2PPathHandler = tapi_path_computation.PostOperationsDeleteP2PPathHandlerFunc(func(params tapi_path_computation.PostOperationsDeleteP2PPathParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_path_computation.PostOperationsDeleteP2PPath has not yet been implemented")
	})
	api.TapiVirtualNetworkPostOperationsDeleteVirtualNetworkServiceHandler = tapi_virtual_network.PostOperationsDeleteVirtualNetworkServiceHandlerFunc(func(params tapi_virtual_network.PostOperationsDeleteVirtualNetworkServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_virtual_network.PostOperationsDeleteVirtualNetworkService has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsGetConnectionDetailsHandler = tapi_connectivity.PostOperationsGetConnectionDetailsHandlerFunc(func(params tapi_connectivity.PostOperationsGetConnectionDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsGetConnectionDetails has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsGetConnectivityServiceDetailsHandler = tapi_connectivity.PostOperationsGetConnectivityServiceDetailsHandlerFunc(func(params tapi_connectivity.PostOperationsGetConnectivityServiceDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsGetConnectivityServiceDetails has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsGetConnectivityServiceListHandler = tapi_connectivity.PostOperationsGetConnectivityServiceListHandlerFunc(func(params tapi_connectivity.PostOperationsGetConnectivityServiceListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsGetConnectivityServiceList has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetLinkDetailsHandler = tapi_topology.PostOperationsGetLinkDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetLinkDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetLinkDetails has not yet been implemented")
	})
	api.TapiOamPostOperationsGetMegHandler = tapi_oam.PostOperationsGetMegHandlerFunc(func(params tapi_oam.PostOperationsGetMegParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsGetMeg has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetNodeDetailsHandler = tapi_topology.PostOperationsGetNodeDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetNodeDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetNodeDetails has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetNodeEdgePointDetailsHandler = tapi_topology.PostOperationsGetNodeEdgePointDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetNodeEdgePointDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetNodeEdgePointDetails has not yet been implemented")
	})
	api.TapiNotificationPostOperationsGetNotificationListHandler = tapi_notification.PostOperationsGetNotificationListHandlerFunc(func(params tapi_notification.PostOperationsGetNotificationListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_notification.PostOperationsGetNotificationList has not yet been implemented")
	})
	api.TapiNotificationPostOperationsGetNotificationSubscriptionServiceDetailsHandler = tapi_notification.PostOperationsGetNotificationSubscriptionServiceDetailsHandlerFunc(func(params tapi_notification.PostOperationsGetNotificationSubscriptionServiceDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_notification.PostOperationsGetNotificationSubscriptionServiceDetails has not yet been implemented")
	})
	api.TapiNotificationPostOperationsGetNotificationSubscriptionServiceListHandler = tapi_notification.PostOperationsGetNotificationSubscriptionServiceListHandlerFunc(func(params tapi_notification.PostOperationsGetNotificationSubscriptionServiceListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_notification.PostOperationsGetNotificationSubscriptionServiceList has not yet been implemented")
	})
	api.TapiOamPostOperationsGetOamJobHandler = tapi_oam.PostOperationsGetOamJobHandlerFunc(func(params tapi_oam.PostOperationsGetOamJobParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsGetOamJob has not yet been implemented")
	})
	api.TapiOamPostOperationsGetOamServiceHandler = tapi_oam.PostOperationsGetOamServiceHandlerFunc(func(params tapi_oam.PostOperationsGetOamServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsGetOamService has not yet been implemented")
	})
	api.TapiOamPostOperationsGetOamServiceEndPointHandler = tapi_oam.PostOperationsGetOamServiceEndPointHandlerFunc(func(params tapi_oam.PostOperationsGetOamServiceEndPointParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsGetOamServiceEndPoint has not yet been implemented")
	})
	api.TapiOamPostOperationsGetOamServiceListHandler = tapi_oam.PostOperationsGetOamServiceListHandlerFunc(func(params tapi_oam.PostOperationsGetOamServiceListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsGetOamServiceList has not yet been implemented")
	})
	api.TapiCommonPostOperationsGetServiceInterfacePointDetailsHandler = tapi_common.PostOperationsGetServiceInterfacePointDetailsHandlerFunc(func(params tapi_common.PostOperationsGetServiceInterfacePointDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.PostOperationsGetServiceInterfacePointDetails has not yet been implemented")
	})
	api.TapiCommonPostOperationsGetServiceInterfacePointListHandler = tapi_common.PostOperationsGetServiceInterfacePointListHandlerFunc(func(params tapi_common.PostOperationsGetServiceInterfacePointListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.PostOperationsGetServiceInterfacePointList has not yet been implemented")
	})
	api.TapiNotificationPostOperationsGetSupportedNotificationTypesHandler = tapi_notification.PostOperationsGetSupportedNotificationTypesHandlerFunc(func(params tapi_notification.PostOperationsGetSupportedNotificationTypesParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_notification.PostOperationsGetSupportedNotificationTypes has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetTopologyDetailsHandler = tapi_topology.PostOperationsGetTopologyDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetTopologyDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetTopologyDetails has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetTopologyListHandler = tapi_topology.PostOperationsGetTopologyListHandlerFunc(func(params tapi_topology.PostOperationsGetTopologyListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetTopologyList has not yet been implemented")
	})
	api.TapiVirtualNetworkPostOperationsGetVirtualNetworkServiceDetailsHandler = tapi_virtual_network.PostOperationsGetVirtualNetworkServiceDetailsHandlerFunc(func(params tapi_virtual_network.PostOperationsGetVirtualNetworkServiceDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_virtual_network.PostOperationsGetVirtualNetworkServiceDetails has not yet been implemented")
	})
	api.TapiVirtualNetworkPostOperationsGetVirtualNetworkServiceListHandler = tapi_virtual_network.PostOperationsGetVirtualNetworkServiceListHandlerFunc(func(params tapi_virtual_network.PostOperationsGetVirtualNetworkServiceListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_virtual_network.PostOperationsGetVirtualNetworkServiceList has not yet been implemented")
	})
	api.TapiPathComputationPostOperationsOptimizeP2PPathHandler = tapi_path_computation.PostOperationsOptimizeP2PPathHandlerFunc(func(params tapi_path_computation.PostOperationsOptimizeP2PPathParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_path_computation.PostOperationsOptimizeP2PPath has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsUpdateConnectivityServiceHandler = tapi_connectivity.PostOperationsUpdateConnectivityServiceHandlerFunc(func(params tapi_connectivity.PostOperationsUpdateConnectivityServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsUpdateConnectivityService has not yet been implemented")
	})
	api.TapiNotificationPostOperationsUpdateNotificationSubscriptionServiceHandler = tapi_notification.PostOperationsUpdateNotificationSubscriptionServiceHandlerFunc(func(params tapi_notification.PostOperationsUpdateNotificationSubscriptionServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_notification.PostOperationsUpdateNotificationSubscriptionService has not yet been implemented")
	})
	api.TapiOamPostOperationsUpdateOamJobHandler = tapi_oam.PostOperationsUpdateOamJobHandlerFunc(func(params tapi_oam.PostOperationsUpdateOamJobParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsUpdateOamJob has not yet been implemented")
	})
	api.TapiOamPostOperationsUpdateOamServiceHandler = tapi_oam.PostOperationsUpdateOamServiceHandlerFunc(func(params tapi_oam.PostOperationsUpdateOamServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsUpdateOamService has not yet been implemented")
	})
	api.TapiOamPostOperationsUpdateOamServiceEndPointHandler = tapi_oam.PostOperationsUpdateOamServiceEndPointHandlerFunc(func(params tapi_oam.PostOperationsUpdateOamServiceEndPointParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_oam.PostOperationsUpdateOamServiceEndPoint has not yet been implemented")
	})
	api.TapiCommonPostOperationsUpdateServiceInterfacePointHandler = tapi_common.PostOperationsUpdateServiceInterfacePointHandlerFunc(func(params tapi_common.PostOperationsUpdateServiceInterfacePointParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.PostOperationsUpdateServiceInterfacePoint has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
