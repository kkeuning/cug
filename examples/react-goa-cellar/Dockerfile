FROM node:5

# Prepare app directory
 RUN mkdir -p /usr/src/app
 ADD . /usr/src/app
 WORKDIR /usr/src/app
 RUN npm install
 RUN apt-get install curl
# Caddy option for production
 RUN curl https://getcaddy.com | bash
 RUN npm run build
# Default command for webpack hmr server
 CMD node webpack/server-docker.js
#
# Expose the app port
 EXPOSE 3000
