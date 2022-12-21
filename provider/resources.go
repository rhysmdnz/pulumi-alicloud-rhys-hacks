// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alicloud

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/rhysmdnz/pulumi-alicloud/provider/pkg/version"
)

// all of the token components used below.
const (
	//mainPkg = "alicloud"
	//mainMod = "index" // the alicloud module
	// packages:
	alicloudPkg = "alicloud"
	// modules:
	alicloudMod            = "index"
	actionTrailMod         = "ActionTrail"
	adbMod                 = "Adb"
	albMod                 = "Alb"
	aliKafaMod             = "AliKafka"
	amqpMod                = "Amqp"
	apiGatewayMod          = "ApiGateway"
	armsMod                = "Arms"
	bastionHostMod         = "BastionHost"
	brainMod               = "Brain"
	casMod                 = "Cas"
	cassandraMod           = "Cassandra"
	cddcMod                = "Cddc"
	cdnMod                 = "Cdn"
	cenMod                 = "Cen"
	clickHouseMod          = "ClickHouse"
	cloudAuthMod           = "CloudAuth"
	cloudConnectMod        = "CloudConnect"
	cloudFirewallMod       = "CloudFirewall"
	cloudStorageGatewayMod = "CloudStorageGateway"
	cloudSsoMod            = "CloudSso"
	cmsMod                 = "Cms"
	cfgMod                 = "Cfg"
	crMod                  = "CR"
	csMod                  = "CS"
	datahubMod             = "Datahub"
	dataWorksMod           = "DataWorks"
	databaseFilesystemMod  = "DatabaseFilesystem"
	databaseGatewayMod     = "DatabaseGateway"
	dcdnMod                = "Dcdn"
	ddsMod                 = "Dds"
	ddosMod                = "Ddos"
	dfsMod                 = "Dfs"
	directMailMod          = "DirectMail"
	dmsMod                 = "Dms"
	dnsMod                 = "Dns"
	drdsMod                = "Drds"
	dtsMod                 = "Dts"
	eaisMod                = "Eais"
	eciMod                 = "Eci"
	ecpMod                 = "Ecp"
	ecsMod                 = "Ecs"
	edasMod                = "Edas"
	edsMod                 = "Eds"
	ehpcMod                = "Ehpc"
	eipAnyCastMod          = "EipAnycast"
	elasticsearchMod       = "ElasticSearch"
	emrMod                 = "Emr"
	ensMod                 = "Ens"
	essMod                 = "Ess"
	eventBridgeMod         = "EventBridge"
	expressConnect         = "ExpressConnect"
	fcMod                  = "FC"
	fnfMod                 = "FNF"
	gaMod                  = "Ga"
	gpdbMod                = "Gpdb"
	graphDatabaseMod       = "GraphDatabase"
	hbaseMod               = "Hbase"
	hbrMod                 = "Hbr"
	iotMod                 = "Iot"
	immMod                 = "Imm"
	impMod                 = "Imp"
	kmsMod                 = "Kms"
	kvstoreMod             = "KVStore"
	lindormMod             = "Lindorm"
	logMod                 = "Log"
	marketPlaceMod         = "MarketPlace"
	maxComputeMod          = "MaxCompute"
	mhubMod                = "Mhub"
	mongoDbMod             = "MongoDB"
	mnsMod                 = "Mns"
	mseMod                 = "Mse"
	nasMod                 = "Nas"
	oosMod                 = "Oos"
	openSearchMod          = "OpenSearch"
	ossMod                 = "Oss"
	otsMod                 = "Ots"
	polarDbMod             = "PolarDB"
	privateLinkMod         = "PrivateLink"
	pvtzMod                = "Pvtz"
	quickBiMod             = "QuickBI"
	quotasMod              = "Quotas"
	ramMod                 = "Ram"
	rdcMod                 = "Rdc"
	resourceManagerMod     = "ResourceManager"
	rocketMqMod            = "RocketMQ"
	rosMod                 = "Ros"
	rdsMod                 = "Rds"
	saeMod                 = "Sae"
	sagMod                 = "Sag"
	serviceMeshMod         = "ServiceMesh"
	sasMod                 = "SimpleApplicationServer"
	scdnMod                = "Scdn"
	sddpMod                = "Sddp"
	securityCenterMod      = "SecurityCenter"
	schedulerXMod          = "SchedulerX"
	slbMod                 = "Slb"
	smsMod                 = "Sms"
	tagMod                 = "Tag"
	tsdbMod                = "Tsdb"
	vpcMod                 = "Vpc"
	vsMod                  = "VideoSurveillance"
	vodMod                 = "Vod"
	vpnMod                 = "Vpn"
	wafMod                 = "Waf"
	yundunMod              = "Yundun"
)

var namespaceMap = map[string]string{
	"alicloud": "AliCloud",
}

// alicloudMember manufactures a type token for the AliCloud package and the given module, file name, and type.
func alicloudMember(moduleTitle string, fn string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	if fn != "" {
		moduleName += "/" + fn
	}
	return tokens.ModuleMember(alicloudPkg + ":" + moduleName + ":" + mem)
}

// alicloudType manufactures a type token for the AWS package and the given module, file name, and type.
func alicloudType(mod string, fn string, typ string) tokens.Type {
	return tokens.Type(alicloudMember(mod, fn, typ))
}

// alicloudTypeDefaultFile manufactures a standard resource token given a module and resource name.  It automatically
// uses the package and names the file by simply lower casing the type's first character.
func alicloudTypeDefaultFile(mod string, typ string) tokens.Type {
	fn := string(unicode.ToLower(rune(typ[0]))) + typ[1:]
	return alicloudType(mod, fn, typ)
}

// dataSource manufactures a standard resource token given a module and resource name.  It automatically uses the AWS
// package and names the file by simply lower casing the data source's first character.
func dataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return alicloudMember(mod, fn, res)
}

