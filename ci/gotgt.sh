#!/bin/bash
#set -x

source ./start_init_test.sh
test_ctrl_stop_start
test_replica_controller_continuous_stop_start
run_data_integrity_test_with_fs_creation
