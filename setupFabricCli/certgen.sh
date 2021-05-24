#!/bin/bash

printPeerCommandHelp() {
    echo
    echo "======= Please provide the following arguments correctly (in order): ======="
    echo -e "\tPeer Admin Credentials Path"
    echo -e "\t\t- You can download the Admin Credentials JSON file from Azure Portal UI, by selecting your Blockchain Resource > Overview pane > Admin Credentials."
    echo
    echo -e "\tPeer Connection Profile Path"
    echo -e "\t\t- You can download the Connection Profile JSON file from Azure Portal UI, by selecting your Blockchain Resource > Overview pane > Connection Profile."
    echo
    echo -e "\tPeer MSP Configuration Path"
    echo -e "\t\t- You can download the MSP Configuration JSON file from Azure Portal UI, by selecting your Blockchain Resource > Overview pane > MSP Configuration."
    echo
    echo -e "\tPeer node name: e.g. \"peer<peer#>\""
    echo
    echo -e "\tOrderer Connection Profile Path"
    echo -e "\t\t- You can download the Connection Profile JSON file from Azure Portal UI, by selecting your Blockchain Resource > Overview pane > Connection Profile."
    echo
    echo -e "\tOrderer node name: e.g. \"orderer<orderer#>\""
    echo
    echo "======= Example: ======="
    echo
    peerProfileRootPath=/var/hyperledger/profiles/peerprofiles/peerOrg
    ordererProfileRootPath=/var/hyperledger/profiles/peerprofiles/ordererOrg
    echo -e "\tsource setupFabricCLI.sh \"peer\" ${peerProfileRootPath}/peerOrg_AdminCredential.json ${peerProfileRootPath}/peerOrg_ConnectionProfile.json ${peerProfileRootPath}/peerOrg_MSPConfiguration.json \"peer1\" ${ordererProfileRootPath}/ordererOrg_ConnectionProfile.json \"orderer1\""
    echo
}

printOrdererCommandHelp() {
    echo
    echo "======= Please provide the following arguments correctly (in order): ======="
    echo -e "\tOrderer Admin Credentials Path"
    echo -e "\t\t- You can download the Admin Credentials JSON file from Azure Portal UI, by selecting your Blockchain Resource > Overview pane > Admin Credentials."
    echo
    echo -e "\tOrderer Connection Profile Path"
    echo -e "\t\t- You can download the Connection Profile JSON file from Azure Portal UI, by selecting your Blockchain Resource > Overview pane > Connection Profile."
    echo
    echo -e "\tOrderer MSP Configuration Path"
    echo -e "\t\t- You can download the MSP Configuration JSON file from Azure Portal UI, by selecting your Blockchain Resource > Overview pane > MSP Configuration."
    echo
    echo -e "\tOrderer node name: e.g. \"orderer<orderer#>\""
    echo
    echo "======= Example: ======="
    echo
    ordererProfileRootPath=/var/hyperledger/profiles/ordererprofiles/ordererOrg
    echo -e "\tsource setupFabricCLI.sh \"orderer\" ${ordererProfileRootPath}/ordererOrg_AdminCredential.json ${ordererProfileRootPath}/ordererOrg_ConnectionProfile.json ${ordererProfileRootPath}/ordererOrg_MSPConfiguration.json \"orderer1\""
    echo
}

