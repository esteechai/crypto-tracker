#!/bin/bash
SSH_KEY_PATH=/home/estee/.ssh/server_access.pub
SSH_USER=estee
HOST_PATH=interns.theninja.life

echo "building frontend..."
cd /home/estee/go/src/crypto-tracker/web
rm -rf build 
npm run build 


rsync -zarvh --delete build/. -e ssh $SSH_USER@$HOST_PATH:/home/estee/crypto_tracker/frontend