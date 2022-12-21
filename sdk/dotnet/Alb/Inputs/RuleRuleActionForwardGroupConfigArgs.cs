// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class RuleRuleActionForwardGroupConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration of session persistence for server groups.
        /// </summary>
        [Input("serverGroupStickySession")]
        public Input<Inputs.RuleRuleActionForwardGroupConfigServerGroupStickySessionArgs>? ServerGroupStickySession { get; set; }

        [Input("serverGroupTuples")]
        private InputList<Inputs.RuleRuleActionForwardGroupConfigServerGroupTupleArgs>? _serverGroupTuples;

        /// <summary>
        /// The destination server group to which requests are forwarded.
        /// </summary>
        public InputList<Inputs.RuleRuleActionForwardGroupConfigServerGroupTupleArgs> ServerGroupTuples
        {
            get => _serverGroupTuples ?? (_serverGroupTuples = new InputList<Inputs.RuleRuleActionForwardGroupConfigServerGroupTupleArgs>());
            set => _serverGroupTuples = value;
        }

        public RuleRuleActionForwardGroupConfigArgs()
        {
        }
        public static new RuleRuleActionForwardGroupConfigArgs Empty => new RuleRuleActionForwardGroupConfigArgs();
    }
}