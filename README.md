go get -u "github.com/stretchr/testify/assert"


curl http://localhost:8000/person/all  |python -m json.tool

curl -X POST http://127.0.0.1:8000/person/add -d 'firstName=chen&lastName=kejun'