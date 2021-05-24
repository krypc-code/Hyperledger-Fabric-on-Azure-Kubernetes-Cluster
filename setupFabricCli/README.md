##Setup Fabric Cli

Open link below to set up fabric cli by following instructions provided on its readme.

    https://github.com/hyperledger/fabric-cli

Generate cryptomaterials from connection profile, admin profile, msp profile for peer and orderer organizations

For Peer organization
    
    cd <rootDir>/Hyperledger-Fabric-On-AKS/setupFabricCli
    PEER_CONNECTION_PROFILE_PATH=<path-to-peerConnectionProfile.json>
    PEER_ADMIN_PROFILE_PATH=<path-to-peerAdminProfile.json>
    PEER_MSP_PROFILE=<path-to-peerMSPProfile.json>
    ./certgen.sh peer $PEER_ADMIN_PROFILE_PATH $PEER_CONNECTION_PROFILE_PATH $PEER_MSP_PROFILE
    
For Orderer organization

    ORDERER_CONNECTION_PROFILE_PATH=<path-to-ordererConnectionProfile.json>
    ORDERER_ADMIN_PROFILE_PATH=<path-to-ordererAdminProfile.json>
    ORDERER_MSP_PROFILE=<path-to-ordererMSPProfile.json>
    ./certgen.sh orderer $ORDERER_ADMIN_PROFILE_PATH $ORDERER_CONNECTION_PROFILE_PATH $ORDERER_MSP_PROFILE

Set up fabric cli environment

    export FABRIC_EXECUTABLE_PATH =<path_to_fabric_binary>
    export PEER_CONNECTION_PROFILE_PATH=<path_to_peer_connection_profile.json>
    export ORDERER_CONNECTION_PROFILE_PATH=<path_to_orderer_connection_profile.json>

    cd genFabricCliGO/main/
    go build -o ../../configCoversion
    cd ../../
    ./configCoversion (this command will create fabric go sdk based config file with name peerorg.yaml)

Initialize fabric cli

    $FABRIC_EXECUTABLE_PATH network set <network-name> <path-to-go-sdk-config.yaml>
    $FABRIC_EXECUTABLE_PATH context set <context-name> --channel <channel-name> --network <network-name> --organization <peer-orgname> --user <admin-user>
    $FABRIC_EXECUTABLE_PATH context use <context-name>

### Channel operations

Use the following command to pull binaries of HLF 2.2
    
    cd genFabricCliGO
    curl -sSL https://goo.gl/6wtTN5 | bash -s 2.2.0
    cat tempConfigtx.yaml | sed -e "s/OrgName/<peerOrgName>/g" > configtx.yaml
    bin/configtxgen -profile OrgsChannel -outputCreateChannelTx ./channel-artifacts/<channel-name>.tx -channelID <channel-name>
    $FABRIC_EXECUTABLE_PATH channel create <channel-name> ./channel-artifacts/<channel-name>.tx
    $FABRIC_EXECUTABLE_PATH channel join <channel-name>

### Chaincode operations

##### Package chaincode

    $FABRIC_EXECUTABLE_PATH lifecycle package <chaincode-label> <chaincode-type> <path>

#### Install chaincode

    $FABRIC_EXECUTABLE_PATH lifecycle install <chaincode-label> <path>
    
#### Approve chaincode

    $FABRIC_EXECUTABLE_PATH lifecycle approve <chaincode-name> <version> <package-id> <sequence> --policy <policy string>

#### Commit chaincode

    $FABRIC_EXECUTABLE_PATH lifecycle commit <chaincode-name> <version> <sequence> --policy <policy>


#### Invoke chaincode

    $FABRIC_EXECUTABLE_PATH  chaincode invoke -h

#### Query chaincode

    $FABRIC_EXECUTABLE_PATH chaincode query -h

#### Create Channel With Custom Policy

New channel with custom policy can be created by editing the policy in the "tempConfigtx.yaml" and then executing the commands in the "Channel Operation" section 
to create a new channel with those custom policies.

For example, let's assume for Application Admin policy we want to change the permissions where in place of Majority only "orgName" organization's admin signatures 
are required to approve the changes. In that case our application section in the yaml file will look like as below.

     Application: &ApplicationDefaults
     
         # Organizations is the list of orgs which are defined as participants on
         # the application side of the network
         Organizations:
     
         # Policies defines the set of policies at this level of the config tree
         # For Application policies, their canonical path is
         #   /Channel/Application/<PolicyName>
         Policies:
             Readers:
                 Type: ImplicitMeta
                 Rule: "ANY Readers"
             Writers:
                 Type: ImplicitMeta
                 Rule: "ANY Writers"
             Admins:
                 Type: Signature
                 Rule: "OR('OrgName.admin')"
             # LifecycleEndorsement:
             #     Type: ImplicitMeta
             #     Rule: "MAJORITY Endorsement"
             # Endorsement:
             #     Type: ImplicitMeta
             #     Rule: "MAJORITY Endorsement"
     
         Capabilities:
             <<: *ApplicationCapabilities
     
After making this change you generate the genesis and channel files using commands as in the "Channel Operation" section and you will get a channel with custom 
policy of you own.