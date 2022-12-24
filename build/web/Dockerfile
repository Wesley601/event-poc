FROM node:18-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY web/package.json /app/package.json
COPY ./web /app

EXPOSE 5173

RUN yarn

CMD ["yarn", "dev"]
