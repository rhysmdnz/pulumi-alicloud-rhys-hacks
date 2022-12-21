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
    public sealed class GetFileSystemsSystemResult
    {
        /// <summary>
        /// (Optional, Available in v1.140.0+) The capacity of the file system.
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// Time of creation.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Description of the FileSystem.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Optional, Available in v1.121.2+) Whether the file system is encrypted. 
        /// * Valid values:
        /// * `0`: The file system is not encrypted.
        /// * `1`: The file system is encrypted with a managed secret key.
        /// * `2`: User management key.
        /// </summary>
        public readonly int EncryptType;
        /// <summary>
        /// The type of the file system.
        /// Valid values:
        /// `standard` (Default),
        /// `extreme`.
        /// </summary>
        public readonly string FileSystemType;
        /// <summary>
        /// ID of the FileSystem.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Optional, Available in v1.140.0+) The id of the KMS key.
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// MeteredSize of the FileSystem.
        /// </summary>
        public readonly int MeteredSize;
        /// <summary>
        /// The protocol type of the file system.
        /// Valid values:
        /// `NFS`,
        /// `SMB` (Available when the `file_system_type` is `standard`).
        /// </summary>
        public readonly string ProtocolType;
        /// <summary>
        /// ID of the region where the FileSystem is located.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The storage type of the file system.
        /// * Valid values:
        /// </summary>
        public readonly string StorageType;
        /// <summary>
        /// (Optional, Available in v1.140.0+) The id of the zone. Each region consists of multiple isolated locations known as zones. Each zone has an independent power supply and network.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetFileSystemsSystemResult(
            int capacity,

            string createTime,

            string description,

            int encryptType,

            string fileSystemType,

            string id,

            string kmsKeyId,

            int meteredSize,

            string protocolType,

            string regionId,

            string storageType,

            string zoneId)
        {
            Capacity = capacity;
            CreateTime = createTime;
            Description = description;
            EncryptType = encryptType;
            FileSystemType = fileSystemType;
            Id = id;
            KmsKeyId = kmsKeyId;
            MeteredSize = meteredSize;
            ProtocolType = protocolType;
            RegionId = regionId;
            StorageType = storageType;
            ZoneId = zoneId;
        }
    }
}