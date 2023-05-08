// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
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
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CpuCoreCount = 1,
    ///         MemorySize = 2,
    ///     });
    /// 
    ///     var defaultImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_18.*64",
    ///         MostRecent = true,
    ///         Owners = "system",
    ///     });
    /// 
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "test_havip_attachment";
    ///     var fooNetwork = new AliCloud.Vpc.Network("fooNetwork", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var fooSwitch = new AliCloud.Vpc.Switch("fooSwitch", new()
    ///     {
    ///         VpcId = fooNetwork.Id,
    ///         CidrBlock = "172.16.0.0/21",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var fooHAVip = new AliCloud.Vpc.HAVip("fooHAVip", new()
    ///     {
    ///         VswitchId = fooSwitch.Id,
    ///         Description = name,
    ///     });
    /// 
    ///     var tfTestFoo = new AliCloud.Ecs.SecurityGroup("tfTestFoo", new()
    ///     {
    ///         Description = "foo",
    ///         VpcId = fooNetwork.Id,
    ///     });
    /// 
    ///     var fooInstance = new AliCloud.Ecs.Instance("fooInstance", new()
    ///     {
    ///         AvailabilityZone = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchId = fooSwitch.Id,
    ///         ImageId = defaultImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceType = defaultInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///         SystemDiskCategory = "cloud_efficiency",
    ///         InternetChargeType = "PayByTraffic",
    ///         InternetMaxBandwidthOut = 5,
    ///         SecurityGroups = new[]
    ///         {
    ///             tfTestFoo.Id,
    ///         },
    ///         InstanceName = name,
    ///         UserData = "echo 'net.ipv4.ip_forward=1'&gt;&gt; /etc/sysctl.conf",
    ///     });
    /// 
    ///     var fooHAVipAttachment = new AliCloud.Vpc.HAVipAttachment("fooHAVipAttachment", new()
    ///     {
    ///         HavipId = fooHAVip.Id,
    ///         InstanceId = fooInstance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The havip attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/hAVipAttachment:HAVipAttachment foo havip-abc123456:i-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/hAVipAttachment:HAVipAttachment")]
    public partial class HAVipAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
        /// </summary>
        [Output("force")]
        public Output<string?> Force { get; private set; } = null!;

        /// <summary>
        /// The havip_id of the havip attachment, the field can't be changed.
        /// </summary>
        [Output("havipId")]
        public Output<string> HavipId { get; private set; } = null!;

        /// <summary>
        /// The instance_id of the havip attachment, the field can't be changed.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// (Available in v1.201.0+) The status of the HaVip instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a HAVipAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HAVipAttachment(string name, HAVipAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/hAVipAttachment:HAVipAttachment", name, args ?? new HAVipAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HAVipAttachment(string name, Input<string> id, HAVipAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/hAVipAttachment:HAVipAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HAVipAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HAVipAttachment Get(string name, Input<string> id, HAVipAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new HAVipAttachment(name, id, state, options);
        }
    }

    public sealed class HAVipAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
        /// </summary>
        [Input("force")]
        public Input<string>? Force { get; set; }

        /// <summary>
        /// The havip_id of the havip attachment, the field can't be changed.
        /// </summary>
        [Input("havipId", required: true)]
        public Input<string> HavipId { get; set; } = null!;

        /// <summary>
        /// The instance_id of the havip attachment, the field can't be changed.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public HAVipAttachmentArgs()
        {
        }
        public static new HAVipAttachmentArgs Empty => new HAVipAttachmentArgs();
    }

    public sealed class HAVipAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
        /// </summary>
        [Input("force")]
        public Input<string>? Force { get; set; }

        /// <summary>
        /// The havip_id of the havip attachment, the field can't be changed.
        /// </summary>
        [Input("havipId")]
        public Input<string>? HavipId { get; set; }

        /// <summary>
        /// The instance_id of the havip attachment, the field can't be changed.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// (Available in v1.201.0+) The status of the HaVip instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public HAVipAttachmentState()
        {
        }
        public static new HAVipAttachmentState Empty => new HAVipAttachmentState();
    }
}
