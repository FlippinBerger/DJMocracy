# ARG NODE_VERSION=18.17.0
#
# FROM node:${NODE_VERSION}-slim as base
#
# ENV NODE_ENV=development
#
# WORKDIR /src
#
# # Build
# FROM base as build
#
# COPY --link package.json yarn.lock ./
# RUN yarn install
#
# # Run
# FROM base
#
# COPY --from=build /src/node_modules /src/node_modules
#
# CMD [ "yarn", "run", "dev" ]

FROM node:20.0.0

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install

COPY . .

EXPOSE 3000

# RUN yarn run build
CMD [ "yarn", "run", "dev" ]
