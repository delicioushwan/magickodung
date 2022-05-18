FROM golang:1.17-alpine AS build

WORKDIR /app
COPY . .
RUN go build


FROM alpine:3
COPY --from=build /app/magickodung .
EXPOSE 9091
CMD [ "/magickodung" ]