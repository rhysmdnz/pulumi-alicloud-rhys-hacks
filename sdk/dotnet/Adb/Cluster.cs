// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Adb
{
    /// <summary>
    /// ## Example Usage
    /// ### Create a ADB MySQL cluster
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "adbClusterconfig";
    ///     var creation = config.Get("creation") ?? "ADB";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = creation,
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var defaultCluster = new AliCloud.Adb.Cluster("defaultCluster", new()
    ///     {
    ///         DbClusterVersion = "3.0",
    ///         DbClusterCategory = "Cluster",
    ///         DbNodeClass = "C8",
    ///         DbNodeCount = 2,
    ///         DbNodeStorage = 200,
    ///         PayType = "PostPaid",
    ///         Description = name,
    ///         VswitchId = defaultSwitch.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ADB cluster can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:adb/cluster:Cluster example am-abc12345678
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:adb/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int> AutoRenewPeriod { get; private set; } = null!;

        [Output("computeResource")]
        public Output<string?> ComputeResource { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.93.0+) The connection string of the ADB cluster.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// Cluster category. Value options: `Basic`, `Cluster`.
        /// </summary>
        [Output("dbClusterCategory")]
        public Output<string> DbClusterCategory { get; private set; } = null!;

        [Output("dbClusterClass")]
        public Output<string?> DbClusterClass { get; private set; } = null!;

        /// <summary>
        /// Cluster version. Value options: `3.0`, Default to `3.0`.
        /// </summary>
        [Output("dbClusterVersion")]
        public Output<string?> DbClusterVersion { get; private set; } = null!;

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Output("dbNodeClass")]
        public Output<string> DbNodeClass { get; private set; } = null!;

        /// <summary>
        /// The db_node_count of cluster node.
        /// </summary>
        [Output("dbNodeCount")]
        public Output<int> DbNodeCount { get; private set; } = null!;

        /// <summary>
        /// The db_node_storage of cluster node.
        /// </summary>
        [Output("dbNodeStorage")]
        public Output<int> DbNodeStorage { get; private set; } = null!;

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("elasticIoResource")]
        public Output<int> ElasticIoResource { get; private set; } = null!;

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Output("maintainTime")]
        public Output<string> MaintainTime { get; private set; } = null!;

        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        [Output("modifyType")]
        public Output<string?> ModifyType { get; private set; } = null!;

        /// <summary>
        /// Field `pay_type` has been deprecated. New field `payment_type` instead.
        /// </summary>
        [Output("payType")]
        public Output<string> PayType { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.196.0+) The connection port of the ADB cluster.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Output("renewalStatus")]
        public Output<string> RenewalStatus { get; private set; } = null!;

        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        [Output("securityIps")]
        public Output<ImmutableArray<string>> SecurityIps { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the DB cluster.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("alicloud:adb/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:adb/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/rhysmdnz/pulumi-alicloud",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        [Input("computeResource")]
        public Input<string>? ComputeResource { get; set; }

        /// <summary>
        /// Cluster category. Value options: `Basic`, `Cluster`.
        /// </summary>
        [Input("dbClusterCategory", required: true)]
        public Input<string> DbClusterCategory { get; set; } = null!;

        [Input("dbClusterClass")]
        public Input<string>? DbClusterClass { get; set; }

        /// <summary>
        /// Cluster version. Value options: `3.0`, Default to `3.0`.
        /// </summary>
        [Input("dbClusterVersion")]
        public Input<string>? DbClusterVersion { get; set; }

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Input("dbNodeClass")]
        public Input<string>? DbNodeClass { get; set; }

        /// <summary>
        /// The db_node_count of cluster node.
        /// </summary>
        [Input("dbNodeCount")]
        public Input<int>? DbNodeCount { get; set; }

        /// <summary>
        /// The db_node_storage of cluster node.
        /// </summary>
        [Input("dbNodeStorage")]
        public Input<int>? DbNodeStorage { get; set; }

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("elasticIoResource")]
        public Input<int>? ElasticIoResource { get; set; }

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        [Input("modifyType")]
        public Input<string>? ModifyType { get; set; }

        /// <summary>
        /// Field `pay_type` has been deprecated. New field `payment_type` instead.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB cluster.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        [Input("computeResource")]
        public Input<string>? ComputeResource { get; set; }

        /// <summary>
        /// (Available in 1.93.0+) The connection string of the ADB cluster.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// Cluster category. Value options: `Basic`, `Cluster`.
        /// </summary>
        [Input("dbClusterCategory")]
        public Input<string>? DbClusterCategory { get; set; }

        [Input("dbClusterClass")]
        public Input<string>? DbClusterClass { get; set; }

        /// <summary>
        /// Cluster version. Value options: `3.0`, Default to `3.0`.
        /// </summary>
        [Input("dbClusterVersion")]
        public Input<string>? DbClusterVersion { get; set; }

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Input("dbNodeClass")]
        public Input<string>? DbNodeClass { get; set; }

        /// <summary>
        /// The db_node_count of cluster node.
        /// </summary>
        [Input("dbNodeCount")]
        public Input<int>? DbNodeCount { get; set; }

        /// <summary>
        /// The db_node_storage of cluster node.
        /// </summary>
        [Input("dbNodeStorage")]
        public Input<int>? DbNodeStorage { get; set; }

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("elasticIoResource")]
        public Input<int>? ElasticIoResource { get; set; }

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("modifyType")]
        public Input<string>? ModifyType { get; set; }

        /// <summary>
        /// Field `pay_type` has been deprecated. New field `payment_type` instead.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `payment_type` supports updating from v1.166.0+.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// (Available in 1.196.0+) The connection port of the ADB cluster.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB cluster.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
