# Oracle cloud IDCS iBridge checker

The IDCS component of the Oracle cloud (<https://cloud.oracle.com/en_US/identity/features)> authenticate and authorize users and apps.
It can be integrated through a .NET app with on prem Microsoft AD.

This go program check the status of ibridge instances and report on synchronized users and groups.

## Steps to config

config.json is the configuration file for the program to obtain :
    a-Tenant name
    b-IDCS instance base URL
    c-Client ID
    d-Client Secret
Client ID and Client secret are obtained after registering and activating the application -this program- in the application section of IDCS. That is the first time, then this program use that to obtain the oauth2 token to execute the API calls.

Hope it helps to have initial visibility over status of ibridge/IDCS synchronization

