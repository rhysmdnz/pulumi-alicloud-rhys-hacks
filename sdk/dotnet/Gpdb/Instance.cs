// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Gpdb
{
    /// <summary>
    /// Provides a AnalyticDB for PostgreSQL instance resource supports replica set instances only. the AnalyticDB for PostgreSQL provides stable, reliable, and automatic scalable database services.
    /// You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/35387.htm)
    /// 
    /// &gt; **NOTE:**  Available in 1.47.0+
    /// 
    /// ## Import
    /// 
    /// AnalyticDB for PostgreSQL can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:gpdb/instance:Instance example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:gpdb/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Field `availability_zone` has been deprecated from provider version 1.187.0. New field `zone_id` instead.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.196.0+) The connection string of the instance.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// Whether to load the sample dataset after the instance is created. Valid values: `true`, `false`.
        /// </summary>
        [Output("createSampleData")]
        public Output<bool> CreateSampleData { get; private set; } = null!;

        /// <summary>
        /// The db instance category. Valid values: `HighAvailability`, `Basic`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Output("dbInstanceCategory")]
        public Output<string> DbInstanceCategory { get; private set; } = null!;

        /// <summary>
        /// The db instance class. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Output("dbInstanceClass")]
        public Output<string?> DbInstanceClass { get; private set; } = null!;

        /// <summary>
        /// The db instance mode. Valid values: `StorageElastic`, `Serverless`, `Classic`.
        /// </summary>
        [Output("dbInstanceMode")]
        public Output<string> DbInstanceMode { get; private set; } = null!;

        /// <summary>
        /// The description of the instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The database engine used by the instance. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// The version of the database engine used by the instance.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// Field `instance_charge_type` has been deprecated from provider version 1.187.0. New field `payment_type` instead.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// The number of nodes. Valid values: `2`, `4`, `8`, `12`, `16`, `24`, `32`, `64`, `96`, `128`.
        /// </summary>
        [Output("instanceGroupCount")]
        public Output<int?> InstanceGroupCount { get; private set; } = null!;

        /// <summary>
        /// The network type of the instance.
        /// </summary>
        [Output("instanceNetworkType")]
        public Output<string> InstanceNetworkType { get; private set; } = null!;

        /// <summary>
        /// The specification of segment nodes.
        /// * When `db_instance_category` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
        /// * When `db_instance_category` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
        /// * When `db_instance_category` is `Serverless`, Valid values: `4C16G`, `8C32G`.
        /// &gt; **NOTE:** This parameter must be passed to create a storage elastic mode instance and a serverless version instance.
        /// </summary>
        [Output("instanceSpec")]
        public Output<string?> InstanceSpec { get; private set; } = null!;

        /// <summary>
        /// The ip whitelist.
        /// </summary>
        [Output("ipWhitelists")]
        public Output<ImmutableArray<Outputs.InstanceIpWhitelist>> IpWhitelists { get; private set; } = null!;

        /// <summary>
        /// The end time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 03:00Z. start time should be later than end time.
        /// </summary>
        [Output("maintainEndTime")]
        public Output<string> MaintainEndTime { get; private set; } = null!;

        /// <summary>
        /// The start time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 02:00Z.
        /// </summary>
        [Output("maintainStartTime")]
        public Output<string> MaintainStartTime { get; private set; } = null!;

        /// <summary>
        /// The number of Master nodes. Valid values: 1 to 2. if it is not filled in, the default value is 1 Master node.
        /// </summary>
        [Output("masterNodeNum")]
        public Output<int?> MasterNodeNum { get; private set; } = null!;

        /// <summary>
        /// The billing method of the instance. Valid values: `Subscription`, `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy the resource, in month. required when `payment_type` is `Subscription`. Valid values: `Year`, `Month`.
        /// </summary>
        [Output("period")]
        public Output<string?> Period { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.196.0+) The connection port of the instance.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// The private ip address.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string?> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// The ID of the enterprise resource group to which the instance belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
        /// </summary>
        [Output("securityIpLists")]
        public Output<ImmutableArray<string>> SecurityIpLists { get; private set; } = null!;

        /// <summary>
        /// Calculate the number of nodes. The value range of the high-availability version of the storage elastic mode is 4 to 512, and the value must be a multiple of 4. The value range of the basic version of the storage elastic mode is 2 to 512, and the value must be a multiple of 2. The-Serverless version has a value range of 2 to 512. The value must be a multiple of 2.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage elastic mode instance and a Serverless version instance. During the public beta of the Serverless version (from 0101, 2022 to 0131, 2022), a maximum of 12 compute nodes can be created.
        /// </summary>
        [Output("segNodeNum")]
        public Output<int> SegNodeNum { get; private set; } = null!;

        /// <summary>
        /// The seg storage type. Valid values: `cloud_essd`, `cloud_efficiency`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage elastic mode instance. Storage Elastic Mode Basic Edition instances only support ESSD cloud disks.
        /// </summary>
        [Output("segStorageType")]
        public Output<string?> SegStorageType { get; private set; } = null!;

        /// <summary>
        /// Enable or disable SSL. Valid values: `0` and `1`.
        /// </summary>
        [Output("sslEnabled")]
        public Output<int> SslEnabled { get; private set; } = null!;

        /// <summary>
        /// The status of the instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The storage capacity. Unit: GB. Value: `50` to `4000`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Output("storageSize")]
        public Output<int> StorageSize { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The used time. When the parameter `period` is `Year`, the `used_time` value is 1 to 3. When the parameter `period` is `Month`, the `used_time` value is 1 to 9.
        /// </summary>
        [Output("usedTime")]
        public Output<string?> UsedTime { get; private set; } = null!;

        /// <summary>
        /// The vpc ID of the resource.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The vswitch id.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The zone ID of the instance.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:gpdb/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:gpdb/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Field `availability_zone` has been deprecated from provider version 1.187.0. New field `zone_id` instead.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Whether to load the sample dataset after the instance is created. Valid values: `true`, `false`.
        /// </summary>
        [Input("createSampleData")]
        public Input<bool>? CreateSampleData { get; set; }

        /// <summary>
        /// The db instance category. Valid values: `HighAvailability`, `Basic`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Input("dbInstanceCategory")]
        public Input<string>? DbInstanceCategory { get; set; }

        /// <summary>
        /// The db instance class. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Input("dbInstanceClass")]
        public Input<string>? DbInstanceClass { get; set; }

        /// <summary>
        /// The db instance mode. Valid values: `StorageElastic`, `Serverless`, `Classic`.
        /// </summary>
        [Input("dbInstanceMode", required: true)]
        public Input<string> DbInstanceMode { get; set; } = null!;

        /// <summary>
        /// The description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The database engine used by the instance. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// The version of the database engine used by the instance.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        /// <summary>
        /// Field `instance_charge_type` has been deprecated from provider version 1.187.0. New field `payment_type` instead.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The number of nodes. Valid values: `2`, `4`, `8`, `12`, `16`, `24`, `32`, `64`, `96`, `128`.
        /// </summary>
        [Input("instanceGroupCount")]
        public Input<int>? InstanceGroupCount { get; set; }

        /// <summary>
        /// The network type of the instance.
        /// </summary>
        [Input("instanceNetworkType")]
        public Input<string>? InstanceNetworkType { get; set; }

        /// <summary>
        /// The specification of segment nodes.
        /// * When `db_instance_category` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
        /// * When `db_instance_category` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
        /// * When `db_instance_category` is `Serverless`, Valid values: `4C16G`, `8C32G`.
        /// &gt; **NOTE:** This parameter must be passed to create a storage elastic mode instance and a serverless version instance.
        /// </summary>
        [Input("instanceSpec")]
        public Input<string>? InstanceSpec { get; set; }

        [Input("ipWhitelists")]
        private InputList<Inputs.InstanceIpWhitelistArgs>? _ipWhitelists;

        /// <summary>
        /// The ip whitelist.
        /// </summary>
        public InputList<Inputs.InstanceIpWhitelistArgs> IpWhitelists
        {
            get => _ipWhitelists ?? (_ipWhitelists = new InputList<Inputs.InstanceIpWhitelistArgs>());
            set => _ipWhitelists = value;
        }

        /// <summary>
        /// The end time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 03:00Z. start time should be later than end time.
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 02:00Z.
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// The number of Master nodes. Valid values: 1 to 2. if it is not filled in, the default value is 1 Master node.
        /// </summary>
        [Input("masterNodeNum")]
        public Input<int>? MasterNodeNum { get; set; }

        /// <summary>
        /// The billing method of the instance. Valid values: `Subscription`, `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy the resource, in month. required when `payment_type` is `Subscription`. Valid values: `Year`, `Month`.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The private ip address.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The ID of the enterprise resource group to which the instance belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityIpLists")]
        private InputList<string>? _securityIpLists;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
        /// </summary>
        [Obsolete(@"Field 'security_ip_list' has been deprecated from version 1.187.0. Use 'ip_whitelist' instead.")]
        public InputList<string> SecurityIpLists
        {
            get => _securityIpLists ?? (_securityIpLists = new InputList<string>());
            set => _securityIpLists = value;
        }

        /// <summary>
        /// Calculate the number of nodes. The value range of the high-availability version of the storage elastic mode is 4 to 512, and the value must be a multiple of 4. The value range of the basic version of the storage elastic mode is 2 to 512, and the value must be a multiple of 2. The-Serverless version has a value range of 2 to 512. The value must be a multiple of 2.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage elastic mode instance and a Serverless version instance. During the public beta of the Serverless version (from 0101, 2022 to 0131, 2022), a maximum of 12 compute nodes can be created.
        /// </summary>
        [Input("segNodeNum")]
        public Input<int>? SegNodeNum { get; set; }

        /// <summary>
        /// The seg storage type. Valid values: `cloud_essd`, `cloud_efficiency`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage elastic mode instance. Storage Elastic Mode Basic Edition instances only support ESSD cloud disks.
        /// </summary>
        [Input("segStorageType")]
        public Input<string>? SegStorageType { get; set; }

        /// <summary>
        /// Enable or disable SSL. Valid values: `0` and `1`.
        /// </summary>
        [Input("sslEnabled")]
        public Input<int>? SslEnabled { get; set; }

        /// <summary>
        /// The storage capacity. Unit: GB. Value: `50` to `4000`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Input("storageSize")]
        public Input<int>? StorageSize { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The used time. When the parameter `period` is `Year`, the `used_time` value is 1 to 3. When the parameter `period` is `Month`, the `used_time` value is 1 to 9.
        /// </summary>
        [Input("usedTime")]
        public Input<string>? UsedTime { get; set; }

        /// <summary>
        /// The vpc ID of the resource.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The vswitch id.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        /// <summary>
        /// The zone ID of the instance.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Field `availability_zone` has been deprecated from provider version 1.187.0. New field `zone_id` instead.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// (Available in 1.196.0+) The connection string of the instance.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// Whether to load the sample dataset after the instance is created. Valid values: `true`, `false`.
        /// </summary>
        [Input("createSampleData")]
        public Input<bool>? CreateSampleData { get; set; }

        /// <summary>
        /// The db instance category. Valid values: `HighAvailability`, `Basic`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Input("dbInstanceCategory")]
        public Input<string>? DbInstanceCategory { get; set; }

        /// <summary>
        /// The db instance class. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Input("dbInstanceClass")]
        public Input<string>? DbInstanceClass { get; set; }

        /// <summary>
        /// The db instance mode. Valid values: `StorageElastic`, `Serverless`, `Classic`.
        /// </summary>
        [Input("dbInstanceMode")]
        public Input<string>? DbInstanceMode { get; set; }

        /// <summary>
        /// The description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The database engine used by the instance. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The version of the database engine used by the instance.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Field `instance_charge_type` has been deprecated from provider version 1.187.0. New field `payment_type` instead.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The number of nodes. Valid values: `2`, `4`, `8`, `12`, `16`, `24`, `32`, `64`, `96`, `128`.
        /// </summary>
        [Input("instanceGroupCount")]
        public Input<int>? InstanceGroupCount { get; set; }

        /// <summary>
        /// The network type of the instance.
        /// </summary>
        [Input("instanceNetworkType")]
        public Input<string>? InstanceNetworkType { get; set; }

        /// <summary>
        /// The specification of segment nodes.
        /// * When `db_instance_category` is `HighAvailability`, Valid values: `2C16G`, `4C32G`, `16C128G`.
        /// * When `db_instance_category` is `Basic`, Valid values: `2C8G`, `4C16G`, `8C32G`, `16C64G`.
        /// * When `db_instance_category` is `Serverless`, Valid values: `4C16G`, `8C32G`.
        /// &gt; **NOTE:** This parameter must be passed to create a storage elastic mode instance and a serverless version instance.
        /// </summary>
        [Input("instanceSpec")]
        public Input<string>? InstanceSpec { get; set; }

        [Input("ipWhitelists")]
        private InputList<Inputs.InstanceIpWhitelistGetArgs>? _ipWhitelists;

        /// <summary>
        /// The ip whitelist.
        /// </summary>
        public InputList<Inputs.InstanceIpWhitelistGetArgs> IpWhitelists
        {
            get => _ipWhitelists ?? (_ipWhitelists = new InputList<Inputs.InstanceIpWhitelistGetArgs>());
            set => _ipWhitelists = value;
        }

        /// <summary>
        /// The end time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 03:00Z. start time should be later than end time.
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the maintenance window for the instance. in the format of HH:mmZ (UTC time), for example 02:00Z.
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// The number of Master nodes. Valid values: 1 to 2. if it is not filled in, the default value is 1 Master node.
        /// </summary>
        [Input("masterNodeNum")]
        public Input<int>? MasterNodeNum { get; set; }

        /// <summary>
        /// The billing method of the instance. Valid values: `Subscription`, `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy the resource, in month. required when `payment_type` is `Subscription`. Valid values: `Year`, `Month`.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// (Available in 1.196.0+) The connection port of the instance.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The private ip address.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The ID of the enterprise resource group to which the instance belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityIpLists")]
        private InputList<string>? _securityIpLists;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
        /// </summary>
        [Obsolete(@"Field 'security_ip_list' has been deprecated from version 1.187.0. Use 'ip_whitelist' instead.")]
        public InputList<string> SecurityIpLists
        {
            get => _securityIpLists ?? (_securityIpLists = new InputList<string>());
            set => _securityIpLists = value;
        }

        /// <summary>
        /// Calculate the number of nodes. The value range of the high-availability version of the storage elastic mode is 4 to 512, and the value must be a multiple of 4. The value range of the basic version of the storage elastic mode is 2 to 512, and the value must be a multiple of 2. The-Serverless version has a value range of 2 to 512. The value must be a multiple of 2.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage elastic mode instance and a Serverless version instance. During the public beta of the Serverless version (from 0101, 2022 to 0131, 2022), a maximum of 12 compute nodes can be created.
        /// </summary>
        [Input("segNodeNum")]
        public Input<int>? SegNodeNum { get; set; }

        /// <summary>
        /// The seg storage type. Valid values: `cloud_essd`, `cloud_efficiency`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage elastic mode instance. Storage Elastic Mode Basic Edition instances only support ESSD cloud disks.
        /// </summary>
        [Input("segStorageType")]
        public Input<string>? SegStorageType { get; set; }

        /// <summary>
        /// Enable or disable SSL. Valid values: `0` and `1`.
        /// </summary>
        [Input("sslEnabled")]
        public Input<int>? SslEnabled { get; set; }

        /// <summary>
        /// The status of the instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The storage capacity. Unit: GB. Value: `50` to `4000`.
        /// &gt; **NOTE:** This parameter must be passed in to create a storage reservation mode instance.
        /// </summary>
        [Input("storageSize")]
        public Input<int>? StorageSize { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The used time. When the parameter `period` is `Year`, the `used_time` value is 1 to 3. When the parameter `period` is `Month`, the `used_time` value is 1 to 9.
        /// </summary>
        [Input("usedTime")]
        public Input<string>? UsedTime { get; set; }

        /// <summary>
        /// The vpc ID of the resource.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The vswitch id.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The zone ID of the instance.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
