FROM node:16-alpine

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY . .

RUN \
    if [ -f yarn.lock ]; then yarn && yarn build; \
    elif [ -f package-lock.json ]; then npm ci && npm run build; \
    elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm i && pnpm run build; \
    else echo "Lockfile not found." && exit 1; \
    fi

EXPOSE 30003
CMD npm run start