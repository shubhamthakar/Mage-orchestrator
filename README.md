
# Mage Orchestrator



## Run Locally

Clone the project

```bash
  git clone https://github.com/shubhamthakar/Mage-orchestrator.git
```

Go to the project directory

```bash
  cd Mage
```

Install dependencies(should have golang installed.)

```bash
   go version
```

Start the datamonk server (in a new terminal)

```bash
  go run datamonk/server/server.go 
```

Start the orchestrator2 server (in a new terminal)

```bash
  go run orchestrator/server/server.go
```

Start the orchestrator1 server (in a new terminal)

```bash
  go run server/server.go
```

Start the client (in a new terminal)

```bash
  go run client/client.go
```


  



Documentation
 https://docs.google.com/document/d/1G1_BiccSTLkZS9FUcnmoHincv7W_7tY4y0aSEep3hdg/edit?usp=sharing
