# MERINIO API #

### How to run the app ? ###
`` docker-compose up``

 

### How to test ? ###

Either Postman 

or 
```
curl --location 'http://localhost:8080/api/branches'

curl --location 'http://localhost:8080/api/branches' \
--header 'Content-Type: application/json' \
--data '{
"name": "branch ajout",
"parent_id": 4,
"requirements": [6,7],
"restrictions":[8]
}'
```