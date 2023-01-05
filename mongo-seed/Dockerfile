FROM mongo

ARG DATABASE_URI

ENV DATABASE_URI ${DATABASE_URI}

COPY ./analytics/consumers/consumers.json /consumers.json
COPY ./analytics/accounts/accounts.json /accounts.json
COPY ./analytics/transactions/transactions.json /transactions.json

CMD mongoimport --uri="mongodb://root:password@mongo_db:27017/analytics?authSource=admin****" --collection consumers --file /consumers.json

# CMD mongoimport --uri=${DATABASE_URI} --collection consumers --drop --jsonArray --file /consumers.json && \
#     mongoimport --uri=${DATABASE_URI} --collection accounts --drop --jsonArray --file /accounts.json && \
#     mongoimport --uri=${DATABASE_URI} --collection transactions --drop --jsonArray --file /transactions.json