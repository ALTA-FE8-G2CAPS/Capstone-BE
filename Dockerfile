FROM golang:1.18

# membuat direktori app
RUN mkdir /app

# set workdir /app
WORKDIR /app
# copy semua file ke app
COPY ./ /app

RUN go mod tidy

RUN go build -o Capstone-BE

CMD ("./Capstone-BE")