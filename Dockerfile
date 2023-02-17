FROM golang:latest

# kopieer de huidige directory naar de container
COPY . /app

# maak de /app directory de huidige directory
WORKDIR /app

# download en installeer de benodigde pakketten
RUN go get -d -v

# bouw het Go-programma
RUN go build -o main .

# stel de poort in waarop de container moet luisteren
EXPOSE 8080

# stel het commando in dat wordt uitgevoerd wanneer de container wordt gestart
CMD ["/app/main"]