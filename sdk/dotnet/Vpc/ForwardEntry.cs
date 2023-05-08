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
    /// Provides a forward resource.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "forward-entry-example-name";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/21",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var defaultNatGateway = new AliCloud.Vpc.NatGateway("defaultNatGateway", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         InternetChargeType = "PayByLcu",
    ///         NatGatewayName = name,
    ///         NatType = "Enhanced",
    ///         VswitchId = defaultSwitch.Id,
    ///     });
    /// 
    ///     var defaultEipAddress = new AliCloud.Ecs.EipAddress("defaultEipAddress", new()
    ///     {
    ///         AddressName = name,
    ///     });
    /// 
    ///     var defaultEipAssociation = new AliCloud.Ecs.EipAssociation("defaultEipAssociation", new()
    ///     {
    ///         AllocationId = defaultEipAddress.Id,
    ///         InstanceId = defaultNatGateway.Id,
    ///     });
    /// 
    ///     var defaultForwardEntry = new AliCloud.Vpc.ForwardEntry("defaultForwardEntry", new()
    ///     {
    ///         ForwardTableId = defaultNatGateway.ForwardTableIds,
    ///         ExternalIp = defaultEipAddress.IpAddress,
    ///         ExternalPort = "80",
    ///         IpProtocol = "tcp",
    ///         InternalIp = "172.16.0.3",
    ///         InternalPort = "8080",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Forward Entry can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/forwardEntry:ForwardEntry foo ftb-1aece3:fwd-232ce2
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/forwardEntry:ForwardEntry")]
    public partial class ForwardEntry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The external ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Output("externalIp")]
        public Output<string> ExternalIp { get; private set; } = null!;

        /// <summary>
        /// The external port, valid value is 1~65535|any.
        /// </summary>
        [Output("externalPort")]
        public Output<string> ExternalPort { get; private set; } = null!;

        /// <summary>
        /// The id of the forward entry on the server.
        /// </summary>
        [Output("forwardEntryId")]
        public Output<string> ForwardEntryId { get; private set; } = null!;

        /// <summary>
        /// The name of forward entry.
        /// </summary>
        [Output("forwardEntryName")]
        public Output<string> ForwardEntryName { get; private set; } = null!;

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "forward_table_ids".
        /// </summary>
        [Output("forwardTableId")]
        public Output<string> ForwardTableId { get; private set; } = null!;

        /// <summary>
        /// The internal ip, must a private ip.
        /// </summary>
        [Output("internalIp")]
        public Output<string> InternalIp { get; private set; } = null!;

        /// <summary>
        /// The internal port, valid value is 1~65535|any.
        /// </summary>
        [Output("internalPort")]
        public Output<string> InternalPort { get; private set; } = null!;

        /// <summary>
        /// The ip protocol, valid value is tcp|udp|any.
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to remove limits on the port range. Default value is `false`.
        /// </summary>
        [Output("portBreak")]
        public Output<bool?> PortBreak { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.119.1+) The status of forward entry.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ForwardEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ForwardEntry(string name, ForwardEntryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/forwardEntry:ForwardEntry", name, args ?? new ForwardEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ForwardEntry(string name, Input<string> id, ForwardEntryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/forwardEntry:ForwardEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ForwardEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ForwardEntry Get(string name, Input<string> id, ForwardEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new ForwardEntry(name, id, state, options);
        }
    }

    public sealed class ForwardEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The external ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Input("externalIp", required: true)]
        public Input<string> ExternalIp { get; set; } = null!;

        /// <summary>
        /// The external port, valid value is 1~65535|any.
        /// </summary>
        [Input("externalPort", required: true)]
        public Input<string> ExternalPort { get; set; } = null!;

        /// <summary>
        /// The name of forward entry.
        /// </summary>
        [Input("forwardEntryName")]
        public Input<string>? ForwardEntryName { get; set; }

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "forward_table_ids".
        /// </summary>
        [Input("forwardTableId", required: true)]
        public Input<string> ForwardTableId { get; set; } = null!;

        /// <summary>
        /// The internal ip, must a private ip.
        /// </summary>
        [Input("internalIp", required: true)]
        public Input<string> InternalIp { get; set; } = null!;

        /// <summary>
        /// The internal port, valid value is 1~65535|any.
        /// </summary>
        [Input("internalPort", required: true)]
        public Input<string> InternalPort { get; set; } = null!;

        /// <summary>
        /// The ip protocol, valid value is tcp|udp|any.
        /// </summary>
        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether to remove limits on the port range. Default value is `false`.
        /// </summary>
        [Input("portBreak")]
        public Input<bool>? PortBreak { get; set; }

        public ForwardEntryArgs()
        {
        }
        public static new ForwardEntryArgs Empty => new ForwardEntryArgs();
    }

    public sealed class ForwardEntryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The external ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Input("externalIp")]
        public Input<string>? ExternalIp { get; set; }

        /// <summary>
        /// The external port, valid value is 1~65535|any.
        /// </summary>
        [Input("externalPort")]
        public Input<string>? ExternalPort { get; set; }

        /// <summary>
        /// The id of the forward entry on the server.
        /// </summary>
        [Input("forwardEntryId")]
        public Input<string>? ForwardEntryId { get; set; }

        /// <summary>
        /// The name of forward entry.
        /// </summary>
        [Input("forwardEntryName")]
        public Input<string>? ForwardEntryName { get; set; }

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "forward_table_ids".
        /// </summary>
        [Input("forwardTableId")]
        public Input<string>? ForwardTableId { get; set; }

        /// <summary>
        /// The internal ip, must a private ip.
        /// </summary>
        [Input("internalIp")]
        public Input<string>? InternalIp { get; set; }

        /// <summary>
        /// The internal port, valid value is 1~65535|any.
        /// </summary>
        [Input("internalPort")]
        public Input<string>? InternalPort { get; set; }

        /// <summary>
        /// The ip protocol, valid value is tcp|udp|any.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.119.1. New field `forward_entry_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether to remove limits on the port range. Default value is `false`.
        /// </summary>
        [Input("portBreak")]
        public Input<bool>? PortBreak { get; set; }

        /// <summary>
        /// (Available in 1.119.1+) The status of forward entry.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ForwardEntryState()
        {
        }
        public static new ForwardEntryState Empty => new ForwardEntryState();
    }
}
