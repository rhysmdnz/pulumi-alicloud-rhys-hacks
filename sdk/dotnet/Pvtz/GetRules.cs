// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Pvtz
{
    public static class GetRules
    {
        /// <summary>
        /// This data source provides the PrivateZone Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.143.0+.
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
        ///     var ids = AliCloud.Pvtz.GetRules.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Pvtz.GetRules.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Rule",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["pvtzRuleId1"] = ids.Apply(getRulesResult =&gt; getRulesResult.Rules[0]?.Id),
        ///         ["pvtzRuleId2"] = nameRegex.Apply(getRulesResult =&gt; getRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRulesResult> InvokeAsync(GetRulesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRulesResult>("alicloud:pvtz/getRules:getRules", args ?? new GetRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the PrivateZone Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.143.0+.
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
        ///     var ids = AliCloud.Pvtz.GetRules.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Pvtz.GetRules.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Rule",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["pvtzRuleId1"] = ids.Apply(getRulesResult =&gt; getRulesResult.Rules[0]?.Id),
        ///         ["pvtzRuleId2"] = nameRegex.Apply(getRulesResult =&gt; getRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRulesResult> Invoke(GetRulesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRulesResult>("alicloud:pvtz/getRules:getRules", args ?? new GetRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Endpoint.
        /// </summary>
        [Input("endpointId")]
        public string? EndpointId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Rule name.
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
        /// The ID of the Endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Rule name.
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
        public readonly string? EndpointId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetRulesRuleResult> Rules;

        [OutputConstructor]
        private GetRulesResult(
            string? endpointId,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetRulesRuleResult> rules)
        {
            EndpointId = endpointId;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Rules = rules;
        }
    }
}