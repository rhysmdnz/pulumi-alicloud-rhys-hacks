// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetBgpPeers
    {
        /// <summary>
        /// This data source provides the Vpc Bgp Peers of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.153.0+.
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
        ///     var ids = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var bgpGroupId = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         BgpGroupId = "example_value",
        ///     });
        /// 
        ///     var routerId = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         RouterId = "example_value",
        ///     });
        /// 
        ///     var status = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         Status = "Available",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpcBgpPeerId1"] = ids.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///         ["vpcBgpPeerId2"] = bgpGroupId.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///         ["vpcBgpPeerId3"] = routerId.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///         ["vpcBgpPeerId4"] = status.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBgpPeersResult> InvokeAsync(GetBgpPeersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBgpPeersResult>("alicloud:vpc/getBgpPeers:getBgpPeers", args ?? new GetBgpPeersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Vpc Bgp Peers of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.153.0+.
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
        ///     var ids = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var bgpGroupId = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         BgpGroupId = "example_value",
        ///     });
        /// 
        ///     var routerId = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         RouterId = "example_value",
        ///     });
        /// 
        ///     var status = AliCloud.Vpc.GetBgpPeers.Invoke(new()
        ///     {
        ///         Status = "Available",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpcBgpPeerId1"] = ids.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///         ["vpcBgpPeerId2"] = bgpGroupId.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///         ["vpcBgpPeerId3"] = routerId.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///         ["vpcBgpPeerId4"] = status.Apply(getBgpPeersResult =&gt; getBgpPeersResult.Peers[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBgpPeersResult> Invoke(GetBgpPeersInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBgpPeersResult>("alicloud:vpc/getBgpPeers:getBgpPeers", args ?? new GetBgpPeersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBgpPeersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the BGP group.
        /// </summary>
        [Input("bgpGroupId")]
        public string? BgpGroupId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Bgp Peer IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the router.
        /// </summary>
        [Input("routerId")]
        public string? RouterId { get; set; }

        /// <summary>
        /// The status of the BGP peer.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetBgpPeersArgs()
        {
        }
        public static new GetBgpPeersArgs Empty => new GetBgpPeersArgs();
    }

    public sealed class GetBgpPeersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the BGP group.
        /// </summary>
        [Input("bgpGroupId")]
        public Input<string>? BgpGroupId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Bgp Peer IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the router.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// The status of the BGP peer.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetBgpPeersInvokeArgs()
        {
        }
        public static new GetBgpPeersInvokeArgs Empty => new GetBgpPeersInvokeArgs();
    }


    [OutputType]
    public sealed class GetBgpPeersResult
    {
        public readonly string? BgpGroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetBgpPeersPeerResult> Peers;
        public readonly string? RouterId;
        public readonly string? Status;

        [OutputConstructor]
        private GetBgpPeersResult(
            string? bgpGroupId,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetBgpPeersPeerResult> peers,

            string? routerId,

            string? status)
        {
            BgpGroupId = bgpGroupId;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Peers = peers;
            RouterId = routerId;
            Status = status;
        }
    }
}