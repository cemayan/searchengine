#Stage 1
FROM node:20.15.1-alpine as builder
WORKDIR /app
COPY web/package*.json .
COPY web/yarn*.lock .
RUN yarn install
COPY web .
ENV VITE_READAPI=http://readapi.local
ENV VITE_WRITEAPI=http://writeapi.local
RUN  yarn build

#Stage 2
FROM nginx:1.27-alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/dist .
ENTRYPOINT ["nginx", "-g", "daemon off;"]