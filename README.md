# sablier-sse

See https://github.com/sablierapp/sablier/issues/433 for context

Steps:

* Run `docker compose up -d`
* Check `http://localhost/sse/events`
* Comment the sablier middleware at `traefik/rules/app-sse.yaml` to see SSE working again
