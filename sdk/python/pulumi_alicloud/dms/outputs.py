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
    'GetEnterpriseInstancesInstanceResult',
    'GetEnterpriseUsersUserResult',
    'GetUserTenantsTenantResult',
]

@pulumi.output_type
class GetEnterpriseInstancesInstanceResult(dict):
    def __init__(__self__, *,
                 data_link_name: str,
                 database_password: str,
                 database_user: str,
                 dba_id: str,
                 dba_nick_name: str,
                 ddl_online: int,
                 ecs_instance_id: str,
                 ecs_region: str,
                 env_type: str,
                 export_timeout: int,
                 host: str,
                 id: str,
                 instance_alias: str,
                 instance_id: str,
                 instance_name: str,
                 instance_source: str,
                 instance_type: str,
                 port: int,
                 query_timeout: int,
                 safe_rule_id: str,
                 sid: str,
                 status: str,
                 use_dsql: int,
                 vpc_id: str):
        """
        :param str data_link_name: The name of the data link for the database instance.
        :param str database_password: The logon password of the database instance.
        :param str database_user: The logon username of the database instance.
        :param str dba_id: The ID of the database administrator (DBA) of the database instance.
        :param str dba_nick_name: The nickname of the DBA.
        :param int ddl_online: Indicates whether the online data description language (DDL) service was enabled for the database instance.
        :param str ecs_instance_id: The ID of the Elastic Compute Service (ECS) instance to which the database instance belongs.
        :param str ecs_region: The region where the database instance resides.
        :param str env_type: The type of the environment to which the database instance belongs.
        :param int export_timeout: The timeout period for exporting the database instance.
        :param str host: The endpoint of the database instance.
        :param str instance_alias: The alias of the database instance.
        :param str instance_id: The ID of the database instance.
        :param str instance_source: The source of the database instance.
        :param str instance_type: The ID of the database instance.
        :param int port: The connection port of the database instance.
        :param int query_timeout: The timeout period for querying the database instance.
        :param str safe_rule_id: The ID of the security rule for the database instance.
        :param str sid: The system ID (SID) of the database instance.
        :param str status: Filter the results by status of the DMS Enterprise Instances. Valid values: `NORMAL`, `UNAVAILABLE`, `UNKNOWN`, `DELETED`, `DISABLE`.
        :param int use_dsql: Indicates whether cross-database query was enabled for the database instance.
        :param str vpc_id: The ID of the Virtual Private Cloud (VPC) to which the database instance belongs.
        """
        pulumi.set(__self__, "data_link_name", data_link_name)
        pulumi.set(__self__, "database_password", database_password)
        pulumi.set(__self__, "database_user", database_user)
        pulumi.set(__self__, "dba_id", dba_id)
        pulumi.set(__self__, "dba_nick_name", dba_nick_name)
        pulumi.set(__self__, "ddl_online", ddl_online)
        pulumi.set(__self__, "ecs_instance_id", ecs_instance_id)
        pulumi.set(__self__, "ecs_region", ecs_region)
        pulumi.set(__self__, "env_type", env_type)
        pulumi.set(__self__, "export_timeout", export_timeout)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_alias", instance_alias)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "instance_source", instance_source)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "query_timeout", query_timeout)
        pulumi.set(__self__, "safe_rule_id", safe_rule_id)
        pulumi.set(__self__, "sid", sid)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "use_dsql", use_dsql)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="dataLinkName")
    def data_link_name(self) -> str:
        """
        The name of the data link for the database instance.
        """
        return pulumi.get(self, "data_link_name")

    @property
    @pulumi.getter(name="databasePassword")
    def database_password(self) -> str:
        """
        The logon password of the database instance.
        """
        return pulumi.get(self, "database_password")

    @property
    @pulumi.getter(name="databaseUser")
    def database_user(self) -> str:
        """
        The logon username of the database instance.
        """
        return pulumi.get(self, "database_user")

    @property
    @pulumi.getter(name="dbaId")
    def dba_id(self) -> str:
        """
        The ID of the database administrator (DBA) of the database instance.
        """
        return pulumi.get(self, "dba_id")

    @property
    @pulumi.getter(name="dbaNickName")
    def dba_nick_name(self) -> str:
        """
        The nickname of the DBA.
        """
        return pulumi.get(self, "dba_nick_name")

    @property
    @pulumi.getter(name="ddlOnline")
    def ddl_online(self) -> int:
        """
        Indicates whether the online data description language (DDL) service was enabled for the database instance.
        """
        return pulumi.get(self, "ddl_online")

    @property
    @pulumi.getter(name="ecsInstanceId")
    def ecs_instance_id(self) -> str:
        """
        The ID of the Elastic Compute Service (ECS) instance to which the database instance belongs.
        """
        return pulumi.get(self, "ecs_instance_id")

    @property
    @pulumi.getter(name="ecsRegion")
    def ecs_region(self) -> str:
        """
        The region where the database instance resides.
        """
        return pulumi.get(self, "ecs_region")

    @property
    @pulumi.getter(name="envType")
    def env_type(self) -> str:
        """
        The type of the environment to which the database instance belongs.
        """
        return pulumi.get(self, "env_type")

    @property
    @pulumi.getter(name="exportTimeout")
    def export_timeout(self) -> int:
        """
        The timeout period for exporting the database instance.
        """
        return pulumi.get(self, "export_timeout")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        The endpoint of the database instance.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceAlias")
    def instance_alias(self) -> str:
        """
        The alias of the database instance.
        """
        return pulumi.get(self, "instance_alias")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The ID of the database instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceSource")
    def instance_source(self) -> str:
        """
        The source of the database instance.
        """
        return pulumi.get(self, "instance_source")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The ID of the database instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The connection port of the database instance.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="queryTimeout")
    def query_timeout(self) -> int:
        """
        The timeout period for querying the database instance.
        """
        return pulumi.get(self, "query_timeout")

    @property
    @pulumi.getter(name="safeRuleId")
    def safe_rule_id(self) -> str:
        """
        The ID of the security rule for the database instance.
        """
        return pulumi.get(self, "safe_rule_id")

    @property
    @pulumi.getter
    def sid(self) -> str:
        """
        The system ID (SID) of the database instance.
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Filter the results by status of the DMS Enterprise Instances. Valid values: `NORMAL`, `UNAVAILABLE`, `UNKNOWN`, `DELETED`, `DISABLE`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="useDsql")
    def use_dsql(self) -> int:
        """
        Indicates whether cross-database query was enabled for the database instance.
        """
        return pulumi.get(self, "use_dsql")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the Virtual Private Cloud (VPC) to which the database instance belongs.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GetEnterpriseUsersUserResult(dict):
    def __init__(__self__, *,
                 id: str,
                 mobile: str,
                 nick_name: str,
                 parent_uid: int,
                 role_ids: Sequence[int],
                 role_names: Sequence[str],
                 status: str,
                 uid: str,
                 user_id: str,
                 user_name: str):
        """
        :param str id: The Alibaba Cloud unique ID (UID) of the user.
        :param str mobile: The DingTalk number or mobile number of the user.
        :param str nick_name: The nickname of the user.
        :param int parent_uid: The Alibaba Cloud unique ID (UID) of the parent account if the user corresponds to a Resource Access Management (RAM) user.
        :param Sequence[int] role_ids: The list ids of the role that the user plays.
        :param Sequence[str] role_names: The list names of the role that he user plays.
        :param str status: The status of the user.
        :param str user_id: The ID of the user.
        :param str user_name: The nickname of the user.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "mobile", mobile)
        pulumi.set(__self__, "nick_name", nick_name)
        pulumi.set(__self__, "parent_uid", parent_uid)
        pulumi.set(__self__, "role_ids", role_ids)
        pulumi.set(__self__, "role_names", role_names)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "uid", uid)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The Alibaba Cloud unique ID (UID) of the user.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mobile(self) -> str:
        """
        The DingTalk number or mobile number of the user.
        """
        return pulumi.get(self, "mobile")

    @property
    @pulumi.getter(name="nickName")
    def nick_name(self) -> str:
        """
        The nickname of the user.
        """
        return pulumi.get(self, "nick_name")

    @property
    @pulumi.getter(name="parentUid")
    def parent_uid(self) -> int:
        """
        The Alibaba Cloud unique ID (UID) of the parent account if the user corresponds to a Resource Access Management (RAM) user.
        """
        return pulumi.get(self, "parent_uid")

    @property
    @pulumi.getter(name="roleIds")
    def role_ids(self) -> Sequence[int]:
        """
        The list ids of the role that the user plays.
        """
        return pulumi.get(self, "role_ids")

    @property
    @pulumi.getter(name="roleNames")
    def role_names(self) -> Sequence[str]:
        """
        The list names of the role that he user plays.
        """
        return pulumi.get(self, "role_names")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the user.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def uid(self) -> str:
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The ID of the user.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The nickname of the user.
        """
        return pulumi.get(self, "user_name")


@pulumi.output_type
class GetUserTenantsTenantResult(dict):
    def __init__(__self__, *,
                 id: str,
                 status: str,
                 tenant_name: str,
                 tid: str):
        """
        :param str id: The user tenant id.
        :param str status: The status of the user tenant.
        :param str tenant_name: The name of the user tenant.
        :param str tid: The user tenant id. Same as id.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "tenant_name", tenant_name)
        pulumi.set(__self__, "tid", tid)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The user tenant id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the user tenant.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> str:
        """
        The name of the user tenant.
        """
        return pulumi.get(self, "tenant_name")

    @property
    @pulumi.getter
    def tid(self) -> str:
        """
        The user tenant id. Same as id.
        """
        return pulumi.get(self, "tid")

