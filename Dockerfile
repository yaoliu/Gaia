FROM alpine:3.10

WORKDIR /app

ADD gaia .

EXPOSE 8000
ENV GOTRACEBACK crash

CMD ["./gaia"]