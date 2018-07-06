#!/bin/sh
TOKEN=my-etcd-token-1
CLUSTER_STATE=new
NAME_1=test-node-1
NAME_2=test-node-2
NAME_3=test-node-3
HOST_1=47.95.242.55
HOST_2=192.168.1.245
HOST_3=192.168.1.248
CLUSTER=${NAME_1}=http://${HOST_1}:2380,${NAME_2}=http://${HOST_2}:2380,${NAME_3}=http://${HOST_3}:2380
CLUSTER_F=${NAME_1}=http://${HOST_1}:2380
case $1 in
        1)
        THIS_NAME=${NAME_1}
        THIS_IP=${HOST_1}
        etcd --data-dir=data.etcd --name ${THIS_NAME} \
                --initial-advertise-peer-urls http://${THIS_IP}:2380 \
                --listen-peer-urls http://${THIS_IP}:2380 \
                --advertise-client-urls http://${THIS_IP}:2379 \
                --listen-client-urls http://${THIS_IP}:2379 \
                --initial-cluster ${CLUSTER} \
                --initial-cluster-state ${CLUSTER_STATE} \
                --initial-cluster-token ${TOKEN}
        ;;

        2)
        # For node 2
        THIS_NAME=${NAME_2}
        THIS_IP=${HOST_2}
        etcd --data-dir=data.etcd --name ${THIS_NAME} \
                --initial-advertise-peer-urls http://${THIS_IP}:2380 \
                --listen-peer-urls http://${THIS_IP}:2380 \
                --advertise-client-urls http://${THIS_IP}:2379 \
                --listen-client-urls http://${THIS_IP}:2379 \
                --initial-cluster ${CLUSTER} \
                --initial-cluster-state ${CLUSTER_STATE} \
                --initial-cluster-token ${TOKEN}
        ;;

        3)
        # For node 3
        THIS_NAME=${NAME_3}
        THIS_IP=${HOST_3}
        etcd --data-dir=data.etcd --name ${THIS_NAME} \
                --initial-advertise-peer-urls http://${THIS_IP}:2380 \
                --listen-peer-urls http://${THIS_IP}:2380 \
                --advertise-client-urls http://${THIS_IP}:2379 \
                --listen-client-urls http://${THIS_IP}:2379 \
                --initial-cluster ${CLUSTER} \
                --initial-cluster-state ${CLUSTER_STATE} \
                --initial-cluster-token ${TOKEN}
                ;;
                
                        *)
        THIS_NAME=${NAME_1}
        THIS_IP=${HOST_1}
        etcd --data-dir=data.etcd --name ${THIS_NAME} \
                --initial-advertise-peer-urls http://${THIS_IP}:2380 \
                --listen-peer-urls http://${THIS_IP}:2380 \
                --advertise-client-urls http://${THIS_IP}:2379 \
                --listen-client-urls http://${THIS_IP}:2379 \
                --initial-cluster ${CLUSTER_F} \
                --initial-cluster-state ${CLUSTER_STATE} \
                --initial-cluster-token ${TOKEN}
        ;;
 esac


