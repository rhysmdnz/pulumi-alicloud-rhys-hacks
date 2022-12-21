// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MongoDB
{
    /// <summary>
    /// Provides a MongoDB Account resource.
    /// 
    /// For information about MongoDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/doc-detail/62154.html).
    /// 
    /// &gt; **NOTE:** Available in v1.148.0+.
    /// 
    /// ## Import
    /// 
    /// MongoDB Account can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:mongodb/account:Account example &lt;instance_id&gt;:&lt;account_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:mongodb/account:Account")]
    public partial class Account : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the account.
        /// * The description must start with a letter, and cannot start with `http://` or `https://`.
        /// * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        /// </summary>
        [Output("accountDescription")]
        public Output<string?> AccountDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the account. Valid values: `root`.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The Password of the Account.
        /// * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&amp;*()_+-=`.
        /// * The password must be `8` to `32` characters in length.
        /// </summary>
        [Output("accountPassword")]
        public Output<string> AccountPassword { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The status of the account. Valid values: `Unavailable`, `Available`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("alicloud:mongodb/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:mongodb/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the account.
        /// * The description must start with a letter, and cannot start with `http://` or `https://`.
        /// * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// The name of the account. Valid values: `root`.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The Password of the Account.
        /// * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&amp;*()_+-=`.
        /// * The password must be `8` to `32` characters in length.
        /// </summary>
        [Input("accountPassword", required: true)]
        public Input<string> AccountPassword { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AccountArgs()
        {
        }
        public static new AccountArgs Empty => new AccountArgs();
    }

    public sealed class AccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the account.
        /// * The description must start with a letter, and cannot start with `http://` or `https://`.
        /// * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// The name of the account. Valid values: `root`.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The Password of the Account.
        /// * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&amp;*()_+-=`.
        /// * The password must be `8` to `32` characters in length.
        /// </summary>
        [Input("accountPassword")]
        public Input<string>? AccountPassword { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The status of the account. Valid values: `Unavailable`, `Available`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AccountState()
        {
        }
        public static new AccountState Empty => new AccountState();
    }
}