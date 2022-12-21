// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides a ECS Image Component resource.
    /// 
    /// For information about ECS Image Component and how to use it, see [What is Image Component](https://www.alibabacloud.com/help/en/doc-detail/200424.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.159.0+.
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
    ///     var @default = AliCloud.ResourceManager.GetResourceGroups.Invoke(new()
    ///     {
    ///         NameRegex = "default",
    ///     });
    /// 
    ///     var example = new AliCloud.Ecs.EcsImageComponent("example", new()
    ///     {
    ///         ComponentType = "Build",
    ///         Content = "RUN yum update -y",
    ///         Description = "example_value",
    ///         ImageComponentName = "example_value",
    ///         ResourceGroupId = @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult).Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Groups[0]?.Id)),
    ///         SystemType = "Linux",
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS Image Component can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ecs/ecsImageComponent:EcsImageComponent example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/ecsImageComponent:EcsImageComponent")]
    public partial class EcsImageComponent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of the image component. Only image building components are supported. Valid values: `Build`.
        /// </summary>
        [Output("componentType")]
        public Output<string> ComponentType { get; private set; } = null!;

        /// <summary>
        /// The content of the image component. The content can consist of up to 127 commands.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        /// </summary>
        [Output("imageComponentName")]
        public Output<string> ImageComponentName { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which to assign the image component.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
        /// </summary>
        [Output("systemType")]
        public Output<string> SystemType { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EcsImageComponent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EcsImageComponent(string name, EcsImageComponentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsImageComponent:EcsImageComponent", name, args ?? new EcsImageComponentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EcsImageComponent(string name, Input<string> id, EcsImageComponentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsImageComponent:EcsImageComponent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EcsImageComponent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EcsImageComponent Get(string name, Input<string> id, EcsImageComponentState? state = null, CustomResourceOptions? options = null)
        {
            return new EcsImageComponent(name, id, state, options);
        }
    }

    public sealed class EcsImageComponentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the image component. Only image building components are supported. Valid values: `Build`.
        /// </summary>
        [Input("componentType")]
        public Input<string>? ComponentType { get; set; }

        /// <summary>
        /// The content of the image component. The content can consist of up to 127 commands.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        /// </summary>
        [Input("imageComponentName")]
        public Input<string>? ImageComponentName { get; set; }

        /// <summary>
        /// The ID of the resource group to which to assign the image component.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
        /// </summary>
        [Input("systemType")]
        public Input<string>? SystemType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EcsImageComponentArgs()
        {
        }
        public static new EcsImageComponentArgs Empty => new EcsImageComponentArgs();
    }

    public sealed class EcsImageComponentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the image component. Only image building components are supported. Valid values: `Build`.
        /// </summary>
        [Input("componentType")]
        public Input<string>? ComponentType { get; set; }

        /// <summary>
        /// The content of the image component. The content can consist of up to 127 commands.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        /// </summary>
        [Input("imageComponentName")]
        public Input<string>? ImageComponentName { get; set; }

        /// <summary>
        /// The ID of the resource group to which to assign the image component.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
        /// </summary>
        [Input("systemType")]
        public Input<string>? SystemType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EcsImageComponentState()
        {
        }
        public static new EcsImageComponentState Empty => new EcsImageComponentState();
    }
}