// aliResource manufactures a standard resource token given a module and resource name.
func aliResource(mod string, res string) tokens.Type {
	return alicloudTypeDefaultFile(mod, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(alicloud.Provider().(*schema.Provider))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "alicloud",
		DisplayName: "Alibaba Cloud",
		Publisher:   "Rhys",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/rhysmdnz/pulumi-alicloud",
		Description: "A Pulumi package for creating and managing AliCloud resources.",
		Keywords:    []string{"pulumi", "alicloud", "category/cloud"},
		License:     "Apache-2.0",
		Homepage:    "https://www.pulumi.com",
		Repository:  "https://github.com/rhysmdnz/pulumi-alicloud",
		GitHubOrg:   "aliyun",
		Config: map[string]*tfbridge.SchemaInfo{
			"ecs_role_name": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ALICLOUD_ECS_ROLE_NAME"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ALICLOUD_REGION"},
				},
			},
			"profile": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ALICLOUD_PROFILE"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Mainmod
			"alicloud_msc_sub_contact":      {Tok: aliResource(alicloudMod, "MscSubContract")},
			"alicloud_msc_sub_subscription": {Tok: aliResource(alicloudMod, "MscSubSubscription")},
			"alicloud_msc_sub_webhook":      {Tok: aliResource(alicloudMod, "MscSubWebhook")},

			// ActionTrail
			"alicloud_actiontrail_trail":                {Tok: aliResource(actionTrailMod, "Trail")},
			"alicloud_actiontrail_history_delivery_job": {Tok: aliResource(actionTrailMod, "HistoryDeliveryJob")},
			"alicloud_actiontrail": {
				Tok:                aliResource(actionTrailMod, "TrailDeprecated"),
				DeprecationMessage: "Resource renamed to `Trail`",
			},

			// Adb
			"alicloud_adb_account":       {Tok: aliResource(adbMod, "Account")},
			"alicloud_adb_backup_policy": {Tok: aliResource(adbMod, "BackupPolicy")},
			"alicloud_adb_cluster":       {Tok: aliResource(adbMod, "Cluster")},
			"alicloud_adb_connection":    {Tok: aliResource(adbMod, "Connection")},
			"alicloud_adb_db_cluster":    {Tok: aliResource(adbMod, "DBCluster")},

			// Alb
			"alicloud_alb_security_policy": {Tok: aliResource(albMod, "SecurityPolicy")},
			"alicloud_alb_server_group":    {Tok: aliResource(albMod, "ServerGroup")},
			"alicloud_alb_load_balancer":   {Tok: aliResource(albMod, "LoadBalancer")},
			"alicloud_alb_rule":            {Tok: aliResource(albMod, "Rule")},
			"alicloud_alb_acl":             {Tok: aliResource(albMod, "Acl")},
			"alicloud_alb_listener": {
				Tok: aliResource(albMod, "Listener"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// This property was deprecated in favor of x_forwarded_for_config, and continuing to generate a
					// binding causes the .NET SDK to not compile, so we skip:
					"xforwarded_for_config": {
						Omit: true,
					},
				},
			},
			"alicloud_alb_health_check_template":                      {Tok: aliResource(albMod, "HealthCheckTemplate")},
			"alicloud_alb_listener_additional_certificate_attachment": {Tok: aliResource(albMod, "ListenerAdditionalCertificateAttachment")},
			"alicloud_alb_listener_acl_attachment":                    {Tok: aliResource(albMod, "ListenerAclAttachment")},
			"alicloud_alb_acl_entry_attachment":                       {Tok: aliResource(albMod, "AclEntryAttachment")},

			// AliKafka
			"alicloud_alikafka_consumer_group":                 {Tok: aliResource(aliKafaMod, "ConsumerGroup")},
			"alicloud_alikafka_instance":                       {Tok: aliResource(aliKafaMod, "Instance")},
			"alicloud_alikafka_sasl_acl":                       {Tok: aliResource(aliKafaMod, "SaslAcl")},
			"alicloud_alikafka_sasl_user":                      {Tok: aliResource(aliKafaMod, "SaslUser")},
			"alicloud_alikafka_instance_allowed_ip_attachment": {Tok: aliResource(aliKafaMod, "InstanceAllowedIpAttachment")},
			"alicloud_alikafka_topic": {
				Tok: aliResource(aliKafaMod, "Topic"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"topic": {
						CSharpName: "TopicName",
					},
				},
			},

			// amqp
			"alicloud_amqp_virtual_host": {Tok: aliResource(amqpMod, "VirtualHost")},
			"alicloud_amqp_queue":        {Tok: aliResource(amqpMod, "Queue")},
			"alicloud_amqp_instance":     {Tok: aliResource(amqpMod, "Instance")},
			"alicloud_amqp_exchange":     {Tok: aliResource(amqpMod, "Exchange")},
			"alicloud_amqp_binding":      {Tok: aliResource(amqpMod, "Binding")},

			// Api Gateway
			"alicloud_api_gateway_group":          {Tok: aliResource(apiGatewayMod, "Group")},
			"alicloud_api_gateway_api":            {Tok: aliResource(apiGatewayMod, "Api")},
			"alicloud_api_gateway_app":            {Tok: aliResource(apiGatewayMod, "App")},
			"alicloud_api_gateway_app_attachment": {Tok: aliResource(apiGatewayMod, "AppAttachment")},
			"alicloud_api_gateway_vpc_access":     {Tok: aliResource(apiGatewayMod, "VpcAccess")},
			"alicloud_api_gateway_backend":        {Tok: aliResource(apiGatewayMod, "Backend")},

			// Arms
			"alicloud_arms_alert_contact":         {Tok: aliResource(armsMod, "AlertContact")},
			"alicloud_arms_alert_contact_group":   {Tok: aliResource(armsMod, "AlertContactGroup")},
			"alicloud_arms_dispatch_rule":         {Tok: aliResource(armsMod, "DispatchRule")},
			"alicloud_arms_prometheus_alert_rule": {Tok: aliResource(armsMod, "PrometheusAlertRule")},

			// BastionHost
			"alicloud_bastionhost_user_group":                               {Tok: aliResource(bastionHostMod, "UserGroup")},
			"alicloud_bastionhost_instance":                                 {Tok: aliResource(bastionHostMod, "Instance")},
			"alicloud_bastionhost_user":                                     {Tok: aliResource(bastionHostMod, "User")},
			"alicloud_bastionhost_user_attachment":                          {Tok: aliResource(bastionHostMod, "UserAttachment")},
			"alicloud_bastionhost_host_group":                               {Tok: aliResource(bastionHostMod, "HostGroup")},
			"alicloud_bastionhost_host_group_account_user_group_attachment": {Tok: aliResource(bastionHostMod, "HostGroupAccountUserGroupAttachment")},
			"alicloud_bastionhost_host_group_account_user_attachment":       {Tok: aliResource(bastionHostMod, "HostGroupAccountUserAttachment")},
			"alicloud_bastionhost_host_account_user_attachment":             {Tok: aliResource(bastionHostMod, "HostAccountUserAttachment")},
			"alicloud_bastionhost_host_account_user_group_attachment":       {Tok: aliResource(bastionHostMod, "HostAccountUserGroupAttachment")},
			"alicloud_bastionhost_host_attachment":                          {Tok: aliResource(bastionHostMod, "HostAttachment")},
			"alicloud_bastionhost_host_account":                             {Tok: aliResource(bastionHostMod, "HostAccount")},
			"alicloud_bastionhost_host":                                     {Tok: aliResource(bastionHostMod, "Host")},
			"alicloud_bastionhost_host_account_share_key_attachment":        {Tok: aliResource(bastionHostMod, "HostAccountShareKeyAttachment")},
			"alicloud_bastionhost_host_share_key":                           {Tok: aliResource(bastionHostMod, "HostShareKey")},

			// Brain
			"alicloud_brain_industrial_pid_project":      {Tok: aliResource(brainMod, "IndustrialPidProject")},
			"alicloud_brain_industrial_pid_organization": {Tok: aliResource(brainMod, "IndustrialPidOrganization")},
			"alicloud_brain_industrial_pid_loop":         {Tok: aliResource(brainMod, "IndustrialPidLoop")},

			// Cas
			"alicloud_cas_certificate": {
				Tok:                aliResource(casMod, "Certificate"),
				DeprecationMessage: "This resource has been deprecated in favour of ServiceCertificate",
			},
			"alicloud_ssl_certificates_service_certificate": {Tok: aliResource(casMod, "ServiceCertificate")},

			// Cassandra
			"alicloud_cassandra_cluster":     {Tok: aliResource(cassandraMod, "Cluster")},
			"alicloud_cassandra_data_center": {Tok: aliResource(cassandraMod, "DataCenter")},
			"alicloud_cassandra_backup_plan": {Tok: aliResource(cassandraMod, "BackupPlan")},

			// CDDC
			"alicloud_cddc_dedicated_host_group":   {Tok: aliResource(cddcMod, "DedicatedHostGroup")},
			"alicloud_cddc_dedicated_host":         {Tok: aliResource(cddcMod, "DedicatedHost")},
			"alicloud_cddc_dedicated_host_account": {Tok: aliResource(cddcMod, "DedicatedHostAccount")},

			//CDN
			"alicloud_cdn_domain": {
				Tok: aliResource(cdnMod, "Domain"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_cdn_domain_config":          {Tok: aliResource(cdnMod, "DomainConfig")},
			"alicloud_cdn_domain_new":             {Tok: aliResource(cdnMod, "DomainNew")},
			"alicloud_cdn_real_time_log_delivery": {Tok: aliResource(cdnMod, "RealTimeLogDelivery")},
			"alicloud_cdn_fc_trigger":             {Tok: aliResource(cdnMod, "FcTrigger")},

			// CEN
			"alicloud_cen_bandwidth_limit": {
				Tok: aliResource(cenMod, "BandwidthLimit"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"bandwidth_limit": {
						CSharpName: "Limit",
					},
				},
			},
			"alicloud_cen_bandwidth_package":            {Tok: aliResource(cenMod, "BandwidthPackage")},
			"alicloud_cen_bandwidth_package_attachment": {Tok: aliResource(cenMod, "BandwidthPackageAttachment")},
			"alicloud_cen_instance":                     {Tok: aliResource(cenMod, "Instance")},
			"alicloud_cen_instance_attachment":          {Tok: aliResource(cenMod, "InstanceAttachment")},
			"alicloud_cen_route_entry":                  {Tok: aliResource(cenMod, "RouteEntry")},
			"alicloud_cen_instance_grant":               {Tok: aliResource(cenMod, "InstanceGrant")},
			"alicloud_cen_flowlog":                      {Tok: aliResource(cenMod, "FlowLog")},
			"alicloud_cen_route_map":                    {Tok: aliResource(cenMod, "RouteMap")},
			"alicloud_cen_private_zone":                 {Tok: aliResource(cenMod, "PrivateZone")},
			"alicloud_cen_vbr_health_check":             {Tok: aliResource(cenMod, "VbrHealthCheck")},
			"alicloud_cen_route_service":                {Tok: aliResource(cenMod, "RouteService")},
			"alicloud_cen_transit_router":               {Tok: aliResource(cenMod, "TransitRouter")},
			"alicloud_cen_transit_router_route_table":   {Tok: aliResource(cenMod, "TransitRouterRouteTable")},
			"alicloud_cen_transit_router_route_table_association": {
				Tok: aliResource(cenMod, "TransitRouterRouteTableAssociation"),
			},
			"alicloud_cen_transit_router_route_table_propagation": {
				Tok: aliResource(cenMod, "TransitRouterRouteTablePropagation"),
			},
			"alicloud_cen_transit_router_route_entry":     {Tok: aliResource(cenMod, "TransitRouterRouteEntry")},
			"alicloud_cen_transit_router_vbr_attachment":  {Tok: aliResource(cenMod, "TransitRouterVbrAttachment")},
			"alicloud_cen_transit_router_vpc_attachment":  {Tok: aliResource(cenMod, "TransitRouterVpcAttachment")},
			"alicloud_cen_transit_router_peer_attachment": {Tok: aliResource(cenMod, "TransitRouterPeerAttachment")},
			"alicloud_cen_traffic_marking_policy":         {Tok: aliResource(cenMod, "TrafficMarkingPolicy")},

			// Clickhouse
			"alicloud_click_house_db_cluster":    {Tok: aliResource(clickHouseMod, "DbCluster")},
			"alicloud_click_house_account":       {Tok: aliResource(clickHouseMod, "Account")},
			"alicloud_click_house_backup_policy": {Tok: aliResource(clickHouseMod, "BackupPolicy")},

			// Cloudauth
			"alicloud_cloudauth_face_config": {Tok: aliResource(cloudAuthMod, "FaceConfig")},

			// CloudConnect
			"alicloud_cloud_connect_network":            {Tok: aliResource(cloudConnectMod, "Network")},
			"alicloud_cloud_connect_network_attachment": {Tok: aliResource(cloudConnectMod, "NetworkAttachment")},
			"alicloud_cloud_connect_network_grant":      {Tok: aliResource(cloudConnectMod, "NetworkGrant")},

			// CloudFirewall
			"alicloud_cloud_firewall_control_policy":       {Tok: aliResource(cloudFirewallMod, "ControlPolicy")},
			"alicloud_cloud_firewall_control_policy_order": {Tok: aliResource(cloudFirewallMod, "ControlPolicyOrder")},
			"alicloud_cloud_firewall_instance":             {Tok: aliResource(cloudFirewallMod, "Instance")},
			"alicloud_cloud_firewall_address_book":         {Tok: aliResource(cloudFirewallMod, "AddressBook")},

			// CloudStorageGateway
			"alicloud_cloud_storage_gateway_storage_bundle": {Tok: aliResource(cloudStorageGatewayMod, "StorageBundle")},
			"alicloud_cloud_storage_gateway_gateway":        {Tok: aliResource(cloudStorageGatewayMod, "Gateway")},
			"alicloud_cloud_storage_gateway_express_sync":   {Tok: aliResource(cloudStorageGatewayMod, "ExpressSync")},
			"alicloud_cloud_storage_gateway_express_sync_share_attachment": {
				Tok: aliResource(cloudStorageGatewayMod, "ExpressSyncShareAttachment"),
			},
			"alicloud_cloud_storage_gateway_gateway_block_volume": {
				Tok: aliResource(cloudStorageGatewayMod, "GatewayBlockVolume"),
			},
			"alicloud_cloud_storage_gateway_gateway_cache_disk": {
				Tok: aliResource(cloudStorageGatewayMod, "GatewayCacheDisk"),
			},
			"alicloud_cloud_storage_gateway_gateway_file_share": {
				Tok: aliResource(cloudStorageGatewayMod, "GatewayFileShare"),
			},
			"alicloud_cloud_storage_gateway_gateway_logging": {
				Tok: aliResource(cloudStorageGatewayMod, "GatewayLogging"),
			},
			"alicloud_cloud_storage_gateway_gateway_smb_user": {
				Tok: aliResource(cloudStorageGatewayMod, "GatewaySmbUser"),
			},

			// cloudSso
			"alicloud_cloud_sso_directory":                         {Tok: aliResource(cloudSsoMod, "Directory")},
			"alicloud_cloud_sso_scim_server_credential":            {Tok: aliResource(cloudSsoMod, "ScimServerCredential")},
			"alicloud_cloud_sso_group":                             {Tok: aliResource(cloudSsoMod, "Group")},
			"alicloud_cloud_sso_access_configuration":              {Tok: aliResource(cloudSsoMod, "AccessConfiguration")},
			"alicloud_cloud_sso_user":                              {Tok: aliResource(cloudSsoMod, "User")},
			"alicloud_cloud_sso_user_attachment":                   {Tok: aliResource(cloudSsoMod, "UserAttachment")},
			"alicloud_cloud_sso_access_assignment":                 {Tok: aliResource(cloudSsoMod, "AccessManagement")},
			"alicloud_cloud_sso_access_configuration_provisioning": {Tok: aliResource(cloudSsoMod, "AccessConfigurationProvisioning")},

			// CMS
			"alicloud_cms_alarm":               {Tok: aliResource(cmsMod, "Alarm")},
			"alicloud_cms_alarm_contact":       {Tok: aliResource(cmsMod, "AlarmContact")},
			"alicloud_cms_alarm_contact_group": {Tok: aliResource(cmsMod, "AlarmContactGroup")},
			"alicloud_cms_site_monitor": {
				Tok: aliResource(cmsMod, "SiteMonitor"),
				Docs: &tfbridge.DocInfo{
					Source: "cms_sitemonitor.markdown",
				},
			},
			"alicloud_cms_group_metric_rule":       {Tok: aliResource(cmsMod, "GroupMetricRule")},
			"alicloud_cms_monitor_group":           {Tok: aliResource(cmsMod, "MonitorGroup")},
			"alicloud_cms_monitor_group_instances": {Tok: aliResource(cmsMod, "MonitorGroupInstances")},
			"alicloud_cms_metric_rule_template":    {Tok: aliResource(cmsMod, "MetricRuleTemplate")},
			"alicloud_cms_dynamic_tag_group":       {Tok: aliResource(cmsMod, "DynamicTagGroup")},
			"alicloud_cms_namespace": {
				Tok: aliResource(cmsMod, "Namespace"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"namespace": {
						CSharpName: "NamespaceName",
					},
				},
			},
			"alicloud_cms_sls_group":               {Tok: aliResource(cmsMod, "SlsGroup")},
			"alicloud_cms_hybrid_monitor_fc_task":  {Tok: aliResource(cmsMod, "HybridMonitorFcTask")},
			"alicloud_cms_hybrid_monitor_sls_task": {Tok: aliResource(cmsMod, "HybridMonitorSlsTask")},
			"alicloud_cms_event_rule":              {Tok: aliResource(cmsMod, "EventRule")},

			// Config
			"alicloud_config_configuration_recorder":    {Tok: aliResource(cfgMod, "ConfigurationRecorder")},
			"alicloud_config_delivery_channel":          {Tok: aliResource(cfgMod, "DeliveryChannel")},
			"alicloud_config_rule":                      {Tok: aliResource(cfgMod, "Rule")},
			"alicloud_config_aggregate_compliance_pack": {Tok: aliResource(cfgMod, "AggregateCompliancePack")},
			"alicloud_config_aggregate_config_rule":     {Tok: aliResource(cfgMod, "AggregateConfigRule")},
			"alicloud_config_aggregator":                {Tok: aliResource(cfgMod, "Aggregator")},
			"alicloud_config_compliance_pack":           {Tok: aliResource(cfgMod, "CompliancePack")},
			"alicloud_config_delivery":                  {Tok: aliResource(cfgMod, "Delivery")},
			"alicloud_config_aggregate_delivery":        {Tok: aliResource(cfgMod, "AggregateDelivery")},

			// CR
			"alicloud_cr_repo":                {Tok: aliResource(crMod, "Repo")},
			"alicloud_cr_namespace":           {Tok: aliResource(crMod, "Namespace")},
			"alicloud_cr_ee_instance":         {Tok: aliResource(crMod, "RegistryEnterpriseInstance")},
			"alicloud_cr_endpoint_acl_policy": {Tok: aliResource(crMod, "EndpointAclPolicy")},
			"alicloud_cr_chart_namespace":     {Tok: aliResource(crMod, "ChartNamespace")},
			"alicloud_cr_chart_repository":    {Tok: aliResource(crMod, "ChartRepository")},
			"alicloud_cr_chain":               {Tok: aliResource(crMod, "Chain")},

			// CS
			"alicloud_container_cluster": {
				Tok: aliResource(csMod, "Cluster"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_cs_application": {
				Tok: aliResource(csMod, "Application"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_cs_kubernetes": {Tok: aliResource(csMod, "Kubernetes")},
			"alicloud_cs_swarm": {
				Tok: aliResource(csMod, "Swarm"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_cs_kubernetes_autoscaler":  {Tok: aliResource(csMod, "KubernetesAutoscaler")},
			"alicloud_cs_managed_kubernetes":     {Tok: aliResource(csMod, "ManagedKubernetes")},
			"alicloud_cs_serverless_kubernetes":  {Tok: aliResource(csMod, "ServerlessKubernetes")},
			"alicloud_cr_ee_repo":                {Tok: aliResource(csMod, "RegistryEnterpriseRepo")},
			"alicloud_cr_ee_namespace":           {Tok: aliResource(csMod, "RegistryEnterpriseNamespace")},
			"alicloud_cr_ee_sync_rule":           {Tok: aliResource(csMod, "RegistryEnterpriseSyncRule")},
			"alicloud_cs_kubernetes_node_pool":   {Tok: aliResource(csMod, "NodePool")},
			"alicloud_cs_edge_kubernetes":        {Tok: aliResource(csMod, "EdgeKubernetes")},
			"alicloud_cs_kubernetes_permissions": {Tok: aliResource(csMod, "KubernetesPermission")},
			"alicloud_cs_autoscaling_config":     {Tok: aliResource(csMod, "AutoscalingConfig")},
			"alicloud_cs_kubernetes_addon":       {Tok: aliResource(csMod, "KubernetesAddon")},

			// DataHub
			"alicloud_datahub_project":      {Tok: aliResource(datahubMod, "Project")},
			"alicloud_datahub_subscription": {Tok: aliResource(datahubMod, "Subscription")},
			"alicloud_datahub_topic":        {Tok: aliResource(datahubMod, "Topic")},

			// Dataworks
			"alicloud_data_works_folder": {Tok: aliResource(dataWorksMod, "Folder")},

			// Database Filesystem
			"alicloud_dbfs_instance":            {Tok: aliResource(databaseFilesystemMod, "Instance")},
			"alicloud_dbfs_instance_attachment": {Tok: aliResource(databaseFilesystemMod, "InstanceAttachment")},
			"alicloud_dbfs_service_linked_role": {Tok: aliResource(databaseFilesystemMod, "ServiceLinkedRole")},
			"alicloud_dbfs_snapshot":            {Tok: aliResource(databaseFilesystemMod, "Snapshot")},

			// Database Gateway
			"alicloud_database_gateway_gateway": {Tok: aliResource(databaseGatewayMod, "Gateway")},

			// DB
			"alicloud_db_account":                         {Tok: aliResource(rdsMod, "Account")},
			"alicloud_db_account_privilege":               {Tok: aliResource(rdsMod, "AccountPrivilege")},
			"alicloud_db_backup_policy":                   {Tok: aliResource(rdsMod, "BackupPolicy")},
			"alicloud_db_connection":                      {Tok: aliResource(rdsMod, "Connection")},
			"alicloud_db_database":                        {Tok: aliResource(rdsMod, "Database")},
			"alicloud_db_instance":                        {Tok: aliResource(rdsMod, "Instance")},
			"alicloud_db_read_write_splitting_connection": {Tok: aliResource(rdsMod, "ReadWriteSplittingConnection")},
			"alicloud_db_readonly_instance":               {Tok: aliResource(rdsMod, "ReadOnlyInstance")},
			"alicloud_rds_parameter_group":                {Tok: aliResource(rdsMod, "RdsParameterGroup")},
			"alicloud_rds_account":                        {Tok: aliResource(rdsMod, "RdsAccount")},
			"alicloud_rds_backup":                         {Tok: aliResource(rdsMod, "RdsBackup")},
			"alicloud_rds_clone_db_instance":              {Tok: aliResource(rdsMod, "RdsCloneDbInstance")},
			"alicloud_rds_upgrade_db_instance":            {Tok: aliResource(rdsMod, "RdsUpgradeDbInstance")},

			// DCDN
			"alicloud_dcdn_domain":        {Tok: aliResource(dcdnMod, "Domain")},
			"alicloud_dcdn_domain_config": {Tok: aliResource(dcdnMod, "DomainConfig")},
			"alicloud_dcdn_ipa_domain":    {Tok: aliResource(dcdnMod, "IpaDomain")},

			// DDOS
			"alicloud_ddoscoo_scheduler_rule":       {Tok: aliResource(ddosMod, "SchedulerRule")},
			"alicloud_ddoscoo_domain_aliResource":   {Tok: aliResource(ddosMod, "DomainResource")},
			"alicloud_ddoscoo_port":                 {Tok: aliResource(ddosMod, "Port")},
			"alicloud_ddos_basic_defense_threshold": {Tok: aliResource(ddosMod, "BasicDefenseThreshold")},
			"alicloud_ddosbgp_ip":                   {Tok: aliResource(ddosMod, "BgpIp")},

			// Dfs
			"alicloud_dfs_access_group": {Tok: aliResource(dfsMod, "AccessGroup")},
			"alicloud_dfs_file_system":  {Tok: aliResource(dfsMod, "FileSystem")},
			"alicloud_dfs_access_rule":  {Tok: aliResource(dfsMod, "AccessRule")},
			"alicloud_dfs_mount_point":  {Tok: aliResource(dfsMod, "MountPoint")},

			// Direct Mail
			"alicloud_direct_mail_receivers":    {Tok: aliResource(directMailMod, "Receivers")},
			"alicloud_direct_mail_mail_address": {Tok: aliResource(directMailMod, "MailAddress")},
			"alicloud_direct_mail_domain":       {Tok: aliResource(directMailMod, "Domain")},
			"alicloud_direct_mail_tag":          {Tok: aliResource(directMailMod, "Tag")},

			// DMS
			"alicloud_dms_enterprise_instance": {Tok: aliResource(dmsMod, "EnterpriseInstance")},
			"alicloud_dms_enterprise_user":     {Tok: aliResource(dmsMod, "EnterpriseUser")},

			// DNS
			"alicloud_dns": {
				Tok:                aliResource(dnsMod, "Domain"),
				DeprecationMessage: "This resource has been deprecated in favour of DnsDomain",
			},
			"alicloud_dns_group":                {Tok: aliResource(dnsMod, "Group")},
			"alicloud_dns_record":               {Tok: aliResource(dnsMod, "Record")},
			"alicloud_dns_domain_attachment":    {Tok: aliResource(dnsMod, "DomainAttachment")},
			"alicloud_dns_instance":             {Tok: aliResource(dnsMod, "Instance")},
			"alicloud_dns_domain":               {Tok: aliResource(dnsMod, "DnsDomain")},
			"alicloud_alidns_domain_group":      {Tok: aliResource(dnsMod, "DomainGroup")},
			"alicloud_alidns_record":            {Tok: aliResource(dnsMod, "AlidnsRecord")},
			"alicloud_alidns_domain":            {Tok: aliResource(dnsMod, "AlidnsDomain")},
			"alicloud_alidns_instance":          {Tok: aliResource(dnsMod, "AlidnsInstance")},
			"alicloud_alidns_domain_attachment": {Tok: aliResource(dnsMod, "AlidnsDomainAttachment")},
			"alicloud_alidns_custom_line":       {Tok: aliResource(dnsMod, "CustomLine")},
			"alicloud_alidns_gtm_instance":      {Tok: aliResource(dnsMod, "GtmInstance")},
			"alicloud_alidns_monitor_config":    {Tok: aliResource(dnsMod, "MonitorConfig")},
			"alicloud_alidns_address_pool":      {Tok: aliResource(dnsMod, "AddressPool")},
			"alicloud_alidns_access_strategy":   {Tok: aliResource(dnsMod, "AccessStrategy")},

			// Drds
			"alicloud_drds_instance": {Tok: aliResource(drdsMod, "Instance")},

			// Dts
			"alicloud_dts_job_monitor_rule":         {Tok: aliResource(dtsMod, "JobMonitorRule")},
			"alicloud_dts_migration_instance":       {Tok: aliResource(dtsMod, "MigrationInstance")},
			"alicloud_dts_migration_job":            {Tok: aliResource(dtsMod, "MigrationJob")},
			"alicloud_dts_subscription_job":         {Tok: aliResource(dtsMod, "SubscriptionJob")},
			"alicloud_dts_synchronization_instance": {Tok: aliResource(dtsMod, "SynchronizationInstance")},
			"alicloud_dts_synchronization_job":      {Tok: aliResource(dtsMod, "SynchronizationJob")},
			"alicloud_dts_consumer_channel":         {Tok: aliResource(dtsMod, "ConsumerChannel")},

			// Eais
			"alicloud_eais_instance": {Tok: aliResource(eaisMod, "Instance")},

			// Eci
			"alicloud_eci_openapi_image_cache": {
				Tok: aliResource(eciMod, "OpenApiImageCache"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_eci_image_cache":     {Tok: aliResource(eciMod, "ImageCache")},
			"alicloud_eci_container_group": {Tok: aliResource(eciMod, "ContainerGroup")},
			"alicloud_eci_virtual_node":    {Tok: aliResource(eciMod, "VirtualNode")},

			// Ecp
			"alicloud_ecp_instance": {Tok: aliResource(ecpMod, "Instance")},
			"alicloud_ecp_key_pair": {Tok: aliResource(ecpMod, "KeyPair")},

			// ECS
			"alicloud_auto_provisioning_group": {Tok: aliResource(ecsMod, "AutoProvisioningGroup")},
			"alicloud_disk":                    {Tok: aliResource(ecsMod, "Disk")},
			"alicloud_disk_attachment":         {Tok: aliResource(ecsMod, "DiskAttachment")},
			"alicloud_launch_template":         {Tok: aliResource(ecsMod, "LaunchTemplate")},
			"alicloud_eip": {
				Tok:                aliResource(ecsMod, "Eip"),
				DeprecationMessage: "This resource has been deprecated in favour of the EipAddress aliResource",
			},
			"alicloud_eip_address":            {Tok: aliResource(ecsMod, "EipAddress")},
			"alicloud_eip_association":        {Tok: aliResource(ecsMod, "EipAssociation")},
			"alicloud_instance":               {Tok: aliResource(ecsMod, "Instance")},
			"alicloud_key_pair":               {Tok: aliResource(ecsMod, "KeyPair")},
			"alicloud_key_pair_attachment":    {Tok: aliResource(ecsMod, "KeyPairAttachment")},
			"alicloud_image":                  {Tok: aliResource(ecsMod, "Image")},
			"alicloud_image_copy":             {Tok: aliResource(ecsMod, "ImageCopy")},
			"alicloud_image_export":           {Tok: aliResource(ecsMod, "ImageExport")},
			"alicloud_image_share_permission": {Tok: aliResource(ecsMod, "ImageSharePermission")},
			"alicloud_security_group":         {Tok: aliResource(ecsMod, "SecurityGroup")},
			"alicloud_security_group_rule":    {Tok: aliResource(ecsMod, "SecurityGroupRule")},
			"alicloud_reserved_instance":      {Tok: aliResource(ecsMod, "ReservedInstance")},
			"alicloud_snapshot":               {Tok: aliResource(ecsMod, "Snapshot")},
			"alicloud_snapshot_policy":        {Tok: aliResource(ecsMod, "SnapshotPolicy")},
			"alicloud_copy_image": {
				Tok: aliResource(ecsMod, "CopyImage"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_image_import": {Tok: aliResource(ecsMod, "ImageImport")},
			"alicloud_ecs_dedicated_host": {
				Tok: aliResource(ecsMod, "DedicatedHost"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"network_attributes": {
						MaxItemsOne: tfbridge.False(),
						Name:        "networkAttributes",
					},
				},
			},
			"alicloud_ecs_hpc_cluster":                     {Tok: aliResource(ecsMod, "HpcCluster")},
			"alicloud_ecs_command":                         {Tok: aliResource(ecsMod, "Command")},
			"alicloud_ecs_auto_snapshot_policy":            {Tok: aliResource(ecsMod, "AutoSnapshotPolicy")},
			"alicloud_ecs_snapshot":                        {Tok: aliResource(ecsMod, "EcsSnapshot")},
			"alicloud_ecs_storage_capacity_unit":           {Tok: aliResource(ecsMod, "StorageCapacityUnit")},
			"alicloud_ecs_launch_template":                 {Tok: aliResource(ecsMod, "EcsLaunchTemplate")},
			"alicloud_ecs_key_pair":                        {Tok: aliResource(ecsMod, "EcsKeyPair")},
			"alicloud_ecs_key_pair_attachment":             {Tok: aliResource(ecsMod, "EcsKeyPairAttachment")},
			"alicloud_ecs_auto_snapshot_policy_attachment": {Tok: aliResource(ecsMod, "EcsAutoSnapshotPolicyAttachment")},
			"alicloud_ecs_disk":                            {Tok: aliResource(ecsMod, "EcsDisk")},
			"alicloud_ecs_disk_attachment":                 {Tok: aliResource(ecsMod, "EcsDiskAttachment")},
			"alicloud_ecs_network_interface":               {Tok: aliResource(ecsMod, "EcsNetworkInterface")},
			"alicloud_ecs_network_interface_attachment":    {Tok: aliResource(ecsMod, "EcsNetworkInterfaceAttachment")},
			"alicloud_ecs_deployment_set":                  {Tok: aliResource(ecsMod, "EcsDeploymentSet")},
			"alicloud_ecs_dedicated_host_cluster":          {Tok: aliResource(ecsMod, "EcsDedicatedHostCluster")},
			"alicloud_ecs_session_manager_status":          {Tok: aliResource(ecsMod, "EcsSessionManagerStatus")},
			"alicloud_ecs_prefix_list":                     {Tok: aliResource(ecsMod, "EcsPrefixList")},
			"alicloud_ecs_image_component":                 {Tok: aliResource(ecsMod, "EcsImageComponent")},
			"alicloud_ecs_snapshot_group":                  {Tok: aliResource(ecsMod, "EcsSnapshotGroup")},
			"alicloud_ecs_image_pipeline":                  {Tok: aliResource(ecsMod, "EcsImagePipeline")},
			"alicloud_ecs_invocation":                      {Tok: aliResource(ecsMod, "EcsInvocation")},
			"alicloud_ecs_network_interface_permission":    {Tok: aliResource(ecsMod, "EcsNetworkInterfacePermission")},
			"alicloud_ecs_instance_set":                    {Tok: aliResource(ecsMod, "EcsInstanceSet")},
			"alicloud_ecs_activation":                      {Tok: aliResource(ecsMod, "Activation")},

			// Edas
			"alicloud_edas_application":                 {Tok: aliResource(edasMod, "Application")},
			"alicloud_edas_deploy_group":                {Tok: aliResource(edasMod, "DeployGroup")},
			"alicloud_edas_application_scale":           {Tok: aliResource(edasMod, "ApplicationScale")},
			"alicloud_edas_slb_attachment":              {Tok: aliResource(edasMod, "SlbAttachment")},
			"alicloud_edas_cluster":                     {Tok: aliResource(edasMod, "Cluster")},
			"alicloud_edas_instance_cluster_attachment": {Tok: aliResource(edasMod, "InstanceClusterAttachment")},
			"alicloud_edas_application_deployment":      {Tok: aliResource(edasMod, "ApplicationDeployment")},
			"alicloud_edas_k8s_cluster":                 {Tok: aliResource(edasMod, "K8sCluster")},
			"alicloud_edas_k8s_application":             {Tok: aliResource(edasMod, "K8sApplication")},
			"alicloud_edas_namespace":                   {Tok: aliResource(edasMod, "Namespace")},

			// Eds
			"alicloud_ecd_policy_group":             {Tok: aliResource(edsMod, "EcdPolicyGroup")},
			"alicloud_ecd_simple_office_site":       {Tok: aliResource(edsMod, "SimpleOfficeSite")},
			"alicloud_ecd_nas_file_system":          {Tok: aliResource(edsMod, "NasFileSystem")},
			"alicloud_ecd_desktop":                  {Tok: aliResource(edsMod, "Desktop")},
			"alicloud_ecd_network_package":          {Tok: aliResource(edsMod, "NetworkPackage")},
			"alicloud_ecd_user":                     {Tok: aliResource(edsMod, "User")},
			"alicloud_ecd_image":                    {Tok: aliResource(edsMod, "Image")},
			"alicloud_ecd_command":                  {Tok: aliResource(edsMod, "Command")},
			"alicloud_ecd_snapshot":                 {Tok: aliResource(edsMod, "Snapshot")},
			"alicloud_ecd_bundle":                   {Tok: aliResource(edsMod, "Bundle")},
			"alicloud_ecd_ad_connector_directory":   {Tok: aliResource(edsMod, "AdConnectorDirectory")},
			"alicloud_ecd_ram_directory":            {Tok: aliResource(edsMod, "RamDirectory")},
			"alicloud_ecd_ad_connector_office_site": {Tok: aliResource(edsMod, "AdConnectorOfficeSite")},
			"alicloud_ecd_custom_property":          {Tok: aliResource(edsMod, "CustomProperty")},

			// Ehpc
			"alicloud_ehpc_job_template": {Tok: aliResource(ehpcMod, "JobTemplate")},
			"alicloud_ehpc_cluster":      {Tok: aliResource(ehpcMod, "Cluster")},

			// EipAnycast
			"alicloud_eipanycast_anycast_eip_address":            {Tok: aliResource(eipAnyCastMod, "AnycastEipAddress")},
			"alicloud_eipanycast_anycast_eip_address_attachment": {Tok: aliResource(eipAnyCastMod, "AnycastEipAddressAttachment")},

			// Ens
			"alicloud_ens_key_pair": {Tok: aliResource(ensMod, "KeyPair")},

			// ESS
			"alicloud_ess_alarm":      {Tok: aliResource(essMod, "Alarm")},
			"alicloud_ess_attachment": {Tok: aliResource(essMod, "Attachment")},
			"alicloud_ess_lifecycle_hook": {
				Tok: aliResource(essMod, "LifecycleHook"),
				Docs: &tfbridge.DocInfo{
					Source: "ess_scaling_lifecycle_hook.html.markdown",
				},
			},
			"alicloud_ess_scaling_configuration":       {Tok: aliResource(essMod, "ScalingConfiguration")},
			"alicloud_ess_scaling_group":               {Tok: aliResource(essMod, "ScalingGroup")},
			"alicloud_ess_scaling_rule":                {Tok: aliResource(essMod, "ScalingRule")},
			"alicloud_ess_schedule":                    {Tok: aliResource(essMod, "Schedule")},
			"alicloud_ess_notification":                {Tok: aliResource(essMod, "Notification")},
			"alicloud_ess_scalinggroup_vserver_groups": {Tok: aliResource(essMod, "ScalingGroupVServerGroups")},
			"alicloud_ess_scheduled_task":              {Tok: aliResource(essMod, "ScheduledTask")},
			"alicloud_ess_alb_server_group_attachment": {Tok: aliResource(essMod, "AlbServerGroupAttachment")},
			"alicloud_ess_eci_scaling_configuration":   {Tok: aliResource(essMod, "EciScalingConfiguration")},
			"alicloud_ess_suspend_process":             {Tok: aliResource(essMod, "SuspendProcess")},

			// ElasticSearch
			"alicloud_elasticsearch_instance": {
				Tok: aliResource(elasticsearchMod, "Instance"),
				Docs: &tfbridge.DocInfo{
					Source: "elasticsearch.html.markdown",
				},
			},

			// EMR
			"alicloud_emr_cluster": {Tok: aliResource(emrMod, "Cluster")},

			// Eventbridge
			"alicloud_event_bridge_event_bus": {Tok: aliResource(eventBridgeMod, "EventBus")},
			"alicloud_event_bridge_rule":      {Tok: aliResource(eventBridgeMod, "Rule")},
			"alicloud_event_bridge_slr": {
				Tok: aliResource(eventBridgeMod, "Slr"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_event_bridge_event_source":        {Tok: aliResource(eventBridgeMod, "EventSource")},
			"alicloud_event_bridge_service_linked_role": {Tok: aliResource(eventBridgeMod, "ServiceLinkedRole")},

			// ExpressConnect
			"alicloud_express_connect_physical_connection":   {Tok: aliResource(expressConnect, "PhysicalConnection")},
			"alicloud_express_connect_virtual_border_router": {Tok: aliResource(expressConnect, "VirtualBorderRouter")},

			// FC
			"alicloud_fc_function":                     {Tok: aliResource(fcMod, "Function")},
			"alicloud_fc_service":                      {Tok: aliResource(fcMod, "Service")},
			"alicloud_fc_trigger":                      {Tok: aliResource(fcMod, "Trigger")},
			"alicloud_fc_custom_domain":                {Tok: aliResource(fcMod, "CustomDomain")},
			"alicloud_fc_function_async_invoke_config": {Tok: aliResource(fcMod, "FunctionAsyncInvokeConfig")},
			"alicloud_fc_alias":                        {Tok: aliResource(fcMod, "Alias")},
			"alicloud_fc_layer_version":                {Tok: aliResource(fcMod, "LayerVersion")},

			// FNF
			"alicloud_fnf_schedule":  {Tok: aliResource(fnfMod, "Schedule")},
			"alicloud_fnf_flow":      {Tok: aliResource(fnfMod, "Flow")},
			"alicloud_fnf_execution": {Tok: aliResource(fnfMod, "Execution")},

			// Ga
			"alicloud_ga_listener":                        {Tok: aliResource(gaMod, "Listener")},
			"alicloud_ga_accelerator":                     {Tok: aliResource(gaMod, "Accelerator")},
			"alicloud_ga_bandwidth_package":               {Tok: aliResource(gaMod, "BandwidthPackage")},
			"alicloud_ga_endpoint_group":                  {Tok: aliResource(gaMod, "EndpointGroup")},
			"alicloud_ga_ip_set":                          {Tok: aliResource(gaMod, "IpSet")},
			"alicloud_ga_bandwidth_package_attachment":    {Tok: aliResource(gaMod, "BandwidthPackageAttachment")},
			"alicloud_ga_forwarding_rule":                 {Tok: aliResource(gaMod, "ForwardingRule")},
			"alicloud_ga_acl":                             {Tok: aliResource(gaMod, "Acl")},
			"alicloud_ga_acl_attachment":                  {Tok: aliResource(gaMod, "AclAttachment")},
			"alicloud_ga_additional_certificate":          {Tok: aliResource(gaMod, "AdditionalCertificate")},
			"alicloud_ga_accelerator_spare_ip_attachment": {Tok: aliResource(gaMod, "AcceleratorSpareIpAttachment")},

			// Gpdb
			"alicloud_gpdb_connection":       {Tok: aliResource(gpdbMod, "Connection")},
			"alicloud_gpdb_instance":         {Tok: aliResource(gpdbMod, "Instance")},
			"alicloud_gpdb_elastic_instance": {Tok: aliResource(gpdbMod, "ElasticInstance")},
			"alicloud_gpdb_account":          {Tok: aliResource(gpdbMod, "Account")},

			// Graph Database
			"alicloud_graph_database_db_instance": {Tok: aliResource(graphDatabaseMod, "DbInstance")},

			// Hbase
			"alicloud_hbase_instance": {Tok: aliResource(hbaseMod, "Instance")},

			// Hbr
			"alicloud_hbr_vault":              {Tok: aliResource(hbrMod, "Vault")},
			"alicloud_hbr_oss_backup_plan":    {Tok: aliResource(hbrMod, "OssBackupPlan")},
			"alicloud_hbr_nas_backup_plan":    {Tok: aliResource(hbrMod, "NasBackupPlan")},
			"alicloud_hbr_ecs_backup_plan":    {Tok: aliResource(hbrMod, "EcsBackupPlan")},
			"alicloud_hbr_ecs_backup_client":  {Tok: aliResource(hbrMod, "EcsBackupClient")},
			"alicloud_hbr_restore_job":        {Tok: aliResource(hbrMod, "RestoreJob")},
			"alicloud_hbr_server_backup_plan": {Tok: aliResource(hbrMod, "ServerBackupPlan")},
			"alicloud_hbr_replication_vault":  {Tok: aliResource(hbrMod, "ReplicationVault")},
			"alicloud_hbr_ots_backup_plan":    {Tok: aliResource(hbrMod, "OtsBackupPlan")},
			"alicloud_hbr_hana_instance":      {Tok: aliResource(hbrMod, "HanaInstance")},
			"alicloud_hbr_hana_backup_plan":   {Tok: aliResource(hbrMod, "HanaBackupPlan")},

			// Iot
			"alicloud_iot_device_group": {Tok: aliResource(iotMod, "DeviceGroup")},

			// Imm
			"alicloud_imm_project": {
				Tok: aliResource(immMod, "Project"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"project": {
						CSharpName: "ProjectName",
					},
				},
			},

			// Imp
			"alicloud_imp_app_template": {Tok: aliResource(impMod, "AppTemplate")},

			// KMS
			"alicloud_kms_key": {
				Tok: aliResource(kmsMod, "Key"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"key_state": {
						CSharpName: "KeyStatus",
					},
				},
			},
			"alicloud_kms_key_version": {Tok: aliResource(kmsMod, "KeyVersion")},
			"alicloud_kms_ciphertext":  {Tok: aliResource(kmsMod, "Ciphertext")},
			"alicloud_kms_secret":      {Tok: aliResource(kmsMod, "Secret")},
			"alicloud_kms_alias":       {Tok: aliResource(kmsMod, "Alias")},

			// KVStore
			"alicloud_kvstore_backup_policy":    {Tok: aliResource(kvstoreMod, "BackupPolicy")},
			"alicloud_kvstore_instance":         {Tok: aliResource(kvstoreMod, "Instance")},
			"alicloud_kvstore_account":          {Tok: aliResource(kvstoreMod, "Account")},
			"alicloud_kvstore_connection":       {Tok: aliResource(kvstoreMod, "Connection")},
			"alicloud_kvstore_audit_log_config": {Tok: aliResource(kvstoreMod, "AuditLogConfig")},

			// Lindorm
			"alicloud_lindorm_instance": {Tok: aliResource(lindormMod, "Instance")},

			// Log
			"alicloud_log_machine_group":   {Tok: aliResource(logMod, "MachineGroup")},
			"alicloud_log_project":         {Tok: aliResource(logMod, "Project")},
			"alicloud_log_store":           {Tok: aliResource(logMod, "Store")},
			"alicloud_log_store_index":     {Tok: aliResource(logMod, "StoreIndex")},
			"alicloud_logtail_attachment":  {Tok: aliResource(logMod, "LogTailAttachment")},
			"alicloud_logtail_config":      {Tok: aliResource(logMod, "LogTailConfig")},
			"alicloud_log_alert":           {Tok: aliResource(logMod, "Alert")},
			"alicloud_log_audit":           {Tok: aliResource(logMod, "Audit")},
			"alicloud_log_dashboard":       {Tok: aliResource(logMod, "Dashboard")},
			"alicloud_log_etl":             {Tok: aliResource(logMod, "Etl")},
			"alicloud_log_oss_shipper":     {Tok: aliResource(logMod, "OssShipper")},
			"alicloud_log_ingestion":       {Tok: aliResource(logMod, "Ingestion")},
			"alicloud_log_resource":        {Tok: aliResource(logMod, "Resource")},
			"alicloud_log_resource_record": {Tok: aliResource(logMod, "ResourceRecord")},

			// Marketplace
			"alicloud_market_order": {Tok: aliResource(marketPlaceMod, "Order")},

			// MaxCompute
			"alicloud_maxcompute_project": {Tok: aliResource(maxComputeMod, "Project")},

			// MHub
			"alicloud_mhub_product": {Tok: aliResource(mhubMod, "Product")},
			"alicloud_mhub_app":     {Tok: aliResource(mhubMod, "App")},

			// MongoDb
			"alicloud_mongodb_instance":                         {Tok: aliResource(mongoDbMod, "Instance")},
			"alicloud_mongodb_sharding_instance":                {Tok: aliResource(mongoDbMod, "ShardingInstance")},
			"alicloud_mongodb_audit_policy":                     {Tok: aliResource(mongoDbMod, "AuditPolicy")},
			"alicloud_mongodb_account":                          {Tok: aliResource(mongoDbMod, "Account")},
			"alicloud_mongodb_serverless_instance":              {Tok: aliResource(mongoDbMod, "ServerlessInstance")},
			"alicloud_mongodb_sharding_network_public_address":  {Tok: aliResource(mongoDbMod, "ShardingNetworkPublicAddress")},
			"alicloud_mongodb_sharding_network_private_address": {Tok: aliResource(mongoDbMod, "ShardingNetworkPrivateAddress")},

			// Mns
			"alicloud_mns_queue":              {Tok: aliResource(mnsMod, "Queue")},
			"alicloud_mns_topic":              {Tok: aliResource(mnsMod, "Topic")},
			"alicloud_mns_topic_subscription": {Tok: aliResource(mnsMod, "TopicSubscription")},

			// Mse
			"alicloud_mse_cluster":          {Tok: aliResource(mseMod, "Cluster")},
			"alicloud_mse_gateway":          {Tok: aliResource(mseMod, "Gateway")},
			"alicloud_mse_znode":            {Tok: aliResource(mseMod, "Znode")},
			"alicloud_mse_engine_namespace": {Tok: aliResource(mseMod, "EngineNamespace")},

			// Nas
			"alicloud_nas_access_group":         {Tok: aliResource(nasMod, "AccessGroup")},
			"alicloud_nas_access_rule":          {Tok: aliResource(nasMod, "AccessRule")},
			"alicloud_nas_file_system":          {Tok: aliResource(nasMod, "FileSystem")},
			"alicloud_nas_mount_target":         {Tok: aliResource(nasMod, "MountTarget")},
			"alicloud_nas_snapshot":             {Tok: aliResource(nasMod, "Snapshot")},
			"alicloud_nas_lifecycle_policy":     {Tok: aliResource(nasMod, "LifecyclePolicy")},
			"alicloud_nas_fileset":              {Tok: aliResource(nasMod, "Fileset")},
			"alicloud_nas_data_flow":            {Tok: aliResource(nasMod, "DataFlow")},
			"alicloud_nas_auto_snapshot_policy": {Tok: aliResource(nasMod, "AutoSnapshotPolicy")},
			"alicloud_nas_recycle_bin":          {Tok: aliResource(nasMod, "RecycleBin")},

			// Oos
			"alicloud_oos_template":            {Tok: aliResource(oosMod, "Template")},
			"alicloud_oos_execution":           {Tok: aliResource(oosMod, "Execution")},
			"alicloud_oos_application":         {Tok: aliResource(oosMod, "Application")},
			"alicloud_oos_application_group":   {Tok: aliResource(oosMod, "ApplicationGroup")},
			"alicloud_oos_patch_baseline":      {Tok: aliResource(oosMod, "PatchBaseline")},
			"alicloud_oos_service_setting":     {Tok: aliResource(oosMod, "ServiceSetting")},
			"alicloud_oos_parameter":           {Tok: aliResource(oosMod, "Parameter")},
			"alicloud_oos_secret_parameter":    {Tok: aliResource(oosMod, "SecretParameter")},
			"alicloud_oos_state_configuration": {Tok: aliResource(oosMod, "StateConfiguration")},

			// OpenSearch
			"alicloud_open_search_app_group": {Tok: aliResource(openSearchMod, "AppGroup")},

			// Oss
			"alicloud_oss_bucket": {
				Tok: aliResource(ossMod, "Bucket"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"bucket": {
						CSharpName: "BucketName",
					},
				},
			},
			"alicloud_oss_bucket_object":      {Tok: aliResource(ossMod, "BucketObject")},
			"alicloud_oss_bucket_replication": {Tok: aliResource(ossMod, "BucketReplication")},

			// OTS Mod
			"alicloud_ots_instance":            {Tok: aliResource(otsMod, "Instance")},
			"alicloud_ots_instance_attachment": {Tok: aliResource(otsMod, "InstanceAttachment")},
			"alicloud_ots_table":               {Tok: aliResource(otsMod, "Table")},
			"alicloud_ots_tunnel":              {Tok: aliResource(otsMod, "Tunnel")},

			// PrivateLink
			"alicloud_privatelink_vpc_endpoint":         {Tok: aliResource(privateLinkMod, "VpcEndpoint")},
			"alicloud_privatelink_vpc_endpoint_service": {Tok: aliResource(privateLinkMod, "VpcEndpointService")},
			"alicloud_privatelink_vpc_endpoint_zone":    {Tok: aliResource(privateLinkMod, "VpcEndpointZone")},
			"alicloud_privatelink_vpc_endpoint_service_resource": {
				Tok: aliResource(privateLinkMod, "VpcEndpointServiceResource"),
			},
			"alicloud_privatelink_vpc_endpoint_service_user": {
				Tok: aliResource(privateLinkMod, "VpcEndpointServiceUser"),
			},
			"alicloud_privatelink_vpc_endpoint_connection": {
				Tok: aliResource(privateLinkMod, "VpcEndpointServiceConnection"),
			},

			// PolarDB
			"alicloud_polardb_account": {Tok: aliResource(polarDbMod, "Account")},
			"alicloud_polardb_account_privilege": {
				Tok: aliResource(polarDbMod, "AccountPrivilege"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"account_privilege": {
						CSharpName: "Privilege",
					},
				},
			},
			"alicloud_polardb_backup_policy":           {Tok: aliResource(polarDbMod, "BackupPolicy")},
			"alicloud_polardb_cluster":                 {Tok: aliResource(polarDbMod, "Cluster")},
			"alicloud_polardb_database":                {Tok: aliResource(polarDbMod, "Database")},
			"alicloud_polardb_endpoint_address":        {Tok: aliResource(polarDbMod, "EndpointAddress")},
			"alicloud_polardb_endpoint":                {Tok: aliResource(polarDbMod, "Endpoint")},
			"alicloud_polardb_global_database_network": {Tok: aliResource(polarDbMod, "GlobalDatabaseNetwork")},

			// Pvtz
			"alicloud_pvtz_zone":                   {Tok: aliResource(pvtzMod, "Zone")},
			"alicloud_pvtz_zone_attachment":        {Tok: aliResource(pvtzMod, "ZoneAttachment")},
			"alicloud_pvtz_zone_record":            {Tok: aliResource(pvtzMod, "ZoneRecord")},
			"alicloud_pvtz_user_vpc_authorization": {Tok: aliResource(pvtzMod, "UserVpcAuthorization")},
			"alicloud_pvtz_endpoint":               {Tok: aliResource(pvtzMod, "Endpoint")},
			"alicloud_pvtz_rule":                   {Tok: aliResource(pvtzMod, "Rule")},
			"alicloud_pvtz_rule_attachment":        {Tok: aliResource(pvtzMod, "RuleAttachment")},

			// QuickBI
			"alicloud_quick_bi_user": {Tok: aliResource(quickBiMod, "User")},

			// Quotas
			"alicloud_quotas_application_info": {
				Tok: aliResource(quotasMod, "ApplicationInfo"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_quotas_quota_alarm":       {Tok: aliResource(quotasMod, "QuotaAlarm")},
			"alicloud_quotas_quota_application": {Tok: aliResource(quotasMod, "QuotaApplication")},

			// Ram
			"alicloud_ram_access_key": {Tok: aliResource(ramMod, "AccessKey")},
			"alicloud_ram_account_alias": {
				Tok: aliResource(ramMod, "AccountAlias"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"account_alias": {
						CSharpName: "Alias",
					},
				},
			},
			"alicloud_ram_alias":                   {Tok: aliResource(ramMod, "Alias")},
			"alicloud_ram_group":                   {Tok: aliResource(ramMod, "Group")},
			"alicloud_ram_group_membership":        {Tok: aliResource(ramMod, "GroupMembership")},
			"alicloud_ram_group_policy_attachment": {Tok: aliResource(ramMod, "GroupPolicyAttachment")},
			"alicloud_ram_login_profile":           {Tok: aliResource(ramMod, "LoginProfile")},
			"alicloud_ram_policy":                  {Tok: aliResource(ramMod, "Policy")},
			"alicloud_ram_role":                    {Tok: aliResource(ramMod, "Role")},
			"alicloud_ram_role_attachment":         {Tok: aliResource(ramMod, "RoleAttachment")},
			"alicloud_ram_role_policy_attachment":  {Tok: aliResource(ramMod, "RolePolicyAttachment")},
			"alicloud_ram_user":                    {Tok: aliResource(ramMod, "User")},
			"alicloud_ram_user_policy_attachment":  {Tok: aliResource(ramMod, "UserPolicyAttachment")},
			"alicloud_ram_account_password_policy": {Tok: aliResource(ramMod, "AccountPasswordPolicy")},
			"alicloud_ram_saml_provider":           {Tok: aliResource(ramMod, "SamlProvider")},
			"alicloud_ram_security_preference":     {Tok: aliResource(ramMod, "SecurityPreference")},

			// Rdc
			"alicloud_rdc_organization": {Tok: aliResource(rdcMod, "Organization")},

			// ResourceManager
			"alicloud_resource_manager_handshake":          {Tok: aliResource(resourceManagerMod, "Handshake")},
			"alicloud_resource_manager_folder":             {Tok: aliResource(resourceManagerMod, "Folder")},
			"alicloud_resource_manager_resource_group":     {Tok: aliResource(resourceManagerMod, "ResourceGroup")},
			"alicloud_resource_manager_role":               {Tok: aliResource(resourceManagerMod, "Role")},
			"alicloud_resource_manager_account":            {Tok: aliResource(resourceManagerMod, "Account")},
			"alicloud_resource_manager_policy":             {Tok: aliResource(resourceManagerMod, "Policy")},
			"alicloud_resource_manager_resource_directory": {Tok: aliResource(resourceManagerMod, "ResourceDirectory")},
			"alicloud_resource_manager_policy_version":     {Tok: aliResource(resourceManagerMod, "PolicyVersion")},
			"alicloud_resource_manager_policy_attachment":  {Tok: aliResource(resourceManagerMod, "PolicyAttachment")},
			"alicloud_resource_manager_shared_resource":    {Tok: aliResource(resourceManagerMod, "SharedResource")},
			"alicloud_resource_manager_shared_target":      {Tok: aliResource(resourceManagerMod, "SharedTarget")},
			"alicloud_resource_manager_resource_share":     {Tok: aliResource(resourceManagerMod, "ResourceShare")},
			"alicloud_resource_manager_control_policy":     {Tok: aliResource(resourceManagerMod, "ControlPolicy")},
			"alicloud_resource_manager_control_policy_attachment": {
				Tok: aliResource(resourceManagerMod, "ControlPolicyAttachment"),
			},
			"alicloud_resource_manager_service_linked_role":     {Tok: aliResource(resourceManagerMod, "ServiceLinkedRole")},
			"alicloud_resource_manager_delegated_administrator": {Tok: aliResource(resourceManagerMod, "DelegatedAdministrator")},

			// RocketMQ
			"alicloud_ons_group":    {Tok: aliResource(rocketMqMod, "Group")},
			"alicloud_ons_instance": {Tok: aliResource(rocketMqMod, "Instance")},
			"alicloud_ons_topic": {
				Tok: aliResource(rocketMqMod, "Topic"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"topic": {
						CSharpName: "TopicDeprecated",
					},
				},
			},

			// ros
			"alicloud_ros_change_set":       {Tok: aliResource(rosMod, "ChangeSet")},
			"alicloud_ros_stack_group":      {Tok: aliResource(rosMod, "StackGroup")},
			"alicloud_ros_stack":            {Tok: aliResource(rosMod, "Stack")},
			"alicloud_ros_template":         {Tok: aliResource(rosMod, "Template")},
			"alicloud_ros_stack_instance":   {Tok: aliResource(rosMod, "StackInstance")},
			"alicloud_ros_template_scratch": {Tok: aliResource(rosMod, "TemplateScratch")},

			// Sae
			"alicloud_sae_namespace":                {Tok: aliResource(saeMod, "Namespace")},
			"alicloud_sae_config_map":               {Tok: aliResource(saeMod, "ConfigMap")},
			"alicloud_sae_application":              {Tok: aliResource(saeMod, "Application")},
			"alicloud_sae_ingress":                  {Tok: aliResource(saeMod, "Ingress")},
			"alicloud_sae_application_scaling_rule": {Tok: aliResource(saeMod, "ApplicationScalingRule")},
			"alicloud_sae_grey_tag_route":           {Tok: aliResource(saeMod, "GreyTagRoute")},
			"alicloud_sae_load_balancer_internet":   {Tok: aliResource(saeMod, "LoadBalancerInternet")},
			"alicloud_sae_load_balancer_intranet":   {Tok: aliResource(saeMod, "LoadBalancerIntranet")},

			// Sag
			"alicloud_sag_acl":         {Tok: aliResource(rocketMqMod, "Acl")},
			"alicloud_sag_acl_rule":    {Tok: aliResource(rocketMqMod, "AclRule")},
			"alicloud_sag_client_user": {Tok: aliResource(rocketMqMod, "ClientUser")},
			"alicloud_sag_dnat_entry":  {Tok: aliResource(rocketMqMod, "DnatEntry")},
			"alicloud_sag_qos":         {Tok: aliResource(rocketMqMod, "Qos")},
			"alicloud_sag_qos_car":     {Tok: aliResource(rocketMqMod, "QosCar")},
			"alicloud_sag_qos_policy":  {Tok: aliResource(rocketMqMod, "QosPolicy")},
			"alicloud_sag_snat_entry":  {Tok: aliResource(rocketMqMod, "SnatEntry")},
			// All above need remapped :/
			"alicloud_smartag_flow_log": {Tok: aliResource(sagMod, "SmartagFlowLog")},

			// ServiceMesh
			"alicloud_service_mesh_service_mesh":    {Tok: aliResource(serviceMeshMod, "ServiceMesh")},
			"alicloud_service_mesh_user_permission": {Tok: aliResource(serviceMeshMod, "UserPermission")},

			// Sas
			"alicloud_simple_application_server_instance":      {Tok: aliResource(sasMod, "Instance")},
			"alicloud_simple_application_server_custom_image":  {Tok: aliResource(sasMod, "CustomImage")},
			"alicloud_simple_application_server_firewall_rule": {Tok: aliResource(sasMod, "FirewallRule")},
			"alicloud_simple_application_server_snapshot":      {Tok: aliResource(sasMod, "Snapshot")},

			// Scdn
			"alicloud_scdn_domain_config": {Tok: aliResource(scdnMod, "DomainConfig")},
			"alicloud_scdn_domain":        {Tok: aliResource(scdnMod, "Domain")},

			// SchedulerX
			"alicloud_schedulerx_namespace": {Tok: aliResource(schedulerXMod, "Namespace")},

			// Sddp
			"alicloud_sddp_rule":       {Tok: aliResource(sddpMod, "Rule")},
			"alicloud_sddp_config":     {Tok: aliResource(sddpMod, "Config")},
			"alicloud_sddp_instance":   {Tok: aliResource(sddpMod, "Instance")},
			"alicloud_sddp_data_limit": {Tok: aliResource(sddpMod, "DataLimit")},

			// Security Center
			"alicloud_security_center_group":               {Tok: aliResource(securityCenterMod, "Group")},
			"alicloud_security_center_service_linked_role": {Tok: aliResource(securityCenterMod, "ServiceLinkedRole")},

			// Slb
			"alicloud_slb": {
				Tok:                aliResource(slbMod, "LoadBalancer"),
				DeprecationMessage: "This resource has been deprecated in favour of the ApplicationLoadBalancer resource",
			},
			"alicloud_slb_acl":            {Tok: aliResource(slbMod, "Acl")},
			"alicloud_slb_attachment":     {Tok: aliResource(slbMod, "Attachment")},
			"alicloud_slb_listener":       {Tok: aliResource(slbMod, "Listener")},
			"alicloud_slb_rule":           {Tok: aliResource(slbMod, "Rule")},
			"alicloud_slb_server_group":   {Tok: aliResource(slbMod, "ServerGroup")},
			"alicloud_slb_backend_server": {Tok: aliResource(slbMod, "BackendServer")},
			"alicloud_slb_ca_certificate": {
				Tok: aliResource(slbMod, "CaCertificate"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ca_certificate": {
						CSharpName: "Certificate",
					},
				},
			},
			"alicloud_slb_domain_extension":          {Tok: aliResource(slbMod, "DomainExtension")},
			"alicloud_slb_master_slave_server_group": {Tok: aliResource(slbMod, "MasterSlaveServerGroup")},
			"alicloud_slb_server_certificate": {
				Tok: aliResource(slbMod, "ServerCertificate"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"server_certificate": {
						CSharpName: "Certificate",
					},
				},
			},
			"alicloud_slb_load_balancer":                  {Tok: aliResource(slbMod, "ApplicationLoadBalancer")},
			"alicloud_slb_tls_cipher_policy":              {Tok: aliResource(slbMod, "TlsCipherPolicy")},
			"alicloud_slb_acl_entry_attachment":           {Tok: aliResource(slbMod, "AclEntryAttachment")},
			"alicloud_slb_server_group_server_attachment": {Tok: aliResource(slbMod, "ServerGroupServerAttachment")},

			// sms
			"alicloud_sms_short_url": {Tok: aliResource(smsMod, "ShortUrl")},

			// Tsdb
			"alicloud_tsdb_instance": {Tok: aliResource(tsdbMod, "Instance")},

			// VPC
			"alicloud_subnet": {
				Tok: aliResource(vpcMod, "Subnet"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
				DeprecationMessage: "This resource has been deprecated and replaced by the Switch resource.",
			},
			"alicloud_vpc":                          {Tok: aliResource(vpcMod, "Network")},
			"alicloud_vswitch":                      {Tok: aliResource(vpcMod, "Switch")},
			"alicloud_network_acl":                  {Tok: aliResource(vpcMod, "NetworkAcl")},
			"alicloud_network_acl_attachment":       {Tok: aliResource(vpcMod, "NetworkAclAttachment")},
			"alicloud_network_acl_entries":          {Tok: aliResource(vpcMod, "NetworkAclEntries")},
			"alicloud_network_interface":            {Tok: aliResource(vpcMod, "NetworkInterface")},
			"alicloud_network_interface_attachment": {Tok: aliResource(vpcMod, "NetworkInterfaceAttachment")},
			"alicloud_snat_entry": {
				Tok: aliResource(vpcMod, "SnatEntry"),
				Docs: &tfbridge.DocInfo{
					Source: "snat.html.markdown",
				},
			},
			"alicloud_route_entry":                         {Tok: aliResource(vpcMod, "RouteEntry")},
			"alicloud_route_table":                         {Tok: aliResource(vpcMod, "RouteTable")},
			"alicloud_route_table_attachment":              {Tok: aliResource(vpcMod, "RouteTableAttachment")},
			"alicloud_router_interface":                    {Tok: aliResource(vpcMod, "RouterInterface")},
			"alicloud_router_interface_connection":         {Tok: aliResource(vpcMod, "RouterInterfaceConnection")},
			"alicloud_nat_gateway":                         {Tok: aliResource(vpcMod, "NatGateway")},
			"alicloud_forward_entry":                       {Tok: aliResource(vpcMod, "ForwardEntry")},
			"alicloud_havip":                               {Tok: aliResource(vpcMod, "HAVip")},
			"alicloud_havip_attachment":                    {Tok: aliResource(vpcMod, "HAVipAttachment")},
			"alicloud_common_bandwidth_package":            {Tok: aliResource(vpcMod, "CommonBandwithPackage")},
			"alicloud_common_bandwidth_package_attachment": {Tok: aliResource(vpcMod, "CommonBandwithPackageAttachment")},
			"alicloud_vpc_flow_log":                        {Tok: aliResource(vpcMod, "FlowLog")},
			"alicloud_vpc_dhcp_options_set":                {Tok: aliResource(vpcMod, "DhcpOptionsSet")},
			"alicloud_vpc_nat_ip_cidr": {
				Tok: aliResource(vpcMod, "NatIpCidr"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"nat_ip_cidr": {
						CSharpName: "NatIpCidrBlock",
					},
				},
			},
			"alicloud_vpc_nat_ip": {
				Tok: aliResource(vpcMod, "NatIp"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"nat_ip": {
						CSharpName: "NatIpAddress",
					},
				},
			},
			"alicloud_vpc_traffic_mirror_filter":              {Tok: aliResource(vpcMod, "TrafficMirrorFilter")},
			"alicloud_vpc_traffic_mirror_filter_egress_rule":  {Tok: aliResource(vpcMod, "TrafficMirrorFilterEgressRule")},
			"alicloud_vpc_traffic_mirror_filter_ingress_rule": {Tok: aliResource(vpcMod, "TrafficMirrorFilterIngressRule")},
			"alicloud_vpc_traffic_mirror_session":             {Tok: aliResource(vpcMod, "TrafficMirrorSession")},
			"alicloud_vpc_ipv6_egress_rule":                   {Tok: aliResource(vpcMod, "Ipv6EgressRule")},
			"alicloud_vpc_ipv6_gateway":                       {Tok: aliResource(vpcMod, "Ipv6Gateway")},
			"alicloud_vpc_ipv6_internet_bandwidth":            {Tok: aliResource(vpcMod, "Ipv6InternetBandwidth")},
			"alicloud_vpc_vbr_ha":                             {Tok: aliResource(vpcMod, "VbrHa")},
			"alicloud_vpc_dhcp_options_set_attachment":        {Tok: aliResource(vpcMod, "DhcpOptionsSetAttachment")},
			"alicloud_vpc_bgp_peer":                           {Tok: aliResource(vpcMod, "BgpPeer")},
			"alicloud_vpc_bgp_network":                        {Tok: aliResource(vpcMod, "BgpNetwork")},
			"alicloud_vpc_bgp_group":                          {Tok: aliResource(vpcMod, "BgpGroup")},
			"alicloud_vpc_ipv4_gateway":                       {Tok: aliResource(vpcMod, "Ipv4Gateway")},
			"alicloud_vpc_prefix_list":                        {Tok: aliResource(vpcMod, "PrefixList")},

			// Vod
			"alicloud_vod_domain": {Tok: aliResource(vodMod, "Domain")},

			// VPN
			"alicloud_ssl_vpn_client_cert":        {Tok: aliResource(vpnMod, "SslVpnClientCert")},
			"alicloud_ssl_vpn_server":             {Tok: aliResource(vpnMod, "SslVpnServer")},
			"alicloud_vpn_connection":             {Tok: aliResource(vpnMod, "Connection")},
			"alicloud_vpn_customer_gateway":       {Tok: aliResource(vpnMod, "CustomerGateway")},
			"alicloud_vpn_gateway":                {Tok: aliResource(vpnMod, "Gateway")},
			"alicloud_vpn_route_entry":            {Tok: aliResource(vpnMod, "RouteEntry")},
			"alicloud_vpn_ipsec_server":           {Tok: aliResource(vpnMod, "IpsecServer")},
			"alicloud_vpn_pbr_route_entry":        {Tok: aliResource(vpnMod, "PbrRouteEntry")},
			"alicloud_vpn_gateway_vpn_attachment": {Tok: aliResource(vpnMod, "GatewayVpnAttachment")},

			// Vs
			"alicloud_video_surveillance_system_group": {Tok: aliResource(vsMod, "SystemGroup")},

			// Waf
			"alicloud_waf_domain": {
				Tok: aliResource(wafMod, "Domain"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"domain": {
						CSharpName: "DomainDeprecated",
					},
				},
			},
			"alicloud_waf_instance":          {Tok: aliResource(wafMod, "Instance")},
			"alicloud_waf_protection_module": {Tok: aliResource(wafMod, "ProtectionModule")},
			"alicloud_waf_certificate": {
				Tok: aliResource(wafMod, "Certificate"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"certificate": {
						CSharpName: "CertificateContents",
					},
				},
			},

			// Yundun
			"alicloud_yundun_bastionhost_instance": {
				Tok: aliResource(yundunMod, "BastionHostInstance"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_yundun_dbaudit_instance": {Tok: aliResource(yundunMod, "DBAuditInstance")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"alicloud_account":                              {Tok: dataSource(alicloudMod, "getAccount")},
			"alicloud_zones":                                {Tok: dataSource(alicloudMod, "getZones")},
			"alicloud_regions":                              {Tok: dataSource(alicloudMod, "getRegions")},
			"alicloud_caller_identity":                      {Tok: dataSource(alicloudMod, "getCallerIdentity")},
			"alicloud_file_crc64_checksum":                  {Tok: dataSource(alicloudMod, "getFileCrc64Checksum")},
			"alicloud_msc_sub_contacts":                     {Tok: dataSource(alicloudMod, "getMscSubContacts")},
			"alicloud_msc_sub_subscriptions":                {Tok: dataSource(alicloudMod, "getMscSubSubscriptions")},
			"alicloud_msc_sub_webhooks":                     {Tok: dataSource(alicloudMod, "getMscSubWebhooks")},
			"alicloud_msc_sub_contact_verification_message": {Tok: dataSource(alicloudMod, "getMscSubContactVerificationMessage")},

			// ActionTrail
			"alicloud_actiontrails": {
				Tok:                dataSource(actionTrailMod, "getTrailsDeprecated"),
				DeprecationMessage: "DataSource has been renamed to `getTrails`",
			},
			"alicloud_actiontrail_trails":                {Tok: dataSource(actionTrailMod, "getTrails")},
			"alicloud_actiontrail_history_delivery_jobs": {Tok: dataSource(actionTrailMod, "getHistoryDeliveryJobs")},

			// AliKafka
			"alicloud_alikafka_consumer_groups": {Tok: dataSource(actionTrailMod, "getConsumerGroups")},
			"alicloud_alikafka_instances":       {Tok: dataSource(actionTrailMod, "getInstances")},
			"alicloud_alikafka_sasl_acls":       {Tok: dataSource(actionTrailMod, "getSaslAcls")},
			"alicloud_alikafka_sasl_users":      {Tok: dataSource(actionTrailMod, "getSaslUsers")},
			"alicloud_alikafka_topics":          {Tok: dataSource(actionTrailMod, "getTopics")},

			// Alb
			"alicloud_alb_security_policies":      {Tok: dataSource(albMod, "getSecurityPolicies")},
			"alicloud_alb_server_groups":          {Tok: dataSource(albMod, "getServerGroups")},
			"alicloud_alb_load_balancers":         {Tok: dataSource(albMod, "getLoadBalancers")},
			"alicloud_alb_zones":                  {Tok: dataSource(albMod, "getZones")},
			"alicloud_alb_rules":                  {Tok: dataSource(albMod, "getRules")},
			"alicloud_alb_acls":                   {Tok: dataSource(albMod, "getAcls")},
			"alicloud_alb_listeners":              {Tok: dataSource(albMod, "getListeners")},
			"alicloud_alb_health_check_templates": {Tok: dataSource(albMod, "getHealthCheckTemplates")},

			// amqp
			"alicloud_amqp_virtual_hosts": {Tok: dataSource(amqpMod, "getVirtualHosts")},
			"alicloud_amqp_queues":        {Tok: dataSource(amqpMod, "getQueues")},
			"alicloud_amqp_instances":     {Tok: dataSource(amqpMod, "getInstances")},
			"alicloud_amqp_exchanges":     {Tok: dataSource(amqpMod, "getExchanges")},
			"alicloud_amqp_bindings":      {Tok: dataSource(amqpMod, "getBindings")},

			// ApiGateway
			"alicloud_api_gateway_groups":   {Tok: dataSource(apiGatewayMod, "getGroups")},
			"alicloud_api_gateway_apis":     {Tok: dataSource(apiGatewayMod, "getApis")},
			"alicloud_api_gateway_apps":     {Tok: dataSource(apiGatewayMod, "getApps")},
			"alicloud_api_gateway_service":  {Tok: dataSource(apiGatewayMod, "getService")},
			"alicloud_api_gateway_backends": {Tok: dataSource(apiGatewayMod, "getBackends")},

			// Arms
			"alicloud_arms_alert_contacts":         {Tok: dataSource(armsMod, "getAlertContacts")},
			"alicloud_arms_alert_contact_groups":   {Tok: dataSource(armsMod, "getAlertContactGroups")},
			"alicloud_arms_dispatch_rules":         {Tok: dataSource(armsMod, "getDispatchRules")},
			"alicloud_arms_prometheus_alert_rules": {Tok: dataSource(armsMod, "getPrometheusAlertRules")},

			// BastionHost
			"alicloud_bastionhost_user_groups":     {Tok: dataSource(bastionHostMod, "getUserGroups")},
			"alicloud_bastionhost_instances":       {Tok: dataSource(bastionHostMod, "getInstances")},
			"alicloud_bastionhost_users":           {Tok: dataSource(bastionHostMod, "getUsers")},
			"alicloud_bastionhost_host_groups":     {Tok: dataSource(bastionHostMod, "getHostGroups")},
			"alicloud_bastionhost_host_accounts":   {Tok: dataSource(bastionHostMod, "getHostAccounts")},
			"alicloud_bastionhost_hosts":           {Tok: dataSource(bastionHostMod, "getHosts")},
			"alicloud_bastionhost_host_share_keys": {Tok: dataSource(bastionHostMod, "getHostShareKeys")},

			// Brain
			"alicloud_brain_industrial_pid_projects":      {Tok: dataSource(brainMod, "getIndustrialPidProjects")},
			"alicloud_brain_industrial_pid_organizations": {Tok: dataSource(brainMod, "getIndustrialPidOrganizations")},
			"alicloud_brain_industrial_service":           {Tok: dataSource(brainMod, "getIndustrialSerice")},
			"alicloud_brain_industrial_pid_loops":         {Tok: dataSource(brainMod, "getIndustrialPidLoops")},

			// Cas
			"alicloud_cas_certificates": {
				Tok:                dataSource(casMod, "getCertificates"),
				DeprecationMessage: "This resource has been deprecated in favour of getServiceCertificates",
			},
			"alicloud_ssl_certificates_service_certificates": {Tok: dataSource(casMod, "getServiceCertificates")},

			// Cassandra
			"alicloud_cassandra_clusters":     {Tok: dataSource(cassandraMod, "getClusters")},
			"alicloud_cassandra_data_centers": {Tok: dataSource(cassandraMod, "getDataCenters")},
			"alicloud_cassandra_zones":        {Tok: dataSource(cassandraMod, "getZones")},
			"alicloud_cassandra_backup_plans": {Tok: dataSource(cassandraMod, "getBackupPlans")},

			// Cddc
			"alicloud_cddc_dedicated_host_groups":   {Tok: dataSource(cddcMod, "getDedicatedHostGroups")},
			"alicloud_cddc_dedicated_hosts":         {Tok: dataSource(cddcMod, "getDedicatedHosts")},
			"alicloud_cddc_zones":                   {Tok: dataSource(cddcMod, "getZones")},
			"alicloud_cddc_host_ecs_level_infos":    {Tok: dataSource(cddcMod, "getHostEcsLevelInfos")},
			"alicloud_cddc_dedicated_host_accounts": {Tok: dataSource(cddcMod, "getDedicatedHostAccounts")},

			// Cdn
			"alicloud_cdn_service":                  {Tok: dataSource(cdnMod, "getService")},
			"alicloud_cdn_real_time_log_deliveries": {Tok: dataSource(cdnMod, "getRealTimeLogDeliveries")},
			"alicloud_cdn_ip_info":                  {Tok: dataSource(cdnMod, "getIpInfo")},
			"alicloud_cdn_blocked_regions":          {Tok: dataSource(cdnMod, "getBlockedRegions")},

			// Cen
			"alicloud_cen_bandwidth_limits":            {Tok: dataSource(cenMod, "getBandwidthLimits")},
			"alicloud_cen_bandwidth_packages":          {Tok: dataSource(cenMod, "getBandwidthPackages")},
			"alicloud_cen_instances":                   {Tok: dataSource(cenMod, "getInstances")},
			"alicloud_cen_route_entries":               {Tok: dataSource(cenMod, "getRouteEntries")},
			"alicloud_cen_region_route_entries":        {Tok: dataSource(cenMod, "getRegionRouteEntries")},
			"alicloud_cen_flowlogs":                    {Tok: dataSource(cenMod, "getFlowlogs")},
			"alicloud_cen_route_maps":                  {Tok: dataSource(cenMod, "getRouteMaps")},
			"alicloud_cen_private_zones":               {Tok: dataSource(cenMod, "getPrivateZones")},
			"alicloud_cen_instance_attachments":        {Tok: dataSource(cenMod, "getInstanceAttachments")},
			"alicloud_cen_vbr_health_checks":           {Tok: dataSource(cenMod, "getVbrHealthChecks")},
			"alicloud_cen_route_services":              {Tok: dataSource(cenMod, "getRouteServices")},
			"alicloud_cen_transit_routers":             {Tok: dataSource(cenMod, "getTransitRouters")},
			"alicloud_cen_transit_router_route_tables": {Tok: dataSource(cenMod, "getTransitRouterRouteTables")},
			"alicloud_cen_transit_router_route_table_associations": {
				Tok: dataSource(cenMod, "getTransitRouterRouteTableAssociations"),
			},
			"alicloud_cen_transit_router_route_table_propagations": {
				Tok: dataSource(cenMod, "getTransitRouterRouteTablePropagations"),
			},
			"alicloud_cen_transit_router_route_entries":       {Tok: dataSource(cenMod, "getTransitRouterRouteEntries")},
			"alicloud_cen_transit_router_vbr_attachments":     {Tok: dataSource(cenMod, "getTransitRouterVbrAttachments")},
			"alicloud_cen_transit_router_vpc_attachments":     {Tok: dataSource(cenMod, "getTransitRouterVpcAttachments")},
			"alicloud_cen_transit_router_peer_attachments":    {Tok: dataSource(cenMod, "getTransitRouterPeerAttachments")},
			"alicloud_cen_transit_router_service":             {Tok: dataSource(cenMod, "getTransitRouterService")},
			"alicloud_cen_transit_router_available_resources": {Tok: dataSource(cenMod, "getTransitRouterAvailableResources")},
			"alicloud_cen_traffic_marking_policies":           {Tok: dataSource(cenMod, "getTrafficMarkingPolicies")},

			// Clickhouse
			"alicloud_click_house_db_clusters":     {Tok: dataSource(clickHouseMod, "getDbClusters")},
			"alicloud_click_house_accounts":        {Tok: dataSource(clickHouseMod, "getAccounts")},
			"alicloud_click_house_regions":         {Tok: dataSource(clickHouseMod, "getRegions")},
			"alicloud_click_house_backup_policies": {Tok: dataSource(clickHouseMod, "getBackupPolicies")},

			// CloudAuth
			"alicloud_cloudauth_face_configs": {Tok: dataSource(cloudAuthMod, "getFaceConfigs")},

			// CloudConnect
			"alicloud_cloud_connect_networks": {Tok: dataSource(cloudConnectMod, "getNetworks")},

			// Cloudfirewall
			"alicloud_cloud_firewall_control_policies": {Tok: dataSource(cloudFirewallMod, "getControlPolicies")},
			"alicloud_cloud_firewall_instances":        {Tok: dataSource(cloudFirewallMod, "getInstances")},
			"alicloud_cloud_firewall_address_books":    {Tok: dataSource(cloudFirewallMod, "getAddressBooks")},

			// CloudStorageGateway
			"alicloud_cloud_storage_gateway_storage_bundles": {
				Tok: dataSource(cloudStorageGatewayMod, "getStorageBundles"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_cloud_storage_gateway_service":               {Tok: dataSource(cloudStorageGatewayMod, "getService")},
			"alicloud_cloud_storage_gateway_gateways":              {Tok: dataSource(cloudStorageGatewayMod, "getGateways")},
			"alicloud_cloud_storage_gateway_express_syncs":         {Tok: dataSource(cloudStorageGatewayMod, "getExpressSyncs")},
			"alicloud_cloud_storage_gateway_gateway_block_volumes": {Tok: dataSource(cloudStorageGatewayMod, "getGatewayBlockVolumes")},
			"alicloud_cloud_storage_gateway_gateway_cache_disks":   {Tok: dataSource(cloudStorageGatewayMod, "getGatewayCacheDisks")},
			"alicloud_cloud_storage_gateway_gateway_file_shares":   {Tok: dataSource(cloudStorageGatewayMod, "getGatewayFileShares")},
			"alicloud_cloud_storage_gateway_gateway_smb_users":     {Tok: dataSource(cloudStorageGatewayMod, "getGatewaySmbUsers")},
			"alicloud_cloud_storage_gateway_stocks":                {Tok: dataSource(cloudStorageGatewayMod, "getStocks")},

			// CloudSSO
			"alicloud_cloud_sso_directories":             {Tok: dataSource(cloudSsoMod, "getDirectories")},
			"alicloud_cloud_sso_scim_server_credentials": {Tok: dataSource(cloudSsoMod, "getScimServerCredentials")},
			"alicloud_cloud_sso_groups":                  {Tok: dataSource(cloudSsoMod, "getGroups")},
			"alicloud_cloud_sso_access_configurations":   {Tok: dataSource(cloudSsoMod, "getAccessConfigurations")},
			"alicloud_cloud_sso_users":                   {Tok: dataSource(cloudSsoMod, "getUsers")},
			"alicloud_cloud_sso_service":                 {Tok: dataSource(cloudSsoMod, "getService")},

			// CMS
			"alicloud_cms_alarm_contacts":       {Tok: dataSource(cmsMod, "getAlarmContacts")},
			"alicloud_cms_alarm_contact_groups": {Tok: dataSource(cmsMod, "getAlarmContactGroups")},
			"alicloud_cms_group_metric_rules":   {Tok: dataSource(cmsMod, "getGroupMetricRules")},
			"alicloud_cms_service":              {Tok: dataSource(cmsMod, "getService")},
			"alicloud_cms_monitor_groups":       {Tok: dataSource(cmsMod, "getMonitorGroups")},
			"alicloud_cms_monitor_group_instanceses": {
				Tok: dataSource(cmsMod, "getMonitorGroupInstances"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_cms_metric_rule_templates":    {Tok: dataSource(cmsMod, "getMetricRuleTemplates")},
			"alicloud_cms_dynamic_tag_groups":       {Tok: dataSource(cmsMod, "getDynamicTagGroups")},
			"alicloud_cms_namespaces":               {Tok: dataSource(cmsMod, "getNamespaces")},
			"alicloud_cms_sls_groups":               {Tok: dataSource(cmsMod, "getSlsGroups")},
			"alicloud_cms_hybrid_monitor_datas":     {Tok: dataSource(cmsMod, "getHybridMonitorDatas")},
			"alicloud_cms_hybrid_monitor_fc_tasks":  {Tok: dataSource(cmsMod, "getHybridMonitorFcTasks")},
			"alicloud_cms_hybrid_monitor_sls_tasks": {Tok: dataSource(cmsMod, "getHybridMonitorSlsTasks")},
			"alicloud_cms_event_rules":              {Tok: dataSource(cmsMod, "getEventRules")},

			// Config
			"alicloud_config_configuration_recorders":    {Tok: dataSource(cfgMod, "getConfigurationRecorders")},
			"alicloud_config_delivery_channels":          {Tok: dataSource(cfgMod, "getDeliveryChannels")},
			"alicloud_config_rules":                      {Tok: dataSource(cfgMod, "getRules")},
			"alicloud_config_aggregate_compliance_packs": {Tok: dataSource(cfgMod, "getAggregateCompliancePacks")},
			"alicloud_config_aggregate_config_rules":     {Tok: dataSource(cfgMod, "getAggregateConfigRules")},
			"alicloud_config_aggregators":                {Tok: dataSource(cfgMod, "getAggregators")},
			"alicloud_config_compliance_packs":           {Tok: dataSource(cfgMod, "getCompliancePacks")},
			"alicloud_config_deliveries":                 {Tok: dataSource(cfgMod, "getDeliveries")},
			"alicloud_config_aggregate_deliveries":       {Tok: dataSource(cfgMod, "getAggregateDeliveries")},

			// Cr
			"alicloud_cr_namespaces":            {Tok: dataSource(crMod, "getNamespaces")},
			"alicloud_cr_repos":                 {Tok: dataSource(crMod, "getRepos")},
			"alicloud_cr_service":               {Tok: dataSource(crMod, "getService")},
			"alicloud_cr_endpoint_acl_service":  {Tok: dataSource(crMod, "getEndpointAclService")},
			"alicloud_cr_endpoint_acl_policies": {Tok: dataSource(crMod, "getEndpointAclPolicies")},
			"alicloud_cr_chart_repositories":    {Tok: dataSource(crMod, "getChartRepositories")},
			"alicloud_cr_chart_namespaces":      {Tok: dataSource(crMod, "getChartNamespaces")},
			"alicloud_cr_chains":                {Tok: dataSource(crMod, "getChains")},

			// Cs
			"alicloud_cs_kubernetes_clusters":            {Tok: dataSource(csMod, "getKubernetesClusters")},
			"alicloud_cs_edge_kubernetes_clusters":       {Tok: dataSource(csMod, "getEdgeKubernetesClusters")},
			"alicloud_cs_managed_kubernetes_clusters":    {Tok: dataSource(csMod, "getManagedKubernetesClusters")},
			"alicloud_cs_serverless_kubernetes_clusters": {Tok: dataSource(csMod, "getServerlessKubernetesClusters")},
			"alicloud_cr_ee_instances":                   {Tok: dataSource(csMod, "getRegistryEnterpriseInstances")},
			"alicloud_cr_ee_namespaces":                  {Tok: dataSource(csMod, "getRegistryEnterpriseNamespaces")},
			"alicloud_cr_ee_repos":                       {Tok: dataSource(csMod, "getRegistryEnterpriseRepos")},
			"alicloud_cr_ee_sync_rules":                  {Tok: dataSource(csMod, "getRegistryEnterpriseSyncRules")},
			"alicloud_ack_service":                       {Tok: dataSource(csMod, "getAckService")},
			"alicloud_cs_kubernetes_permissions":         {Tok: dataSource(csMod, "getKubernetesPermission")},
			"alicloud_cs_kubernetes_addons":              {Tok: dataSource(csMod, "getKubernetesAddons")},
			"alicloud_cs_kubernetes_addon_metadata":      {Tok: dataSource(csMod, "getKubernetesAddonMetadata")},
			"alicloud_cs_kubernetes_version":             {Tok: dataSource(csMod, "getKubernetesVersion")},

			// Database Filesystem
			"alicloud_dbfs_instances": {Tok: dataSource(databaseFilesystemMod, "getInstances")},
			"alicloud_dbfs_snapshots": {Tok: dataSource(databaseFilesystemMod, "getSnapshots")},

			// Database Gateway
			"alicloud_database_gateway_gateways": {Tok: dataSource(databaseGatewayMod, "getGateways")},

			// Datahub
			"alicloud_datahub_service": {Tok: dataSource(datahubMod, "getService")},

			// DataWorks
			"alicloud_data_works_folders": {Tok: dataSource(dataWorksMod, "getFolders")},
			"alicloud_data_works_service": {Tok: dataSource(dataWorksMod, "getService")},

			// Dcdn
			"alicloud_dcdn_domains":     {Tok: dataSource(dcdnMod, "getDomains")},
			"alicloud_dcdn_service":     {Tok: dataSource(dcdnMod, "getService")},
			"alicloud_dcdn_ipa_domains": {Tok: dataSource(dcdnMod, "getIpaDomains")},

			// Dds
			"alicloud_mongo_instances": {
				Tok: dataSource(ddsMod, "getMongoInstances"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},

			// Ddos
			"alicloud_ddoscoo_instances":        {Tok: dataSource(ddosMod, "getDdosCooInstances")},
			"alicloud_ddosbgp_instances":        {Tok: dataSource(ddosMod, "getDdosBgpInstances")},
			"alicloud_ddoscoo_domain_resources": {Tok: dataSource(ddosMod, "getDdosCooDomainResources")},
			"alicloud_ddoscoo_ports":            {Tok: dataSource(ddosMod, "getDdosCooPorts")},
			"alicloud_ddosbgp_ips":              {Tok: dataSource(ddosMod, "getDdosBgpIps")},

			// Dfs
			"alicloud_dfs_access_groups": {Tok: dataSource(dfsMod, "getAccessGroups")},
			"alicloud_dfs_zones":         {Tok: dataSource(dfsMod, "getZones")},
			"alicloud_dfs_file_systems":  {Tok: dataSource(dfsMod, "getFileSystems")},
			"alicloud_dfs_access_rules":  {Tok: dataSource(dfsMod, "getAccessRules")},
			"alicloud_dfs_mount_points":  {Tok: dataSource(dfsMod, "getMountPoints")},

			// Direct Mail
			"alicloud_direct_mail_receiverses":    {Tok: dataSource(directMailMod, "getReceivers")},
			"alicloud_direct_mail_mail_addresses": {Tok: dataSource(directMailMod, "getMailAddresses")},
			"alicloud_direct_mail_domains":        {Tok: dataSource(directMailMod, "getDomains")},
			"alicloud_direct_mail_tags":           {Tok: dataSource(directMailMod, "getTags")},

			// Dms
			"alicloud_dms_enterprise_instances": {Tok: dataSource(dmsMod, "getEnterpriseInstances")},
			"alicloud_dms_enterprise_users":     {Tok: dataSource(dmsMod, "getEnterpriseUsers")},
			"alicloud_dms_user_tenants":         {Tok: dataSource(dmsMod, "getUserTenants")},

			// Dns
			"alicloud_dns_domain_groups":        {Tok: dataSource(dnsMod, "getDomainGroups")},
			"alicloud_dns_domain_records":       {Tok: dataSource(dnsMod, "getDomainRecords")},
			"alicloud_dns_domains":              {Tok: dataSource(dnsMod, "getDomains")},
			"alicloud_dns_groups":               {Tok: dataSource(dnsMod, "getGroups")},
			"alicloud_dns_records":              {Tok: dataSource(dnsMod, "getRecords")},
			"alicloud_dns_resolution_lines":     {Tok: dataSource(dnsMod, "getResolutionLines")},
			"alicloud_dns_domain_txt_guid":      {Tok: dataSource(dnsMod, "getDomainTxtGuid")},
			"alicloud_dns_instances":            {Tok: dataSource(dnsMod, "getInstances")},
			"alicloud_alidns_domain_groups":     {Tok: dataSource(dnsMod, "getAlidnsDomainGroups")},
			"alicloud_alidns_records":           {Tok: dataSource(dnsMod, "getAlidnsRecords")},
			"alicloud_alidns_domains":           {Tok: dataSource(dnsMod, "getAlidnsDomains")},
			"alicloud_alidns_instances":         {Tok: dataSource(dnsMod, "getAlidnsInstances")},
			"alicloud_alidns_gtm_instances":     {Tok: dataSource(dnsMod, "getGtmInstances")},
			"alicloud_alidns_custom_lines":      {Tok: dataSource(dnsMod, "getCustomLines")},
			"alicloud_alidns_address_pools":     {Tok: dataSource(dnsMod, "getAddressPools")},
			"alicloud_alidns_access_strategies": {Tok: dataSource(dnsMod, "getAccessStrategies")},

			// Drds
			"alicloud_drds_instances": {Tok: dataSource(drdsMod, "getInstances")},

			// Dts
			"alicloud_dts_subscription_jobs":    {Tok: dataSource(dtsMod, "getSubscriptionJobs")},
			"alicloud_dts_synchronization_jobs": {Tok: dataSource(dtsMod, "getSynchronizationJobs")},
			"alicloud_dts_consumer_channels":    {Tok: dataSource(dtsMod, "getConsumerChannels")},
			"alicloud_dts_migration_jobs":       {Tok: dataSource(dtsMod, "getMigrationJobs")},

			// Eais
			"alicloud_eais_instances": {Tok: dataSource(eaisMod, "getInstances")},

			// Eci
			"alicloud_eci_image_caches":     {Tok: dataSource(eciMod, "getImageCaches")},
			"alicloud_eci_container_groups": {Tok: dataSource(eciMod, "getContainerGroups")},
			"alicloud_eci_virtual_nodes":    {Tok: dataSource(eciMod, "getVirtualNodes")},
			"alicloud_eci_zones":            {Tok: dataSource(eciMod, "getZones")},

			// Ecp
			"alicloud_ecp_key_pairs":      {Tok: dataSource(ecpMod, "getKeyPairs")},
			"alicloud_ecp_instance_types": {Tok: dataSource(ecpMod, "getInstanceTypes")},
			"alicloud_ecp_instances":      {Tok: dataSource(ecpMod, "getInstances")},
			"alicloud_ecp_zones":          {Tok: dataSource(ecpMod, "getZones")},

			// Ecs
			"alicloud_images":               {Tok: dataSource(ecsMod, "getImages")},
			"alicloud_instance_types":       {Tok: dataSource(ecsMod, "getInstanceTypes")},
			"alicloud_instances":            {Tok: dataSource(ecsMod, "getInstances")},
			"alicloud_key_pairs":            {Tok: dataSource(ecsMod, "getKeyPairs")},
			"alicloud_security_group_rules": {Tok: dataSource(ecsMod, "getSecurityGroupRules")},
			"alicloud_security_groups":      {Tok: dataSource(ecsMod, "getSecurityGroups")},
			"alicloud_disks":                {Tok: dataSource(ecsMod, "getDisks")},
			"alicloud_eips": {
				Tok:                dataSource(ecsMod, "getEips"),
				DeprecationMessage: "This function has been deprecated in favour of the getEipAddresses function",
			},
			"alicloud_eip_addresses":                     {Tok: dataSource(ecsMod, "getEipAddresses")},
			"alicloud_instance_type_families":            {Tok: dataSource(ecsMod, "getInstanceTypeFamilies")},
			"alicloud_network_interfaces":                {Tok: dataSource(ecsMod, "getNetworkInterfaces")},
			"alicloud_snapshots":                         {Tok: dataSource(ecsMod, "getSnapshots")},
			"alicloud_ecs_dedicated_hosts":               {Tok: dataSource(ecsMod, "getDedicatedHosts")},
			"alicloud_ecs_hpc_clusters":                  {Tok: dataSource(ecsMod, "getHpcClusters")},
			"alicloud_ecs_commands":                      {Tok: dataSource(ecsMod, "getCommands")},
			"alicloud_ecs_auto_snapshot_policies":        {Tok: dataSource(ecsMod, "getAutoSnapshotPolicies")},
			"alicloud_ecs_snapshots":                     {Tok: dataSource(ecsMod, "getEcsSnapshots")},
			"alicloud_ecs_launch_templates":              {Tok: dataSource(ecsMod, "getEcsLaunchTemplates")},
			"alicloud_ecs_key_pairs":                     {Tok: dataSource(ecsMod, "getEcsKeyPairs")},
			"alicloud_ecs_disks":                         {Tok: dataSource(ecsMod, "getEcsDisks")},
			"alicloud_ecs_network_interfaces":            {Tok: dataSource(ecsMod, "getEcsNetworkInterfaces")},
			"alicloud_ecs_deployment_sets":               {Tok: dataSource(ecsMod, "getEcsDeploymentSets")},
			"alicloud_ecs_dedicated_host_clusters":       {Tok: dataSource(ecsMod, "getEcsDedicatedHostClusters")},
			"alicloud_ecs_prefix_lists":                  {Tok: dataSource(ecsMod, "getEcsPrefixLists")},
			"alicloud_ecs_image_components":              {Tok: dataSource(ecsMod, "getEcsImageComponents")},
			"alicloud_ecs_snapshot_groups":               {Tok: dataSource(ecsMod, "getEcsSnapshotGroups")},
			"alicloud_ecs_storage_capacity_units":        {Tok: dataSource(ecsMod, "getEcsStorageCapacityUnits")},
			"alicloud_ecs_image_pipelines":               {Tok: dataSource(ecsMod, "getEcsImagePipeline")},
			"alicloud_ecs_invocations":                   {Tok: dataSource(ecsMod, "getEcsInvocations")},
			"alicloud_ecs_network_interface_permissions": {Tok: dataSource(ecsMod, "getEcsNetworkInterfacePermissions")},
			"alicloud_ecs_activations":                   {Tok: dataSource(ecsMod, "getActivations")},

			// Edas
			"alicloud_edas_applications":  {Tok: dataSource(edasMod, "getApplications")},
			"alicloud_edas_deploy_groups": {Tok: dataSource(edasMod, "getDeployGroups")},
			"alicloud_edas_clusters":      {Tok: dataSource(edasMod, "getClusters")},
			"alicloud_edas_service":       {Tok: dataSource(edasMod, "getService")},
			"alicloud_edas_namespaces":    {Tok: dataSource(edasMod, "getNamespaces")},

			// Eds
			"alicloud_ecd_policy_groups":             {Tok: dataSource(edsMod, "getPolicyGroups")},
			"alicloud_ecd_simple_office_sites":       {Tok: dataSource(edsMod, "getSimpleOfficeSites")},
			"alicloud_ecd_nas_file_systems":          {Tok: dataSource(edsMod, "getNasFileSystems")},
			"alicloud_ecd_bundles":                   {Tok: dataSource(edsMod, "getBundles")},
			"alicloud_ecd_desktops":                  {Tok: dataSource(edsMod, "getDesktops")},
			"alicloud_ecd_network_packages":          {Tok: dataSource(edsMod, "getNetworkPackages")},
			"alicloud_ecd_users":                     {Tok: dataSource(edsMod, "getUsers")},
			"alicloud_ecd_images":                    {Tok: dataSource(edsMod, "getImages")},
			"alicloud_ecd_commands":                  {Tok: dataSource(edsMod, "getCommands")},
			"alicloud_ecd_snapshots":                 {Tok: dataSource(edsMod, "getSnapshots")},
			"alicloud_ecd_desktop_types":             {Tok: dataSource(edsMod, "getDesktopTypes")},
			"alicloud_ecd_ad_connector_directories":  {Tok: dataSource(edsMod, "getAdConnectorDirectories")},
			"alicloud_ecd_ram_directories":           {Tok: dataSource(edsMod, "getRamDirectories")},
			"alicloud_ecd_zones":                     {Tok: dataSource(edsMod, "getZones")},
			"alicloud_ecd_ad_connector_office_sites": {Tok: dataSource(edsMod, "getAdConnectorOfficeSites")},
			"alicloud_ecd_custom_properties":         {Tok: dataSource(edsMod, "getCustomProperties")},

			// Ehpc
			"alicloud_ehpc_job_templates": {Tok: dataSource(ehpcMod, "getJobTemplates")},
			"alicloud_ehpc_clusters":      {Tok: dataSource(ehpcMod, "getClusters")},

			// EipAnycast
			"alicloud_eipanycast_anycast_eip_addresses": {Tok: dataSource(eipAnyCastMod, "getAnycastEipAddresses")},

			// Elasticsearch
			"alicloud_elasticsearch_instances": {
				Tok: dataSource(elasticsearchMod, "getInstances"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_elasticsearch_zones": {Tok: dataSource(elasticsearchMod, "getZones")},

			// Emr
			"alicloud_emr_disk_types":     {Tok: dataSource(emrMod, "getDiskTypes")},
			"alicloud_emr_instance_types": {Tok: dataSource(emrMod, "getInstanceTypes")},
			"alicloud_emr_main_versions":  {Tok: dataSource(emrMod, "getMainVersions")},
			"alicloud_emr_clusters":       {Tok: dataSource(emrMod, "getClusters")},

			// Ens
			"alicloud_ens_key_pairs": {Tok: dataSource(ensMod, "getKeyPairs")},

			// Ess
			"alicloud_ess_scaling_configurations": {Tok: dataSource(essMod, "getScalingConfigurations")},
			"alicloud_ess_scaling_groups":         {Tok: dataSource(essMod, "getScalingGroups")},
			"alicloud_ess_scaling_rules":          {Tok: dataSource(essMod, "getScalingRules")},
			"alicloud_ess_alarms":                 {Tok: dataSource(essMod, "getAlarms")},
			"alicloud_ess_notifications":          {Tok: dataSource(essMod, "getNotifications")},
			"alicloud_ess_scheduled_tasks":        {Tok: dataSource(essMod, "getScheduledTasks")},
			"alicloud_ess_lifecycle_hooks":        {Tok: dataSource(essMod, "getLifecycleHooks")},

			// event bridge
			"alicloud_event_bridge_service":       {Tok: dataSource(eventBridgeMod, "getService")},
			"alicloud_event_bridge_event_buses":   {Tok: dataSource(eventBridgeMod, "getEventBuses")},
			"alicloud_event_bridge_rules":         {Tok: dataSource(eventBridgeMod, "getRules")},
			"alicloud_event_bridge_event_sources": {Tok: dataSource(eventBridgeMod, "getEventSources")},

			// Express Connect
			"alicloud_express_connect_access_points": {Tok: dataSource(expressConnect, "getAccessPoints")},
			"alicloud_express_connect_physical_connection_service": {
				Tok: dataSource(expressConnect, "getPhysicalConnectionService"),
			},
			"alicloud_express_connect_physical_connections":   {Tok: dataSource(expressConnect, "getPhysicalConnections")},
			"alicloud_express_connect_virtual_border_routers": {Tok: dataSource(expressConnect, "getVirtualBorderRouters")},

			// Fc
			"alicloud_fc_functions":      {Tok: dataSource(fcMod, "getFunctions")},
			"alicloud_fc_services":       {Tok: dataSource(fcMod, "getServices")},
			"alicloud_fc_triggers":       {Tok: dataSource(fcMod, "getTriggers")},
			"alicloud_fc_zones":          {Tok: dataSource(fcMod, "getZones")},
			"alicloud_fc_custom_domains": {Tok: dataSource(fcMod, "getCustomDomains")},
			"alicloud_fc_service":        {Tok: dataSource(fcMod, "getService")},

			// FNF
			"alicloud_fnf_schedules":  {Tok: dataSource(fnfMod, "getSchedules")},
			"alicloud_fnf_flows":      {Tok: dataSource(fnfMod, "getFlows")},
			"alicloud_fnf_service":    {Tok: dataSource(fnfMod, "getService")},
			"alicloud_fnf_executions": {Tok: dataSource(fnfMod, "getExecutions")},

			// Ga
			"alicloud_ga_listeners":                        {Tok: dataSource(gaMod, "getListeners")},
			"alicloud_ga_accelerators":                     {Tok: dataSource(gaMod, "getAccelerators")},
			"alicloud_ga_bandwidth_packages":               {Tok: dataSource(gaMod, "getBandwidthPackages")},
			"alicloud_ga_endpoint_groups":                  {Tok: dataSource(gaMod, "getEndpointGroups")},
			"alicloud_ga_ip_sets":                          {Tok: dataSource(gaMod, "getIpSets")},
			"alicloud_ga_forwarding_rules":                 {Tok: dataSource(gaMod, "getForwardingRules")},
			"alicloud_ga_additional_certificates":          {Tok: dataSource(gaMod, "getAdditionalCertificates")},
			"alicloud_ga_acls":                             {Tok: dataSource(gaMod, "getAcls")},
			"alicloud_ga_accelerator_spare_ip_attachments": {Tok: dataSource(gaMod, "getAcceleratorSpareIpAttachments")},

			// Gpdb
			"alicloud_gpdb_instances": {Tok: dataSource(gpdbMod, "getInstances")},
			"alicloud_gpdb_zones":     {Tok: dataSource(gpdbMod, "getZones")},
			"alicloud_gpdb_accounts":  {Tok: dataSource(gpdbMod, "getAccounts")},

			// Graph Database
			"alicloud_graph_database_db_instances": {Tok: dataSource(graphDatabaseMod, "getDbInstances")},

			// Hbase
			"alicloud_hbase_instances":      {Tok: dataSource(hbaseMod, "getInstances")},
			"alicloud_hbase_zones":          {Tok: dataSource(hbaseMod, "getZones")},
			"alicloud_hbase_instance_types": {Tok: dataSource(hbaseMod, "getInstanceTypes")},

			// hbr
			"alicloud_hbr_vaults":                    {Tok: dataSource(hbrMod, "getVaults")},
			"alicloud_hbr_oss_backup_plans":          {Tok: dataSource(hbrMod, "getOssBackupPlans")},
			"alicloud_hbr_ecs_backup_plans":          {Tok: dataSource(hbrMod, "getEcsBackupPlans")},
			"alicloud_hbr_nas_backup_plans":          {Tok: dataSource(hbrMod, "getNasBackupPlans")},
			"alicloud_hbr_ecs_backup_clients":        {Tok: dataSource(hbrMod, "getEcsBackupClients")},
			"alicloud_hbr_snapshots":                 {Tok: dataSource(hbrMod, "getSnapshots")},
			"alicloud_hbr_restore_jobs":              {Tok: dataSource(hbrMod, "getRestoreJobs")},
			"alicloud_hbr_backup_jobs":               {Tok: dataSource(hbrMod, "getBackupJobs")},
			"alicloud_hbr_server_backup_plans":       {Tok: dataSource(hbrMod, "getServerBackupPlans")},
			"alicloud_hbr_replication_vault_regions": {Tok: dataSource(hbrMod, "getReplicationVaultRegions")},
			"alicloud_hbr_ots_backup_plans":          {Tok: dataSource(hbrMod, "getOtsBackupPlans")},
			"alicloud_hbr_ots_snapshots":             {Tok: dataSource(hbrMod, "getOtsSnapshots")},
			"alicloud_hbr_hana_instances":            {Tok: dataSource(hbrMod, "getHanaInstances")},
			"alicloud_hbr_hana_backup_plans":         {Tok: dataSource(hbrMod, "getHanaBackupPlans")},

			// iot
			"alicloud_iot_service":       {Tok: dataSource(iotMod, "getService")},
			"alicloud_iot_device_groups": {Tok: dataSource(iotMod, "getDeviceGroups")},

			// Imm
			"alicloud_imm_projects": {Tok: dataSource(immMod, "getProjects")},

			// Imp
			"alicloud_imp_app_templates": {Tok: dataSource(impMod, "getAppTemplates")},

			// Kms
			"alicloud_kms_ciphertext":      {Tok: dataSource(kmsMod, "getCiphertext")},
			"alicloud_kms_plaintext":       {Tok: dataSource(kmsMod, "getPlaintext")},
			"alicloud_kms_keys":            {Tok: dataSource(kmsMod, "getKeys")},
			"alicloud_kms_aliases":         {Tok: dataSource(kmsMod, "getAliases")},
			"alicloud_kms_key_versions":    {Tok: dataSource(kmsMod, "getKeyVersions")},
			"alicloud_kms_secrets":         {Tok: dataSource(kmsMod, "getSecrets")},
			"alicloud_kms_secret_versions": {Tok: dataSource(kmsMod, "getSecretVersions")},
			"alicloud_kms_service":         {Tok: dataSource(kmsMod, "getService")},

			// KvStore
			"alicloud_kvstore_instances":        {Tok: dataSource(kvstoreMod, "getInstances")},
			"alicloud_kvstore_instance_classes": {Tok: dataSource(kvstoreMod, "getInstanceClasses")},
			"alicloud_kvstore_instance_engines": {Tok: dataSource(kvstoreMod, "getInstanceEngines")},
			"alicloud_kvstore_zones":            {Tok: dataSource(kvstoreMod, "getZones")},
			"alicloud_kvstore_connections":      {Tok: dataSource(kvstoreMod, "getConnections")},
			"alicloud_kvstore_accounts":         {Tok: dataSource(kvstoreMod, "getAccounts")},
			"alicloud_kvstore_permission":       {Tok: dataSource(kvstoreMod, "getPermission")},

			// Lindorm
			"alicloud_lindorm_instances": {Tok: dataSource(lindormMod, "getInstances")},

			// Log
			"alicloud_log_service":        {Tok: dataSource(logMod, "getService")},
			"alicloud_log_projects":       {Tok: dataSource(logMod, "getProjects")},
			"alicloud_log_stores":         {Tok: dataSource(logMod, "getStores")},
			"alicloud_log_alert_resource": {Tok: dataSource(logMod, "getAlertResource")},

			// Marketplace
			"alicloud_market_products": {Tok: dataSource(marketPlaceMod, "getProducts")},
			"alicloud_market_product":  {Tok: dataSource(marketPlaceMod, "getProduct")},

			// MaxCompute
			"alicloud_maxcompute_service": {Tok: dataSource(maxComputeMod, "getService")},

			// MHub
			"alicloud_mhub_products": {Tok: dataSource(mhubMod, "getProducts")},
			"alicloud_mhub_apps":     {Tok: dataSource(mhubMod, "getApps")},

			// Mns
			"alicloud_mns_queues":              {Tok: dataSource(mnsMod, "getQueues")},
			"alicloud_mns_topic_subscriptions": {Tok: dataSource(mnsMod, "getTopicSubscriptions")},
			"alicloud_mns_topics":              {Tok: dataSource(mnsMod, "getTopics")},
			"alicloud_mns_service":             {Tok: dataSource(mnsMod, "getService")},

			// Mongo
			"alicloud_mongodb_instances":                          {Tok: dataSource(mongoDbMod, "getInstances")},
			"alicloud_mongodb_zones":                              {Tok: dataSource(mongoDbMod, "getZones")},
			"alicloud_mongodb_serverless_instances":               {Tok: dataSource(mongoDbMod, "getServerlessInstances")},
			"alicloud_mongodb_accounts":                           {Tok: dataSource(mongoDbMod, "getAccounts")},
			"alicloud_mongodb_audit_policies":                     {Tok: dataSource(mongoDbMod, "getAuditPolicies")},
			"alicloud_mongodb_sharding_network_public_addresses":  {Tok: dataSource(mongoDbMod, "getShardingNetworkPublicAddresses")},
			"alicloud_mongodb_sharding_network_private_addresses": {Tok: dataSource(mongoDbMod, "getShardingNetworkPrivateAddresses")},

			// Mse
			"alicloud_mse_clusters":          {Tok: dataSource(mseMod, "getClusters")},
			"alicloud_mse_gateways":          {Tok: dataSource(mseMod, "getGateways")},
			"alicloud_mse_znodes":            {Tok: dataSource(mseMod, "getZnodes")},
			"alicloud_mse_engine_namespaces": {Tok: dataSource(mseMod, "getEngineNamespaces")},

			// Nas
			"alicloud_nas_access_groups":          {Tok: dataSource(nasMod, "getAccessGroups")},
			"alicloud_nas_access_rules":           {Tok: dataSource(nasMod, "getAccessRules")},
			"alicloud_nas_file_systems":           {Tok: dataSource(nasMod, "getFileSystems")},
			"alicloud_nas_mount_targets":          {Tok: dataSource(nasMod, "getMountTargets")},
			"alicloud_nas_protocols":              {Tok: dataSource(nasMod, "getProtocols")},
			"alicloud_nas_service":                {Tok: dataSource(nasMod, "getService")},
			"alicloud_nas_zones":                  {Tok: dataSource(nasMod, "getZones")},
			"alicloud_nas_snapshots":              {Tok: dataSource(nasMod, "getSnapshots")},
			"alicloud_nas_lifecycle_policies":     {Tok: dataSource(nasMod, "getLifecyclePolicies")},
			"alicloud_nas_filesets":               {Tok: dataSource(nasMod, "getFilesets")},
			"alicloud_nas_data_flows":             {Tok: dataSource(nasMod, "getDataFlows")},
			"alicloud_nas_auto_snapshot_policies": {Tok: dataSource(nasMod, "getAutoSnapshotPolicies")},

			// Oos
			"alicloud_oos_templates":            {Tok: dataSource(oosMod, "getTemplates")},
			"alicloud_oos_executions":           {Tok: dataSource(oosMod, "getExecutions")},
			"alicloud_oos_applications":         {Tok: dataSource(oosMod, "getApplications")},
			"alicloud_oos_patch_baselines":      {Tok: dataSource(oosMod, "getPatchBaselines")},
			"alicloud_oos_application_groups":   {Tok: dataSource(oosMod, "getApplicationGroups")},
			"alicloud_oos_secret_parameters":    {Tok: dataSource(oosMod, "getSecretParameters")},
			"alicloud_oos_state_configurations": {Tok: dataSource(oosMod, "getStateConfigurations")},
			"alicloud_oos_parameters":           {Tok: dataSource(oosMod, "getParameters")},

			// OpenSearch
			"alicloud_open_search_app_groups": {Tok: dataSource(openSearchMod, "getAppGroups")},

			// Oss
			"alicloud_oss_bucket_objects": {Tok: dataSource(ossMod, "getBucketObjects")},
			"alicloud_oss_buckets":        {Tok: dataSource(ossMod, "getBuckets")},
			"alicloud_oss_service":        {Tok: dataSource(ossMod, "getService")},

			// Ots
			"alicloud_ots_service": {Tok: dataSource(otsMod, "getService")},
			"alicloud_ots_tunnels": {Tok: dataSource(otsMod, "getTunnels")},

			// PrivateLink
			"alicloud_privatelink_vpc_endpoints":      {Tok: dataSource(privateLinkMod, "getVpcEndpoints")},
			"alicloud_privatelink_service":            {Tok: dataSource(privateLinkMod, "getService")},
			"alicloud_privatelink_vpc_endpoint_zones": {Tok: dataSource(privateLinkMod, "getVpcEndpointZones")},
			"alicloud_privatelink_vpc_endpoint_services": {
				Tok: dataSource(privateLinkMod, "getVpcEndpointServices"),
			},
			"alicloud_privatelink_vpc_endpoint_service_resources": {
				Tok: dataSource(privateLinkMod, "getVpcEndpointServiceResources"),
			},
			"alicloud_privatelink_vpc_endpoint_service_users": {
				Tok: dataSource(privateLinkMod, "getVpcEndpointServiceUsers"),
			},
			"alicloud_privatelink_vpc_endpoint_connections": {
				Tok: dataSource(privateLinkMod, "getVpcEndpointConnections"),
			},

			// PolarDb
			"alicloud_polardb_clusters":                 {Tok: dataSource(polarDbMod, "getClusters")},
			"alicloud_polardb_endpoints":                {Tok: dataSource(polarDbMod, "getEndpoints")},
			"alicloud_polardb_accounts":                 {Tok: dataSource(polarDbMod, "getAccounts")},
			"alicloud_polardb_databases":                {Tok: dataSource(polarDbMod, "getDatabases")},
			"alicloud_polardb_zones":                    {Tok: dataSource(polarDbMod, "getZones")},
			"alicloud_polardb_node_classes":             {Tok: dataSource(polarDbMod, "getNodeClasses")},
			"alicloud_polardb_global_database_networks": {Tok: dataSource(polarDbMod, "getGlobalDatabaseNetworks")},

			// Pvtr
			"alicloud_pvtz_zone_records":   {Tok: dataSource(pvtzMod, "getZoneRecords")},
			"alicloud_pvtz_zones":          {Tok: dataSource(pvtzMod, "getZones")},
			"alicloud_pvtz_service":        {Tok: dataSource(pvtzMod, "getService")},
			"alicloud_pvtz_endpoints":      {Tok: dataSource(pvtzMod, "getEndpoints")},
			"alicloud_pvtz_resolver_zones": {Tok: dataSource(pvtzMod, "getResolverZones")},
			"alicloud_pvtz_rules":          {Tok: dataSource(pvtzMod, "getRules")},

			// quickbi
			"alicloud_quick_bi_users": {Tok: dataSource(quickBiMod, "getUsers")},

			// quotas
			"alicloud_quotas_quotas": {Tok: dataSource(quotasMod, "getQuotas")},
			"alicloud_quotas_application_infos": {
				Tok: dataSource(quotasMod, "getApplicationInfos"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_quotas_quota_alarms":       {Tok: dataSource(quotasMod, "getQuotaAlarms")},
			"alicloud_quotas_quota_applications": {Tok: dataSource(quotasMod, "getQuotaApplications")},

			// Ram
			"alicloud_ram_account_alias": {
				Tok: dataSource(ramMod, "getAccountAlias"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_ram_account_aliases": {Tok: dataSource(ramMod, "getAccountAliases")},
			"alicloud_ram_groups":          {Tok: dataSource(ramMod, "getGroups")},
			"alicloud_ram_policies":        {Tok: dataSource(ramMod, "getPolicies")},
			"alicloud_ram_roles":           {Tok: dataSource(ramMod, "getRoles")},
			"alicloud_ram_users":           {Tok: dataSource(ramMod, "getUsers")},
			"alicloud_ram_saml_providers":  {Tok: dataSource(ramMod, "getSamlProviders")},

			// Rdc
			"alicloud_rdc_organizations": {Tok: dataSource(rdcMod, "getOrganizations")},

			// Rds
			"alicloud_db_instances":        {Tok: dataSource(rdsMod, "getInstances")},
			"alicloud_db_instance_classes": {Tok: dataSource(rdsMod, "getInstanceClasses")},
			"alicloud_db_instance_engines": {Tok: dataSource(rdsMod, "getInstanceEngines")},
			"alicloud_db_zones":            {Tok: dataSource(rdsMod, "getZones")},
			"alicloud_rds_parameter_groups": {
				Tok: dataSource(rdsMod, "getRdsParameterGroups"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_rds_accounts":              {Tok: dataSource(rdsMod, "getAccounts")},
			"alicloud_rds_backups":               {Tok: dataSource(rdsMod, "getRdsBackups")},
			"alicloud_rds_modify_parameter_logs": {Tok: dataSource(rdsMod, "getModifyParameterLogs")},

			// ResourceManager
			"alicloud_resource_manager_folders":          {Tok: dataSource(resourceManagerMod, "getFolders")},
			"alicloud_resource_manager_resource_groups":  {Tok: dataSource(resourceManagerMod, "getResourceGroups")},
			"alicloud_resource_manager_policy_versions":  {Tok: dataSource(resourceManagerMod, "getPolicyVersions")},
			"alicloud_resource_manager_handshakes":       {Tok: dataSource(resourceManagerMod, "getHandshakes")},
			"alicloud_resource_manager_accounts":         {Tok: dataSource(resourceManagerMod, "getAccounts")},
			"alicloud_resource_manager_roles":            {Tok: dataSource(resourceManagerMod, "getRoles")},
			"alicloud_resource_manager_policies":         {Tok: dataSource(resourceManagerMod, "getPolicies")},
			"alicloud_resource_manager_shared_resources": {Tok: dataSource(resourceManagerMod, "getSharedResources")},
			"alicloud_resource_manager_shared_targets":   {Tok: dataSource(resourceManagerMod, "getSharedTargets")},
			"alicloud_resource_manager_resource_shares":  {Tok: dataSource(resourceManagerMod, "getResourceShares")},
			"alicloud_resource_manager_control_policies": {Tok: dataSource(resourceManagerMod, "getControlPolicies")},
			"alicloud_resource_manager_control_policy_attachments": {
				Tok: dataSource(resourceManagerMod, "getControlPolicyAttachments"),
			},
			"alicloud_resource_manager_resource_directories": {
				Tok: dataSource(resourceManagerMod, "getResourceDirectories"),
			},
			"alicloud_resource_manager_policy_attachments": {
				Tok: dataSource(resourceManagerMod, "getPolicyAttachments"),
			},
			"alicloud_resource_manager_delegated_administrators": {
				Tok: dataSource(resourceManagerMod, "getDelegatedAdministrators"),
			},

			// RocketMq
			"alicloud_ons_groups":    {Tok: dataSource(rocketMqMod, "getGroups")},
			"alicloud_ons_instances": {Tok: dataSource(rocketMqMod, "getInstances")},
			"alicloud_ons_topics":    {Tok: dataSource(rocketMqMod, "getTopics")},
			"alicloud_ons_service":   {Tok: dataSource(rocketMqMod, "getService")},

			//Ros
			"alicloud_ros_change_sets":        {Tok: dataSource(rosMod, "getChangeSets")},
			"alicloud_ros_stack_groups":       {Tok: dataSource(rosMod, "getStackGroups")},
			"alicloud_ros_stacks":             {Tok: dataSource(rosMod, "getStacks")},
			"alicloud_ros_templates":          {Tok: dataSource(rosMod, "getTemplates")},
			"alicloud_ros_stack_instances":    {Tok: dataSource(rosMod, "getStackInstances")},
			"alicloud_ros_regions":            {Tok: dataSource(rosMod, "getRegions")},
			"alicloud_ros_template_scratches": {Tok: dataSource(rosMod, "getTemplateScratches")},

			// Sae
			"alicloud_sae_service":                   {Tok: dataSource(saeMod, "getService")},
			"alicloud_sae_namespaces":                {Tok: dataSource(saeMod, "getNamespaces")},
			"alicloud_sae_config_maps":               {Tok: dataSource(saeMod, "getConfigMaps")},
			"alicloud_sae_applications":              {Tok: dataSource(saeMod, "getApplications")},
			"alicloud_sae_ingresses":                 {Tok: dataSource(saeMod, "getIngresses")},
			"alicloud_sae_instance_specifications":   {Tok: dataSource(saeMod, "getInstanceSpecifications")},
			"alicloud_sae_application_scaling_rules": {Tok: dataSource(saeMod, "getApplicationScalingRules")},
			"alicloud_sae_grey_tag_routes":           {Tok: dataSource(saeMod, "getGreyTagRoutes")},

			// Sag
			"alicloud_sag_acls":          {Tok: dataSource(sagMod, "getAcls")},
			"alicloud_smartag_flow_logs": {Tok: dataSource(sagMod, "getSmartagFlowLogs")},

			// Sas
			"alicloud_simple_application_server_instances":      {Tok: dataSource(sasMod, "getInstances")},
			"alicloud_simple_application_server_images":         {Tok: dataSource(sasMod, "getImages")},
			"alicloud_simple_application_server_plans":          {Tok: dataSource(sasMod, "getServerPlans")},
			"alicloud_simple_application_server_custom_images":  {Tok: dataSource(sasMod, "getServerCustomImages")},
			"alicloud_simple_application_server_disks":          {Tok: dataSource(sasMod, "getServerDisks")},
			"alicloud_simple_application_server_firewall_rules": {Tok: dataSource(sasMod, "getServerFirewallRules")},
			"alicloud_simple_application_server_snapshots":      {Tok: dataSource(sasMod, "getServerSnapshots")},

			// Scdn
			"alicloud_scdn_domains": {Tok: dataSource(scdnMod, "getDomains")},

			// SchedulerX
			"alicloud_schedulerx_namespaces": {Tok: dataSource(schedulerXMod, "getNamespaces")},

			// sddp
			"alicloud_sddp_rules":       {Tok: dataSource(sddpMod, "getRules")},
			"alicloud_sddp_configs":     {Tok: dataSource(sddpMod, "getConfigs")},
			"alicloud_sddp_instances":   {Tok: dataSource(sddpMod, "getInstances")},
			"alicloud_sddp_data_limits": {Tok: dataSource(sddpMod, "getDataLimits")},

			// Security Center
			"alicloud_security_center_groups": {Tok: dataSource(securityCenterMod, "getGroups")},

			// Service Mesh
			"alicloud_service_mesh_service_meshes": {Tok: dataSource(serviceMeshMod, "getServiceMeshes")},
			"alicloud_service_mesh_versions":       {Tok: dataSource(serviceMeshMod, "getVersions")},

			// Slb
			"alicloud_slb_attachments":                {Tok: dataSource(slbMod, "getAttachments")},
			"alicloud_slb_listeners":                  {Tok: dataSource(slbMod, "getListeners")},
			"alicloud_slb_rules":                      {Tok: dataSource(slbMod, "getRules")},
			"alicloud_slb_server_groups":              {Tok: dataSource(slbMod, "getServerGroups")},
			"alicloud_slbs":                           {Tok: dataSource(slbMod, "getLoadBalancers")},
			"alicloud_slb_acls":                       {Tok: dataSource(slbMod, "getAcls")},
			"alicloud_slb_backend_servers":            {Tok: dataSource(slbMod, "getBackendServers")},
			"alicloud_slb_ca_certificates":            {Tok: dataSource(slbMod, "getCaCertificates")},
			"alicloud_slb_domain_extensions":          {Tok: dataSource(slbMod, "getDomainExtensions")},
			"alicloud_slb_master_slave_server_groups": {Tok: dataSource(slbMod, "getMasterSlaveServerGroups")},
			"alicloud_slb_server_certificates":        {Tok: dataSource(slbMod, "getServerCertificates")},
			"alicloud_slb_zones":                      {Tok: dataSource(slbMod, "getZones")},
			"alicloud_slb_load_balancers":             {Tok: dataSource(slbMod, "getApplicationLoadBalancers")},
			"alicloud_slb_tls_cipher_policies":        {Tok: dataSource(slbMod, "getTlsCipherPolicies")},

			// Tag
			"alicloud_tag_meta_tags": {Tok: dataSource(tagMod, "getMetaTags")},

			// tsdb
			"alicloud_tsdb_instances": {Tok: dataSource(tsdbMod, "getInstances")},
			"alicloud_tsdb_zones":     {Tok: dataSource(tsdbMod, "getZones")},

			// Vpc
			"alicloud_vpcs":                      {Tok: dataSource(vpcMod, "getNetworks")},
			"alicloud_router_interfaces":         {Tok: dataSource(vpcMod, "getRouterInterfaces")},
			"alicloud_forward_entries":           {Tok: dataSource(vpcMod, "getForwardEntries")},
			"alicloud_nat_gateways":              {Tok: dataSource(vpcMod, "getNatGateways")},
			"alicloud_route_entries":             {Tok: dataSource(vpcMod, "getRouteEntries")},
			"alicloud_route_tables":              {Tok: dataSource(vpcMod, "getRouteTables")},
			"alicloud_snat_entries":              {Tok: dataSource(vpcMod, "getSnatEntries")},
			"alicloud_common_bandwidth_packages": {Tok: dataSource(vpcMod, "getCommonBandwidthPackages")},
			"alicloud_enhanced_nat_available_zones": {
				Tok: dataSource(vpcMod, "getEnhancedNatAvailableZones"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_havips":                     {Tok: dataSource(vpcMod, "getHavips")},
			"alicloud_vpc_flow_logs":              {Tok: dataSource(vpcMod, "getVpcFlowLogs")},
			"alicloud_network_acls":               {Tok: dataSource(vpcMod, "getNetworkAcls")},
			"alicloud_vpc_dhcp_options_sets":      {Tok: dataSource(vpcMod, "getDhcpOptionsSets")},
			"alicloud_vpc_nat_ip_cidrs":           {Tok: dataSource(vpcMod, "getNatIpCidrs")},
			"alicloud_vpc_nat_ips":                {Tok: dataSource(vpcMod, "getNatIps")},
			"alicloud_vpc_traffic_mirror_filters": {Tok: dataSource(vpcMod, "getTrafficMirrorFilters")},
			"alicloud_vpc_traffic_mirror_filter_egress_rules": {
				Tok: dataSource(vpcMod, "getTrafficMirrorFilterEgressRules"),
			},
			"alicloud_vpc_traffic_mirror_filter_ingress_rules": {
				Tok: dataSource(vpcMod, "getTrafficMirrorFilterIngressRules"),
			},
			"alicloud_vpc_traffic_mirror_service":   {Tok: dataSource(vpcMod, "getTrafficMirrorService")},
			"alicloud_vpc_ipv6_addresses":           {Tok: dataSource(vpcMod, "getIpv6Addresses")},
			"alicloud_vpc_ipv6_egress_rules":        {Tok: dataSource(vpcMod, "getIpv6EgressRules")},
			"alicloud_vpc_ipv6_gateways":            {Tok: dataSource(vpcMod, "getIpv6Gateways")},
			"alicloud_vpc_ipv6_internet_bandwidths": {Tok: dataSource(vpcMod, "getIpv6InternetBandwidths")},
			"alicloud_vpc_traffic_mirror_sessions":  {Tok: dataSource(vpcMod, "getTrafficMirrorSessions")},
			"alicloud_vpc_bgp_groups":               {Tok: dataSource(vpcMod, "getBgpGroups")},
			"alicloud_vpc_bgp_networks":             {Tok: dataSource(vpcMod, "getBgpNetworks")},
			"alicloud_vpc_bgp_peers":                {Tok: dataSource(vpcMod, "getBgpPeers")},
			"alicloud_vpc_ipv4_gateways":            {Tok: dataSource(vpcMod, "getIpv4Gateways")},
			"alicloud_vpc_prefix_lists":             {Tok: dataSource(vpcMod, "getPrefixLists")},

			// Vod
			"alicloud_vod_domains": {Tok: dataSource(vodMod, "getDomains")},

			// Video Surveillance
			"alicloud_vs_service":                       {Tok: dataSource(vsMod, "getService")},
			"alicloud_video_surveillance_system_groups": {Tok: dataSource(vsMod, "getSystemGroups")},

			// Vpn
			"alicloud_vpn_connections":             {Tok: dataSource(vpnMod, "getConnections")},
			"alicloud_vpn_customer_gateways":       {Tok: dataSource(vpnMod, "getCustomerGateways")},
			"alicloud_vpn_gateways":                {Tok: dataSource(vpnMod, "getGateways")},
			"alicloud_vswitches":                   {Tok: dataSource(vpcMod, "getSwitches")},
			"alicloud_ssl_vpn_client_certs":        {Tok: dataSource(vpcMod, "getSslVpnClientCerts")},
			"alicloud_ssl_vpn_servers":             {Tok: dataSource(vpcMod, "getSslVpnServers")},
			"alicloud_vpn_ipsec_servers":           {Tok: dataSource(vpcMod, "getIpsecServers")},
			"alicloud_vpn_pbr_route_entries":       {Tok: dataSource(vpcMod, "getPbrRouteEntries")},
			"alicloud_vpn_gateway_vpn_attachments": {Tok: dataSource(vpnMod, "getGatewayVpnAttachments")},

			// Yundun
			"alicloud_yundun_bastionhost_instances": {
				Tok: dataSource(yundunMod, "getBastionHostInstances"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
			"alicloud_yundun_dbaudit_instance": {
				Tok: dataSource(yundunMod, "getDBAuditInstance"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},

			// Adb
			"alicloud_adb_clusters":    {Tok: dataSource(adbMod, "getClusters")},
			"alicloud_adb_zones":       {Tok: dataSource(adbMod, "getZones")},
			"alicloud_adb_db_clusters": {Tok: dataSource(adbMod, "getDBClusters")},

			// Waf
			"alicloud_waf_domains":      {Tok: dataSource(wafMod, "getDomains")},
			"alicloud_waf_instances":    {Tok: dataSource(wafMod, "getInstances")},
			"alicloud_waf_certificates": {Tok: dataSource(wafMod, "getCertificates")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/rhysmdnz/pulumi-%[1]s/sdk/", alicloudPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				alicloudPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: namespaceMap,
		},
	}

	prov.RenameResourceWithAlias("alicloud_ddosbgp_instance", aliResource(dnsMod, "DdosBgpInstance"),
		aliResource(ddosMod, "DdosBgpInstance"), dnsMod, ddosMod, nil)
	prov.RenameResourceWithAlias("alicloud_ddoscoo_instance", aliResource(dnsMod, "DdosCooInstance"),
		aliResource(ddosMod, "DdosCooInstance"), dnsMod, ddosMod, nil)

	prov.RenameDataSource("alicloud_ots_instance_attachments", dataSource(ossMod, "getInstanceAttachments"),
		dataSource(otsMod, "getInstanceAttachments"), ossMod, otsMod, nil)
	prov.RenameDataSource("alicloud_ots_instances", dataSource(ossMod, "getInstances"),
		dataSource(otsMod, "getInstances"), ossMod, otsMod, nil)
	prov.RenameDataSource("alicloud_ots_tables", dataSource(ossMod, "getTables"),
		dataSource(otsMod, "getTables"), ossMod, otsMod, nil)

	prov.SetAutonaming(255, "-")

	return prov
}
