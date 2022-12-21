// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ddos
{
    public static class GetDdosBgpIps
    {
        /// <summary>
        /// This data source provides the Ddos Bgp Ips of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.180.0+.
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
        ///     var ids = AliCloud.Ddos.GetDdosBgpIps.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ddosbgpIpId1"] = ids.Apply(getDdosBgpIpsResult =&gt; getDdosBgpIpsResult.Ips[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDdosBgpIpsResult> InvokeAsync(GetDdosBgpIpsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDdosBgpIpsResult>("alicloud:ddos/getDdosBgpIps:getDdosBgpIps", args ?? new GetDdosBgpIpsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ddos Bgp Ips of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.180.0+.
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
        ///     var ids = AliCloud.Ddos.GetDdosBgpIps.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ddosbgpIpId1"] = ids.Apply(getDdosBgpIpsResult =&gt; getDdosBgpIpsResult.Ips[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDdosBgpIpsResult> Invoke(GetDdosBgpIpsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDdosBgpIpsResult>("alicloud:ddos/getDdosBgpIps:getDdosBgpIps", args ?? new GetDdosBgpIpsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDdosBgpIpsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Ip IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the native protection enterprise instance to be operated.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The product name. Valid Value:`ECS`, `SLB`, `EIP`, `WAF`.
        /// </summary>
        [Input("productName")]
        public string? ProductName { get; set; }

        /// <summary>
        /// The current state of the IP address.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetDdosBgpIpsArgs()
        {
        }
        public static new GetDdosBgpIpsArgs Empty => new GetDdosBgpIpsArgs();
    }

    public sealed class GetDdosBgpIpsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Ip IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the native protection enterprise instance to be operated.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The product name. Valid Value:`ECS`, `SLB`, `EIP`, `WAF`.
        /// </summary>
        [Input("productName")]
        public Input<string>? ProductName { get; set; }

        /// <summary>
        /// The current state of the IP address.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetDdosBgpIpsInvokeArgs()
        {
        }
        public static new GetDdosBgpIpsInvokeArgs Empty => new GetDdosBgpIpsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDdosBgpIpsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string InstanceId;
        public readonly ImmutableArray<Outputs.GetDdosBgpIpsIpResult> Ips;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? ProductName;
        public readonly string? Status;

        [OutputConstructor]
        private GetDdosBgpIpsResult(
            string id,

            ImmutableArray<string> ids,

            string instanceId,

            ImmutableArray<Outputs.GetDdosBgpIpsIpResult> ips,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? productName,

            string? status)
        {
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            Ips = ips;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            ProductName = productName;
            Status = status;
        }
    }
}