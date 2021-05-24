## Hyperledger Fabric on Azure Kubernetes Service using ARM🥁


- Microsoft provides ARM template for creating HLF network on AKS via `https://github.com/Azure/Hyperledger-Fabric-on-Azure-Kubernetes-Service`.
Since this repo only covers HLF 1.4 and there is a need for HL 2.x in the market due it's widely popular features and improvments, we have created this repository with HLF 2.2. support following the same order as the Microsoft Sample.
  
- In this repo we create custom ARM template for resource utilization on Azure and provide commands for executing the whole lifecycle.
This repository carries custom built fabric-go-cli as added feature to enhance developer experience.
The whole process is explained in stages through different Readme files, each describing a stage in the lifecycle.


- The whole process of creating infrastucture, deploying fabric components , creation of consortiums and channels is summarized here.
There types pf profiles are required for the whole process to work which are :

- Connection profile : Connection profile contains the urls about peer orderer etc.
- Admin profile : Admin profile contains certificates of admin , name of admin etc.
- MSP profile : TLS certificates for organisations and other credentials.

For performing operations to create these certificates ,you can create an app directory which will pull content from `getconnecter.sh` .
The shell script will make the profiles required for custom deployment.

- On completion of profile generation environment variables has to be set on azure for orgname and resource group of peer.
- Same procedure is to be followed for orderer till the environment variables are set for it.


- If the user is interested the user can follow steps that guide in creation of custom user for fabric ca by following , fabric ca documentation.

- User can set environment variables for storage account to be used following the instructions. This step is to be completed after creation of a storage accounr and creation of a container inside the storage account on azure cloud. The steps are provided.
- Once the content is moved a dynamically generated sas token must be created and used in the template.


After setting up the storage account fabric cli can be initialized from setupFabricCli folder.
You can generate crypto material and create folder structure using certgen as mentioned in Hyperledger fabric cli go documentation.
This folder structure is needed for fabric cli.

- Fabric cli environment will setup the yaml file for use by go sdk ,once connection profile is set we can run go code for the command line interface.

        -- go build
- Store and execute build to initialize fabric cli.
- Following the fabric cli documentation , you have to create a network and link context to the network.
- Using this context fabric cli can be initialized.

Once consortium is done you can create channel, join channel and install chaincode.

The last section on external chaincode execution can be performed by using connection and metadata.json as explained.
Connection will have things like address port etc,metadata will have type and label of chaincode etc.



## Follow 🛠
- [ARM Template Deployment : For customization of microsoft arm template for user defined organisation details.](fabric-part-1.md)
- [Generating Profiles : The admin,msp and connection profiles for orderer and peer.](fabric-part-2.md)
- [Fabric CA Operations : Step if you wish to create a custom user for fabric ca.](fabric-ca.md)
- [Consortium Creation : Creation of multiple organisations and their crypto materials.](fabric-part-3.md)
- [Setup Fabric Cli Go : Command line interface for interacting with the blockchain system.](fabric-part-4.md)
- [Channel Operations : Creation of channels and basic operations on them.](fabric-part-5.md)
- [Chaincode Operations : Basic chaincode interactions.](fabric-part-6.md)
- [External Chaincode : Executing chaincode on an external container. This an advanced feature.](fabric-part-7.md)
- Example chaincode `asset-transfer-basic` and sample commands are provided in `chaincode-sample` folder.
