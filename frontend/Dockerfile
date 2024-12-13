FROM node:22.12.0-alpine3.21

RUN apk add --no-cache tini
ENTRYPOINT ["/sbin/tini", "--"]

RUN npm install -g npm@latest

WORKDIR /app

ENV TZ=America/Sao_Paulo

COPY --chown=node:node package*.json ./
RUN npm ci --omit=dev && npm cache clean --force
COPY --chown=node:node . .

EXPOSE 3000
USER node
CMD [ "node", "src/server.js" ]
