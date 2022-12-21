// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static class GetBandwidthPackages
    {
        /// <summary>
        /// This data source provides CEN Bandwidth Packages available to the user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AliCloud.Cen.GetBandwidthPackages.Invoke(new()
        ///     {
        ///         InstanceId = "cen-id1",
        ///         NameRegex = "^foo",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCenBandwidthPackageId"] = example.Apply(getBandwidthPackagesResult =&gt; getBandwidthPackagesResult.Packages[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBandwidthPackagesResult> InvokeAsync(GetBandwidthPackagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBandwidthPackagesResult>("alicloud:cen/getBandwidthPackages:getBandwidthPackages", args ?? new GetBandwidthPackagesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides CEN Bandwidth Packages available to the user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AliCloud.Cen.GetBandwidthPackages.Invoke(new()
        ///     {
        ///         InstanceId = "cen-id1",
        ///         NameRegex = "^foo",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstCenBandwidthPackageId"] = example.Apply(getBandwidthPackagesResult =&gt; getBandwidthPackagesResult.Packages[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBandwidthPackagesResult> Invoke(GetBandwidthPackagesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBandwidthPackagesResult>("alicloud:cen/getBandwidthPackages:getBandwidthPackages", args ?? new GetBandwidthPackagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBandwidthPackagesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// Limit search to a list of specific CEN Bandwidth Package IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// -Indicates whether to include renewal data. Valid values: `true`: Return renewal data in the response. `false`: Do not return renewal data in the response.
        /// </summary>
        [Input("includeReservationData")]
        public bool? IncludeReservationData { get; set; }

        /// <summary>
        /// ID of a CEN instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// A regex string to filter CEN Bandwidth Package by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Status of the CEN Bandwidth Package in CEN instance, Valid value: `Idle` and `InUse`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetBandwidthPackagesArgs()
        {
        }
        public static new GetBandwidthPackagesArgs Empty => new GetBandwidthPackagesArgs();
    }

    public sealed class GetBandwidthPackagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// Limit search to a list of specific CEN Bandwidth Package IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// -Indicates whether to include renewal data. Valid values: `true`: Return renewal data in the response. `false`: Do not return renewal data in the response.
        /// </summary>
        [Input("includeReservationData")]
        public Input<bool>? IncludeReservationData { get; set; }

        /// <summary>
        /// ID of a CEN instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// A regex string to filter CEN Bandwidth Package by name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Status of the CEN Bandwidth Package in CEN instance, Valid value: `Idle` and `InUse`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetBandwidthPackagesInvokeArgs()
        {
        }
        public static new GetBandwidthPackagesInvokeArgs Empty => new GetBandwidthPackagesInvokeArgs();
    }


    [OutputType]
    public sealed class GetBandwidthPackagesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of specific CEN Bandwidth Package IDs.
        /// * `names` (Available in 1.98.0+) - A list of CEN Bandwidth Package Names.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly bool? IncludeReservationData;
        /// <summary>
        /// The ID of the CEN instance that are associated with the bandwidth package.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of CEN bandwidth package. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBandwidthPackagesPackageResult> Packages;
        /// <summary>
        /// Status of the CEN Bandwidth Package in CEN instance, including `Idle` and `InUse`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetBandwidthPackagesResult(
            string id,

            ImmutableArray<string> ids,

            bool? includeReservationData,

            string? instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetBandwidthPackagesPackageResult> packages,

            string? status)
        {
            Id = id;
            Ids = ids;
            IncludeReservationData = includeReservationData;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Packages = packages;
            Status = status;
        }
    }
}