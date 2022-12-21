// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ros
{
    public static class GetRegions
    {
        /// <summary>
        /// This data source provides the Ros Regions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.145.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var all = AliCloud.Ros.GetRegions.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["rosRegionRegionId1"] = all.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.RegionId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegionsResult> InvokeAsync(GetRegionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionsResult>("alicloud:ros/getRegions:getRegions", args ?? new GetRegionsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ros Regions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.145.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var all = AliCloud.Ros.GetRegions.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["rosRegionRegionId1"] = all.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.RegionId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegionsResult> Invoke(GetRegionsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionsResult>("alicloud:ros/getRegions:getRegions", args ?? new GetRegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetRegionsArgs()
        {
        }
        public static new GetRegionsArgs Empty => new GetRegionsArgs();
    }

    public sealed class GetRegionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetRegionsInvokeArgs()
        {
        }
        public static new GetRegionsInvokeArgs Empty => new GetRegionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetRegionsRegionResult> Regions;

        [OutputConstructor]
        private GetRegionsResult(
            string id,

            string? outputFile,

            ImmutableArray<Outputs.GetRegionsRegionResult> regions)
        {
            Id = id;
            OutputFile = outputFile;
            Regions = regions;
        }
    }
}