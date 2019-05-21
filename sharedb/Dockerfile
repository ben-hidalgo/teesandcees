FROM node:10.15.3-jessie

WORKDIR /usr/local/sharedb

ADD package.json package.json
ADD package-lock.json package-lock.json

RUN npm i

COPY . .

CMD ["node", "server.js"]
#CMD ["sleep", "9999999"]
