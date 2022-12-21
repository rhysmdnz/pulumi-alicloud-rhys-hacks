// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud
{
    /// <summary>
    /// Provides a Msc Sub Webhook resource.
    /// 
    /// &gt; **NOTE:** Available in v1.141.0+.
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
    ///     var example = new AliCloud.MscSubWebhook("example", new()
    ///     {
    ///         ServerUrl = "example_value",
    ///         WebhookName = "example_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Msc Sub Webhook can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:index/mscSubWebhook:MscSubWebhook example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:index/mscSubWebhook:MscSubWebhook")]
    public partial class MscSubWebhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The serverUrl of the Webhook. This url must start with `https://oapi.dingtalk.com/robot/send?access_token=`.
        /// </summary>
        [Output("serverUrl")]
        public Output<string> ServerUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the Webhook. **Note:** The name must be `2` to `12` characters in length, and can contain uppercase and lowercase letters.
        /// </summary>
        [Output("webhookName")]
        public Output<string> WebhookName { get; private set; } = null!;


        /// <summary>
        /// Create a MscSubWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MscSubWebhook(string name, MscSubWebhookArgs args, CustomResourceOptions? options = null)
            : base("alicloud:index/mscSubWebhook:MscSubWebhook", name, args ?? new MscSubWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MscSubWebhook(string name, Input<string> id, MscSubWebhookState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:index/mscSubWebhook:MscSubWebhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MscSubWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MscSubWebhook Get(string name, Input<string> id, MscSubWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new MscSubWebhook(name, id, state, options);
        }
    }

    public sealed class MscSubWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The serverUrl of the Webhook. This url must start with `https://oapi.dingtalk.com/robot/send?access_token=`.
        /// </summary>
        [Input("serverUrl", required: true)]
        public Input<string> ServerUrl { get; set; } = null!;

        /// <summary>
        /// The name of the Webhook. **Note:** The name must be `2` to `12` characters in length, and can contain uppercase and lowercase letters.
        /// </summary>
        [Input("webhookName", required: true)]
        public Input<string> WebhookName { get; set; } = null!;

        public MscSubWebhookArgs()
        {
        }
        public static new MscSubWebhookArgs Empty => new MscSubWebhookArgs();
    }

    public sealed class MscSubWebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The serverUrl of the Webhook. This url must start with `https://oapi.dingtalk.com/robot/send?access_token=`.
        /// </summary>
        [Input("serverUrl")]
        public Input<string>? ServerUrl { get; set; }

        /// <summary>
        /// The name of the Webhook. **Note:** The name must be `2` to `12` characters in length, and can contain uppercase and lowercase letters.
        /// </summary>
        [Input("webhookName")]
        public Input<string>? WebhookName { get; set; }

        public MscSubWebhookState()
        {
        }
        public static new MscSubWebhookState Empty => new MscSubWebhookState();
    }
}