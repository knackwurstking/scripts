#!/bin/bash

LOG_DIR=${HOME}/.log
[ ! -e ${LOG_DIR} ] && mkdir ${LOG_DIR}

LOG=${LOG_DIR}/backup-documents.log

ROOT=/mnt/media
SRC=${ROOT}/Data/Documents
BACKUP_PATH=${ROOT}/backups
DST=${BACKUP_PATH}/Data/Documents

echo "[INFO] [$(date)] Backup ${SRC} --> ${DST}" >> ${LOG}

tar --create \
    --gzip \
    --preserve-permissions \
    --file ${DST}/backup-$(date +%Y-%m-%d).tar.gz \
    ${SRC} 2>> ${LOG}

exit 0
