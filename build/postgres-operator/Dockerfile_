ARG BASEOS
ARG BASEVER
ARG PREFIX
FROM ${PREFIX}/pgo-base:${BASEOS}-${BASEVER}

LABEL name="postgres-operator" \
	summary="RadonDB PostgreSQL Operator" \
	description="RadonDB PostgreSQL Operator"

ADD bin/postgres-operator /usr/local/bin

USER 2

CMD ["postgres-operator"]
