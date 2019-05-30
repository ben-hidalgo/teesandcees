FROM node:12.3.1-stretch

RUN apt-get update
RUN apt-get install -q -y vim

WORKDIR /usr/local/sharedb

ADD package.json package.json
ADD package-lock.json package-lock.json

RUN npm install

COPY . .

#CMD ["node", "server.js"]
CMD ["sleep", "9999999"]
