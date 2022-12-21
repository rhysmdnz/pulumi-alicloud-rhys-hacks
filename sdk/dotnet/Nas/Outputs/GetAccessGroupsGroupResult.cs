// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas.Outputs
{

    [OutputType]
    public sealed class GetAccessGroupsGroupResult
    {
        /// <summary>
        /// The name of access group.
        /// </summary>
        public readonly string AccessGroupName;
        /// <summary>
        /// Filter results by a specific AccessGroupType.
        /// </summary>
        public readonly string AccessGroupType;
        /// <summary>
        /// Filter results by a specific Description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// This ID of this AccessGroup. It is formatted to ``&lt;access_group_id&gt;:&lt;file_system_type&gt;``. Before version 1.95.0, the value is `access_group_name`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// MountTargetCount block of the AccessGroup
        /// </summary>
        public readonly int MountTargetCount;
        /// <summary>
        /// RuleCount of the AccessGroup.
        /// </summary>
        public readonly int RuleCount;
        /// <summary>
        /// Field `type` has been deprecated from version 1.95.0. Use `access_group_type` instead.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAccessGroupsGroupResult(
            string accessGroupName,

            string accessGroupType,

            string description,

            string id,

            int mountTargetCount,

            int ruleCount,

            string type)
        {
            AccessGroupName = accessGroupName;
            AccessGroupType = accessGroupType;
            Description = description;
            Id = id;
            MountTargetCount = mountTargetCount;
            RuleCount = ruleCount;
            Type = type;
        }
    }
}