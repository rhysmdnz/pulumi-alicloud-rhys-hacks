// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager.Outputs
{

    [OutputType]
    public sealed class GetRolesRoleResult
    {
        public readonly string Arn;
        public readonly string AssumeRolePolicyDocument;
        public readonly string Description;
        /// <summary>
        /// The ID of the role.
        /// * `role_id`- The ID of the role.
        /// * `role_name`- The name of the role.
        /// * `arn`- The Alibaba Cloud Resource Name (ARN) of the RAM role.
        /// * `create_date`- (Removed form v1.114.0) The time when the RAM role was created.
        /// * `update_date`- The time when the RAM role was updated.
        /// * `description`- The description of the RAM role.
        /// * `max_session_duration`- The maximum session duration of the RAM role.
        /// * `assume_role_policy_document`- (Available in v1.114.0+) The assume role policy document.
        /// </summary>
        public readonly string Id;
        public readonly int MaxSessionDuration;
        public readonly string RoleId;
        public readonly string RoleName;
        public readonly string UpdateDate;

        [OutputConstructor]
        private GetRolesRoleResult(
            string arn,

            string assumeRolePolicyDocument,

            string description,

            string id,

            int maxSessionDuration,

            string roleId,

            string roleName,

            string updateDate)
        {
            Arn = arn;
            AssumeRolePolicyDocument = assumeRolePolicyDocument;
            Description = description;
            Id = id;
            MaxSessionDuration = maxSessionDuration;
            RoleId = roleId;
            RoleName = roleName;
            UpdateDate = updateDate;
        }
    }
}