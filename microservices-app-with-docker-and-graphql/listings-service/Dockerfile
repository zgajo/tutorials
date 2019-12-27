FROM node:12

COPY . /opt/app

WORKDIR /opt/app

RUN yarn

RUN yarn build

CMD yarn watch