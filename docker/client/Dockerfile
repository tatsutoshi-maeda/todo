FROM node:12.16.1
WORKDIR /client
RUN npm install
COPY client/. .
ENTRYPOINT [ "npm", "start" ]