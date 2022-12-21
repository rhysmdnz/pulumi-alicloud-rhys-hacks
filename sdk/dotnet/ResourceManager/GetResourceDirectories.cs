// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager
{
    public static class GetResourceDirectories
    {
        /// <summary>
        /// This data source provides the Resource Manager Resource Directories of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.86.0+.
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
        ///     var @default = AliCloud.ResourceManager.GetResourceDirectories.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["resourceDirectoryId"] = @default.Apply(getResourceDirectoriesResult =&gt; getResourceDirectoriesResult).Apply(@default =&gt; @default.Apply(getResourceDirectoriesResult =&gt; getResourceDirectoriesResult.Directories[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResourceDirectoriesResult> InvokeAsync(GetResourceDirectoriesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceDirectoriesResult>("alicloud:resourcemanager/getResourceDirectories:getResourceDirectories", args ?? new GetResourceDirectoriesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Resource Manager Resource Directories of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.86.0+.
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
        ///     var @default = AliCloud.ResourceManager.GetResourceDirectories.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["resourceDirectoryId"] = @default.Apply(getResourceDirectoriesResult =&gt; getResourceDirectoriesResult).Apply(@default =&gt; @default.Apply(getResourceDirectoriesResult =&gt; getResourceDirectoriesResult.Directories[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetResourceDirectoriesResult> Invoke(GetResourceDirectoriesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourceDirectoriesResult>("alicloud:resourcemanager/getResourceDirectories:getResourceDirectories", args ?? new GetResourceDirectoriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceDirectoriesArgs : global::Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetResourceDirectoriesArgs()
        {
        }
        public static new GetResourceDirectoriesArgs Empty => new GetResourceDirectoriesArgs();
    }

    public sealed class GetResourceDirectoriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetResourceDirectoriesInvokeArgs()
        {
        }
        public static new GetResourceDirectoriesInvokeArgs Empty => new GetResourceDirectoriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceDirectoriesResult
    {
        /// <summary>
        /// A list of resource directories. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourceDirectoriesDirectoryResult> Directories;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetResourceDirectoriesResult(
            ImmutableArray<Outputs.GetResourceDirectoriesDirectoryResult> directories,

            string id,

            string? outputFile)
        {
            Directories = directories;
            Id = id;
            OutputFile = outputFile;
        }
    }
}