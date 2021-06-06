# Setting Up Fabric Go Cli
![gopher.png](../images/gopher.png)



- The Hyperledger Fabric CLI is a tool used to interact with Fabric networks.


## Installation

- >     Note : Follow fabric-cli documentation for more info


`    git clone https://github.com/hyperledger/fabric-cli
`

- Add the binary to your PATH

- Execute fabric for more information


### Getting Started


```

    Add a Network with fabric network set
    Add a Context with fabric context set
    Use the new context with fabric context use
    You're all set... Have fun!

```


#### Network

- A network is a direct reference to a Fabric-SDK-Go configuration. This configuration contains all of the necessary details for interacting with a Fabric network at a global scope.

#### Context

- A context defines the scope for interactions with the network. An example of this would be: As Admin, I want peer peer0.org1.example.com in organization Org1 to join channel mychannel. In this example, the context would include the identity, peer, organization, and channel.

#### Built-in Commands

- Built-in commands can be found in /cmd/fabric/commands.
  https://github.com/hyperledger/fabric-cli/blob/main/cmd/fabric/commands
- These commands can serve as examples for building future commands like
- `plugin chaincode install ....`


## Generate cryptomaterials from connection profile, admin profile, msp profile for peer and orderer organizations

- For Peer organization


```
cd <rootDir>/Hyperledger-Fabric-On-AKS/setupFabricCli
PEER_CONNECTION_PROFILE_PATH=<path-to-peerConnectionProfile.json>
PEER_ADMIN_PROFILE_PATH=<path-to-peerAdminProfile.json>
PEER_MSP_PROFILE=<path-to-peerMSPProfile.json>
./certgen.sh peer $PEER_ADMIN_PROFILE_PATH $PEER_CONNECTION_PROFILE_PATH $PEER_MSP_PROFILE
```


- For Orderer organization



```
ORDERER_CONNECTION_PROFILE_PATH=<path-to-ordererConnectionProfile.json>
ORDERER_ADMIN_PROFILE_PATH=<path-to-ordererAdminProfile.json>
ORDERER_MSP_PROFILE=<path-to-ordererMSPProfile.json>
./certgen.sh orderer $ORDERER_ADMIN_PROFILE_PATH $ORDERER_CONNECTION_PROFILE_PATH $ORDERER_MSP_PROFILE

```


### Set up fabric cli environment
- The golang program takes peer and orderer connection profile and using that information it will generate the yaml file required for go cli in the current directory. This go program is built first and then used to execute actions needed for cli preparation.

```
export FABRIC_EXECUTABLE_PATH =<path_to_fabric_binary>
export PEER_CONNECTION_PROFILE_PATH=<path_to_peer_connection_profile.json>
export ORDERER_CONNECTION_PROFILE_PATH=<path_to_orderer_connection_profile.json>

cd genFabricCliGOConfig/main/
go build -o ../../configCoversion
cd ../../
./configCoversion (this command will create fabric go sdk based config file with name peerorg.yaml)

```

> Note : If there is another Peer Organisation the whole process must be duplicated till the point of creation of the new yaml file representing the additional organisation.
This must carry a seperate context for the peer organisation.


### Initialize fabric cli



```
$FABRIC_EXECUTABLE_PATH network set <network-name> <path-to-go-sdk-config.yaml>
$FABRIC_EXECUTABLE_PATH context set <context-name> --channel <channel-name> --network <network-name> --organization <peer-orgname> --user <admin-user>
$FABRIC_EXECUTABLE_PATH context use <context-name>
```
  
## Follow Up
- Continue deployment by following channel operations.
- [Channel Operations](ChannelOperations.md)
  