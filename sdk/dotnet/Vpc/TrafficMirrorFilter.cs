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
    /// Provides a VPC Traffic Mirror Filter resource.
    /// 
    /// For information about VPC Traffic Mirror Filter and how to use it, see [What is Traffic Mirror Filter](https://www.alibabacloud.com/help/doc-detail/207513.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.140.0+.
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
    ///     var example = new AliCloud.Vpc.TrafficMirrorFilter("example", new()
    ///     {
    ///         TrafficMirrorFilterName = "example_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Traffic Mirror Filter can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/trafficMirrorFilter:TrafficMirrorFilter example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/trafficMirrorFilter:TrafficMirrorFilter")]
    public partial class TrafficMirrorFilter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The dry run.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The state of the filter. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`. `Creating`: The filter is being created. `Created`: The filter is created. `Modifying`: The filter is being modified. `Deleting`: The filter is being deleted.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The description of the filter. The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Output("trafficMirrorFilterDescription")]
        public Output<string?> TrafficMirrorFilterDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the filter. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Output("trafficMirrorFilterName")]
        public Output<string?> TrafficMirrorFilterName { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficMirrorFilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficMirrorFilter(string name, TrafficMirrorFilterArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/trafficMirrorFilter:TrafficMirrorFilter", name, args ?? new TrafficMirrorFilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficMirrorFilter(string name, Input<string> id, TrafficMirrorFilterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/trafficMirrorFilter:TrafficMirrorFilter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TrafficMirrorFilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficMirrorFilter Get(string name, Input<string> id, TrafficMirrorFilterState? state = null, CustomResourceOptions? options = null)
        {
            return new TrafficMirrorFilter(name, id, state, options);
        }
    }

    public sealed class TrafficMirrorFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The description of the filter. The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("trafficMirrorFilterDescription")]
        public Input<string>? TrafficMirrorFilterDescription { get; set; }

        /// <summary>
        /// The name of the filter. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("trafficMirrorFilterName")]
        public Input<string>? TrafficMirrorFilterName { get; set; }

        public TrafficMirrorFilterArgs()
        {
        }
        public static new TrafficMirrorFilterArgs Empty => new TrafficMirrorFilterArgs();
    }

    public sealed class TrafficMirrorFilterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The state of the filter. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`. `Creating`: The filter is being created. `Created`: The filter is created. `Modifying`: The filter is being modified. `Deleting`: The filter is being deleted.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The description of the filter. The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("trafficMirrorFilterDescription")]
        public Input<string>? TrafficMirrorFilterDescription { get; set; }

        /// <summary>
        /// The name of the filter. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("trafficMirrorFilterName")]
        public Input<string>? TrafficMirrorFilterName { get; set; }

        public TrafficMirrorFilterState()
        {
        }
        public static new TrafficMirrorFilterState Empty => new TrafficMirrorFilterState();
    }
}
