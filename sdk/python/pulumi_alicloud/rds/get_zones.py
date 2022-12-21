# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetZonesResult',
    'AwaitableGetZonesResult',
    'get_zones',
    'get_zones_output',
]

@pulumi.output_type
class GetZonesResult:
    """
    A collection of values returned by getZones.
    """
    def __init__(__self__, category=None, db_instance_class=None, db_instance_storage_type=None, engine=None, engine_version=None, id=None, ids=None, instance_charge_type=None, multi=None, multi_zone=None, output_file=None, zones=None):
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        pulumi.set(__self__, "category", category)
        if db_instance_class and not isinstance(db_instance_class, str):
            raise TypeError("Expected argument 'db_instance_class' to be a str")
        pulumi.set(__self__, "db_instance_class", db_instance_class)
        if db_instance_storage_type and not isinstance(db_instance_storage_type, str):
            raise TypeError("Expected argument 'db_instance_storage_type' to be a str")
        pulumi.set(__self__, "db_instance_storage_type", db_instance_storage_type)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_charge_type and not isinstance(instance_charge_type, str):
            raise TypeError("Expected argument 'instance_charge_type' to be a str")
        pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if multi and not isinstance(multi, bool):
            raise TypeError("Expected argument 'multi' to be a bool")
        pulumi.set(__self__, "multi", multi)
        if multi_zone and not isinstance(multi_zone, bool):
            raise TypeError("Expected argument 'multi_zone' to be a bool")
        pulumi.set(__self__, "multi_zone", multi_zone)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="dbInstanceClass")
    def db_instance_class(self) -> Optional[str]:
        return pulumi.get(self, "db_instance_class")

    @property
    @pulumi.getter(name="dbInstanceStorageType")
    def db_instance_storage_type(self) -> Optional[str]:
        return pulumi.get(self, "db_instance_storage_type")

    @property
    @pulumi.getter
    def engine(self) -> Optional[str]:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[str]:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        A list of zone IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter
    def multi(self) -> Optional[bool]:
        return pulumi.get(self, "multi")

    @property
    @pulumi.getter(name="multiZone")
    def multi_zone(self) -> Optional[bool]:
        return pulumi.get(self, "multi_zone")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def zones(self) -> Sequence['outputs.GetZonesZoneResult']:
        """
        A list of availability zones. Each element contains the following attributes:
        """
        return pulumi.get(self, "zones")


class AwaitableGetZonesResult(GetZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZonesResult(
            category=self.category,
            db_instance_class=self.db_instance_class,
            db_instance_storage_type=self.db_instance_storage_type,
            engine=self.engine,
            engine_version=self.engine_version,
            id=self.id,
            ids=self.ids,
            instance_charge_type=self.instance_charge_type,
            multi=self.multi,
            multi_zone=self.multi_zone,
            output_file=self.output_file,
            zones=self.zones)


def get_zones(category: Optional[str] = None,
              db_instance_class: Optional[str] = None,
              db_instance_storage_type: Optional[str] = None,
              engine: Optional[str] = None,
              engine_version: Optional[str] = None,
              instance_charge_type: Optional[str] = None,
              multi: Optional[bool] = None,
              multi_zone: Optional[bool] = None,
              output_file: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZonesResult:
    """
    This data source provides availability zones for RDS that can be accessed by an Alibaba Cloud account within the region configured in the provider.

    > **NOTE:** Available in v1.73.0+.


    :param str category: DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
    :param str db_instance_storage_type: The DB instance storage space required by the user. Valid values: "cloud_ssd", "local_ssd", "cloud_essd", "cloud_essd2", "cloud_essd3".
    :param str engine: Database type. Valid values: "MySQL", "SQLServer", "PostgreSQL", "PPAS", "MariaDB". If not set, it will match all of engines.
    :param str engine_version: Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
    :param str instance_charge_type: Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param bool multi: It has been deprecated from version 1.137.0 and using `multi_zone` instead.
    :param bool multi_zone: Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
    """
    __args__ = dict()
    __args__['category'] = category
    __args__['dbInstanceClass'] = db_instance_class
    __args__['dbInstanceStorageType'] = db_instance_storage_type
    __args__['engine'] = engine
    __args__['engineVersion'] = engine_version
    __args__['instanceChargeType'] = instance_charge_type
    __args__['multi'] = multi
    __args__['multiZone'] = multi_zone
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:rds/getZones:getZones', __args__, opts=opts, typ=GetZonesResult).value

    return AwaitableGetZonesResult(
        category=__ret__.category,
        db_instance_class=__ret__.db_instance_class,
        db_instance_storage_type=__ret__.db_instance_storage_type,
        engine=__ret__.engine,
        engine_version=__ret__.engine_version,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_charge_type=__ret__.instance_charge_type,
        multi=__ret__.multi,
        multi_zone=__ret__.multi_zone,
        output_file=__ret__.output_file,
        zones=__ret__.zones)


@_utilities.lift_output_func(get_zones)
def get_zones_output(category: Optional[pulumi.Input[Optional[str]]] = None,
                     db_instance_class: Optional[pulumi.Input[Optional[str]]] = None,
                     db_instance_storage_type: Optional[pulumi.Input[Optional[str]]] = None,
                     engine: Optional[pulumi.Input[Optional[str]]] = None,
                     engine_version: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_charge_type: Optional[pulumi.Input[Optional[str]]] = None,
                     multi: Optional[pulumi.Input[Optional[bool]]] = None,
                     multi_zone: Optional[pulumi.Input[Optional[bool]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZonesResult]:
    """
    This data source provides availability zones for RDS that can be accessed by an Alibaba Cloud account within the region configured in the provider.

    > **NOTE:** Available in v1.73.0+.


    :param str category: DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
    :param str db_instance_storage_type: The DB instance storage space required by the user. Valid values: "cloud_ssd", "local_ssd", "cloud_essd", "cloud_essd2", "cloud_essd3".
    :param str engine: Database type. Valid values: "MySQL", "SQLServer", "PostgreSQL", "PPAS", "MariaDB". If not set, it will match all of engines.
    :param str engine_version: Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
    :param str instance_charge_type: Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param bool multi: It has been deprecated from version 1.137.0 and using `multi_zone` instead.
    :param bool multi_zone: Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
    """
    ...