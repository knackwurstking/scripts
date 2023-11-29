#!/bin/bash

USER=git
BACKUP_LOCATION=/mnt/media/backups
BACKUP_LOG=$HOME/.backup.log

echo -e "\033[1;33mRunning backup for user \033[1;32m$USER\033[1;33m on \033[1;32m$(date)\033[0m" >> $BACKUP_LOG

tar --create --gzip --preserve-permissions \
  --file ${BACKUP_LOCATION}/${USER}/backup-$(date +%Y-%m-%d).tar.gz \
  ${HOME} 2>> $BACKUP_LOG

exit 0
