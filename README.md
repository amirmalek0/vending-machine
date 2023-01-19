# vending-machine
vending machine

this app could run in two mod
# 1 - restAPI
     - run via docker:
        docker-compose up -d --build
        or go run main.go
        it will serve app on 8000 port
        using postman file located doc directory you can make request

# 2 - command line interface( cli )
Instead of using restAPI you can use command line interface for call app functions 
### Step 1 
build app using below command
```bash
go build
```
### Step 2 two 
Print available machine list

```bash
./vending-machine machines
```

Buy a coca
```bash
./vending-machine coca MACHINE_NAME COIN(int)
```

Buy a coffee
```bash
./vending-machine coffee MACHINE_NAME COIN(int)
```
