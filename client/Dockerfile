# pull official base image
FROM node:15.14.0-alpine3.10

# set working directory
WORKDIR /app

RUN npm install -g create-react-app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY yarn.lock ./
COPY ./public ./public
COPY ./src ./src
RUN yarn install --

EXPOSE 3000

# start app
CMD ["yarn", "start"]
