# Build Stage
# First pull Golang image
FROM node:18-alpine

WORKDIR /app

COPY url-shortener/package*.json ./
RUN npm install

COPY . ./

WORKDIR /app/url-shortener
RUN npm run build

CMD ["npm", "start"]