# Deploying the template 🛠

- The template can be deployed using Azure CLI or Powershell.

- Sample command to deploy from Azure CLI is as follows:
- `az deployment group create --name <deployment> --resource-group <resourceGroup> --subscription <subscriptionId> --template-uri https://<baseURI>/mainTemplate.json`


## Next Steps

-Once Peer and Orderer are deployed


# Running native HLF operations
![fabricApp.png](images/Deployment.png)


- Sample application for performing the HLF operations. The commands are provided to Create new user identity and install your own chaincode.

## Download application files
- The first setup for running application is to download all the application files in a folder say app.
- Create app folder move inside app:
- `mkdir app
  cd app`

- Execute below command to download all the required files and packages:
- ``curl https://raw.githubusercontent.com/Azure/Hyperledger-Fabric-on-Azure-Kubernetes-Cluster/master/application/setup.sh | bash``


- This command takes some time as it loads all the packages. After successful execution of the command, you can see a `node_modules` folder in the current directory. All the required packages are loaded in the `node_modules` folder.


## Generate connection profile and admin profile

- Create profile directory inside the app folder
- `cd app
  mkdir ./profile`

- Set these environment variables on Azure cloud shell

- `ORGNAME=<orgname>`
- `AKS_RESOURCE_GROUP=<resourceGroup>`

- Execute below comand to generate connection profile and admin profile of the organization
#
- `./getConnector.sh $AKS_RESOURCE_GROUP | sed -e "s/{action}/gateway/g"| xargs curl > ./profile/$ORGNAME-ccp.json`


- `./getConnector.sh $AKS_RESOURCE_GROUP | sed -e "s/{action}/admin/g"| xargs curl > ./profile/$ORGNAME-admin.json`


- `./getConnector.sh $AKS_RESOURCE_GROUP | sed -e "s/{action}/msp/g"| xargs curl > ./profile/$ORGNAME-msp.json`


It will create connection profile,admin profile and msp profile of the organization inside the profile folder with name <orgname>-ccp.json and <orgname>-admin.json respectively.


- Similarly, generate connection profile, admin profile and msp profile for each orderer and peer organization.

  ## Follow up
   -  Continue deployment by creating consortiums.
   -  [Create Consortium](CreatingConsortiums.md)
