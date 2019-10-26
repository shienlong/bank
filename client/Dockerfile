FROM node:13-alpine
ARG NODE_ENV
ENV NODE_ENV=development
ARG APP_ENV
ENV APP_ENV=development

RUN mkdir /app
WORKDIR /app

COPY package.json /app/package.json
COPY yarn.lock /app/yarn.lock
RUN yarn install

COPY . /app