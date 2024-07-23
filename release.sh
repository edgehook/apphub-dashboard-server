#!/bin/bash

Version=$1
#Version=${Version:-0.2.0}

TARGET_ROOT=`cd "$(dirname "$0")"; pwd`
RELEASE_NAME=release_$Version
RELEASE_ROOT=${TARGET_ROOT}/${RELEASE_NAME}

echo RELEASE_NAME=$RELEASE_NAME

cd ${TARGET_ROOT}/

echo "build x86_64 dashboard... "
make dashboard
echo "build arm64 dashboard... "
make dashboard-ARM64
echo "build arm dashboard... "
make dashboard-ARM
echo "build windows exe dashboard"
make dashboard.exe

if [ ! -d ${RELEASE_ROOT} ]; then
	rm -rf ${RELEASE_ROOT}		
fi

[ ! -e ${RELEASE_ROOT} ] && mkdir -p ${RELEASE_ROOT}

cp -a frontend ${RELEASE_ROOT}/
cp -a conf ${RELEASE_ROOT}/
cp  -a dashboard dashboard-ARM dashboard-ARM64 dashboard.exe ${RELEASE_ROOT}/
[ -a "startup.sh" ] && cp -a startup.sh ${RELEASE_ROOT}/
[ -a "dashboard.service" ] && cp -a dashboard.service  ${RELEASE_ROOT}/dashboard.service

echo "Syncing...."
sync;sync;
sync
sync

#tar cJvf ${RELEASE_NAME}.tar.xz ${RELEASE_NAME}
tar zcvf ${RELEASE_NAME}.tar.gz ${RELEASE_NAME}

rm -rf ${RELEASE_ROOT}
echo "[Done]"

