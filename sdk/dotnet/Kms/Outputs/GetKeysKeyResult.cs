// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Kms.Outputs
{

    [OutputType]
    public sealed class GetKeysKeyResult
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the key.
        /// </summary>
        public readonly string Arn;
        public readonly string AutomaticRotation;
        /// <summary>
        /// Creation date of key.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The owner of the key.
        /// * `automatic_rotation` -(Available in 1.123.1+) Specifies whether to enable automatic key rotation.
        /// * `key_id` -(Available in 1.123.1+)  ID of the key.
        /// * `key_spec` -(Available in 1.123.1+)  The type of the CMK.
        /// * `key_usage` -(Available in 1.123.1+)  The usage of CMK.
        /// * `last_rotation_date` -(Available in 1.123.1+)  The date and time the last rotation was performed.
        /// * `material_expire_time` -(Available in 1.123.1+)  The time and date the key material for the CMK expires.
        /// * `next_rotation_date` -(Available in 1.123.1+)  The time the next rotation is scheduled for execution.
        /// * `origin` -(Available in 1.123.1+)  The source of the key material for the CMK.
        /// * `protection_level` -(Available in 1.123.1+)  The protection level of the CMK.
        /// * `rotation_interval` -(Available in 1.123.1+)  The period of automatic key rotation.
        /// * `primary_key_version` -(Available in 1.123.1+)  The ID of the current primary key version of the symmetric CMK.
        /// </summary>
        public readonly string Creator;
        /// <summary>
        /// Deletion date of key.
        /// </summary>
        public readonly string DeleteDate;
        /// <summary>
        /// Description of the key.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the key.
        /// </summary>
        public readonly string Id;
        public readonly string KeyId;
        public readonly string KeySpec;
        public readonly string KeyUsage;
        public readonly string LastRotationDate;
        public readonly string MaterialExpireTime;
        public readonly string NextRotationDate;
        public readonly string Origin;
        public readonly string PrimaryKeyVersion;
        public readonly string ProtectionLevel;
        public readonly string RotationInterval;
        /// <summary>
        /// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetKeysKeyResult(
            string arn,

            string automaticRotation,

            string creationDate,

            string creator,

            string deleteDate,

            string description,

            string id,

            string keyId,

            string keySpec,

            string keyUsage,

            string lastRotationDate,

            string materialExpireTime,

            string nextRotationDate,

            string origin,

            string primaryKeyVersion,

            string protectionLevel,

            string rotationInterval,

            string status)
        {
            Arn = arn;
            AutomaticRotation = automaticRotation;
            CreationDate = creationDate;
            Creator = creator;
            DeleteDate = deleteDate;
            Description = description;
            Id = id;
            KeyId = keyId;
            KeySpec = keySpec;
            KeyUsage = keyUsage;
            LastRotationDate = lastRotationDate;
            MaterialExpireTime = materialExpireTime;
            NextRotationDate = nextRotationDate;
            Origin = origin;
            PrimaryKeyVersion = primaryKeyVersion;
            ProtectionLevel = protectionLevel;
            RotationInterval = rotationInterval;
            Status = status;
        }
    }
}