# Node API Dockerfile
FROM node:18-alpine
WORKDIR /app

# copy package files and install
COPY package.json package-lock.json ./
RUN npm ci --only=production

# copy source
COPY . ./

ENV PORT=3000
ENV SECRET=secret
EXPOSE 3000
CMD ["node", "src/index.js"]
