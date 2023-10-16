# Build NextJS Image
FROM node:lts-alpine3.18 AS npmbase

# Install dependencies only when needed
FROM npmbase AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
WORKDIR /frontend

# Install dependencies based on the preferred package manager
COPY links-admin/package.json links-admin/package-lock.json* ./
RUN npm ci

# Rebuild the source code only when needed
FROM npmbase AS frontendbuilder
WORKDIR /frontend
COPY --from=deps /frontend/node_modules ./node_modules
COPY links-admin/ ./

ENV NEXT_TELEMETRY_DISABLED 1

RUN npm run build

# Build Server Deps
FROM golang:1.21.3 AS SERVERBUILD
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /links

## Final Image
FROM alpine:3.18.4 

COPY --from=SERVERBUILD /links /links
COPY --from=frontendbuilder /frontend/out /frontend
# Run
CMD ["/links"]