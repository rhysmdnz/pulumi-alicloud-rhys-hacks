// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static class GetRules
    {
        /// <summary>
        /// This data source provides the rules associated with a server load balancer listener.
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
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "slbrulebasicconfig";
        ///     var defaultZones = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableDiskCategory = "cloud_efficiency",
        ///         AvailableResourceCreation = "VSwitch",
        ///     });
        /// 
        ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
        ///     {
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
        ///     {
        ///         VpcId = defaultNetwork.Id,
        ///         CidrBlock = "172.16.0.0/16",
        ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VswitchName = name,
        ///     });
        /// 
        ///     var defaultApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", new()
        ///     {
        ///         LoadBalancerName = name,
        ///         VswitchId = defaultSwitch.Id,
        ///     });
        /// 
        ///     var defaultListener = new AliCloud.Slb.Listener("defaultListener", new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///         BackendPort = 22,
        ///         FrontendPort = 22,
        ///         Protocol = "http",
        ///         Bandwidth = 5,
        ///         HealthCheckConnectPort = 20,
        ///     });
        /// 
        ///     var defaultServerGroup = new AliCloud.Slb.ServerGroup("defaultServerGroup", new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///     });
        /// 
        ///     var defaultRule = new AliCloud.Slb.Rule("defaultRule", new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///         FrontendPort = defaultListener.FrontendPort,
        ///         Domain = "*.aliyun.com",
        ///         Url = "/image",
        ///         ServerGroupId = defaultServerGroup.Id,
        ///     });
        /// 
        ///     var sampleDs = AliCloud.Slb.GetRules.Invoke(new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///         FrontendPort = 22,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSlbRuleId"] = sampleDs.Apply(getRulesResult =&gt; getRulesResult.SlbRules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRulesResult> InvokeAsync(GetRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRulesResult>("alicloud:slb/getRules:getRules", args ?? new GetRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the rules associated with a server load balancer listener.
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
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "slbrulebasicconfig";
        ///     var defaultZones = AliCloud.GetZones.Invoke(new()
        ///     {
        ///         AvailableDiskCategory = "cloud_efficiency",
        ///         AvailableResourceCreation = "VSwitch",
        ///     });
        /// 
        ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
        ///     {
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
        ///     {
        ///         VpcId = defaultNetwork.Id,
        ///         CidrBlock = "172.16.0.0/16",
        ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VswitchName = name,
        ///     });
        /// 
        ///     var defaultApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", new()
        ///     {
        ///         LoadBalancerName = name,
        ///         VswitchId = defaultSwitch.Id,
        ///     });
        /// 
        ///     var defaultListener = new AliCloud.Slb.Listener("defaultListener", new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///         BackendPort = 22,
        ///         FrontendPort = 22,
        ///         Protocol = "http",
        ///         Bandwidth = 5,
        ///         HealthCheckConnectPort = 20,
        ///     });
        /// 
        ///     var defaultServerGroup = new AliCloud.Slb.ServerGroup("defaultServerGroup", new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///     });
        /// 
        ///     var defaultRule = new AliCloud.Slb.Rule("defaultRule", new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///         FrontendPort = defaultListener.FrontendPort,
        ///         Domain = "*.aliyun.com",
        ///         Url = "/image",
        ///         ServerGroupId = defaultServerGroup.Id,
        ///     });
        /// 
        ///     var sampleDs = AliCloud.Slb.GetRules.Invoke(new()
        ///     {
        ///         LoadBalancerId = defaultApplicationLoadBalancer.Id,
        ///         FrontendPort = 22,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSlbRuleId"] = sampleDs.Apply(getRulesResult =&gt; getRulesResult.SlbRules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRulesResult> Invoke(GetRulesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRulesResult>("alicloud:slb/getRules:getRules", args ?? new GetRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// SLB listener port.
        /// </summary>
        [Input("frontendPort", required: true)]
        public int FrontendPort { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of rules IDs to filter results.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of the SLB with listener rules.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public string LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetRulesArgs()
        {
        }
        public static new GetRulesArgs Empty => new GetRulesArgs();
    }

    public sealed class GetRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// SLB listener port.
        /// </summary>
        [Input("frontendPort", required: true)]
        public Input<int> FrontendPort { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of rules IDs to filter results.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of the SLB with listener rules.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetRulesInvokeArgs()
        {
        }
        public static new GetRulesInvokeArgs Empty => new GetRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRulesResult
    {
        public readonly int FrontendPort;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of SLB listener rules IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string LoadBalancerId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of SLB listener rules names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of SLB listener rules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRulesSlbRuleResult> SlbRules;

        [OutputConstructor]
        private GetRulesResult(
            int frontendPort,

            string id,

            ImmutableArray<string> ids,

            string loadBalancerId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetRulesSlbRuleResult> slbRules)
        {
            FrontendPort = frontendPort;
            Id = id;
            Ids = ids;
            LoadBalancerId = loadBalancerId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            SlbRules = slbRules;
        }
    }
}