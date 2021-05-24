# Optional

- Fabric CA Operations

>   Optional Section Start

```
This entire section is optional.
```
The last step is to import organization's admin user identity in the wallet.


`npm run importAdmin -- -o <orgName>
`

The above command executes importAdmin.js to import the admin user identity into the wallet. The script reads admin identity from the admin profile <orgname>-admin.json and imports it in wallet for executing HLF operations.

The scripts use file system wallet to store the identites. It creates a wallet as per the path specified in ".wallet" field in the connection profile. By default, ".wallet" field is initalized with <orgname>, which means a folder named <orgname> is created in the current directory to store the identities. If you want to create wallet at some other path, modify ".wallet" field in the connection profile before running enroll admin user and any other HLF operations.


Similarly, import admin user identity for each organization.


Refer command help for more details on the arguments passed in the command

`npm run importAdmin -- -h
`

### User identity generation

Execute below commands in the given order to generate new user identites for the HLF organization.

> Note: Before starting with user identity generation steps, ensure that the initial setup of the application is done.

### Set below enviroment variables on azure cloud shell


- Organization name for which user identity is to be generated. Name of new user identity. Identity will be registered with the Fabric-CA using this name.
    - `ORGNAME=<orgname>`
    - `USER_IDENTITY=<username>`

### Register and enroll new user

To register and enroll new user, execute the below command that executes registerUser.js. It saves the generated user identity in the wallet.

`npm run registerUser -- -o $ORGNAME -u $USER_IDENTITY
`

> Note: Admin user identity is used to issue register command for the new user. Hence, it is mandatory to have the admin user identity in the wallet before executing this command. Otherwise, this command will fail.

Refer command help for more details on the arguments passed in the command

`npm run registerUser -- -h
`
>   Optional Section End 
