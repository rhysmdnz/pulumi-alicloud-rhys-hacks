// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos
{
    public static class GetApplicationGroups
    {
        /// <summary>
        /// This data source provides the Oos Application Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.146.0+.
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
        ///     var ids = AliCloud.Oos.GetApplicationGroups.Invoke(new()
        ///     {
        ///         ApplicationName = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "my-ApplicationGroup-1",
        ///             "my-ApplicationGroup-2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Oos.GetApplicationGroups.Invoke(new()
        ///     {
        ///         ApplicationName = "example_value",
        ///         NameRegex = "^my-ApplicationGroup",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["oosApplicationGroupId1"] = ids.Apply(getApplicationGroupsResult =&gt; getApplicationGroupsResult.Groups[0]?.Id),
        ///         ["oosApplicationGroupId2"] = nameRegex.Apply(getApplicationGroupsResult =&gt; getApplicationGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationGroupsResult> InvokeAsync(GetApplicationGroupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationGroupsResult>("alicloud:oos/getApplicationGroups:getApplicationGroups", args ?? new GetApplicationGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Oos Application Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.146.0+.
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
        ///     var ids = AliCloud.Oos.GetApplicationGroups.Invoke(new()
        ///     {
        ///         ApplicationName = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "my-ApplicationGroup-1",
        ///             "my-ApplicationGroup-2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Oos.GetApplicationGroups.Invoke(new()
        ///     {
        ///         ApplicationName = "example_value",
        ///         NameRegex = "^my-ApplicationGroup",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["oosApplicationGroupId1"] = ids.Apply(getApplicationGroupsResult =&gt; getApplicationGroupsResult.Groups[0]?.Id),
        ///         ["oosApplicationGroupId2"] = nameRegex.Apply(getApplicationGroupsResult =&gt; getApplicationGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationGroupsResult> Invoke(GetApplicationGroupsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApplicationGroupsResult>("alicloud:oos/getApplicationGroups:getApplicationGroups", args ?? new GetApplicationGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Application.
        /// </summary>
        [Input("applicationName", required: true)]
        public string ApplicationName { get; set; } = null!;

        /// <summary>
        /// The region ID of the deployment.
        /// </summary>
        [Input("deployRegionId")]
        public string? DeployRegionId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Application Group IDs. Its element value is same as Application Group Name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Application Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetApplicationGroupsArgs()
        {
        }
        public static new GetApplicationGroupsArgs Empty => new GetApplicationGroupsArgs();
    }

    public sealed class GetApplicationGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Application.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// The region ID of the deployment.
        /// </summary>
        [Input("deployRegionId")]
        public Input<string>? DeployRegionId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Application Group IDs. Its element value is same as Application Group Name.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Application Group name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetApplicationGroupsInvokeArgs()
        {
        }
        public static new GetApplicationGroupsInvokeArgs Empty => new GetApplicationGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationGroupsResult
    {
        public readonly string ApplicationName;
        public readonly string? DeployRegionId;
        public readonly ImmutableArray<Outputs.GetApplicationGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetApplicationGroupsResult(
            string applicationName,

            string? deployRegionId,

            ImmutableArray<Outputs.GetApplicationGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            ApplicationName = applicationName;
            DeployRegionId = deployRegionId;
            Groups = groups;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}