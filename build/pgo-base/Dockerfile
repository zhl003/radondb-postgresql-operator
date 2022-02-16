ARG BASE_IMAGE_OS
ARG DOCKERBASEREGISTRY
FROM ${DOCKERBASEREGISTRY}${BASE_IMAGE_OS}

ARG BASEOS
ARG PGVERSION
ARG PG_FULL
ARG PACKAGER
ARG RELVER


LABEL vendor="RadonDB" \
	url="https://radondb.com" \
	release="${RELVER}" \
	postgresql.version.major="${PGVERSION}" \
	postgresql.version="${PG_FULL}" \
	org.opencontainers.image.vendor="RadonDB" \
	io.openshift.tags="postgresql,postgres,sql,nosql,radondb" \
	io.k8s.description="Trusted open source PostgreSQL-as-a-Service"

COPY licenses /licenses

RUN ${PACKAGER} -y update && ${PACKAGER} -y clean all

RUN if [ "$BASEOS" = "centos8" ]; then \
    ${PACKAGER} -qy module disable postgresql ; \
fi

# Create module file to disable postgres module, microdnf cannot do this with the current version
RUN if [ "$BASEOS" = "ubi8" ] ; then \
    echo "[postgresql]"  >> /etc/dnf/modules.d/postgresql.module \
    && echo "name=postgresql"  >> /etc/dnf/modules.d/postgresql.module \
    && echo "stream=10"  >> /etc/dnf/modules.d/postgresql.module \
    && echo "profiles="  >> /etc/dnf/modules.d/postgresql.module \
    && echo "state=disabled"  >> /etc/dnf/modules.d/postgresql.module ; \
fi

# RadonDB PostgreSQL repository
ADD conf/radondbpg${PGVERSION}.repo /etc/yum.repos.d/
