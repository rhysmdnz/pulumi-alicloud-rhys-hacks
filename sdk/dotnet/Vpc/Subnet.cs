// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    [Obsolete(@"This resource has been deprecated and replaced by the Switch resource.")]
    [AliCloudResourceType("alicloud:vpc/subnet:Subnet")]
    public partial class Subnet : global::Pulumi.CustomResource
    {
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("enableIpv6")]
        public Output<bool?> EnableIpv6 { get; private set; } = null!;

        [Output("ipv6CidrBlock")]
        public Output<string> Ipv6CidrBlock { get; private set; } = null!;

        [Output("ipv6CidrBlockMask")]
        public Output<int?> Ipv6CidrBlockMask { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        [Output("vswitchName")]
        public Output<string> VswitchName { get; private set; } = null!;

        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/subnet:Subnet", name, args ?? new SubnetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/subnet:Subnet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subnet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnet Get(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
        {
            return new Subnet(name, id, state, options);
        }
    }

    public sealed class SubnetArgs : global::Pulumi.ResourceArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        [Input("ipv6CidrBlockMask")]
        public Input<int>? Ipv6CidrBlockMask { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        [Input("vswitchName")]
        public Input<string>? VswitchName { get; set; }

        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public SubnetArgs()
        {
        }
        public static new SubnetArgs Empty => new SubnetArgs();
    }

    public sealed class SubnetState : global::Pulumi.ResourceArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        [Input("ipv6CidrBlockMask")]
        public Input<int>? Ipv6CidrBlockMask { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("vswitchName")]
        public Input<string>? VswitchName { get; set; }

        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public SubnetState()
        {
        }
        public static new SubnetState Empty => new SubnetState();
    }
}
