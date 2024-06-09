# syntax=docker/dockerfile:1
FROM node:20.12.0

# Install xdg-utils for the xdg-open command
RUN apt-get update && apt-get install -y xdg-utils

WORKDIR /app

COPY package.json ./
COPY yarn.lock ./
RUN yarn install

COPY . ./

EXPOSE 8083

CMD ["yarn", "dev"]
