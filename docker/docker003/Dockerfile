FROM node:8-alpine
LABEL maintainer "tallongsun@gmail.com"

ENV PROJECT_DIR=/app

WORKDIR  $PROJECT_DIR
COPY package.json  $PROJECT_DIR
RUN npm install
COPY .  $PROJECT_DIR

ENV APP_PORT=8080

EXPOSE $APP_PORT
HEALTHCHECK CMD curl --fail http://localhost:$APP_PORT || exit 1

CMD ["npm" ,"start"]
