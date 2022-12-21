// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Pvtz.Outputs
{

    [OutputType]
    public sealed class GetRulesRuleResult
    {
        /// <summary>
        /// The List of the VPC. See the following `Block bind_vpcs`. **NOTE:** Available in v1.158.0+.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRulesRuleBindVpcResult> BindVpcs;
        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The ID of the Endpoint.
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// The Name of the Endpoint.
        /// </summary>
        public readonly string EndpointName;
        public readonly ImmutableArray<Outputs.GetRulesRuleForwardIpResult> ForwardIps;
        /// <summary>
        /// The ID of the Rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        public readonly string RuleId;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// The type of the rule.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The name of the forwarding zone.
        /// </summary>
        public readonly string ZoneName;

        [OutputConstructor]
        private GetRulesRuleResult(
            ImmutableArray<Outputs.GetRulesRuleBindVpcResult> bindVpcs,

            string createTime,

            string endpointId,

            string endpointName,

            ImmutableArray<Outputs.GetRulesRuleForwardIpResult> forwardIps,

            string id,

            string ruleId,

            string ruleName,

            string type,

            string zoneName)
        {
            BindVpcs = bindVpcs;
            CreateTime = createTime;
            EndpointId = endpointId;
            EndpointName = endpointName;
            ForwardIps = forwardIps;
            Id = id;
            RuleId = ruleId;
            RuleName = ruleName;
            Type = type;
            ZoneName = zoneName;
        }
    }
}