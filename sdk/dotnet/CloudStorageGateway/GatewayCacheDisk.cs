// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudStorageGateway
{
    /// <summary>
    /// Provides a Cloud Storage Gateway Gateway Cache Disk resource.
    /// 
    /// For information about Cloud Storage Gateway Gateway Cache Disk and how to use it, see [What is Gateway Cache Disk](https://www.alibabacloud.com/help/zh/doc-detail/170294.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.144.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleStocks = AliCloud.CloudStorageGateway.GetStocks.Invoke(new()
    ///     {
    ///         GatewayClass = "Standard",
    ///     });
    /// 
    ///     var vpc = new AliCloud.Vpc.Network("vpc", new()
    ///     {
    ///         VpcName = "example_value",
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var exampleSwitch = new AliCloud.Vpc.Switch("exampleSwitch", new()
    ///     {
    ///         VpcId = vpc.Id,
    ///         CidrBlock = "172.16.0.0/21",
    ///         ZoneId = exampleStocks.Apply(getStocksResult =&gt; getStocksResult.Stocks[0]?.ZoneId),
    ///         VswitchName = "example_value",
    ///     });
    /// 
    ///     var exampleStorageBundle = new AliCloud.CloudStorageGateway.StorageBundle("exampleStorageBundle", new()
    ///     {
    ///         StorageBundleName = "example_value",
    ///     });
    /// 
    ///     var exampleGateway = new AliCloud.CloudStorageGateway.Gateway("exampleGateway", new()
    ///     {
    ///         Description = "tf-acctestDesalone",
    ///         GatewayClass = "Standard",
    ///         Type = "File",
    ///         PaymentType = "PayAsYouGo",
    ///         VswitchId = exampleSwitch.Id,
    ///         ReleaseAfterExpiration = true,
    ///         PublicNetworkBandwidth = 10,
    ///         StorageBundleId = exampleStorageBundle.Id,
    ///         Location = "Cloud",
    ///         GatewayName = "example_value",
    ///     });
    /// 
    ///     var exampleGatewayCacheDisk = new AliCloud.CloudStorageGateway.GatewayCacheDisk("exampleGatewayCacheDisk", new()
    ///     {
    ///         CacheDiskCategory = "cloud_efficiency",
    ///         GatewayId = alicloud_cloud_storage_gateway_gateways.Example.Id,
    ///         CacheDiskSizeInGb = 50,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Storage Gateway Gateway Cache Disk can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk example &lt;gateway_id&gt;:&lt;cache_id&gt;:&lt;local_file_path&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk")]
    public partial class GatewayCacheDisk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The cache disk type. Valid values: `cloud_efficiency`, `cloud_ssd`.
        /// </summary>
        [Output("cacheDiskCategory")]
        public Output<string> CacheDiskCategory { get; private set; } = null!;

        /// <summary>
        /// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
        /// </summary>
        [Output("cacheDiskSizeInGb")]
        public Output<int> CacheDiskSizeInGb { get; private set; } = null!;

        /// <summary>
        /// The ID of the cache.
        /// </summary>
        [Output("cacheId")]
        public Output<string> CacheId { get; private set; } = null!;

        /// <summary>
        /// The ID of the gateway.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// The cache disk inside the device name.
        /// </summary>
        [Output("localFilePath")]
        public Output<string> LocalFilePath { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. Valid values: `0`, `1`, `2`. `0`: Normal. `1`: Is about to expire. `2`: Has expired.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayCacheDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayCacheDisk(string name, GatewayCacheDiskArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk", name, args ?? new GatewayCacheDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayCacheDisk(string name, Input<string> id, GatewayCacheDiskState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayCacheDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayCacheDisk Get(string name, Input<string> id, GatewayCacheDiskState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayCacheDisk(name, id, state, options);
        }
    }

    public sealed class GatewayCacheDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cache disk type. Valid values: `cloud_efficiency`, `cloud_ssd`.
        /// </summary>
        [Input("cacheDiskCategory")]
        public Input<string>? CacheDiskCategory { get; set; }

        /// <summary>
        /// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
        /// </summary>
        [Input("cacheDiskSizeInGb", required: true)]
        public Input<int> CacheDiskSizeInGb { get; set; } = null!;

        /// <summary>
        /// The ID of the gateway.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        public GatewayCacheDiskArgs()
        {
        }
        public static new GatewayCacheDiskArgs Empty => new GatewayCacheDiskArgs();
    }

    public sealed class GatewayCacheDiskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cache disk type. Valid values: `cloud_efficiency`, `cloud_ssd`.
        /// </summary>
        [Input("cacheDiskCategory")]
        public Input<string>? CacheDiskCategory { get; set; }

        /// <summary>
        /// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
        /// </summary>
        [Input("cacheDiskSizeInGb")]
        public Input<int>? CacheDiskSizeInGb { get; set; }

        /// <summary>
        /// The ID of the cache.
        /// </summary>
        [Input("cacheId")]
        public Input<string>? CacheId { get; set; }

        /// <summary>
        /// The ID of the gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// The cache disk inside the device name.
        /// </summary>
        [Input("localFilePath")]
        public Input<string>? LocalFilePath { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `0`, `1`, `2`. `0`: Normal. `1`: Is about to expire. `2`: Has expired.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public GatewayCacheDiskState()
        {
        }
        public static new GatewayCacheDiskState Empty => new GatewayCacheDiskState();
    }
}