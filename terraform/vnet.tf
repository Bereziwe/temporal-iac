
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test-rg" {
  name     = "example-resources"
  location = "East US"
}

resource "azurerm_virtual_network" "test-vnet" {
  name                = "test-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.test-rg.location
  resource_group_name = azurerm_resource_group.test-rg.name
}

resource "azurerm_subnet" "test-subnet" {
  name                 = "test-subnet"
  resource_group_name  = azurerm_resource_group.test-rg.name
  virtual_network_name = azurerm_virtual_network.test-vnet.name
  address_prefixes     = ["10.0.1.0/24"]
}