isFile() {
    if ! [[ "$1" == /* ]]; then
        echo
        echo "The path: \"$1\" is not an absolute path! Please give absolute path to the file!"
        echo
        
        return 1
    elif [ -r "$1" ]; then
        return 0 
    else
        echo
        echo "Unable to find file: $1 or it does not have read access! Please give absolute path to the file!"
        echo
        
        return 1
    fi
}

checkNodeName() {
    node=$(cat $2 | jq '.'$1's.'$3'' | sed 's/grpcs:\/\///g' | tr -d '"')
    if [ "$node" = null ]; then
        echo
        echo "Invalid node name: \"$3\" OR invalid connection profile file for $1 organization"
        echo

        return 1
    else
        return 0
    fi
}

createOrgMSP() {
    orgName=$1
    adminProfilePath=$2
    connectionProfilePath=$3
    mspProfilePath=$4

    rm -rf ./cryptoconfig/${nodeType}/${orgName}/*
    mkdir -p ./cryptoconfig/${nodeType}/${orgName}/msp/{admincerts,cacerts,keystore,signcerts,tlscacerts}
    mkdir -p ./cryptoconfig/${nodeType}/${orgName}/users/admin.${orgName}/msp/{admincerts,cacerts,keystore,signcerts,tlscacerts}
    mkdir -p ./cryptoconfig/${nodeType}/${orgName}/tls
    mkdir -p ./cryptoconfig/${nodeType}/${orgName}/ca/tlsca
    #admincerts
    cat ${adminProfilePath} | jq '.cert' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/msp/admincerts/cert.pem
    cat ${adminProfilePath} | jq '.cert' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/users/admin.${orgName}/msp/admincerts/cert.pem

    #cacerts
    cat ${mspProfilePath} | jq '.cacerts' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/msp/cacerts/rca.pem
    cat ${mspProfilePath} | jq '.cacerts' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/users/admin.${orgName}/msp/cacerts/rca.pem


    cat $connectionProfilePath | jq '.certificateAuthorities.'${orgName}'CA.tlsCACerts.pem' | tr -d '"' | sed 's/\\n/\n/g' > ./cryptoconfig/${nodeType}/${orgName}/ca/tlsca/cert.pem

   for row in  $(cat $connectionProfilePath | jq '.organizations.'${orgName}'.'$nodeType's[]'); do

    folderName=$(sed -e 's/^"//' -e 's/"$//' <<<"$row")
    checkNodeName ${nodeType} ${connectionProfilePath} ${row}
     mkdir -p ./cryptoconfig/${nodeType}/${orgName}/${folderName}/msp/tlscacerts
    #tlscacerts
    cat ${connectionProfilePath} | jq '.'$nodeType's.'$row'.tlsCACerts.pem' | tr -d '"' | sed 's/\\n/\n/g' > \
    ./cryptoconfig/${nodeType}/${orgName}/$folderName/msp/tlscacerts/ca.crt

   done
   echo "path nodeType/orgName "${nodeType}"/"${orgName}
    #signcerts
    cp ./cryptoconfig/${nodeType}/${orgName}/msp/admincerts/cert.pem ./cryptoconfig/${nodeType}/${orgName}/msp/signcerts/cert.pem
    cp ./cryptoconfig/${nodeType}/${orgName}/msp/admincerts/cert.pem ./cryptoconfig/${nodeType}/${orgName}/users/admin.${orgName}/msp/signcerts/admin.${orgName}@${orgName}-cert.pem


    #keystore
    cat ${adminProfilePath} | jq '.private_key' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/msp/keystore/key.pem
    cat ${adminProfilePath} | jq '.private_key' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/users/admin.${orgName}/msp/keystore/priv_sk


    #admin-tls-cert
    cat ${adminProfilePath} | jq '.tls_cert' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/tls/cert.pem

    #admin-tls-key
    cat ${adminProfilePath} | jq '.tls_private_key' | tr -d '"' | base64 -d > ./cryptoconfig/${nodeType}/${orgName}/tls/key.pem
}

createOrdererTLSCA() {
    rm -rf ./orderer/${ordererOrgName}/*
    mkdir -p ./orderer/${ordererOrgName}/msp/tlscacerts

    #tlscacerts
    cat ${ordererConnectionProfilePath} | jq '.orderers."'$ordererNodeName'.'${ordererOrgName}'".tlsCACerts.pem' | tr -d '"' | sed 's/\\n/\n/g' > \
    ./orderer/${ordererOrgName}/msp/tlscacerts/ca.crt
}

setEnvVars() {
    orgName=$1
    connectionProfilePath=$2
    nodeName=$3

    export CORE_PEER_LOCALMSPID="${orgName}"
    export CORE_PEER_ADDRESS=$(cat ${connectionProfilePath} | jq '.'$nodeType's."'$nodeName'.'${orgName}'".url' | sed 's/grpcs:\/\///g' | tr -d '"')
    export CORE_PEER_TLS_ENABLED="true"
    export CORE_PEER_TLS_ROOTCERT_FILE=$(pwd)/${nodeType}/${orgName}/msp/tlscacerts/ca.crt
    export CORE_PEER_TLS_CLIENTAUTHREQUIRED="true"
    export CORE_PEER_TLS_CLIENTCERT_FILE=$(pwd)/${nodeType}/${orgName}/tls/cert.pem
    export CORE_PEER_TLS_CLIENTKEY_FILE=$(pwd)/${nodeType}/${orgName}/tls/key.pem
    export CORE_PEER_TLS_CLIENTROOTCAS_FILES=$(pwd)/${nodeType}/${orgName}/msp/tlscacerts/ca.crt
    export CORE_PEER_MSPCONFIGPATH=$(pwd)/${nodeType}/${orgName}/msp

    if [ "${nodeType}" = "peer" ]; then
        ordererConnectionProfilePath=$4
        ordererNodeName=$5
        export ORDERER_ENDPOINT=$(cat ${ordererConnectionProfilePath} | jq '.orderers."'$ordererNodeName'.'${ordererOrgName}'".url' | sed 's/grpcs:\/\///g' | tr -d '"')
        export ORDERER_TLS_CERT=$(pwd)/orderer/${ordererOrgName}/msp/tlscacerts/ca.crt
    fi 
}

peerArgsCount=4
ordererArgsCount=4

nodeType=$1
if [ "${nodeType}" = "peer" ]; then
    echo
    echo "======= Configuring environment for your ${nodeType} organization! ======="
    echo

    if [ $# -ne $peerArgsCount ]; then
        echo "Invalid number of arguments while setting up Fabric CLI environment for ${nodeType} organization!"
        printPeerCommandHelp
        return;
    fi

    peerAdminProfilePath=$2
    peerConnectionProfilePath=$3
    peerMSPProfilePath=$4

    if ! isFile ${peerAdminProfilePath} || ! isFile ${peerConnectionProfilePath} \
    || ! isFile ${peerMSPProfilePath}; then
        printPeerCommandHelp
        return;
    fi

    peerOrgName=$(cat ${peerAdminProfilePath} | jq '.msp_id' | tr -d '"')

    createOrgMSP $peerOrgName $peerAdminProfilePath $peerConnectionProfilePath $peerMSPProfilePath 

    echo
    echo "======= Successfully configured environment for your ${nodeType} organization! ======="
    echo
elif [ "${nodeType}" = "orderer" ]; then
    echo
    echo "======= Configuring environment for your ${nodeType} organization! ======="
    echo

    if [ $# -ne $ordererArgsCount ]; then
        echo "Invalid number of arguments while setting up Fabric CLI environment for ${nodeType} organization!"
        printOrdererCommandHelp
        return;
    fi

    ordererAdminProfilePath=$2
    ordererConnectionProfilePath=$3
    ordererMSPProfilePath=$4

    if ! isFile ${ordererAdminProfilePath} || ! isFile ${ordererConnectionProfilePath} \
    || ! isFile ${ordererMSPProfilePath}; then
        printOrdererCommandHelp
        return;
    fi

    ordererOrgName=$(cat ${ordererConnectionProfilePath} | jq '.name' | tr -d '"')

    createOrgMSP $ordererOrgName $ordererAdminProfilePath $ordererConnectionProfilePath $ordererMSPProfilePath 
    # setEnvVars $ordererOrgName $ordererConnectionProfilePath $ordererNodeName

    echo
    echo "======= Successfully configured environment for your ${nodeType} organization! ======="
    echo
else
    echo
    echo "Failed to configure environment for node type: \"${nodeType}\"! Should be either \"peer\" or \"orderer\"!"
    echo
fi
