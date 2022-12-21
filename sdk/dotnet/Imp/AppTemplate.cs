// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Imp
{
    /// <summary>
    /// Provides a Apsara Agile Live (IMP) App Template resource.
    /// 
    /// For information about Apsara Agile Live (IMP) App Template and how to use it, see [What is App Template](https://help.aliyun.com/document_detail/270121.html).
    /// 
    /// &gt; **NOTE:** Available in v1.137.0+.
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
    ///     var example = new AliCloud.Imp.AppTemplate("example", new()
    ///     {
    ///         AppTemplateName = "example_value",
    ///         ComponentLists = new[]
    ///         {
    ///             "component.live",
    ///             "component.liveRecord",
    ///         },
    ///         IntegrationMode = "paasSDK",
    ///         Scene = "business",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Apsara Agile Live (IMP) App Template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:imp/appTemplate:AppTemplate example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:imp/appTemplate:AppTemplate")]
    public partial class AppTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("appTemplateName")]
        public Output<string> AppTemplateName { get; private set; } = null!;

        /// <summary>
        /// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
        /// </summary>
        [Output("componentLists")]
        public Output<ImmutableArray<string>> ComponentLists { get; private set; } = null!;

        /// <summary>
        /// Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
        /// </summary>
        [Output("configLists")]
        public Output<ImmutableArray<Outputs.AppTemplateConfigList>> ConfigLists { get; private set; } = null!;

        /// <summary>
        /// Integration mode. Valid values:
        /// * paasSDK: Integrated SDK.
        /// * standardRoom: Model Room.
        /// </summary>
        [Output("integrationMode")]
        public Output<string?> IntegrationMode { get; private set; } = null!;

        /// <summary>
        /// Application Template scenario. Valid values: ["business", "classroom"].
        /// </summary>
        [Output("scene")]
        public Output<string?> Scene { get; private set; } = null!;

        /// <summary>
        /// Application template usage status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AppTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppTemplate(string name, AppTemplateArgs args, CustomResourceOptions? options = null)
            : base("alicloud:imp/appTemplate:AppTemplate", name, args ?? new AppTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppTemplate(string name, Input<string> id, AppTemplateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:imp/appTemplate:AppTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppTemplate Get(string name, Input<string> id, AppTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new AppTemplate(name, id, state, options);
        }
    }

    public sealed class AppTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("appTemplateName", required: true)]
        public Input<string> AppTemplateName { get; set; } = null!;

        [Input("componentLists", required: true)]
        private InputList<string>? _componentLists;

        /// <summary>
        /// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
        /// </summary>
        public InputList<string> ComponentLists
        {
            get => _componentLists ?? (_componentLists = new InputList<string>());
            set => _componentLists = value;
        }

        [Input("configLists")]
        private InputList<Inputs.AppTemplateConfigListArgs>? _configLists;

        /// <summary>
        /// Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
        /// </summary>
        public InputList<Inputs.AppTemplateConfigListArgs> ConfigLists
        {
            get => _configLists ?? (_configLists = new InputList<Inputs.AppTemplateConfigListArgs>());
            set => _configLists = value;
        }

        /// <summary>
        /// Integration mode. Valid values:
        /// * paasSDK: Integrated SDK.
        /// * standardRoom: Model Room.
        /// </summary>
        [Input("integrationMode")]
        public Input<string>? IntegrationMode { get; set; }

        /// <summary>
        /// Application Template scenario. Valid values: ["business", "classroom"].
        /// </summary>
        [Input("scene")]
        public Input<string>? Scene { get; set; }

        public AppTemplateArgs()
        {
        }
        public static new AppTemplateArgs Empty => new AppTemplateArgs();
    }

    public sealed class AppTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("appTemplateName")]
        public Input<string>? AppTemplateName { get; set; }

        [Input("componentLists")]
        private InputList<string>? _componentLists;

        /// <summary>
        /// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
        /// </summary>
        public InputList<string> ComponentLists
        {
            get => _componentLists ?? (_componentLists = new InputList<string>());
            set => _componentLists = value;
        }

        [Input("configLists")]
        private InputList<Inputs.AppTemplateConfigListGetArgs>? _configLists;

        /// <summary>
        /// Configuration list. It have several default configs after the resource is created. See the following `Block config_list`.
        /// </summary>
        public InputList<Inputs.AppTemplateConfigListGetArgs> ConfigLists
        {
            get => _configLists ?? (_configLists = new InputList<Inputs.AppTemplateConfigListGetArgs>());
            set => _configLists = value;
        }

        /// <summary>
        /// Integration mode. Valid values:
        /// * paasSDK: Integrated SDK.
        /// * standardRoom: Model Room.
        /// </summary>
        [Input("integrationMode")]
        public Input<string>? IntegrationMode { get; set; }

        /// <summary>
        /// Application Template scenario. Valid values: ["business", "classroom"].
        /// </summary>
        [Input("scene")]
        public Input<string>? Scene { get; set; }

        /// <summary>
        /// Application template usage status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AppTemplateState()
        {
        }
        public static new AppTemplateState Empty => new AppTemplateState();
    }
}