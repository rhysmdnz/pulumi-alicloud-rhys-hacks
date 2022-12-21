// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dms.Outputs
{

    [OutputType]
    public sealed class GetEnterpriseUsersUserResult
    {
        /// <summary>
        /// The Alibaba Cloud unique ID (UID) of the user.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The DingTalk number or mobile number of the user.
        /// </summary>
        public readonly string Mobile;
        /// <summary>
        /// The nickname of the user.
        /// </summary>
        public readonly string NickName;
        /// <summary>
        /// The Alibaba Cloud unique ID (UID) of the parent account if the user corresponds to a Resource Access Management (RAM) user.
        /// </summary>
        public readonly int ParentUid;
        /// <summary>
        /// The list ids of the role that the user plays.
        /// </summary>
        public readonly ImmutableArray<int> RoleIds;
        /// <summary>
        /// The list names of the role that he user plays.
        /// </summary>
        public readonly ImmutableArray<string> RoleNames;
        /// <summary>
        /// The status of the user.
        /// </summary>
        public readonly string Status;
        public readonly string Uid;
        /// <summary>
        /// The ID of the user.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// The nickname of the user.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetEnterpriseUsersUserResult(
            string id,

            string mobile,

            string nickName,

            int parentUid,

            ImmutableArray<int> roleIds,

            ImmutableArray<string> roleNames,

            string status,

            string uid,

            string userId,

            string userName)
        {
            Id = id;
            Mobile = mobile;
            NickName = nickName;
            ParentUid = parentUid;
            RoleIds = roleIds;
            RoleNames = roleNames;
            Status = status;
            Uid = uid;
            UserId = userId;
            UserName = userName;
        }
    }
}