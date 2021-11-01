up:
	cd deployments/docker-compose && docker-compose up -d

ps:
	cd deployments/docker-compose && docker-compose ps

stop:
	cd deployments/docker-compose && docker-compose stop

exec:
	cd deployments/docker-compose && docker-compose exec webhook bash
