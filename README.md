# Ekow 

A simple api backed by redis and deployed using 


## Running 

- `git clone https://github.com/s1ntaxe770r/ekow`

- `cd ekow && docker-compose up -d`




## Create an entry

```bash
curl -d '{"key":"12","value":"false"}' -H "Content-Type: application/json" -X POST http://localhost:8080/set
```

### Response
```bash
{
  "SUCCESS": {
    "key": "12",
    "Value": "false",
    "status": "SUCCESS"
  }
}

```
## Retrive the entry

```bash 
 curl -d '{"key":"12"}' -H "Content-Type: application/json" -X POST http://localhost:8080/get | jq
```

### Response 
```bash
{
  "SUCCESS": {
    "key": "12",
    "value": "false"
  }
}

```
## OR!

:mag: explore the swagger docs at http://prod.evil.corp/swagger/index.html



# Deploying :rocket:




