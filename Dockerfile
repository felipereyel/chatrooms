# Build the Go binary
FROM golang:1.20-alpine AS goapp
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go  .
COPY gosrc/ gosrc/
RUN go build -o ./goapp

# Build the Vue app
FROM node:18-alpine as vueapp
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY ./src ./src
COPY index.html tsconfig.json tsconfig.node.json vite.config.ts ./
RUN npm run build

# Build the final image
FROM alpine:latest as release
WORKDIR /app

COPY --from=goapp /app/goapp /goapp
COPY --from=vueapp /app/dist /dist

ENV PUBLIC_DIR /dist
CMD ["/goapp", "serve"]