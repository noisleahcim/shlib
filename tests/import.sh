#!/bin/sh

./shlib import --all
log_info "From import all"

./shlib import --module='logging'
log_info "From import logging module"

./shlib import --module='logging' --function='log_info'
log_info "From import log_info function"
