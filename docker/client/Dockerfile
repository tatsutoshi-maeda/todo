FROM node:12.16.1
WORKDIR /client
RUN yarn install
COPY client/. .
ENTRYPOINT [ "yarn", "start" ]