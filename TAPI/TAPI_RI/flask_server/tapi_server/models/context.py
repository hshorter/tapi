# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from tapi_server.models.base_model_ import Model
from tapi_server.models.global_class import GlobalClass  # noqa: F401,E501
from tapi_server.models.name_and_value import NameAndValue  # noqa: F401,E501
from tapi_server.models.service_interface_point import ServiceInterfacePoint  # noqa: F401,E501
from tapi_server import util


class Context(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, uuid: str=None, name: List[NameAndValue]=None, service_interface_point: List[ServiceInterfacePoint]=None):  # noqa: E501
        """Context - a model defined in Swagger

        :param uuid: The uuid of this Context.  # noqa: E501
        :type uuid: str
        :param name: The name of this Context.  # noqa: E501
        :type name: List[NameAndValue]
        :param service_interface_point: The service_interface_point of this Context.  # noqa: E501
        :type service_interface_point: List[ServiceInterfacePoint]
        """
        self.swagger_types = {
            'uuid': str,
            'name': List[NameAndValue],
            'service_interface_point': List[ServiceInterfacePoint]
        }

        self.attribute_map = {
            'uuid': 'uuid',
            'name': 'name',
            'service_interface_point': 'service-interface-point'
        }

        self._uuid = uuid
        self._name = name
        self._service_interface_point = service_interface_point

    @classmethod
    def from_dict(cls, dikt) -> 'Context':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The context of this Context.  # noqa: E501
        :rtype: Context
        """
        return util.deserialize_model(dikt, cls)

    @property
    def uuid(self) -> str:
        """Gets the uuid of this Context.

        UUID: An identifier that is universally unique within an identifier space, where the identifier space is itself globally unique, and immutable. An UUID carries no semantics with respect to the purpose or state of the entity. UUID here uses string representation as defined in RFC 4122.  The canonical representation uses lowercase characters. Pattern: [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-' + '[0-9a-fA-F]{4}-[0-9a-fA-F]{12}  Example of a UUID in string representation: f81d4fae-7dec-11d0-a765-00a0c91e6bf6  # noqa: E501

        :return: The uuid of this Context.
        :rtype: str
        """
        return self._uuid

    @uuid.setter
    def uuid(self, uuid: str):
        """Sets the uuid of this Context.

        UUID: An identifier that is universally unique within an identifier space, where the identifier space is itself globally unique, and immutable. An UUID carries no semantics with respect to the purpose or state of the entity. UUID here uses string representation as defined in RFC 4122.  The canonical representation uses lowercase characters. Pattern: [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-' + '[0-9a-fA-F]{4}-[0-9a-fA-F]{12}  Example of a UUID in string representation: f81d4fae-7dec-11d0-a765-00a0c91e6bf6  # noqa: E501

        :param uuid: The uuid of this Context.
        :type uuid: str
        """

        self._uuid = uuid

    @property
    def name(self) -> List[NameAndValue]:
        """Gets the name of this Context.

        List of names. A property of an entity with a value that is unique in some namespace but may change during the life of the entity. A name carries no semantics with respect to the purpose of the entity.  # noqa: E501

        :return: The name of this Context.
        :rtype: List[NameAndValue]
        """
        return self._name

    @name.setter
    def name(self, name: List[NameAndValue]):
        """Sets the name of this Context.

        List of names. A property of an entity with a value that is unique in some namespace but may change during the life of the entity. A name carries no semantics with respect to the purpose of the entity.  # noqa: E501

        :param name: The name of this Context.
        :type name: List[NameAndValue]
        """

        self._name = name

    @property
    def service_interface_point(self) -> List[ServiceInterfacePoint]:
        """Gets the service_interface_point of this Context.


        :return: The service_interface_point of this Context.
        :rtype: List[ServiceInterfacePoint]
        """
        return self._service_interface_point

    @service_interface_point.setter
    def service_interface_point(self, service_interface_point: List[ServiceInterfacePoint]):
        """Sets the service_interface_point of this Context.


        :param service_interface_point: The service_interface_point of this Context.
        :type service_interface_point: List[ServiceInterfacePoint]
        """

        self._service_interface_point = service_interface_point