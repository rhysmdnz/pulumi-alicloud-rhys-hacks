// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dfs.Outputs
{

    [OutputType]
    public sealed class GetAccessRulesRuleResult
    {
        /// <summary>
        /// The resource ID of the Access Group.
        /// </summary>
        public readonly string AccessGroupId;
        /// <summary>
        /// The ID of the Access Rule.
        /// </summary>
        public readonly string AccessRuleId;
        /// <summary>
        /// The created time of the Access Rule.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the Access Rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The resource ID of Access Rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The NetworkSegment of the Access Rule.
        /// </summary>
        public readonly string NetworkSegment;
        /// <summary>
        /// The priority of the Access Rule.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
        /// </summary>
        public readonly string RwAccessType;

        [OutputConstructor]
        private GetAccessRulesRuleResult(
            string accessGroupId,

            string accessRuleId,

            string createTime,

            string description,

            string id,

            string networkSegment,

            int priority,

            string rwAccessType)
        {
            AccessGroupId = accessGroupId;
            AccessRuleId = accessRuleId;
            CreateTime = createTime;
            Description = description;
            Id = id;
            NetworkSegment = networkSegment;
            Priority = priority;
            RwAccessType = rwAccessType;
        }
    }
}