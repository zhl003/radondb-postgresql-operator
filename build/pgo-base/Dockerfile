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
ARG EPEL7_RPM=https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
ARG EPEL8_RPM=https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm
ENV TZ=Asia/Shanghai
COPY conf/pip.conf pip.conf
ENV PIP_CONFIG_FILE pip.conf
RUN sed -e "s|^mirrorlist=|#mirrorlist=|g" \
         -e "s|^#baseurl=http://mirror.centos.org/\$contentdir/\$releasever|baseurl=https://mirrors.tuna.tsinghua.edu.cn/centos-vault/8.5.2111|g" \
         -i.bak \
         /etc/yum.repos.d/CentOS-*.repo
COPY licenses /licenses

RUN ${PACKAGER} -y update && ${PACKAGER} -y clean all

RUN if [ "$BASEOS" = "centos8" ]; then \
    ${PACKAGER} -qy module disable postgresql ; \
fi

# RadonDB PostgreSQL repository
ADD conf/radondbpg${PGVERSION}.repo /etc/yum.repos.d/
