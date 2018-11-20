# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from tapi_server.models.base_model_ import Model
from tapi_server.models.cost_characteristic import CostCharacteristic  # noqa: F401,E501
from tapi_server import util


class TransferCostPac(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, cost_characteristic: List[CostCharacteristic]=None):  # noqa: E501
        """TransferCostPac - a model defined in Swagger

        :param cost_characteristic: The cost_characteristic of this TransferCostPac.  # noqa: E501
        :type cost_characteristic: List[CostCharacteristic]
        """
        self.swagger_types = {
            'cost_characteristic': List[CostCharacteristic]
        }

        self.attribute_map = {
            'cost_characteristic': 'cost-characteristic'
        }

        self._cost_characteristic = cost_characteristic

    @classmethod
    def from_dict(cls, dikt) -> 'TransferCostPac':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The transfer-cost-pac of this TransferCostPac.  # noqa: E501
        :rtype: TransferCostPac
        """
        return util.deserialize_model(dikt, cls)

    @property
    def cost_characteristic(self) -> List[CostCharacteristic]:
        """Gets the cost_characteristic of this TransferCostPac.

        The list of costs where each cost relates to some aspect of the TopologicalEntity.  # noqa: E501

        :return: The cost_characteristic of this TransferCostPac.
        :rtype: List[CostCharacteristic]
        """
        return self._cost_characteristic

    @cost_characteristic.setter
    def cost_characteristic(self, cost_characteristic: List[CostCharacteristic]):
        """Sets the cost_characteristic of this TransferCostPac.

        The list of costs where each cost relates to some aspect of the TopologicalEntity.  # noqa: E501

        :param cost_characteristic: The cost_characteristic of this TransferCostPac.
        :type cost_characteristic: List[CostCharacteristic]
        """

        self._cost_characteristic = cost_characteristic
