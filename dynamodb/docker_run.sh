sudo docker network create lambda-local
sudo docker run -d --network lambda-local --name dynamolocal -p 8000:8000 amazon/dynamodb-local:latest -jar DynamoDBLocal.jar -sharedDb