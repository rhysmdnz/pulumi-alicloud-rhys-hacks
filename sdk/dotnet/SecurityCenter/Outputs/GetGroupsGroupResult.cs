// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.SecurityCenter.Outputs
{

    [OutputType]
    public sealed class GetGroupsGroupResult
    {
        /// <summary>
        /// GroupFlag, '0' mean default group(created by system), '1' means customer defined group.
        /// </summary>
        public readonly int GroupFlag;
        /// <summary>
        /// The ID of Group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The name of Group.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// The ID of the Group(same as the group_id).
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGroupsGroupResult(
            int groupFlag,

            string groupId,

            string groupName,

            string id)
        {
            GroupFlag = groupFlag;
            GroupId = groupId;
            GroupName = groupName;
            Id = id;
        }
    }
}