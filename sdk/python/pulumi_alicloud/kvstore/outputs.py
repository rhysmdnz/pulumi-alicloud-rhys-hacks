# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceParameter',
    'GetAccountsAccountResult',
    'GetConnectionsConnectionResult',
    'GetInstanceClassesClassResult',
    'GetInstanceEnginesInstanceEngineResult',
    'GetInstancesInstanceResult',
    'GetZonesZoneResult',
]

@pulumi.output_type
class InstanceParameter(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class GetAccountsAccountResult(dict):
    def __init__(__self__, *,
                 account_name: str,
                 account_privilege: str,
                 account_type: str,
                 description: str,
                 id: str,
                 instance_id: str,
                 status: str):
        """
        :param str account_name: The name of the account.
        :param str account_privilege: The privilege of account access database.
        :param str account_type: Privilege type of account.
        :param str description: The description of account.
        :param str id: The ID of the Account.
        :param str instance_id: The Id of instance in which account belongs.
        :param str status: The status of account.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_privilege", account_privilege)
        pulumi.set(__self__, "account_type", account_type)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        """
        The name of the account.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountPrivilege")
    def account_privilege(self) -> str:
        """
        The privilege of account access database.
        """
        return pulumi.get(self, "account_privilege")

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> str:
        """
        Privilege type of account.
        """
        return pulumi.get(self, "account_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of account.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Account.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The Id of instance in which account belongs.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of account.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetConnectionsConnectionResult(dict):
    def __init__(__self__, *,
                 connection_string: str,
                 db_instance_net_type: str,
                 expired_time: str,
                 id: str,
                 instance_id: str,
                 ip_address: str,
                 port: str,
                 upgradeable: str,
                 vpc_id: str,
                 vpc_instance_id: str,
                 vswitch_id: str):
        """
        :param str connection_string: The connection string of the instance.
        :param str db_instance_net_type: The network type of the instance.
        :param str expired_time: The expiration time of the classic network address.
        :param str ip_address: The IP address of the instance.
        :param str port: The port number of the instance.
        :param str upgradeable: The remaining validity period of the endpoint of the classic network.
        :param str vpc_id: The ID of the VPC where the instance is deployed.
        :param str vpc_instance_id: The ID of the instance. It is returned only when the value of the DBInstanceNetType parameter is 2 (indicating VPC).
        :param str vswitch_id: The ID of the VSwitch.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "db_instance_net_type", db_instance_net_type)
        pulumi.set(__self__, "expired_time", expired_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "ip_address", ip_address)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "upgradeable", upgradeable)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vpc_instance_id", vpc_instance_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        The connection string of the instance.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="dbInstanceNetType")
    def db_instance_net_type(self) -> str:
        """
        The network type of the instance.
        """
        return pulumi.get(self, "db_instance_net_type")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> str:
        """
        The expiration time of the classic network address.
        """
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The IP address of the instance.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def port(self) -> str:
        """
        The port number of the instance.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def upgradeable(self) -> str:
        """
        The remaining validity period of the endpoint of the classic network.
        """
        return pulumi.get(self, "upgradeable")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the VPC where the instance is deployed.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcInstanceId")
    def vpc_instance_id(self) -> str:
        """
        The ID of the instance. It is returned only when the value of the DBInstanceNetType parameter is 2 (indicating VPC).
        """
        return pulumi.get(self, "vpc_instance_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> str:
        """
        The ID of the VSwitch.
        """
        return pulumi.get(self, "vswitch_id")


@pulumi.output_type
class GetInstanceClassesClassResult(dict):
    def __init__(__self__, *,
                 instance_class: str,
                 price: str):
        """
        :param str instance_class: KVStore available instance class.
        """
        pulumi.set(__self__, "instance_class", instance_class)
        pulumi.set(__self__, "price", price)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> str:
        """
        KVStore available instance class.
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter
    def price(self) -> str:
        return pulumi.get(self, "price")


@pulumi.output_type
class GetInstanceEnginesInstanceEngineResult(dict):
    def __init__(__self__, *,
                 engine: str,
                 engine_version: str,
                 zone_id: str):
        """
        :param str engine: Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
        :param str engine_version: Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
        :param str zone_id: The Zone to launch the KVStore instance.
        """
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        """
        Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The Zone to launch the KVStore instance.
        """
        return pulumi.get(self, "zone_id")


@pulumi.output_type
class GetInstancesInstanceResult(dict):
    def __init__(__self__, *,
                 architecture_type: str,
                 auto_renew: bool,
                 auto_renew_period: int,
                 availability_zone: str,
                 bandwidth: int,
                 capacity: int,
                 charge_type: str,
                 config: Mapping[str, Any],
                 connection_domain: str,
                 connection_mode: str,
                 connections: int,
                 create_time: str,
                 db_instance_id: str,
                 db_instance_name: str,
                 destroy_time: str,
                 end_time: str,
                 engine_version: str,
                 expire_time: str,
                 has_renew_change_order: bool,
                 id: str,
                 instance_class: str,
                 instance_release_protection: bool,
                 instance_type: str,
                 is_rds: bool,
                 maintain_end_time: str,
                 maintain_start_time: str,
                 max_connections: int,
                 name: str,
                 network_type: str,
                 node_type: str,
                 package_type: str,
                 payment_type: str,
                 port: int,
                 private_ip: str,
                 qps: int,
                 region_id: str,
                 replacate_id: str,
                 resource_group_id: str,
                 search_key: str,
                 secondary_zone_id: str,
                 security_group_id: str,
                 security_ip_group_attribute: str,
                 security_ip_group_name: str,
                 security_ips: Sequence[str],
                 ssl_enable: str,
                 status: str,
                 tags: Mapping[str, Any],
                 user_name: str,
                 vpc_auth_mode: str,
                 vpc_cloud_instance_id: str,
                 vpc_id: str,
                 vswitch_id: str,
                 zone_id: str):
        """
        :param str architecture_type: The type of the architecture. Valid values: `cluster`, `standard` and `SplitRW`.
        :param str availability_zone: It has been deprecated from provider version 1.101.0 and `zone_id` instead.
        :param int bandwidth: Instance bandwidth limit. Unit: Mbit/s.
        :param int capacity: Capacity of the applied ApsaraDB for the instance. Unit: MB.
        :param str charge_type: It has been deprecated from provider version 1.101.0 and `payment_type` instead.
        :param Mapping[str, Any] config: The parameter configuration of the instance.
        :param str connection_domain: Instance connection domain (only Intranet access supported).
        :param str connection_mode: The connection mode of the instance.
        :param int connections: IIt has been deprecated from provider version 1.101.0 and `max_connections` instead.
        :param str create_time: Creation time of the instance.
        :param str db_instance_id: The ID of the instance.
        :param str db_instance_name: The name of the instance.
        :param str destroy_time: The time when the instance was destroyed.
        :param str end_time: Expiration time. Pay-As-You-Go instances are never expire.
        :param str engine_version: The engine version. Valid values: `2.8`, `4.0`, `5.0`, `6.0`.
        :param str expire_time: It has been deprecated from provider version 1.101.0 and `end_time` instead.
        :param bool has_renew_change_order: Indicates whether there was an order of renewal with configuration change that had not taken effect.
        :param str id: The ID of the instance.
        :param str instance_class: Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
        :param str instance_type: The engine type of the KVStore DBInstance. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
        :param bool is_rds: Indicates whether the instance is managed by Relational Database Service (RDS).
        :param int max_connections: Instance connection quantity limit. Unit: count.
        :param str name: It has been deprecated from provider version 1.101.0 and `db_instance_name` instead.
        :param str network_type: The type of the network. Valid values: `CLASSIC`, `VPC`.
        :param str node_type: The node type of the instance.
        :param str package_type: The type of the package.
        :param str payment_type: The payment type. Valid values: `PostPaid`, `PrePaid`.
        :param int port: The service port of the instance.
        :param str private_ip: Private IP address of the instance.
        :param int qps: The queries per second (QPS) supported by the instance.
        :param str region_id: Region ID the instance belongs to.
        :param str replacate_id: The logical ID of the replica instance.
        :param str resource_group_id: The ID of the resource group.
        :param str search_key: The name of the instance.
        :param str secondary_zone_id: (Optional, Available in 1.128.0+) The ID of the secondary zone to which you want to migrate the ApsaraDB for Redis instance.
        :param str status: The status of the KVStore DBInstance. Valid values: `Changing`, `CleaningUpExpiredData`, `Creating`, `Flushing`, `HASwitching`, `Inactive`, `MajorVersionUpgrading`, `Migrating`, `NetworkModifying`, `Normal`, `Rebooting`, `SSLModifying`, `Transforming`, `ZoneMigrating`.
        :param Mapping[str, Any] tags: Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
        :param str user_name: The username of the instance.
        :param str vpc_cloud_instance_id: Connection port of the instance.
        :param str vpc_id: Used to retrieve instances belong to specified VPC.
        :param str vswitch_id: Used to retrieve instances belong to specified `vswitch` resources.
        :param str zone_id: The ID of the zone.
        """
        pulumi.set(__self__, "architecture_type", architecture_type)
        pulumi.set(__self__, "auto_renew", auto_renew)
        pulumi.set(__self__, "auto_renew_period", auto_renew_period)
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "connection_domain", connection_domain)
        pulumi.set(__self__, "connection_mode", connection_mode)
        pulumi.set(__self__, "connections", connections)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "db_instance_name", db_instance_name)
        pulumi.set(__self__, "destroy_time", destroy_time)
        pulumi.set(__self__, "end_time", end_time)
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "has_renew_change_order", has_renew_change_order)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_class", instance_class)
        pulumi.set(__self__, "instance_release_protection", instance_release_protection)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "is_rds", is_rds)
        pulumi.set(__self__, "maintain_end_time", maintain_end_time)
        pulumi.set(__self__, "maintain_start_time", maintain_start_time)
        pulumi.set(__self__, "max_connections", max_connections)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "network_type", network_type)
        pulumi.set(__self__, "node_type", node_type)
        pulumi.set(__self__, "package_type", package_type)
        pulumi.set(__self__, "payment_type", payment_type)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "private_ip", private_ip)
        pulumi.set(__self__, "qps", qps)
        pulumi.set(__self__, "region_id", region_id)
        pulumi.set(__self__, "replacate_id", replacate_id)
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        pulumi.set(__self__, "search_key", search_key)
        pulumi.set(__self__, "secondary_zone_id", secondary_zone_id)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "security_ip_group_attribute", security_ip_group_attribute)
        pulumi.set(__self__, "security_ip_group_name", security_ip_group_name)
        pulumi.set(__self__, "security_ips", security_ips)
        pulumi.set(__self__, "ssl_enable", ssl_enable)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "user_name", user_name)
        pulumi.set(__self__, "vpc_auth_mode", vpc_auth_mode)
        pulumi.set(__self__, "vpc_cloud_instance_id", vpc_cloud_instance_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="architectureType")
    def architecture_type(self) -> str:
        """
        The type of the architecture. Valid values: `cluster`, `standard` and `SplitRW`.
        """
        return pulumi.get(self, "architecture_type")

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> bool:
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="autoRenewPeriod")
    def auto_renew_period(self) -> int:
        return pulumi.get(self, "auto_renew_period")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        It has been deprecated from provider version 1.101.0 and `zone_id` instead.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        """
        Instance bandwidth limit. Unit: Mbit/s.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def capacity(self) -> int:
        """
        Capacity of the applied ApsaraDB for the instance. Unit: MB.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        """
        It has been deprecated from provider version 1.101.0 and `payment_type` instead.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter
    def config(self) -> Mapping[str, Any]:
        """
        The parameter configuration of the instance.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="connectionDomain")
    def connection_domain(self) -> str:
        """
        Instance connection domain (only Intranet access supported).
        """
        return pulumi.get(self, "connection_domain")

    @property
    @pulumi.getter(name="connectionMode")
    def connection_mode(self) -> str:
        """
        The connection mode of the instance.
        """
        return pulumi.get(self, "connection_mode")

    @property
    @pulumi.getter
    def connections(self) -> int:
        """
        IIt has been deprecated from provider version 1.101.0 and `max_connections` instead.
        """
        return pulumi.get(self, "connections")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of the instance.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> str:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="dbInstanceName")
    def db_instance_name(self) -> str:
        """
        The name of the instance.
        """
        return pulumi.get(self, "db_instance_name")

    @property
    @pulumi.getter(name="destroyTime")
    def destroy_time(self) -> str:
        """
        The time when the instance was destroyed.
        """
        return pulumi.get(self, "destroy_time")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        Expiration time. Pay-As-You-Go instances are never expire.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        """
        The engine version. Valid values: `2.8`, `4.0`, `5.0`, `6.0`.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        It has been deprecated from provider version 1.101.0 and `end_time` instead.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="hasRenewChangeOrder")
    def has_renew_change_order(self) -> bool:
        """
        Indicates whether there was an order of renewal with configuration change that had not taken effect.
        """
        return pulumi.get(self, "has_renew_change_order")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> str:
        """
        Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="instanceReleaseProtection")
    def instance_release_protection(self) -> bool:
        return pulumi.get(self, "instance_release_protection")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The engine type of the KVStore DBInstance. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="isRds")
    def is_rds(self) -> bool:
        """
        Indicates whether the instance is managed by Relational Database Service (RDS).
        """
        return pulumi.get(self, "is_rds")

    @property
    @pulumi.getter(name="maintainEndTime")
    def maintain_end_time(self) -> str:
        return pulumi.get(self, "maintain_end_time")

    @property
    @pulumi.getter(name="maintainStartTime")
    def maintain_start_time(self) -> str:
        return pulumi.get(self, "maintain_start_time")

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> int:
        """
        Instance connection quantity limit. Unit: count.
        """
        return pulumi.get(self, "max_connections")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        It has been deprecated from provider version 1.101.0 and `db_instance_name` instead.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        """
        The type of the network. Valid values: `CLASSIC`, `VPC`.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> str:
        """
        The node type of the instance.
        """
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> str:
        """
        The type of the package.
        """
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> str:
        """
        The payment type. Valid values: `PostPaid`, `PrePaid`.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The service port of the instance.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        Private IP address of the instance.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter
    def qps(self) -> int:
        """
        The queries per second (QPS) supported by the instance.
        """
        return pulumi.get(self, "qps")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        """
        Region ID the instance belongs to.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="replacateId")
    def replacate_id(self) -> str:
        """
        The logical ID of the replica instance.
        """
        return pulumi.get(self, "replacate_id")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> str:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="searchKey")
    def search_key(self) -> str:
        """
        The name of the instance.
        """
        return pulumi.get(self, "search_key")

    @property
    @pulumi.getter(name="secondaryZoneId")
    def secondary_zone_id(self) -> str:
        """
        (Optional, Available in 1.128.0+) The ID of the secondary zone to which you want to migrate the ApsaraDB for Redis instance.
        """
        return pulumi.get(self, "secondary_zone_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="securityIpGroupAttribute")
    def security_ip_group_attribute(self) -> str:
        return pulumi.get(self, "security_ip_group_attribute")

    @property
    @pulumi.getter(name="securityIpGroupName")
    def security_ip_group_name(self) -> str:
        return pulumi.get(self, "security_ip_group_name")

    @property
    @pulumi.getter(name="securityIps")
    def security_ips(self) -> Sequence[str]:
        return pulumi.get(self, "security_ips")

    @property
    @pulumi.getter(name="sslEnable")
    def ssl_enable(self) -> str:
        return pulumi.get(self, "ssl_enable")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the KVStore DBInstance. Valid values: `Changing`, `CleaningUpExpiredData`, `Creating`, `Flushing`, `HASwitching`, `Inactive`, `MajorVersionUpgrading`, `Migrating`, `NetworkModifying`, `Normal`, `Rebooting`, `SSLModifying`, `Transforming`, `ZoneMigrating`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        """
        Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The username of the instance.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="vpcAuthMode")
    def vpc_auth_mode(self) -> str:
        return pulumi.get(self, "vpc_auth_mode")

    @property
    @pulumi.getter(name="vpcCloudInstanceId")
    def vpc_cloud_instance_id(self) -> str:
        """
        Connection port of the instance.
        """
        return pulumi.get(self, "vpc_cloud_instance_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        Used to retrieve instances belong to specified VPC.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> str:
        """
        Used to retrieve instances belong to specified `vswitch` resources.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The ID of the zone.
        """
        return pulumi.get(self, "zone_id")


@pulumi.output_type
class GetZonesZoneResult(dict):
    def __init__(__self__, *,
                 id: str,
                 multi_zone_ids: Sequence[str]):
        """
        :param str id: ID of the zone.
        :param Sequence[str] multi_zone_ids: A list of zone ids in which the multi zone.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "multi_zone_ids", multi_zone_ids)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the zone.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="multiZoneIds")
    def multi_zone_ids(self) -> Sequence[str]:
        """
        A list of zone ids in which the multi zone.
        """
        return pulumi.get(self, "multi_zone_ids")

