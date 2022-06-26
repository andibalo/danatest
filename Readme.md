## Getting Started

To start the application locally, run the following commands:

```bash
cd ./infra
docker-compose -f docker-compose.dev.yaml up -d  
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Reason for using SQL database

I use SQL because of the extensive support with GORM library as well as the atomicity attribute
