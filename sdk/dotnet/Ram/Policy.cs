// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new RAM Policy.
    ///     var policy = new AliCloud.Ram.Policy("policy", new()
    ///     {
    ///         Description = "this is a policy test",
    ///         Force = true,
    ///         PolicyDocument = @"  {
    ///     ""Statement"": [
    ///       {
    ///         ""Action"": [
    ///           ""oss:ListObjects"",
    ///           ""oss:GetObject""
    ///         ],
    ///         ""Effect"": ""Allow"",
    ///         ""Resource"": [
    ///           ""acs:oss:*:*:mybucket"",
    ///           ""acs:oss:*:*:mybucket/*""
    ///         ]
    ///       }
    ///     ],
    ///       ""Version"": ""1""
    ///   }
    ///   
    /// ",
    ///         PolicyName = "policyName",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RAM policy can be imported using the id or name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ram/policy:Policy example my-policy
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ram/policy:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The policy attachment count.
        /// </summary>
        [Output("attachmentCount")]
        public Output<int> AttachmentCount { get; private set; } = null!;

        /// <summary>
        /// The default version of policy.
        /// </summary>
        [Output("defaultVersion")]
        public Output<string> DefaultVersion { get; private set; } = null!;

        /// <summary>
        /// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from provider version 1.114.0 and `policy_document` instead.
        /// </summary>
        [Output("document")]
        public Output<string> Document { get; private set; } = null!;

        /// <summary>
        /// This parameter is used for resource destroy. Default value is `false`.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from provider version 1.114.0 and `policy_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Document of the RAM policy. It is required when the `statement` is not specified.
        /// </summary>
        [Output("policyDocument")]
        public Output<string> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
        /// </summary>
        [Output("rotateStrategy")]
        public Output<string?> RotateStrategy { get; private set; } = null!;

        /// <summary>
        /// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
        /// </summary>
        [Output("statements")]
        public Output<ImmutableArray<Outputs.PolicyStatement>> Statements { get; private set; } = null!;

        /// <summary>
        /// The policy type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;

        /// <summary>
        /// The ID of default version policy.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ram/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ram/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// It has been deprecated from provider version 1.114.0 and `policy_document` instead.
        /// </summary>
        [Input("document")]
        public Input<string>? Document { get; set; }

        /// <summary>
        /// This parameter is used for resource destroy. Default value is `false`.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// It has been deprecated from provider version 1.114.0 and `policy_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Document of the RAM policy. It is required when the `statement` is not specified.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        /// <summary>
        /// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
        /// </summary>
        [Input("rotateStrategy")]
        public Input<string>? RotateStrategy { get; set; }

        [Input("statements")]
        private InputList<Inputs.PolicyStatementArgs>? _statements;

        /// <summary>
        /// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
        /// </summary>
        [Obsolete(@"Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace. ")]
        public InputList<Inputs.PolicyStatementArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.PolicyStatementArgs>());
            set => _statements = value;
        }

        /// <summary>
        /// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }

    public sealed class PolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy attachment count.
        /// </summary>
        [Input("attachmentCount")]
        public Input<int>? AttachmentCount { get; set; }

        /// <summary>
        /// The default version of policy.
        /// </summary>
        [Input("defaultVersion")]
        public Input<string>? DefaultVersion { get; set; }

        /// <summary>
        /// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// It has been deprecated from provider version 1.114.0 and `policy_document` instead.
        /// </summary>
        [Input("document")]
        public Input<string>? Document { get; set; }

        /// <summary>
        /// This parameter is used for resource destroy. Default value is `false`.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// It has been deprecated from provider version 1.114.0 and `policy_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Document of the RAM policy. It is required when the `statement` is not specified.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        /// <summary>
        /// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
        /// </summary>
        [Input("rotateStrategy")]
        public Input<string>? RotateStrategy { get; set; }

        [Input("statements")]
        private InputList<Inputs.PolicyStatementGetArgs>? _statements;

        /// <summary>
        /// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
        /// </summary>
        [Obsolete(@"Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace. ")]
        public InputList<Inputs.PolicyStatementGetArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.PolicyStatementGetArgs>());
            set => _statements = value;
        }

        /// <summary>
        /// The policy type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The ID of default version policy.
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public PolicyState()
        {
        }
        public static new PolicyState Empty => new PolicyState();
    }
}