# Run native HLF operations
To help customers get started with executing Hyperledger native commands on HLF network on AKS. The sample application is provided that uses fabric NodeJS SDK to perform the HLF operations. The commands are provided to Create new user identity and install your own chaincode.

### Before you begin
Follow the below commands for the initial setup of the application:

- [ Download application files](#downloadFiles)
- [ Generate connection profile and admin profile](#profileGen)
- [Import admin user identity](#importAdmin)

After completing the initial setup, you can use the SDK to achieve the below operations:
-  [User identity generation](#fabricca)
-  [Chaincode operations](#chaincode)

The above-mentioned commands can be executed from Azure Cloud Shell.

<a name="downloadFiles"></a>
### Download application files
The first setup for running application is to download all the application files in a folder say ```app```.

Create ```app``` folder and enter into the folder:
```bash
mkdir app
cd app
```

Execute below command to download all the required files and packages:
```bash
curl https://raw.githubusercontent.com/Azure/Hyperledger-Fabric-on-Azure-Kubernetes-Service/master/application/setup.sh | bash
```

This command takes some time as it loads all the packages. After successful execution of the command, you can see a ```node_modules``` folder in the current directory. All the required packages are loaded in the ```node_modules``` folder.

<a name="profileGen"></a>
### Generate connection profile and admin profile
Create ```profile``` directory inside the ```app``` folder
```bash
cd app
mkdir ./profile
```

Set these environment variables on Azure cloud shell
```bash
# Organization name whose connection profile is to be generated
ORGNAME=<orgname>
# Organization AKS cluster resource group
AKS_RESOURCE_GROUP=<resourceGroup>
```

Execute below comand to generate connection profile and admin profile of the organization
```bash
./getConnector.sh $AKS_RESOURCE_GROUP | sed -e "s/{action}/gateway/g"| xargs curl > ./profile/$ORGNAME-ccp.json
./getConnector.sh $AKS_RESOURCE_GROUP | sed -e "s/{action}/admin/g"| xargs curl > ./profile/$ORGNAME-admin.json
./getConnector.sh $AKS_RESOURCE_GROUP | sed -e "s/{action}/msp/g"| xargs curl > ./profile/$ORGNAME-msp.json
```
It will create connection profile,admin profile and msp profile of the organization inside the ```profile``` folder with name ```<orgname>-ccp.json``` and ```<orgname>-admin.json``` respectively.

Similarly, generate connection profile, admin profile and msp profile for each orderer and peer organization.

<a name="importAdmin"></a>
### Import admin user identity
The last step is to import organization's admin user identity in the wallet.

```
npm run importAdmin -- -o <orgName>
```
The above command executes importAdmin.js to import the admin user identity into the wallet. The script reads admin identity from the admin profile ```<orgname>-admin.json``` and imports it in wallet for executing HLF operations.

The scripts use file system wallet to store the identites. It creates a wallet as per the path specified in ".wallet" field in the connection profile. By default, ".wallet" field is initalized with ```<orgname>```, which means a folder named ```<orgname>``` is created in the current directory to store the identities. If you want to create wallet at some other path, modify ".wallet" field in the connection profile before running enroll admin user and any other HLF operations.

Similarly, import admin user identity for each organization.

Refer command help for more details on the arguments passed in the command
```
npm run importAdmin -- -h
```

<a name="fabricca"></a>
### User identity generation
Execute below commands in the given order to generate new user identites for the HLF organization.

> **_Note:_** Before starting with user identity generation steps, ensure that the initial setup of the application is done.

#### Set below enviroment variables on azure cloud shell
```
# Organization name for which user identity is to be generated
ORGNAME=<orgname>
# Name of new user identity. Identity will be registered with the Fabric-CA using this name.
USER_IDENTITY=<username>
```
#### Register and enroll new user
To register and enroll new user, execute the below command that executes registerUser.js. It saves the generated user identity in the wallet.
```bash
npm run registerUser -- -o $ORGNAME -u $USER_IDENTITY
```
> **_Note:_** Admin user identity is used to issue register command for the new user. Hence, it is mandatory to have the admin user identity in the wallet before executing this command. Otherwise, this command will fail.

Refer command help for more details on the arguments passed in the command
```bash
npm run registerUser -- -h
```
<a name="chaincode"></a>

### Consortium Creation

    cd Hyperledger-Fabric-On-AKS/azhlfTool 
    npm install
    npm run setup
Set up environment variables
Set environment variables for the orderer organization's client

    ORDERER_ORG_SUBSCRIPTION=<ordererOrgSubscription>
    ORDERER_ORG_RESOURCE_GROUP=<ordererOrgResourceGroup>
    ORDERER_ORG_NAME=<ordererOrgName>
    ORDERER_ADMIN_IDENTITY="admin.$ORDERER_ORG_NAME"
    CHANNEL_NAME=<channelName>

Set environment variables for the peer organization's client

    PEER_ORG_SUBSCRIPTION=<peerOrgSubscritpion>
    PEER_ORG_RESOURCE_GROUP=<peerOrgResourceGroup>
    PEER_ORG_NAME=<peerOrgName>
    PEER_ADMIN_IDENTITY="admin.$PEER_ORG_NAME"
    CHANNEL_NAME=<channelName>

Set environment variables for an Azure storage account

    STORAGE_SUBSCRIPTION=<subscriptionId>
    STORAGE_RESOURCE_GROUP=<azureFileShareResourceGroup>
    STORAGE_ACCOUNT=<azureStorageAccountName>
    STORAGE_LOCATION=<azureStorageAccountLocation>
    STORAGE_FILE_SHARE=<azureFileShareName>

Use the following commands to create an Azure storage account. If you already have Azure storage account, skip this step.

    az account set --subscription $STORAGE_SUBSCRIPTION
    az group create -l $STORAGE_LOCATION -n $STORAGE_RESOURCE_GROUP
    az storage account create -n $STORAGE_ACCOUNT -g  $STORAGE_RESOURCE_GROUP -l $STORAGE_LOCATION --sku Standard_LRS

Use the following commands to create a file share in the Azure storage account. If you already have a file share, skip this step.

    STORAGE_KEY=$(az storage account keys list --resource-group $STORAGE_RESOURCE_GROUP  --account-name $STORAGE_ACCOUNT --query "[0].value" | tr -d '"')
    az storage share create  --account-name $STORAGE_ACCOUNT  --account-key $STORAGE_KEY  --name $STORAGE_FILE_SHARE

Use the following commands to generate a connection string for an Azure file share.

    STORAGE_KEY=$(az storage account keys list --resource-group $STORAGE_RESOURCE_GROUP  --account-name $STORAGE_ACCOUNT --query "[0].value" | tr -d '"')
    SAS_TOKEN=$(az storage account generate-sas --account-key $STORAGE_KEY --account-name $STORAGE_ACCOUNT --expiry `date -u -d "1 day" '+%Y-%m-%dT%H:%MZ'` --https-only --permissions lruwd --resource-types sco --services f | tr -d '"')
    AZURE_FILE_CONNECTION_STRING=https://$STORAGE_ACCOUNT.file.core.windows.net/$STORAGE_FILE_SHARE?$SAS_TOKEN

#### Import an organization connection profile, admin user identity, and MSP

Use the following commands to fetch the organization's connection profile, admin user identity, and Managed Service Provider (MSP) from the Azure Kubernetes Service cluster and store these identities in the client application's local store. An example of a local store is the azhlfTool/stores directory.

For the orderer organization:



    ./azhlf adminProfile import fromAzure -o $ORDERER_ORG_NAME -g $ORDERER_ORG_RESOURCE_GROUP -s $ORDERER_ORG_SUBSCRIPTION
    ./azhlf connectionProfile import fromAzure -g $ORDERER_ORG_RESOURCE_GROUP -s $ORDERER_ORG_SUBSCRIPTION -o $ORDERER_ORG_NAME   
    ./azhlf msp import fromAzure -g $ORDERER_ORG_RESOURCE_GROUP -s $ORDERER_ORG_SUBSCRIPTION -o $ORDERER_ORG_NAME

For the peer organization:


    ./azhlf adminProfile import fromAzure -g $PEER_ORG_RESOURCE_GROUP -s $PEER_ORG_SUBSCRIPTION -o $PEER_ORG_NAME
    ./azhlf connectionProfile import fromAzure -g $PEER_ORG_RESOURCE_GROUP -s $PEER_ORG_SUBSCRIPTION -o $PEER_ORG_NAME
    ./azhlf msp import fromAzure -g $PEER_ORG_RESOURCE_GROUP -s $PEER_ORG_SUBSCRIPTION -o $PEER_ORG_NAME

#### Add a peer organization for consortium management

Run the following commands in the given order to add a peer organization in a channel and consortium:


From the peer organization's client, upload the peer organization's MSP on Azure Storage.

    ./azhlf msp export toAzureStorage -f  $AZURE_FILE_CONNECTION_STRING -o $PEER_ORG_NAME

From the orderer organization's client, download the peer organization's MSP from Azure Storage. Then issue the command to add the peer organization in the channel and consortium.

    ./azhlf msp import fromAzureStorage -o $PEER_ORG_NAME -f $AZURE_FILE_CONNECTION_STRING
    ./azhlf consortium join -o $ORDERER_ORG_NAME  -u $ORDERER_ADMIN_IDENTITY -p $PEER_ORG_NAME

#### Navigate

Navigate to setupcli for channel and chaincode operations

    cd ..
    cd <rootDir>/Hyperledger-Fabric-On-AKS/setupFabricCli
