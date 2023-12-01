#!/bin/bash

COLOR_RESET="\033[0m"
COLOR_BOLD_YELLOW="\033[1;33m"
COLOR_BOLD_GREEN="\033[1;32m"

USER=git

DST=/mnt/media/backups

LOG_DIR=${HOME}/.log
LOG=${LOG_DIR}/backup.log

[ ! -e ${LOG_DIR} ] && mkdir ${LOG_DIR}

echo -e \
    "${COLOR_BOLD_YELLOW}Running backup for user " \
    "${COLOR_BOLD_GREEN}${USER}" \
    "${COLOR_BOLD_YELLOW} on " \
    "${COLOR_BOLD_GREEN}$(date)${COLOR_RESET}" >> ${LOG}

tar --create \
    --gzip \
    --preserve-permissions \
    --file ${DST}/${USER}/backup-$(date +%Y-%m-%d).tar.gz \
    ${HOME} 2>> ${LOG}

exit 0
