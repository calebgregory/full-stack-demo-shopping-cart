resource="$RESOURCE"
id="$ID"
create="$CREATE"
update="$UPDATE"

if [ -z "$resource" ] || [ -z "$id" ] || [ -z "$create" ] || [ -z "$update" ]
then
  echo "Must include RESOURCE, ID, CREATE, UPDATE as env var"
  exit 1
fi

url="localhost:3000/$resource/get-all"
printf "\n$url\n  "
curl -X POST "$url" -H "Content-Type: application/json" -d '{}'

url="localhost:3000/$resource/create"
printf "\n$url\n  "
curl -X POST "$url" -H "Content-Type: application/json" -d "{\"$resource\":{\"id\":$id,$create}}"

url="localhost:3000/$resource/update"
printf "\n$url\n  "
curl -X POST "$url" -H "Content-Type: application/json" -d "{\"$resource\":{\"id\":$id,$update}}"

url="localhost:3000/$resource/delete"
printf "\n$url\n  "
curl -X POST "$url" -H "Content-Type: application/json" -d "{\"$resource\":{\"id\":$id}}"

url="localhost:3000/$resource/get-one"
printf "\n$url\n  "
curl -X POST "$url" -H "Content-Type: application/json" -d "{\"id\":$id}"

