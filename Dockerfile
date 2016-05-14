FROM alpine:3.3

WORKDIR /app

ENV VERSION 0.1
ENV BINARY product-service_${VERSION}_linux_amd64


COPY api/seed.json /app/api/seed.json
ADD https://github.com/yunspace/product-service/releases/download/v${VERSION}/${BINARY}.tar.gz /app
RUN tar xzf /app/${BINARY}.tar.gz -C /app/ \
	&& rm /app/${BINARY}.tar.gz

EXPOSE 8080
CMD ["./product-service"]