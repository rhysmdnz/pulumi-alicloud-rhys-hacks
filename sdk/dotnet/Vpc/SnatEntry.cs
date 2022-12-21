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
    /// Provides a snat resource.
    /// 
    /// ## Import
    /// 
    /// Snat Entry can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/snatEntry:SnatEntry foo stb-1aece3:snat-232ce2
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/snatEntry:SnatEntry")]
    public partial class SnatEntry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the snat entry on the server.
        /// </summary>
        [Output("snatEntryId")]
        public Output<string> SnatEntryId { get; private set; } = null!;

        /// <summary>
        /// The name of snat entry.
        /// </summary>
        [Output("snatEntryName")]
        public Output<string?> SnatEntryName { get; private set; } = null!;

        /// <summary>
        /// The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Output("snatIp")]
        public Output<string> SnatIp { get; private set; } = null!;

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "snat_table_ids".
        /// </summary>
        [Output("snatTableId")]
        public Output<string> SnatTableId { get; private set; } = null!;

        /// <summary>
        /// The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        /// </summary>
        [Output("sourceCidr")]
        public Output<string> SourceCidr { get; private set; } = null!;

        /// <summary>
        /// The vswitch ID.
        /// </summary>
        [Output("sourceVswitchId")]
        public Output<string> SourceVswitchId { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.119.1+) The status of snat entry.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SnatEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnatEntry(string name, SnatEntryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/snatEntry:SnatEntry", name, args ?? new SnatEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnatEntry(string name, Input<string> id, SnatEntryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/snatEntry:SnatEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnatEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnatEntry Get(string name, Input<string> id, SnatEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new SnatEntry(name, id, state, options);
        }
    }

    public sealed class SnatEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of snat entry.
        /// </summary>
        [Input("snatEntryName")]
        public Input<string>? SnatEntryName { get; set; }

        /// <summary>
        /// The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Input("snatIp", required: true)]
        public Input<string> SnatIp { get; set; } = null!;

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "snat_table_ids".
        /// </summary>
        [Input("snatTableId", required: true)]
        public Input<string> SnatTableId { get; set; } = null!;

        /// <summary>
        /// The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        /// </summary>
        [Input("sourceCidr")]
        public Input<string>? SourceCidr { get; set; }

        /// <summary>
        /// The vswitch ID.
        /// </summary>
        [Input("sourceVswitchId")]
        public Input<string>? SourceVswitchId { get; set; }

        public SnatEntryArgs()
        {
        }
        public static new SnatEntryArgs Empty => new SnatEntryArgs();
    }

    public sealed class SnatEntryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the snat entry on the server.
        /// </summary>
        [Input("snatEntryId")]
        public Input<string>? SnatEntryId { get; set; }

        /// <summary>
        /// The name of snat entry.
        /// </summary>
        [Input("snatEntryName")]
        public Input<string>? SnatEntryName { get; set; }

        /// <summary>
        /// The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidth_packages`.
        /// </summary>
        [Input("snatIp")]
        public Input<string>? SnatIp { get; set; }

        /// <summary>
        /// The value can get from `alicloud.vpc.NatGateway` Attributes "snat_table_ids".
        /// </summary>
        [Input("snatTableId")]
        public Input<string>? SnatTableId { get; set; }

        /// <summary>
        /// The private network segment of Ecs. This parameter and the `source_vswitch_id` parameter are mutually exclusive and cannot appear at the same time.
        /// </summary>
        [Input("sourceCidr")]
        public Input<string>? SourceCidr { get; set; }

        /// <summary>
        /// The vswitch ID.
        /// </summary>
        [Input("sourceVswitchId")]
        public Input<string>? SourceVswitchId { get; set; }

        /// <summary>
        /// (Available in 1.119.1+) The status of snat entry.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SnatEntryState()
        {
        }
        public static new SnatEntryState Empty => new SnatEntryState();
    }
}