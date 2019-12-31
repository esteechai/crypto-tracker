#!/bin/bash
SSH_KEY_PATH=/home/estee/.ssh/server_access.pub
SSH_USER=estee
HOST_PATH=interns.theninja.life

echo "build go binary start..."
cd /home/estee/go/src/crypto-tracker/server
rm server
go build

echo "update server binary..."
scp -i $SSH_KEY_PATH server $SSH_USER@$HOST_PATH:/home/estee/crypto_tracker/backend

echo "update template..."
rsync -zarvh --delete /home/estee/go/src/crypto-tracker/server/template/. -e ssh $SSH_USER@$HOST_PATH:/home/estee/crypto_tracker/backend/template

echo "update migrations..."
rsync -zarvh --delete /home/estee/go/src/crypto-tracker/server/migrations/. -e ssh $SSH_USER@$HOST_PATH:/home/estee/crypto_tracker/migrations

ssh -i $SSH_KEY_PATH $SSH_USER@$HOST_PATH migrate -database "postgres://esteecoinbase:dev@localhost:5432/esteecoinbase?sslmode=disable" -path ./crypto_tracker/migrations up

echo "building frontend..."
cd /home/estee/go/src/crypto-tracker/web
rm -rf build 
npm run build 


rsync -zarvh --delete build/. -e ssh $SSH_USER@$HOST_PATH:/home/estee/crypto_tracker/frontend
