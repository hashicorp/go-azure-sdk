package v2024_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/adminrulecollections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/adminrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/applicationgatewayprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/applicationgatewayprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/applicationgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/applicationgatewaywafdynamicmanifests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/applicationsecuritygroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/availabledelegations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/availableservicealiases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/azurefirewalls"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/bastionhosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/bastionshareablelink"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/bgpservicecommunities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/checkdnsavailabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/cloudservicepublicipaddresses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/connectionmonitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/connectivityconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/customipprefixes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/ddoscustompolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/ddosprotectionplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/dscpconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/dscpconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/endpointservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuitarptable"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuitauthorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuitconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuitpeerings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuitroutestable"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuitroutestablesummary"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuits"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecircuitstats"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressrouteconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecrossconnectionarptable"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecrossconnectionpeerings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecrossconnectionroutetable"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecrossconnectionroutetablesummary"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutecrossconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutegateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressroutelinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressrouteportauthorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressrouteports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressrouteportslocations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressrouteproviderports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/expressrouteserviceproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/firewallpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/firewallpolicyrulecollectiongroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/flowlogs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/ipallocations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/ipampools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/ipgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/loadbalancers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/localnetworkgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/natgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkinterfaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkmanageractiveconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkmanageractiveconnectivityconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkmanagerconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkmanagereffectiveconnectivityconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkmanagereffectivesecurityadminrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkmanagerroutingconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkmanagers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecuritygroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeteraccessrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeterassociableresourcetypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeterassociations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeterlinkreferences"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeterlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeterloggingconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeterprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networksecurityperimeters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkvirtualappliances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkwatchers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/p2svpngateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/packetcaptures"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/peerexpressroutecircuitconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/privatednszonegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/privateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/privatelinkservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/publicipaddresses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/publicipprefixes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/reachabilityanalysisintent"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/reachabilityanalysisintents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/reachabilityanalysisruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/routefilterrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/routefilters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/routes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/routetables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/routingrulecollections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/routingrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/scopeconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/securityadminconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/securitypartnerproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/securityrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/securityuserconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/securityuserrulecollections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/securityuserrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/serviceendpointpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/serviceendpointpolicydefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/servicetags"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/staticcidrs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/staticmembers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/subnets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/trafficanalytics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/verifierworkspaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/vipswap"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualappliancesites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualapplianceskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworkgatewayconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworkgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworkpeerings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworktap"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworktaps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualrouterpeerings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualrouters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualwans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/vmsspublicipaddresses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/vpngateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/vpnlinkconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/vpnserverconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/vpnsites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/webapplicationfirewallpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/webcategories"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdminRuleCollections                             *adminrulecollections.AdminRuleCollectionsClient
	AdminRules                                       *adminrules.AdminRulesClient
	ApplicationGatewayPrivateEndpointConnections     *applicationgatewayprivateendpointconnections.ApplicationGatewayPrivateEndpointConnectionsClient
	ApplicationGatewayPrivateLinkResources           *applicationgatewayprivatelinkresources.ApplicationGatewayPrivateLinkResourcesClient
	ApplicationGatewayWafDynamicManifests            *applicationgatewaywafdynamicmanifests.ApplicationGatewayWafDynamicManifestsClient
	ApplicationGateways                              *applicationgateways.ApplicationGatewaysClient
	ApplicationSecurityGroups                        *applicationsecuritygroups.ApplicationSecurityGroupsClient
	AvailableDelegations                             *availabledelegations.AvailableDelegationsClient
	AvailableServiceAliases                          *availableservicealiases.AvailableServiceAliasesClient
	AzureFirewalls                                   *azurefirewalls.AzureFirewallsClient
	BastionHosts                                     *bastionhosts.BastionHostsClient
	BastionShareableLink                             *bastionshareablelink.BastionShareableLinkClient
	BgpServiceCommunities                            *bgpservicecommunities.BgpServiceCommunitiesClient
	CheckDnsAvailabilities                           *checkdnsavailabilities.CheckDnsAvailabilitiesClient
	CloudServicePublicIPAddresses                    *cloudservicepublicipaddresses.CloudServicePublicIPAddressesClient
	ConnectionMonitors                               *connectionmonitors.ConnectionMonitorsClient
	ConnectivityConfigurations                       *connectivityconfigurations.ConnectivityConfigurationsClient
	CustomIPPrefixes                                 *customipprefixes.CustomIPPrefixesClient
	DdosCustomPolicies                               *ddoscustompolicies.DdosCustomPoliciesClient
	DdosProtectionPlans                              *ddosprotectionplans.DdosProtectionPlansClient
	DscpConfiguration                                *dscpconfiguration.DscpConfigurationClient
	DscpConfigurations                               *dscpconfigurations.DscpConfigurationsClient
	EndpointServices                                 *endpointservices.EndpointServicesClient
	ExpressRouteCircuitArpTable                      *expressroutecircuitarptable.ExpressRouteCircuitArpTableClient
	ExpressRouteCircuitAuthorizations                *expressroutecircuitauthorizations.ExpressRouteCircuitAuthorizationsClient
	ExpressRouteCircuitConnections                   *expressroutecircuitconnections.ExpressRouteCircuitConnectionsClient
	ExpressRouteCircuitPeerings                      *expressroutecircuitpeerings.ExpressRouteCircuitPeeringsClient
	ExpressRouteCircuitRoutesTable                   *expressroutecircuitroutestable.ExpressRouteCircuitRoutesTableClient
	ExpressRouteCircuitRoutesTableSummary            *expressroutecircuitroutestablesummary.ExpressRouteCircuitRoutesTableSummaryClient
	ExpressRouteCircuitStats                         *expressroutecircuitstats.ExpressRouteCircuitStatsClient
	ExpressRouteCircuits                             *expressroutecircuits.ExpressRouteCircuitsClient
	ExpressRouteConnections                          *expressrouteconnections.ExpressRouteConnectionsClient
	ExpressRouteCrossConnectionArpTable              *expressroutecrossconnectionarptable.ExpressRouteCrossConnectionArpTableClient
	ExpressRouteCrossConnectionPeerings              *expressroutecrossconnectionpeerings.ExpressRouteCrossConnectionPeeringsClient
	ExpressRouteCrossConnectionRouteTable            *expressroutecrossconnectionroutetable.ExpressRouteCrossConnectionRouteTableClient
	ExpressRouteCrossConnectionRouteTableSummary     *expressroutecrossconnectionroutetablesummary.ExpressRouteCrossConnectionRouteTableSummaryClient
	ExpressRouteCrossConnections                     *expressroutecrossconnections.ExpressRouteCrossConnectionsClient
	ExpressRouteGateways                             *expressroutegateways.ExpressRouteGatewaysClient
	ExpressRouteLinks                                *expressroutelinks.ExpressRouteLinksClient
	ExpressRoutePortAuthorizations                   *expressrouteportauthorizations.ExpressRoutePortAuthorizationsClient
	ExpressRoutePorts                                *expressrouteports.ExpressRoutePortsClient
	ExpressRoutePortsLocations                       *expressrouteportslocations.ExpressRoutePortsLocationsClient
	ExpressRouteProviderPorts                        *expressrouteproviderports.ExpressRouteProviderPortsClient
	ExpressRouteServiceProviders                     *expressrouteserviceproviders.ExpressRouteServiceProvidersClient
	FirewallPolicies                                 *firewallpolicies.FirewallPoliciesClient
	FirewallPolicyRuleCollectionGroups               *firewallpolicyrulecollectiongroups.FirewallPolicyRuleCollectionGroupsClient
	FlowLogs                                         *flowlogs.FlowLogsClient
	IPAllocations                                    *ipallocations.IPAllocationsClient
	IPGroups                                         *ipgroups.IPGroupsClient
	IPamPools                                        *ipampools.IPamPoolsClient
	LoadBalancers                                    *loadbalancers.LoadBalancersClient
	LocalNetworkGateways                             *localnetworkgateways.LocalNetworkGatewaysClient
	NatGateways                                      *natgateways.NatGatewaysClient
	NetworkGroups                                    *networkgroups.NetworkGroupsClient
	NetworkInterfaces                                *networkinterfaces.NetworkInterfacesClient
	NetworkManagerActiveConfigurations               *networkmanageractiveconfigurations.NetworkManagerActiveConfigurationsClient
	NetworkManagerActiveConnectivityConfigurations   *networkmanageractiveconnectivityconfigurations.NetworkManagerActiveConnectivityConfigurationsClient
	NetworkManagerConnections                        *networkmanagerconnections.NetworkManagerConnectionsClient
	NetworkManagerEffectiveConnectivityConfiguration *networkmanagereffectiveconnectivityconfiguration.NetworkManagerEffectiveConnectivityConfigurationClient
	NetworkManagerEffectiveSecurityAdminRules        *networkmanagereffectivesecurityadminrules.NetworkManagerEffectiveSecurityAdminRulesClient
	NetworkManagerRoutingConfigurations              *networkmanagerroutingconfigurations.NetworkManagerRoutingConfigurationsClient
	NetworkManagers                                  *networkmanagers.NetworkManagersClient
	NetworkProfiles                                  *networkprofiles.NetworkProfilesClient
	NetworkSecurityGroups                            *networksecuritygroups.NetworkSecurityGroupsClient
	NetworkSecurityPerimeterAccessRules              *networksecurityperimeteraccessrules.NetworkSecurityPerimeterAccessRulesClient
	NetworkSecurityPerimeterAssociableResourceTypes  *networksecurityperimeterassociableresourcetypes.NetworkSecurityPerimeterAssociableResourceTypesClient
	NetworkSecurityPerimeterAssociations             *networksecurityperimeterassociations.NetworkSecurityPerimeterAssociationsClient
	NetworkSecurityPerimeterLinkReferences           *networksecurityperimeterlinkreferences.NetworkSecurityPerimeterLinkReferencesClient
	NetworkSecurityPerimeterLinks                    *networksecurityperimeterlinks.NetworkSecurityPerimeterLinksClient
	NetworkSecurityPerimeterLoggingConfigurations    *networksecurityperimeterloggingconfigurations.NetworkSecurityPerimeterLoggingConfigurationsClient
	NetworkSecurityPerimeterProfiles                 *networksecurityperimeterprofiles.NetworkSecurityPerimeterProfilesClient
	NetworkSecurityPerimeters                        *networksecurityperimeters.NetworkSecurityPerimetersClient
	NetworkVirtualAppliances                         *networkvirtualappliances.NetworkVirtualAppliancesClient
	NetworkWatchers                                  *networkwatchers.NetworkWatchersClient
	P2sVpnGateways                                   *p2svpngateways.P2sVpnGatewaysClient
	PacketCaptures                                   *packetcaptures.PacketCapturesClient
	PeerExpressRouteCircuitConnections               *peerexpressroutecircuitconnections.PeerExpressRouteCircuitConnectionsClient
	PrivateDnsZoneGroups                             *privatednszonegroups.PrivateDnsZoneGroupsClient
	PrivateEndpoints                                 *privateendpoints.PrivateEndpointsClient
	PrivateLinkServices                              *privatelinkservices.PrivateLinkServicesClient
	PublicIPAddresses                                *publicipaddresses.PublicIPAddressesClient
	PublicIPPrefixes                                 *publicipprefixes.PublicIPPrefixesClient
	ReachabilityAnalysisIntent                       *reachabilityanalysisintent.ReachabilityAnalysisIntentClient
	ReachabilityAnalysisIntents                      *reachabilityanalysisintents.ReachabilityAnalysisIntentsClient
	ReachabilityAnalysisRuns                         *reachabilityanalysisruns.ReachabilityAnalysisRunsClient
	RouteFilterRules                                 *routefilterrules.RouteFilterRulesClient
	RouteFilters                                     *routefilters.RouteFiltersClient
	RouteTables                                      *routetables.RouteTablesClient
	Routes                                           *routes.RoutesClient
	RoutingRuleCollections                           *routingrulecollections.RoutingRuleCollectionsClient
	RoutingRules                                     *routingrules.RoutingRulesClient
	ScopeConnections                                 *scopeconnections.ScopeConnectionsClient
	SecurityAdminConfigurations                      *securityadminconfigurations.SecurityAdminConfigurationsClient
	SecurityPartnerProviders                         *securitypartnerproviders.SecurityPartnerProvidersClient
	SecurityRules                                    *securityrules.SecurityRulesClient
	SecurityUserConfigurations                       *securityuserconfigurations.SecurityUserConfigurationsClient
	SecurityUserRuleCollections                      *securityuserrulecollections.SecurityUserRuleCollectionsClient
	SecurityUserRules                                *securityuserrules.SecurityUserRulesClient
	ServiceEndpointPolicies                          *serviceendpointpolicies.ServiceEndpointPoliciesClient
	ServiceEndpointPolicyDefinitions                 *serviceendpointpolicydefinitions.ServiceEndpointPolicyDefinitionsClient
	ServiceTags                                      *servicetags.ServiceTagsClient
	StaticCidrs                                      *staticcidrs.StaticCidrsClient
	StaticMembers                                    *staticmembers.StaticMembersClient
	Subnets                                          *subnets.SubnetsClient
	TrafficAnalytics                                 *trafficanalytics.TrafficAnalyticsClient
	Usages                                           *usages.UsagesClient
	VMSSPublicIPAddresses                            *vmsspublicipaddresses.VMSSPublicIPAddressesClient
	VerifierWorkspaces                               *verifierworkspaces.VerifierWorkspacesClient
	VipSwap                                          *vipswap.VipSwapClient
	VirtualApplianceSites                            *virtualappliancesites.VirtualApplianceSitesClient
	VirtualApplianceSkus                             *virtualapplianceskus.VirtualApplianceSkusClient
	VirtualNetworkGatewayConnections                 *virtualnetworkgatewayconnections.VirtualNetworkGatewayConnectionsClient
	VirtualNetworkGateways                           *virtualnetworkgateways.VirtualNetworkGatewaysClient
	VirtualNetworkPeerings                           *virtualnetworkpeerings.VirtualNetworkPeeringsClient
	VirtualNetworkTap                                *virtualnetworktap.VirtualNetworkTapClient
	VirtualNetworkTaps                               *virtualnetworktaps.VirtualNetworkTapsClient
	VirtualNetworks                                  *virtualnetworks.VirtualNetworksClient
	VirtualRouterPeerings                            *virtualrouterpeerings.VirtualRouterPeeringsClient
	VirtualRouters                                   *virtualrouters.VirtualRoutersClient
	VirtualWANs                                      *virtualwans.VirtualWANsClient
	VpnGateways                                      *vpngateways.VpnGatewaysClient
	VpnLinkConnections                               *vpnlinkconnections.VpnLinkConnectionsClient
	VpnServerConfigurations                          *vpnserverconfigurations.VpnServerConfigurationsClient
	VpnSites                                         *vpnsites.VpnSitesClient
	WebApplicationFirewallPolicies                   *webapplicationfirewallpolicies.WebApplicationFirewallPoliciesClient
	WebCategories                                    *webcategories.WebCategoriesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	adminRuleCollectionsClient, err := adminrulecollections.NewAdminRuleCollectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdminRuleCollections client: %+v", err)
	}
	configureFunc(adminRuleCollectionsClient.Client)

	adminRulesClient, err := adminrules.NewAdminRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdminRules client: %+v", err)
	}
	configureFunc(adminRulesClient.Client)

	applicationGatewayPrivateEndpointConnectionsClient, err := applicationgatewayprivateendpointconnections.NewApplicationGatewayPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGatewayPrivateEndpointConnections client: %+v", err)
	}
	configureFunc(applicationGatewayPrivateEndpointConnectionsClient.Client)

	applicationGatewayPrivateLinkResourcesClient, err := applicationgatewayprivatelinkresources.NewApplicationGatewayPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGatewayPrivateLinkResources client: %+v", err)
	}
	configureFunc(applicationGatewayPrivateLinkResourcesClient.Client)

	applicationGatewayWafDynamicManifestsClient, err := applicationgatewaywafdynamicmanifests.NewApplicationGatewayWafDynamicManifestsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGatewayWafDynamicManifests client: %+v", err)
	}
	configureFunc(applicationGatewayWafDynamicManifestsClient.Client)

	applicationGatewaysClient, err := applicationgateways.NewApplicationGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGateways client: %+v", err)
	}
	configureFunc(applicationGatewaysClient.Client)

	applicationSecurityGroupsClient, err := applicationsecuritygroups.NewApplicationSecurityGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationSecurityGroups client: %+v", err)
	}
	configureFunc(applicationSecurityGroupsClient.Client)

	availableDelegationsClient, err := availabledelegations.NewAvailableDelegationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableDelegations client: %+v", err)
	}
	configureFunc(availableDelegationsClient.Client)

	availableServiceAliasesClient, err := availableservicealiases.NewAvailableServiceAliasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableServiceAliases client: %+v", err)
	}
	configureFunc(availableServiceAliasesClient.Client)

	azureFirewallsClient, err := azurefirewalls.NewAzureFirewallsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AzureFirewalls client: %+v", err)
	}
	configureFunc(azureFirewallsClient.Client)

	bastionHostsClient, err := bastionhosts.NewBastionHostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BastionHosts client: %+v", err)
	}
	configureFunc(bastionHostsClient.Client)

	bastionShareableLinkClient, err := bastionshareablelink.NewBastionShareableLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BastionShareableLink client: %+v", err)
	}
	configureFunc(bastionShareableLinkClient.Client)

	bgpServiceCommunitiesClient, err := bgpservicecommunities.NewBgpServiceCommunitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BgpServiceCommunities client: %+v", err)
	}
	configureFunc(bgpServiceCommunitiesClient.Client)

	checkDnsAvailabilitiesClient, err := checkdnsavailabilities.NewCheckDnsAvailabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckDnsAvailabilities client: %+v", err)
	}
	configureFunc(checkDnsAvailabilitiesClient.Client)

	cloudServicePublicIPAddressesClient, err := cloudservicepublicipaddresses.NewCloudServicePublicIPAddressesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudServicePublicIPAddresses client: %+v", err)
	}
	configureFunc(cloudServicePublicIPAddressesClient.Client)

	connectionMonitorsClient, err := connectionmonitors.NewConnectionMonitorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectionMonitors client: %+v", err)
	}
	configureFunc(connectionMonitorsClient.Client)

	connectivityConfigurationsClient, err := connectivityconfigurations.NewConnectivityConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectivityConfigurations client: %+v", err)
	}
	configureFunc(connectivityConfigurationsClient.Client)

	customIPPrefixesClient, err := customipprefixes.NewCustomIPPrefixesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomIPPrefixes client: %+v", err)
	}
	configureFunc(customIPPrefixesClient.Client)

	ddosCustomPoliciesClient, err := ddoscustompolicies.NewDdosCustomPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DdosCustomPolicies client: %+v", err)
	}
	configureFunc(ddosCustomPoliciesClient.Client)

	ddosProtectionPlansClient, err := ddosprotectionplans.NewDdosProtectionPlansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DdosProtectionPlans client: %+v", err)
	}
	configureFunc(ddosProtectionPlansClient.Client)

	dscpConfigurationClient, err := dscpconfiguration.NewDscpConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DscpConfiguration client: %+v", err)
	}
	configureFunc(dscpConfigurationClient.Client)

	dscpConfigurationsClient, err := dscpconfigurations.NewDscpConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DscpConfigurations client: %+v", err)
	}
	configureFunc(dscpConfigurationsClient.Client)

	endpointServicesClient, err := endpointservices.NewEndpointServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EndpointServices client: %+v", err)
	}
	configureFunc(endpointServicesClient.Client)

	expressRouteCircuitArpTableClient, err := expressroutecircuitarptable.NewExpressRouteCircuitArpTableClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuitArpTable client: %+v", err)
	}
	configureFunc(expressRouteCircuitArpTableClient.Client)

	expressRouteCircuitAuthorizationsClient, err := expressroutecircuitauthorizations.NewExpressRouteCircuitAuthorizationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuitAuthorizations client: %+v", err)
	}
	configureFunc(expressRouteCircuitAuthorizationsClient.Client)

	expressRouteCircuitConnectionsClient, err := expressroutecircuitconnections.NewExpressRouteCircuitConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuitConnections client: %+v", err)
	}
	configureFunc(expressRouteCircuitConnectionsClient.Client)

	expressRouteCircuitPeeringsClient, err := expressroutecircuitpeerings.NewExpressRouteCircuitPeeringsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuitPeerings client: %+v", err)
	}
	configureFunc(expressRouteCircuitPeeringsClient.Client)

	expressRouteCircuitRoutesTableClient, err := expressroutecircuitroutestable.NewExpressRouteCircuitRoutesTableClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuitRoutesTable client: %+v", err)
	}
	configureFunc(expressRouteCircuitRoutesTableClient.Client)

	expressRouteCircuitRoutesTableSummaryClient, err := expressroutecircuitroutestablesummary.NewExpressRouteCircuitRoutesTableSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuitRoutesTableSummary client: %+v", err)
	}
	configureFunc(expressRouteCircuitRoutesTableSummaryClient.Client)

	expressRouteCircuitStatsClient, err := expressroutecircuitstats.NewExpressRouteCircuitStatsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuitStats client: %+v", err)
	}
	configureFunc(expressRouteCircuitStatsClient.Client)

	expressRouteCircuitsClient, err := expressroutecircuits.NewExpressRouteCircuitsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCircuits client: %+v", err)
	}
	configureFunc(expressRouteCircuitsClient.Client)

	expressRouteConnectionsClient, err := expressrouteconnections.NewExpressRouteConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteConnections client: %+v", err)
	}
	configureFunc(expressRouteConnectionsClient.Client)

	expressRouteCrossConnectionArpTableClient, err := expressroutecrossconnectionarptable.NewExpressRouteCrossConnectionArpTableClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCrossConnectionArpTable client: %+v", err)
	}
	configureFunc(expressRouteCrossConnectionArpTableClient.Client)

	expressRouteCrossConnectionPeeringsClient, err := expressroutecrossconnectionpeerings.NewExpressRouteCrossConnectionPeeringsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCrossConnectionPeerings client: %+v", err)
	}
	configureFunc(expressRouteCrossConnectionPeeringsClient.Client)

	expressRouteCrossConnectionRouteTableClient, err := expressroutecrossconnectionroutetable.NewExpressRouteCrossConnectionRouteTableClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCrossConnectionRouteTable client: %+v", err)
	}
	configureFunc(expressRouteCrossConnectionRouteTableClient.Client)

	expressRouteCrossConnectionRouteTableSummaryClient, err := expressroutecrossconnectionroutetablesummary.NewExpressRouteCrossConnectionRouteTableSummaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCrossConnectionRouteTableSummary client: %+v", err)
	}
	configureFunc(expressRouteCrossConnectionRouteTableSummaryClient.Client)

	expressRouteCrossConnectionsClient, err := expressroutecrossconnections.NewExpressRouteCrossConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteCrossConnections client: %+v", err)
	}
	configureFunc(expressRouteCrossConnectionsClient.Client)

	expressRouteGatewaysClient, err := expressroutegateways.NewExpressRouteGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteGateways client: %+v", err)
	}
	configureFunc(expressRouteGatewaysClient.Client)

	expressRouteLinksClient, err := expressroutelinks.NewExpressRouteLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteLinks client: %+v", err)
	}
	configureFunc(expressRouteLinksClient.Client)

	expressRoutePortAuthorizationsClient, err := expressrouteportauthorizations.NewExpressRoutePortAuthorizationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRoutePortAuthorizations client: %+v", err)
	}
	configureFunc(expressRoutePortAuthorizationsClient.Client)

	expressRoutePortsClient, err := expressrouteports.NewExpressRoutePortsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRoutePorts client: %+v", err)
	}
	configureFunc(expressRoutePortsClient.Client)

	expressRoutePortsLocationsClient, err := expressrouteportslocations.NewExpressRoutePortsLocationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRoutePortsLocations client: %+v", err)
	}
	configureFunc(expressRoutePortsLocationsClient.Client)

	expressRouteProviderPortsClient, err := expressrouteproviderports.NewExpressRouteProviderPortsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteProviderPorts client: %+v", err)
	}
	configureFunc(expressRouteProviderPortsClient.Client)

	expressRouteServiceProvidersClient, err := expressrouteserviceproviders.NewExpressRouteServiceProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExpressRouteServiceProviders client: %+v", err)
	}
	configureFunc(expressRouteServiceProvidersClient.Client)

	firewallPoliciesClient, err := firewallpolicies.NewFirewallPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallPolicies client: %+v", err)
	}
	configureFunc(firewallPoliciesClient.Client)

	firewallPolicyRuleCollectionGroupsClient, err := firewallpolicyrulecollectiongroups.NewFirewallPolicyRuleCollectionGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallPolicyRuleCollectionGroups client: %+v", err)
	}
	configureFunc(firewallPolicyRuleCollectionGroupsClient.Client)

	flowLogsClient, err := flowlogs.NewFlowLogsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FlowLogs client: %+v", err)
	}
	configureFunc(flowLogsClient.Client)

	iPAllocationsClient, err := ipallocations.NewIPAllocationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IPAllocations client: %+v", err)
	}
	configureFunc(iPAllocationsClient.Client)

	iPGroupsClient, err := ipgroups.NewIPGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IPGroups client: %+v", err)
	}
	configureFunc(iPGroupsClient.Client)

	iPamPoolsClient, err := ipampools.NewIPamPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IPamPools client: %+v", err)
	}
	configureFunc(iPamPoolsClient.Client)

	loadBalancersClient, err := loadbalancers.NewLoadBalancersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LoadBalancers client: %+v", err)
	}
	configureFunc(loadBalancersClient.Client)

	localNetworkGatewaysClient, err := localnetworkgateways.NewLocalNetworkGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LocalNetworkGateways client: %+v", err)
	}
	configureFunc(localNetworkGatewaysClient.Client)

	natGatewaysClient, err := natgateways.NewNatGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NatGateways client: %+v", err)
	}
	configureFunc(natGatewaysClient.Client)

	networkGroupsClient, err := networkgroups.NewNetworkGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkGroups client: %+v", err)
	}
	configureFunc(networkGroupsClient.Client)

	networkInterfacesClient, err := networkinterfaces.NewNetworkInterfacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkInterfaces client: %+v", err)
	}
	configureFunc(networkInterfacesClient.Client)

	networkManagerActiveConfigurationsClient, err := networkmanageractiveconfigurations.NewNetworkManagerActiveConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkManagerActiveConfigurations client: %+v", err)
	}
	configureFunc(networkManagerActiveConfigurationsClient.Client)

	networkManagerActiveConnectivityConfigurationsClient, err := networkmanageractiveconnectivityconfigurations.NewNetworkManagerActiveConnectivityConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkManagerActiveConnectivityConfigurations client: %+v", err)
	}
	configureFunc(networkManagerActiveConnectivityConfigurationsClient.Client)

	networkManagerConnectionsClient, err := networkmanagerconnections.NewNetworkManagerConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkManagerConnections client: %+v", err)
	}
	configureFunc(networkManagerConnectionsClient.Client)

	networkManagerEffectiveConnectivityConfigurationClient, err := networkmanagereffectiveconnectivityconfiguration.NewNetworkManagerEffectiveConnectivityConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkManagerEffectiveConnectivityConfiguration client: %+v", err)
	}
	configureFunc(networkManagerEffectiveConnectivityConfigurationClient.Client)

	networkManagerEffectiveSecurityAdminRulesClient, err := networkmanagereffectivesecurityadminrules.NewNetworkManagerEffectiveSecurityAdminRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkManagerEffectiveSecurityAdminRules client: %+v", err)
	}
	configureFunc(networkManagerEffectiveSecurityAdminRulesClient.Client)

	networkManagerRoutingConfigurationsClient, err := networkmanagerroutingconfigurations.NewNetworkManagerRoutingConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkManagerRoutingConfigurations client: %+v", err)
	}
	configureFunc(networkManagerRoutingConfigurationsClient.Client)

	networkManagersClient, err := networkmanagers.NewNetworkManagersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkManagers client: %+v", err)
	}
	configureFunc(networkManagersClient.Client)

	networkProfilesClient, err := networkprofiles.NewNetworkProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkProfiles client: %+v", err)
	}
	configureFunc(networkProfilesClient.Client)

	networkSecurityGroupsClient, err := networksecuritygroups.NewNetworkSecurityGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityGroups client: %+v", err)
	}
	configureFunc(networkSecurityGroupsClient.Client)

	networkSecurityPerimeterAccessRulesClient, err := networksecurityperimeteraccessrules.NewNetworkSecurityPerimeterAccessRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterAccessRules client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterAccessRulesClient.Client)

	networkSecurityPerimeterAssociableResourceTypesClient, err := networksecurityperimeterassociableresourcetypes.NewNetworkSecurityPerimeterAssociableResourceTypesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterAssociableResourceTypes client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterAssociableResourceTypesClient.Client)

	networkSecurityPerimeterAssociationsClient, err := networksecurityperimeterassociations.NewNetworkSecurityPerimeterAssociationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterAssociations client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterAssociationsClient.Client)

	networkSecurityPerimeterLinkReferencesClient, err := networksecurityperimeterlinkreferences.NewNetworkSecurityPerimeterLinkReferencesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterLinkReferences client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterLinkReferencesClient.Client)

	networkSecurityPerimeterLinksClient, err := networksecurityperimeterlinks.NewNetworkSecurityPerimeterLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterLinks client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterLinksClient.Client)

	networkSecurityPerimeterLoggingConfigurationsClient, err := networksecurityperimeterloggingconfigurations.NewNetworkSecurityPerimeterLoggingConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterLoggingConfigurations client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterLoggingConfigurationsClient.Client)

	networkSecurityPerimeterProfilesClient, err := networksecurityperimeterprofiles.NewNetworkSecurityPerimeterProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterProfiles client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterProfilesClient.Client)

	networkSecurityPerimetersClient, err := networksecurityperimeters.NewNetworkSecurityPerimetersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeters client: %+v", err)
	}
	configureFunc(networkSecurityPerimetersClient.Client)

	networkVirtualAppliancesClient, err := networkvirtualappliances.NewNetworkVirtualAppliancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkVirtualAppliances client: %+v", err)
	}
	configureFunc(networkVirtualAppliancesClient.Client)

	networkWatchersClient, err := networkwatchers.NewNetworkWatchersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkWatchers client: %+v", err)
	}
	configureFunc(networkWatchersClient.Client)

	p2sVpnGatewaysClient, err := p2svpngateways.NewP2sVpnGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building P2sVpnGateways client: %+v", err)
	}
	configureFunc(p2sVpnGatewaysClient.Client)

	packetCapturesClient, err := packetcaptures.NewPacketCapturesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PacketCaptures client: %+v", err)
	}
	configureFunc(packetCapturesClient.Client)

	peerExpressRouteCircuitConnectionsClient, err := peerexpressroutecircuitconnections.NewPeerExpressRouteCircuitConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PeerExpressRouteCircuitConnections client: %+v", err)
	}
	configureFunc(peerExpressRouteCircuitConnectionsClient.Client)

	privateDnsZoneGroupsClient, err := privatednszonegroups.NewPrivateDnsZoneGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateDnsZoneGroups client: %+v", err)
	}
	configureFunc(privateDnsZoneGroupsClient.Client)

	privateEndpointsClient, err := privateendpoints.NewPrivateEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpoints client: %+v", err)
	}
	configureFunc(privateEndpointsClient.Client)

	privateLinkServicesClient, err := privatelinkservices.NewPrivateLinkServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkServices client: %+v", err)
	}
	configureFunc(privateLinkServicesClient.Client)

	publicIPAddressesClient, err := publicipaddresses.NewPublicIPAddressesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicIPAddresses client: %+v", err)
	}
	configureFunc(publicIPAddressesClient.Client)

	publicIPPrefixesClient, err := publicipprefixes.NewPublicIPPrefixesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicIPPrefixes client: %+v", err)
	}
	configureFunc(publicIPPrefixesClient.Client)

	reachabilityAnalysisIntentClient, err := reachabilityanalysisintent.NewReachabilityAnalysisIntentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReachabilityAnalysisIntent client: %+v", err)
	}
	configureFunc(reachabilityAnalysisIntentClient.Client)

	reachabilityAnalysisIntentsClient, err := reachabilityanalysisintents.NewReachabilityAnalysisIntentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReachabilityAnalysisIntents client: %+v", err)
	}
	configureFunc(reachabilityAnalysisIntentsClient.Client)

	reachabilityAnalysisRunsClient, err := reachabilityanalysisruns.NewReachabilityAnalysisRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReachabilityAnalysisRuns client: %+v", err)
	}
	configureFunc(reachabilityAnalysisRunsClient.Client)

	routeFilterRulesClient, err := routefilterrules.NewRouteFilterRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RouteFilterRules client: %+v", err)
	}
	configureFunc(routeFilterRulesClient.Client)

	routeFiltersClient, err := routefilters.NewRouteFiltersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RouteFilters client: %+v", err)
	}
	configureFunc(routeFiltersClient.Client)

	routeTablesClient, err := routetables.NewRouteTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RouteTables client: %+v", err)
	}
	configureFunc(routeTablesClient.Client)

	routesClient, err := routes.NewRoutesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Routes client: %+v", err)
	}
	configureFunc(routesClient.Client)

	routingRuleCollectionsClient, err := routingrulecollections.NewRoutingRuleCollectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoutingRuleCollections client: %+v", err)
	}
	configureFunc(routingRuleCollectionsClient.Client)

	routingRulesClient, err := routingrules.NewRoutingRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoutingRules client: %+v", err)
	}
	configureFunc(routingRulesClient.Client)

	scopeConnectionsClient, err := scopeconnections.NewScopeConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScopeConnections client: %+v", err)
	}
	configureFunc(scopeConnectionsClient.Client)

	securityAdminConfigurationsClient, err := securityadminconfigurations.NewSecurityAdminConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityAdminConfigurations client: %+v", err)
	}
	configureFunc(securityAdminConfigurationsClient.Client)

	securityPartnerProvidersClient, err := securitypartnerproviders.NewSecurityPartnerProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityPartnerProviders client: %+v", err)
	}
	configureFunc(securityPartnerProvidersClient.Client)

	securityRulesClient, err := securityrules.NewSecurityRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityRules client: %+v", err)
	}
	configureFunc(securityRulesClient.Client)

	securityUserConfigurationsClient, err := securityuserconfigurations.NewSecurityUserConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityUserConfigurations client: %+v", err)
	}
	configureFunc(securityUserConfigurationsClient.Client)

	securityUserRuleCollectionsClient, err := securityuserrulecollections.NewSecurityUserRuleCollectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityUserRuleCollections client: %+v", err)
	}
	configureFunc(securityUserRuleCollectionsClient.Client)

	securityUserRulesClient, err := securityuserrules.NewSecurityUserRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityUserRules client: %+v", err)
	}
	configureFunc(securityUserRulesClient.Client)

	serviceEndpointPoliciesClient, err := serviceendpointpolicies.NewServiceEndpointPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceEndpointPolicies client: %+v", err)
	}
	configureFunc(serviceEndpointPoliciesClient.Client)

	serviceEndpointPolicyDefinitionsClient, err := serviceendpointpolicydefinitions.NewServiceEndpointPolicyDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceEndpointPolicyDefinitions client: %+v", err)
	}
	configureFunc(serviceEndpointPolicyDefinitionsClient.Client)

	serviceTagsClient, err := servicetags.NewServiceTagsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceTags client: %+v", err)
	}
	configureFunc(serviceTagsClient.Client)

	staticCidrsClient, err := staticcidrs.NewStaticCidrsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticCidrs client: %+v", err)
	}
	configureFunc(staticCidrsClient.Client)

	staticMembersClient, err := staticmembers.NewStaticMembersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticMembers client: %+v", err)
	}
	configureFunc(staticMembersClient.Client)

	subnetsClient, err := subnets.NewSubnetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subnets client: %+v", err)
	}
	configureFunc(subnetsClient.Client)

	trafficAnalyticsClient, err := trafficanalytics.NewTrafficAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TrafficAnalytics client: %+v", err)
	}
	configureFunc(trafficAnalyticsClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	vMSSPublicIPAddressesClient, err := vmsspublicipaddresses.NewVMSSPublicIPAddressesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMSSPublicIPAddresses client: %+v", err)
	}
	configureFunc(vMSSPublicIPAddressesClient.Client)

	verifierWorkspacesClient, err := verifierworkspaces.NewVerifierWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VerifierWorkspaces client: %+v", err)
	}
	configureFunc(verifierWorkspacesClient.Client)

	vipSwapClient, err := vipswap.NewVipSwapClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VipSwap client: %+v", err)
	}
	configureFunc(vipSwapClient.Client)

	virtualApplianceSitesClient, err := virtualappliancesites.NewVirtualApplianceSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualApplianceSites client: %+v", err)
	}
	configureFunc(virtualApplianceSitesClient.Client)

	virtualApplianceSkusClient, err := virtualapplianceskus.NewVirtualApplianceSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualApplianceSkus client: %+v", err)
	}
	configureFunc(virtualApplianceSkusClient.Client)

	virtualNetworkGatewayConnectionsClient, err := virtualnetworkgatewayconnections.NewVirtualNetworkGatewayConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkGatewayConnections client: %+v", err)
	}
	configureFunc(virtualNetworkGatewayConnectionsClient.Client)

	virtualNetworkGatewaysClient, err := virtualnetworkgateways.NewVirtualNetworkGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkGateways client: %+v", err)
	}
	configureFunc(virtualNetworkGatewaysClient.Client)

	virtualNetworkPeeringsClient, err := virtualnetworkpeerings.NewVirtualNetworkPeeringsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkPeerings client: %+v", err)
	}
	configureFunc(virtualNetworkPeeringsClient.Client)

	virtualNetworkTapClient, err := virtualnetworktap.NewVirtualNetworkTapClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkTap client: %+v", err)
	}
	configureFunc(virtualNetworkTapClient.Client)

	virtualNetworkTapsClient, err := virtualnetworktaps.NewVirtualNetworkTapsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkTaps client: %+v", err)
	}
	configureFunc(virtualNetworkTapsClient.Client)

	virtualNetworksClient, err := virtualnetworks.NewVirtualNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworks client: %+v", err)
	}
	configureFunc(virtualNetworksClient.Client)

	virtualRouterPeeringsClient, err := virtualrouterpeerings.NewVirtualRouterPeeringsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualRouterPeerings client: %+v", err)
	}
	configureFunc(virtualRouterPeeringsClient.Client)

	virtualRoutersClient, err := virtualrouters.NewVirtualRoutersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualRouters client: %+v", err)
	}
	configureFunc(virtualRoutersClient.Client)

	virtualWANsClient, err := virtualwans.NewVirtualWANsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualWANs client: %+v", err)
	}
	configureFunc(virtualWANsClient.Client)

	vpnGatewaysClient, err := vpngateways.NewVpnGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VpnGateways client: %+v", err)
	}
	configureFunc(vpnGatewaysClient.Client)

	vpnLinkConnectionsClient, err := vpnlinkconnections.NewVpnLinkConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VpnLinkConnections client: %+v", err)
	}
	configureFunc(vpnLinkConnectionsClient.Client)

	vpnServerConfigurationsClient, err := vpnserverconfigurations.NewVpnServerConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VpnServerConfigurations client: %+v", err)
	}
	configureFunc(vpnServerConfigurationsClient.Client)

	vpnSitesClient, err := vpnsites.NewVpnSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VpnSites client: %+v", err)
	}
	configureFunc(vpnSitesClient.Client)

	webApplicationFirewallPoliciesClient, err := webapplicationfirewallpolicies.NewWebApplicationFirewallPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebApplicationFirewallPolicies client: %+v", err)
	}
	configureFunc(webApplicationFirewallPoliciesClient.Client)

	webCategoriesClient, err := webcategories.NewWebCategoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebCategories client: %+v", err)
	}
	configureFunc(webCategoriesClient.Client)

	return &Client{
		AdminRuleCollections: adminRuleCollectionsClient,
		AdminRules:           adminRulesClient,
		ApplicationGatewayPrivateEndpointConnections:     applicationGatewayPrivateEndpointConnectionsClient,
		ApplicationGatewayPrivateLinkResources:           applicationGatewayPrivateLinkResourcesClient,
		ApplicationGatewayWafDynamicManifests:            applicationGatewayWafDynamicManifestsClient,
		ApplicationGateways:                              applicationGatewaysClient,
		ApplicationSecurityGroups:                        applicationSecurityGroupsClient,
		AvailableDelegations:                             availableDelegationsClient,
		AvailableServiceAliases:                          availableServiceAliasesClient,
		AzureFirewalls:                                   azureFirewallsClient,
		BastionHosts:                                     bastionHostsClient,
		BastionShareableLink:                             bastionShareableLinkClient,
		BgpServiceCommunities:                            bgpServiceCommunitiesClient,
		CheckDnsAvailabilities:                           checkDnsAvailabilitiesClient,
		CloudServicePublicIPAddresses:                    cloudServicePublicIPAddressesClient,
		ConnectionMonitors:                               connectionMonitorsClient,
		ConnectivityConfigurations:                       connectivityConfigurationsClient,
		CustomIPPrefixes:                                 customIPPrefixesClient,
		DdosCustomPolicies:                               ddosCustomPoliciesClient,
		DdosProtectionPlans:                              ddosProtectionPlansClient,
		DscpConfiguration:                                dscpConfigurationClient,
		DscpConfigurations:                               dscpConfigurationsClient,
		EndpointServices:                                 endpointServicesClient,
		ExpressRouteCircuitArpTable:                      expressRouteCircuitArpTableClient,
		ExpressRouteCircuitAuthorizations:                expressRouteCircuitAuthorizationsClient,
		ExpressRouteCircuitConnections:                   expressRouteCircuitConnectionsClient,
		ExpressRouteCircuitPeerings:                      expressRouteCircuitPeeringsClient,
		ExpressRouteCircuitRoutesTable:                   expressRouteCircuitRoutesTableClient,
		ExpressRouteCircuitRoutesTableSummary:            expressRouteCircuitRoutesTableSummaryClient,
		ExpressRouteCircuitStats:                         expressRouteCircuitStatsClient,
		ExpressRouteCircuits:                             expressRouteCircuitsClient,
		ExpressRouteConnections:                          expressRouteConnectionsClient,
		ExpressRouteCrossConnectionArpTable:              expressRouteCrossConnectionArpTableClient,
		ExpressRouteCrossConnectionPeerings:              expressRouteCrossConnectionPeeringsClient,
		ExpressRouteCrossConnectionRouteTable:            expressRouteCrossConnectionRouteTableClient,
		ExpressRouteCrossConnectionRouteTableSummary:     expressRouteCrossConnectionRouteTableSummaryClient,
		ExpressRouteCrossConnections:                     expressRouteCrossConnectionsClient,
		ExpressRouteGateways:                             expressRouteGatewaysClient,
		ExpressRouteLinks:                                expressRouteLinksClient,
		ExpressRoutePortAuthorizations:                   expressRoutePortAuthorizationsClient,
		ExpressRoutePorts:                                expressRoutePortsClient,
		ExpressRoutePortsLocations:                       expressRoutePortsLocationsClient,
		ExpressRouteProviderPorts:                        expressRouteProviderPortsClient,
		ExpressRouteServiceProviders:                     expressRouteServiceProvidersClient,
		FirewallPolicies:                                 firewallPoliciesClient,
		FirewallPolicyRuleCollectionGroups:               firewallPolicyRuleCollectionGroupsClient,
		FlowLogs:                                         flowLogsClient,
		IPAllocations:                                    iPAllocationsClient,
		IPGroups:                                         iPGroupsClient,
		IPamPools:                                        iPamPoolsClient,
		LoadBalancers:                                    loadBalancersClient,
		LocalNetworkGateways:                             localNetworkGatewaysClient,
		NatGateways:                                      natGatewaysClient,
		NetworkGroups:                                    networkGroupsClient,
		NetworkInterfaces:                                networkInterfacesClient,
		NetworkManagerActiveConfigurations:               networkManagerActiveConfigurationsClient,
		NetworkManagerActiveConnectivityConfigurations:   networkManagerActiveConnectivityConfigurationsClient,
		NetworkManagerConnections:                        networkManagerConnectionsClient,
		NetworkManagerEffectiveConnectivityConfiguration: networkManagerEffectiveConnectivityConfigurationClient,
		NetworkManagerEffectiveSecurityAdminRules:        networkManagerEffectiveSecurityAdminRulesClient,
		NetworkManagerRoutingConfigurations:              networkManagerRoutingConfigurationsClient,
		NetworkManagers:                                  networkManagersClient,
		NetworkProfiles:                                  networkProfilesClient,
		NetworkSecurityGroups:                            networkSecurityGroupsClient,
		NetworkSecurityPerimeterAccessRules:              networkSecurityPerimeterAccessRulesClient,
		NetworkSecurityPerimeterAssociableResourceTypes:  networkSecurityPerimeterAssociableResourceTypesClient,
		NetworkSecurityPerimeterAssociations:             networkSecurityPerimeterAssociationsClient,
		NetworkSecurityPerimeterLinkReferences:           networkSecurityPerimeterLinkReferencesClient,
		NetworkSecurityPerimeterLinks:                    networkSecurityPerimeterLinksClient,
		NetworkSecurityPerimeterLoggingConfigurations:    networkSecurityPerimeterLoggingConfigurationsClient,
		NetworkSecurityPerimeterProfiles:                 networkSecurityPerimeterProfilesClient,
		NetworkSecurityPerimeters:                        networkSecurityPerimetersClient,
		NetworkVirtualAppliances:                         networkVirtualAppliancesClient,
		NetworkWatchers:                                  networkWatchersClient,
		P2sVpnGateways:                                   p2sVpnGatewaysClient,
		PacketCaptures:                                   packetCapturesClient,
		PeerExpressRouteCircuitConnections:               peerExpressRouteCircuitConnectionsClient,
		PrivateDnsZoneGroups:                             privateDnsZoneGroupsClient,
		PrivateEndpoints:                                 privateEndpointsClient,
		PrivateLinkServices:                              privateLinkServicesClient,
		PublicIPAddresses:                                publicIPAddressesClient,
		PublicIPPrefixes:                                 publicIPPrefixesClient,
		ReachabilityAnalysisIntent:                       reachabilityAnalysisIntentClient,
		ReachabilityAnalysisIntents:                      reachabilityAnalysisIntentsClient,
		ReachabilityAnalysisRuns:                         reachabilityAnalysisRunsClient,
		RouteFilterRules:                                 routeFilterRulesClient,
		RouteFilters:                                     routeFiltersClient,
		RouteTables:                                      routeTablesClient,
		Routes:                                           routesClient,
		RoutingRuleCollections:                           routingRuleCollectionsClient,
		RoutingRules:                                     routingRulesClient,
		ScopeConnections:                                 scopeConnectionsClient,
		SecurityAdminConfigurations:                      securityAdminConfigurationsClient,
		SecurityPartnerProviders:                         securityPartnerProvidersClient,
		SecurityRules:                                    securityRulesClient,
		SecurityUserConfigurations:                       securityUserConfigurationsClient,
		SecurityUserRuleCollections:                      securityUserRuleCollectionsClient,
		SecurityUserRules:                                securityUserRulesClient,
		ServiceEndpointPolicies:                          serviceEndpointPoliciesClient,
		ServiceEndpointPolicyDefinitions:                 serviceEndpointPolicyDefinitionsClient,
		ServiceTags:                                      serviceTagsClient,
		StaticCidrs:                                      staticCidrsClient,
		StaticMembers:                                    staticMembersClient,
		Subnets:                                          subnetsClient,
		TrafficAnalytics:                                 trafficAnalyticsClient,
		Usages:                                           usagesClient,
		VMSSPublicIPAddresses:                            vMSSPublicIPAddressesClient,
		VerifierWorkspaces:                               verifierWorkspacesClient,
		VipSwap:                                          vipSwapClient,
		VirtualApplianceSites:                            virtualApplianceSitesClient,
		VirtualApplianceSkus:                             virtualApplianceSkusClient,
		VirtualNetworkGatewayConnections:                 virtualNetworkGatewayConnectionsClient,
		VirtualNetworkGateways:                           virtualNetworkGatewaysClient,
		VirtualNetworkPeerings:                           virtualNetworkPeeringsClient,
		VirtualNetworkTap:                                virtualNetworkTapClient,
		VirtualNetworkTaps:                               virtualNetworkTapsClient,
		VirtualNetworks:                                  virtualNetworksClient,
		VirtualRouterPeerings:                            virtualRouterPeeringsClient,
		VirtualRouters:                                   virtualRoutersClient,
		VirtualWANs:                                      virtualWANsClient,
		VpnGateways:                                      vpnGatewaysClient,
		VpnLinkConnections:                               vpnLinkConnectionsClient,
		VpnServerConfigurations:                          vpnServerConfigurationsClient,
		VpnSites:                                         vpnSitesClient,
		WebApplicationFirewallPolicies:                   webApplicationFirewallPoliciesClient,
		WebCategories:                                    webCategoriesClient,
	}, nil
}
