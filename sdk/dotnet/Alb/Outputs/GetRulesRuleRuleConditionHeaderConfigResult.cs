// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class GetRulesRuleRuleConditionHeaderConfigResult
    {
        /// <summary>
        /// The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Add one or more IP addresses or IP address segments.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetRulesRuleRuleConditionHeaderConfigResult(
            string key,

            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }
}